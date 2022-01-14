package qos

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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Ipv6_Protocol_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) bool) *oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) bool) *oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Await(t testing.TB, timeout time.Duration, val oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Ipv6_Protocol_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) bool) *oc.Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/protocol to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_ProtocolPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath extracts the value of the leaf Protocol from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union.
func convertQos_Classifier_Term_Conditions_Ipv6_ProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union{
		Metadata: md,
	}
	val := parent.Protocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-address to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/state/source-flow-label to the batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath extracts the value of the leaf SourceFlowLabel from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SourceFlowLabel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2Path) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_L2{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2Path) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_L2 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2PathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_L2
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_L2{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2PathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_L2 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_L2
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_L2{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_L2) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions_L2{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions_L2)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_L2Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_L2) bool) *oc.Qos_Classifier_Term_Conditions_L2Watcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_L2Watcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Conditions_L2{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_L2)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_L2) bool) *oc.Qos_Classifier_Term_Conditions_L2Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_L2Path) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_L2) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_L2{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_L2) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_L2) bool) *oc.Qos_Classifier_Term_Conditions_L2Watcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2Path(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac-mask to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath extracts the value of the leaf DestinationMacMask from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_DestinationMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_DestinationMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_L2_DestinationMacPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_L2_DestinationMacPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_DestinationMacPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_DestinationMacPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/destination-mac to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_L2_DestinationMacPath extracts the value of the leaf DestinationMac from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_DestinationMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_EthertypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_EthertypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_L2_Ethertype_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_L2_EthertypePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) bool) *oc.Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_L2_EthertypePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) bool) *oc.Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_EthertypePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Await(t testing.TB, timeout time.Duration, val oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union) *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_L2_Ethertype_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) bool) *oc.Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_EthertypePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/ethertype to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_L2_EthertypePath extracts the value of the leaf Ethertype from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union.
func convertQos_Classifier_Term_Conditions_L2_EthertypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union{
		Metadata: md,
	}
	val := parent.Ethertype
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac-mask to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath extracts the value of the leaf SourceMacMask from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_SourceMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_SourceMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_L2_SourceMacPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_L2{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_L2_SourceMacPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_SourceMacPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_L2_SourceMacPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/state/source-mac to the batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_L2_SourceMacPath extracts the value of the leaf SourceMac from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_SourceMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, false, false)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Mpls
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Qos_Classifier_Term_Conditions_Mpls)))
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_MplsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool) *oc.Qos_Classifier_Term_Conditions_MplsWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_MplsWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", gs, queryPath, false, false)
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool) *oc.Qos_Classifier_Term_Conditions_MplsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_MplsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Await(t testing.TB, timeout time.Duration, val *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls to the batch object.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls) bool) *oc.Qos_Classifier_Term_Conditions_MplsWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_MplsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls to the batch object.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, false)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	w := &oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher{}
	gs := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", gs, queryPath, true, false)
		return convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Await(t testing.TB, timeout time.Duration, val oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	c := &oc.CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) bool) *oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher {
	t.Helper()
	return watch_Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/state/end-label-value to the batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath extracts the value of the leaf EndLabelValue from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union.
func convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union{
		Metadata: md,
	}
	val := parent.EndLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
