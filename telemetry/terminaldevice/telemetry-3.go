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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Counters", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Counters", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-discard to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath extracts the value of the leaf TlvDiscard from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvDiscardPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TlvDiscard
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Counters", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Counters", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ethernet_Lldp_Counters{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Counters", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/counters/tlv-unknown to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath extracts the value of the leaf TlvUnknown from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_Lldp_Counters_TlvUnknownPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.TlvUnknown
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
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
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/state/enabled to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath extracts the value of the leaf Enabled from its parent oc.TerminalDevice_Channel_Ethernet_Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_NeighborWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_Lldp_NeighborWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_NeighborWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_NeighborPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_NeighborWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_Lldp_NeighborWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_NeighborWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_NeighborPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/age to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath extracts the value of the leaf Age from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_AgePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.Age
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath extracts the value of the leaf ChassisId from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath) Lookup(t testing.TB) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath) Get(t testing.TB) oc.E_LldpTypes_ChassisIdType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpTypes_ChassisIdType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePathAny) Get(t testing.TB) []oc.E_LldpTypes_ChassisIdType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_LldpTypes_ChassisIdType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_ChassisIdTypeWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_ChassisIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_LldpTypes_ChassisIdType) *oc.QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_ChassisIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_ChassisIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_ChassisIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_ChassisIdTypeWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_ChassisIdType) bool) *oc.E_LldpTypes_ChassisIdTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/chassis-id-type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath extracts the value of the leaf ChassisIdType from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpTypes_ChassisIdType.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ChassisIdTypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedE_LldpTypes_ChassisIdType {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/id to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath extracts the value of the leaf Id from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInt64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath) Await(t testing.TB, timeout time.Duration, val int64) *oc.QualifiedInt64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInt64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInt64 {
	t.Helper()
	c := &oc.CollectionInt64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInt64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	w := &oc.Int64Watcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInt64) bool) *oc.Int64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/last-update to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath extracts the value of the leaf LastUpdate from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_LastUpdatePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedInt64 {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath extracts the value of the leaf ManagementAddress from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/management-address-type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath extracts the value of the leaf ManagementAddressType from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_ManagementAddressTypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-description to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath extracts the value of the leaf PortDescription from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortDescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath extracts the value of the leaf PortId from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath) Lookup(t testing.TB) *oc.QualifiedE_LldpTypes_PortIdType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath) Get(t testing.TB) oc.E_LldpTypes_PortIdType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_LldpTypes_PortIdType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_LldpTypes_PortIdType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePathAny) Get(t testing.TB) []oc.E_LldpTypes_PortIdType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_LldpTypes_PortIdType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_PortIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_PortIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_PortIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_PortIdType) bool) *oc.E_LldpTypes_PortIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_PortIdTypeWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_PortIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_PortIdType) bool) *oc.E_LldpTypes_PortIdTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_LldpTypes_PortIdType) *oc.QualifiedE_LldpTypes_PortIdType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_LldpTypes_PortIdType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_LldpTypes_PortIdType {
	t.Helper()
	c := &oc.CollectionE_LldpTypes_PortIdType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_LldpTypes_PortIdType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_PortIdType) bool) *oc.E_LldpTypes_PortIdTypeWatcher {
	t.Helper()
	w := &oc.E_LldpTypes_PortIdTypeWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_LldpTypes_PortIdType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_LldpTypes_PortIdType) bool) *oc.E_LldpTypes_PortIdTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/port-id-type to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath extracts the value of the leaf PortIdType from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_LldpTypes_PortIdType.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_PortIdTypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedE_LldpTypes_PortIdType {
	t.Helper()
	qv := &oc.QualifiedE_LldpTypes_PortIdType{
		Metadata: md,
	}
	val := parent.PortIdType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-description to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath extracts the value of the leaf SystemDescription from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemDescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/state/system-name to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath extracts the value of the leaf SystemName from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_SystemNamePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) bool) *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_TlvPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/neighbors/neighbor/custom-tlvs/tlv/state/oui to the batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath extracts the value of the leaf Oui from its parent oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv_OuiPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp_Neighbor_Tlv) *oc.QualifiedString {
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
