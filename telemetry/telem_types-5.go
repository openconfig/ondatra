package telemetry

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ygot/ygot"
)

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSizeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSizeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSizeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSizeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_LspBufferSize, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4ReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4ReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4ReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4ReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_PrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_PrefixWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_PrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_FlagsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_FlagsWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_FlagsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_FlagsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Flags, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_PrefixSid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_TagWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_TagWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_TagWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_TagWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64Watcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64Watcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64 samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_Subtlv_Tag64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv4Reachability_Prefix_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6ReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6ReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6ReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6ReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_PrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_PrefixWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_PrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_FlagsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_FlagsWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_FlagsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_FlagsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Flags, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_PrefixSid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_TagWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_TagWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_TagWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_TagWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64Watcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64Watcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64 samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_Subtlv_Tag64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIpv6Reachability_Prefix_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttributeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_InstanceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_InstanceWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_InstanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_InstanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_AvailableBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ExtendedAdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv4NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_Ipv6NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LanAdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkDelayVariation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLossWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkLoss, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_LinkProtectionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_MinMaxLinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_ResidualBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriorityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_SetupPriority, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_TeDefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UnconstrainedLsp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_Subtlv_UtilizedBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsisNeighborAttribute_Neighbor_Instance_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsnWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsnWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsnWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsnWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_InstanceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_InstanceWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_InstanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_InstanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_AvailableBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ExtendedAdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv4NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_Ipv6NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LanAdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributesWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkDelayVariation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLossWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLossWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLossWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLossWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkLoss, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_LinkProtectionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_MinMaxLinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_ResidualBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriorityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriorityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriorityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriorityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_SetupPriority, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_TeDefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UnconstrainedLsp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_Subtlv_UtilizedBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MtIsn_Neighbor_Instance_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopologyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopologyWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopologyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopologyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_TopologyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_TopologyWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_TopologyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_TopologyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_MultiTopology_Topology, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_NlpidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_NlpidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_NlpidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_NlpidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Nlpid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOiWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOiWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOiWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOiWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_PurgeOi, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_UndefinedTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference is a *NetworkInstance_Protocol_Isis_Level_RoutePreference with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_RoutePreference // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_RoutePreference sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_RoutePreference {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_RoutePreference sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference) SetVal(v *NetworkInstance_Protocol_Isis_Level_RoutePreference) *QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_RoutePreference is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_RoutePreference samples.
type CollectionNetworkInstance_Protocol_Isis_Level_RoutePreference struct {
	W    *NetworkInstance_Protocol_Isis_Level_RoutePreferenceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_RoutePreference) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_RoutePreferenceWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_RoutePreference samples.
type NetworkInstance_Protocol_Isis_Level_RoutePreferenceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_RoutePreferenceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_RoutePreference, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters is a *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters) SetVal(v *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters) *QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_SystemLevelCounters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters samples.
type CollectionNetworkInstance_Protocol_Isis_Level_SystemLevelCounters struct {
	W    *NetworkInstance_Protocol_Isis_Level_SystemLevelCountersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_SystemLevelCounters) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_SystemLevelCountersWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_SystemLevelCounters samples.
type NetworkInstance_Protocol_Isis_Level_SystemLevelCountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_SystemLevelCountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_SystemLevelCounters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
