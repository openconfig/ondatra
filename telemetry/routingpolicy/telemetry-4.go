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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/state/community-set-ref to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath extracts the value of the leaf CommunitySetRef from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CommunitySetRef
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Await(t testing.TB, timeout time.Duration, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/state/communities to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath extracts the value of the leaf Communities from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice{
		Metadata: md,
	}
	val := parent.Communities
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Lookup(t testing.TB) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Get(t testing.TB) oc.E_SetCommunity_Method {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SetCommunity_Method
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny) Get(t testing.TB) []oc.E_SetCommunity_Method {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SetCommunity_Method
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SetCommunity_Method {
	t.Helper()
	c := &oc.CollectionE_SetCommunity_Method{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SetCommunity_Method) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	w := &oc.E_SetCommunity_MethodWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SetCommunity_Method)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Await(t testing.TB, timeout time.Duration, val oc.E_SetCommunity_Method) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SetCommunity_Method) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SetCommunity_Method {
	t.Helper()
	c := &oc.CollectionE_SetCommunity_Method{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SetCommunity_Method) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	w := &oc.E_SetCommunity_MethodWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SetCommunity_Method) bool) *oc.E_SetCommunity_MethodWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/method to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath extracts the value of the leaf Method from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SetCommunity_Method.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *oc.QualifiedE_SetCommunity_Method {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Lookup(t testing.TB) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Get(t testing.TB) oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny) Get(t testing.TB) []oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpPolicy_BgpSetCommunityOptionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	c := &oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	w := &oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Await(t testing.TB, timeout time.Duration, val oc.E_BgpPolicy_BgpSetCommunityOptionType) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	c := &oc.CollectionE_BgpPolicy_BgpSetCommunityOptionType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	w := &oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType) bool) *oc.E_BgpPolicy_BgpSetCommunityOptionTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/state/options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath extracts the value of the leaf Options from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/state/ext-community-set-ref to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath extracts the value of the leaf ExtCommunitySetRef from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ExtCommunitySetRef
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-local-pref to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath extracts the value of the leaf SetLocalPref from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SetLocalPref
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Get(t testing.TB) oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Await(t testing.TB, timeout time.Duration, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/state/set-med to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath extracts the value of the leaf SetMed from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union{
		Metadata: md,
	}
	val := parent.SetMed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
