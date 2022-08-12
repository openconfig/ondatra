package system

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

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_HostnamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, false)
	if ok {
		return convertSystem_Dns_HostEntry_HostnamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_HostnamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_HostnamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_HostnamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a ONCE subscription.
func (n *System_Dns_HostEntry_HostnamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_HostnamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_HostnamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_HostEntry_HostnamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_HostnamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_HostnamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntry_HostnamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry/state/hostname failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/hostname to the batch object.
func (n *System_Dns_HostEntry_HostnamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_HostnamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_HostnamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_HostEntry_HostnamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_HostnamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_HostnamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/hostname to the batch object.
func (n *System_Dns_HostEntry_HostnamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_HostEntry_HostnamePath extracts the value of the leaf Hostname from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Dns_HostEntry_HostnamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Hostname
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, false)
	if ok {
		return convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a ONCE subscription.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv4AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv4AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address to the batch object.
func (n *System_Dns_HostEntry_Ipv4AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv4AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_HostEntry_Ipv4AddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv4AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv4-address to the batch object.
func (n *System_Dns_HostEntry_Ipv4AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_HostEntry_Ipv4AddressPath extracts the value of the leaf Ipv4Address from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_Ipv4AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Ipv4Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns_HostEntry{}
	md, ok := oc.Lookup(t, n, "System_Dns_HostEntry", goStruct, true, false)
	if ok {
		return convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_HostEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_HostEntry", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a ONCE subscription.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv6AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_HostEntry", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv6AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address to the batch object.
func (n *System_Dns_HostEntry_Ipv6AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_HostEntry_Ipv6AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.System_Dns_HostEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_HostEntry{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_HostEntry", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_HostEntry_Ipv6AddressPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_HostEntry_Ipv6AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/host-entries/host-entry/state/ipv6-address to the batch object.
func (n *System_Dns_HostEntry_Ipv6AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_HostEntry_Ipv6AddressPath extracts the value of the leaf Ipv6Address from its parent oc.System_Dns_HostEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_HostEntry_Ipv6AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_HostEntry) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Ipv6Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/state/search with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_SearchPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.System_Dns{}
	md, ok := oc.Lookup(t, n, "System_Dns", goStruct, true, false)
	if ok {
		return convertSystem_Dns_SearchPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/state/search with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_SearchPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/state/search with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_SearchPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_SearchPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/state/search with a ONCE subscription.
func (n *System_Dns_SearchPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/state/search with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_SearchPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_SearchPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	gs := &oc.System_Dns{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_SearchPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/state/search with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_SearchPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_SearchPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/state/search with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_SearchPath) Await(t testing.TB, timeout time.Duration, val []string) *oc.QualifiedStringSlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedStringSlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/state/search failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/state/search to the batch object.
func (n *System_Dns_SearchPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/state/search with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_SearchPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionStringSlice {
	t.Helper()
	c := &oc.CollectionStringSlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedStringSlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_SearchPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	w := &oc.StringSliceWatcher{}
	structs := map[string]*oc.System_Dns{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_SearchPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedStringSlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/state/search with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_SearchPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedStringSlice) bool) *oc.StringSliceWatcher {
	t.Helper()
	return watch_System_Dns_SearchPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/state/search to the batch object.
func (n *System_Dns_SearchPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_SearchPath extracts the value of the leaf Search from its parent oc.System_Dns
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertSystem_Dns_SearchPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Search
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_ServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_Dns_Server {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/servers/server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_ServerPath) Get(t testing.TB) *oc.System_Dns_Server {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_ServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Dns_Server {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Dns_Server
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server with a ONCE subscription.
func (n *System_Dns_ServerPathAny) Get(t testing.TB) []*oc.System_Dns_Server {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Dns_Server
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_ServerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns_Server {
	t.Helper()
	c := &oc.CollectionSystem_Dns_Server{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns_Server) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Dns_Server{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Dns_Server)))
		return false
	})
	return c
}

func watch_System_Dns_ServerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	w := &oc.System_Dns_ServerWatcher{}
	gs := &oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_Server", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Dns_Server{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns_Server)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_ServerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	return watch_System_Dns_ServerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/servers/server with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_ServerPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Dns_Server) *oc.QualifiedSystem_Dns_Server {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Dns_Server) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/servers/server failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/servers/server to the batch object.
func (n *System_Dns_ServerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_ServerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Dns_Server {
	t.Helper()
	c := &oc.CollectionSystem_Dns_Server{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Dns_Server) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_ServerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	w := &oc.System_Dns_ServerWatcher{}
	structs := map[string]*oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_Server{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_Server", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Dns_Server{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Dns_Server)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_ServerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Dns_Server) bool) *oc.System_Dns_ServerWatcher {
	t.Helper()
	return watch_System_Dns_ServerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/servers/server to the batch object.
func (n *System_Dns_ServerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_Server_AddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, true, false)
	if ok {
		return convertSystem_Dns_Server_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_Server_AddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_Server_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_Server_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server/state/address with a ONCE subscription.
func (n *System_Dns_Server_AddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_Server_AddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_Server_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_Server_AddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/servers/server/state/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/servers/server/state/address to the batch object.
func (n *System_Dns_Server_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_AddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_Server{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_Server", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_Server_AddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Dns_Server_AddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/servers/server/state/address to the batch object.
func (n *System_Dns_Server_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_Server_AddressPath extracts the value of the leaf Address from its parent oc.System_Dns_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Dns_Server_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_Server) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Dns_Server_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Dns_Server{}
	md, ok := oc.Lookup(t, n, "System_Dns_Server", goStruct, true, false)
	if ok {
		return convertSystem_Dns_Server_PortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPort())
}

// Get fetches the value at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Dns_Server_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Dns_Server_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Dns_Server{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Dns_Server", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Dns_Server_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/dns/servers/server/state/port with a ONCE subscription.
func (n *System_Dns_Server_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_PortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_PortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Dns_Server", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Dns_Server_PortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_PortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Dns_Server_PortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Dns_Server_PortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/dns/servers/server/state/port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/dns/servers/server/state/port to the batch object.
func (n *System_Dns_Server_PortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Dns_Server_PortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Dns_Server_PortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.System_Dns_Server{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Dns_Server{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Dns_Server", structs[pre], queryPath, true, false)
			qv := convertSystem_Dns_Server_PortPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/dns/servers/server/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Dns_Server_PortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Dns_Server_PortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/dns/servers/server/state/port to the batch object.
func (n *System_Dns_Server_PortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Dns_Server_PortPath extracts the value of the leaf Port from its parent oc.System_Dns_Server
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Dns_Server_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Dns_Server) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Port
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/state/domain-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_DomainNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, false)
	if ok {
		return convertSystem_DomainNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/state/domain-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_DomainNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/state/domain-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_DomainNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_DomainNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/state/domain-name with a ONCE subscription.
func (n *System_DomainNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/domain-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_DomainNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_DomainNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_DomainNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/domain-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_DomainNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_DomainNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/state/domain-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_DomainNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/state/domain-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/state/domain-name to the batch object.
func (n *System_DomainNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/domain-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_DomainNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_DomainNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System", structs[pre], queryPath, true, false)
			qv := convertSystem_DomainNamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/domain-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_DomainNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_DomainNamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/state/domain-name to the batch object.
func (n *System_DomainNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_DomainNamePath extracts the value of the leaf DomainName from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_DomainNamePath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DomainName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/gnmi-pathz-policies with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GnmiPathzPoliciesPath) Lookup(t testing.TB) *oc.QualifiedSystem_GnmiPathzPolicies {
	t.Helper()
	goStruct := &oc.System_GnmiPathzPolicies{}
	md, ok := oc.Lookup(t, n, "System_GnmiPathzPolicies", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_GnmiPathzPolicies{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/gnmi-pathz-policies with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GnmiPathzPoliciesPath) Get(t testing.TB) *oc.System_GnmiPathzPolicies {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/gnmi-pathz-policies with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GnmiPathzPoliciesPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GnmiPathzPolicies {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GnmiPathzPolicies
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GnmiPathzPolicies{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GnmiPathzPolicies", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_GnmiPathzPolicies{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/gnmi-pathz-policies with a ONCE subscription.
func (n *System_GnmiPathzPoliciesPathAny) Get(t testing.TB) []*oc.System_GnmiPathzPolicies {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_GnmiPathzPolicies
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPoliciesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GnmiPathzPolicies {
	t.Helper()
	c := &oc.CollectionSystem_GnmiPathzPolicies{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GnmiPathzPolicies) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_GnmiPathzPolicies{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_GnmiPathzPolicies)))
		return false
	})
	return c
}

func watch_System_GnmiPathzPoliciesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies) bool) *oc.System_GnmiPathzPoliciesWatcher {
	t.Helper()
	w := &oc.System_GnmiPathzPoliciesWatcher{}
	gs := &oc.System_GnmiPathzPolicies{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GnmiPathzPolicies", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_GnmiPathzPolicies{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GnmiPathzPolicies)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPoliciesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies) bool) *oc.System_GnmiPathzPoliciesWatcher {
	t.Helper()
	return watch_System_GnmiPathzPoliciesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/gnmi-pathz-policies with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GnmiPathzPoliciesPath) Await(t testing.TB, timeout time.Duration, val *oc.System_GnmiPathzPolicies) *oc.QualifiedSystem_GnmiPathzPolicies {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_GnmiPathzPolicies) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/gnmi-pathz-policies failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies to the batch object.
func (n *System_GnmiPathzPoliciesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPoliciesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GnmiPathzPolicies {
	t.Helper()
	c := &oc.CollectionSystem_GnmiPathzPolicies{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GnmiPathzPolicies) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPoliciesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies) bool) *oc.System_GnmiPathzPoliciesWatcher {
	t.Helper()
	w := &oc.System_GnmiPathzPoliciesWatcher{}
	structs := map[string]*oc.System_GnmiPathzPolicies{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GnmiPathzPolicies{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GnmiPathzPolicies", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_GnmiPathzPolicies{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GnmiPathzPolicies)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPoliciesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies) bool) *oc.System_GnmiPathzPoliciesWatcher {
	t.Helper()
	return watch_System_GnmiPathzPoliciesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies to the batch object.
func (n *System_GnmiPathzPoliciesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GnmiPathzPolicies_PolicyPath) Lookup(t testing.TB) *oc.QualifiedSystem_GnmiPathzPolicies_Policy {
	t.Helper()
	goStruct := &oc.System_GnmiPathzPolicies_Policy{}
	md, ok := oc.Lookup(t, n, "System_GnmiPathzPolicies_Policy", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_GnmiPathzPolicies_Policy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GnmiPathzPolicies_PolicyPath) Get(t testing.TB) *oc.System_GnmiPathzPolicies_Policy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GnmiPathzPolicies_PolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GnmiPathzPolicies_Policy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GnmiPathzPolicies_Policy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GnmiPathzPolicies_Policy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_GnmiPathzPolicies_Policy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a ONCE subscription.
func (n *System_GnmiPathzPolicies_PolicyPathAny) Get(t testing.TB) []*oc.System_GnmiPathzPolicies_Policy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_GnmiPathzPolicies_Policy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_PolicyPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GnmiPathzPolicies_Policy {
	t.Helper()
	c := &oc.CollectionSystem_GnmiPathzPolicies_Policy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GnmiPathzPolicies_Policy) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_GnmiPathzPolicies_Policy{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_GnmiPathzPolicies_Policy)))
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_PolicyPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies_Policy) bool) *oc.System_GnmiPathzPolicies_PolicyWatcher {
	t.Helper()
	w := &oc.System_GnmiPathzPolicies_PolicyWatcher{}
	gs := &oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GnmiPathzPolicies_Policy", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_GnmiPathzPolicies_Policy{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GnmiPathzPolicies_Policy)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_PolicyPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies_Policy) bool) *oc.System_GnmiPathzPolicies_PolicyWatcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_PolicyPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GnmiPathzPolicies_PolicyPath) Await(t testing.TB, timeout time.Duration, val *oc.System_GnmiPathzPolicies_Policy) *oc.QualifiedSystem_GnmiPathzPolicies_Policy {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_GnmiPathzPolicies_Policy) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/gnmi-pathz-policies/policies/policy failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy to the batch object.
func (n *System_GnmiPathzPolicies_PolicyPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_PolicyPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GnmiPathzPolicies_Policy {
	t.Helper()
	c := &oc.CollectionSystem_GnmiPathzPolicies_Policy{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GnmiPathzPolicies_Policy) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_PolicyPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies_Policy) bool) *oc.System_GnmiPathzPolicies_PolicyWatcher {
	t.Helper()
	w := &oc.System_GnmiPathzPolicies_PolicyWatcher{}
	structs := map[string]*oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GnmiPathzPolicies_Policy{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_GnmiPathzPolicies_Policy{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GnmiPathzPolicies_Policy)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_PolicyPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GnmiPathzPolicies_Policy) bool) *oc.System_GnmiPathzPolicies_PolicyWatcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_PolicyPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy to the batch object.
func (n *System_GnmiPathzPolicies_PolicyPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_GnmiPathzPolicies_Policy{}
	md, ok := oc.Lookup(t, n, "System_GnmiPathzPolicies_Policy", goStruct, true, false)
	if ok {
		return convertSystem_GnmiPathzPolicies_Policy_CreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GnmiPathzPolicies_Policy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GnmiPathzPolicies_Policy_CreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a ONCE subscription.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_Policy_CreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GnmiPathzPolicies_Policy", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GnmiPathzPolicies_Policy_CreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_Policy_CreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on to the batch object.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_Policy_CreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GnmiPathzPolicies_Policy{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", structs[pre], queryPath, true, false)
			qv := convertSystem_GnmiPathzPolicies_Policy_CreatedOnPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_Policy_CreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/created-on to the batch object.
func (n *System_GnmiPathzPolicies_Policy_CreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GnmiPathzPolicies_Policy_CreatedOnPath extracts the value of the leaf CreatedOn from its parent oc.System_GnmiPathzPolicies_Policy
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_GnmiPathzPolicies_Policy_CreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GnmiPathzPolicies_Policy) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GnmiPathzPolicies_Policy_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GnmiPathzPolicies_Policy{}
	md, ok := oc.Lookup(t, n, "System_GnmiPathzPolicies_Policy", goStruct, true, false)
	if ok {
		return convertSystem_GnmiPathzPolicies_Policy_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GnmiPathzPolicies_Policy_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GnmiPathzPolicies_Policy_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GnmiPathzPolicies_Policy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GnmiPathzPolicies_Policy_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a ONCE subscription.
func (n *System_GnmiPathzPolicies_Policy_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_Policy_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_Policy_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GnmiPathzPolicies_Policy", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GnmiPathzPolicies_Policy_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_Policy_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_Policy_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GnmiPathzPolicies_Policy_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id to the batch object.
func (n *System_GnmiPathzPolicies_Policy_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_Policy_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_Policy_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GnmiPathzPolicies_Policy{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", structs[pre], queryPath, true, false)
			qv := convertSystem_GnmiPathzPolicies_Policy_IdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_Policy_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_Policy_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/id to the batch object.
func (n *System_GnmiPathzPolicies_Policy_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GnmiPathzPolicies_Policy_IdPath extracts the value of the leaf Id from its parent oc.System_GnmiPathzPolicies_Policy
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GnmiPathzPolicies_Policy_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GnmiPathzPolicies_Policy) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GnmiPathzPolicies_Policy_VersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GnmiPathzPolicies_Policy{}
	md, ok := oc.Lookup(t, n, "System_GnmiPathzPolicies_Policy", goStruct, true, false)
	if ok {
		return convertSystem_GnmiPathzPolicies_Policy_VersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GnmiPathzPolicies_Policy_VersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GnmiPathzPolicies_Policy_VersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GnmiPathzPolicies_Policy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GnmiPathzPolicies_Policy_VersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a ONCE subscription.
func (n *System_GnmiPathzPolicies_Policy_VersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_Policy_VersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_Policy_VersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GnmiPathzPolicies_Policy", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GnmiPathzPolicies_Policy_VersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_Policy_VersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_Policy_VersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GnmiPathzPolicies_Policy_VersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version to the batch object.
func (n *System_GnmiPathzPolicies_Policy_VersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GnmiPathzPolicies_Policy_VersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GnmiPathzPolicies_Policy_VersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GnmiPathzPolicies_Policy{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GnmiPathzPolicies_Policy{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GnmiPathzPolicies_Policy", structs[pre], queryPath, true, false)
			qv := convertSystem_GnmiPathzPolicies_Policy_VersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GnmiPathzPolicies_Policy_VersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GnmiPathzPolicies_Policy_VersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/gnmi-pathz-policies/policies/policy/state/version to the batch object.
func (n *System_GnmiPathzPolicies_Policy_VersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GnmiPathzPolicies_Policy_VersionPath extracts the value of the leaf Version from its parent oc.System_GnmiPathzPolicies_Policy
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GnmiPathzPolicies_Policy_VersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GnmiPathzPolicies_Policy) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Version
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-credentials with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcCredentialsPath) Lookup(t testing.TB) *oc.QualifiedSystem_GrpcCredentials {
	t.Helper()
	goStruct := &oc.System_GrpcCredentials{}
	md, ok := oc.Lookup(t, n, "System_GrpcCredentials", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_GrpcCredentials{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-credentials with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcCredentialsPath) Get(t testing.TB) *oc.System_GrpcCredentials {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-credentials with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcCredentialsPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GrpcCredentials {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GrpcCredentials
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcCredentials{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcCredentials", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_GrpcCredentials{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-credentials with a ONCE subscription.
func (n *System_GrpcCredentialsPathAny) Get(t testing.TB) []*oc.System_GrpcCredentials {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_GrpcCredentials
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentialsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcCredentials {
	t.Helper()
	c := &oc.CollectionSystem_GrpcCredentials{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcCredentials) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_GrpcCredentials{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_GrpcCredentials)))
		return false
	})
	return c
}

func watch_System_GrpcCredentialsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials) bool) *oc.System_GrpcCredentialsWatcher {
	t.Helper()
	w := &oc.System_GrpcCredentialsWatcher{}
	gs := &oc.System_GrpcCredentials{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcCredentials", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_GrpcCredentials{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcCredentials)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentialsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials) bool) *oc.System_GrpcCredentialsWatcher {
	t.Helper()
	return watch_System_GrpcCredentialsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-credentials with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcCredentialsPath) Await(t testing.TB, timeout time.Duration, val *oc.System_GrpcCredentials) *oc.QualifiedSystem_GrpcCredentials {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_GrpcCredentials) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-credentials failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-credentials to the batch object.
func (n *System_GrpcCredentialsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentialsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcCredentials {
	t.Helper()
	c := &oc.CollectionSystem_GrpcCredentials{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcCredentials) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentialsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials) bool) *oc.System_GrpcCredentialsWatcher {
	t.Helper()
	w := &oc.System_GrpcCredentialsWatcher{}
	structs := map[string]*oc.System_GrpcCredentials{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcCredentials{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcCredentials", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_GrpcCredentials{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcCredentials)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentialsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials) bool) *oc.System_GrpcCredentialsWatcher {
	t.Helper()
	return watch_System_GrpcCredentialsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-credentials to the batch object.
func (n *System_GrpcCredentialsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/grpc-credentials/entities/entity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcCredentials_EntityPath) Lookup(t testing.TB) *oc.QualifiedSystem_GrpcCredentials_Entity {
	t.Helper()
	goStruct := &oc.System_GrpcCredentials_Entity{}
	md, ok := oc.Lookup(t, n, "System_GrpcCredentials_Entity", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_GrpcCredentials_Entity{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-credentials/entities/entity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcCredentials_EntityPath) Get(t testing.TB) *oc.System_GrpcCredentials_Entity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-credentials/entities/entity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcCredentials_EntityPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GrpcCredentials_Entity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GrpcCredentials_Entity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcCredentials_Entity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcCredentials_Entity", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_GrpcCredentials_Entity{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-credentials/entities/entity with a ONCE subscription.
func (n *System_GrpcCredentials_EntityPathAny) Get(t testing.TB) []*oc.System_GrpcCredentials_Entity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_GrpcCredentials_Entity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_EntityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcCredentials_Entity {
	t.Helper()
	c := &oc.CollectionSystem_GrpcCredentials_Entity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcCredentials_Entity) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_GrpcCredentials_Entity{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_GrpcCredentials_Entity)))
		return false
	})
	return c
}

func watch_System_GrpcCredentials_EntityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials_Entity) bool) *oc.System_GrpcCredentials_EntityWatcher {
	t.Helper()
	w := &oc.System_GrpcCredentials_EntityWatcher{}
	gs := &oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcCredentials_Entity", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_GrpcCredentials_Entity{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcCredentials_Entity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_EntityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials_Entity) bool) *oc.System_GrpcCredentials_EntityWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_EntityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-credentials/entities/entity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcCredentials_EntityPath) Await(t testing.TB, timeout time.Duration, val *oc.System_GrpcCredentials_Entity) *oc.QualifiedSystem_GrpcCredentials_Entity {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_GrpcCredentials_Entity) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-credentials/entities/entity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity to the batch object.
func (n *System_GrpcCredentials_EntityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_EntityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcCredentials_Entity {
	t.Helper()
	c := &oc.CollectionSystem_GrpcCredentials_Entity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcCredentials_Entity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_EntityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials_Entity) bool) *oc.System_GrpcCredentials_EntityWatcher {
	t.Helper()
	w := &oc.System_GrpcCredentials_EntityWatcher{}
	structs := map[string]*oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcCredentials_Entity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcCredentials_Entity", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_GrpcCredentials_Entity{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcCredentials_Entity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_EntityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcCredentials_Entity) bool) *oc.System_GrpcCredentials_EntityWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_EntityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity to the batch object.
func (n *System_GrpcCredentials_EntityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcCredentials_Entity_CreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_GrpcCredentials_Entity{}
	md, ok := oc.Lookup(t, n, "System_GrpcCredentials_Entity", goStruct, true, false)
	if ok {
		return convertSystem_GrpcCredentials_Entity_CreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcCredentials_Entity_CreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcCredentials_Entity_CreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcCredentials_Entity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcCredentials_Entity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcCredentials_Entity_CreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a ONCE subscription.
func (n *System_GrpcCredentials_Entity_CreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_CreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_CreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcCredentials_Entity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcCredentials_Entity_CreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_CreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_CreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcCredentials_Entity_CreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/created-on to the batch object.
func (n *System_GrpcCredentials_Entity_CreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_CreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_CreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcCredentials_Entity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcCredentials_Entity", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcCredentials_Entity_CreatedOnPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_CreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_CreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/created-on to the batch object.
func (n *System_GrpcCredentials_Entity_CreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcCredentials_Entity_CreatedOnPath extracts the value of the leaf CreatedOn from its parent oc.System_GrpcCredentials_Entity
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_GrpcCredentials_Entity_CreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcCredentials_Entity) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcCredentials_Entity_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcCredentials_Entity{}
	md, ok := oc.Lookup(t, n, "System_GrpcCredentials_Entity", goStruct, true, false)
	if ok {
		return convertSystem_GrpcCredentials_Entity_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcCredentials_Entity_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcCredentials_Entity_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcCredentials_Entity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcCredentials_Entity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcCredentials_Entity_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a ONCE subscription.
func (n *System_GrpcCredentials_Entity_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcCredentials_Entity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcCredentials_Entity_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcCredentials_Entity_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-credentials/entities/entity/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/id to the batch object.
func (n *System_GrpcCredentials_Entity_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcCredentials_Entity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcCredentials_Entity", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcCredentials_Entity_IdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/id to the batch object.
func (n *System_GrpcCredentials_Entity_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcCredentials_Entity_IdPath extracts the value of the leaf Id from its parent oc.System_GrpcCredentials_Entity
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcCredentials_Entity_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcCredentials_Entity) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcCredentials_Entity_KindPath) Lookup(t testing.TB) *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	goStruct := &oc.System_GrpcCredentials_Entity{}
	md, ok := oc.Lookup(t, n, "System_GrpcCredentials_Entity", goStruct, true, false)
	if ok {
		return convertSystem_GrpcCredentials_Entity_KindPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcCredentials_Entity_KindPath) Get(t testing.TB) oc.E_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcCredentials_Entity_KindPathAny) Lookup(t testing.TB) []*oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcCredentials_Entity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcCredentials_Entity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcCredentials_Entity_KindPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a ONCE subscription.
func (n *System_GrpcCredentials_Entity_KindPathAny) Get(t testing.TB) []oc.E_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_GnsiCert_GrpcCredentialEntityKind
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_KindPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	c := &oc.CollectionE_GnsiCert_GrpcCredentialEntityKind{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_KindPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind) bool) *oc.E_GnsiCert_GrpcCredentialEntityKindWatcher {
	t.Helper()
	w := &oc.E_GnsiCert_GrpcCredentialEntityKindWatcher{}
	gs := &oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcCredentials_Entity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcCredentials_Entity_KindPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_KindPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind) bool) *oc.E_GnsiCert_GrpcCredentialEntityKindWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_KindPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcCredentials_Entity_KindPath) Await(t testing.TB, timeout time.Duration, val oc.E_GnsiCert_GrpcCredentialEntityKind) *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-credentials/entities/entity/state/kind failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/kind to the batch object.
func (n *System_GrpcCredentials_Entity_KindPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_KindPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	c := &oc.CollectionE_GnsiCert_GrpcCredentialEntityKind{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_KindPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind) bool) *oc.E_GnsiCert_GrpcCredentialEntityKindWatcher {
	t.Helper()
	w := &oc.E_GnsiCert_GrpcCredentialEntityKindWatcher{}
	structs := map[string]*oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcCredentials_Entity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcCredentials_Entity", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcCredentials_Entity_KindPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/kind with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_KindPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind) bool) *oc.E_GnsiCert_GrpcCredentialEntityKindWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_KindPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/kind to the batch object.
func (n *System_GrpcCredentials_Entity_KindPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcCredentials_Entity_KindPath extracts the value of the leaf Kind from its parent oc.System_GrpcCredentials_Entity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind.
func convertSystem_GrpcCredentials_Entity_KindPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcCredentials_Entity) *oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind {
	t.Helper()
	qv := &oc.QualifiedE_GnsiCert_GrpcCredentialEntityKind{
		Metadata: md,
	}
	val := parent.Kind
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcCredentials_Entity_VersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcCredentials_Entity{}
	md, ok := oc.Lookup(t, n, "System_GrpcCredentials_Entity", goStruct, true, false)
	if ok {
		return convertSystem_GrpcCredentials_Entity_VersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcCredentials_Entity_VersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcCredentials_Entity_VersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcCredentials_Entity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcCredentials_Entity", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcCredentials_Entity_VersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a ONCE subscription.
func (n *System_GrpcCredentials_Entity_VersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_VersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_VersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcCredentials_Entity", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcCredentials_Entity_VersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_VersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_VersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcCredentials_Entity_VersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-credentials/entities/entity/state/version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/version to the batch object.
func (n *System_GrpcCredentials_Entity_VersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcCredentials_Entity_VersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcCredentials_Entity_VersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcCredentials_Entity{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcCredentials_Entity{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcCredentials_Entity", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcCredentials_Entity_VersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-credentials/entities/entity/state/version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcCredentials_Entity_VersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcCredentials_Entity_VersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-credentials/entities/entity/state/version to the batch object.
func (n *System_GrpcCredentials_Entity_VersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcCredentials_Entity_VersionPath extracts the value of the leaf Version from its parent oc.System_GrpcCredentials_Entity
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcCredentials_Entity_VersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcCredentials_Entity) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Version
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_GrpcServer {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServerPath) Get(t testing.TB) *oc.System_GrpcServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_GrpcServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_GrpcServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server with a ONCE subscription.
func (n *System_GrpcServerPathAny) Get(t testing.TB) []*oc.System_GrpcServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_GrpcServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcServer {
	t.Helper()
	c := &oc.CollectionSystem_GrpcServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcServer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_GrpcServer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_GrpcServer)))
		return false
	})
	return c
}

func watch_System_GrpcServerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	w := &oc.System_GrpcServerWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_GrpcServer{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcServer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	return watch_System_GrpcServerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServerPath) Await(t testing.TB, timeout time.Duration, val *oc.System_GrpcServer) *oc.QualifiedSystem_GrpcServer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_GrpcServer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server to the batch object.
func (n *System_GrpcServerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_GrpcServer {
	t.Helper()
	c := &oc.CollectionSystem_GrpcServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_GrpcServer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	w := &oc.System_GrpcServerWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_GrpcServer{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_GrpcServer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_GrpcServer) bool) *oc.System_GrpcServerWatcher {
	t.Helper()
	return watch_System_GrpcServerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server to the batch object.
func (n *System_GrpcServerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_AuthzPolicyCreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_AuthzPolicyCreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a ONCE subscription.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_AuthzPolicyCreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_AuthzPolicyCreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_AuthzPolicyCreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on to the batch object.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_AuthzPolicyCreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_AuthzPolicyCreatedOnPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_AuthzPolicyCreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-created-on to the batch object.
func (n *System_GrpcServer_AuthzPolicyCreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_AuthzPolicyCreatedOnPath extracts the value of the leaf AuthzPolicyCreatedOn from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_GrpcServer_AuthzPolicyCreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.AuthzPolicyCreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_AuthzPolicyVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_AuthzPolicyVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_AuthzPolicyVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_AuthzPolicyVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_AuthzPolicyVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a ONCE subscription.
func (n *System_GrpcServer_AuthzPolicyVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_AuthzPolicyVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_AuthzPolicyVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_AuthzPolicyVersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_AuthzPolicyVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_AuthzPolicyVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_AuthzPolicyVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version to the batch object.
func (n *System_GrpcServer_AuthzPolicyVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_AuthzPolicyVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_AuthzPolicyVersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_AuthzPolicyVersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_AuthzPolicyVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_AuthzPolicyVersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/authz-policy-version to the batch object.
func (n *System_GrpcServer_AuthzPolicyVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_AuthzPolicyVersionPath extracts the value of the leaf AuthzPolicyVersion from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_AuthzPolicyVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AuthzPolicyVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_CaTrustBundleCreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CaTrustBundleCreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a ONCE subscription.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CaTrustBundleCreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_CaTrustBundleCreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_CaTrustBundleCreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on to the batch object.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CaTrustBundleCreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_CaTrustBundleCreatedOnPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_CaTrustBundleCreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-created-on to the batch object.
func (n *System_GrpcServer_CaTrustBundleCreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_CaTrustBundleCreatedOnPath extracts the value of the leaf CaTrustBundleCreatedOn from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_GrpcServer_CaTrustBundleCreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CaTrustBundleCreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CaTrustBundleIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_CaTrustBundleIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CaTrustBundleIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CaTrustBundleIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CaTrustBundleIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a ONCE subscription.
func (n *System_GrpcServer_CaTrustBundleIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CaTrustBundleIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CaTrustBundleIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_CaTrustBundleIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CaTrustBundleIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CaTrustBundleIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_CaTrustBundleIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id to the batch object.
func (n *System_GrpcServer_CaTrustBundleIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CaTrustBundleIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CaTrustBundleIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_CaTrustBundleIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CaTrustBundleIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CaTrustBundleIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-id to the batch object.
func (n *System_GrpcServer_CaTrustBundleIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_CaTrustBundleIdPath extracts the value of the leaf CaTrustBundleId from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_CaTrustBundleIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CaTrustBundleId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CaTrustBundleVersionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_CaTrustBundleVersionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CaTrustBundleVersionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CaTrustBundleVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CaTrustBundleVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a ONCE subscription.
func (n *System_GrpcServer_CaTrustBundleVersionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CaTrustBundleVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CaTrustBundleVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_CaTrustBundleVersionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CaTrustBundleVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CaTrustBundleVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_CaTrustBundleVersionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version to the batch object.
func (n *System_GrpcServer_CaTrustBundleVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CaTrustBundleVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CaTrustBundleVersionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_CaTrustBundleVersionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CaTrustBundleVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CaTrustBundleVersionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/ca-trust-bundle-version to the batch object.
func (n *System_GrpcServer_CaTrustBundleVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_CaTrustBundleVersionPath extracts the value of the leaf CaTrustBundleVersion from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_CaTrustBundleVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CaTrustBundleVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CertificateCreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_CertificateCreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CertificateCreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CertificateCreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CertificateCreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a ONCE subscription.
func (n *System_GrpcServer_CertificateCreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateCreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateCreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_CertificateCreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateCreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateCreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_CertificateCreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on to the batch object.
func (n *System_GrpcServer_CertificateCreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateCreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateCreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_CertificateCreatedOnPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateCreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateCreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-created-on to the batch object.
func (n *System_GrpcServer_CertificateCreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_CertificateCreatedOnPath extracts the value of the leaf CertificateCreatedOn from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_GrpcServer_CertificateCreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CertificateCreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CertificateIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_CertificateIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CertificateIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CertificateIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CertificateIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a ONCE subscription.
func (n *System_GrpcServer_CertificateIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_CertificateIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_CertificateIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id to the batch object.
func (n *System_GrpcServer_CertificateIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_CertificateIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-id to the batch object.
func (n *System_GrpcServer_CertificateIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_CertificateIdPath extracts the value of the leaf CertificateId from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_GrpcServer_CertificateIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CertificateId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_CertificateRevocationListBundleCreatedOnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_CertificateRevocationListBundleCreatedOnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a ONCE subscription.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateRevocationListBundleCreatedOnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_CertificateRevocationListBundleCreatedOnPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateRevocationListBundleCreatedOnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on to the batch object.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_CertificateRevocationListBundleCreatedOnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_GrpcServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_GrpcServer", structs[pre], queryPath, true, false)
			qv := convertSystem_GrpcServer_CertificateRevocationListBundleCreatedOnPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_GrpcServer_CertificateRevocationListBundleCreatedOnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/certificate-revocation-list-bundle-created-on to the batch object.
func (n *System_GrpcServer_CertificateRevocationListBundleCreatedOnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_CertificateRevocationListBundleCreatedOnPath extracts the value of the leaf CertificateRevocationListBundleCreatedOn from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_GrpcServer_CertificateRevocationListBundleCreatedOnPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CertificateRevocationListBundleCreatedOn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
