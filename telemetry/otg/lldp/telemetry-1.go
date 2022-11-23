package lldp

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry/otg"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath) Lookup(t testing.TB) *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath) Get(t testing.TB) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPathAny) Lookup(t testing.TB) []*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPathAny) Get(t testing.TB) []*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities)))
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath) Await(t testing.TB, timeout time.Duration, val *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/enabled to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath extracts the value of the leaf Enabled from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath) Lookup(t testing.TB) *oc.QualifiedE_Capabilities_Name {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath) Get(t testing.TB) oc.E_Capabilities_Name {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Capabilities_Name {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Capabilities_Name
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePathAny) Get(t testing.TB) []oc.E_Capabilities_Name {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Capabilities_Name
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Capabilities_Name {
	t.Helper()
	c := &oc.CollectionE_Capabilities_Name{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Capabilities_Name) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Capabilities_Name) bool) *oc.E_Capabilities_NameWatcher {
	t.Helper()
	w := &oc.E_Capabilities_NameWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Capabilities_Name)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Capabilities_Name) bool) *oc.E_Capabilities_NameWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath) Await(t testing.TB, timeout time.Duration, val oc.E_Capabilities_Name) *oc.QualifiedE_Capabilities_Name {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Capabilities_Name) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Capabilities_Name {
	t.Helper()
	c := &oc.CollectionE_Capabilities_Name{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Capabilities_Name) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Capabilities_Name) bool) *oc.E_Capabilities_NameWatcher {
	t.Helper()
	w := &oc.E_Capabilities_NameWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Capabilities_Name)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Capabilities_Name) bool) *oc.E_Capabilities_NameWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/capabilities/state/name to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath extracts the value of the leaf Name from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Capabilities_Name.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) *oc.QualifiedE_Capabilities_Name {
	t.Helper()
	qv := &oc.QualifiedE_Capabilities_Name{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath extracts the value of the leaf ChassisId from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ChassisId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath) Lookup(t testing.TB) *oc.QualifiedE_LldpNeighbor_ChassisIdType {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath) Get(t testing.TB) oc.E_LldpNeighbor_ChassisIdType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpNeighbor_ChassisIdType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpNeighbor_ChassisIdType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePathAny) Get(t testing.TB) []oc.E_LldpNeighbor_ChassisIdType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_LldpNeighbor_ChassisIdType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpNeighbor_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpNeighbor_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpNeighbor_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_ChassisIdType) bool) *oc.E_LldpNeighbor_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpNeighbor_ChassisIdTypeWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpNeighbor_ChassisIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_ChassisIdType) bool) *oc.E_LldpNeighbor_ChassisIdTypeWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_LldpNeighbor_ChassisIdType) *oc.QualifiedE_LldpNeighbor_ChassisIdType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_LldpNeighbor_ChassisIdType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpNeighbor_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpNeighbor_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpNeighbor_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_ChassisIdType) bool) *oc.E_LldpNeighbor_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpNeighbor_ChassisIdTypeWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpNeighbor_ChassisIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_ChassisIdType) bool) *oc.E_LldpNeighbor_ChassisIdTypeWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/chassis-id-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath extracts the value of the leaf ChassisIdType from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpNeighbor_ChassisIdType.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ChassisIdTypePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedE_LldpNeighbor_ChassisIdType {
	t.Helper()
	qv := &oc.QualifiedE_LldpNeighbor_ChassisIdType{
		Metadata: md,
	}
	val := parent.ChassisIdType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath) Lookup(t testing.TB) *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, false, false)
	if ok {
		return (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath) Get(t testing.TB) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPathAny) Lookup(t testing.TB) []*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPathAny) Get(t testing.TB) []*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv)))
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", gs, queryPath, false, false)
		qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath) Await(t testing.TB, timeout time.Duration, val *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	c := &oc.CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher {
	t.Helper()
	w := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) bool) *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/custom-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath extracts the value of the leaf CustomType from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_CustomTypePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.CustomType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath extracts the value of the leaf Oui from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Oui
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/custom-tlv/state/oui-subtype to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath extracts the value of the leaf OuiSubtype from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv_OuiSubtypePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OuiSubtype
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetLastUpdate())
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/last_update to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath extracts the value of the leaf LastUpdate from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_LastUpdatePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.LastUpdate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath extracts the value of the leaf ManagementAddress from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ManagementAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
