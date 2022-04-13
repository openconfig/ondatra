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

// QualifiedE_Port_Link is a E_Port_Link with a corresponding timestamp.
type QualifiedE_Port_Link struct {
	*genutil.Metadata
	val     E_Port_Link // val is the sample value.
	present bool
}

func (q *QualifiedE_Port_Link) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Port_Link sample, erroring out if not present.
func (q *QualifiedE_Port_Link) Val(t testing.TB) E_Port_Link {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Port_Link sample.
func (q *QualifiedE_Port_Link) SetVal(v E_Port_Link) *QualifiedE_Port_Link {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Port_Link) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Port_Link is a telemetry Collection whose Await method returns a slice of E_Port_Link samples.
type CollectionE_Port_Link struct {
	W    *E_Port_LinkWatcher
	Data []*QualifiedE_Port_Link
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Port_Link) Await(t testing.TB) []*QualifiedE_Port_Link {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Port_LinkWatcher observes a stream of E_Port_Link samples.
type E_Port_LinkWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Port_Link
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Port_LinkWatcher) Await(t testing.TB) (*QualifiedE_Port_Link, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBool is a bool with a corresponding timestamp.
type QualifiedBool struct {
	*genutil.Metadata
	val     bool // val is the sample value.
	present bool
}

func (q *QualifiedBool) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the bool sample, erroring out if not present.
func (q *QualifiedBool) Val(t testing.TB) bool {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the bool sample.
func (q *QualifiedBool) SetVal(v bool) *QualifiedBool {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBool) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBool is a telemetry Collection whose Await method returns a slice of bool samples.
type CollectionBool struct {
	W    *BoolWatcher
	Data []*QualifiedBool
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBool) Await(t testing.TB) []*QualifiedBool {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BoolWatcher observes a stream of bool samples.
type BoolWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBool
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BoolWatcher) Await(t testing.TB) (*QualifiedBool, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFloat32 is a float32 with a corresponding timestamp.
type QualifiedFloat32 struct {
	*genutil.Metadata
	val     float32 // val is the sample value.
	present bool
}

func (q *QualifiedFloat32) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the float32 sample, erroring out if not present.
func (q *QualifiedFloat32) Val(t testing.TB) float32 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the float32 sample.
func (q *QualifiedFloat32) SetVal(v float32) *QualifiedFloat32 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFloat32) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFloat32 is a telemetry Collection whose Await method returns a slice of float32 samples.
type CollectionFloat32 struct {
	W    *Float32Watcher
	Data []*QualifiedFloat32
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFloat32) Await(t testing.TB) []*QualifiedFloat32 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Float32Watcher observes a stream of float32 samples.
type Float32Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFloat32
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Float32Watcher) Await(t testing.TB) (*QualifiedFloat32, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInt64 is a int64 with a corresponding timestamp.
type QualifiedInt64 struct {
	*genutil.Metadata
	val     int64 // val is the sample value.
	present bool
}

func (q *QualifiedInt64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the int64 sample, erroring out if not present.
func (q *QualifiedInt64) Val(t testing.TB) int64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the int64 sample.
func (q *QualifiedInt64) SetVal(v int64) *QualifiedInt64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInt64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInt64 is a telemetry Collection whose Await method returns a slice of int64 samples.
type CollectionInt64 struct {
	W    *Int64Watcher
	Data []*QualifiedInt64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInt64) Await(t testing.TB) []*QualifiedInt64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Int64Watcher observes a stream of int64 samples.
type Int64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInt64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Int64Watcher) Await(t testing.TB) (*QualifiedInt64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedString is a string with a corresponding timestamp.
type QualifiedString struct {
	*genutil.Metadata
	val     string // val is the sample value.
	present bool
}

func (q *QualifiedString) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the string sample, erroring out if not present.
func (q *QualifiedString) Val(t testing.TB) string {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the string sample.
func (q *QualifiedString) SetVal(v string) *QualifiedString {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedString) IsPresent() bool {
	return q != nil && q.present
}

// CollectionString is a telemetry Collection whose Await method returns a slice of string samples.
type CollectionString struct {
	W    *StringWatcher
	Data []*QualifiedString
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionString) Await(t testing.TB) []*QualifiedString {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// StringWatcher observes a stream of string samples.
type StringWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedString
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *StringWatcher) Await(t testing.TB) (*QualifiedString, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedUint64 is a uint64 with a corresponding timestamp.
type QualifiedUint64 struct {
	*genutil.Metadata
	val     uint64 // val is the sample value.
	present bool
}

func (q *QualifiedUint64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the uint64 sample, erroring out if not present.
func (q *QualifiedUint64) Val(t testing.TB) uint64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the uint64 sample.
func (q *QualifiedUint64) SetVal(v uint64) *QualifiedUint64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint64 is a telemetry Collection whose Await method returns a slice of uint64 samples.
type CollectionUint64 struct {
	W    *Uint64Watcher
	Data []*QualifiedUint64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint64) Await(t testing.TB) []*QualifiedUint64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint64Watcher observes a stream of uint64 samples.
type Uint64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint64Watcher) Await(t testing.TB) (*QualifiedUint64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
