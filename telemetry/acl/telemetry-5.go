package acl

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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_MplsPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_MplsPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Mpls {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_MplsPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Mpls
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_MplsPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Mpls {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Mpls
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_MplsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Mpls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Mpls) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry_Mpls{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry_Mpls)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_MplsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls) bool) *oc.Acl_AclSet_AclEntry_MplsWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_MplsWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", gs, queryPath, false, false)
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Mpls{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Mpls)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_MplsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls) bool) *oc.Acl_AclSet_AclEntry_MplsWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_MplsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_MplsPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Mpls) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls to the batch object.
func (n *Acl_AclSet_AclEntry_MplsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_MplsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Mpls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Mpls) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_MplsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls) bool) *oc.Acl_AclSet_AclEntry_MplsWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_MplsWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedAcl_AclSet_AclEntry_Mpls{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Mpls)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_MplsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls) bool) *oc.Acl_AclSet_AclEntry_MplsWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_MplsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls to the batch object.
func (n *Acl_AclSet_AclEntry_MplsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_EndLabelValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_EndLabelValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Await(t testing.TB, timeout time.Duration, val oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/end-label-value to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_EndLabelValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath extracts the value of the leaf EndLabelValue from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union.
func convertAcl_AclSet_AclEntry_Mpls_EndLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union{
		Metadata: md,
	}
	val := parent.EndLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Get(t testing.TB) oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny) Get(t testing.TB) []oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_StartLabelValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_StartLabelValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Await(t testing.TB, timeout time.Duration, val oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) bool) *oc.Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/start-label-value to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_StartLabelValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath extracts the value of the leaf StartLabelValue from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union.
func convertAcl_AclSet_AclEntry_Mpls_StartLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union{
		Metadata: md,
	}
	val := parent.StartLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_TrafficClassPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_TrafficClassPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/traffic-class to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TrafficClassPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath extracts the value of the leaf TrafficClass from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Mpls_TrafficClassPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Mpls", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_TtlValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_TtlValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_Mpls_TtlValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Mpls{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Mpls", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_Mpls_TtlValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/mpls/state/ttl-value to the batch object.
func (n *Acl_AclSet_AclEntry_Mpls_TtlValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_Mpls_TtlValuePath extracts the value of the leaf TtlValue from its parent oc.Acl_AclSet_AclEntry_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertAcl_AclSet_AclEntry_Mpls_TtlValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Mpls) *oc.QualifiedUint8 {
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

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, true, false)
	if ok {
		return convertAcl_AclSet_AclEntry_SequenceIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_SequenceIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_SequenceIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_SequenceIdPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_SequenceIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertAcl_AclSet_AclEntry_SequenceIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_SequenceIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id to the batch object.
func (n *Acl_AclSet_AclEntry_SequenceIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_SequenceIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_SequenceIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry", structs[pre], queryPath, true, false)
			qv := convertAcl_AclSet_AclEntry_SequenceIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_SequenceIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_SequenceIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/state/sequence-id to the batch object.
func (n *Acl_AclSet_AclEntry_SequenceIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_AclSet_AclEntry_SequenceIdPath extracts the value of the leaf SequenceId from its parent oc.Acl_AclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_AclSet_AclEntry_SequenceIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SequenceId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_TransportPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Transport {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Transport", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Transport{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_TransportPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Transport {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_TransportPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Transport {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Transport
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Transport{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_TransportPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Transport {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Transport
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_TransportPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Transport {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Transport{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Transport) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_AclSet_AclEntry_Transport{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_AclSet_AclEntry_Transport)))
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_TransportPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport) bool) *oc.Acl_AclSet_AclEntry_TransportWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_TransportWatcher{}
	gs := &oc.Acl_AclSet_AclEntry_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", gs, queryPath, false, false)
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Transport{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Transport)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_TransportPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport) bool) *oc.Acl_AclSet_AclEntry_TransportWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_TransportPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_AclSet_AclEntry_TransportPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_AclSet_AclEntry_Transport) *oc.QualifiedAcl_AclSet_AclEntry_Transport {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_AclSet_AclEntry_Transport) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport to the batch object.
func (n *Acl_AclSet_AclEntry_TransportPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_AclSet_AclEntry_TransportPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_AclSet_AclEntry_Transport {
	t.Helper()
	c := &oc.CollectionAcl_AclSet_AclEntry_Transport{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_AclSet_AclEntry_Transport) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_AclSet_AclEntry_TransportPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport) bool) *oc.Acl_AclSet_AclEntry_TransportWatcher {
	t.Helper()
	w := &oc.Acl_AclSet_AclEntry_TransportWatcher{}
	structs := map[string]*oc.Acl_AclSet_AclEntry_Transport{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Acl_AclSet_AclEntry_Transport{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Acl_AclSet_AclEntry_Transport", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedAcl_AclSet_AclEntry_Transport{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_AclSet_AclEntry_Transport)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_AclSet_AclEntry_TransportPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_AclSet_AclEntry_Transport) bool) *oc.Acl_AclSet_AclEntry_TransportWatcher {
	t.Helper()
	return watch_Acl_AclSet_AclEntry_TransportPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/transport to the batch object.
func (n *Acl_AclSet_AclEntry_TransportPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
