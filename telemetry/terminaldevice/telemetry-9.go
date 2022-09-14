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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_TtiMsgRecvPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_TtiMsgRecvPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_TtiMsgRecvPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_TtiMsgRecvPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_TtiMsgRecvPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv to the batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_TtiMsgRecvPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Otn_TtiMsgRecvPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_TtiMsgRecvPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-recv to the batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgRecvPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_TtiMsgRecvPath extracts the value of the leaf TtiMsgRecv from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Otn_TtiMsgRecvPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TtiMsgRecv
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_TtiMsgTransmitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_TtiMsgTransmitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit to the batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/tti-msg-transmit to the batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath extracts the value of the leaf TtiMsgTransmit from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TtiMsgTransmit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_UnavailableSecondsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_UnavailableSecondsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_UnavailableSecondsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_UnavailableSecondsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_UnavailableSecondsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds to the batch object.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_UnavailableSecondsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Otn_UnavailableSecondsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_UnavailableSecondsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/unavailable-seconds to the batch object.
func (n *TerminalDevice_Channel_Otn_UnavailableSecondsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_UnavailableSecondsPath extracts the value of the leaf UnavailableSeconds from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Otn_UnavailableSecondsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.UnavailableSeconds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_RateClassPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_RateClassPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_RateClassPath) Get(t testing.TB) oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_RateClassPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_RateClassPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a ONCE subscription.
func (n *TerminalDevice_Channel_RateClassPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_RateClassPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_RateClassPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_RateClassPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_RateClassPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_RateClassPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_RateClassPath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class to the batch object.
func (n *TerminalDevice_Channel_RateClassPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_RateClassPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_RateClassPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_RateClassPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_RateClassPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_RateClassPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/rate-class to the batch object.
func (n *TerminalDevice_Channel_RateClassPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_RateClassPath extracts the value of the leaf RateClass from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE.
func convertTerminalDevice_Channel_RateClassPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE{
		Metadata: md,
	}
	val := parent.RateClass
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_TestSignalPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_TestSignalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_TestSignalPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_TestSignalPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_TestSignalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a ONCE subscription.
func (n *TerminalDevice_Channel_TestSignalPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_TestSignalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_TestSignalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_TestSignalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_TestSignalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_TestSignalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_TestSignalPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal to the batch object.
func (n *TerminalDevice_Channel_TestSignalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_TestSignalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_TestSignalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_TestSignalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_TestSignalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_TestSignalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/test-signal to the batch object.
func (n *TerminalDevice_Channel_TestSignalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_TestSignalPath extracts the value of the leaf TestSignal from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertTerminalDevice_Channel_TestSignalPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.TestSignal
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_TribProtocolPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_TribProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_TribProtocolPath) Get(t testing.TB) oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_TribProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_TribProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a ONCE subscription.
func (n *TerminalDevice_Channel_TribProtocolPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_TribProtocolPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_TribProtocolPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_TribProtocolPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_TribProtocolPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_TribProtocolPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_TribProtocolPath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol to the batch object.
func (n *TerminalDevice_Channel_TribProtocolPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_TribProtocolPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_TribProtocolPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_TribProtocolPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_TribProtocolPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_TribProtocolPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/trib-protocol to the batch object.
func (n *TerminalDevice_Channel_TribProtocolPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_TribProtocolPath extracts the value of the leaf TribProtocol from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE.
func convertTerminalDevice_Channel_TribProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE{
		Metadata: md,
	}
	val := parent.TribProtocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_ModePath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Mode {
	t.Helper()
	goStruct := &oc.TerminalDevice_Mode{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Mode", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Mode{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_ModePath) Get(t testing.TB) *oc.TerminalDevice_Mode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_ModePathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Mode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Mode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Mode{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Mode", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Mode{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode with a ONCE subscription.
func (n *TerminalDevice_ModePathAny) Get(t testing.TB) []*oc.TerminalDevice_Mode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Mode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_ModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Mode {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Mode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Mode) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Mode{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Mode)))
		return false
	})
	return c
}

func watch_TerminalDevice_ModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Mode) bool) *oc.TerminalDevice_ModeWatcher {
	t.Helper()
	w := &oc.TerminalDevice_ModeWatcher{}
	gs := &oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Mode", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Mode{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Mode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_ModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Mode) bool) *oc.TerminalDevice_ModeWatcher {
	t.Helper()
	return watch_TerminalDevice_ModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/operational-modes/mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_ModePath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Mode) *oc.QualifiedTerminalDevice_Mode {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Mode) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/operational-modes/mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode to the batch object.
func (n *TerminalDevice_ModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_ModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Mode {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Mode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Mode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_ModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Mode) bool) *oc.TerminalDevice_ModeWatcher {
	t.Helper()
	w := &oc.TerminalDevice_ModeWatcher{}
	structs := map[string]*oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Mode{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Mode", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Mode{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Mode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_ModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Mode) bool) *oc.TerminalDevice_ModeWatcher {
	t.Helper()
	return watch_TerminalDevice_ModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode to the batch object.
func (n *TerminalDevice_ModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Mode_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Mode{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Mode", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Mode_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Mode_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Mode_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Mode{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Mode", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Mode_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a ONCE subscription.
func (n *TerminalDevice_Mode_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Mode_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Mode_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Mode", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Mode_DescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Mode_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Mode_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Mode_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description to the batch object.
func (n *TerminalDevice_Mode_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Mode_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Mode_DescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Mode{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Mode", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Mode_DescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Mode_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Mode_DescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode/state/description to the batch object.
func (n *TerminalDevice_Mode_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Mode_DescriptionPath extracts the value of the leaf Description from its parent oc.TerminalDevice_Mode
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Mode_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Mode) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Mode_ModeIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Mode{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Mode", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Mode_ModeIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Mode_ModeIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Mode_ModeIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Mode{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Mode", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Mode_ModeIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a ONCE subscription.
func (n *TerminalDevice_Mode_ModeIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Mode_ModeIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Mode_ModeIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Mode", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Mode_ModeIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Mode_ModeIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_TerminalDevice_Mode_ModeIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Mode_ModeIdPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id to the batch object.
func (n *TerminalDevice_Mode_ModeIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Mode_ModeIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Mode_ModeIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Mode{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Mode", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Mode_ModeIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Mode_ModeIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_TerminalDevice_Mode_ModeIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode/state/mode-id to the batch object.
func (n *TerminalDevice_Mode_ModeIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Mode_ModeIdPath extracts the value of the leaf ModeId from its parent oc.TerminalDevice_Mode
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertTerminalDevice_Mode_ModeIdPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Mode) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.ModeId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Mode_VendorIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Mode{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Mode", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Mode_VendorIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Mode_VendorIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Mode_VendorIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Mode{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Mode", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Mode_VendorIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a ONCE subscription.
func (n *TerminalDevice_Mode_VendorIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Mode_VendorIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Mode_VendorIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Mode", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Mode_VendorIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Mode_VendorIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Mode_VendorIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Mode_VendorIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id to the batch object.
func (n *TerminalDevice_Mode_VendorIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Mode_VendorIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Mode_VendorIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Mode{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Mode{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Mode", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Mode_VendorIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Mode_VendorIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Mode_VendorIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/operational-modes/mode/state/vendor-id to the batch object.
func (n *TerminalDevice_Mode_VendorIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Mode_VendorIdPath extracts the value of the leaf VendorId from its parent oc.TerminalDevice_Mode
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Mode_VendorIdPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Mode) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.VendorId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
