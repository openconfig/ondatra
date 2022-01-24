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

// Lookup fetches the value at /openconfig-system/system/ssh-server/state/protocol-version with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_ProtocolVersionPath) Lookup(t testing.TB) *oc.QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, false)
	if ok {
		return convertSystem_SshServer_ProtocolVersionPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_SshServer_ProtocolVersion{
		Metadata: md,
	}).SetVal(goStruct.GetProtocolVersion())
}

// Get fetches the value at /openconfig-system/system/ssh-server/state/protocol-version with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_ProtocolVersionPath) Get(t testing.TB) oc.E_SshServer_ProtocolVersion {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/state/protocol-version with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_ProtocolVersionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SshServer_ProtocolVersion
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_ProtocolVersionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/state/protocol-version with a ONCE subscription.
func (n *System_SshServer_ProtocolVersionPathAny) Get(t testing.TB) []oc.E_SshServer_ProtocolVersion {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SshServer_ProtocolVersion
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/protocol-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_ProtocolVersionPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SshServer_ProtocolVersion {
	t.Helper()
	c := &oc.CollectionE_SshServer_ProtocolVersion{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SshServer_ProtocolVersion) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_SshServer_ProtocolVersionPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_SshServer_ProtocolVersion) bool) *oc.E_SshServer_ProtocolVersionWatcher {
	t.Helper()
	w := &oc.E_SshServer_ProtocolVersionWatcher{}
	gs := &oc.System_SshServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_SshServer", gs, queryPath, true, false)
		return convertSystem_SshServer_ProtocolVersionPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_SshServer_ProtocolVersion)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/protocol-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_ProtocolVersionPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SshServer_ProtocolVersion) bool) *oc.E_SshServer_ProtocolVersionWatcher {
	t.Helper()
	return watch_System_SshServer_ProtocolVersionPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ssh-server/state/protocol-version with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_SshServer_ProtocolVersionPath) Await(t testing.TB, timeout time.Duration, val oc.E_SshServer_ProtocolVersion) *oc.QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_SshServer_ProtocolVersion) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ssh-server/state/protocol-version failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ssh-server/state/protocol-version to the batch object.
func (n *System_SshServer_ProtocolVersionPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/protocol-version with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_ProtocolVersionPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_SshServer_ProtocolVersion {
	t.Helper()
	c := &oc.CollectionE_SshServer_ProtocolVersion{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_SshServer_ProtocolVersion) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/protocol-version with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_ProtocolVersionPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_SshServer_ProtocolVersion) bool) *oc.E_SshServer_ProtocolVersionWatcher {
	t.Helper()
	return watch_System_SshServer_ProtocolVersionPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ssh-server/state/protocol-version to the batch object.
func (n *System_SshServer_ProtocolVersionPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_SshServer_ProtocolVersionPath extracts the value of the leaf ProtocolVersion from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SshServer_ProtocolVersion.
func convertSystem_SshServer_ProtocolVersionPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	qv := &oc.QualifiedE_SshServer_ProtocolVersion{
		Metadata: md,
	}
	val := parent.ProtocolVersion
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/state/rate-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_RateLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, false)
	if ok {
		return convertSystem_SshServer_RateLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ssh-server/state/rate-limit with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_RateLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/state/rate-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_RateLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_RateLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/state/rate-limit with a ONCE subscription.
func (n *System_SshServer_RateLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/rate-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_RateLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_SshServer_RateLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_SshServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_SshServer", gs, queryPath, true, false)
		return convertSystem_SshServer_RateLimitPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/rate-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_RateLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_SshServer_RateLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ssh-server/state/rate-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_SshServer_RateLimitPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ssh-server/state/rate-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ssh-server/state/rate-limit to the batch object.
func (n *System_SshServer_RateLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/rate-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_RateLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/rate-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_RateLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_SshServer_RateLimitPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ssh-server/state/rate-limit to the batch object.
func (n *System_SshServer_RateLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_SshServer_RateLimitPath extracts the value of the leaf RateLimit from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_SshServer_RateLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.RateLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/state/session-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_SessionLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, false)
	if ok {
		return convertSystem_SshServer_SessionLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ssh-server/state/session-limit with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_SessionLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/state/session-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_SessionLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_SessionLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/state/session-limit with a ONCE subscription.
func (n *System_SshServer_SessionLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/session-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_SessionLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_SshServer_SessionLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_SshServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_SshServer", gs, queryPath, true, false)
		return convertSystem_SshServer_SessionLimitPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/session-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_SessionLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_SshServer_SessionLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ssh-server/state/session-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_SshServer_SessionLimitPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ssh-server/state/session-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ssh-server/state/session-limit to the batch object.
func (n *System_SshServer_SessionLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/session-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_SessionLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/session-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_SessionLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_SshServer_SessionLimitPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ssh-server/state/session-limit to the batch object.
func (n *System_SshServer_SessionLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_SshServer_SessionLimitPath extracts the value of the leaf SessionLimit from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_SshServer_SessionLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SessionLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ssh-server/state/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_SshServer_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_SshServer{}
	md, ok := oc.Lookup(t, n, "System_SshServer", goStruct, true, false)
	if ok {
		return convertSystem_SshServer_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ssh-server/state/timeout with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_SshServer_TimeoutPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ssh-server/state/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_SshServer_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_SshServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_SshServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_SshServer_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ssh-server/state/timeout with a ONCE subscription.
func (n *System_SshServer_TimeoutPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_TimeoutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_SshServer_TimeoutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_SshServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_SshServer", gs, queryPath, true, false)
		return convertSystem_SshServer_TimeoutPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_TimeoutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_SshServer_TimeoutPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/ssh-server/state/timeout with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_SshServer_TimeoutPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/ssh-server/state/timeout failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/ssh-server/state/timeout to the batch object.
func (n *System_SshServer_TimeoutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/ssh-server/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_SshServer_TimeoutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/ssh-server/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_SshServer_TimeoutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_SshServer_TimeoutPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/ssh-server/state/timeout to the batch object.
func (n *System_SshServer_TimeoutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_SshServer_TimeoutPath extracts the value of the leaf Timeout from its parent oc.System_SshServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_SshServer_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.System_SshServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Timeout
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_TelnetServer {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSystem_TelnetServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServerPath) Get(t testing.TB) *oc.System_TelnetServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_TelnetServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_TelnetServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_TelnetServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server with a ONCE subscription.
func (n *System_TelnetServerPathAny) Get(t testing.TB) []*oc.System_TelnetServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_TelnetServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServerPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_TelnetServer {
	t.Helper()
	c := &oc.CollectionSystem_TelnetServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_TelnetServer) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedSystem_TelnetServer{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.System_TelnetServer)))
		return false
	})
	return c
}

func watch_System_TelnetServerPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedSystem_TelnetServer) bool) *oc.System_TelnetServerWatcher {
	t.Helper()
	w := &oc.System_TelnetServerWatcher{}
	gs := &oc.System_TelnetServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_TelnetServer", gs, queryPath, false, false)
		return (&oc.QualifiedSystem_TelnetServer{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedSystem_TelnetServer)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServerPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_TelnetServer) bool) *oc.System_TelnetServerWatcher {
	t.Helper()
	return watch_System_TelnetServerPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/telnet-server with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_TelnetServerPath) Await(t testing.TB, timeout time.Duration, val *oc.System_TelnetServer) *oc.QualifiedSystem_TelnetServer {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedSystem_TelnetServer) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/telnet-server failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/telnet-server to the batch object.
func (n *System_TelnetServerPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServerPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionSystem_TelnetServer {
	t.Helper()
	c := &oc.CollectionSystem_TelnetServer{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedSystem_TelnetServer) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServerPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedSystem_TelnetServer) bool) *oc.System_TelnetServerWatcher {
	t.Helper()
	return watch_System_TelnetServerPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/telnet-server to the batch object.
func (n *System_TelnetServerPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/state/enable with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_EnablePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, false)
	if ok {
		return convertSystem_TelnetServer_EnablePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnable())
}

// Get fetches the value at /openconfig-system/system/telnet-server/state/enable with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_EnablePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/state/enable with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_EnablePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_EnablePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/state/enable with a ONCE subscription.
func (n *System_TelnetServer_EnablePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/enable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_EnablePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_TelnetServer_EnablePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.System_TelnetServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_TelnetServer", gs, queryPath, true, false)
		return convertSystem_TelnetServer_EnablePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/enable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_EnablePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_TelnetServer_EnablePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/telnet-server/state/enable with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_TelnetServer_EnablePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/telnet-server/state/enable failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/telnet-server/state/enable to the batch object.
func (n *System_TelnetServer_EnablePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/enable with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_EnablePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/enable with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_EnablePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_System_TelnetServer_EnablePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/telnet-server/state/enable to the batch object.
func (n *System_TelnetServer_EnablePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_TelnetServer_EnablePath extracts the value of the leaf Enable from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_TelnetServer_EnablePath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enable
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/state/rate-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_RateLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, false)
	if ok {
		return convertSystem_TelnetServer_RateLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server/state/rate-limit with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_RateLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/state/rate-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_RateLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_RateLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/state/rate-limit with a ONCE subscription.
func (n *System_TelnetServer_RateLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/rate-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_RateLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_TelnetServer_RateLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_TelnetServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_TelnetServer", gs, queryPath, true, false)
		return convertSystem_TelnetServer_RateLimitPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/rate-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_RateLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_TelnetServer_RateLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/telnet-server/state/rate-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_TelnetServer_RateLimitPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/telnet-server/state/rate-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/telnet-server/state/rate-limit to the batch object.
func (n *System_TelnetServer_RateLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/rate-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_RateLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/rate-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_RateLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_TelnetServer_RateLimitPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/telnet-server/state/rate-limit to the batch object.
func (n *System_TelnetServer_RateLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_TelnetServer_RateLimitPath extracts the value of the leaf RateLimit from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_TelnetServer_RateLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.RateLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/state/session-limit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_SessionLimitPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, false)
	if ok {
		return convertSystem_TelnetServer_SessionLimitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server/state/session-limit with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_SessionLimitPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/state/session-limit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_SessionLimitPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_SessionLimitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/state/session-limit with a ONCE subscription.
func (n *System_TelnetServer_SessionLimitPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/session-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_SessionLimitPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_TelnetServer_SessionLimitPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_TelnetServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_TelnetServer", gs, queryPath, true, false)
		return convertSystem_TelnetServer_SessionLimitPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/session-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_SessionLimitPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_TelnetServer_SessionLimitPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/telnet-server/state/session-limit with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_TelnetServer_SessionLimitPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/telnet-server/state/session-limit failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/telnet-server/state/session-limit to the batch object.
func (n *System_TelnetServer_SessionLimitPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/session-limit with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_SessionLimitPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/session-limit with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_SessionLimitPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_TelnetServer_SessionLimitPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/telnet-server/state/session-limit to the batch object.
func (n *System_TelnetServer_SessionLimitPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_TelnetServer_SessionLimitPath extracts the value of the leaf SessionLimit from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_TelnetServer_SessionLimitPath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SessionLimit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/telnet-server/state/timeout with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_TelnetServer_TimeoutPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_TelnetServer{}
	md, ok := oc.Lookup(t, n, "System_TelnetServer", goStruct, true, false)
	if ok {
		return convertSystem_TelnetServer_TimeoutPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/telnet-server/state/timeout with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_TelnetServer_TimeoutPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/telnet-server/state/timeout with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_TelnetServer_TimeoutPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_TelnetServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_TelnetServer", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertSystem_TelnetServer_TimeoutPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/telnet-server/state/timeout with a ONCE subscription.
func (n *System_TelnetServer_TimeoutPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_TimeoutPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_System_TelnetServer_TimeoutPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.System_TelnetServer{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "System_TelnetServer", gs, queryPath, true, false)
		return convertSystem_TelnetServer_TimeoutPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_TimeoutPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_TelnetServer_TimeoutPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-system/system/telnet-server/state/timeout with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *System_TelnetServer_TimeoutPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-system/system/telnet-server/state/timeout failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-system/system/telnet-server/state/timeout to the batch object.
func (n *System_TelnetServer_TimeoutPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-system/system/telnet-server/state/timeout with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *System_TelnetServer_TimeoutPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-system/system/telnet-server/state/timeout with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *System_TelnetServer_TimeoutPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_System_TelnetServer_TimeoutPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-system/system/telnet-server/state/timeout to the batch object.
func (n *System_TelnetServer_TimeoutPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertSystem_TelnetServer_TimeoutPath extracts the value of the leaf Timeout from its parent oc.System_TelnetServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_TelnetServer_TimeoutPath(t testing.TB, md *genutil.Metadata, parent *oc.System_TelnetServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Timeout
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
