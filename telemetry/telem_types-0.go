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

// QualifiedAcl is a *Acl with a corresponding timestamp.
type QualifiedAcl struct {
	*genutil.Metadata
	val     *Acl // val is the sample value.
	present bool
}

func (q *QualifiedAcl) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl sample, erroring out if not present.
func (q *QualifiedAcl) Val(t testing.TB) *Acl {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl sample.
func (q *QualifiedAcl) SetVal(v *Acl) *QualifiedAcl {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl is a telemetry Collection whose Await method returns a slice of *Acl samples.
type CollectionAcl struct {
	W    *AclWatcher
	Data []*QualifiedAcl
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl) Await(t testing.TB) []*QualifiedAcl {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// AclWatcher observes a stream of *Acl samples.
type AclWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *AclWatcher) Await(t testing.TB) (*QualifiedAcl, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet is a *Acl_AclSet with a corresponding timestamp.
type QualifiedAcl_AclSet struct {
	*genutil.Metadata
	val     *Acl_AclSet // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet sample, erroring out if not present.
func (q *QualifiedAcl_AclSet) Val(t testing.TB) *Acl_AclSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet sample.
func (q *QualifiedAcl_AclSet) SetVal(v *Acl_AclSet) *QualifiedAcl_AclSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet is a telemetry Collection whose Await method returns a slice of *Acl_AclSet samples.
type CollectionAcl_AclSet struct {
	W    *Acl_AclSetWatcher
	Data []*QualifiedAcl_AclSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet) Await(t testing.TB) []*QualifiedAcl_AclSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSetWatcher observes a stream of *Acl_AclSet samples.
type Acl_AclSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSetWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry is a *Acl_AclSet_AclEntry with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry) Val(t testing.TB) *Acl_AclSet_AclEntry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry sample.
func (q *QualifiedAcl_AclSet_AclEntry) SetVal(v *Acl_AclSet_AclEntry) *QualifiedAcl_AclSet_AclEntry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry samples.
type CollectionAcl_AclSet_AclEntry struct {
	W    *Acl_AclSet_AclEntryWatcher
	Data []*QualifiedAcl_AclSet_AclEntry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntryWatcher observes a stream of *Acl_AclSet_AclEntry samples.
type Acl_AclSet_AclEntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntryWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Actions is a *Acl_AclSet_AclEntry_Actions with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Actions struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_Actions // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Actions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_Actions sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Actions) Val(t testing.TB) *Acl_AclSet_AclEntry_Actions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_Actions sample.
func (q *QualifiedAcl_AclSet_AclEntry_Actions) SetVal(v *Acl_AclSet_AclEntry_Actions) *QualifiedAcl_AclSet_AclEntry_Actions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Actions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Actions is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_Actions samples.
type CollectionAcl_AclSet_AclEntry_Actions struct {
	W    *Acl_AclSet_AclEntry_ActionsWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Actions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Actions) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Actions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_ActionsWatcher observes a stream of *Acl_AclSet_AclEntry_Actions samples.
type Acl_AclSet_AclEntry_ActionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Actions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_ActionsWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Actions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_InputInterface is a *Acl_AclSet_AclEntry_InputInterface with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_InputInterface struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_InputInterface // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_InputInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_InputInterface sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_InputInterface) Val(t testing.TB) *Acl_AclSet_AclEntry_InputInterface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_InputInterface sample.
func (q *QualifiedAcl_AclSet_AclEntry_InputInterface) SetVal(v *Acl_AclSet_AclEntry_InputInterface) *QualifiedAcl_AclSet_AclEntry_InputInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_InputInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_InputInterface is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_InputInterface samples.
type CollectionAcl_AclSet_AclEntry_InputInterface struct {
	W    *Acl_AclSet_AclEntry_InputInterfaceWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_InputInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_InputInterface) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_InputInterfaceWatcher observes a stream of *Acl_AclSet_AclEntry_InputInterface samples.
type Acl_AclSet_AclEntry_InputInterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_InputInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_InputInterfaceWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_InputInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef is a *Acl_AclSet_AclEntry_InputInterface_InterfaceRef with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_InputInterface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_InputInterface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) Val(t testing.TB) *Acl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_InputInterface_InterfaceRef sample.
func (q *QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) SetVal(v *Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_InputInterface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_InputInterface_InterfaceRef samples.
type CollectionAcl_AclSet_AclEntry_InputInterface_InterfaceRef struct {
	W    *Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_InputInterface_InterfaceRef) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher observes a stream of *Acl_AclSet_AclEntry_InputInterface_InterfaceRef samples.
type Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_InputInterface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Ipv4 is a *Acl_AclSet_AclEntry_Ipv4 with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Ipv4 struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_Ipv4 // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Ipv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_Ipv4 sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv4) Val(t testing.TB) *Acl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_Ipv4 sample.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv4) SetVal(v *Acl_AclSet_AclEntry_Ipv4) *QualifiedAcl_AclSet_AclEntry_Ipv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Ipv4 is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_Ipv4 samples.
type CollectionAcl_AclSet_AclEntry_Ipv4 struct {
	W    *Acl_AclSet_AclEntry_Ipv4Watcher
	Data []*QualifiedAcl_AclSet_AclEntry_Ipv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Ipv4) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Ipv4Watcher observes a stream of *Acl_AclSet_AclEntry_Ipv4 samples.
type Acl_AclSet_AclEntry_Ipv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Ipv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Ipv4Watcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Ipv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Ipv6 is a *Acl_AclSet_AclEntry_Ipv6 with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Ipv6 struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_Ipv6 // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Ipv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_Ipv6 sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv6) Val(t testing.TB) *Acl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_Ipv6 sample.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv6) SetVal(v *Acl_AclSet_AclEntry_Ipv6) *QualifiedAcl_AclSet_AclEntry_Ipv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Ipv6 is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_Ipv6 samples.
type CollectionAcl_AclSet_AclEntry_Ipv6 struct {
	W    *Acl_AclSet_AclEntry_Ipv6Watcher
	Data []*QualifiedAcl_AclSet_AclEntry_Ipv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Ipv6) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Ipv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Ipv6Watcher observes a stream of *Acl_AclSet_AclEntry_Ipv6 samples.
type Acl_AclSet_AclEntry_Ipv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Ipv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Ipv6Watcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Ipv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_L2 is a *Acl_AclSet_AclEntry_L2 with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_L2 struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_L2 // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_L2) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_L2 sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_L2) Val(t testing.TB) *Acl_AclSet_AclEntry_L2 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_L2 sample.
func (q *QualifiedAcl_AclSet_AclEntry_L2) SetVal(v *Acl_AclSet_AclEntry_L2) *QualifiedAcl_AclSet_AclEntry_L2 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_L2) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_L2 is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_L2 samples.
type CollectionAcl_AclSet_AclEntry_L2 struct {
	W    *Acl_AclSet_AclEntry_L2Watcher
	Data []*QualifiedAcl_AclSet_AclEntry_L2
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_L2) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_L2 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_L2Watcher observes a stream of *Acl_AclSet_AclEntry_L2 samples.
type Acl_AclSet_AclEntry_L2Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_L2
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_L2Watcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_L2, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Mpls is a *Acl_AclSet_AclEntry_Mpls with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Mpls struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_Mpls // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Mpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_Mpls sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls) Val(t testing.TB) *Acl_AclSet_AclEntry_Mpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_Mpls sample.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls) SetVal(v *Acl_AclSet_AclEntry_Mpls) *QualifiedAcl_AclSet_AclEntry_Mpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Mpls is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_Mpls samples.
type CollectionAcl_AclSet_AclEntry_Mpls struct {
	W    *Acl_AclSet_AclEntry_MplsWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Mpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Mpls) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Mpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_MplsWatcher observes a stream of *Acl_AclSet_AclEntry_Mpls samples.
type Acl_AclSet_AclEntry_MplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Mpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_MplsWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Mpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Transport is a *Acl_AclSet_AclEntry_Transport with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Transport struct {
	*genutil.Metadata
	val     *Acl_AclSet_AclEntry_Transport // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Transport) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_AclSet_AclEntry_Transport sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Transport) Val(t testing.TB) *Acl_AclSet_AclEntry_Transport {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_AclSet_AclEntry_Transport sample.
func (q *QualifiedAcl_AclSet_AclEntry_Transport) SetVal(v *Acl_AclSet_AclEntry_Transport) *QualifiedAcl_AclSet_AclEntry_Transport {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Transport) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Transport is a telemetry Collection whose Await method returns a slice of *Acl_AclSet_AclEntry_Transport samples.
type CollectionAcl_AclSet_AclEntry_Transport struct {
	W    *Acl_AclSet_AclEntry_TransportWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Transport
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Transport) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Transport {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_TransportWatcher observes a stream of *Acl_AclSet_AclEntry_Transport samples.
type Acl_AclSet_AclEntry_TransportWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Transport
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_TransportWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Transport, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_Interface is a *Acl_Interface with a corresponding timestamp.
type QualifiedAcl_Interface struct {
	*genutil.Metadata
	val     *Acl_Interface // val is the sample value.
	present bool
}

func (q *QualifiedAcl_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_Interface sample, erroring out if not present.
func (q *QualifiedAcl_Interface) Val(t testing.TB) *Acl_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_Interface sample.
func (q *QualifiedAcl_Interface) SetVal(v *Acl_Interface) *QualifiedAcl_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_Interface is a telemetry Collection whose Await method returns a slice of *Acl_Interface samples.
type CollectionAcl_Interface struct {
	W    *Acl_InterfaceWatcher
	Data []*QualifiedAcl_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_Interface) Await(t testing.TB) []*QualifiedAcl_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_InterfaceWatcher observes a stream of *Acl_Interface samples.
type Acl_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_InterfaceWatcher) Await(t testing.TB) (*QualifiedAcl_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_Interface_EgressAclSet is a *Acl_Interface_EgressAclSet with a corresponding timestamp.
type QualifiedAcl_Interface_EgressAclSet struct {
	*genutil.Metadata
	val     *Acl_Interface_EgressAclSet // val is the sample value.
	present bool
}

func (q *QualifiedAcl_Interface_EgressAclSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_Interface_EgressAclSet sample, erroring out if not present.
func (q *QualifiedAcl_Interface_EgressAclSet) Val(t testing.TB) *Acl_Interface_EgressAclSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_Interface_EgressAclSet sample.
func (q *QualifiedAcl_Interface_EgressAclSet) SetVal(v *Acl_Interface_EgressAclSet) *QualifiedAcl_Interface_EgressAclSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_Interface_EgressAclSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_Interface_EgressAclSet is a telemetry Collection whose Await method returns a slice of *Acl_Interface_EgressAclSet samples.
type CollectionAcl_Interface_EgressAclSet struct {
	W    *Acl_Interface_EgressAclSetWatcher
	Data []*QualifiedAcl_Interface_EgressAclSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_Interface_EgressAclSet) Await(t testing.TB) []*QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_Interface_EgressAclSetWatcher observes a stream of *Acl_Interface_EgressAclSet samples.
type Acl_Interface_EgressAclSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_Interface_EgressAclSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_Interface_EgressAclSetWatcher) Await(t testing.TB) (*QualifiedAcl_Interface_EgressAclSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_Interface_EgressAclSet_AclEntry is a *Acl_Interface_EgressAclSet_AclEntry with a corresponding timestamp.
type QualifiedAcl_Interface_EgressAclSet_AclEntry struct {
	*genutil.Metadata
	val     *Acl_Interface_EgressAclSet_AclEntry // val is the sample value.
	present bool
}

func (q *QualifiedAcl_Interface_EgressAclSet_AclEntry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_Interface_EgressAclSet_AclEntry sample, erroring out if not present.
func (q *QualifiedAcl_Interface_EgressAclSet_AclEntry) Val(t testing.TB) *Acl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_Interface_EgressAclSet_AclEntry sample.
func (q *QualifiedAcl_Interface_EgressAclSet_AclEntry) SetVal(v *Acl_Interface_EgressAclSet_AclEntry) *QualifiedAcl_Interface_EgressAclSet_AclEntry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_Interface_EgressAclSet_AclEntry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_Interface_EgressAclSet_AclEntry is a telemetry Collection whose Await method returns a slice of *Acl_Interface_EgressAclSet_AclEntry samples.
type CollectionAcl_Interface_EgressAclSet_AclEntry struct {
	W    *Acl_Interface_EgressAclSet_AclEntryWatcher
	Data []*QualifiedAcl_Interface_EgressAclSet_AclEntry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_Interface_EgressAclSet_AclEntry) Await(t testing.TB) []*QualifiedAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_Interface_EgressAclSet_AclEntryWatcher observes a stream of *Acl_Interface_EgressAclSet_AclEntry samples.
type Acl_Interface_EgressAclSet_AclEntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_Interface_EgressAclSet_AclEntry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_Interface_EgressAclSet_AclEntryWatcher) Await(t testing.TB) (*QualifiedAcl_Interface_EgressAclSet_AclEntry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_Interface_IngressAclSet is a *Acl_Interface_IngressAclSet with a corresponding timestamp.
type QualifiedAcl_Interface_IngressAclSet struct {
	*genutil.Metadata
	val     *Acl_Interface_IngressAclSet // val is the sample value.
	present bool
}

func (q *QualifiedAcl_Interface_IngressAclSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_Interface_IngressAclSet sample, erroring out if not present.
func (q *QualifiedAcl_Interface_IngressAclSet) Val(t testing.TB) *Acl_Interface_IngressAclSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_Interface_IngressAclSet sample.
func (q *QualifiedAcl_Interface_IngressAclSet) SetVal(v *Acl_Interface_IngressAclSet) *QualifiedAcl_Interface_IngressAclSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_Interface_IngressAclSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_Interface_IngressAclSet is a telemetry Collection whose Await method returns a slice of *Acl_Interface_IngressAclSet samples.
type CollectionAcl_Interface_IngressAclSet struct {
	W    *Acl_Interface_IngressAclSetWatcher
	Data []*QualifiedAcl_Interface_IngressAclSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_Interface_IngressAclSet) Await(t testing.TB) []*QualifiedAcl_Interface_IngressAclSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_Interface_IngressAclSetWatcher observes a stream of *Acl_Interface_IngressAclSet samples.
type Acl_Interface_IngressAclSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_Interface_IngressAclSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_Interface_IngressAclSetWatcher) Await(t testing.TB) (*QualifiedAcl_Interface_IngressAclSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_Interface_IngressAclSet_AclEntry is a *Acl_Interface_IngressAclSet_AclEntry with a corresponding timestamp.
type QualifiedAcl_Interface_IngressAclSet_AclEntry struct {
	*genutil.Metadata
	val     *Acl_Interface_IngressAclSet_AclEntry // val is the sample value.
	present bool
}

func (q *QualifiedAcl_Interface_IngressAclSet_AclEntry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_Interface_IngressAclSet_AclEntry sample, erroring out if not present.
func (q *QualifiedAcl_Interface_IngressAclSet_AclEntry) Val(t testing.TB) *Acl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_Interface_IngressAclSet_AclEntry sample.
func (q *QualifiedAcl_Interface_IngressAclSet_AclEntry) SetVal(v *Acl_Interface_IngressAclSet_AclEntry) *QualifiedAcl_Interface_IngressAclSet_AclEntry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_Interface_IngressAclSet_AclEntry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_Interface_IngressAclSet_AclEntry is a telemetry Collection whose Await method returns a slice of *Acl_Interface_IngressAclSet_AclEntry samples.
type CollectionAcl_Interface_IngressAclSet_AclEntry struct {
	W    *Acl_Interface_IngressAclSet_AclEntryWatcher
	Data []*QualifiedAcl_Interface_IngressAclSet_AclEntry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_Interface_IngressAclSet_AclEntry) Await(t testing.TB) []*QualifiedAcl_Interface_IngressAclSet_AclEntry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_Interface_IngressAclSet_AclEntryWatcher observes a stream of *Acl_Interface_IngressAclSet_AclEntry samples.
type Acl_Interface_IngressAclSet_AclEntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_Interface_IngressAclSet_AclEntry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_Interface_IngressAclSet_AclEntryWatcher) Await(t testing.TB) (*QualifiedAcl_Interface_IngressAclSet_AclEntry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_Interface_InterfaceRef is a *Acl_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedAcl_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *Acl_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedAcl_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Acl_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedAcl_Interface_InterfaceRef) Val(t testing.TB) *Acl_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Acl_Interface_InterfaceRef sample.
func (q *QualifiedAcl_Interface_InterfaceRef) SetVal(v *Acl_Interface_InterfaceRef) *QualifiedAcl_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Acl_Interface_InterfaceRef samples.
type CollectionAcl_Interface_InterfaceRef struct {
	W    *Acl_Interface_InterfaceRefWatcher
	Data []*QualifiedAcl_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedAcl_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_Interface_InterfaceRefWatcher observes a stream of *Acl_Interface_InterfaceRef samples.
type Acl_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedAcl_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent is a *Component with a corresponding timestamp.
type QualifiedComponent struct {
	*genutil.Metadata
	val     *Component // val is the sample value.
	present bool
}

func (q *QualifiedComponent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component sample, erroring out if not present.
func (q *QualifiedComponent) Val(t testing.TB) *Component {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component sample.
func (q *QualifiedComponent) SetVal(v *Component) *QualifiedComponent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent is a telemetry Collection whose Await method returns a slice of *Component samples.
type CollectionComponent struct {
	W    *ComponentWatcher
	Data []*QualifiedComponent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent) Await(t testing.TB) []*QualifiedComponent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// ComponentWatcher observes a stream of *Component samples.
type ComponentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *ComponentWatcher) Await(t testing.TB) (*QualifiedComponent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Backplane is a *Component_Backplane with a corresponding timestamp.
type QualifiedComponent_Backplane struct {
	*genutil.Metadata
	val     *Component_Backplane // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Backplane) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Backplane sample, erroring out if not present.
func (q *QualifiedComponent_Backplane) Val(t testing.TB) *Component_Backplane {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Backplane sample.
func (q *QualifiedComponent_Backplane) SetVal(v *Component_Backplane) *QualifiedComponent_Backplane {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Backplane) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Backplane is a telemetry Collection whose Await method returns a slice of *Component_Backplane samples.
type CollectionComponent_Backplane struct {
	W    *Component_BackplaneWatcher
	Data []*QualifiedComponent_Backplane
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Backplane) Await(t testing.TB) []*QualifiedComponent_Backplane {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_BackplaneWatcher observes a stream of *Component_Backplane samples.
type Component_BackplaneWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Backplane
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_BackplaneWatcher) Await(t testing.TB) (*QualifiedComponent_Backplane, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Chassis is a *Component_Chassis with a corresponding timestamp.
type QualifiedComponent_Chassis struct {
	*genutil.Metadata
	val     *Component_Chassis // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Chassis) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Chassis sample, erroring out if not present.
func (q *QualifiedComponent_Chassis) Val(t testing.TB) *Component_Chassis {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Chassis sample.
func (q *QualifiedComponent_Chassis) SetVal(v *Component_Chassis) *QualifiedComponent_Chassis {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Chassis) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Chassis is a telemetry Collection whose Await method returns a slice of *Component_Chassis samples.
type CollectionComponent_Chassis struct {
	W    *Component_ChassisWatcher
	Data []*QualifiedComponent_Chassis
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Chassis) Await(t testing.TB) []*QualifiedComponent_Chassis {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_ChassisWatcher observes a stream of *Component_Chassis samples.
type Component_ChassisWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Chassis
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_ChassisWatcher) Await(t testing.TB) (*QualifiedComponent_Chassis, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Cpu is a *Component_Cpu with a corresponding timestamp.
type QualifiedComponent_Cpu struct {
	*genutil.Metadata
	val     *Component_Cpu // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Cpu) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Cpu sample, erroring out if not present.
func (q *QualifiedComponent_Cpu) Val(t testing.TB) *Component_Cpu {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Cpu sample.
func (q *QualifiedComponent_Cpu) SetVal(v *Component_Cpu) *QualifiedComponent_Cpu {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Cpu) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Cpu is a telemetry Collection whose Await method returns a slice of *Component_Cpu samples.
type CollectionComponent_Cpu struct {
	W    *Component_CpuWatcher
	Data []*QualifiedComponent_Cpu
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Cpu) Await(t testing.TB) []*QualifiedComponent_Cpu {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_CpuWatcher observes a stream of *Component_Cpu samples.
type Component_CpuWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Cpu
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_CpuWatcher) Await(t testing.TB) (*QualifiedComponent_Cpu, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Cpu_Utilization is a *Component_Cpu_Utilization with a corresponding timestamp.
type QualifiedComponent_Cpu_Utilization struct {
	*genutil.Metadata
	val     *Component_Cpu_Utilization // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Cpu_Utilization) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Cpu_Utilization sample, erroring out if not present.
func (q *QualifiedComponent_Cpu_Utilization) Val(t testing.TB) *Component_Cpu_Utilization {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Cpu_Utilization sample.
func (q *QualifiedComponent_Cpu_Utilization) SetVal(v *Component_Cpu_Utilization) *QualifiedComponent_Cpu_Utilization {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Cpu_Utilization) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Cpu_Utilization is a telemetry Collection whose Await method returns a slice of *Component_Cpu_Utilization samples.
type CollectionComponent_Cpu_Utilization struct {
	W    *Component_Cpu_UtilizationWatcher
	Data []*QualifiedComponent_Cpu_Utilization
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Cpu_Utilization) Await(t testing.TB) []*QualifiedComponent_Cpu_Utilization {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Cpu_UtilizationWatcher observes a stream of *Component_Cpu_Utilization samples.
type Component_Cpu_UtilizationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Cpu_Utilization
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Cpu_UtilizationWatcher) Await(t testing.TB) (*QualifiedComponent_Cpu_Utilization, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Fabric is a *Component_Fabric with a corresponding timestamp.
type QualifiedComponent_Fabric struct {
	*genutil.Metadata
	val     *Component_Fabric // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Fabric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Fabric sample, erroring out if not present.
func (q *QualifiedComponent_Fabric) Val(t testing.TB) *Component_Fabric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Fabric sample.
func (q *QualifiedComponent_Fabric) SetVal(v *Component_Fabric) *QualifiedComponent_Fabric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Fabric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Fabric is a telemetry Collection whose Await method returns a slice of *Component_Fabric samples.
type CollectionComponent_Fabric struct {
	W    *Component_FabricWatcher
	Data []*QualifiedComponent_Fabric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Fabric) Await(t testing.TB) []*QualifiedComponent_Fabric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_FabricWatcher observes a stream of *Component_Fabric samples.
type Component_FabricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Fabric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_FabricWatcher) Await(t testing.TB) (*QualifiedComponent_Fabric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Fan is a *Component_Fan with a corresponding timestamp.
type QualifiedComponent_Fan struct {
	*genutil.Metadata
	val     *Component_Fan // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Fan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Fan sample, erroring out if not present.
func (q *QualifiedComponent_Fan) Val(t testing.TB) *Component_Fan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Fan sample.
func (q *QualifiedComponent_Fan) SetVal(v *Component_Fan) *QualifiedComponent_Fan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Fan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Fan is a telemetry Collection whose Await method returns a slice of *Component_Fan samples.
type CollectionComponent_Fan struct {
	W    *Component_FanWatcher
	Data []*QualifiedComponent_Fan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Fan) Await(t testing.TB) []*QualifiedComponent_Fan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_FanWatcher observes a stream of *Component_Fan samples.
type Component_FanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Fan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_FanWatcher) Await(t testing.TB) (*QualifiedComponent_Fan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_IntegratedCircuit is a *Component_IntegratedCircuit with a corresponding timestamp.
type QualifiedComponent_IntegratedCircuit struct {
	*genutil.Metadata
	val     *Component_IntegratedCircuit // val is the sample value.
	present bool
}

func (q *QualifiedComponent_IntegratedCircuit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_IntegratedCircuit sample, erroring out if not present.
func (q *QualifiedComponent_IntegratedCircuit) Val(t testing.TB) *Component_IntegratedCircuit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_IntegratedCircuit sample.
func (q *QualifiedComponent_IntegratedCircuit) SetVal(v *Component_IntegratedCircuit) *QualifiedComponent_IntegratedCircuit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_IntegratedCircuit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_IntegratedCircuit is a telemetry Collection whose Await method returns a slice of *Component_IntegratedCircuit samples.
type CollectionComponent_IntegratedCircuit struct {
	W    *Component_IntegratedCircuitWatcher
	Data []*QualifiedComponent_IntegratedCircuit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_IntegratedCircuit) Await(t testing.TB) []*QualifiedComponent_IntegratedCircuit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_IntegratedCircuitWatcher observes a stream of *Component_IntegratedCircuit samples.
type Component_IntegratedCircuitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_IntegratedCircuit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_IntegratedCircuitWatcher) Await(t testing.TB) (*QualifiedComponent_IntegratedCircuit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity is a *Component_IntegratedCircuit_BackplaneFacingCapacity with a corresponding timestamp.
type QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity struct {
	*genutil.Metadata
	val     *Component_IntegratedCircuit_BackplaneFacingCapacity // val is the sample value.
	present bool
}

func (q *QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_IntegratedCircuit_BackplaneFacingCapacity sample, erroring out if not present.
func (q *QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) Val(t testing.TB) *Component_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_IntegratedCircuit_BackplaneFacingCapacity sample.
func (q *QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) SetVal(v *Component_IntegratedCircuit_BackplaneFacingCapacity) *QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_IntegratedCircuit_BackplaneFacingCapacity is a telemetry Collection whose Await method returns a slice of *Component_IntegratedCircuit_BackplaneFacingCapacity samples.
type CollectionComponent_IntegratedCircuit_BackplaneFacingCapacity struct {
	W    *Component_IntegratedCircuit_BackplaneFacingCapacityWatcher
	Data []*QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_IntegratedCircuit_BackplaneFacingCapacity) Await(t testing.TB) []*QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_IntegratedCircuit_BackplaneFacingCapacityWatcher observes a stream of *Component_IntegratedCircuit_BackplaneFacingCapacity samples.
type Component_IntegratedCircuit_BackplaneFacingCapacityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_IntegratedCircuit_BackplaneFacingCapacityWatcher) Await(t testing.TB) (*QualifiedComponent_IntegratedCircuit_BackplaneFacingCapacity, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_IntegratedCircuit_Memory is a *Component_IntegratedCircuit_Memory with a corresponding timestamp.
type QualifiedComponent_IntegratedCircuit_Memory struct {
	*genutil.Metadata
	val     *Component_IntegratedCircuit_Memory // val is the sample value.
	present bool
}

func (q *QualifiedComponent_IntegratedCircuit_Memory) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_IntegratedCircuit_Memory sample, erroring out if not present.
func (q *QualifiedComponent_IntegratedCircuit_Memory) Val(t testing.TB) *Component_IntegratedCircuit_Memory {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_IntegratedCircuit_Memory sample.
func (q *QualifiedComponent_IntegratedCircuit_Memory) SetVal(v *Component_IntegratedCircuit_Memory) *QualifiedComponent_IntegratedCircuit_Memory {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_IntegratedCircuit_Memory) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_IntegratedCircuit_Memory is a telemetry Collection whose Await method returns a slice of *Component_IntegratedCircuit_Memory samples.
type CollectionComponent_IntegratedCircuit_Memory struct {
	W    *Component_IntegratedCircuit_MemoryWatcher
	Data []*QualifiedComponent_IntegratedCircuit_Memory
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_IntegratedCircuit_Memory) Await(t testing.TB) []*QualifiedComponent_IntegratedCircuit_Memory {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_IntegratedCircuit_MemoryWatcher observes a stream of *Component_IntegratedCircuit_Memory samples.
type Component_IntegratedCircuit_MemoryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_IntegratedCircuit_Memory
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_IntegratedCircuit_MemoryWatcher) Await(t testing.TB) (*QualifiedComponent_IntegratedCircuit_Memory, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_LastSwitchoverReason is a *Component_LastSwitchoverReason with a corresponding timestamp.
type QualifiedComponent_LastSwitchoverReason struct {
	*genutil.Metadata
	val     *Component_LastSwitchoverReason // val is the sample value.
	present bool
}

func (q *QualifiedComponent_LastSwitchoverReason) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_LastSwitchoverReason sample, erroring out if not present.
func (q *QualifiedComponent_LastSwitchoverReason) Val(t testing.TB) *Component_LastSwitchoverReason {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_LastSwitchoverReason sample.
func (q *QualifiedComponent_LastSwitchoverReason) SetVal(v *Component_LastSwitchoverReason) *QualifiedComponent_LastSwitchoverReason {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_LastSwitchoverReason) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_LastSwitchoverReason is a telemetry Collection whose Await method returns a slice of *Component_LastSwitchoverReason samples.
type CollectionComponent_LastSwitchoverReason struct {
	W    *Component_LastSwitchoverReasonWatcher
	Data []*QualifiedComponent_LastSwitchoverReason
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_LastSwitchoverReason) Await(t testing.TB) []*QualifiedComponent_LastSwitchoverReason {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_LastSwitchoverReasonWatcher observes a stream of *Component_LastSwitchoverReason samples.
type Component_LastSwitchoverReasonWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_LastSwitchoverReason
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_LastSwitchoverReasonWatcher) Await(t testing.TB) (*QualifiedComponent_LastSwitchoverReason, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Memory is a *Component_Memory with a corresponding timestamp.
type QualifiedComponent_Memory struct {
	*genutil.Metadata
	val     *Component_Memory // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Memory) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Memory sample, erroring out if not present.
func (q *QualifiedComponent_Memory) Val(t testing.TB) *Component_Memory {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Memory sample.
func (q *QualifiedComponent_Memory) SetVal(v *Component_Memory) *QualifiedComponent_Memory {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Memory) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Memory is a telemetry Collection whose Await method returns a slice of *Component_Memory samples.
type CollectionComponent_Memory struct {
	W    *Component_MemoryWatcher
	Data []*QualifiedComponent_Memory
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Memory) Await(t testing.TB) []*QualifiedComponent_Memory {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_MemoryWatcher observes a stream of *Component_Memory samples.
type Component_MemoryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Memory
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_MemoryWatcher) Await(t testing.TB) (*QualifiedComponent_Memory, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Pcie is a *Component_Pcie with a corresponding timestamp.
type QualifiedComponent_Pcie struct {
	*genutil.Metadata
	val     *Component_Pcie // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Pcie) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Pcie sample, erroring out if not present.
func (q *QualifiedComponent_Pcie) Val(t testing.TB) *Component_Pcie {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Pcie sample.
func (q *QualifiedComponent_Pcie) SetVal(v *Component_Pcie) *QualifiedComponent_Pcie {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Pcie) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Pcie is a telemetry Collection whose Await method returns a slice of *Component_Pcie samples.
type CollectionComponent_Pcie struct {
	W    *Component_PcieWatcher
	Data []*QualifiedComponent_Pcie
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Pcie) Await(t testing.TB) []*QualifiedComponent_Pcie {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_PcieWatcher observes a stream of *Component_Pcie samples.
type Component_PcieWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Pcie
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_PcieWatcher) Await(t testing.TB) (*QualifiedComponent_Pcie, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Pcie_CorrectableErrors is a *Component_Pcie_CorrectableErrors with a corresponding timestamp.
type QualifiedComponent_Pcie_CorrectableErrors struct {
	*genutil.Metadata
	val     *Component_Pcie_CorrectableErrors // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Pcie_CorrectableErrors) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Pcie_CorrectableErrors sample, erroring out if not present.
func (q *QualifiedComponent_Pcie_CorrectableErrors) Val(t testing.TB) *Component_Pcie_CorrectableErrors {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Pcie_CorrectableErrors sample.
func (q *QualifiedComponent_Pcie_CorrectableErrors) SetVal(v *Component_Pcie_CorrectableErrors) *QualifiedComponent_Pcie_CorrectableErrors {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Pcie_CorrectableErrors) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Pcie_CorrectableErrors is a telemetry Collection whose Await method returns a slice of *Component_Pcie_CorrectableErrors samples.
type CollectionComponent_Pcie_CorrectableErrors struct {
	W    *Component_Pcie_CorrectableErrorsWatcher
	Data []*QualifiedComponent_Pcie_CorrectableErrors
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Pcie_CorrectableErrors) Await(t testing.TB) []*QualifiedComponent_Pcie_CorrectableErrors {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Pcie_CorrectableErrorsWatcher observes a stream of *Component_Pcie_CorrectableErrors samples.
type Component_Pcie_CorrectableErrorsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Pcie_CorrectableErrors
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Pcie_CorrectableErrorsWatcher) Await(t testing.TB) (*QualifiedComponent_Pcie_CorrectableErrors, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Pcie_FatalErrors is a *Component_Pcie_FatalErrors with a corresponding timestamp.
type QualifiedComponent_Pcie_FatalErrors struct {
	*genutil.Metadata
	val     *Component_Pcie_FatalErrors // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Pcie_FatalErrors) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Pcie_FatalErrors sample, erroring out if not present.
func (q *QualifiedComponent_Pcie_FatalErrors) Val(t testing.TB) *Component_Pcie_FatalErrors {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Pcie_FatalErrors sample.
func (q *QualifiedComponent_Pcie_FatalErrors) SetVal(v *Component_Pcie_FatalErrors) *QualifiedComponent_Pcie_FatalErrors {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Pcie_FatalErrors) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Pcie_FatalErrors is a telemetry Collection whose Await method returns a slice of *Component_Pcie_FatalErrors samples.
type CollectionComponent_Pcie_FatalErrors struct {
	W    *Component_Pcie_FatalErrorsWatcher
	Data []*QualifiedComponent_Pcie_FatalErrors
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Pcie_FatalErrors) Await(t testing.TB) []*QualifiedComponent_Pcie_FatalErrors {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Pcie_FatalErrorsWatcher observes a stream of *Component_Pcie_FatalErrors samples.
type Component_Pcie_FatalErrorsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Pcie_FatalErrors
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Pcie_FatalErrorsWatcher) Await(t testing.TB) (*QualifiedComponent_Pcie_FatalErrors, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Pcie_NonFatalErrors is a *Component_Pcie_NonFatalErrors with a corresponding timestamp.
type QualifiedComponent_Pcie_NonFatalErrors struct {
	*genutil.Metadata
	val     *Component_Pcie_NonFatalErrors // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Pcie_NonFatalErrors) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Pcie_NonFatalErrors sample, erroring out if not present.
func (q *QualifiedComponent_Pcie_NonFatalErrors) Val(t testing.TB) *Component_Pcie_NonFatalErrors {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Pcie_NonFatalErrors sample.
func (q *QualifiedComponent_Pcie_NonFatalErrors) SetVal(v *Component_Pcie_NonFatalErrors) *QualifiedComponent_Pcie_NonFatalErrors {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Pcie_NonFatalErrors) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Pcie_NonFatalErrors is a telemetry Collection whose Await method returns a slice of *Component_Pcie_NonFatalErrors samples.
type CollectionComponent_Pcie_NonFatalErrors struct {
	W    *Component_Pcie_NonFatalErrorsWatcher
	Data []*QualifiedComponent_Pcie_NonFatalErrors
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Pcie_NonFatalErrors) Await(t testing.TB) []*QualifiedComponent_Pcie_NonFatalErrors {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Pcie_NonFatalErrorsWatcher observes a stream of *Component_Pcie_NonFatalErrors samples.
type Component_Pcie_NonFatalErrorsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Pcie_NonFatalErrors
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Pcie_NonFatalErrorsWatcher) Await(t testing.TB) (*QualifiedComponent_Pcie_NonFatalErrors, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Port is a *Component_Port with a corresponding timestamp.
type QualifiedComponent_Port struct {
	*genutil.Metadata
	val     *Component_Port // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Port) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Port sample, erroring out if not present.
func (q *QualifiedComponent_Port) Val(t testing.TB) *Component_Port {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Port sample.
func (q *QualifiedComponent_Port) SetVal(v *Component_Port) *QualifiedComponent_Port {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Port) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Port is a telemetry Collection whose Await method returns a slice of *Component_Port samples.
type CollectionComponent_Port struct {
	W    *Component_PortWatcher
	Data []*QualifiedComponent_Port
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Port) Await(t testing.TB) []*QualifiedComponent_Port {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_PortWatcher observes a stream of *Component_Port samples.
type Component_PortWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Port
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_PortWatcher) Await(t testing.TB) (*QualifiedComponent_Port, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Port_BreakoutMode is a *Component_Port_BreakoutMode with a corresponding timestamp.
type QualifiedComponent_Port_BreakoutMode struct {
	*genutil.Metadata
	val     *Component_Port_BreakoutMode // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Port_BreakoutMode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Port_BreakoutMode sample, erroring out if not present.
func (q *QualifiedComponent_Port_BreakoutMode) Val(t testing.TB) *Component_Port_BreakoutMode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Port_BreakoutMode sample.
func (q *QualifiedComponent_Port_BreakoutMode) SetVal(v *Component_Port_BreakoutMode) *QualifiedComponent_Port_BreakoutMode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Port_BreakoutMode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Port_BreakoutMode is a telemetry Collection whose Await method returns a slice of *Component_Port_BreakoutMode samples.
type CollectionComponent_Port_BreakoutMode struct {
	W    *Component_Port_BreakoutModeWatcher
	Data []*QualifiedComponent_Port_BreakoutMode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Port_BreakoutMode) Await(t testing.TB) []*QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Port_BreakoutModeWatcher observes a stream of *Component_Port_BreakoutMode samples.
type Component_Port_BreakoutModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Port_BreakoutMode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Port_BreakoutModeWatcher) Await(t testing.TB) (*QualifiedComponent_Port_BreakoutMode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Port_BreakoutMode_Group is a *Component_Port_BreakoutMode_Group with a corresponding timestamp.
type QualifiedComponent_Port_BreakoutMode_Group struct {
	*genutil.Metadata
	val     *Component_Port_BreakoutMode_Group // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Port_BreakoutMode_Group) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Port_BreakoutMode_Group sample, erroring out if not present.
func (q *QualifiedComponent_Port_BreakoutMode_Group) Val(t testing.TB) *Component_Port_BreakoutMode_Group {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Port_BreakoutMode_Group sample.
func (q *QualifiedComponent_Port_BreakoutMode_Group) SetVal(v *Component_Port_BreakoutMode_Group) *QualifiedComponent_Port_BreakoutMode_Group {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Port_BreakoutMode_Group) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Port_BreakoutMode_Group is a telemetry Collection whose Await method returns a slice of *Component_Port_BreakoutMode_Group samples.
type CollectionComponent_Port_BreakoutMode_Group struct {
	W    *Component_Port_BreakoutMode_GroupWatcher
	Data []*QualifiedComponent_Port_BreakoutMode_Group
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Port_BreakoutMode_Group) Await(t testing.TB) []*QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Port_BreakoutMode_GroupWatcher observes a stream of *Component_Port_BreakoutMode_Group samples.
type Component_Port_BreakoutMode_GroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Port_BreakoutMode_Group
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Port_BreakoutMode_GroupWatcher) Await(t testing.TB) (*QualifiedComponent_Port_BreakoutMode_Group, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_PowerSupply is a *Component_PowerSupply with a corresponding timestamp.
type QualifiedComponent_PowerSupply struct {
	*genutil.Metadata
	val     *Component_PowerSupply // val is the sample value.
	present bool
}

func (q *QualifiedComponent_PowerSupply) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_PowerSupply sample, erroring out if not present.
func (q *QualifiedComponent_PowerSupply) Val(t testing.TB) *Component_PowerSupply {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_PowerSupply sample.
func (q *QualifiedComponent_PowerSupply) SetVal(v *Component_PowerSupply) *QualifiedComponent_PowerSupply {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_PowerSupply) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_PowerSupply is a telemetry Collection whose Await method returns a slice of *Component_PowerSupply samples.
type CollectionComponent_PowerSupply struct {
	W    *Component_PowerSupplyWatcher
	Data []*QualifiedComponent_PowerSupply
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_PowerSupply) Await(t testing.TB) []*QualifiedComponent_PowerSupply {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_PowerSupplyWatcher observes a stream of *Component_PowerSupply samples.
type Component_PowerSupplyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_PowerSupply
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_PowerSupplyWatcher) Await(t testing.TB) (*QualifiedComponent_PowerSupply, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Property is a *Component_Property with a corresponding timestamp.
type QualifiedComponent_Property struct {
	*genutil.Metadata
	val     *Component_Property // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Property) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Property sample, erroring out if not present.
func (q *QualifiedComponent_Property) Val(t testing.TB) *Component_Property {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Property sample.
func (q *QualifiedComponent_Property) SetVal(v *Component_Property) *QualifiedComponent_Property {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Property) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Property is a telemetry Collection whose Await method returns a slice of *Component_Property samples.
type CollectionComponent_Property struct {
	W    *Component_PropertyWatcher
	Data []*QualifiedComponent_Property
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Property) Await(t testing.TB) []*QualifiedComponent_Property {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_PropertyWatcher observes a stream of *Component_Property samples.
type Component_PropertyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Property
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_PropertyWatcher) Await(t testing.TB) (*QualifiedComponent_Property, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_SoftwareModule is a *Component_SoftwareModule with a corresponding timestamp.
type QualifiedComponent_SoftwareModule struct {
	*genutil.Metadata
	val     *Component_SoftwareModule // val is the sample value.
	present bool
}

func (q *QualifiedComponent_SoftwareModule) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_SoftwareModule sample, erroring out if not present.
func (q *QualifiedComponent_SoftwareModule) Val(t testing.TB) *Component_SoftwareModule {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_SoftwareModule sample.
func (q *QualifiedComponent_SoftwareModule) SetVal(v *Component_SoftwareModule) *QualifiedComponent_SoftwareModule {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_SoftwareModule) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_SoftwareModule is a telemetry Collection whose Await method returns a slice of *Component_SoftwareModule samples.
type CollectionComponent_SoftwareModule struct {
	W    *Component_SoftwareModuleWatcher
	Data []*QualifiedComponent_SoftwareModule
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_SoftwareModule) Await(t testing.TB) []*QualifiedComponent_SoftwareModule {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_SoftwareModuleWatcher observes a stream of *Component_SoftwareModule samples.
type Component_SoftwareModuleWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_SoftwareModule
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_SoftwareModuleWatcher) Await(t testing.TB) (*QualifiedComponent_SoftwareModule, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Storage is a *Component_Storage with a corresponding timestamp.
type QualifiedComponent_Storage struct {
	*genutil.Metadata
	val     *Component_Storage // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Storage) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Storage sample, erroring out if not present.
func (q *QualifiedComponent_Storage) Val(t testing.TB) *Component_Storage {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Storage sample.
func (q *QualifiedComponent_Storage) SetVal(v *Component_Storage) *QualifiedComponent_Storage {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Storage) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Storage is a telemetry Collection whose Await method returns a slice of *Component_Storage samples.
type CollectionComponent_Storage struct {
	W    *Component_StorageWatcher
	Data []*QualifiedComponent_Storage
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Storage) Await(t testing.TB) []*QualifiedComponent_Storage {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_StorageWatcher observes a stream of *Component_Storage samples.
type Component_StorageWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Storage
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_StorageWatcher) Await(t testing.TB) (*QualifiedComponent_Storage, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Subcomponent is a *Component_Subcomponent with a corresponding timestamp.
type QualifiedComponent_Subcomponent struct {
	*genutil.Metadata
	val     *Component_Subcomponent // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Subcomponent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Subcomponent sample, erroring out if not present.
func (q *QualifiedComponent_Subcomponent) Val(t testing.TB) *Component_Subcomponent {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Subcomponent sample.
func (q *QualifiedComponent_Subcomponent) SetVal(v *Component_Subcomponent) *QualifiedComponent_Subcomponent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Subcomponent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Subcomponent is a telemetry Collection whose Await method returns a slice of *Component_Subcomponent samples.
type CollectionComponent_Subcomponent struct {
	W    *Component_SubcomponentWatcher
	Data []*QualifiedComponent_Subcomponent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Subcomponent) Await(t testing.TB) []*QualifiedComponent_Subcomponent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_SubcomponentWatcher observes a stream of *Component_Subcomponent samples.
type Component_SubcomponentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Subcomponent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_SubcomponentWatcher) Await(t testing.TB) (*QualifiedComponent_Subcomponent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Temperature is a *Component_Temperature with a corresponding timestamp.
type QualifiedComponent_Temperature struct {
	*genutil.Metadata
	val     *Component_Temperature // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Temperature) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Temperature sample, erroring out if not present.
func (q *QualifiedComponent_Temperature) Val(t testing.TB) *Component_Temperature {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Temperature sample.
func (q *QualifiedComponent_Temperature) SetVal(v *Component_Temperature) *QualifiedComponent_Temperature {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Temperature) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Temperature is a telemetry Collection whose Await method returns a slice of *Component_Temperature samples.
type CollectionComponent_Temperature struct {
	W    *Component_TemperatureWatcher
	Data []*QualifiedComponent_Temperature
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Temperature) Await(t testing.TB) []*QualifiedComponent_Temperature {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_TemperatureWatcher observes a stream of *Component_Temperature samples.
type Component_TemperatureWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Temperature
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_TemperatureWatcher) Await(t testing.TB) (*QualifiedComponent_Temperature, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver is a *Component_Transceiver with a corresponding timestamp.
type QualifiedComponent_Transceiver struct {
	*genutil.Metadata
	val     *Component_Transceiver // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver) Val(t testing.TB) *Component_Transceiver {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver sample.
func (q *QualifiedComponent_Transceiver) SetVal(v *Component_Transceiver) *QualifiedComponent_Transceiver {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver is a telemetry Collection whose Await method returns a slice of *Component_Transceiver samples.
type CollectionComponent_Transceiver struct {
	W    *Component_TransceiverWatcher
	Data []*QualifiedComponent_Transceiver
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver) Await(t testing.TB) []*QualifiedComponent_Transceiver {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_TransceiverWatcher observes a stream of *Component_Transceiver samples.
type Component_TransceiverWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_TransceiverWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_Channel is a *Component_Transceiver_Channel with a corresponding timestamp.
type QualifiedComponent_Transceiver_Channel struct {
	*genutil.Metadata
	val     *Component_Transceiver_Channel // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_Channel) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_Channel sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_Channel) Val(t testing.TB) *Component_Transceiver_Channel {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_Channel sample.
func (q *QualifiedComponent_Transceiver_Channel) SetVal(v *Component_Transceiver_Channel) *QualifiedComponent_Transceiver_Channel {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_Channel) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_Channel is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_Channel samples.
type CollectionComponent_Transceiver_Channel struct {
	W    *Component_Transceiver_ChannelWatcher
	Data []*QualifiedComponent_Transceiver_Channel
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_Channel) Await(t testing.TB) []*QualifiedComponent_Transceiver_Channel {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_ChannelWatcher observes a stream of *Component_Transceiver_Channel samples.
type Component_Transceiver_ChannelWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_Channel
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_ChannelWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_Channel, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_Channel_InputPower is a *Component_Transceiver_Channel_InputPower with a corresponding timestamp.
type QualifiedComponent_Transceiver_Channel_InputPower struct {
	*genutil.Metadata
	val     *Component_Transceiver_Channel_InputPower // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_Channel_InputPower) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_Channel_InputPower sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_Channel_InputPower) Val(t testing.TB) *Component_Transceiver_Channel_InputPower {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_Channel_InputPower sample.
func (q *QualifiedComponent_Transceiver_Channel_InputPower) SetVal(v *Component_Transceiver_Channel_InputPower) *QualifiedComponent_Transceiver_Channel_InputPower {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_Channel_InputPower) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_Channel_InputPower is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_Channel_InputPower samples.
type CollectionComponent_Transceiver_Channel_InputPower struct {
	W    *Component_Transceiver_Channel_InputPowerWatcher
	Data []*QualifiedComponent_Transceiver_Channel_InputPower
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_Channel_InputPower) Await(t testing.TB) []*QualifiedComponent_Transceiver_Channel_InputPower {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_Channel_InputPowerWatcher observes a stream of *Component_Transceiver_Channel_InputPower samples.
type Component_Transceiver_Channel_InputPowerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_Channel_InputPower
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_Channel_InputPowerWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_Channel_InputPower, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_Channel_LaserBiasCurrent is a *Component_Transceiver_Channel_LaserBiasCurrent with a corresponding timestamp.
type QualifiedComponent_Transceiver_Channel_LaserBiasCurrent struct {
	*genutil.Metadata
	val     *Component_Transceiver_Channel_LaserBiasCurrent // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_Channel_LaserBiasCurrent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_Channel_LaserBiasCurrent sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_Channel_LaserBiasCurrent) Val(t testing.TB) *Component_Transceiver_Channel_LaserBiasCurrent {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_Channel_LaserBiasCurrent sample.
func (q *QualifiedComponent_Transceiver_Channel_LaserBiasCurrent) SetVal(v *Component_Transceiver_Channel_LaserBiasCurrent) *QualifiedComponent_Transceiver_Channel_LaserBiasCurrent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_Channel_LaserBiasCurrent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_Channel_LaserBiasCurrent is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_Channel_LaserBiasCurrent samples.
type CollectionComponent_Transceiver_Channel_LaserBiasCurrent struct {
	W    *Component_Transceiver_Channel_LaserBiasCurrentWatcher
	Data []*QualifiedComponent_Transceiver_Channel_LaserBiasCurrent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_Channel_LaserBiasCurrent) Await(t testing.TB) []*QualifiedComponent_Transceiver_Channel_LaserBiasCurrent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_Channel_LaserBiasCurrentWatcher observes a stream of *Component_Transceiver_Channel_LaserBiasCurrent samples.
type Component_Transceiver_Channel_LaserBiasCurrentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_Channel_LaserBiasCurrent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_Channel_LaserBiasCurrentWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_Channel_LaserBiasCurrent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_Channel_LaserTemperature is a *Component_Transceiver_Channel_LaserTemperature with a corresponding timestamp.
type QualifiedComponent_Transceiver_Channel_LaserTemperature struct {
	*genutil.Metadata
	val     *Component_Transceiver_Channel_LaserTemperature // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_Channel_LaserTemperature) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_Channel_LaserTemperature sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_Channel_LaserTemperature) Val(t testing.TB) *Component_Transceiver_Channel_LaserTemperature {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_Channel_LaserTemperature sample.
func (q *QualifiedComponent_Transceiver_Channel_LaserTemperature) SetVal(v *Component_Transceiver_Channel_LaserTemperature) *QualifiedComponent_Transceiver_Channel_LaserTemperature {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_Channel_LaserTemperature) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_Channel_LaserTemperature is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_Channel_LaserTemperature samples.
type CollectionComponent_Transceiver_Channel_LaserTemperature struct {
	W    *Component_Transceiver_Channel_LaserTemperatureWatcher
	Data []*QualifiedComponent_Transceiver_Channel_LaserTemperature
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_Channel_LaserTemperature) Await(t testing.TB) []*QualifiedComponent_Transceiver_Channel_LaserTemperature {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_Channel_LaserTemperatureWatcher observes a stream of *Component_Transceiver_Channel_LaserTemperature samples.
type Component_Transceiver_Channel_LaserTemperatureWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_Channel_LaserTemperature
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_Channel_LaserTemperatureWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_Channel_LaserTemperature, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_Channel_OutputPower is a *Component_Transceiver_Channel_OutputPower with a corresponding timestamp.
type QualifiedComponent_Transceiver_Channel_OutputPower struct {
	*genutil.Metadata
	val     *Component_Transceiver_Channel_OutputPower // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_Channel_OutputPower) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_Channel_OutputPower sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_Channel_OutputPower) Val(t testing.TB) *Component_Transceiver_Channel_OutputPower {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_Channel_OutputPower sample.
func (q *QualifiedComponent_Transceiver_Channel_OutputPower) SetVal(v *Component_Transceiver_Channel_OutputPower) *QualifiedComponent_Transceiver_Channel_OutputPower {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_Channel_OutputPower) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_Channel_OutputPower is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_Channel_OutputPower samples.
type CollectionComponent_Transceiver_Channel_OutputPower struct {
	W    *Component_Transceiver_Channel_OutputPowerWatcher
	Data []*QualifiedComponent_Transceiver_Channel_OutputPower
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_Channel_OutputPower) Await(t testing.TB) []*QualifiedComponent_Transceiver_Channel_OutputPower {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_Channel_OutputPowerWatcher observes a stream of *Component_Transceiver_Channel_OutputPower samples.
type Component_Transceiver_Channel_OutputPowerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_Channel_OutputPower
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_Channel_OutputPowerWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_Channel_OutputPower, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation is a *Component_Transceiver_Channel_TargetFrequencyDeviation with a corresponding timestamp.
type QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation struct {
	*genutil.Metadata
	val     *Component_Transceiver_Channel_TargetFrequencyDeviation // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_Channel_TargetFrequencyDeviation sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation) Val(t testing.TB) *Component_Transceiver_Channel_TargetFrequencyDeviation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_Channel_TargetFrequencyDeviation sample.
func (q *QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation) SetVal(v *Component_Transceiver_Channel_TargetFrequencyDeviation) *QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_Channel_TargetFrequencyDeviation is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_Channel_TargetFrequencyDeviation samples.
type CollectionComponent_Transceiver_Channel_TargetFrequencyDeviation struct {
	W    *Component_Transceiver_Channel_TargetFrequencyDeviationWatcher
	Data []*QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_Channel_TargetFrequencyDeviation) Await(t testing.TB) []*QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_Channel_TargetFrequencyDeviationWatcher observes a stream of *Component_Transceiver_Channel_TargetFrequencyDeviation samples.
type Component_Transceiver_Channel_TargetFrequencyDeviationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_Channel_TargetFrequencyDeviationWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_Channel_TargetFrequencyDeviation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_Channel_TecCurrent is a *Component_Transceiver_Channel_TecCurrent with a corresponding timestamp.
type QualifiedComponent_Transceiver_Channel_TecCurrent struct {
	*genutil.Metadata
	val     *Component_Transceiver_Channel_TecCurrent // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_Channel_TecCurrent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_Channel_TecCurrent sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_Channel_TecCurrent) Val(t testing.TB) *Component_Transceiver_Channel_TecCurrent {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_Channel_TecCurrent sample.
func (q *QualifiedComponent_Transceiver_Channel_TecCurrent) SetVal(v *Component_Transceiver_Channel_TecCurrent) *QualifiedComponent_Transceiver_Channel_TecCurrent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_Channel_TecCurrent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_Channel_TecCurrent is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_Channel_TecCurrent samples.
type CollectionComponent_Transceiver_Channel_TecCurrent struct {
	W    *Component_Transceiver_Channel_TecCurrentWatcher
	Data []*QualifiedComponent_Transceiver_Channel_TecCurrent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_Channel_TecCurrent) Await(t testing.TB) []*QualifiedComponent_Transceiver_Channel_TecCurrent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_Channel_TecCurrentWatcher observes a stream of *Component_Transceiver_Channel_TecCurrent samples.
type Component_Transceiver_Channel_TecCurrentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_Channel_TecCurrent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_Channel_TecCurrentWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_Channel_TecCurrent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_InputPower is a *Component_Transceiver_InputPower with a corresponding timestamp.
type QualifiedComponent_Transceiver_InputPower struct {
	*genutil.Metadata
	val     *Component_Transceiver_InputPower // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_InputPower) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_InputPower sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_InputPower) Val(t testing.TB) *Component_Transceiver_InputPower {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_InputPower sample.
func (q *QualifiedComponent_Transceiver_InputPower) SetVal(v *Component_Transceiver_InputPower) *QualifiedComponent_Transceiver_InputPower {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_InputPower) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_InputPower is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_InputPower samples.
type CollectionComponent_Transceiver_InputPower struct {
	W    *Component_Transceiver_InputPowerWatcher
	Data []*QualifiedComponent_Transceiver_InputPower
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_InputPower) Await(t testing.TB) []*QualifiedComponent_Transceiver_InputPower {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_InputPowerWatcher observes a stream of *Component_Transceiver_InputPower samples.
type Component_Transceiver_InputPowerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_InputPower
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_InputPowerWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_InputPower, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_LaserBiasCurrent is a *Component_Transceiver_LaserBiasCurrent with a corresponding timestamp.
type QualifiedComponent_Transceiver_LaserBiasCurrent struct {
	*genutil.Metadata
	val     *Component_Transceiver_LaserBiasCurrent // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_LaserBiasCurrent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_LaserBiasCurrent sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_LaserBiasCurrent) Val(t testing.TB) *Component_Transceiver_LaserBiasCurrent {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_LaserBiasCurrent sample.
func (q *QualifiedComponent_Transceiver_LaserBiasCurrent) SetVal(v *Component_Transceiver_LaserBiasCurrent) *QualifiedComponent_Transceiver_LaserBiasCurrent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_LaserBiasCurrent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_LaserBiasCurrent is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_LaserBiasCurrent samples.
type CollectionComponent_Transceiver_LaserBiasCurrent struct {
	W    *Component_Transceiver_LaserBiasCurrentWatcher
	Data []*QualifiedComponent_Transceiver_LaserBiasCurrent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_LaserBiasCurrent) Await(t testing.TB) []*QualifiedComponent_Transceiver_LaserBiasCurrent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_LaserBiasCurrentWatcher observes a stream of *Component_Transceiver_LaserBiasCurrent samples.
type Component_Transceiver_LaserBiasCurrentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_LaserBiasCurrent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_LaserBiasCurrentWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_LaserBiasCurrent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_OutputPower is a *Component_Transceiver_OutputPower with a corresponding timestamp.
type QualifiedComponent_Transceiver_OutputPower struct {
	*genutil.Metadata
	val     *Component_Transceiver_OutputPower // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_OutputPower) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_OutputPower sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_OutputPower) Val(t testing.TB) *Component_Transceiver_OutputPower {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_OutputPower sample.
func (q *QualifiedComponent_Transceiver_OutputPower) SetVal(v *Component_Transceiver_OutputPower) *QualifiedComponent_Transceiver_OutputPower {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_OutputPower) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_OutputPower is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_OutputPower samples.
type CollectionComponent_Transceiver_OutputPower struct {
	W    *Component_Transceiver_OutputPowerWatcher
	Data []*QualifiedComponent_Transceiver_OutputPower
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_OutputPower) Await(t testing.TB) []*QualifiedComponent_Transceiver_OutputPower {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_OutputPowerWatcher observes a stream of *Component_Transceiver_OutputPower samples.
type Component_Transceiver_OutputPowerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_OutputPower
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_OutputPowerWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_OutputPower, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_PostFecBer is a *Component_Transceiver_PostFecBer with a corresponding timestamp.
type QualifiedComponent_Transceiver_PostFecBer struct {
	*genutil.Metadata
	val     *Component_Transceiver_PostFecBer // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_PostFecBer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_PostFecBer sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_PostFecBer) Val(t testing.TB) *Component_Transceiver_PostFecBer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_PostFecBer sample.
func (q *QualifiedComponent_Transceiver_PostFecBer) SetVal(v *Component_Transceiver_PostFecBer) *QualifiedComponent_Transceiver_PostFecBer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_PostFecBer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_PostFecBer is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_PostFecBer samples.
type CollectionComponent_Transceiver_PostFecBer struct {
	W    *Component_Transceiver_PostFecBerWatcher
	Data []*QualifiedComponent_Transceiver_PostFecBer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_PostFecBer) Await(t testing.TB) []*QualifiedComponent_Transceiver_PostFecBer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_PostFecBerWatcher observes a stream of *Component_Transceiver_PostFecBer samples.
type Component_Transceiver_PostFecBerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_PostFecBer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_PostFecBerWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_PostFecBer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_PreFecBer is a *Component_Transceiver_PreFecBer with a corresponding timestamp.
type QualifiedComponent_Transceiver_PreFecBer struct {
	*genutil.Metadata
	val     *Component_Transceiver_PreFecBer // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_PreFecBer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_PreFecBer sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_PreFecBer) Val(t testing.TB) *Component_Transceiver_PreFecBer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_PreFecBer sample.
func (q *QualifiedComponent_Transceiver_PreFecBer) SetVal(v *Component_Transceiver_PreFecBer) *QualifiedComponent_Transceiver_PreFecBer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_PreFecBer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_PreFecBer is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_PreFecBer samples.
type CollectionComponent_Transceiver_PreFecBer struct {
	W    *Component_Transceiver_PreFecBerWatcher
	Data []*QualifiedComponent_Transceiver_PreFecBer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_PreFecBer) Await(t testing.TB) []*QualifiedComponent_Transceiver_PreFecBer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_PreFecBerWatcher observes a stream of *Component_Transceiver_PreFecBer samples.
type Component_Transceiver_PreFecBerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_PreFecBer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_PreFecBerWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_PreFecBer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Transceiver_SupplyVoltage is a *Component_Transceiver_SupplyVoltage with a corresponding timestamp.
type QualifiedComponent_Transceiver_SupplyVoltage struct {
	*genutil.Metadata
	val     *Component_Transceiver_SupplyVoltage // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Transceiver_SupplyVoltage) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Component_Transceiver_SupplyVoltage sample, erroring out if not present.
func (q *QualifiedComponent_Transceiver_SupplyVoltage) Val(t testing.TB) *Component_Transceiver_SupplyVoltage {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Component_Transceiver_SupplyVoltage sample.
func (q *QualifiedComponent_Transceiver_SupplyVoltage) SetVal(v *Component_Transceiver_SupplyVoltage) *QualifiedComponent_Transceiver_SupplyVoltage {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Transceiver_SupplyVoltage) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Transceiver_SupplyVoltage is a telemetry Collection whose Await method returns a slice of *Component_Transceiver_SupplyVoltage samples.
type CollectionComponent_Transceiver_SupplyVoltage struct {
	W    *Component_Transceiver_SupplyVoltageWatcher
	Data []*QualifiedComponent_Transceiver_SupplyVoltage
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Transceiver_SupplyVoltage) Await(t testing.TB) []*QualifiedComponent_Transceiver_SupplyVoltage {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Transceiver_SupplyVoltageWatcher observes a stream of *Component_Transceiver_SupplyVoltage samples.
type Component_Transceiver_SupplyVoltageWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Transceiver_SupplyVoltage
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Transceiver_SupplyVoltageWatcher) Await(t testing.TB) (*QualifiedComponent_Transceiver_SupplyVoltage, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedDevice is a *Device with a corresponding timestamp.
type QualifiedDevice struct {
	*genutil.Metadata
	val     *Device // val is the sample value.
	present bool
}

func (q *QualifiedDevice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Device sample, erroring out if not present.
func (q *QualifiedDevice) Val(t testing.TB) *Device {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Device sample.
func (q *QualifiedDevice) SetVal(v *Device) *QualifiedDevice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedDevice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionDevice is a telemetry Collection whose Await method returns a slice of *Device samples.
type CollectionDevice struct {
	W    *DeviceWatcher
	Data []*QualifiedDevice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionDevice) Await(t testing.TB) []*QualifiedDevice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// DeviceWatcher observes a stream of *Device samples.
type DeviceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedDevice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *DeviceWatcher) Await(t testing.TB) (*QualifiedDevice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow is a *Flow with a corresponding timestamp.
type QualifiedFlow struct {
	*genutil.Metadata
	val     *Flow // val is the sample value.
	present bool
}

func (q *QualifiedFlow) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow sample, erroring out if not present.
func (q *QualifiedFlow) Val(t testing.TB) *Flow {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow sample.
func (q *QualifiedFlow) SetVal(v *Flow) *QualifiedFlow {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow is a telemetry Collection whose Await method returns a slice of *Flow samples.
type CollectionFlow struct {
	W    *FlowWatcher
	Data []*QualifiedFlow
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow) Await(t testing.TB) []*QualifiedFlow {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// FlowWatcher observes a stream of *Flow samples.
type FlowWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *FlowWatcher) Await(t testing.TB) (*QualifiedFlow, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_Counters is a *Flow_Counters with a corresponding timestamp.
type QualifiedFlow_Counters struct {
	*genutil.Metadata
	val     *Flow_Counters // val is the sample value.
	present bool
}

func (q *QualifiedFlow_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_Counters sample, erroring out if not present.
func (q *QualifiedFlow_Counters) Val(t testing.TB) *Flow_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_Counters sample.
func (q *QualifiedFlow_Counters) SetVal(v *Flow_Counters) *QualifiedFlow_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_Counters is a telemetry Collection whose Await method returns a slice of *Flow_Counters samples.
type CollectionFlow_Counters struct {
	W    *Flow_CountersWatcher
	Data []*QualifiedFlow_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_Counters) Await(t testing.TB) []*QualifiedFlow_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_CountersWatcher observes a stream of *Flow_Counters samples.
type Flow_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_CountersWatcher) Await(t testing.TB) (*QualifiedFlow_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_EgressTracking is a *Flow_EgressTracking with a corresponding timestamp.
type QualifiedFlow_EgressTracking struct {
	*genutil.Metadata
	val     *Flow_EgressTracking // val is the sample value.
	present bool
}

func (q *QualifiedFlow_EgressTracking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_EgressTracking sample, erroring out if not present.
func (q *QualifiedFlow_EgressTracking) Val(t testing.TB) *Flow_EgressTracking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_EgressTracking sample.
func (q *QualifiedFlow_EgressTracking) SetVal(v *Flow_EgressTracking) *QualifiedFlow_EgressTracking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_EgressTracking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_EgressTracking is a telemetry Collection whose Await method returns a slice of *Flow_EgressTracking samples.
type CollectionFlow_EgressTracking struct {
	W    *Flow_EgressTrackingWatcher
	Data []*QualifiedFlow_EgressTracking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_EgressTracking) Await(t testing.TB) []*QualifiedFlow_EgressTracking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_EgressTrackingWatcher observes a stream of *Flow_EgressTracking samples.
type Flow_EgressTrackingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_EgressTracking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_EgressTrackingWatcher) Await(t testing.TB) (*QualifiedFlow_EgressTracking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_EgressTracking_Counters is a *Flow_EgressTracking_Counters with a corresponding timestamp.
type QualifiedFlow_EgressTracking_Counters struct {
	*genutil.Metadata
	val     *Flow_EgressTracking_Counters // val is the sample value.
	present bool
}

func (q *QualifiedFlow_EgressTracking_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_EgressTracking_Counters sample, erroring out if not present.
func (q *QualifiedFlow_EgressTracking_Counters) Val(t testing.TB) *Flow_EgressTracking_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_EgressTracking_Counters sample.
func (q *QualifiedFlow_EgressTracking_Counters) SetVal(v *Flow_EgressTracking_Counters) *QualifiedFlow_EgressTracking_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_EgressTracking_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_EgressTracking_Counters is a telemetry Collection whose Await method returns a slice of *Flow_EgressTracking_Counters samples.
type CollectionFlow_EgressTracking_Counters struct {
	W    *Flow_EgressTracking_CountersWatcher
	Data []*QualifiedFlow_EgressTracking_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_EgressTracking_Counters) Await(t testing.TB) []*QualifiedFlow_EgressTracking_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_EgressTracking_CountersWatcher observes a stream of *Flow_EgressTracking_Counters samples.
type Flow_EgressTracking_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_EgressTracking_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_EgressTracking_CountersWatcher) Await(t testing.TB) (*QualifiedFlow_EgressTracking_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_IngressTracking is a *Flow_IngressTracking with a corresponding timestamp.
type QualifiedFlow_IngressTracking struct {
	*genutil.Metadata
	val     *Flow_IngressTracking // val is the sample value.
	present bool
}

func (q *QualifiedFlow_IngressTracking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_IngressTracking sample, erroring out if not present.
func (q *QualifiedFlow_IngressTracking) Val(t testing.TB) *Flow_IngressTracking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_IngressTracking sample.
func (q *QualifiedFlow_IngressTracking) SetVal(v *Flow_IngressTracking) *QualifiedFlow_IngressTracking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_IngressTracking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_IngressTracking is a telemetry Collection whose Await method returns a slice of *Flow_IngressTracking samples.
type CollectionFlow_IngressTracking struct {
	W    *Flow_IngressTrackingWatcher
	Data []*QualifiedFlow_IngressTracking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_IngressTracking) Await(t testing.TB) []*QualifiedFlow_IngressTracking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_IngressTrackingWatcher observes a stream of *Flow_IngressTracking samples.
type Flow_IngressTrackingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_IngressTracking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_IngressTrackingWatcher) Await(t testing.TB) (*QualifiedFlow_IngressTracking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_IngressTracking_Counters is a *Flow_IngressTracking_Counters with a corresponding timestamp.
type QualifiedFlow_IngressTracking_Counters struct {
	*genutil.Metadata
	val     *Flow_IngressTracking_Counters // val is the sample value.
	present bool
}

func (q *QualifiedFlow_IngressTracking_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_IngressTracking_Counters sample, erroring out if not present.
func (q *QualifiedFlow_IngressTracking_Counters) Val(t testing.TB) *Flow_IngressTracking_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_IngressTracking_Counters sample.
func (q *QualifiedFlow_IngressTracking_Counters) SetVal(v *Flow_IngressTracking_Counters) *QualifiedFlow_IngressTracking_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_IngressTracking_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_IngressTracking_Counters is a telemetry Collection whose Await method returns a slice of *Flow_IngressTracking_Counters samples.
type CollectionFlow_IngressTracking_Counters struct {
	W    *Flow_IngressTracking_CountersWatcher
	Data []*QualifiedFlow_IngressTracking_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_IngressTracking_Counters) Await(t testing.TB) []*QualifiedFlow_IngressTracking_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_IngressTracking_CountersWatcher observes a stream of *Flow_IngressTracking_Counters samples.
type Flow_IngressTracking_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_IngressTracking_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_IngressTracking_CountersWatcher) Await(t testing.TB) (*QualifiedFlow_IngressTracking_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_IngressTracking_EgressTracking is a *Flow_IngressTracking_EgressTracking with a corresponding timestamp.
type QualifiedFlow_IngressTracking_EgressTracking struct {
	*genutil.Metadata
	val     *Flow_IngressTracking_EgressTracking // val is the sample value.
	present bool
}

func (q *QualifiedFlow_IngressTracking_EgressTracking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_IngressTracking_EgressTracking sample, erroring out if not present.
func (q *QualifiedFlow_IngressTracking_EgressTracking) Val(t testing.TB) *Flow_IngressTracking_EgressTracking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_IngressTracking_EgressTracking sample.
func (q *QualifiedFlow_IngressTracking_EgressTracking) SetVal(v *Flow_IngressTracking_EgressTracking) *QualifiedFlow_IngressTracking_EgressTracking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_IngressTracking_EgressTracking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_IngressTracking_EgressTracking is a telemetry Collection whose Await method returns a slice of *Flow_IngressTracking_EgressTracking samples.
type CollectionFlow_IngressTracking_EgressTracking struct {
	W    *Flow_IngressTracking_EgressTrackingWatcher
	Data []*QualifiedFlow_IngressTracking_EgressTracking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_IngressTracking_EgressTracking) Await(t testing.TB) []*QualifiedFlow_IngressTracking_EgressTracking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_IngressTracking_EgressTrackingWatcher observes a stream of *Flow_IngressTracking_EgressTracking samples.
type Flow_IngressTracking_EgressTrackingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_IngressTracking_EgressTracking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_IngressTracking_EgressTrackingWatcher) Await(t testing.TB) (*QualifiedFlow_IngressTracking_EgressTracking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_IngressTracking_EgressTracking_Counters is a *Flow_IngressTracking_EgressTracking_Counters with a corresponding timestamp.
type QualifiedFlow_IngressTracking_EgressTracking_Counters struct {
	*genutil.Metadata
	val     *Flow_IngressTracking_EgressTracking_Counters // val is the sample value.
	present bool
}

func (q *QualifiedFlow_IngressTracking_EgressTracking_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_IngressTracking_EgressTracking_Counters sample, erroring out if not present.
func (q *QualifiedFlow_IngressTracking_EgressTracking_Counters) Val(t testing.TB) *Flow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_IngressTracking_EgressTracking_Counters sample.
func (q *QualifiedFlow_IngressTracking_EgressTracking_Counters) SetVal(v *Flow_IngressTracking_EgressTracking_Counters) *QualifiedFlow_IngressTracking_EgressTracking_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_IngressTracking_EgressTracking_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_IngressTracking_EgressTracking_Counters is a telemetry Collection whose Await method returns a slice of *Flow_IngressTracking_EgressTracking_Counters samples.
type CollectionFlow_IngressTracking_EgressTracking_Counters struct {
	W    *Flow_IngressTracking_EgressTracking_CountersWatcher
	Data []*QualifiedFlow_IngressTracking_EgressTracking_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_IngressTracking_EgressTracking_Counters) Await(t testing.TB) []*QualifiedFlow_IngressTracking_EgressTracking_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_IngressTracking_EgressTracking_CountersWatcher observes a stream of *Flow_IngressTracking_EgressTracking_Counters samples.
type Flow_IngressTracking_EgressTracking_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_IngressTracking_EgressTracking_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_IngressTracking_EgressTracking_CountersWatcher) Await(t testing.TB) (*QualifiedFlow_IngressTracking_EgressTracking_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface is a *Interface with a corresponding timestamp.
type QualifiedInterface struct {
	*genutil.Metadata
	val     *Interface // val is the sample value.
	present bool
}

func (q *QualifiedInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface sample, erroring out if not present.
func (q *QualifiedInterface) Val(t testing.TB) *Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface sample.
func (q *QualifiedInterface) SetVal(v *Interface) *QualifiedInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface is a telemetry Collection whose Await method returns a slice of *Interface samples.
type CollectionInterface struct {
	W    *InterfaceWatcher
	Data []*QualifiedInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface) Await(t testing.TB) []*QualifiedInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// InterfaceWatcher observes a stream of *Interface samples.
type InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *InterfaceWatcher) Await(t testing.TB) (*QualifiedInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Aggregation is a *Interface_Aggregation with a corresponding timestamp.
type QualifiedInterface_Aggregation struct {
	*genutil.Metadata
	val     *Interface_Aggregation // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Aggregation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Aggregation sample, erroring out if not present.
func (q *QualifiedInterface_Aggregation) Val(t testing.TB) *Interface_Aggregation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Aggregation sample.
func (q *QualifiedInterface_Aggregation) SetVal(v *Interface_Aggregation) *QualifiedInterface_Aggregation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Aggregation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Aggregation is a telemetry Collection whose Await method returns a slice of *Interface_Aggregation samples.
type CollectionInterface_Aggregation struct {
	W    *Interface_AggregationWatcher
	Data []*QualifiedInterface_Aggregation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Aggregation) Await(t testing.TB) []*QualifiedInterface_Aggregation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_AggregationWatcher observes a stream of *Interface_Aggregation samples.
type Interface_AggregationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Aggregation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_AggregationWatcher) Await(t testing.TB) (*QualifiedInterface_Aggregation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Aggregation_SwitchedVlan is a *Interface_Aggregation_SwitchedVlan with a corresponding timestamp.
type QualifiedInterface_Aggregation_SwitchedVlan struct {
	*genutil.Metadata
	val     *Interface_Aggregation_SwitchedVlan // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Aggregation_SwitchedVlan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Aggregation_SwitchedVlan sample, erroring out if not present.
func (q *QualifiedInterface_Aggregation_SwitchedVlan) Val(t testing.TB) *Interface_Aggregation_SwitchedVlan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Aggregation_SwitchedVlan sample.
func (q *QualifiedInterface_Aggregation_SwitchedVlan) SetVal(v *Interface_Aggregation_SwitchedVlan) *QualifiedInterface_Aggregation_SwitchedVlan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Aggregation_SwitchedVlan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Aggregation_SwitchedVlan is a telemetry Collection whose Await method returns a slice of *Interface_Aggregation_SwitchedVlan samples.
type CollectionInterface_Aggregation_SwitchedVlan struct {
	W    *Interface_Aggregation_SwitchedVlanWatcher
	Data []*QualifiedInterface_Aggregation_SwitchedVlan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Aggregation_SwitchedVlan) Await(t testing.TB) []*QualifiedInterface_Aggregation_SwitchedVlan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Aggregation_SwitchedVlanWatcher observes a stream of *Interface_Aggregation_SwitchedVlan samples.
type Interface_Aggregation_SwitchedVlanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Aggregation_SwitchedVlan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Aggregation_SwitchedVlanWatcher) Await(t testing.TB) (*QualifiedInterface_Aggregation_SwitchedVlan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Counters is a *Interface_Counters with a corresponding timestamp.
type QualifiedInterface_Counters struct {
	*genutil.Metadata
	val     *Interface_Counters // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Counters sample, erroring out if not present.
func (q *QualifiedInterface_Counters) Val(t testing.TB) *Interface_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Counters sample.
func (q *QualifiedInterface_Counters) SetVal(v *Interface_Counters) *QualifiedInterface_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Counters is a telemetry Collection whose Await method returns a slice of *Interface_Counters samples.
type CollectionInterface_Counters struct {
	W    *Interface_CountersWatcher
	Data []*QualifiedInterface_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Counters) Await(t testing.TB) []*QualifiedInterface_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_CountersWatcher observes a stream of *Interface_Counters samples.
type Interface_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_CountersWatcher) Await(t testing.TB) (*QualifiedInterface_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Ethernet is a *Interface_Ethernet with a corresponding timestamp.
type QualifiedInterface_Ethernet struct {
	*genutil.Metadata
	val     *Interface_Ethernet // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ethernet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Ethernet sample, erroring out if not present.
func (q *QualifiedInterface_Ethernet) Val(t testing.TB) *Interface_Ethernet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Ethernet sample.
func (q *QualifiedInterface_Ethernet) SetVal(v *Interface_Ethernet) *QualifiedInterface_Ethernet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ethernet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ethernet is a telemetry Collection whose Await method returns a slice of *Interface_Ethernet samples.
type CollectionInterface_Ethernet struct {
	W    *Interface_EthernetWatcher
	Data []*QualifiedInterface_Ethernet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ethernet) Await(t testing.TB) []*QualifiedInterface_Ethernet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_EthernetWatcher observes a stream of *Interface_Ethernet samples.
type Interface_EthernetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ethernet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_EthernetWatcher) Await(t testing.TB) (*QualifiedInterface_Ethernet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Ethernet_Counters is a *Interface_Ethernet_Counters with a corresponding timestamp.
type QualifiedInterface_Ethernet_Counters struct {
	*genutil.Metadata
	val     *Interface_Ethernet_Counters // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ethernet_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Ethernet_Counters sample, erroring out if not present.
func (q *QualifiedInterface_Ethernet_Counters) Val(t testing.TB) *Interface_Ethernet_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Ethernet_Counters sample.
func (q *QualifiedInterface_Ethernet_Counters) SetVal(v *Interface_Ethernet_Counters) *QualifiedInterface_Ethernet_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ethernet_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ethernet_Counters is a telemetry Collection whose Await method returns a slice of *Interface_Ethernet_Counters samples.
type CollectionInterface_Ethernet_Counters struct {
	W    *Interface_Ethernet_CountersWatcher
	Data []*QualifiedInterface_Ethernet_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ethernet_Counters) Await(t testing.TB) []*QualifiedInterface_Ethernet_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Ethernet_CountersWatcher observes a stream of *Interface_Ethernet_Counters samples.
type Interface_Ethernet_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ethernet_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Ethernet_CountersWatcher) Await(t testing.TB) (*QualifiedInterface_Ethernet_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Ethernet_SwitchedVlan is a *Interface_Ethernet_SwitchedVlan with a corresponding timestamp.
type QualifiedInterface_Ethernet_SwitchedVlan struct {
	*genutil.Metadata
	val     *Interface_Ethernet_SwitchedVlan // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ethernet_SwitchedVlan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Ethernet_SwitchedVlan sample, erroring out if not present.
func (q *QualifiedInterface_Ethernet_SwitchedVlan) Val(t testing.TB) *Interface_Ethernet_SwitchedVlan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Ethernet_SwitchedVlan sample.
func (q *QualifiedInterface_Ethernet_SwitchedVlan) SetVal(v *Interface_Ethernet_SwitchedVlan) *QualifiedInterface_Ethernet_SwitchedVlan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ethernet_SwitchedVlan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ethernet_SwitchedVlan is a telemetry Collection whose Await method returns a slice of *Interface_Ethernet_SwitchedVlan samples.
type CollectionInterface_Ethernet_SwitchedVlan struct {
	W    *Interface_Ethernet_SwitchedVlanWatcher
	Data []*QualifiedInterface_Ethernet_SwitchedVlan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ethernet_SwitchedVlan) Await(t testing.TB) []*QualifiedInterface_Ethernet_SwitchedVlan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Ethernet_SwitchedVlanWatcher observes a stream of *Interface_Ethernet_SwitchedVlan samples.
type Interface_Ethernet_SwitchedVlanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ethernet_SwitchedVlan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Ethernet_SwitchedVlanWatcher) Await(t testing.TB) (*QualifiedInterface_Ethernet_SwitchedVlan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_HoldTime is a *Interface_HoldTime with a corresponding timestamp.
type QualifiedInterface_HoldTime struct {
	*genutil.Metadata
	val     *Interface_HoldTime // val is the sample value.
	present bool
}

func (q *QualifiedInterface_HoldTime) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_HoldTime sample, erroring out if not present.
func (q *QualifiedInterface_HoldTime) Val(t testing.TB) *Interface_HoldTime {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_HoldTime sample.
func (q *QualifiedInterface_HoldTime) SetVal(v *Interface_HoldTime) *QualifiedInterface_HoldTime {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_HoldTime) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_HoldTime is a telemetry Collection whose Await method returns a slice of *Interface_HoldTime samples.
type CollectionInterface_HoldTime struct {
	W    *Interface_HoldTimeWatcher
	Data []*QualifiedInterface_HoldTime
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_HoldTime) Await(t testing.TB) []*QualifiedInterface_HoldTime {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_HoldTimeWatcher observes a stream of *Interface_HoldTime samples.
type Interface_HoldTimeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_HoldTime
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_HoldTimeWatcher) Await(t testing.TB) (*QualifiedInterface_HoldTime, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan is a *Interface_RoutedVlan with a corresponding timestamp.
type QualifiedInterface_RoutedVlan struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan) Val(t testing.TB) *Interface_RoutedVlan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan sample.
func (q *QualifiedInterface_RoutedVlan) SetVal(v *Interface_RoutedVlan) *QualifiedInterface_RoutedVlan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan samples.
type CollectionInterface_RoutedVlan struct {
	W    *Interface_RoutedVlanWatcher
	Data []*QualifiedInterface_RoutedVlan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan) Await(t testing.TB) []*QualifiedInterface_RoutedVlan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlanWatcher observes a stream of *Interface_RoutedVlan samples.
type Interface_RoutedVlanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlanWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4 is a *Interface_RoutedVlan_Ipv4 with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4 struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4 // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4 sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4) Val(t testing.TB) *Interface_RoutedVlan_Ipv4 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4 sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4) SetVal(v *Interface_RoutedVlan_Ipv4) *QualifiedInterface_RoutedVlan_Ipv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4 is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4 samples.
type CollectionInterface_RoutedVlan_Ipv4 struct {
	W    *Interface_RoutedVlan_Ipv4Watcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4Watcher observes a stream of *Interface_RoutedVlan_Ipv4 samples.
type Interface_RoutedVlan_Ipv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4Watcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_Address is a *Interface_RoutedVlan_Ipv4_Address with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_Address struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_Address // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_Address) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_Address sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_Address {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_Address sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address) SetVal(v *Interface_RoutedVlan_Ipv4_Address) *QualifiedInterface_RoutedVlan_Ipv4_Address {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_Address is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_Address samples.
type CollectionInterface_RoutedVlan_Ipv4_Address struct {
	W    *Interface_RoutedVlan_Ipv4_AddressWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_Address
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_Address) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_Address {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_AddressWatcher observes a stream of *Interface_RoutedVlan_Ipv4_Address samples.
type Interface_RoutedVlan_Ipv4_AddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_Address
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_AddressWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_Address, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup is a *Interface_RoutedVlan_Ipv4_Address_VrrpGroup with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_Address_VrrpGroup // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_Address_VrrpGroup sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_Address_VrrpGroup sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) SetVal(v *Interface_RoutedVlan_Ipv4_Address_VrrpGroup) *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_Address_VrrpGroup samples.
type CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup struct {
	W    *Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher observes a stream of *Interface_RoutedVlan_Ipv4_Address_VrrpGroup samples.
type Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_Address_VrrpGroupWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking is a *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) SetVal(v *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking samples.
type CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking struct {
	W    *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher observes a stream of *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking samples.
type Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_Address_VrrpGroup_InterfaceTracking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_Counters is a *Interface_RoutedVlan_Ipv4_Counters with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_Counters struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_Counters // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_Counters sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Counters) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_Counters sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Counters) SetVal(v *Interface_RoutedVlan_Ipv4_Counters) *QualifiedInterface_RoutedVlan_Ipv4_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_Counters is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_Counters samples.
type CollectionInterface_RoutedVlan_Ipv4_Counters struct {
	W    *Interface_RoutedVlan_Ipv4_CountersWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_Counters) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_CountersWatcher observes a stream of *Interface_RoutedVlan_Ipv4_Counters samples.
type Interface_RoutedVlan_Ipv4_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_CountersWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_Neighbor is a *Interface_RoutedVlan_Ipv4_Neighbor with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_Neighbor struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Neighbor) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_Neighbor sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Neighbor) SetVal(v *Interface_RoutedVlan_Ipv4_Neighbor) *QualifiedInterface_RoutedVlan_Ipv4_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_Neighbor samples.
type CollectionInterface_RoutedVlan_Ipv4_Neighbor struct {
	W    *Interface_RoutedVlan_Ipv4_NeighborWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_Neighbor) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_NeighborWatcher observes a stream of *Interface_RoutedVlan_Ipv4_Neighbor samples.
type Interface_RoutedVlan_Ipv4_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_ProxyArp is a *Interface_RoutedVlan_Ipv4_ProxyArp with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_ProxyArp struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_ProxyArp // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_ProxyArp sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_ProxyArp sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) SetVal(v *Interface_RoutedVlan_Ipv4_ProxyArp) *QualifiedInterface_RoutedVlan_Ipv4_ProxyArp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_ProxyArp is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_ProxyArp samples.
type CollectionInterface_RoutedVlan_Ipv4_ProxyArp struct {
	W    *Interface_RoutedVlan_Ipv4_ProxyArpWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_ProxyArp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_ProxyArp) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_ProxyArpWatcher observes a stream of *Interface_RoutedVlan_Ipv4_ProxyArp samples.
type Interface_RoutedVlan_Ipv4_ProxyArpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_ProxyArp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_ProxyArpWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_ProxyArp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_Unnumbered is a *Interface_RoutedVlan_Ipv4_Unnumbered with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_Unnumbered struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_Unnumbered // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_Unnumbered sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_Unnumbered sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) SetVal(v *Interface_RoutedVlan_Ipv4_Unnumbered) *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_Unnumbered is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_Unnumbered samples.
type CollectionInterface_RoutedVlan_Ipv4_Unnumbered struct {
	W    *Interface_RoutedVlan_Ipv4_UnnumberedWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_Unnumbered
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_Unnumbered) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_UnnumberedWatcher observes a stream of *Interface_RoutedVlan_Ipv4_Unnumbered samples.
type Interface_RoutedVlan_Ipv4_UnnumberedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_UnnumberedWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_Unnumbered, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef is a *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) Val(t testing.TB) *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef sample.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) SetVal(v *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef samples.
type CollectionInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef struct {
	W    *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher observes a stream of *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef samples.
type Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6 is a *Interface_RoutedVlan_Ipv6 with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6 struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6 // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6 sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6) Val(t testing.TB) *Interface_RoutedVlan_Ipv6 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6 sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6) SetVal(v *Interface_RoutedVlan_Ipv6) *QualifiedInterface_RoutedVlan_Ipv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6 is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6 samples.
type CollectionInterface_RoutedVlan_Ipv6 struct {
	W    *Interface_RoutedVlan_Ipv6Watcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6Watcher observes a stream of *Interface_RoutedVlan_Ipv6 samples.
type Interface_RoutedVlan_Ipv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6Watcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_Address is a *Interface_RoutedVlan_Ipv6_Address with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_Address struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_Address // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_Address) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_Address sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_Address {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_Address sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address) SetVal(v *Interface_RoutedVlan_Ipv6_Address) *QualifiedInterface_RoutedVlan_Ipv6_Address {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_Address is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_Address samples.
type CollectionInterface_RoutedVlan_Ipv6_Address struct {
	W    *Interface_RoutedVlan_Ipv6_AddressWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_Address
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_Address) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_AddressWatcher observes a stream of *Interface_RoutedVlan_Ipv6_Address samples.
type Interface_RoutedVlan_Ipv6_AddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_Address
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_AddressWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_Address, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup is a *Interface_RoutedVlan_Ipv6_Address_VrrpGroup with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_Address_VrrpGroup // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_Address_VrrpGroup sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_Address_VrrpGroup sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) SetVal(v *Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_Address_VrrpGroup samples.
type CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup struct {
	W    *Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher observes a stream of *Interface_RoutedVlan_Ipv6_Address_VrrpGroup samples.
type Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking is a *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) SetVal(v *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking samples.
type CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking struct {
	W    *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher observes a stream of *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking samples.
type Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_Counters is a *Interface_RoutedVlan_Ipv6_Counters with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_Counters struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_Counters // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_Counters sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Counters) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_Counters sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Counters) SetVal(v *Interface_RoutedVlan_Ipv6_Counters) *QualifiedInterface_RoutedVlan_Ipv6_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_Counters is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_Counters samples.
type CollectionInterface_RoutedVlan_Ipv6_Counters struct {
	W    *Interface_RoutedVlan_Ipv6_CountersWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_Counters) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_CountersWatcher observes a stream of *Interface_RoutedVlan_Ipv6_Counters samples.
type Interface_RoutedVlan_Ipv6_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_CountersWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_Neighbor is a *Interface_RoutedVlan_Ipv6_Neighbor with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_Neighbor struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Neighbor) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_Neighbor sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Neighbor) SetVal(v *Interface_RoutedVlan_Ipv6_Neighbor) *QualifiedInterface_RoutedVlan_Ipv6_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_Neighbor samples.
type CollectionInterface_RoutedVlan_Ipv6_Neighbor struct {
	W    *Interface_RoutedVlan_Ipv6_NeighborWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_Neighbor) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_NeighborWatcher observes a stream of *Interface_RoutedVlan_Ipv6_Neighbor samples.
type Interface_RoutedVlan_Ipv6_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement is a *Interface_RoutedVlan_Ipv6_RouterAdvertisement with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_RouterAdvertisement // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_RouterAdvertisement sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_RouterAdvertisement sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) SetVal(v *Interface_RoutedVlan_Ipv6_RouterAdvertisement) *QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_RouterAdvertisement is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_RouterAdvertisement samples.
type CollectionInterface_RoutedVlan_Ipv6_RouterAdvertisement struct {
	W    *Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_RouterAdvertisement) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher observes a stream of *Interface_RoutedVlan_Ipv6_RouterAdvertisement samples.
type Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_RouterAdvertisementWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_RouterAdvertisement, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_Unnumbered is a *Interface_RoutedVlan_Ipv6_Unnumbered with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_Unnumbered struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_Unnumbered // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_Unnumbered sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_Unnumbered sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) SetVal(v *Interface_RoutedVlan_Ipv6_Unnumbered) *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_Unnumbered is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_Unnumbered samples.
type CollectionInterface_RoutedVlan_Ipv6_Unnumbered struct {
	W    *Interface_RoutedVlan_Ipv6_UnnumberedWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_Unnumbered
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_Unnumbered) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_Unnumbered {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_UnnumberedWatcher observes a stream of *Interface_RoutedVlan_Ipv6_Unnumbered samples.
type Interface_RoutedVlan_Ipv6_UnnumberedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_UnnumberedWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_Unnumbered, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef is a *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef struct {
	*genutil.Metadata
	val     *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) Val(t testing.TB) *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef sample.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) SetVal(v *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef samples.
type CollectionInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef struct {
	W    *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher
	Data []*QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher observes a stream of *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef samples.
type Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface is a *Interface_Subinterface with a corresponding timestamp.
type QualifiedInterface_Subinterface struct {
	*genutil.Metadata
	val     *Interface_Subinterface // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface) Val(t testing.TB) *Interface_Subinterface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface sample.
func (q *QualifiedInterface_Subinterface) SetVal(v *Interface_Subinterface) *QualifiedInterface_Subinterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface samples.
type CollectionInterface_Subinterface struct {
	W    *Interface_SubinterfaceWatcher
	Data []*QualifiedInterface_Subinterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface) Await(t testing.TB) []*QualifiedInterface_Subinterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_SubinterfaceWatcher observes a stream of *Interface_Subinterface samples.
type Interface_SubinterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_SubinterfaceWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Counters is a *Interface_Subinterface_Counters with a corresponding timestamp.
type QualifiedInterface_Subinterface_Counters struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Counters // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Counters sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Counters) Val(t testing.TB) *Interface_Subinterface_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Counters sample.
func (q *QualifiedInterface_Subinterface_Counters) SetVal(v *Interface_Subinterface_Counters) *QualifiedInterface_Subinterface_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Counters is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Counters samples.
type CollectionInterface_Subinterface_Counters struct {
	W    *Interface_Subinterface_CountersWatcher
	Data []*QualifiedInterface_Subinterface_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Counters) Await(t testing.TB) []*QualifiedInterface_Subinterface_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_CountersWatcher observes a stream of *Interface_Subinterface_Counters samples.
type Interface_Subinterface_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_CountersWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4 is a *Interface_Subinterface_Ipv4 with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4 struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4 // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4 sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4) Val(t testing.TB) *Interface_Subinterface_Ipv4 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4 sample.
func (q *QualifiedInterface_Subinterface_Ipv4) SetVal(v *Interface_Subinterface_Ipv4) *QualifiedInterface_Subinterface_Ipv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4 is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4 samples.
type CollectionInterface_Subinterface_Ipv4 struct {
	W    *Interface_Subinterface_Ipv4Watcher
	Data []*QualifiedInterface_Subinterface_Ipv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4Watcher observes a stream of *Interface_Subinterface_Ipv4 samples.
type Interface_Subinterface_Ipv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4Watcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_Address is a *Interface_Subinterface_Ipv4_Address with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_Address struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_Address // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_Address) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_Address sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_Address) Val(t testing.TB) *Interface_Subinterface_Ipv4_Address {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_Address sample.
func (q *QualifiedInterface_Subinterface_Ipv4_Address) SetVal(v *Interface_Subinterface_Ipv4_Address) *QualifiedInterface_Subinterface_Ipv4_Address {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_Address) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_Address is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_Address samples.
type CollectionInterface_Subinterface_Ipv4_Address struct {
	W    *Interface_Subinterface_Ipv4_AddressWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_Address
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_Address) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_Address {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_AddressWatcher observes a stream of *Interface_Subinterface_Ipv4_Address samples.
type Interface_Subinterface_Ipv4_AddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_Address
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_AddressWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_Address, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup is a *Interface_Subinterface_Ipv4_Address_VrrpGroup with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_Address_VrrpGroup // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_Address_VrrpGroup sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup) Val(t testing.TB) *Interface_Subinterface_Ipv4_Address_VrrpGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_Address_VrrpGroup sample.
func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup) SetVal(v *Interface_Subinterface_Ipv4_Address_VrrpGroup) *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_Address_VrrpGroup is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_Address_VrrpGroup samples.
type CollectionInterface_Subinterface_Ipv4_Address_VrrpGroup struct {
	W    *Interface_Subinterface_Ipv4_Address_VrrpGroupWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_Address_VrrpGroup) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_Address_VrrpGroupWatcher observes a stream of *Interface_Subinterface_Ipv4_Address_VrrpGroup samples.
type Interface_Subinterface_Ipv4_Address_VrrpGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_Address_VrrpGroupWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking is a *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking) Val(t testing.TB) *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking sample.
func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking) SetVal(v *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking) *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking samples.
type CollectionInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking struct {
	W    *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher observes a stream of *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking samples.
type Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTrackingWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_Address_VrrpGroup_InterfaceTracking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_Counters is a *Interface_Subinterface_Ipv4_Counters with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_Counters struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_Counters // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_Counters sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_Counters) Val(t testing.TB) *Interface_Subinterface_Ipv4_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_Counters sample.
func (q *QualifiedInterface_Subinterface_Ipv4_Counters) SetVal(v *Interface_Subinterface_Ipv4_Counters) *QualifiedInterface_Subinterface_Ipv4_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_Counters is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_Counters samples.
type CollectionInterface_Subinterface_Ipv4_Counters struct {
	W    *Interface_Subinterface_Ipv4_CountersWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_Counters) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_CountersWatcher observes a stream of *Interface_Subinterface_Ipv4_Counters samples.
type Interface_Subinterface_Ipv4_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_CountersWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_Neighbor is a *Interface_Subinterface_Ipv4_Neighbor with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_Neighbor struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_Neighbor) Val(t testing.TB) *Interface_Subinterface_Ipv4_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_Neighbor sample.
func (q *QualifiedInterface_Subinterface_Ipv4_Neighbor) SetVal(v *Interface_Subinterface_Ipv4_Neighbor) *QualifiedInterface_Subinterface_Ipv4_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_Neighbor samples.
type CollectionInterface_Subinterface_Ipv4_Neighbor struct {
	W    *Interface_Subinterface_Ipv4_NeighborWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_Neighbor) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_NeighborWatcher observes a stream of *Interface_Subinterface_Ipv4_Neighbor samples.
type Interface_Subinterface_Ipv4_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_ProxyArp is a *Interface_Subinterface_Ipv4_ProxyArp with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_ProxyArp struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_ProxyArp // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_ProxyArp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_ProxyArp sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_ProxyArp) Val(t testing.TB) *Interface_Subinterface_Ipv4_ProxyArp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_ProxyArp sample.
func (q *QualifiedInterface_Subinterface_Ipv4_ProxyArp) SetVal(v *Interface_Subinterface_Ipv4_ProxyArp) *QualifiedInterface_Subinterface_Ipv4_ProxyArp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_ProxyArp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_ProxyArp is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_ProxyArp samples.
type CollectionInterface_Subinterface_Ipv4_ProxyArp struct {
	W    *Interface_Subinterface_Ipv4_ProxyArpWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_ProxyArp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_ProxyArp) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_ProxyArp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_ProxyArpWatcher observes a stream of *Interface_Subinterface_Ipv4_ProxyArp samples.
type Interface_Subinterface_Ipv4_ProxyArpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_ProxyArp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_ProxyArpWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_ProxyArp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_Unnumbered is a *Interface_Subinterface_Ipv4_Unnumbered with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_Unnumbered struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_Unnumbered // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_Unnumbered sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered) Val(t testing.TB) *Interface_Subinterface_Ipv4_Unnumbered {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_Unnumbered sample.
func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered) SetVal(v *Interface_Subinterface_Ipv4_Unnumbered) *QualifiedInterface_Subinterface_Ipv4_Unnumbered {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_Unnumbered is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_Unnumbered samples.
type CollectionInterface_Subinterface_Ipv4_Unnumbered struct {
	W    *Interface_Subinterface_Ipv4_UnnumberedWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_Unnumbered
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_Unnumbered) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_Unnumbered {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_UnnumberedWatcher observes a stream of *Interface_Subinterface_Ipv4_Unnumbered samples.
type Interface_Subinterface_Ipv4_UnnumberedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_Unnumbered
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_UnnumberedWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_Unnumbered, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef is a *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) Val(t testing.TB) *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef sample.
func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) SetVal(v *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef) *QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef samples.
type CollectionInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef struct {
	W    *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher
	Data []*QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher observes a stream of *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef samples.
type Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6 is a *Interface_Subinterface_Ipv6 with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6 struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6 // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6 sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6) Val(t testing.TB) *Interface_Subinterface_Ipv6 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6 sample.
func (q *QualifiedInterface_Subinterface_Ipv6) SetVal(v *Interface_Subinterface_Ipv6) *QualifiedInterface_Subinterface_Ipv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6 is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6 samples.
type CollectionInterface_Subinterface_Ipv6 struct {
	W    *Interface_Subinterface_Ipv6Watcher
	Data []*QualifiedInterface_Subinterface_Ipv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6Watcher observes a stream of *Interface_Subinterface_Ipv6 samples.
type Interface_Subinterface_Ipv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6Watcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Address is a *Interface_Subinterface_Ipv6_Address with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Address struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Address // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Address) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Address sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Address) Val(t testing.TB) *Interface_Subinterface_Ipv6_Address {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Address sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Address) SetVal(v *Interface_Subinterface_Ipv6_Address) *QualifiedInterface_Subinterface_Ipv6_Address {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Address) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Address is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Address samples.
type CollectionInterface_Subinterface_Ipv6_Address struct {
	W    *Interface_Subinterface_Ipv6_AddressWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Address
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Address) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Address {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_AddressWatcher observes a stream of *Interface_Subinterface_Ipv6_Address samples.
type Interface_Subinterface_Ipv6_AddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Address
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_AddressWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Address, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup is a *Interface_Subinterface_Ipv6_Address_VrrpGroup with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Address_VrrpGroup // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Address_VrrpGroup sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) Val(t testing.TB) *Interface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Address_VrrpGroup sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) SetVal(v *Interface_Subinterface_Ipv6_Address_VrrpGroup) *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Address_VrrpGroup samples.
type CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup struct {
	W    *Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher observes a stream of *Interface_Subinterface_Ipv6_Address_VrrpGroup samples.
type Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_Address_VrrpGroupWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking is a *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) Val(t testing.TB) *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) SetVal(v *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking samples.
type CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking struct {
	W    *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher observes a stream of *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking samples.
type Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Autoconf is a *Interface_Subinterface_Ipv6_Autoconf with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Autoconf struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Autoconf // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Autoconf) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Autoconf sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Autoconf) Val(t testing.TB) *Interface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Autoconf sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Autoconf) SetVal(v *Interface_Subinterface_Ipv6_Autoconf) *QualifiedInterface_Subinterface_Ipv6_Autoconf {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Autoconf) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Autoconf is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Autoconf samples.
type CollectionInterface_Subinterface_Ipv6_Autoconf struct {
	W    *Interface_Subinterface_Ipv6_AutoconfWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Autoconf
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Autoconf) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_AutoconfWatcher observes a stream of *Interface_Subinterface_Ipv6_Autoconf samples.
type Interface_Subinterface_Ipv6_AutoconfWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Autoconf
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_AutoconfWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Autoconf, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Counters is a *Interface_Subinterface_Ipv6_Counters with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Counters struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Counters // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Counters sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Counters) Val(t testing.TB) *Interface_Subinterface_Ipv6_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Counters sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Counters) SetVal(v *Interface_Subinterface_Ipv6_Counters) *QualifiedInterface_Subinterface_Ipv6_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Counters is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Counters samples.
type CollectionInterface_Subinterface_Ipv6_Counters struct {
	W    *Interface_Subinterface_Ipv6_CountersWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Counters) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_CountersWatcher observes a stream of *Interface_Subinterface_Ipv6_Counters samples.
type Interface_Subinterface_Ipv6_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_CountersWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Neighbor is a *Interface_Subinterface_Ipv6_Neighbor with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Neighbor struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Neighbor) Val(t testing.TB) *Interface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Neighbor sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Neighbor) SetVal(v *Interface_Subinterface_Ipv6_Neighbor) *QualifiedInterface_Subinterface_Ipv6_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Neighbor samples.
type CollectionInterface_Subinterface_Ipv6_Neighbor struct {
	W    *Interface_Subinterface_Ipv6_NeighborWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Neighbor) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_NeighborWatcher observes a stream of *Interface_Subinterface_Ipv6_Neighbor samples.
type Interface_Subinterface_Ipv6_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement is a *Interface_Subinterface_Ipv6_RouterAdvertisement with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_RouterAdvertisement // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_RouterAdvertisement sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) Val(t testing.TB) *Interface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_RouterAdvertisement sample.
func (q *QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) SetVal(v *Interface_Subinterface_Ipv6_RouterAdvertisement) *QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_RouterAdvertisement is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_RouterAdvertisement samples.
type CollectionInterface_Subinterface_Ipv6_RouterAdvertisement struct {
	W    *Interface_Subinterface_Ipv6_RouterAdvertisementWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_RouterAdvertisement) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_RouterAdvertisementWatcher observes a stream of *Interface_Subinterface_Ipv6_RouterAdvertisement samples.
type Interface_Subinterface_Ipv6_RouterAdvertisementWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_RouterAdvertisementWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Unnumbered is a *Interface_Subinterface_Ipv6_Unnumbered with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Unnumbered struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Unnumbered // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Unnumbered sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered) Val(t testing.TB) *Interface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Unnumbered sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered) SetVal(v *Interface_Subinterface_Ipv6_Unnumbered) *QualifiedInterface_Subinterface_Ipv6_Unnumbered {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Unnumbered is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Unnumbered samples.
type CollectionInterface_Subinterface_Ipv6_Unnumbered struct {
	W    *Interface_Subinterface_Ipv6_UnnumberedWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Unnumbered
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Unnumbered) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_UnnumberedWatcher observes a stream of *Interface_Subinterface_Ipv6_Unnumbered samples.
type Interface_Subinterface_Ipv6_UnnumberedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Unnumbered
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_UnnumberedWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Unnumbered, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef is a *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef with a corresponding timestamp.
type QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) Val(t testing.TB) *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef sample.
func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) SetVal(v *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef samples.
type CollectionInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef struct {
	W    *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher
	Data []*QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef) Await(t testing.TB) []*QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher observes a stream of *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef samples.
type Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan is a *Interface_Subinterface_Vlan with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan) Val(t testing.TB) *Interface_Subinterface_Vlan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan sample.
func (q *QualifiedInterface_Subinterface_Vlan) SetVal(v *Interface_Subinterface_Vlan) *QualifiedInterface_Subinterface_Vlan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan samples.
type CollectionInterface_Subinterface_Vlan struct {
	W    *Interface_Subinterface_VlanWatcher
	Data []*QualifiedInterface_Subinterface_Vlan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_VlanWatcher observes a stream of *Interface_Subinterface_Vlan samples.
type Interface_Subinterface_VlanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_VlanWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_EgressMapping is a *Interface_Subinterface_Vlan_EgressMapping with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_EgressMapping struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_EgressMapping // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_EgressMapping) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_EgressMapping sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_EgressMapping) Val(t testing.TB) *Interface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_EgressMapping sample.
func (q *QualifiedInterface_Subinterface_Vlan_EgressMapping) SetVal(v *Interface_Subinterface_Vlan_EgressMapping) *QualifiedInterface_Subinterface_Vlan_EgressMapping {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_EgressMapping) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_EgressMapping is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_EgressMapping samples.
type CollectionInterface_Subinterface_Vlan_EgressMapping struct {
	W    *Interface_Subinterface_Vlan_EgressMappingWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_EgressMapping
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_EgressMapping) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_EgressMappingWatcher observes a stream of *Interface_Subinterface_Vlan_EgressMapping samples.
type Interface_Subinterface_Vlan_EgressMappingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_EgressMapping
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_EgressMappingWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_EgressMapping, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_IngressMapping is a *Interface_Subinterface_Vlan_IngressMapping with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_IngressMapping struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_IngressMapping // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_IngressMapping) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_IngressMapping sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_IngressMapping) Val(t testing.TB) *Interface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_IngressMapping sample.
func (q *QualifiedInterface_Subinterface_Vlan_IngressMapping) SetVal(v *Interface_Subinterface_Vlan_IngressMapping) *QualifiedInterface_Subinterface_Vlan_IngressMapping {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_IngressMapping) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_IngressMapping is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_IngressMapping samples.
type CollectionInterface_Subinterface_Vlan_IngressMapping struct {
	W    *Interface_Subinterface_Vlan_IngressMappingWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_IngressMapping
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_IngressMapping) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_IngressMapping {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_IngressMappingWatcher observes a stream of *Interface_Subinterface_Vlan_IngressMapping samples.
type Interface_Subinterface_Vlan_IngressMappingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_IngressMapping
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_IngressMappingWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_IngressMapping, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match is a *Interface_Subinterface_Vlan_Match with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match) Val(t testing.TB) *Interface_Subinterface_Vlan_Match {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match) SetVal(v *Interface_Subinterface_Vlan_Match) *QualifiedInterface_Subinterface_Vlan_Match {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match samples.
type CollectionInterface_Subinterface_Vlan_Match struct {
	W    *Interface_Subinterface_Vlan_MatchWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_MatchWatcher observes a stream of *Interface_Subinterface_Vlan_Match samples.
type Interface_Subinterface_Vlan_MatchWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_MatchWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged is a *Interface_Subinterface_Vlan_Match_DoubleTagged with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_DoubleTagged // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_DoubleTagged sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_DoubleTagged sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) SetVal(v *Interface_Subinterface_Vlan_Match_DoubleTagged) *QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_DoubleTagged is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_DoubleTagged samples.
type CollectionInterface_Subinterface_Vlan_Match_DoubleTagged struct {
	W    *Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_DoubleTagged) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher observes a stream of *Interface_Subinterface_Vlan_Match_DoubleTagged samples.
type Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_DoubleTaggedWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_DoubleTagged, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList is a *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) SetVal(v *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList) *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList samples.
type CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList struct {
	W    *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher observes a stream of *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerList samples.
type Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerListWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerList, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange is a *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) SetVal(v *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange samples.
type CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange struct {
	W    *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher observes a stream of *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange samples.
type Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRangeWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerOuterRange, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange is a *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) SetVal(v *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange samples.
type CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange struct {
	W    *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher observes a stream of *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRange samples.
type Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_DoubleTaggedInnerRangeWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedInnerRange, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList is a *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) SetVal(v *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList) *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList samples.
type CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList struct {
	W    *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher observes a stream of *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterList samples.
type Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterListWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterList, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange is a *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) SetVal(v *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange samples.
type CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange struct {
	W    *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher observes a stream of *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRange samples.
type Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_DoubleTaggedOuterRangeWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_DoubleTaggedOuterRange, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_SingleTagged is a *Interface_Subinterface_Vlan_Match_SingleTagged with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_SingleTagged struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_SingleTagged // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_SingleTagged sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_SingleTagged sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) SetVal(v *Interface_Subinterface_Vlan_Match_SingleTagged) *QualifiedInterface_Subinterface_Vlan_Match_SingleTagged {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTagged) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_SingleTagged is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_SingleTagged samples.
type CollectionInterface_Subinterface_Vlan_Match_SingleTagged struct {
	W    *Interface_Subinterface_Vlan_Match_SingleTaggedWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_SingleTagged
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_SingleTagged) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_SingleTagged {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_SingleTaggedWatcher observes a stream of *Interface_Subinterface_Vlan_Match_SingleTagged samples.
type Interface_Subinterface_Vlan_Match_SingleTaggedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_SingleTagged
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_SingleTaggedWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_SingleTagged, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList is a *Interface_Subinterface_Vlan_Match_SingleTaggedList with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_SingleTaggedList // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_SingleTaggedList sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_SingleTaggedList sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) SetVal(v *Interface_Subinterface_Vlan_Match_SingleTaggedList) *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_SingleTaggedList is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_SingleTaggedList samples.
type CollectionInterface_Subinterface_Vlan_Match_SingleTaggedList struct {
	W    *Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_SingleTaggedList) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher observes a stream of *Interface_Subinterface_Vlan_Match_SingleTaggedList samples.
type Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_SingleTaggedListWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedList, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange is a *Interface_Subinterface_Vlan_Match_SingleTaggedRange with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange struct {
	*genutil.Metadata
	val     *Interface_Subinterface_Vlan_Match_SingleTaggedRange // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Subinterface_Vlan_Match_SingleTaggedRange sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) Val(t testing.TB) *Interface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Subinterface_Vlan_Match_SingleTaggedRange sample.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) SetVal(v *Interface_Subinterface_Vlan_Match_SingleTaggedRange) *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_Match_SingleTaggedRange is a telemetry Collection whose Await method returns a slice of *Interface_Subinterface_Vlan_Match_SingleTaggedRange samples.
type CollectionInterface_Subinterface_Vlan_Match_SingleTaggedRange struct {
	W    *Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_Match_SingleTaggedRange) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher observes a stream of *Interface_Subinterface_Vlan_Match_SingleTaggedRange samples.
type Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_Match_SingleTaggedRangeWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_Match_SingleTaggedRange, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedKeychain is a *Keychain with a corresponding timestamp.
type QualifiedKeychain struct {
	*genutil.Metadata
	val     *Keychain // val is the sample value.
	present bool
}

func (q *QualifiedKeychain) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Keychain sample, erroring out if not present.
func (q *QualifiedKeychain) Val(t testing.TB) *Keychain {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Keychain sample.
func (q *QualifiedKeychain) SetVal(v *Keychain) *QualifiedKeychain {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedKeychain) IsPresent() bool {
	return q != nil && q.present
}

// CollectionKeychain is a telemetry Collection whose Await method returns a slice of *Keychain samples.
type CollectionKeychain struct {
	W    *KeychainWatcher
	Data []*QualifiedKeychain
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionKeychain) Await(t testing.TB) []*QualifiedKeychain {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// KeychainWatcher observes a stream of *Keychain samples.
type KeychainWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedKeychain
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *KeychainWatcher) Await(t testing.TB) (*QualifiedKeychain, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedKeychain_Key is a *Keychain_Key with a corresponding timestamp.
type QualifiedKeychain_Key struct {
	*genutil.Metadata
	val     *Keychain_Key // val is the sample value.
	present bool
}

func (q *QualifiedKeychain_Key) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Keychain_Key sample, erroring out if not present.
func (q *QualifiedKeychain_Key) Val(t testing.TB) *Keychain_Key {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Keychain_Key sample.
func (q *QualifiedKeychain_Key) SetVal(v *Keychain_Key) *QualifiedKeychain_Key {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedKeychain_Key) IsPresent() bool {
	return q != nil && q.present
}

// CollectionKeychain_Key is a telemetry Collection whose Await method returns a slice of *Keychain_Key samples.
type CollectionKeychain_Key struct {
	W    *Keychain_KeyWatcher
	Data []*QualifiedKeychain_Key
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionKeychain_Key) Await(t testing.TB) []*QualifiedKeychain_Key {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Keychain_KeyWatcher observes a stream of *Keychain_Key samples.
type Keychain_KeyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedKeychain_Key
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Keychain_KeyWatcher) Await(t testing.TB) (*QualifiedKeychain_Key, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedKeychain_Key_ReceiveLifetime is a *Keychain_Key_ReceiveLifetime with a corresponding timestamp.
type QualifiedKeychain_Key_ReceiveLifetime struct {
	*genutil.Metadata
	val     *Keychain_Key_ReceiveLifetime // val is the sample value.
	present bool
}

func (q *QualifiedKeychain_Key_ReceiveLifetime) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Keychain_Key_ReceiveLifetime sample, erroring out if not present.
func (q *QualifiedKeychain_Key_ReceiveLifetime) Val(t testing.TB) *Keychain_Key_ReceiveLifetime {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Keychain_Key_ReceiveLifetime sample.
func (q *QualifiedKeychain_Key_ReceiveLifetime) SetVal(v *Keychain_Key_ReceiveLifetime) *QualifiedKeychain_Key_ReceiveLifetime {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedKeychain_Key_ReceiveLifetime) IsPresent() bool {
	return q != nil && q.present
}

// CollectionKeychain_Key_ReceiveLifetime is a telemetry Collection whose Await method returns a slice of *Keychain_Key_ReceiveLifetime samples.
type CollectionKeychain_Key_ReceiveLifetime struct {
	W    *Keychain_Key_ReceiveLifetimeWatcher
	Data []*QualifiedKeychain_Key_ReceiveLifetime
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionKeychain_Key_ReceiveLifetime) Await(t testing.TB) []*QualifiedKeychain_Key_ReceiveLifetime {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Keychain_Key_ReceiveLifetimeWatcher observes a stream of *Keychain_Key_ReceiveLifetime samples.
type Keychain_Key_ReceiveLifetimeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedKeychain_Key_ReceiveLifetime
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Keychain_Key_ReceiveLifetimeWatcher) Await(t testing.TB) (*QualifiedKeychain_Key_ReceiveLifetime, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedKeychain_Key_SendLifetime is a *Keychain_Key_SendLifetime with a corresponding timestamp.
type QualifiedKeychain_Key_SendLifetime struct {
	*genutil.Metadata
	val     *Keychain_Key_SendLifetime // val is the sample value.
	present bool
}

func (q *QualifiedKeychain_Key_SendLifetime) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Keychain_Key_SendLifetime sample, erroring out if not present.
func (q *QualifiedKeychain_Key_SendLifetime) Val(t testing.TB) *Keychain_Key_SendLifetime {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Keychain_Key_SendLifetime sample.
func (q *QualifiedKeychain_Key_SendLifetime) SetVal(v *Keychain_Key_SendLifetime) *QualifiedKeychain_Key_SendLifetime {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedKeychain_Key_SendLifetime) IsPresent() bool {
	return q != nil && q.present
}

// CollectionKeychain_Key_SendLifetime is a telemetry Collection whose Await method returns a slice of *Keychain_Key_SendLifetime samples.
type CollectionKeychain_Key_SendLifetime struct {
	W    *Keychain_Key_SendLifetimeWatcher
	Data []*QualifiedKeychain_Key_SendLifetime
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionKeychain_Key_SendLifetime) Await(t testing.TB) []*QualifiedKeychain_Key_SendLifetime {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Keychain_Key_SendLifetimeWatcher observes a stream of *Keychain_Key_SendLifetime samples.
type Keychain_Key_SendLifetimeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedKeychain_Key_SendLifetime
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Keychain_Key_SendLifetimeWatcher) Await(t testing.TB) (*QualifiedKeychain_Key_SendLifetime, bool) {
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

// QualifiedLacp_Interface is a *Lacp_Interface with a corresponding timestamp.
type QualifiedLacp_Interface struct {
	*genutil.Metadata
	val     *Lacp_Interface // val is the sample value.
	present bool
}

func (q *QualifiedLacp_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lacp_Interface sample, erroring out if not present.
func (q *QualifiedLacp_Interface) Val(t testing.TB) *Lacp_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lacp_Interface sample.
func (q *QualifiedLacp_Interface) SetVal(v *Lacp_Interface) *QualifiedLacp_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLacp_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLacp_Interface is a telemetry Collection whose Await method returns a slice of *Lacp_Interface samples.
type CollectionLacp_Interface struct {
	W    *Lacp_InterfaceWatcher
	Data []*QualifiedLacp_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLacp_Interface) Await(t testing.TB) []*QualifiedLacp_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lacp_InterfaceWatcher observes a stream of *Lacp_Interface samples.
type Lacp_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLacp_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lacp_InterfaceWatcher) Await(t testing.TB) (*QualifiedLacp_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLacp_Interface_Member is a *Lacp_Interface_Member with a corresponding timestamp.
type QualifiedLacp_Interface_Member struct {
	*genutil.Metadata
	val     *Lacp_Interface_Member // val is the sample value.
	present bool
}

func (q *QualifiedLacp_Interface_Member) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lacp_Interface_Member sample, erroring out if not present.
func (q *QualifiedLacp_Interface_Member) Val(t testing.TB) *Lacp_Interface_Member {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lacp_Interface_Member sample.
func (q *QualifiedLacp_Interface_Member) SetVal(v *Lacp_Interface_Member) *QualifiedLacp_Interface_Member {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLacp_Interface_Member) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLacp_Interface_Member is a telemetry Collection whose Await method returns a slice of *Lacp_Interface_Member samples.
type CollectionLacp_Interface_Member struct {
	W    *Lacp_Interface_MemberWatcher
	Data []*QualifiedLacp_Interface_Member
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLacp_Interface_Member) Await(t testing.TB) []*QualifiedLacp_Interface_Member {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lacp_Interface_MemberWatcher observes a stream of *Lacp_Interface_Member samples.
type Lacp_Interface_MemberWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLacp_Interface_Member
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lacp_Interface_MemberWatcher) Await(t testing.TB) (*QualifiedLacp_Interface_Member, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLacp_Interface_Member_Counters is a *Lacp_Interface_Member_Counters with a corresponding timestamp.
type QualifiedLacp_Interface_Member_Counters struct {
	*genutil.Metadata
	val     *Lacp_Interface_Member_Counters // val is the sample value.
	present bool
}

func (q *QualifiedLacp_Interface_Member_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lacp_Interface_Member_Counters sample, erroring out if not present.
func (q *QualifiedLacp_Interface_Member_Counters) Val(t testing.TB) *Lacp_Interface_Member_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lacp_Interface_Member_Counters sample.
func (q *QualifiedLacp_Interface_Member_Counters) SetVal(v *Lacp_Interface_Member_Counters) *QualifiedLacp_Interface_Member_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLacp_Interface_Member_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLacp_Interface_Member_Counters is a telemetry Collection whose Await method returns a slice of *Lacp_Interface_Member_Counters samples.
type CollectionLacp_Interface_Member_Counters struct {
	W    *Lacp_Interface_Member_CountersWatcher
	Data []*QualifiedLacp_Interface_Member_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLacp_Interface_Member_Counters) Await(t testing.TB) []*QualifiedLacp_Interface_Member_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lacp_Interface_Member_CountersWatcher observes a stream of *Lacp_Interface_Member_Counters samples.
type Lacp_Interface_Member_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLacp_Interface_Member_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lacp_Interface_Member_CountersWatcher) Await(t testing.TB) (*QualifiedLacp_Interface_Member_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldp is a *Lldp with a corresponding timestamp.
type QualifiedLldp struct {
	*genutil.Metadata
	val     *Lldp // val is the sample value.
	present bool
}

func (q *QualifiedLldp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lldp sample, erroring out if not present.
func (q *QualifiedLldp) Val(t testing.TB) *Lldp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lldp sample.
func (q *QualifiedLldp) SetVal(v *Lldp) *QualifiedLldp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldp is a telemetry Collection whose Await method returns a slice of *Lldp samples.
type CollectionLldp struct {
	W    *LldpWatcher
	Data []*QualifiedLldp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldp) Await(t testing.TB) []*QualifiedLldp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LldpWatcher observes a stream of *Lldp samples.
type LldpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LldpWatcher) Await(t testing.TB) (*QualifiedLldp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldp_Counters is a *Lldp_Counters with a corresponding timestamp.
type QualifiedLldp_Counters struct {
	*genutil.Metadata
	val     *Lldp_Counters // val is the sample value.
	present bool
}

func (q *QualifiedLldp_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lldp_Counters sample, erroring out if not present.
func (q *QualifiedLldp_Counters) Val(t testing.TB) *Lldp_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lldp_Counters sample.
func (q *QualifiedLldp_Counters) SetVal(v *Lldp_Counters) *QualifiedLldp_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldp_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldp_Counters is a telemetry Collection whose Await method returns a slice of *Lldp_Counters samples.
type CollectionLldp_Counters struct {
	W    *Lldp_CountersWatcher
	Data []*QualifiedLldp_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldp_Counters) Await(t testing.TB) []*QualifiedLldp_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lldp_CountersWatcher observes a stream of *Lldp_Counters samples.
type Lldp_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldp_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lldp_CountersWatcher) Await(t testing.TB) (*QualifiedLldp_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldp_Interface is a *Lldp_Interface with a corresponding timestamp.
type QualifiedLldp_Interface struct {
	*genutil.Metadata
	val     *Lldp_Interface // val is the sample value.
	present bool
}

func (q *QualifiedLldp_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lldp_Interface sample, erroring out if not present.
func (q *QualifiedLldp_Interface) Val(t testing.TB) *Lldp_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lldp_Interface sample.
func (q *QualifiedLldp_Interface) SetVal(v *Lldp_Interface) *QualifiedLldp_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldp_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldp_Interface is a telemetry Collection whose Await method returns a slice of *Lldp_Interface samples.
type CollectionLldp_Interface struct {
	W    *Lldp_InterfaceWatcher
	Data []*QualifiedLldp_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldp_Interface) Await(t testing.TB) []*QualifiedLldp_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lldp_InterfaceWatcher observes a stream of *Lldp_Interface samples.
type Lldp_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldp_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lldp_InterfaceWatcher) Await(t testing.TB) (*QualifiedLldp_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldp_Interface_Counters is a *Lldp_Interface_Counters with a corresponding timestamp.
type QualifiedLldp_Interface_Counters struct {
	*genutil.Metadata
	val     *Lldp_Interface_Counters // val is the sample value.
	present bool
}

func (q *QualifiedLldp_Interface_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lldp_Interface_Counters sample, erroring out if not present.
func (q *QualifiedLldp_Interface_Counters) Val(t testing.TB) *Lldp_Interface_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lldp_Interface_Counters sample.
func (q *QualifiedLldp_Interface_Counters) SetVal(v *Lldp_Interface_Counters) *QualifiedLldp_Interface_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldp_Interface_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldp_Interface_Counters is a telemetry Collection whose Await method returns a slice of *Lldp_Interface_Counters samples.
type CollectionLldp_Interface_Counters struct {
	W    *Lldp_Interface_CountersWatcher
	Data []*QualifiedLldp_Interface_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldp_Interface_Counters) Await(t testing.TB) []*QualifiedLldp_Interface_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lldp_Interface_CountersWatcher observes a stream of *Lldp_Interface_Counters samples.
type Lldp_Interface_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldp_Interface_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lldp_Interface_CountersWatcher) Await(t testing.TB) (*QualifiedLldp_Interface_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldp_Interface_Neighbor is a *Lldp_Interface_Neighbor with a corresponding timestamp.
type QualifiedLldp_Interface_Neighbor struct {
	*genutil.Metadata
	val     *Lldp_Interface_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedLldp_Interface_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lldp_Interface_Neighbor sample, erroring out if not present.
func (q *QualifiedLldp_Interface_Neighbor) Val(t testing.TB) *Lldp_Interface_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lldp_Interface_Neighbor sample.
func (q *QualifiedLldp_Interface_Neighbor) SetVal(v *Lldp_Interface_Neighbor) *QualifiedLldp_Interface_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldp_Interface_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldp_Interface_Neighbor is a telemetry Collection whose Await method returns a slice of *Lldp_Interface_Neighbor samples.
type CollectionLldp_Interface_Neighbor struct {
	W    *Lldp_Interface_NeighborWatcher
	Data []*QualifiedLldp_Interface_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldp_Interface_Neighbor) Await(t testing.TB) []*QualifiedLldp_Interface_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lldp_Interface_NeighborWatcher observes a stream of *Lldp_Interface_Neighbor samples.
type Lldp_Interface_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldp_Interface_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lldp_Interface_NeighborWatcher) Await(t testing.TB) (*QualifiedLldp_Interface_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldp_Interface_Neighbor_Capability is a *Lldp_Interface_Neighbor_Capability with a corresponding timestamp.
type QualifiedLldp_Interface_Neighbor_Capability struct {
	*genutil.Metadata
	val     *Lldp_Interface_Neighbor_Capability // val is the sample value.
	present bool
}

func (q *QualifiedLldp_Interface_Neighbor_Capability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lldp_Interface_Neighbor_Capability sample, erroring out if not present.
func (q *QualifiedLldp_Interface_Neighbor_Capability) Val(t testing.TB) *Lldp_Interface_Neighbor_Capability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lldp_Interface_Neighbor_Capability sample.
func (q *QualifiedLldp_Interface_Neighbor_Capability) SetVal(v *Lldp_Interface_Neighbor_Capability) *QualifiedLldp_Interface_Neighbor_Capability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldp_Interface_Neighbor_Capability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldp_Interface_Neighbor_Capability is a telemetry Collection whose Await method returns a slice of *Lldp_Interface_Neighbor_Capability samples.
type CollectionLldp_Interface_Neighbor_Capability struct {
	W    *Lldp_Interface_Neighbor_CapabilityWatcher
	Data []*QualifiedLldp_Interface_Neighbor_Capability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldp_Interface_Neighbor_Capability) Await(t testing.TB) []*QualifiedLldp_Interface_Neighbor_Capability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lldp_Interface_Neighbor_CapabilityWatcher observes a stream of *Lldp_Interface_Neighbor_Capability samples.
type Lldp_Interface_Neighbor_CapabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldp_Interface_Neighbor_Capability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lldp_Interface_Neighbor_CapabilityWatcher) Await(t testing.TB) (*QualifiedLldp_Interface_Neighbor_Capability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldp_Interface_Neighbor_Tlv is a *Lldp_Interface_Neighbor_Tlv with a corresponding timestamp.
type QualifiedLldp_Interface_Neighbor_Tlv struct {
	*genutil.Metadata
	val     *Lldp_Interface_Neighbor_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedLldp_Interface_Neighbor_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Lldp_Interface_Neighbor_Tlv sample, erroring out if not present.
func (q *QualifiedLldp_Interface_Neighbor_Tlv) Val(t testing.TB) *Lldp_Interface_Neighbor_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Lldp_Interface_Neighbor_Tlv sample.
func (q *QualifiedLldp_Interface_Neighbor_Tlv) SetVal(v *Lldp_Interface_Neighbor_Tlv) *QualifiedLldp_Interface_Neighbor_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldp_Interface_Neighbor_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldp_Interface_Neighbor_Tlv is a telemetry Collection whose Await method returns a slice of *Lldp_Interface_Neighbor_Tlv samples.
type CollectionLldp_Interface_Neighbor_Tlv struct {
	W    *Lldp_Interface_Neighbor_TlvWatcher
	Data []*QualifiedLldp_Interface_Neighbor_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldp_Interface_Neighbor_Tlv) Await(t testing.TB) []*QualifiedLldp_Interface_Neighbor_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Lldp_Interface_Neighbor_TlvWatcher observes a stream of *Lldp_Interface_Neighbor_Tlv samples.
type Lldp_Interface_Neighbor_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldp_Interface_Neighbor_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Lldp_Interface_Neighbor_TlvWatcher) Await(t testing.TB) (*QualifiedLldp_Interface_Neighbor_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
