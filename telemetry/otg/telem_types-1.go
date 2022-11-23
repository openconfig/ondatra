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

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachabilityWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_PrefixWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4ExternalReachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachabilityWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_PrefixWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv4InternalReachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6ReachabilityWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_PrefixWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributesWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Ipv6Reachability_Prefix_PrefixAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachabilityWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_NeighborWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_IsReachability_Neighbor, bool) {
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

// QualifiedLldpInterface is a *LldpInterface with a corresponding timestamp.
type QualifiedLldpInterface struct {
	*genutil.Metadata
	val     *LldpInterface // val is the sample value.
	present bool
}

func (q *QualifiedLldpInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *LldpInterface sample, erroring out if not present.
func (q *QualifiedLldpInterface) Val(t testing.TB) *LldpInterface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *LldpInterface sample.
func (q *QualifiedLldpInterface) SetVal(v *LldpInterface) *QualifiedLldpInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldpInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldpInterface is a telemetry Collection whose Await method returns a slice of *LldpInterface samples.
type CollectionLldpInterface struct {
	W    *LldpInterfaceWatcher
	Data []*QualifiedLldpInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldpInterface) Await(t testing.TB) []*QualifiedLldpInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LldpInterfaceWatcher observes a stream of *LldpInterface samples.
type LldpInterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldpInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LldpInterfaceWatcher) Await(t testing.TB) (*QualifiedLldpInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldpInterface_Counters is a *LldpInterface_Counters with a corresponding timestamp.
type QualifiedLldpInterface_Counters struct {
	*genutil.Metadata
	val     *LldpInterface_Counters // val is the sample value.
	present bool
}

func (q *QualifiedLldpInterface_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *LldpInterface_Counters sample, erroring out if not present.
func (q *QualifiedLldpInterface_Counters) Val(t testing.TB) *LldpInterface_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *LldpInterface_Counters sample.
func (q *QualifiedLldpInterface_Counters) SetVal(v *LldpInterface_Counters) *QualifiedLldpInterface_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldpInterface_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldpInterface_Counters is a telemetry Collection whose Await method returns a slice of *LldpInterface_Counters samples.
type CollectionLldpInterface_Counters struct {
	W    *LldpInterface_CountersWatcher
	Data []*QualifiedLldpInterface_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldpInterface_Counters) Await(t testing.TB) []*QualifiedLldpInterface_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LldpInterface_CountersWatcher observes a stream of *LldpInterface_Counters samples.
type LldpInterface_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldpInterface_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LldpInterface_CountersWatcher) Await(t testing.TB) (*QualifiedLldpInterface_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldpInterface_LldpNeighborDatabase is a *LldpInterface_LldpNeighborDatabase with a corresponding timestamp.
type QualifiedLldpInterface_LldpNeighborDatabase struct {
	*genutil.Metadata
	val     *LldpInterface_LldpNeighborDatabase // val is the sample value.
	present bool
}

func (q *QualifiedLldpInterface_LldpNeighborDatabase) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *LldpInterface_LldpNeighborDatabase sample, erroring out if not present.
func (q *QualifiedLldpInterface_LldpNeighborDatabase) Val(t testing.TB) *LldpInterface_LldpNeighborDatabase {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *LldpInterface_LldpNeighborDatabase sample.
func (q *QualifiedLldpInterface_LldpNeighborDatabase) SetVal(v *LldpInterface_LldpNeighborDatabase) *QualifiedLldpInterface_LldpNeighborDatabase {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldpInterface_LldpNeighborDatabase) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldpInterface_LldpNeighborDatabase is a telemetry Collection whose Await method returns a slice of *LldpInterface_LldpNeighborDatabase samples.
type CollectionLldpInterface_LldpNeighborDatabase struct {
	W    *LldpInterface_LldpNeighborDatabaseWatcher
	Data []*QualifiedLldpInterface_LldpNeighborDatabase
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldpInterface_LldpNeighborDatabase) Await(t testing.TB) []*QualifiedLldpInterface_LldpNeighborDatabase {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LldpInterface_LldpNeighborDatabaseWatcher observes a stream of *LldpInterface_LldpNeighborDatabase samples.
type LldpInterface_LldpNeighborDatabaseWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldpInterface_LldpNeighborDatabase
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LldpInterface_LldpNeighborDatabaseWatcher) Await(t testing.TB) (*QualifiedLldpInterface_LldpNeighborDatabase, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor is a *LldpInterface_LldpNeighborDatabase_LldpNeighbor with a corresponding timestamp.
type QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor struct {
	*genutil.Metadata
	val     *LldpInterface_LldpNeighborDatabase_LldpNeighbor // val is the sample value.
	present bool
}

func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *LldpInterface_LldpNeighborDatabase_LldpNeighbor sample, erroring out if not present.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) Val(t testing.TB) *LldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *LldpInterface_LldpNeighborDatabase_LldpNeighbor sample.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) SetVal(v *LldpInterface_LldpNeighborDatabase_LldpNeighbor) *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor is a telemetry Collection whose Await method returns a slice of *LldpInterface_LldpNeighborDatabase_LldpNeighbor samples.
type CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor struct {
	W    *LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher
	Data []*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor) Await(t testing.TB) []*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher observes a stream of *LldpInterface_LldpNeighborDatabase_LldpNeighbor samples.
type LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LldpInterface_LldpNeighborDatabase_LldpNeighborWatcher) Await(t testing.TB) (*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities is a *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities with a corresponding timestamp.
type QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities struct {
	*genutil.Metadata
	val     *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities // val is the sample value.
	present bool
}

func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities sample, erroring out if not present.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) Val(t testing.TB) *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities sample.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) SetVal(v *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities is a telemetry Collection whose Await method returns a slice of *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities samples.
type CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities struct {
	W    *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher
	Data []*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities) Await(t testing.TB) []*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher observes a stream of *LldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities samples.
type LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CapabilitiesWatcher) Await(t testing.TB) (*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_Capabilities, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv is a *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv with a corresponding timestamp.
type QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv struct {
	*genutil.Metadata
	val     *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv // val is the sample value.
	present bool
}

func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv sample, erroring out if not present.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) Val(t testing.TB) *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv sample.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) SetVal(v *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv is a telemetry Collection whose Await method returns a slice of *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv samples.
type CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv struct {
	W    *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher
	Data []*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv) Await(t testing.TB) []*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher observes a stream of *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv samples.
type LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *LldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlvWatcher) Await(t testing.TB) (*QualifiedLldpInterface_LldpNeighborDatabase_LldpNeighbor_CustomTlv, bool) {
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

// QualifiedE_Capabilities_Name is a E_Capabilities_Name with a corresponding timestamp.
type QualifiedE_Capabilities_Name struct {
	*genutil.Metadata
	val     E_Capabilities_Name // val is the sample value.
	present bool
}

func (q *QualifiedE_Capabilities_Name) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Capabilities_Name sample, erroring out if not present.
func (q *QualifiedE_Capabilities_Name) Val(t testing.TB) E_Capabilities_Name {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Capabilities_Name sample.
func (q *QualifiedE_Capabilities_Name) SetVal(v E_Capabilities_Name) *QualifiedE_Capabilities_Name {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Capabilities_Name) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Capabilities_Name is a telemetry Collection whose Await method returns a slice of E_Capabilities_Name samples.
type CollectionE_Capabilities_Name struct {
	W    *E_Capabilities_NameWatcher
	Data []*QualifiedE_Capabilities_Name
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Capabilities_Name) Await(t testing.TB) []*QualifiedE_Capabilities_Name {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Capabilities_NameWatcher observes a stream of E_Capabilities_Name samples.
type E_Capabilities_NameWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Capabilities_Name
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Capabilities_NameWatcher) Await(t testing.TB) (*QualifiedE_Capabilities_Name, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType is a E_ExtendedIpv4Reachability_Prefix_RedistributionType with a corresponding timestamp.
type QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType struct {
	*genutil.Metadata
	val     E_ExtendedIpv4Reachability_Prefix_RedistributionType // val is the sample value.
	present bool
}

func (q *QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_ExtendedIpv4Reachability_Prefix_RedistributionType sample, erroring out if not present.
func (q *QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) Val(t testing.TB) E_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_ExtendedIpv4Reachability_Prefix_RedistributionType sample.
func (q *QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) SetVal(v E_ExtendedIpv4Reachability_Prefix_RedistributionType) *QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_ExtendedIpv4Reachability_Prefix_RedistributionType is a telemetry Collection whose Await method returns a slice of E_ExtendedIpv4Reachability_Prefix_RedistributionType samples.
type CollectionE_ExtendedIpv4Reachability_Prefix_RedistributionType struct {
	W    *E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher
	Data []*QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_ExtendedIpv4Reachability_Prefix_RedistributionType) Await(t testing.TB) []*QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher observes a stream of E_ExtendedIpv4Reachability_Prefix_RedistributionType samples.
type E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_ExtendedIpv4Reachability_Prefix_RedistributionTypeWatcher) Await(t testing.TB) (*QualifiedE_ExtendedIpv4Reachability_Prefix_RedistributionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Ipv4ExternalReachability_Prefix_OriginType is a E_Ipv4ExternalReachability_Prefix_OriginType with a corresponding timestamp.
type QualifiedE_Ipv4ExternalReachability_Prefix_OriginType struct {
	*genutil.Metadata
	val     E_Ipv4ExternalReachability_Prefix_OriginType // val is the sample value.
	present bool
}

func (q *QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Ipv4ExternalReachability_Prefix_OriginType sample, erroring out if not present.
func (q *QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) Val(t testing.TB) E_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Ipv4ExternalReachability_Prefix_OriginType sample.
func (q *QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) SetVal(v E_Ipv4ExternalReachability_Prefix_OriginType) *QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Ipv4ExternalReachability_Prefix_OriginType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Ipv4ExternalReachability_Prefix_OriginType is a telemetry Collection whose Await method returns a slice of E_Ipv4ExternalReachability_Prefix_OriginType samples.
type CollectionE_Ipv4ExternalReachability_Prefix_OriginType struct {
	W    *E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher
	Data []*QualifiedE_Ipv4ExternalReachability_Prefix_OriginType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Ipv4ExternalReachability_Prefix_OriginType) Await(t testing.TB) []*QualifiedE_Ipv4ExternalReachability_Prefix_OriginType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher observes a stream of E_Ipv4ExternalReachability_Prefix_OriginType samples.
type E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Ipv4ExternalReachability_Prefix_OriginType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Ipv4ExternalReachability_Prefix_OriginTypeWatcher) Await(t testing.TB) (*QualifiedE_Ipv4ExternalReachability_Prefix_OriginType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType is a E_Ipv4ExternalReachability_Prefix_RedistributionType with a corresponding timestamp.
type QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType struct {
	*genutil.Metadata
	val     E_Ipv4ExternalReachability_Prefix_RedistributionType // val is the sample value.
	present bool
}

func (q *QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Ipv4ExternalReachability_Prefix_RedistributionType sample, erroring out if not present.
func (q *QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) Val(t testing.TB) E_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Ipv4ExternalReachability_Prefix_RedistributionType sample.
func (q *QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) SetVal(v E_Ipv4ExternalReachability_Prefix_RedistributionType) *QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType is a telemetry Collection whose Await method returns a slice of E_Ipv4ExternalReachability_Prefix_RedistributionType samples.
type CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType struct {
	W    *E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher
	Data []*QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Ipv4ExternalReachability_Prefix_RedistributionType) Await(t testing.TB) []*QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher observes a stream of E_Ipv4ExternalReachability_Prefix_RedistributionType samples.
type E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Ipv4ExternalReachability_Prefix_RedistributionTypeWatcher) Await(t testing.TB) (*QualifiedE_Ipv4ExternalReachability_Prefix_RedistributionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
