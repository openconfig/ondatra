package discovery

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

// Lookup fetches the value at /open-traffic-generator-discovery/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *InterfacePath) Lookup(t testing.TB) *oc.QualifiedInterface {
	t.Helper()
	goStruct := &oc.Interface{}
	md, ok := oc.Lookup(t, n, "Interface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-discovery/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *InterfacePath) Get(t testing.TB) *oc.Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-discovery/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-discovery/interfaces/interface with a ONCE subscription.
func (n *InterfacePathAny) Get(t testing.TB) []*oc.Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface {
	t.Helper()
	c := &oc.CollectionInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface)))
		return false
	})
	return c
}

func watch_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface) bool) *oc.InterfaceWatcher {
	t.Helper()
	w := &oc.InterfaceWatcher{}
	gs := &oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface) bool) *oc.InterfaceWatcher {
	t.Helper()
	return watch_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-discovery/interfaces/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *InterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Interface) *oc.QualifiedInterface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-discovery/interfaces/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface to the batch object.
func (n *InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface {
	t.Helper()
	c := &oc.CollectionInterface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface) bool) *oc.InterfaceWatcher {
	t.Helper()
	w := &oc.InterfaceWatcher{}
	structs := map[string]*oc.Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface) bool) *oc.InterfaceWatcher {
	t.Helper()
	return watch_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface to the batch object.
func (n *InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ipv4NeighborPath) Lookup(t testing.TB) *oc.QualifiedInterface_Ipv4Neighbor {
	t.Helper()
	goStruct := &oc.Interface_Ipv4Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Ipv4Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_Ipv4Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ipv4NeighborPath) Get(t testing.TB) *oc.Interface_Ipv4Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ipv4NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Ipv4Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Ipv4Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ipv4Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ipv4Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Ipv4Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a ONCE subscription.
func (n *Interface_Ipv4NeighborPathAny) Get(t testing.TB) []*oc.Interface_Ipv4Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Ipv4Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv4NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ipv4Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_Ipv4Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ipv4Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_Ipv4Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_Ipv4Neighbor)))
		return false
	})
	return c
}

func watch_Interface_Ipv4NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ipv4Neighbor) bool) *oc.Interface_Ipv4NeighborWatcher {
	t.Helper()
	w := &oc.Interface_Ipv4NeighborWatcher{}
	gs := &oc.Interface_Ipv4Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ipv4Neighbor", gs, queryPath, false, false)
		qv := (&oc.QualifiedInterface_Ipv4Neighbor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ipv4Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv4NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ipv4Neighbor) bool) *oc.Interface_Ipv4NeighborWatcher {
	t.Helper()
	return watch_Interface_Ipv4NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ipv4NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_Ipv4Neighbor) *oc.QualifiedInterface_Ipv4Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_Ipv4Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor to the batch object.
func (n *Interface_Ipv4NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv4NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_Ipv4Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_Ipv4Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_Ipv4Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv4NeighborPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_Ipv4Neighbor) bool) *oc.Interface_Ipv4NeighborWatcher {
	t.Helper()
	w := &oc.Interface_Ipv4NeighborWatcher{}
	structs := map[string]*oc.Interface_Ipv4Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ipv4Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ipv4Neighbor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedInterface_Ipv4Neighbor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_Ipv4Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv4NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_Ipv4Neighbor) bool) *oc.Interface_Ipv4NeighborWatcher {
	t.Helper()
	return watch_Interface_Ipv4NeighborPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor to the batch object.
func (n *Interface_Ipv4NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Ipv4Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Ipv4Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_Ipv4Neighbor_Ipv4AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Ipv4Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Ipv4Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_Ipv4Neighbor_Ipv4AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a ONCE subscription.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv4Neighbor_Ipv4AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_Ipv4Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_Ipv4Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertInterface_Ipv4Neighbor_Ipv4AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ipv4Neighbor_Ipv4AddressPath(t, n, timeout, predicate)
}

// Await observes values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address to the batch object.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_Ipv4Neighbor_Ipv4AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.Interface_Ipv4Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.Interface_Ipv4Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "Interface_Ipv4Neighbor", structs[pre], queryPath, true, false)
			qv := convertInterface_Ipv4Neighbor_Ipv4AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_Ipv4Neighbor_Ipv4AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /open-traffic-generator-discovery/interfaces/interface/ipv4-neighbors/ipv4-neighbor/state/ipv4-address to the batch object.
func (n *Interface_Ipv4Neighbor_Ipv4AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_Ipv4Neighbor_Ipv4AddressPath extracts the value of the leaf Ipv4Address from its parent oc.Interface_Ipv4Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Ipv4Neighbor_Ipv4AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Ipv4Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ipv4Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
