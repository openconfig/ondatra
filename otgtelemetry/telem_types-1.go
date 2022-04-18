package otgtelemetry

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

// QualifiedMeta is a *Meta with a corresponding timestamp.
type QualifiedMeta struct {
	*genutil.Metadata
	val     *Meta // val is the sample value.
	present bool
}

func (q *QualifiedMeta) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Meta sample, erroring out if not present.
func (q *QualifiedMeta) Val(t testing.TB) *Meta {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Meta sample.
func (q *QualifiedMeta) SetVal(v *Meta) *QualifiedMeta {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedMeta) IsPresent() bool {
	return q != nil && q.present
}

// CollectionMeta is a telemetry Collection whose Await method returns a slice of *Meta samples.
type CollectionMeta struct {
	W    *MetaWatcher
	Data []*QualifiedMeta
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionMeta) Await(t testing.TB) []*QualifiedMeta {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// MetaWatcher observes a stream of *Meta samples.
type MetaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedMeta
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *MetaWatcher) Await(t testing.TB) (*QualifiedMeta, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedMeta_Window is a *Meta_Window with a corresponding timestamp.
type QualifiedMeta_Window struct {
	*genutil.Metadata
	val     *Meta_Window // val is the sample value.
	present bool
}

func (q *QualifiedMeta_Window) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Meta_Window sample, erroring out if not present.
func (q *QualifiedMeta_Window) Val(t testing.TB) *Meta_Window {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Meta_Window sample.
func (q *QualifiedMeta_Window) SetVal(v *Meta_Window) *QualifiedMeta_Window {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedMeta_Window) IsPresent() bool {
	return q != nil && q.present
}

// CollectionMeta_Window is a telemetry Collection whose Await method returns a slice of *Meta_Window samples.
type CollectionMeta_Window struct {
	W    *Meta_WindowWatcher
	Data []*QualifiedMeta_Window
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionMeta_Window) Await(t testing.TB) []*QualifiedMeta_Window {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Meta_WindowWatcher observes a stream of *Meta_Window samples.
type Meta_WindowWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedMeta_Window
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Meta_WindowWatcher) Await(t testing.TB) (*QualifiedMeta_Window, bool) {
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
