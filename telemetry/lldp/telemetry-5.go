package lldp

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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Capability{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Capability", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_Capability_EnabledPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Capability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_Capability_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Capability_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Lldp_Interface_Neighbor_Capability{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_Capability_EnabledPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Capability_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled to the batch object.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Capability_EnabledPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Capability{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Capability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_Capability_EnabledPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Capability_EnabledPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/enabled to the batch object.
func (n *Lldp_Interface_Neighbor_Capability_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_Capability_EnabledPath extracts the value of the leaf Enabled from its parent oc.Lldp_Interface_Neighbor_Capability
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLldp_Interface_Neighbor_Capability_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor_Capability) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_Capability_NamePath) Lookup(t testing.TB) *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor_Capability{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor_Capability", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_Capability_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_Capability_NamePath) Get(t testing.TB) oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_Capability_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor_Capability{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_Capability_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_Capability_NamePathAny) Get(t testing.TB) []oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Capability_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_LLDP_SYSTEM_CAPABILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Capability_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) bool) *oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher{}
	gs := &oc.Lldp_Interface_Neighbor_Capability{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_Capability_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Capability_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) bool) *oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Capability_NamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_Capability_NamePath) Await(t testing.TB, timeout time.Duration, val oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITY) *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name to the batch object.
func (n *Lldp_Interface_Neighbor_Capability_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_Capability_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_LLDP_SYSTEM_CAPABILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_Capability_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) bool) *oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor_Capability{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor_Capability{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor_Capability", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_Capability_NamePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_Capability_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) bool) *oc.E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_Capability_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/capabilities/capability/state/name to the batch object.
func (n *Lldp_Interface_Neighbor_Capability_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_Capability_NamePath extracts the value of the leaf Name from its parent oc.Lldp_Interface_Neighbor_Capability
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY.
func convertLldp_Interface_Neighbor_Capability_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor_Capability) *oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	qv := &oc.QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_ChassisIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_ChassisIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_ChassisIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_ChassisIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_ChassisIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_ChassisIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_ChassisIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_ChassisIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_ChassisIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_ChassisIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_ChassisIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_ChassisIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id to the batch object.
func (n *Lldp_Interface_Neighbor_ChassisIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_ChassisIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_ChassisIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_ChassisIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_ChassisIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_ChassisIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id to the batch object.
func (n *Lldp_Interface_Neighbor_ChassisIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_ChassisIdPath extracts the value of the leaf ChassisId from its parent oc.Lldp_Interface_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_Neighbor_ChassisIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePath) Lookup(t testing.TB) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_ChassisIdTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePath) Get(t testing.TB) oc.E_LldpTypes_ChassisIdType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpTypes_ChassisIdType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_ChassisIdTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePathAny) Get(t testing.TB) []oc.E_LldpTypes_ChassisIdType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_LldpTypes_ChassisIdType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_ChassisIdTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_ChassisIdTypeWatcher{}
	gs := &oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_ChassisIdTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_ChassisIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_ChassisIdTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_LldpTypes_ChassisIdType) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type to the batch object.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_ChassisIdTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_ChassisIdTypeWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_ChassisIdTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_ChassisIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_ChassisIdTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type to the batch object.
func (n *Lldp_Interface_Neighbor_ChassisIdTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_ChassisIdTypePath extracts the value of the leaf ChassisIdType from its parent oc.Lldp_Interface_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpTypes_ChassisIdType.
func convertLldp_Interface_Neighbor_ChassisIdTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	qv := &oc.QualifiedE_LldpTypes_ChassisIdType{
		Metadata: md,
	}
	val := parent.ChassisIdType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id to the batch object.
func (n *Lldp_Interface_Neighbor_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_IdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id to the batch object.
func (n *Lldp_Interface_Neighbor_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_IdPath extracts the value of the leaf Id from its parent oc.Lldp_Interface_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldp_Interface_Neighbor_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_Interface_Neighbor_LastUpdatePath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Lldp_Interface_Neighbor{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface_Neighbor", goStruct, true, false)
	if ok {
		return convertLldp_Interface_Neighbor_LastUpdatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_Interface_Neighbor_LastUpdatePath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_Interface_Neighbor_LastUpdatePathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldp_Interface_Neighbor_LastUpdatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a ONCE subscription.
func (n *Lldp_Interface_Neighbor_LastUpdatePathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_LastUpdatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_LastUpdatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Lldp_Interface_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldp_Interface_Neighbor_LastUpdatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_LastUpdatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_LastUpdatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Lldp_Interface_Neighbor_LastUpdatePath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update to the batch object.
func (n *Lldp_Interface_Neighbor_LastUpdatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Lldp_Interface_Neighbor_LastUpdatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Lldp_Interface_Neighbor_LastUpdatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	structs := map[string]*oc.Lldp_Interface_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Lldp_Interface_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Lldp_Interface_Neighbor", structs[pre], queryPath, true, false)
			qv := convertLldp_Interface_Neighbor_LastUpdatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Lldp_Interface_Neighbor_LastUpdatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_Lldp_Interface_Neighbor_LastUpdatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update to the batch object.
func (n *Lldp_Interface_Neighbor_LastUpdatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldp_Interface_Neighbor_LastUpdatePath extracts the value of the leaf LastUpdate from its parent oc.Lldp_Interface_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertLldp_Interface_Neighbor_LastUpdatePath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp_Interface_Neighbor) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.LastUpdate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
