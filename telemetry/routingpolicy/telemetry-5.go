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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Get(t testing.TB) oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPathAny) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", gs, queryPath, true, false)
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Await(t testing.TB, timeout time.Duration, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-next-hop to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath extracts the value of the leaf SetNextHop from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union{
		Metadata: md,
	}
	val := parent.SetNextHop
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Lookup(t testing.TB) *oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Get(t testing.TB) oc.E_BgpTypes_BgpOriginAttrType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpTypes_BgpOriginAttrType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPathAny) Get(t testing.TB) []oc.E_BgpTypes_BgpOriginAttrType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpTypes_BgpOriginAttrType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	c := &oc.CollectionE_BgpTypes_BgpOriginAttrType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpTypes_BgpOriginAttrType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpTypes_BgpOriginAttrType) bool) *oc.E_BgpTypes_BgpOriginAttrTypeWatcher {
	t.Helper()
	w := &oc.E_BgpTypes_BgpOriginAttrTypeWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", gs, queryPath, true, false)
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpTypes_BgpOriginAttrType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpTypes_BgpOriginAttrType) bool) *oc.E_BgpTypes_BgpOriginAttrTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_BgpTypes_BgpOriginAttrType) *oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_BgpTypes_BgpOriginAttrType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	c := &oc.CollectionE_BgpTypes_BgpOriginAttrType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpTypes_BgpOriginAttrType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpTypes_BgpOriginAttrType) bool) *oc.E_BgpTypes_BgpOriginAttrTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-route-origin to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath extracts the value of the leaf SetRouteOrigin from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpTypes_BgpOriginAttrType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	qv := &oc.QualifiedE_BgpTypes_BgpOriginAttrType{
		Metadata: md,
	}
	val := parent.SetRouteOrigin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Lookup(t testing.TB) *oc.QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Get(t testing.TB) oc.E_RoutingPolicy_PolicyResultType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPathAny) Lookup(t testing.TB) []*oc.QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_RoutingPolicy_PolicyResultType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPathAny) Get(t testing.TB) []oc.E_RoutingPolicy_PolicyResultType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_RoutingPolicy_PolicyResultType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_RoutingPolicy_PolicyResultType {
	t.Helper()
	c := &oc.CollectionE_RoutingPolicy_PolicyResultType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_RoutingPolicy_PolicyResultType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_RoutingPolicy_PolicyResultType) bool) *oc.E_RoutingPolicy_PolicyResultTypeWatcher {
	t.Helper()
	w := &oc.E_RoutingPolicy_PolicyResultTypeWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions", gs, queryPath, true, false)
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_RoutingPolicy_PolicyResultType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_RoutingPolicy_PolicyResultType) bool) *oc.E_RoutingPolicy_PolicyResultTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Await(t testing.TB, timeout time.Duration, val oc.E_RoutingPolicy_PolicyResultType) *oc.QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_RoutingPolicy_PolicyResultType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_RoutingPolicy_PolicyResultType {
	t.Helper()
	c := &oc.CollectionE_RoutingPolicy_PolicyResultType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_RoutingPolicy_PolicyResultType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_RoutingPolicy_PolicyResultType) bool) *oc.E_RoutingPolicy_PolicyResultTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/state/policy-result to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath extracts the value of the leaf PolicyResult from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_RoutingPolicy_PolicyResultType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions) *oc.QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	qv := &oc.QualifiedE_RoutingPolicy_PolicyResultType{
		Metadata: md,
	}
	val := parent.PolicyResult
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", gs, queryPath, true, false)
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Await(t testing.TB, timeout time.Duration, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/state/tag to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath extracts the value of the leaf Tag from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice{
		Metadata: md,
	}
	val := parent.Tag
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Lookup(t testing.TB) *oc.QualifiedE_SetTag_Mode {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Get(t testing.TB) oc.E_SetTag_Mode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_SetTag_Mode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SetTag_Mode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePathAny) Get(t testing.TB) []oc.E_SetTag_Mode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SetTag_Mode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SetTag_Mode {
	t.Helper()
	c := &oc.CollectionE_SetTag_Mode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SetTag_Mode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SetTag_Mode) bool) *oc.E_SetTag_ModeWatcher {
	t.Helper()
	w := &oc.E_SetTag_ModeWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", gs, queryPath, true, false)
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SetTag_Mode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SetTag_Mode) bool) *oc.E_SetTag_ModeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Await(t testing.TB, timeout time.Duration, val oc.E_SetTag_Mode) *oc.QualifiedE_SetTag_Mode {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SetTag_Mode) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SetTag_Mode {
	t.Helper()
	c := &oc.CollectionE_SetTag_Mode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SetTag_Mode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SetTag_Mode) bool) *oc.E_SetTag_ModeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/state/mode to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath extracts the value of the leaf Mode from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SetTag_Mode.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) *oc.QualifiedE_SetTag_Mode {
	t.Helper()
	qv := &oc.QualifiedE_SetTag_Mode{
		Metadata: md,
	}
	val := parent.Mode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferencePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference", gs, queryPath, true, false)
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/reference/state/tag-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath extracts the value of the leaf TagSet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference_TagSetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TagSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_ConditionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions", gs, queryPath, false, false)
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_ConditionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_ConditionsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_ConditionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
