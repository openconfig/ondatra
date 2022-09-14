package terminaldevice

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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui-subtype to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath extracts the value of the leaf OuiSubtype from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiSubtypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath) Lookup(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePathAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	w := &oc.Int32Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath) Await(t testing.TB, timeout time.Duration, val int32) *oc.QualifiedInt32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt32 {
	t.Helper()
	c := &oc.CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	w := &oc.Int32Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt32) bool) *oc.Int32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath extracts the value of the leaf Type from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedInt32.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) *oc.QualifiedInt32 {
	t.Helper()
	qv := &oc.QualifiedInt32{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath) Lookup(t testing.TB) *oc.QualifiedBinary {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath) Get(t testing.TB) oc.Binary {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePathAny) Lookup(t testing.TB) []*oc.QualifiedBinary {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBinary
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePathAny) Get(t testing.TB) []oc.Binary {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBinary {
	t.Helper()
	c := &oc.CollectionBinary{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBinary) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	w := &oc.BinaryWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBinary)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath) Await(t testing.TB, timeout time.Duration, val oc.Binary) *oc.QualifiedBinary {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBinary) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBinary {
	t.Helper()
	c := &oc.CollectionBinary{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBinary) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	w := &oc.BinaryWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBinary)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBinary) bool) *oc.BinaryWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/value to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath extracts the value of the leaf Value from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedBinary.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_ValuePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) *oc.QualifiedBinary {
	t.Helper()
	qv := &oc.QualifiedBinary{
		Metadata: md,
	}
	val := parent.Value
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/ttl to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath extracts the value of the leaf Ttl from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_TtlPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Ttl
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetSnooping())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Lldp{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/snooping to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath extracts the value of the leaf Snooping from its parent oc.TerminalDevice_Channel_Ethernet_Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Snooping
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_OutBlockErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_OutBlockErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutBlockErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_OutBlockErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutBlockErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutBlockErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_OutBlockErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutBlockErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-block-errors to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutBlockErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_OutBlockErrorsPath extracts the value of the leaf OutBlockErrors from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_OutBlockErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutBlockErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_OutCrcErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_OutCrcErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutCrcErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_OutCrcErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutCrcErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutCrcErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_OutCrcErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutCrcErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-crc-errors to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutCrcErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_OutCrcErrorsPath extracts the value of the leaf OutCrcErrors from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_OutCrcErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutCrcErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_OutMacControlFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_OutMacControlFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutMacControlFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_OutMacControlFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutMacControlFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutMacControlFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_OutMacControlFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutMacControlFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-control-frames to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutMacControlFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_OutMacControlFramesPath extracts the value of the leaf OutMacControlFrames from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_OutMacControlFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutMacControlFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_OutMacErrorsTxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_OutMacErrorsTxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_OutMacErrorsTxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutMacErrorsTxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_OutMacErrorsTxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutMacErrorsTxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-errors-tx to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutMacErrorsTxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_OutMacErrorsTxPath extracts the value of the leaf OutMacErrorsTx from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_OutMacErrorsTxPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutMacErrorsTx
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_OutMacPauseFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_OutMacPauseFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_OutMacPauseFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutMacPauseFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_OutMacPauseFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutMacPauseFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-mac-pause-frames to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutMacPauseFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_OutMacPauseFramesPath extracts the value of the leaf OutMacPauseFrames from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_OutMacPauseFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutMacPauseFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-pcs-bip-errors to the batch object.
func (n *TerminalDevice_Channel_Ethernet_OutPcsBipErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath extracts the value of the leaf OutPcsBipErrors from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_OutPcsBipErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutPcsBipErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Out_8021QFramesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Out_8021QFramesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Out_8021QFramesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Out_8021QFramesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Out_8021QFramesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Out_8021QFramesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Out_8021QFramesPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Out_8021QFramesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/out-8021q-frames to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Out_8021QFramesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Out_8021QFramesPath extracts the value of the leaf Out_8021QFrames from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_Out_8021QFramesPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Out_8021QFrames
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ethernet_PostFecBer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ethernet_PostFecBer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ethernet_PostFecBer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_PostFecBer {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_PostFecBer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Ethernet_PostFecBer)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer) bool) *oc.TerminalDevice_Channel_Ethernet_PostFecBerWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_PostFecBerWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer) bool) *oc.TerminalDevice_Channel_Ethernet_PostFecBerWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Ethernet_PostFecBer) *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_PostFecBer {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_PostFecBer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer) bool) *oc.TerminalDevice_Channel_Ethernet_PostFecBerWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_PostFecBerWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_PostFecBer) bool) *oc.TerminalDevice_Channel_Ethernet_PostFecBerWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_PostFecBer_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_PostFecBer_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_AvgPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/avg to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_PostFecBer_AvgPath extracts the value of the leaf Avg from its parent oc.TerminalDevice_Channel_Ethernet_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Ethernet_PostFecBer_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_PostFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_PostFecBer_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_PostFecBer_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_InstantPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/instant to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_PostFecBer_InstantPath extracts the value of the leaf Instant from its parent oc.TerminalDevice_Channel_Ethernet_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Ethernet_PostFecBer_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_PostFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Instant
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/interval to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath extracts the value of the leaf Interval from its parent oc.TerminalDevice_Channel_Ethernet_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_PostFecBer_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_PostFecBer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxPath extracts the value of the leaf Max from its parent oc.TerminalDevice_Channel_Ethernet_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_PostFecBer) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_PostFecBer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_PostFecBer", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/state/post-fec-ber/max-time to the batch object.
func (n *TerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath extracts the value of the leaf MaxTime from its parent oc.TerminalDevice_Channel_Ethernet_PostFecBer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_PostFecBer_MaxTimePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_PostFecBer) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.MaxTime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
