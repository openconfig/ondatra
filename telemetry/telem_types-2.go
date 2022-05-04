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

// QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup is a *NetworkInstance_PolicyForwarding_PathSelectionGroup with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_PathSelectionGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_PathSelectionGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup) Val(t testing.TB) *NetworkInstance_PolicyForwarding_PathSelectionGroup {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_PathSelectionGroup sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup) SetVal(v *NetworkInstance_PolicyForwarding_PathSelectionGroup) *QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_PathSelectionGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_PathSelectionGroup samples.
type CollectionNetworkInstance_PolicyForwarding_PathSelectionGroup struct {
	W    *NetworkInstance_PolicyForwarding_PathSelectionGroupWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_PathSelectionGroup) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_PathSelectionGroupWatcher observes a stream of *NetworkInstance_PolicyForwarding_PathSelectionGroup samples.
type NetworkInstance_PolicyForwarding_PathSelectionGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_PathSelectionGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_PathSelectionGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy is a *NetworkInstance_PolicyForwarding_Policy with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy) SetVal(v *NetworkInstance_PolicyForwarding_Policy) *QualifiedNetworkInstance_PolicyForwarding_Policy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy samples.
type CollectionNetworkInstance_PolicyForwarding_Policy struct {
	W    *NetworkInstance_PolicyForwarding_PolicyWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_PolicyWatcher observes a stream of *NetworkInstance_PolicyForwarding_Policy samples.
type NetworkInstance_PolicyForwarding_PolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_PolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule is a *NetworkInstance_PolicyForwarding_Policy_Rule with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule struct {
	W    *NetworkInstance_PolicyForwarding_Policy_RuleWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_RuleWatcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule samples.
type NetworkInstance_PolicyForwarding_Policy_RuleWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_RuleWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action is a *NetworkInstance_PolicyForwarding_Policy_Rule_Action with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule_Action // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Action sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule_Action {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Action sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule_Action) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule_Action samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_ActionWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_ActionWatcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule_Action samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_ActionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_ActionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre is a *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGreWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGreWatcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target is a *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_TargetWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_TargetWatcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_TargetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_TargetWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Action_EncapsulateGre_Target, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 is a *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4Watcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4Watcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4 samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 is a *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6Watcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6Watcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6 samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2 is a *NetworkInstance_PolicyForwarding_Policy_Rule_L2 with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2 struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule_L2 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_L2 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule_L2 {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_L2 sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule_L2) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_L2 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule_L2 samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_L2 struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_L2Watcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_L2) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_L2Watcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule_L2 samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_L2Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_L2Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport is a *NetworkInstance_PolicyForwarding_Policy_Rule_Transport with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport struct {
	*genutil.Metadata
	val     *NetworkInstance_PolicyForwarding_Policy_Rule_Transport // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Transport sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport) Val(t testing.TB) *NetworkInstance_PolicyForwarding_Policy_Rule_Transport {
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

// SetVal sets the value of the *NetworkInstance_PolicyForwarding_Policy_Rule_Transport sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport) SetVal(v *NetworkInstance_PolicyForwarding_Policy_Rule_Transport) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport is a telemetry Collection whose Await method returns a slice of *NetworkInstance_PolicyForwarding_Policy_Rule_Transport samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_TransportWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_TransportWatcher observes a stream of *NetworkInstance_PolicyForwarding_Policy_Rule_Transport samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_TransportWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_TransportWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol is a *NetworkInstance_Protocol with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol) Val(t testing.TB) *NetworkInstance_Protocol {
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

// SetVal sets the value of the *NetworkInstance_Protocol sample.
func (q *QualifiedNetworkInstance_Protocol) SetVal(v *NetworkInstance_Protocol) *QualifiedNetworkInstance_Protocol {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol samples.
type CollectionNetworkInstance_Protocol struct {
	W    *NetworkInstance_ProtocolWatcher
	Data []*QualifiedNetworkInstance_Protocol
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ProtocolWatcher observes a stream of *NetworkInstance_Protocol samples.
type NetworkInstance_ProtocolWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ProtocolWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Aggregate is a *NetworkInstance_Protocol_Aggregate with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Aggregate struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Aggregate // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Aggregate) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Aggregate sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Aggregate) Val(t testing.TB) *NetworkInstance_Protocol_Aggregate {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Aggregate sample.
func (q *QualifiedNetworkInstance_Protocol_Aggregate) SetVal(v *NetworkInstance_Protocol_Aggregate) *QualifiedNetworkInstance_Protocol_Aggregate {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Aggregate) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Aggregate is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Aggregate samples.
type CollectionNetworkInstance_Protocol_Aggregate struct {
	W    *NetworkInstance_Protocol_AggregateWatcher
	Data []*QualifiedNetworkInstance_Protocol_Aggregate
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Aggregate) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Aggregate {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_AggregateWatcher observes a stream of *NetworkInstance_Protocol_Aggregate samples.
type NetworkInstance_Protocol_AggregateWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Aggregate
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_AggregateWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Aggregate, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp is a *NetworkInstance_Protocol_Bgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp) SetVal(v *NetworkInstance_Protocol_Bgp) *QualifiedNetworkInstance_Protocol_Bgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp samples.
type CollectionNetworkInstance_Protocol_Bgp struct {
	W    *NetworkInstance_Protocol_BgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_BgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp samples.
type NetworkInstance_Protocol_BgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_BgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global is a *NetworkInstance_Protocol_Bgp_Global with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global) SetVal(v *NetworkInstance_Protocol_Bgp_Global) *QualifiedNetworkInstance_Protocol_Bgp_Global {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global samples.
type CollectionNetworkInstance_Protocol_Bgp_Global struct {
	W    *NetworkInstance_Protocol_Bgp_GlobalWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_GlobalWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global samples.
type NetworkInstance_Protocol_Bgp_GlobalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_GlobalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafiWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafiWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafiWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafiWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_AddPaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_Ipv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpnWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpnWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpnWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpnWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnEvpn_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVplsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L2VpnVpls_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_RouteSelectionOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4Watcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4 samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6Watcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6 samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp is a *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp) SetVal(v *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp) *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp struct {
	W    *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_IbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_IbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp samples.
type NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_IbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_IbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_AfiSafi_UseMultiplePaths_Ibgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation is a *NetworkInstance_Protocol_Bgp_Global_Confederation with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_Confederation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_Confederation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_Confederation {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_Confederation sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation) SetVal(v *NetworkInstance_Protocol_Bgp_Global_Confederation) *QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_Confederation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_Confederation samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_Confederation struct {
	W    *NetworkInstance_Protocol_Bgp_Global_ConfederationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_Confederation) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_ConfederationWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_Confederation samples.
type NetworkInstance_Protocol_Bgp_Global_ConfederationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_ConfederationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_Confederation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance is a *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance) SetVal(v *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance) *QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance struct {
	W    *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistanceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistanceWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance samples.
type NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_DefaultRouteDistanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_DefaultRouteDistance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix is a *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix) SetVal(v *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix) *QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix struct {
	W    *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefixWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix samples.
type NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_DynamicNeighborPrefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart is a *NetworkInstance_Protocol_Bgp_Global_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_GracefulRestart {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Bgp_Global_GracefulRestart) *QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Bgp_Global_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_GracefulRestart samples.
type NetworkInstance_Protocol_Bgp_Global_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions is a *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions) SetVal(v *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions) *QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions struct {
	W    *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions samples.
type NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_RouteSelectionOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_RouteSelectionOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_Global_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_Global_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp is a *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp) SetVal(v *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp) *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp samples.
type CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp struct {
	W    *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_IbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_IbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp samples.
type NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_IbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_IbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Global_UseMultiplePaths_Ibgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor is a *NetworkInstance_Protocol_Bgp_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor samples.
type NetworkInstance_Protocol_Bgp_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafiWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafiWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafiWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafiWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_AddPaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_ApplyPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Ipv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpnWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpnWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpnWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpnWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnEvpn_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVplsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L2VpnVpls_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_PrefixesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_PrefixesWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_PrefixesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_PrefixesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_Prefixes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4Watcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4 samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6Watcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6 samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AfiSafi_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy is a *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy samples.
type NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ApplyPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions is a *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions samples.
type NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_AsPathOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_AsPathOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop is a *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihopWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihopWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop samples.
type NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EbgpMultihop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd is a *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_EnableBfdWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfd samples.
type NetworkInstance_Protocol_Bgp_Neighbor_EnableBfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_EnableBfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_EnableBfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling is a *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandlingWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandlingWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling samples.
type NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandlingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_ErrorHandlingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_ErrorHandling, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart is a *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart samples.
type NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions is a *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions samples.
type NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_LoggingOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_LoggingOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages is a *NetworkInstance_Protocol_Bgp_Neighbor_Messages with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_Messages // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Messages sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_Messages {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Messages sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_Messages) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_Messages samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_MessagesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_MessagesWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_Messages samples.
type NetworkInstance_Protocol_Bgp_Neighbor_MessagesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_MessagesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received is a *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_Messages_ReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_Messages_ReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Received samples.
type NetworkInstance_Protocol_Bgp_Neighbor_Messages_ReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_Messages_ReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Received, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent is a *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_Messages_SentWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_Messages_SentWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent samples.
type NetworkInstance_Protocol_Bgp_Neighbor_Messages_SentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_Messages_SentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Messages_Sent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues is a *NetworkInstance_Protocol_Bgp_Neighbor_Queues with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_Queues // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Queues sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_Queues {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Queues sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_Queues) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_Queues is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_Queues samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_Queues struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_QueuesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_Queues) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_QueuesWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_Queues samples.
type NetworkInstance_Protocol_Bgp_Neighbor_QueuesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_QueuesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Queues, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector is a *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflectorWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_RouteReflectorWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector samples.
type NetworkInstance_Protocol_Bgp_Neighbor_RouteReflectorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflectorWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers is a *NetworkInstance_Protocol_Bgp_Neighbor_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_Timers {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_Timers) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_Timers samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_Timers struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_TimersWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_Timers samples.
type NetworkInstance_Protocol_Bgp_Neighbor_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport is a *NetworkInstance_Protocol_Bgp_Neighbor_Transport with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_Transport // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Transport sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_Transport {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_Transport sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_Transport) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_Transport is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_Transport samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_Transport struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_TransportWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_Transport) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_TransportWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_Transport samples.
type NetworkInstance_Protocol_Bgp_Neighbor_TransportWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_TransportWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_Transport, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup is a *NetworkInstance_Protocol_Bgp_PeerGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroupWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup samples.
type NetworkInstance_Protocol_Bgp_PeerGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafiWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafiWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafiWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafiWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_AddPaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_ApplyPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
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

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
