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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/interface to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath extracts the value of the leaf Interface from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/state/subinterface to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsRestrictedType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Await(t testing.TB, timeout time.Duration, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/state/neighbor-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath extracts the value of the leaf NeighborSet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NeighborSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsRestrictedType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Await(t testing.TB, timeout time.Duration, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	c := &oc.CollectionE_PolicyTypes_MatchSetOptionsRestrictedType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	w := &oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) bool) *oc.E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/match-set-options to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, true, false)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", structs[pre], queryPath, true, false)
			qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/state/prefix-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath extracts the value of the leaf PrefixSet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PrefixSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet)))
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher{}
	gs := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", gs, queryPath, false, false)
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Await(t testing.TB, timeout time.Duration, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	c := &oc.CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher {
	t.Helper()
	w := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher{}
	structs := map[string]*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) bool) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher {
	t.Helper()
	return watch_RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set to the batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}
