package routingpolicy

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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath) Lookup(t testing.TB) *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath) Get(t testing.TB) []oc.E_BgpTypes_AFI_SAFI_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPathAny) Get(t testing.TB) [][]oc.E_BgpTypes_AFI_SAFI_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_BgpTypes_AFI_SAFI_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	c := &oc.CollectionE_BgpTypes_AFI_SAFI_TYPESlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) bool) *oc.E_BgpTypes_AFI_SAFI_TYPESliceWatcher {
	t.Helper()
	w := &oc.E_BgpTypes_AFI_SAFI_TYPESliceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) bool) *oc.E_BgpTypes_AFI_SAFI_TYPESliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath) Await(t testing.TB, timeout time.Duration, val []oc.E_BgpTypes_AFI_SAFI_TYPE) *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	c := &oc.CollectionE_BgpTypes_AFI_SAFI_TYPESlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) bool) *oc.E_BgpTypes_AFI_SAFI_TYPESliceWatcher {
	t.Helper()
	w := &oc.E_BgpTypes_AFI_SAFI_TYPESliceWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) bool) *oc.E_BgpTypes_AFI_SAFI_TYPESliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/afi-safi-in to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath extracts the value of the leaf AfiSafiIn from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AfiSafiInPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) *oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	qv := &oc.QualifiedE_BgpTypes_AFI_SAFI_TYPESlice{
		Metadata: md,
	}
	val := parent.AfiSafiIn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath) Get(t testing.TB) oc.E_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPathAny) Get(t testing.TB) []oc.E_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_ATTRIBUTE_COMPARISON
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath) Await(t testing.TB, timeout time.Duration, val oc.E_PolicyTypes_ATTRIBUTE_COMPARISON) *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/operator to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath extracts the value of the leaf Operator from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_OperatorPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON{
		Metadata: md,
	}
	val := parent.Operator
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/as-path-length/state/value to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath extracts the value of the leaf Value from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength_ValuePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Value
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath) Get(t testing.TB) oc.E_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPathAny) Get(t testing.TB) []oc.E_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_ATTRIBUTE_COMPARISON
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath) Await(t testing.TB, timeout time.Duration, val oc.E_PolicyTypes_ATTRIBUTE_COMPARISON) *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) bool) *oc.E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/operator to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath extracts the value of the leaf Operator from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_OperatorPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) *oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON{
		Metadata: md,
	}
	val := parent.Operator
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/community-count/state/value to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath extracts the value of the leaf Value from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount_ValuePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Value
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/community-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath extracts the value of the leaf CommunitySet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunitySetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CommunitySet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/state/ext-community-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath extracts the value of the leaf ExtCommunitySet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_ExtCommunitySetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ExtCommunitySet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
