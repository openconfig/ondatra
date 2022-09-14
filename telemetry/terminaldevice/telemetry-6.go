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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_IndexPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a ONCE subscription.
func (n *TerminalDevice_Channel_IndexPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_IndexPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_IndexPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_IndexPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_IndexPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_IndexPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_IndexPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index to the batch object.
func (n *TerminalDevice_Channel_IndexPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_IndexPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_IndexPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
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
			qv := convertTerminalDevice_Channel_IndexPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_IndexPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_IndexPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/index to the batch object.
func (n *TerminalDevice_Channel_IndexPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_IndexPath extracts the value of the leaf Index from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_IngressPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ingress {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ingress{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_IngressPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ingress {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_IngressPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ingress {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ingress
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ingress{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription.
func (n *TerminalDevice_Channel_IngressPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ingress {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ingress
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_IngressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ingress {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ingress{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ingress) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Ingress{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Ingress)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_IngressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ingress) bool) *oc.TerminalDevice_Channel_IngressWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_IngressWatcher{}
	gs := &oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ingress", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Ingress{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ingress)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_IngressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ingress) bool) *oc.TerminalDevice_Channel_IngressWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_IngressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_IngressPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Ingress) *oc.QualifiedTerminalDevice_Channel_Ingress {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Ingress) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress to the batch object.
func (n *TerminalDevice_Channel_IngressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_IngressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Ingress {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Ingress{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Ingress) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_IngressPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ingress) bool) *oc.TerminalDevice_Channel_IngressWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_IngressWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ingress{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ingress", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Ingress{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Ingress)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_IngressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Ingress) bool) *oc.TerminalDevice_Channel_IngressWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_IngressPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress to the batch object.
func (n *TerminalDevice_Channel_IngressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ingress_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ingress_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ingress_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a ONCE subscription.
func (n *TerminalDevice_Channel_Ingress_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ingress_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ingress", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ingress_InterfacePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ingress_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface to the batch object.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ingress_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ingress_InterfacePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ingress{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ingress", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ingress_InterfacePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ingress_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ingress_InterfacePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/interface to the batch object.
func (n *TerminalDevice_Channel_Ingress_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ingress_InterfacePath extracts the value of the leaf Interface from its parent oc.TerminalDevice_Channel_Ingress
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ingress_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ingress) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a ONCE subscription.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ingress_PhysicalChannelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	gs := &oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ingress", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ingress_PhysicalChannelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Await(t testing.TB, timeout time.Duration, val []uint16) *oc.QualifiedUint16Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel to the batch object.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16Slice {
	t.Helper()
	c := &oc.CollectionUint16Slice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ingress_PhysicalChannelPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	w := &oc.Uint16SliceWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ingress{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ingress", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16Slice) bool) *oc.Uint16SliceWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ingress_PhysicalChannelPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/physical-channel to the batch object.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ingress_PhysicalChannelPath extracts the value of the leaf PhysicalChannel from its parent oc.TerminalDevice_Channel_Ingress
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ingress) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.PhysicalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Ingress_TransceiverPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ingress_TransceiverPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ingress_TransceiverPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a ONCE subscription.
func (n *TerminalDevice_Channel_Ingress_TransceiverPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ingress_TransceiverPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Ingress", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Ingress_TransceiverPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ingress_TransceiverPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver to the batch object.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Ingress_TransceiverPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Ingress_TransceiverPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Ingress{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Ingress{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Ingress", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Ingress_TransceiverPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Ingress_TransceiverPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Ingress_TransceiverPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/state/transceiver to the batch object.
func (n *TerminalDevice_Channel_Ingress_TransceiverPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Ingress_TransceiverPath extracts the value of the leaf Transceiver from its parent oc.TerminalDevice_Channel_Ingress
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ingress_TransceiverPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ingress) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Transceiver
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_LinkStatePath) Lookup(t testing.TB) *oc.QualifiedE_Channel_LinkState {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_LinkStatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_LinkStatePath) Get(t testing.TB) oc.E_Channel_LinkState {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_LinkStatePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Channel_LinkState {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Channel_LinkState
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_LinkStatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a ONCE subscription.
func (n *TerminalDevice_Channel_LinkStatePathAny) Get(t testing.TB) []oc.E_Channel_LinkState {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Channel_LinkState
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_LinkStatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Channel_LinkState {
	t.Helper()
	c := &oc.CollectionE_Channel_LinkState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Channel_LinkState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_LinkStatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Channel_LinkState) bool) *oc.E_Channel_LinkStateWatcher {
	t.Helper()
	w := &oc.E_Channel_LinkStateWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_LinkStatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Channel_LinkState)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_LinkStatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Channel_LinkState) bool) *oc.E_Channel_LinkStateWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_LinkStatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_LinkStatePath) Await(t testing.TB, timeout time.Duration, val oc.E_Channel_LinkState) *oc.QualifiedE_Channel_LinkState {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Channel_LinkState) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state to the batch object.
func (n *TerminalDevice_Channel_LinkStatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_LinkStatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Channel_LinkState {
	t.Helper()
	c := &oc.CollectionE_Channel_LinkState{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Channel_LinkState) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_LinkStatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Channel_LinkState) bool) *oc.E_Channel_LinkStateWatcher {
	t.Helper()
	w := &oc.E_Channel_LinkStateWatcher{}
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
			qv := convertTerminalDevice_Channel_LinkStatePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Channel_LinkState)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_LinkStatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Channel_LinkState) bool) *oc.E_Channel_LinkStateWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_LinkStatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/link-state to the batch object.
func (n *TerminalDevice_Channel_LinkStatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_LinkStatePath extracts the value of the leaf LinkState from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Channel_LinkState.
func convertTerminalDevice_Channel_LinkStatePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_Channel_LinkState {
	t.Helper()
	qv := &oc.QualifiedE_Channel_LinkState{
		Metadata: md,
	}
	val := parent.LinkState
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_LogicalChannelTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Get(t testing.TB) oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_LogicalChannelTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_LogicalChannelTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a ONCE subscription.
func (n *TerminalDevice_Channel_LogicalChannelTypePathAny) Get(t testing.TB) []oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_LogicalChannelTypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPEWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_LogicalChannelTypePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_LogicalChannelTypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type to the batch object.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_LogicalChannelTypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_LogicalChannelTypePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPEWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPEWatcher{}
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
			qv := convertTerminalDevice_Channel_LogicalChannelTypePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_LogicalChannelTypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) bool) *oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPEWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_LogicalChannelTypePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/logical-channel-type to the batch object.
func (n *TerminalDevice_Channel_LogicalChannelTypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_LogicalChannelTypePath extracts the value of the leaf LogicalChannelType from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE.
func convertTerminalDevice_Channel_LogicalChannelTypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE{
		Metadata: md,
	}
	val := parent.LogicalChannelType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_LoopbackModePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_LoopbackModePath(t, md, goStruct)
	}
	return (&oc.QualifiedE_TransportTypes_LoopbackModeType{
		Metadata: md,
	}).SetVal(goStruct.GetLoopbackMode())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_LoopbackModePath) Get(t testing.TB) oc.E_TransportTypes_LoopbackModeType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_LoopbackModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_LoopbackModeType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_LoopbackModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a ONCE subscription.
func (n *TerminalDevice_Channel_LoopbackModePathAny) Get(t testing.TB) []oc.E_TransportTypes_LoopbackModeType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_LoopbackModeType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_LoopbackModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_LoopbackModeType {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_LoopbackModeType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_LoopbackModeType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_LoopbackModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LoopbackModeType) bool) *oc.E_TransportTypes_LoopbackModeTypeWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_LoopbackModeTypeWatcher{}
	gs := &oc.TerminalDevice_Channel{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_LoopbackModePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_LoopbackModeType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_LoopbackModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LoopbackModeType) bool) *oc.E_TransportTypes_LoopbackModeTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_LoopbackModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_LoopbackModePath) Await(t testing.TB, timeout time.Duration, val oc.E_TransportTypes_LoopbackModeType) *oc.QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_TransportTypes_LoopbackModeType) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode to the batch object.
func (n *TerminalDevice_Channel_LoopbackModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_LoopbackModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_TransportTypes_LoopbackModeType {
	t.Helper()
	c := &oc.CollectionE_TransportTypes_LoopbackModeType{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_TransportTypes_LoopbackModeType) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_LoopbackModePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LoopbackModeType) bool) *oc.E_TransportTypes_LoopbackModeTypeWatcher {
	t.Helper()
	w := &oc.E_TransportTypes_LoopbackModeTypeWatcher{}
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
			qv := convertTerminalDevice_Channel_LoopbackModePath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_TransportTypes_LoopbackModeType)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_LoopbackModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_TransportTypes_LoopbackModeType) bool) *oc.E_TransportTypes_LoopbackModeTypeWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_LoopbackModePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/state/loopback-mode to the batch object.
func (n *TerminalDevice_Channel_LoopbackModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_LoopbackModePath extracts the value of the leaf LoopbackMode from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_LoopbackModeType.
func convertTerminalDevice_Channel_LoopbackModePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_LoopbackModeType{
		Metadata: md,
	}
	val := parent.LoopbackMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_OtnPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Otn {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Otn{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_OtnPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Otn {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_OtnPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Otn {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Otn
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Otn{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription.
func (n *TerminalDevice_Channel_OtnPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Otn {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Otn
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_OtnPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Otn {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Otn{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Otn) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Otn{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Otn)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_OtnPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn) bool) *oc.TerminalDevice_Channel_OtnWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_OtnWatcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Otn{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Otn)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_OtnPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn) bool) *oc.TerminalDevice_Channel_OtnWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_OtnPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_OtnPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Otn) *oc.QualifiedTerminalDevice_Channel_Otn {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Otn) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn to the batch object.
func (n *TerminalDevice_Channel_OtnPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_OtnPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Otn {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Otn{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Otn) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_OtnPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn) bool) *oc.TerminalDevice_Channel_OtnWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_OtnWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
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
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Otn{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Otn)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_OtnPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn) bool) *oc.TerminalDevice_Channel_OtnWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_OtnPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn to the batch object.
func (n *TerminalDevice_Channel_OtnPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_BackgroundBlockErrorsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertTerminalDevice_Channel_Otn_BackgroundBlockErrorsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_BackgroundBlockErrorsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors to the batch object.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_BackgroundBlockErrorsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertTerminalDevice_Channel_Otn_BackgroundBlockErrorsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_BackgroundBlockErrorsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/background-block-errors to the batch object.
func (n *TerminalDevice_Channel_Otn_BackgroundBlockErrorsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_BackgroundBlockErrorsPath extracts the value of the leaf BackgroundBlockErrors from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Otn_BackgroundBlockErrorsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.BackgroundBlockErrors
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_CodeViolationsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertTerminalDevice_Channel_Otn_CodeViolationsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_CodeViolationsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_CodeViolationsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_CodeViolationsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations to the batch object.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_CodeViolationsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertTerminalDevice_Channel_Otn_CodeViolationsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_CodeViolationsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/code-violations to the batch object.
func (n *TerminalDevice_Channel_Otn_CodeViolationsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_CodeViolationsPath extracts the value of the leaf CodeViolations from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Otn_CodeViolationsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.CodeViolations
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_ErroredBlocksPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertTerminalDevice_Channel_Otn_ErroredBlocksPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_ErroredBlocksPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_ErroredBlocksPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_ErroredBlocksPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks to the batch object.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_ErroredBlocksPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertTerminalDevice_Channel_Otn_ErroredBlocksPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_ErroredBlocksPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-blocks to the batch object.
func (n *TerminalDevice_Channel_Otn_ErroredBlocksPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_ErroredBlocksPath extracts the value of the leaf ErroredBlocks from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Otn_ErroredBlocksPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ErroredBlocks
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_ErroredSecondsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
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
		qv := convertTerminalDevice_Channel_Otn_ErroredSecondsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_ErroredSecondsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_ErroredSecondsPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_ErroredSecondsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds to the batch object.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_ErroredSecondsPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
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
			qv := convertTerminalDevice_Channel_Otn_ErroredSecondsPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_ErroredSecondsPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/errored-seconds to the batch object.
func (n *TerminalDevice_Channel_Otn_ErroredSecondsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_ErroredSecondsPath extracts the value of the leaf ErroredSeconds from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Otn_ErroredSecondsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ErroredSeconds
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_EsnrPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Otn_Esnr {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn_Esnr", goStruct, false, false)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Otn_Esnr{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_EsnrPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Otn_Esnr {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_EsnrPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Otn_Esnr {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Otn_Esnr
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Otn_Esnr{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_EsnrPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Otn_Esnr {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Otn_Esnr
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_EsnrPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Otn_Esnr {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Otn_Esnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Otn_Esnr) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedTerminalDevice_Channel_Otn_Esnr{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.TerminalDevice_Channel_Otn_Esnr)))
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_EsnrPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn_Esnr) bool) *oc.TerminalDevice_Channel_Otn_EsnrWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Otn_EsnrWatcher{}
	gs := &oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", gs, queryPath, false, false)
		qv := (&oc.QualifiedTerminalDevice_Channel_Otn_Esnr{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Otn_Esnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_EsnrPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn_Esnr) bool) *oc.TerminalDevice_Channel_Otn_EsnrWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_EsnrPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_EsnrPath) Await(t testing.TB, timeout time.Duration, val *oc.TerminalDevice_Channel_Otn_Esnr) *oc.QualifiedTerminalDevice_Channel_Otn_Esnr {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedTerminalDevice_Channel_Otn_Esnr) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr to the batch object.
func (n *TerminalDevice_Channel_Otn_EsnrPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_EsnrPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionTerminalDevice_Channel_Otn_Esnr {
	t.Helper()
	c := &oc.CollectionTerminalDevice_Channel_Otn_Esnr{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedTerminalDevice_Channel_Otn_Esnr) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_EsnrPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn_Esnr) bool) *oc.TerminalDevice_Channel_Otn_EsnrWatcher {
	t.Helper()
	w := &oc.TerminalDevice_Channel_Otn_EsnrWatcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedTerminalDevice_Channel_Otn_Esnr{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedTerminalDevice_Channel_Otn_Esnr)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_EsnrPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedTerminalDevice_Channel_Otn_Esnr) bool) *oc.TerminalDevice_Channel_Otn_EsnrWatcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_EsnrPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr to the batch object.
func (n *TerminalDevice_Channel_Otn_EsnrPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn_Esnr", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_Esnr_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_Esnr_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_AvgPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_Esnr_AvgPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_AvgPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_AvgPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Otn_Esnr_AvgPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_AvgPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/avg to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_AvgPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_Esnr_AvgPath extracts the value of the leaf Avg from its parent oc.TerminalDevice_Channel_Otn_Esnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Otn_Esnr_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn_Esnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn_Esnr", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_Esnr_InstantPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_Esnr_InstantPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_InstantPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_Esnr_InstantPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_InstantPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_InstantPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Otn_Esnr_InstantPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_InstantPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/instant to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_InstantPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_Esnr_InstantPath extracts the value of the leaf Instant from its parent oc.TerminalDevice_Channel_Otn_Esnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Otn_Esnr_InstantPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn_Esnr) *oc.QualifiedFloat64 {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn_Esnr", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_Esnr_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_Esnr_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_IntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_Esnr_IntervalPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_IntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_IntervalPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Otn_Esnr_IntervalPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_IntervalPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/interval to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_IntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_Esnr_IntervalPath extracts the value of the leaf Interval from its parent oc.TerminalDevice_Channel_Otn_Esnr
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertTerminalDevice_Channel_Otn_Esnr_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn_Esnr) *oc.QualifiedUint64 {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn_Esnr", goStruct, true, false)
	if ok {
		return convertTerminalDevice_Channel_Otn_Esnr_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn_Esnr{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_Esnr_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_MaxPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	gs := &oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertTerminalDevice_Channel_Otn_Esnr_MaxPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedFloat64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_MaxPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPath) Await(t testing.TB, timeout time.Duration, val float64) *oc.QualifiedFloat64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedFloat64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionFloat64 {
	t.Helper()
	c := &oc.CollectionFloat64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedFloat64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_TerminalDevice_Channel_Otn_Esnr_MaxPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	w := &oc.Float64Watcher{}
	structs := map[string]*oc.TerminalDevice_Channel_Otn_Esnr{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.TerminalDevice_Channel_Otn_Esnr{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "TerminalDevice_Channel_Otn_Esnr", structs[pre], queryPath, true, false)
			qv := convertTerminalDevice_Channel_Otn_Esnr_MaxPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedFloat64) bool) *oc.Float64Watcher {
	t.Helper()
	return watch_TerminalDevice_Channel_Otn_Esnr_MaxPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/state/esnr/max to the batch object.
func (n *TerminalDevice_Channel_Otn_Esnr_MaxPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertTerminalDevice_Channel_Otn_Esnr_MaxPath extracts the value of the leaf Max from its parent oc.TerminalDevice_Channel_Otn_Esnr
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Otn_Esnr_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn_Esnr) *oc.QualifiedFloat64 {
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
