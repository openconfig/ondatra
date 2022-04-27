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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Mpls
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions_Mpls)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_MplsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool) *oc.Qos_Classifier_Term_Conditions_MplsWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_MplsWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool) *oc.Qos_Classifier_Term_Conditions_MplsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_MplsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls to the batch object.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_MplsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool) *oc.Qos_Classifier_Term_Conditions_MplsWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_MplsWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool) *oc.Qos_Classifier_Term_Conditions_MplsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_MplsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls to the batch object.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Await(t testing.TB, timeout time.Duration, val oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath extracts the value of the leaf EndLabelValue from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union.
func convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union{
		Metadata: md,
	}
	val := parent.EndLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Await(t testing.TB, timeout time.Duration, val oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/start-label-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath extracts the value of the leaf StartLabelValue from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union.
func convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union{
		Metadata: md,
	}
	val := parent.StartLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/traffic-class to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath extracts the value of the leaf TrafficClass from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.TrafficClass
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_TtlValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_TtlValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/ttl-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath extracts the value of the leaf TtlValue from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.TtlValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Transport{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Transport {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_TransportPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Transport
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Transport{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_TransportPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Transport {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Transport
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Transport{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Transport) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions_Transport{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions_Transport)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_TransportPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport) bool) *oc.Qos_Classifier_Term_Conditions_TransportWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_TransportWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Transport{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Transport)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport) bool) *oc.Qos_Classifier_Term_Conditions_TransportWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_TransportPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions_Transport) *oc.QualifiedQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Transport) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport to the batch object.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_TransportPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Transport{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Transport) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_TransportPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport) bool) *oc.Qos_Classifier_Term_Conditions_TransportWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_TransportWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Transport{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Transport{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Transport)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_TransportPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport) bool) *oc.Qos_Classifier_Term_Conditions_TransportWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_TransportPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport to the batch object.
func (n *Qos_Classifier_Term_Conditions_TransportPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Transport_DestinationPort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Transport_DestinationPortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Transport_DestinationPortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Await(t testing.TB, timeout time.Duration, val oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port to the batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Transport_DestinationPort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Transport{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/destination-port to the batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath extracts the value of the leaf DestinationPort from its parent oc.Qos_Classifier_Term_Conditions_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union.
func convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Transport) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union{
		Metadata: md,
	}
	val := parent.DestinationPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Transport_SourcePort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Transport_SourcePortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Transport_SourcePortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Await(t testing.TB, timeout time.Duration, val oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port to the batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Transport_SourcePort_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Transport{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) bool) *oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/source-port to the batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Transport_SourcePortPath extracts the value of the leaf SourcePort from its parent oc.Qos_Classifier_Term_Conditions_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union.
func convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Transport) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union{
		Metadata: md,
	}
	val := parent.SourcePort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Lookup(t testing.TB) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Get(t testing.TB) []oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny) Get(t testing.TB) [][]oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_PacketMatchTypes_TCP_FLAGS
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	c := &oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool) *oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher {
	t.Helper()
	w := &oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool) *oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Await(t testing.TB, timeout time.Duration, val []oc.E_PacketMatchTypes_TCP_FLAGS) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags to the batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	c := &oc.CollectionE_PacketMatchTypes_TCP_FLAGSSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool) *oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher {
	t.Helper()
	w := &oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term_Conditions_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term_Conditions_Transport{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) bool) *oc.E_PacketMatchTypes_TCP_FLAGSSliceWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/state/tcp-flags to the batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath extracts the value of the leaf TcpFlags from its parent oc.Qos_Classifier_Term_Conditions_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice.
func convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Transport) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	qv := &oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice{
		Metadata: md,
	}
	val := parent.TcpFlags
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a ONCE subscription.
func (n *Qos_Classifier_Term_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_Term_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/state/id to the batch object.
func (n *Qos_Classifier_Term_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Classifier_Term{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier_Term{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier_Term", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_Term_IdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/state/id to the batch object.
func (n *Qos_Classifier_Term_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_IdPath extracts the value of the leaf Id from its parent oc.Qos_Classifier_Term
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Qos_Classifier_Type {
	t.Helper()
	goStruct := &oc.Qos_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier", goStruct, true, false)
	if ok {
		return convertQos_Classifier_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/state/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_TypePath) Get(t testing.TB) oc.E_Qos_Classifier_Type {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Qos_Classifier_Type {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Qos_Classifier_Type
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/state/type with a ONCE subscription.
func (n *Qos_Classifier_TypePathAny) Get(t testing.TB) []oc.E_Qos_Classifier_Type {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Qos_Classifier_Type
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Qos_Classifier_Type {
	t.Helper()
	c := &oc.CollectionE_Qos_Classifier_Type{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Qos_Classifier_Type) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Qos_Classifier_Type) bool) *oc.E_Qos_Classifier_TypeWatcher {
	t.Helper()
	w := &oc.E_Qos_Classifier_TypeWatcher{}
	gs := &oc.Qos_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Classifier_TypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Qos_Classifier_Type)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Qos_Classifier_Type) bool) *oc.E_Qos_Classifier_TypeWatcher {
	t.Helper()
	return watch_Qos_Classifier_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_TypePath) Await(t testing.TB, timeout time.Duration, val oc.E_Qos_Classifier_Type) *oc.QualifiedE_Qos_Classifier_Type {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Qos_Classifier_Type) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/state/type to the batch object.
func (n *Qos_Classifier_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Qos_Classifier_Type {
	t.Helper()
	c := &oc.CollectionE_Qos_Classifier_Type{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Qos_Classifier_Type) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_TypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Qos_Classifier_Type) bool) *oc.E_Qos_Classifier_TypeWatcher {
	t.Helper()
	w := &oc.E_Qos_Classifier_TypeWatcher{}
	structs := map[string]*oc.Qos_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Classifier{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Classifier", structs[pre], queryPath, true, false)
			qv := convertQos_Classifier_TypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Qos_Classifier_Type)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Qos_Classifier_Type) bool) *oc.E_Qos_Classifier_TypeWatcher {
	t.Helper()
	return watch_Qos_Classifier_TypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/state/type to the batch object.
func (n *Qos_Classifier_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_TypePath extracts the value of the leaf Type from its parent oc.Qos_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Qos_Classifier_Type.
func convertQos_Classifier_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier) *oc.QualifiedE_Qos_Classifier_Type {
	t.Helper()
	qv := &oc.QualifiedE_Qos_Classifier_Type{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroupPath) Lookup(t testing.TB) *oc.QualifiedQos_ForwardingGroup {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_ForwardingGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroupPath) Get(t testing.TB) *oc.Qos_ForwardingGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_ForwardingGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_ForwardingGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_ForwardingGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription.
func (n *Qos_ForwardingGroupPathAny) Get(t testing.TB) []*oc.Qos_ForwardingGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_ForwardingGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_ForwardingGroup {
	t.Helper()
	c := &oc.CollectionQos_ForwardingGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_ForwardingGroup) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_ForwardingGroup{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_ForwardingGroup)))
		return false
	})
	return c
}

func watch_Qos_ForwardingGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_ForwardingGroup) bool) *oc.Qos_ForwardingGroupWatcher {
	t.Helper()
	w := &oc.Qos_ForwardingGroupWatcher{}
	gs := &oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_ForwardingGroup", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_ForwardingGroup{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_ForwardingGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_ForwardingGroup) bool) *oc.Qos_ForwardingGroupWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_ForwardingGroupPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_ForwardingGroup) *oc.QualifiedQos_ForwardingGroup {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_ForwardingGroup) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/forwarding-groups/forwarding-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group to the batch object.
func (n *Qos_ForwardingGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_ForwardingGroup {
	t.Helper()
	c := &oc.CollectionQos_ForwardingGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_ForwardingGroup) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroupPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_ForwardingGroup) bool) *oc.Qos_ForwardingGroupWatcher {
	t.Helper()
	w := &oc.Qos_ForwardingGroupWatcher{}
	structs := map[string]*oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_ForwardingGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_ForwardingGroup", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_ForwardingGroup{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_ForwardingGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_ForwardingGroup) bool) *oc.Qos_ForwardingGroupWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroupPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group to the batch object.
func (n *Qos_ForwardingGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, false)
	if ok {
		return convertQos_ForwardingGroup_FabricPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_FabricPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_FabricPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a ONCE subscription.
func (n *Qos_ForwardingGroup_FabricPriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_FabricPriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_ForwardingGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_ForwardingGroup_FabricPriorityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_FabricPriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority to the batch object.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_FabricPriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_FabricPriorityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_ForwardingGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_ForwardingGroup", structs[pre], queryPath, true, false)
			qv := convertQos_ForwardingGroup_FabricPriorityPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_FabricPriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_FabricPriorityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/fabric-priority to the batch object.
func (n *Qos_ForwardingGroup_FabricPriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_ForwardingGroup_FabricPriorityPath extracts the value of the leaf FabricPriority from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_ForwardingGroup_FabricPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.FabricPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, false)
	if ok {
		return convertQos_ForwardingGroup_MulticastOutputQueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_MulticastOutputQueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a ONCE subscription.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_MulticastOutputQueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_ForwardingGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_ForwardingGroup_MulticastOutputQueuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_MulticastOutputQueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue to the batch object.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_MulticastOutputQueuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_ForwardingGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_ForwardingGroup", structs[pre], queryPath, true, false)
			qv := convertQos_ForwardingGroup_MulticastOutputQueuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_MulticastOutputQueuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/multicast-output-queue to the batch object.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_ForwardingGroup_MulticastOutputQueuePath extracts the value of the leaf MulticastOutputQueue from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_MulticastOutputQueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MulticastOutputQueue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, false)
	if ok {
		return convertQos_ForwardingGroup_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a ONCE subscription.
func (n *Qos_ForwardingGroup_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_ForwardingGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_ForwardingGroup_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_ForwardingGroup_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name to the batch object.
func (n *Qos_ForwardingGroup_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_ForwardingGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_ForwardingGroup", structs[pre], queryPath, true, false)
			qv := convertQos_ForwardingGroup_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/name to the batch object.
func (n *Qos_ForwardingGroup_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_ForwardingGroup_NamePath extracts the value of the leaf Name from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_OutputQueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, false)
	if ok {
		return convertQos_ForwardingGroup_OutputQueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_OutputQueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_OutputQueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_OutputQueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a ONCE subscription.
func (n *Qos_ForwardingGroup_OutputQueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_OutputQueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_OutputQueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_ForwardingGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_ForwardingGroup_OutputQueuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_OutputQueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_OutputQueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_ForwardingGroup_OutputQueuePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue to the batch object.
func (n *Qos_ForwardingGroup_OutputQueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_OutputQueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_OutputQueuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_ForwardingGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_ForwardingGroup", structs[pre], queryPath, true, false)
			qv := convertQos_ForwardingGroup_OutputQueuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_OutputQueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_OutputQueuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/output-queue to the batch object.
func (n *Qos_ForwardingGroup_OutputQueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_ForwardingGroup_OutputQueuePath extracts the value of the leaf OutputQueue from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_OutputQueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OutputQueue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, false)
	if ok {
		return convertQos_ForwardingGroup_UnicastOutputQueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_UnicastOutputQueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a ONCE subscription.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_UnicastOutputQueuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_ForwardingGroup", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_ForwardingGroup_UnicastOutputQueuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_UnicastOutputQueuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue to the batch object.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_ForwardingGroup_UnicastOutputQueuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_ForwardingGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_ForwardingGroup{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_ForwardingGroup", structs[pre], queryPath, true, false)
			qv := convertQos_ForwardingGroup_UnicastOutputQueuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_ForwardingGroup_UnicastOutputQueuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/forwarding-groups/forwarding-group/state/unicast-output-queue to the batch object.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_ForwardingGroup_UnicastOutputQueuePath extracts the value of the leaf UnicastOutputQueue from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_UnicastOutputQueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.UnicastOutputQueue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_InterfacePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface {
	t.Helper()
	goStruct := &oc.Qos_Interface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_InterfacePath) Get(t testing.TB) *oc.Qos_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface with a ONCE subscription.
func (n *Qos_InterfacePathAny) Get(t testing.TB) []*oc.Qos_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface {
	t.Helper()
	c := &oc.CollectionQos_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface)))
		return false
	})
	return c
}

func watch_Qos_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface) bool) *oc.Qos_InterfaceWatcher {
	t.Helper()
	w := &oc.Qos_InterfaceWatcher{}
	gs := &oc.Qos_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface) bool) *oc.Qos_InterfaceWatcher {
	t.Helper()
	return watch_Qos_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_InterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface) *oc.QualifiedQos_Interface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface to the batch object.
func (n *Qos_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface {
	t.Helper()
	c := &oc.CollectionQos_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface) bool) *oc.Qos_InterfaceWatcher {
	t.Helper()
	w := &oc.Qos_InterfaceWatcher{}
	structs := map[string]*oc.Qos_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface) bool) *oc.Qos_InterfaceWatcher {
	t.Helper()
	return watch_Qos_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface to the batch object.
func (n *Qos_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InputPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InputPath) Get(t testing.TB) *oc.Qos_Interface_Input {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription.
func (n *Qos_Interface_InputPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InputPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input)))
		return false
	})
	return c
}

func watch_Qos_Interface_InputPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input) bool) *oc.Qos_Interface_InputWatcher {
	t.Helper()
	w := &oc.Qos_Interface_InputWatcher{}
	gs := &oc.Qos_Interface_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface_Input{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InputPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input) bool) *oc.Qos_Interface_InputWatcher {
	t.Helper()
	return watch_Qos_Interface_InputPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_InputPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input) *oc.QualifiedQos_Interface_Input {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input to the batch object.
func (n *Qos_Interface_InputPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_InputPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_InputPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input) bool) *oc.Qos_Interface_InputWatcher {
	t.Helper()
	w := &oc.Qos_Interface_InputWatcher{}
	structs := map[string]*oc.Qos_Interface_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface_Input{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_InputPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input) bool) *oc.Qos_Interface_InputWatcher {
	t.Helper()
	return watch_Qos_Interface_InputPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input to the batch object.
func (n *Qos_Interface_InputPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, true, false)
	if ok {
		return convertQos_Interface_Input_BufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_BufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertQos_Interface_Input_BufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Input_BufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_BufferAllocationProfilePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Interface_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertQos_Interface_Input_BufferAllocationProfilePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_BufferAllocationProfilePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_BufferAllocationProfilePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_BufferAllocationProfilePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Qos_Interface_Input{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input", structs[pre], queryPath, true, false)
			qv := convertQos_Interface_Input_BufferAllocationProfilePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_BufferAllocationProfilePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_BufferAllocationProfilePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/state/buffer-allocation-profile to the batch object.
func (n *Qos_Interface_Input_BufferAllocationProfilePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Interface_Input_BufferAllocationProfilePath extracts the value of the leaf BufferAllocationProfile from its parent oc.Qos_Interface_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_BufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.BufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_ClassifierPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_Classifier {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_ClassifierPath) Get(t testing.TB) *oc.Qos_Interface_Input_Classifier {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_ClassifierPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_Classifier {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_Classifier
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription.
func (n *Qos_Interface_Input_ClassifierPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_Classifier {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_Classifier
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_ClassifierPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_Classifier {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_Classifier{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_Classifier) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Interface_Input_Classifier{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Interface_Input_Classifier)))
		return false
	})
	return c
}

func watch_Qos_Interface_Input_ClassifierPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Classifier) bool) *oc.Qos_Interface_Input_ClassifierWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_ClassifierWatcher{}
	gs := &oc.Qos_Interface_Input_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Interface_Input_Classifier", gs, queryPath, false, false)
		qv := (&oc.QualifiedQos_Interface_Input_Classifier{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_Classifier)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_ClassifierPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Classifier) bool) *oc.Qos_Interface_Input_ClassifierWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_ClassifierPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Interface_Input_ClassifierPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Interface_Input_Classifier) *oc.QualifiedQos_Interface_Input_Classifier {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Interface_Input_Classifier) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier to the batch object.
func (n *Qos_Interface_Input_ClassifierPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Interface_Input_ClassifierPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Interface_Input_Classifier {
	t.Helper()
	c := &oc.CollectionQos_Interface_Input_Classifier{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Interface_Input_Classifier) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Interface_Input_ClassifierPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Classifier) bool) *oc.Qos_Interface_Input_ClassifierWatcher {
	t.Helper()
	w := &oc.Qos_Interface_Input_ClassifierWatcher{}
	structs := map[string]*oc.Qos_Interface_Input_Classifier{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Qos_Interface_Input_Classifier{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Qos_Interface_Input_Classifier", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedQos_Interface_Input_Classifier{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Interface_Input_Classifier)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Interface_Input_ClassifierPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Interface_Input_Classifier) bool) *oc.Qos_Interface_Input_ClassifierWatcher {
	t.Helper()
	return watch_Qos_Interface_Input_ClassifierPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier to the batch object.
func (n *Qos_Interface_Input_ClassifierPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
