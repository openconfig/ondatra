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

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/management-address-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath extracts the value of the leaf ManagementAddressType from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_ManagementAddressTypePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ManagementAddressType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/neighbor_id to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath extracts the value of the leaf NeighborId from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_NeighborIdPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NeighborId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-description to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath extracts the value of the leaf PortDescription from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortDescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PortDescription
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath extracts the value of the leaf PortId from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PortId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath) Lookup(t testing.TB) *oc.QualifiedE_LldpNeighbor_PortIdType {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath) Get(t testing.TB) oc.E_LldpNeighbor_PortIdType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpNeighbor_PortIdType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpNeighbor_PortIdType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePathAny) Get(t testing.TB) []oc.E_LldpNeighbor_PortIdType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_LldpNeighbor_PortIdType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpNeighbor_PortIdType {
	t.Helper()
	c := &oc.CollectionE_LldpNeighbor_PortIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpNeighbor_PortIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_PortIdType) bool) *oc.E_LldpNeighbor_PortIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpNeighbor_PortIdTypeWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpNeighbor_PortIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_PortIdType) bool) *oc.E_LldpNeighbor_PortIdTypeWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_LldpNeighbor_PortIdType) *oc.QualifiedE_LldpNeighbor_PortIdType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_LldpNeighbor_PortIdType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpNeighbor_PortIdType {
	t.Helper()
	c := &oc.CollectionE_LldpNeighbor_PortIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpNeighbor_PortIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_PortIdType) bool) *oc.E_LldpNeighbor_PortIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpNeighbor_PortIdTypeWatcher{}
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpNeighbor_PortIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpNeighbor_PortIdType) bool) *oc.E_LldpNeighbor_PortIdTypeWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/port-id-type to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath extracts the value of the leaf PortIdType from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpNeighbor_PortIdType.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_PortIdTypePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedE_LldpNeighbor_PortIdType {
	t.Helper()
	qv := &oc.QualifiedE_LldpNeighbor_PortIdType{
		Metadata: md,
	}
	val := parent.PortIdType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-description to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath extracts the value of the leaf SystemDescription from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemDescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SystemDescription
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/system-name to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath extracts the value of the leaf SystemName from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_SystemNamePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SystemName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	md, ok := oc.Lookup(t, n, "LldpInterface_LldpNeighborDatabase_LldpNeighbor", goStruct, true, false)
	if ok {
		return convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetTtl())
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
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
		qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a ONCE subscription.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface_LldpNeighborDatabase_LldpNeighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
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
			qv := convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/lldp-neighbor-database/lldp-neighbors/lldp-neighbor/state/ttl to the batch object.
func (n *LldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath extracts the value of the leaf Ttl from its parent oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLldpInterface_LldpNeighborDatabase_LldpNeighbor_TtlPath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface_LldpNeighborDatabase_LldpNeighbor) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Ttl
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LldpInterface_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LldpInterface{}
	md, ok := oc.Lookup(t, n, "LldpInterface", goStruct, true, false)
	if ok {
		return convertLldpInterface_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LldpInterface_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LldpInterface_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LldpInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LldpInterface", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertLldpInterface_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a ONCE subscription.
func (n *LldpInterface_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_NamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_NamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.LldpInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "LldpInterface", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertLldpInterface_NamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_NamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_NamePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *LldpInterface_NamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-lldp/lldps/lldp-interface/state/name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/name to the batch object.
func (n *LldpInterface_NamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *LldpInterface_NamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_LldpInterface_NamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.LldpInterface{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.LldpInterface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "LldpInterface", structs[pre], queryPath, true, false)
			qv := convertLldpInterface_NamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-lldp/lldps/lldp-interface/state/name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *LldpInterface_NamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_LldpInterface_NamePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-lldp/lldps/lldp-interface/state/name to the batch object.
func (n *LldpInterface_NamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertLldpInterface_NamePath extracts the value of the leaf Name from its parent oc.LldpInterface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLldpInterface_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.LldpInterface) *oc.QualifiedString {
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
