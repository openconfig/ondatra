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

// QualifiedE_UnicastIpv4Prefix_Origin is a E_UnicastIpv4Prefix_Origin with a corresponding timestamp.
type QualifiedE_UnicastIpv4Prefix_Origin struct {
	*genutil.Metadata
	val     E_UnicastIpv4Prefix_Origin // val is the sample value.
	present bool
}

func (q *QualifiedE_UnicastIpv4Prefix_Origin) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_UnicastIpv4Prefix_Origin sample, erroring out if not present.
func (q *QualifiedE_UnicastIpv4Prefix_Origin) Val(t testing.TB) E_UnicastIpv4Prefix_Origin {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_UnicastIpv4Prefix_Origin sample.
func (q *QualifiedE_UnicastIpv4Prefix_Origin) SetVal(v E_UnicastIpv4Prefix_Origin) *QualifiedE_UnicastIpv4Prefix_Origin {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_UnicastIpv4Prefix_Origin) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_UnicastIpv4Prefix_Origin is a telemetry Collection whose Await method returns a slice of E_UnicastIpv4Prefix_Origin samples.
type CollectionE_UnicastIpv4Prefix_Origin struct {
	W    *E_UnicastIpv4Prefix_OriginWatcher
	Data []*QualifiedE_UnicastIpv4Prefix_Origin
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_UnicastIpv4Prefix_Origin) Await(t testing.TB) []*QualifiedE_UnicastIpv4Prefix_Origin {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_UnicastIpv4Prefix_OriginWatcher observes a stream of E_UnicastIpv4Prefix_Origin samples.
type E_UnicastIpv4Prefix_OriginWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_UnicastIpv4Prefix_Origin
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_UnicastIpv4Prefix_OriginWatcher) Await(t testing.TB) (*QualifiedE_UnicastIpv4Prefix_Origin, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_UnicastIpv6Prefix_Origin is a E_UnicastIpv6Prefix_Origin with a corresponding timestamp.
type QualifiedE_UnicastIpv6Prefix_Origin struct {
	*genutil.Metadata
	val     E_UnicastIpv6Prefix_Origin // val is the sample value.
	present bool
}

func (q *QualifiedE_UnicastIpv6Prefix_Origin) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_UnicastIpv6Prefix_Origin sample, erroring out if not present.
func (q *QualifiedE_UnicastIpv6Prefix_Origin) Val(t testing.TB) E_UnicastIpv6Prefix_Origin {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_UnicastIpv6Prefix_Origin sample.
func (q *QualifiedE_UnicastIpv6Prefix_Origin) SetVal(v E_UnicastIpv6Prefix_Origin) *QualifiedE_UnicastIpv6Prefix_Origin {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_UnicastIpv6Prefix_Origin) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_UnicastIpv6Prefix_Origin is a telemetry Collection whose Await method returns a slice of E_UnicastIpv6Prefix_Origin samples.
type CollectionE_UnicastIpv6Prefix_Origin struct {
	W    *E_UnicastIpv6Prefix_OriginWatcher
	Data []*QualifiedE_UnicastIpv6Prefix_Origin
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_UnicastIpv6Prefix_Origin) Await(t testing.TB) []*QualifiedE_UnicastIpv6Prefix_Origin {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_UnicastIpv6Prefix_OriginWatcher observes a stream of E_UnicastIpv6Prefix_Origin samples.
type E_UnicastIpv6Prefix_OriginWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_UnicastIpv6Prefix_Origin
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_UnicastIpv6Prefix_OriginWatcher) Await(t testing.TB) (*QualifiedE_UnicastIpv6Prefix_Origin, bool) {
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

// QualifiedUint32 is a uint32 with a corresponding timestamp.
type QualifiedUint32 struct {
	*genutil.Metadata
	val     uint32 // val is the sample value.
	present bool
}

func (q *QualifiedUint32) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the uint32 sample, erroring out if not present.
func (q *QualifiedUint32) Val(t testing.TB) uint32 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the uint32 sample.
func (q *QualifiedUint32) SetVal(v uint32) *QualifiedUint32 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint32) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint32 is a telemetry Collection whose Await method returns a slice of uint32 samples.
type CollectionUint32 struct {
	W    *Uint32Watcher
	Data []*QualifiedUint32
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint32) Await(t testing.TB) []*QualifiedUint32 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint32Watcher observes a stream of uint32 samples.
type Uint32Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint32
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint32Watcher) Await(t testing.TB) (*QualifiedUint32, bool) {
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
