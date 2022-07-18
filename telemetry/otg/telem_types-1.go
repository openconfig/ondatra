package otg

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ygot/ygot"
)

// QualifiedIsisRouter is a *IsisRouter with a corresponding timestamp.
type QualifiedIsisRouter struct {
	*genutil.Metadata
	val     *IsisRouter // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter sample, erroring out if not present.
func (q *QualifiedIsisRouter) Val(t testing.TB) *IsisRouter {
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

// SetVal sets the value of the *IsisRouter sample.
func (q *QualifiedIsisRouter) SetVal(v *IsisRouter) *QualifiedIsisRouter {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter is a telemetry Collection whose Await method returns a slice of *IsisRouter samples.
type CollectionIsisRouter struct {
	W    *IsisRouterWatcher
	Data []*QualifiedIsisRouter
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter) Await(t testing.TB) []*QualifiedIsisRouter {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouterWatcher observes a stream of *IsisRouter samples.
type IsisRouterWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouterWatcher) Await(t testing.TB) (*QualifiedIsisRouter, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_Counters is a *IsisRouter_Counters with a corresponding timestamp.
type QualifiedIsisRouter_Counters struct {
	*genutil.Metadata
	val     *IsisRouter_Counters // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_Counters sample, erroring out if not present.
func (q *QualifiedIsisRouter_Counters) Val(t testing.TB) *IsisRouter_Counters {
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

// SetVal sets the value of the *IsisRouter_Counters sample.
func (q *QualifiedIsisRouter_Counters) SetVal(v *IsisRouter_Counters) *QualifiedIsisRouter_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_Counters is a telemetry Collection whose Await method returns a slice of *IsisRouter_Counters samples.
type CollectionIsisRouter_Counters struct {
	W    *IsisRouter_CountersWatcher
	Data []*QualifiedIsisRouter_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_Counters) Await(t testing.TB) []*QualifiedIsisRouter_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_CountersWatcher observes a stream of *IsisRouter_Counters samples.
type IsisRouter_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_CountersWatcher) Await(t testing.TB) (*QualifiedIsisRouter_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_Counters_Level1 is a *IsisRouter_Counters_Level1 with a corresponding timestamp.
type QualifiedIsisRouter_Counters_Level1 struct {
	*genutil.Metadata
	val     *IsisRouter_Counters_Level1 // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_Counters_Level1) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_Counters_Level1 sample, erroring out if not present.
func (q *QualifiedIsisRouter_Counters_Level1) Val(t testing.TB) *IsisRouter_Counters_Level1 {
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

// SetVal sets the value of the *IsisRouter_Counters_Level1 sample.
func (q *QualifiedIsisRouter_Counters_Level1) SetVal(v *IsisRouter_Counters_Level1) *QualifiedIsisRouter_Counters_Level1 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_Counters_Level1) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_Counters_Level1 is a telemetry Collection whose Await method returns a slice of *IsisRouter_Counters_Level1 samples.
type CollectionIsisRouter_Counters_Level1 struct {
	W    *IsisRouter_Counters_Level1Watcher
	Data []*QualifiedIsisRouter_Counters_Level1
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_Counters_Level1) Await(t testing.TB) []*QualifiedIsisRouter_Counters_Level1 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_Counters_Level1Watcher observes a stream of *IsisRouter_Counters_Level1 samples.
type IsisRouter_Counters_Level1Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_Counters_Level1
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_Counters_Level1Watcher) Await(t testing.TB) (*QualifiedIsisRouter_Counters_Level1, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_Counters_Level2 is a *IsisRouter_Counters_Level2 with a corresponding timestamp.
type QualifiedIsisRouter_Counters_Level2 struct {
	*genutil.Metadata
	val     *IsisRouter_Counters_Level2 // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_Counters_Level2) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_Counters_Level2 sample, erroring out if not present.
func (q *QualifiedIsisRouter_Counters_Level2) Val(t testing.TB) *IsisRouter_Counters_Level2 {
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

// SetVal sets the value of the *IsisRouter_Counters_Level2 sample.
func (q *QualifiedIsisRouter_Counters_Level2) SetVal(v *IsisRouter_Counters_Level2) *QualifiedIsisRouter_Counters_Level2 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_Counters_Level2) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_Counters_Level2 is a telemetry Collection whose Await method returns a slice of *IsisRouter_Counters_Level2 samples.
type CollectionIsisRouter_Counters_Level2 struct {
	W    *IsisRouter_Counters_Level2Watcher
	Data []*QualifiedIsisRouter_Counters_Level2
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_Counters_Level2) Await(t testing.TB) []*QualifiedIsisRouter_Counters_Level2 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_Counters_Level2Watcher observes a stream of *IsisRouter_Counters_Level2 samples.
type IsisRouter_Counters_Level2Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_Counters_Level2
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_Counters_Level2Watcher) Await(t testing.TB) (*QualifiedIsisRouter_Counters_Level2, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLacp is a *Lacp with a corresponding timestamp.
type QualifiedLacp struct {
	*genutil.Metadata
	val     *Lacp // val is the sample value.
	present bool
}

func (q *QualifiedLacp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lacp sample, erroring out if not present.
func (q *QualifiedLacp) Val(t testing.TB) *Lacp {
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

// SetVal sets the value of the *Lacp sample.
func (q *QualifiedLacp) SetVal(v *Lacp) *QualifiedLacp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLacp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLacp is a telemetry Collection whose Await method returns a slice of *Lacp samples.
type CollectionLacp struct {
	W    *LacpWatcher
	Data []*QualifiedLacp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLacp) Await(t testing.TB) []*QualifiedLacp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LacpWatcher observes a stream of *Lacp samples.
type LacpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLacp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LacpWatcher) Await(t testing.TB) (*QualifiedLacp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLacp_LagMember is a *Lacp_LagMember with a corresponding timestamp.
type QualifiedLacp_LagMember struct {
	*genutil.Metadata
	val     *Lacp_LagMember // val is the sample value.
	present bool
}

func (q *QualifiedLacp_LagMember) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lacp_LagMember sample, erroring out if not present.
func (q *QualifiedLacp_LagMember) Val(t testing.TB) *Lacp_LagMember {
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

// SetVal sets the value of the *Lacp_LagMember sample.
func (q *QualifiedLacp_LagMember) SetVal(v *Lacp_LagMember) *QualifiedLacp_LagMember {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLacp_LagMember) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLacp_LagMember is a telemetry Collection whose Await method returns a slice of *Lacp_LagMember samples.
type CollectionLacp_LagMember struct {
	W    *Lacp_LagMemberWatcher
	Data []*QualifiedLacp_LagMember
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLacp_LagMember) Await(t testing.TB) []*QualifiedLacp_LagMember {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lacp_LagMemberWatcher observes a stream of *Lacp_LagMember samples.
type Lacp_LagMemberWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLacp_LagMember
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lacp_LagMemberWatcher) Await(t testing.TB) (*QualifiedLacp_LagMember, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLacp_LagMember_Counters is a *Lacp_LagMember_Counters with a corresponding timestamp.
type QualifiedLacp_LagMember_Counters struct {
	*genutil.Metadata
	val     *Lacp_LagMember_Counters // val is the sample value.
	present bool
}

func (q *QualifiedLacp_LagMember_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lacp_LagMember_Counters sample, erroring out if not present.
func (q *QualifiedLacp_LagMember_Counters) Val(t testing.TB) *Lacp_LagMember_Counters {
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

// SetVal sets the value of the *Lacp_LagMember_Counters sample.
func (q *QualifiedLacp_LagMember_Counters) SetVal(v *Lacp_LagMember_Counters) *QualifiedLacp_LagMember_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLacp_LagMember_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLacp_LagMember_Counters is a telemetry Collection whose Await method returns a slice of *Lacp_LagMember_Counters samples.
type CollectionLacp_LagMember_Counters struct {
	W    *Lacp_LagMember_CountersWatcher
	Data []*QualifiedLacp_LagMember_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLacp_LagMember_Counters) Await(t testing.TB) []*QualifiedLacp_LagMember_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lacp_LagMember_CountersWatcher observes a stream of *Lacp_LagMember_Counters samples.
type Lacp_LagMember_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLacp_LagMember_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lacp_LagMember_CountersWatcher) Await(t testing.TB) (*QualifiedLacp_LagMember_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLag is a *Lag with a corresponding timestamp.
type QualifiedLag struct {
	*genutil.Metadata
	val     *Lag // val is the sample value.
	present bool
}

func (q *QualifiedLag) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lag sample, erroring out if not present.
func (q *QualifiedLag) Val(t testing.TB) *Lag {
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

// SetVal sets the value of the *Lag sample.
func (q *QualifiedLag) SetVal(v *Lag) *QualifiedLag {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLag) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLag is a telemetry Collection whose Await method returns a slice of *Lag samples.
type CollectionLag struct {
	W    *LagWatcher
	Data []*QualifiedLag
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLag) Await(t testing.TB) []*QualifiedLag {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LagWatcher observes a stream of *Lag samples.
type LagWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLag
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LagWatcher) Await(t testing.TB) (*QualifiedLag, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLag_Counters is a *Lag_Counters with a corresponding timestamp.
type QualifiedLag_Counters struct {
	*genutil.Metadata
	val     *Lag_Counters // val is the sample value.
	present bool
}

func (q *QualifiedLag_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lag_Counters sample, erroring out if not present.
func (q *QualifiedLag_Counters) Val(t testing.TB) *Lag_Counters {
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

// SetVal sets the value of the *Lag_Counters sample.
func (q *QualifiedLag_Counters) SetVal(v *Lag_Counters) *QualifiedLag_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLag_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLag_Counters is a telemetry Collection whose Await method returns a slice of *Lag_Counters samples.
type CollectionLag_Counters struct {
	W    *Lag_CountersWatcher
	Data []*QualifiedLag_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLag_Counters) Await(t testing.TB) []*QualifiedLag_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lag_CountersWatcher observes a stream of *Lag_Counters samples.
type Lag_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLag_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lag_CountersWatcher) Await(t testing.TB) (*QualifiedLag_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedPort is a *Port with a corresponding timestamp.
type QualifiedPort struct {
	*genutil.Metadata
	val     *Port // val is the sample value.
	present bool
}

func (q *QualifiedPort) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Port sample, erroring out if not present.
func (q *QualifiedPort) Val(t testing.TB) *Port {
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

// SetVal sets the value of the *Port sample.
func (q *QualifiedPort) SetVal(v *Port) *QualifiedPort {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedPort) IsPresent() bool {
	return q != nil && q.present
}

// CollectionPort is a telemetry Collection whose Await method returns a slice of *Port samples.
type CollectionPort struct {
	W    *PortWatcher
	Data []*QualifiedPort
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionPort) Await(t testing.TB) []*QualifiedPort {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// PortWatcher observes a stream of *Port samples.
type PortWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedPort
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *PortWatcher) Await(t testing.TB) (*QualifiedPort, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedPort_Counters is a *Port_Counters with a corresponding timestamp.
type QualifiedPort_Counters struct {
	*genutil.Metadata
	val     *Port_Counters // val is the sample value.
	present bool
}

func (q *QualifiedPort_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Port_Counters sample, erroring out if not present.
func (q *QualifiedPort_Counters) Val(t testing.TB) *Port_Counters {
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

// SetVal sets the value of the *Port_Counters sample.
func (q *QualifiedPort_Counters) SetVal(v *Port_Counters) *QualifiedPort_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedPort_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionPort_Counters is a telemetry Collection whose Await method returns a slice of *Port_Counters samples.
type CollectionPort_Counters struct {
	W    *Port_CountersWatcher
	Data []*QualifiedPort_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionPort_Counters) Await(t testing.TB) []*QualifiedPort_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Port_CountersWatcher observes a stream of *Port_Counters samples.
type Port_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedPort_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Port_CountersWatcher) Await(t testing.TB) (*QualifiedPort_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpPeer_SessionState is a E_BgpPeer_SessionState with a corresponding timestamp.
type QualifiedE_BgpPeer_SessionState struct {
	*genutil.Metadata
	val     E_BgpPeer_SessionState // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpPeer_SessionState) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpPeer_SessionState sample, erroring out if not present.
func (q *QualifiedE_BgpPeer_SessionState) Val(t testing.TB) E_BgpPeer_SessionState {
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

// SetVal sets the value of the E_BgpPeer_SessionState sample.
func (q *QualifiedE_BgpPeer_SessionState) SetVal(v E_BgpPeer_SessionState) *QualifiedE_BgpPeer_SessionState {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpPeer_SessionState) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpPeer_SessionState is a telemetry Collection whose Await method returns a slice of E_BgpPeer_SessionState samples.
type CollectionE_BgpPeer_SessionState struct {
	W    *E_BgpPeer_SessionStateWatcher
	Data []*QualifiedE_BgpPeer_SessionState
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpPeer_SessionState) Await(t testing.TB) []*QualifiedE_BgpPeer_SessionState {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpPeer_SessionStateWatcher observes a stream of E_BgpPeer_SessionState samples.
type E_BgpPeer_SessionStateWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpPeer_SessionState
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpPeer_SessionStateWatcher) Await(t testing.TB) (*QualifiedE_BgpPeer_SessionState, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lag_OperStatus is a E_Lag_OperStatus with a corresponding timestamp.
type QualifiedE_Lag_OperStatus struct {
	*genutil.Metadata
	val     E_Lag_OperStatus // val is the sample value.
	present bool
}

func (q *QualifiedE_Lag_OperStatus) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lag_OperStatus sample, erroring out if not present.
func (q *QualifiedE_Lag_OperStatus) Val(t testing.TB) E_Lag_OperStatus {
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

// SetVal sets the value of the E_Lag_OperStatus sample.
func (q *QualifiedE_Lag_OperStatus) SetVal(v E_Lag_OperStatus) *QualifiedE_Lag_OperStatus {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lag_OperStatus) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lag_OperStatus is a telemetry Collection whose Await method returns a slice of E_Lag_OperStatus samples.
type CollectionE_Lag_OperStatus struct {
	W    *E_Lag_OperStatusWatcher
	Data []*QualifiedE_Lag_OperStatus
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lag_OperStatus) Await(t testing.TB) []*QualifiedE_Lag_OperStatus {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lag_OperStatusWatcher observes a stream of E_Lag_OperStatus samples.
type E_Lag_OperStatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lag_OperStatus
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lag_OperStatusWatcher) Await(t testing.TB) (*QualifiedE_Lag_OperStatus, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType is a E_OpenTrafficGeneratorLacp_LacpActivityType with a corresponding timestamp.
type QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType struct {
	*genutil.Metadata
	val     E_OpenTrafficGeneratorLacp_LacpActivityType // val is the sample value.
	present bool
}

func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OpenTrafficGeneratorLacp_LacpActivityType sample, erroring out if not present.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) Val(t testing.TB) E_OpenTrafficGeneratorLacp_LacpActivityType {
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

// SetVal sets the value of the E_OpenTrafficGeneratorLacp_LacpActivityType sample.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) SetVal(v E_OpenTrafficGeneratorLacp_LacpActivityType) *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OpenTrafficGeneratorLacp_LacpActivityType is a telemetry Collection whose Await method returns a slice of E_OpenTrafficGeneratorLacp_LacpActivityType samples.
type CollectionE_OpenTrafficGeneratorLacp_LacpActivityType struct {
	W    *E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher
	Data []*QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OpenTrafficGeneratorLacp_LacpActivityType) Await(t testing.TB) []*QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher observes a stream of E_OpenTrafficGeneratorLacp_LacpActivityType samples.
type E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher) Await(t testing.TB) (*QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
