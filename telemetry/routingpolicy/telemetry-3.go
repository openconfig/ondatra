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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/asn to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath extracts the value of the leaf Asn from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Asn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/state/repeat-n to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath extracts the value of the leaf RepeatN from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.RepeatN
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Await(t testing.TB, timeout time.Duration, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/state/communities to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath extracts the value of the leaf Communities from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice{
		Metadata: md,
	}
	val := parent.Communities
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Lookup(t testing.TB) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Get(t testing.TB) oc.E_SetCommunity_Method {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SetCommunity_Method
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny) Get(t testing.TB) []oc.E_SetCommunity_Method {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SetCommunity_Method
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SetCommunity_Method {
	t.Helper()
	c := &oc.CollectionE_SetCommunity_Method{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SetCommunity_Method) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	w := &oc.E_SetCommunity_MethodWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SetCommunity_Method)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Await(t testing.TB, timeout time.Duration, val oc.E_SetCommunity_Method) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SetCommunity_Method) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SetCommunity_Method {
	t.Helper()
	c := &oc.CollectionE_SetCommunity_Method{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SetCommunity_Method) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	w := &oc.E_SetCommunity_MethodWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SetCommunity_Method)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/method to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath extracts the value of the leaf Method from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SetCommunity_Method.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	qv := &oc.QualifiedE_SetCommunity_Method{
		Metadata: md,
	}
	val := parent.Method
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Lookup(t testing.TB) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Get(t testing.TB) oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny) Get(t testing.TB) []oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpPolicy_BgpSetCommunityOptionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	c := &oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	w := &oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Await(t testing.TB, timeout time.Duration, val oc.E_BgpPolicy_BgpSetCommunityOptionType) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	c := &oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	w := &oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/state/options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath extracts the value of the leaf Options from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	qv := &oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType{
		Metadata: md,
	}
	val := parent.Options
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
