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

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_PortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_PortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_PortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_PortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_PortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/port with a ONCE subscription.
func (n *System_GrpcServer_PortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_PortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_PortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_PortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_PortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_GrpcServer_PortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_PortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/port to the batch object.
func (n *System_GrpcServer_PortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_PortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_PortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
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
			qv := convertSystem_GrpcServer_PortPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_PortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_GrpcServer_PortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/port to the batch object.
func (n *System_GrpcServer_PortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_PortPath extracts the value of the leaf Port from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_GrpcServer_PortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedUint16 {
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

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/services with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_ServicesPath) Lookup(t testing.TB) *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_ServicesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/services with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_ServicesPath) Get(t testing.TB) []oc.E_SystemGrpc_GRPC_SERVICE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/services with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_ServicesPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_ServicesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/services with a ONCE subscription.
func (n *System_GrpcServer_ServicesPathAny) Get(t testing.TB) [][]oc.E_SystemGrpc_GRPC_SERVICE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_SystemGrpc_GRPC_SERVICE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/services with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_ServicesPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	c := &oc.CollectionE_SystemGrpc_GRPC_SERVICESlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_ServicesPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice) bool) *oc.E_SystemGrpc_GRPC_SERVICESliceWatcher {
	t.Helper()
	w := &oc.E_SystemGrpc_GRPC_SERVICESliceWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_ServicesPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/services with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_ServicesPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice) bool) *oc.E_SystemGrpc_GRPC_SERVICESliceWatcher {
	t.Helper()
	return watch_System_GrpcServer_ServicesPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/services with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_ServicesPath) Await(t testing.TB, timeout time.Duration, val []oc.E_SystemGrpc_GRPC_SERVICE) *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/services failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/services to the batch object.
func (n *System_GrpcServer_ServicesPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/services with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_ServicesPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	c := &oc.CollectionE_SystemGrpc_GRPC_SERVICESlice{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_ServicesPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice) bool) *oc.E_SystemGrpc_GRPC_SERVICESliceWatcher {
	t.Helper()
	w := &oc.E_SystemGrpc_GRPC_SERVICESliceWatcher{}
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
			qv := convertSystem_GrpcServer_ServicesPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/services with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_ServicesPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice) bool) *oc.E_SystemGrpc_GRPC_SERVICESliceWatcher {
	t.Helper()
	return watch_System_GrpcServer_ServicesPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/services to the batch object.
func (n *System_GrpcServer_ServicesPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_ServicesPath extracts the value of the leaf Services from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice.
func convertSystem_GrpcServer_ServicesPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	qv := &oc.QualifiedE_SystemGrpc_GRPC_SERVICESlice{
		Metadata: md,
	}
	val := parent.Services
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_GrpcServer_TransportSecurityPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_GrpcServer{}
	md, ok := oc.Lookup(t, n, "System_GrpcServer", goStruct, true, false)
	if ok {
		return convertSystem_GrpcServer_TransportSecurityPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetTransportSecurity())
}

// Get fetches the value at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_GrpcServer_TransportSecurityPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_GrpcServer_TransportSecurityPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_GrpcServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_GrpcServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_GrpcServer_TransportSecurityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a ONCE subscription.
func (n *System_GrpcServer_TransportSecurityPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_TransportSecurityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_TransportSecurityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_GrpcServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_GrpcServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_GrpcServer_TransportSecurityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_TransportSecurityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_GrpcServer_TransportSecurityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_GrpcServer_TransportSecurityPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/transport-security to the batch object.
func (n *System_GrpcServer_TransportSecurityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_GrpcServer_TransportSecurityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_GrpcServer_TransportSecurityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
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
			qv := convertSystem_GrpcServer_TransportSecurityPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/grpc-servers/grpc-server/state/transport-security with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_GrpcServer_TransportSecurityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_GrpcServer_TransportSecurityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/grpc-servers/grpc-server/state/transport-security to the batch object.
func (n *System_GrpcServer_TransportSecurityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_GrpcServer_TransportSecurityPath extracts the value of the leaf TransportSecurity from its parent oc.System_GrpcServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_GrpcServer_TransportSecurityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_GrpcServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.TransportSecurity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/state/hostname with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_HostnamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, false)
	if ok {
		return convertSystem_HostnamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/state/hostname with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_HostnamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/state/hostname with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_HostnamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertSystem_HostnamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/state/hostname with a ONCE subscription.
func (n *System_HostnamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_HostnamePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_HostnamePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_HostnamePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_HostnamePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_HostnamePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/state/hostname with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_HostnamePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/state/hostname failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/state/hostname to the batch object.
func (n *System_HostnamePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/state/hostname with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_HostnamePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_HostnamePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
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
			qv := convertSystem_HostnamePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/state/hostname with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_HostnamePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_HostnamePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/state/hostname to the batch object.
func (n *System_HostnamePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_HostnamePath extracts the value of the leaf Hostname from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_HostnamePath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/license with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_LicensePath) Lookup(t testing.TB) *oc.QualifiedSystem_License {
	t.Helper()
	goStruct := &oc.System_License{}
	md, ok := oc.Lookup(t, n, "System_License", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_LicensePath) Get(t testing.TB) *oc.System_License {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_LicensePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license with a ONCE subscription.
func (n *System_LicensePathAny) Get(t testing.TB) []*oc.System_License {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_License
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LicensePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License {
	t.Helper()
	c := &oc.CollectionSystem_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_License{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_License)))
		return false
	})
	return c
}

func watch_System_LicensePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License) bool) *oc.System_LicenseWatcher {
	t.Helper()
	w := &oc.System_LicenseWatcher{}
	gs := &oc.System_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_License{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LicensePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License) bool) *oc.System_LicenseWatcher {
	t.Helper()
	return watch_System_LicensePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_LicensePath) Await(t testing.TB, timeout time.Duration, val *oc.System_License) *oc.QualifiedSystem_License {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_License) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license to the batch object.
func (n *System_LicensePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LicensePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License {
	t.Helper()
	c := &oc.CollectionSystem_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_LicensePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License) bool) *oc.System_LicenseWatcher {
	t.Helper()
	w := &oc.System_LicenseWatcher{}
	structs := map[string]*oc.System_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_License{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LicensePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License) bool) *oc.System_LicenseWatcher {
	t.Helper()
	return watch_System_LicensePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license to the batch object.
func (n *System_LicensePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_LicensePath) Lookup(t testing.TB) *oc.QualifiedSystem_License_License {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_LicensePath) Get(t testing.TB) *oc.System_License_License {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_LicensePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License_License {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License_License
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license with a ONCE subscription.
func (n *System_License_LicensePathAny) Get(t testing.TB) []*oc.System_License_License {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_License_License
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_LicensePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License {
	t.Helper()
	c := &oc.CollectionSystem_License_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_License_License{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_License_License)))
		return false
	})
	return c
}

func watch_System_License_LicensePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License_License) bool) *oc.System_License_LicenseWatcher {
	t.Helper()
	w := &oc.System_License_LicenseWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_License_License{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License_License)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_LicensePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License) bool) *oc.System_License_LicenseWatcher {
	t.Helper()
	return watch_System_License_LicensePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_LicensePath) Await(t testing.TB, timeout time.Duration, val *oc.System_License_License) *oc.QualifiedSystem_License_License {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_License_License) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license to the batch object.
func (n *System_License_LicensePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_LicensePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License {
	t.Helper()
	c := &oc.CollectionSystem_License_License{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_LicensePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License_License) bool) *oc.System_License_LicenseWatcher {
	t.Helper()
	w := &oc.System_License_LicenseWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_License_License{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License_License)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_LicensePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License) bool) *oc.System_License_LicenseWatcher {
	t.Helper()
	return watch_System_License_LicensePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license to the batch object.
func (n *System_License_LicensePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ActivePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ActivePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetActive())
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ActivePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ActivePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ActivePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/active with a ONCE subscription.
func (n *System_License_License_ActivePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ActivePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ActivePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_ActivePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ActivePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ActivePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ActivePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/active failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/active to the batch object.
func (n *System_License_License_ActivePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ActivePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ActivePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_ActivePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/active with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ActivePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ActivePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/active to the batch object.
func (n *System_License_License_ActivePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ActivePath extracts the value of the leaf Active from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_ActivePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Active
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/description with a ONCE subscription.
func (n *System_License_License_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_DescriptionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_DescriptionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_DescriptionPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_DescriptionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_DescriptionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_DescriptionPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/description failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/description to the batch object.
func (n *System_License_License_DescriptionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_DescriptionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_DescriptionPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_DescriptionPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/description with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_DescriptionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_DescriptionPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/description to the batch object.
func (n *System_License_License_DescriptionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_DescriptionPath extracts the value of the leaf Description from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_License_License_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ExpirationDatePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ExpirationDatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ExpirationDatePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ExpirationDatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ExpirationDatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a ONCE subscription.
func (n *System_License_License_ExpirationDatePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpirationDatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ExpirationDatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_ExpirationDatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpirationDatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_ExpirationDatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ExpirationDatePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/expiration-date failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expiration-date to the batch object.
func (n *System_License_License_ExpirationDatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpirationDatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ExpirationDatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_ExpirationDatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expiration-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpirationDatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_ExpirationDatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expiration-date to the batch object.
func (n *System_License_License_ExpirationDatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ExpirationDatePath extracts the value of the leaf ExpirationDate from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_License_License_ExpirationDatePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.ExpirationDate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ExpiredPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ExpiredPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ExpiredPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ExpiredPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ExpiredPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/expired with a ONCE subscription.
func (n *System_License_License_ExpiredPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpiredPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ExpiredPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_ExpiredPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpiredPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ExpiredPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ExpiredPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/expired failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expired to the batch object.
func (n *System_License_License_ExpiredPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ExpiredPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ExpiredPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_ExpiredPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/expired with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ExpiredPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ExpiredPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/expired to the batch object.
func (n *System_License_License_ExpiredPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ExpiredPath extracts the value of the leaf Expired from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_ExpiredPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Expired
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_InUsePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_InUsePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_InUsePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_InUsePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_InUsePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/in-use with a ONCE subscription.
func (n *System_License_License_InUsePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_InUsePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_InUsePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_InUsePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_InUsePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_InUsePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_InUsePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/in-use failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/in-use to the batch object.
func (n *System_License_License_InUsePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_InUsePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_InUsePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_InUsePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/in-use with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_InUsePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_InUsePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/in-use to the batch object.
func (n *System_License_License_InUsePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_InUsePath extracts the value of the leaf InUse from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_InUsePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.InUse
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_IssueDatePath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_IssueDatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_IssueDatePath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_IssueDatePathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_IssueDatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/issue-date with a ONCE subscription.
func (n *System_License_License_IssueDatePathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_IssueDatePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_IssueDatePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_IssueDatePath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_IssueDatePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_IssueDatePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_IssueDatePath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/issue-date failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/issue-date to the batch object.
func (n *System_License_License_IssueDatePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_IssueDatePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_IssueDatePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_IssueDatePath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/issue-date with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_IssueDatePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_System_License_License_IssueDatePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/issue-date to the batch object.
func (n *System_License_License_IssueDatePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_IssueDatePath extracts the value of the leaf IssueDate from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertSystem_License_License_IssueDatePath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.IssueDate
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_LicenseDataPath) Lookup(t testing.TB) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_LicenseDataPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_LicenseDataPath) Get(t testing.TB) oc.System_License_License_LicenseData_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_LicenseDataPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_License_License_LicenseData_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_LicenseDataPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/license-data with a ONCE subscription.
func (n *System_License_License_LicenseDataPathAny) Get(t testing.TB) []oc.System_License_License_LicenseData_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.System_License_License_LicenseData_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseDataPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License_LicenseData_Union {
	t.Helper()
	c := &oc.CollectionSystem_License_License_LicenseData_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License_LicenseData_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_LicenseDataPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License_License_LicenseData_Union) bool) *oc.System_License_License_LicenseData_UnionWatcher {
	t.Helper()
	w := &oc.System_License_License_LicenseData_UnionWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_LicenseDataPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License_License_LicenseData_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseDataPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License_LicenseData_Union) bool) *oc.System_License_License_LicenseData_UnionWatcher {
	t.Helper()
	return watch_System_License_License_LicenseDataPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_LicenseDataPath) Await(t testing.TB, timeout time.Duration, val oc.System_License_License_LicenseData_Union) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_License_License_LicenseData_Union) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/license-data failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-data to the batch object.
func (n *System_License_License_LicenseDataPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseDataPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_License_License_LicenseData_Union {
	t.Helper()
	c := &oc.CollectionSystem_License_License_LicenseData_Union{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_License_License_LicenseData_Union) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_LicenseDataPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_License_License_LicenseData_Union) bool) *oc.System_License_License_LicenseData_UnionWatcher {
	t.Helper()
	w := &oc.System_License_License_LicenseData_UnionWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_LicenseDataPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_License_License_LicenseData_Union)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-data with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseDataPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_License_License_LicenseData_Union) bool) *oc.System_License_License_LicenseData_UnionWatcher {
	t.Helper()
	return watch_System_License_License_LicenseDataPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-data to the batch object.
func (n *System_License_License_LicenseDataPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_LicenseDataPath extracts the value of the leaf LicenseData from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_License_License_LicenseData_Union.
func convertSystem_License_License_LicenseDataPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	qv := &oc.QualifiedSystem_License_License_LicenseData_Union{
		Metadata: md,
	}
	val := parent.LicenseData
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_LicenseIdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_LicenseIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_LicenseIdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_LicenseIdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_LicenseIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/license-id with a ONCE subscription.
func (n *System_License_License_LicenseIdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseIdPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_LicenseIdPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_LicenseIdPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseIdPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_LicenseIdPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_LicenseIdPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/license-id failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-id to the batch object.
func (n *System_License_License_LicenseIdPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_LicenseIdPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_LicenseIdPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_LicenseIdPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/license-id with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_LicenseIdPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_License_License_LicenseIdPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/license-id to the batch object.
func (n *System_License_License_LicenseIdPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_LicenseIdPath extracts the value of the leaf LicenseId from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_License_License_LicenseIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LicenseId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_License_License_ValidPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_License_License{}
	md, ok := oc.Lookup(t, n, "System_License_License", goStruct, true, false)
	if ok {
		return convertSystem_License_License_ValidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_License_License_ValidPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_License_License_ValidPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_License_License{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_License_License", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_License_License_ValidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/license/licenses/license/state/valid with a ONCE subscription.
func (n *System_License_License_ValidPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ValidPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ValidPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_License_License", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_License_License_ValidPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ValidPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ValidPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_License_License_ValidPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/license/licenses/license/state/valid failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/license/licenses/license/state/valid to the batch object.
func (n *System_License_License_ValidPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_License_License_ValidPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_License_License_ValidPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	structs := map[string]*oc.System_License_License{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_License_License{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_License_License", structs[pre], queryPath, true, false)
			qv := convertSystem_License_License_ValidPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/license/licenses/license/state/valid with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_License_License_ValidPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_License_License_ValidPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/license/licenses/license/state/valid to the batch object.
func (n *System_License_License_ValidPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_License_License_ValidPath extracts the value of the leaf Valid from its parent oc.System_License_License
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_License_License_ValidPath(t testing.TB, md *genutil.Metadata, parent *oc.System_License_License) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Valid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_LoggingPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging {
	t.Helper()
	goStruct := &oc.System_Logging{}
	md, ok := oc.Lookup(t, n, "System_Logging", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_LoggingPath) Get(t testing.TB) *oc.System_Logging {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_LoggingPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging with a ONCE subscription.
func (n *System_LoggingPathAny) Get(t testing.TB) []*oc.System_Logging {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LoggingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging {
	t.Helper()
	c := &oc.CollectionSystem_Logging{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Logging{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Logging)))
		return false
	})
	return c
}

func watch_System_LoggingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging) bool) *oc.System_LoggingWatcher {
	t.Helper()
	w := &oc.System_LoggingWatcher{}
	gs := &oc.System_Logging{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Logging{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LoggingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging) bool) *oc.System_LoggingWatcher {
	t.Helper()
	return watch_System_LoggingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_LoggingPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Logging) *oc.QualifiedSystem_Logging {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Logging) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging to the batch object.
func (n *System_LoggingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_LoggingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging {
	t.Helper()
	c := &oc.CollectionSystem_Logging{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_LoggingPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging) bool) *oc.System_LoggingWatcher {
	t.Helper()
	w := &oc.System_LoggingWatcher{}
	structs := map[string]*oc.System_Logging{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Logging{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_LoggingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging) bool) *oc.System_LoggingWatcher {
	t.Helper()
	return watch_System_LoggingPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging to the batch object.
func (n *System_LoggingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/logging/console with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_ConsolePath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_Console {
	t.Helper()
	goStruct := &oc.System_Logging_Console{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Logging_Console{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_ConsolePath) Get(t testing.TB) *oc.System_Logging_Console {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_ConsolePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_Console {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_Console
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_Console{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console with a ONCE subscription.
func (n *System_Logging_ConsolePathAny) Get(t testing.TB) []*oc.System_Logging_Console {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_Console
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_ConsolePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_Console {
	t.Helper()
	c := &oc.CollectionSystem_Logging_Console{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_Console) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Logging_Console{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Logging_Console)))
		return false
	})
	return c
}

func watch_System_Logging_ConsolePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console) bool) *oc.System_Logging_ConsoleWatcher {
	t.Helper()
	w := &oc.System_Logging_ConsoleWatcher{}
	gs := &oc.System_Logging_Console{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_Console", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Logging_Console{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_Console)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_ConsolePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console) bool) *oc.System_Logging_ConsoleWatcher {
	t.Helper()
	return watch_System_Logging_ConsolePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/console with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_ConsolePath) Await(t testing.TB, timeout time.Duration, val *oc.System_Logging_Console) *oc.QualifiedSystem_Logging_Console {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Logging_Console) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/console failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/console to the batch object.
func (n *System_Logging_ConsolePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_ConsolePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_Console {
	t.Helper()
	c := &oc.CollectionSystem_Logging_Console{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_Console) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_ConsolePathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console) bool) *oc.System_Logging_ConsoleWatcher {
	t.Helper()
	w := &oc.System_Logging_ConsoleWatcher{}
	structs := map[string]*oc.System_Logging_Console{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_Console{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_Console", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Logging_Console{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_Console)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_ConsolePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console) bool) *oc.System_Logging_ConsoleWatcher {
	t.Helper()
	return watch_System_Logging_ConsolePathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/console to the batch object.
func (n *System_Logging_ConsolePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_Console_SelectorPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_Console_Selector {
	t.Helper()
	goStruct := &oc.System_Logging_Console_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console_Selector", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Logging_Console_Selector{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_Console_SelectorPath) Get(t testing.TB) *oc.System_Logging_Console_Selector {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_Console_SelectorPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_Console_Selector {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_Console_Selector
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console_Selector", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_Console_Selector{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription.
func (n *System_Logging_Console_SelectorPathAny) Get(t testing.TB) []*oc.System_Logging_Console_Selector {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_Console_Selector
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console/selectors/selector with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_Console_SelectorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_Console_Selector {
	t.Helper()
	c := &oc.CollectionSystem_Logging_Console_Selector{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_Console_Selector) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Logging_Console_Selector{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Logging_Console_Selector)))
		return false
	})
	return c
}

func watch_System_Logging_Console_SelectorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console_Selector) bool) *oc.System_Logging_Console_SelectorWatcher {
	t.Helper()
	w := &oc.System_Logging_Console_SelectorWatcher{}
	gs := &oc.System_Logging_Console_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_Console_Selector", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Logging_Console_Selector{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_Console_Selector)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console/selectors/selector with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_Console_SelectorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console_Selector) bool) *oc.System_Logging_Console_SelectorWatcher {
	t.Helper()
	return watch_System_Logging_Console_SelectorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/console/selectors/selector with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_Console_SelectorPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Logging_Console_Selector) *oc.QualifiedSystem_Logging_Console_Selector {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Logging_Console_Selector) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/console/selectors/selector failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/console/selectors/selector to the batch object.
func (n *System_Logging_Console_SelectorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console/selectors/selector with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_Console_SelectorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_Console_Selector {
	t.Helper()
	c := &oc.CollectionSystem_Logging_Console_Selector{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_Console_Selector) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_Console_SelectorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console_Selector) bool) *oc.System_Logging_Console_SelectorWatcher {
	t.Helper()
	w := &oc.System_Logging_Console_SelectorWatcher{}
	structs := map[string]*oc.System_Logging_Console_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_Console_Selector{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_Console_Selector", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Logging_Console_Selector{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_Console_Selector)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console/selectors/selector with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_Console_SelectorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_Console_Selector) bool) *oc.System_Logging_Console_SelectorWatcher {
	t.Helper()
	return watch_System_Logging_Console_SelectorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/console/selectors/selector to the batch object.
func (n *System_Logging_Console_SelectorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/logging/console/selectors/selector/state/facility with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_Console_Selector_FacilityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	goStruct := &oc.System_Logging_Console_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console_Selector", goStruct, true, false)
	if ok {
		return convertSystem_Logging_Console_Selector_FacilityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console/selectors/selector/state/facility with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_Console_Selector_FacilityPath) Get(t testing.TB) oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console/selectors/selector/state/facility with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_Console_Selector_FacilityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console_Selector", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_Console_Selector_FacilityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console/selectors/selector/state/facility with a ONCE subscription.
func (n *System_Logging_Console_Selector_FacilityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SYSLOG_FACILITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console/selectors/selector/state/facility with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_Console_Selector_FacilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SYSLOG_FACILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_Console_Selector_FacilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	w := &oc.E_SystemLogging_SYSLOG_FACILITYWatcher{}
	gs := &oc.System_Logging_Console_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_Console_Selector", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Logging_Console_Selector_FacilityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console/selectors/selector/state/facility with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_Console_Selector_FacilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	return watch_System_Logging_Console_Selector_FacilityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/console/selectors/selector/state/facility with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_Console_Selector_FacilityPath) Await(t testing.TB, timeout time.Duration, val oc.E_SystemLogging_SYSLOG_FACILITY) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/console/selectors/selector/state/facility failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/console/selectors/selector/state/facility to the batch object.
func (n *System_Logging_Console_Selector_FacilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console/selectors/selector/state/facility with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_Console_Selector_FacilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SYSLOG_FACILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_Console_Selector_FacilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	w := &oc.E_SystemLogging_SYSLOG_FACILITYWatcher{}
	structs := map[string]*oc.System_Logging_Console_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_Console_Selector{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_Console_Selector", structs[pre], queryPath, true, false)
			qv := convertSystem_Logging_Console_Selector_FacilityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console/selectors/selector/state/facility with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_Console_Selector_FacilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	return watch_System_Logging_Console_Selector_FacilityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/console/selectors/selector/state/facility to the batch object.
func (n *System_Logging_Console_Selector_FacilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Logging_Console_Selector_FacilityPath extracts the value of the leaf Facility from its parent oc.System_Logging_Console_Selector
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY.
func convertSystem_Logging_Console_Selector_FacilityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_Console_Selector) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SYSLOG_FACILITY{
		Metadata: md,
	}
	val := parent.Facility
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/console/selectors/selector/state/severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_Console_Selector_SeverityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	goStruct := &oc.System_Logging_Console_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console_Selector", goStruct, true, false)
	if ok {
		return convertSystem_Logging_Console_Selector_SeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console/selectors/selector/state/severity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_Console_Selector_SeverityPath) Get(t testing.TB) oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console/selectors/selector/state/severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_Console_Selector_SeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SyslogSeverity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console_Selector", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_Console_Selector_SeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console/selectors/selector/state/severity with a ONCE subscription.
func (n *System_Logging_Console_Selector_SeverityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SyslogSeverity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console/selectors/selector/state/severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_Console_Selector_SeverityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SyslogSeverity {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SyslogSeverity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SyslogSeverity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_Console_Selector_SeverityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SyslogSeverity) bool) *oc.E_SystemLogging_SyslogSeverityWatcher {
	t.Helper()
	w := &oc.E_SystemLogging_SyslogSeverityWatcher{}
	gs := &oc.System_Logging_Console_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_Console_Selector", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Logging_Console_Selector_SeverityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemLogging_SyslogSeverity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console/selectors/selector/state/severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_Console_Selector_SeverityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SyslogSeverity) bool) *oc.E_SystemLogging_SyslogSeverityWatcher {
	t.Helper()
	return watch_System_Logging_Console_Selector_SeverityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/console/selectors/selector/state/severity with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_Console_Selector_SeverityPath) Await(t testing.TB, timeout time.Duration, val oc.E_SystemLogging_SyslogSeverity) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SystemLogging_SyslogSeverity) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/console/selectors/selector/state/severity failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/console/selectors/selector/state/severity to the batch object.
func (n *System_Logging_Console_Selector_SeverityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/console/selectors/selector/state/severity with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_Console_Selector_SeverityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SyslogSeverity {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SyslogSeverity{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SyslogSeverity) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_Console_Selector_SeverityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SyslogSeverity) bool) *oc.E_SystemLogging_SyslogSeverityWatcher {
	t.Helper()
	w := &oc.E_SystemLogging_SyslogSeverityWatcher{}
	structs := map[string]*oc.System_Logging_Console_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_Console_Selector{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_Console_Selector", structs[pre], queryPath, true, false)
			qv := convertSystem_Logging_Console_Selector_SeverityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemLogging_SyslogSeverity)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/console/selectors/selector/state/severity with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_Console_Selector_SeverityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SyslogSeverity) bool) *oc.E_SystemLogging_SyslogSeverityWatcher {
	t.Helper()
	return watch_System_Logging_Console_Selector_SeverityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/console/selectors/selector/state/severity to the batch object.
func (n *System_Logging_Console_Selector_SeverityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Logging_Console_Selector_SeverityPath extracts the value of the leaf Severity from its parent oc.System_Logging_Console_Selector
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SyslogSeverity.
func convertSystem_Logging_Console_Selector_SeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_Console_Selector) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SyslogSeverity{
		Metadata: md,
	}
	val := parent.Severity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_RemoteServer {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Logging_RemoteServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServerPath) Get(t testing.TB) *oc.System_Logging_RemoteServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_RemoteServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_RemoteServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_RemoteServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription.
func (n *System_Logging_RemoteServerPathAny) Get(t testing.TB) []*oc.System_Logging_RemoteServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_RemoteServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_RemoteServer {
	t.Helper()
	c := &oc.CollectionSystem_Logging_RemoteServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_RemoteServer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Logging_RemoteServer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Logging_RemoteServer)))
		return false
	})
	return c
}

func watch_System_Logging_RemoteServerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer) bool) *oc.System_Logging_RemoteServerWatcher {
	t.Helper()
	w := &oc.System_Logging_RemoteServerWatcher{}
	gs := &oc.System_Logging_RemoteServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_RemoteServer", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Logging_RemoteServer{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_RemoteServer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer) bool) *oc.System_Logging_RemoteServerWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/remote-servers/remote-server with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_RemoteServerPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Logging_RemoteServer) *oc.QualifiedSystem_Logging_RemoteServer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Logging_RemoteServer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/remote-servers/remote-server failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server to the batch object.
func (n *System_Logging_RemoteServerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_RemoteServer {
	t.Helper()
	c := &oc.CollectionSystem_Logging_RemoteServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_RemoteServer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServerPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer) bool) *oc.System_Logging_RemoteServerWatcher {
	t.Helper()
	w := &oc.System_Logging_RemoteServerWatcher{}
	structs := map[string]*oc.System_Logging_RemoteServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_RemoteServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_RemoteServer", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Logging_RemoteServer{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_RemoteServer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer) bool) *oc.System_Logging_RemoteServerWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServerPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server to the batch object.
func (n *System_Logging_RemoteServerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_HostPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer", goStruct, true, false)
	if ok {
		return convertSystem_Logging_RemoteServer_HostPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_HostPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_HostPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_HostPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a ONCE subscription.
func (n *System_Logging_RemoteServer_HostPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_HostPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_HostPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.System_Logging_RemoteServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_RemoteServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Logging_RemoteServer_HostPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_HostPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_HostPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_RemoteServer_HostPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/remote-servers/remote-server/state/host failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/state/host to the batch object.
func (n *System_Logging_RemoteServer_HostPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_HostPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_HostPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	structs := map[string]*oc.System_Logging_RemoteServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_RemoteServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_RemoteServer", structs[pre], queryPath, true, false)
			qv := convertSystem_Logging_RemoteServer_HostPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/host with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_HostPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_HostPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/state/host to the batch object.
func (n *System_Logging_RemoteServer_HostPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Logging_RemoteServer_HostPath extracts the value of the leaf Host from its parent oc.System_Logging_RemoteServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Logging_RemoteServer_HostPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Host
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_RemotePortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer", goStruct, true, false)
	if ok {
		return convertSystem_Logging_RemoteServer_RemotePortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetRemotePort())
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_RemotePortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_RemotePortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_RemotePortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a ONCE subscription.
func (n *System_Logging_RemoteServer_RemotePortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_RemotePortPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_RemotePortPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_Logging_RemoteServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_RemoteServer", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Logging_RemoteServer_RemotePortPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_RemotePortPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_RemotePortPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_RemoteServer_RemotePortPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port to the batch object.
func (n *System_Logging_RemoteServer_RemotePortPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_RemotePortPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_RemotePortPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	structs := map[string]*oc.System_Logging_RemoteServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_RemoteServer{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_RemoteServer", structs[pre], queryPath, true, false)
			qv := convertSystem_Logging_RemoteServer_RemotePortPath(t, md, structs[pre])
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

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_RemotePortPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_RemotePortPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/state/remote-port to the batch object.
func (n *System_Logging_RemoteServer_RemotePortPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Logging_RemoteServer_RemotePortPath extracts the value of the leaf RemotePort from its parent oc.System_Logging_RemoteServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Logging_RemoteServer_RemotePortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.RemotePort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_SelectorPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_RemoteServer_Selector {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer_Selector", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_Logging_RemoteServer_Selector{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_SelectorPath) Get(t testing.TB) *oc.System_Logging_RemoteServer_Selector {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_SelectorPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_RemoteServer_Selector {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_RemoteServer_Selector
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer_Selector", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_RemoteServer_Selector{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription.
func (n *System_Logging_RemoteServer_SelectorPathAny) Get(t testing.TB) []*oc.System_Logging_RemoteServer_Selector {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_RemoteServer_Selector
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_SelectorPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_RemoteServer_Selector {
	t.Helper()
	c := &oc.CollectionSystem_Logging_RemoteServer_Selector{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_RemoteServer_Selector) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_Logging_RemoteServer_Selector{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_Logging_RemoteServer_Selector)))
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_SelectorPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer_Selector) bool) *oc.System_Logging_RemoteServer_SelectorWatcher {
	t.Helper()
	w := &oc.System_Logging_RemoteServer_SelectorWatcher{}
	gs := &oc.System_Logging_RemoteServer_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_RemoteServer_Selector", gs, queryPath, false, false)
		qv := (&oc.QualifiedSystem_Logging_RemoteServer_Selector{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_RemoteServer_Selector)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_SelectorPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer_Selector) bool) *oc.System_Logging_RemoteServer_SelectorWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_SelectorPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_RemoteServer_SelectorPath) Await(t testing.TB, timeout time.Duration, val *oc.System_Logging_RemoteServer_Selector) *oc.QualifiedSystem_Logging_RemoteServer_Selector {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_Logging_RemoteServer_Selector) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector to the batch object.
func (n *System_Logging_RemoteServer_SelectorPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_SelectorPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_Logging_RemoteServer_Selector {
	t.Helper()
	c := &oc.CollectionSystem_Logging_RemoteServer_Selector{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_Logging_RemoteServer_Selector) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_SelectorPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer_Selector) bool) *oc.System_Logging_RemoteServer_SelectorWatcher {
	t.Helper()
	w := &oc.System_Logging_RemoteServer_SelectorWatcher{}
	structs := map[string]*oc.System_Logging_RemoteServer_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_RemoteServer_Selector{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_RemoteServer_Selector", structs[pre], queryPath, false, false)
			qv := (&oc.QualifiedSystem_Logging_RemoteServer_Selector{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_Logging_RemoteServer_Selector)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_SelectorPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_Logging_RemoteServer_Selector) bool) *oc.System_Logging_RemoteServer_SelectorWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_SelectorPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector to the batch object.
func (n *System_Logging_RemoteServer_SelectorPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer_Selector", goStruct, true, false)
	if ok {
		return convertSystem_Logging_RemoteServer_Selector_FacilityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Get(t testing.TB) oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_Selector_FacilityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer_Selector", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_Selector_FacilityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a ONCE subscription.
func (n *System_Logging_RemoteServer_Selector_FacilityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SYSLOG_FACILITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SYSLOG_FACILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_Selector_FacilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	w := &oc.E_SystemLogging_SYSLOG_FACILITYWatcher{}
	gs := &oc.System_Logging_RemoteServer_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_Logging_RemoteServer_Selector", gs, queryPath, true, false)
		return []genutil.QualifiedValue{convertSystem_Logging_RemoteServer_Selector_FacilityPath(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_Selector_FacilityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Await(t testing.TB, timeout time.Duration, val oc.E_SystemLogging_SYSLOG_FACILITY) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility to the batch object.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_Logging_RemoteServer_Selector_FacilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	c := &oc.CollectionE_SystemLogging_SYSLOG_FACILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_Logging_RemoteServer_Selector_FacilityPathAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	w := &oc.E_SystemLogging_SYSLOG_FACILITYWatcher{}
	structs := map[string]*oc.System_Logging_RemoteServer_Selector{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &oc.System_Logging_RemoteServer_Selector{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], oc.GetSchema(), "System_Logging_RemoteServer_Selector", structs[pre], queryPath, true, false)
			qv := convertSystem_Logging_RemoteServer_Selector_FacilityPath(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_Logging_RemoteServer_Selector_FacilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY) bool) *oc.E_SystemLogging_SYSLOG_FACILITYWatcher {
	t.Helper()
	return watch_System_Logging_RemoteServer_Selector_FacilityPathAny(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/state/facility to the batch object.
func (n *System_Logging_RemoteServer_Selector_FacilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_Logging_RemoteServer_Selector_FacilityPath extracts the value of the leaf Facility from its parent oc.System_Logging_RemoteServer_Selector
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY.
func convertSystem_Logging_RemoteServer_Selector_FacilityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer_Selector) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SYSLOG_FACILITY{
		Metadata: md,
	}
	val := parent.Facility
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
