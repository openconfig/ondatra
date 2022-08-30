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

// QualifiedNetworkInstance is a *NetworkInstance with a corresponding timestamp.
type QualifiedNetworkInstance struct {
	*genutil.Metadata
	val     *NetworkInstance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance sample, erroring out if not present.
func (q *QualifiedNetworkInstance) Val(t testing.TB) *NetworkInstance {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance sample.
func (q *QualifiedNetworkInstance) SetVal(v *NetworkInstance) *QualifiedNetworkInstance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance is a telemetry Collection whose Await method returns a slice of *NetworkInstance samples.
type CollectionNetworkInstance struct {
	W    *NetworkInstanceWatcher
	Data []*QualifiedNetworkInstance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance) Await(t testing.TB) []*QualifiedNetworkInstance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstanceWatcher observes a stream of *NetworkInstance samples.
type NetworkInstanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts is a *NetworkInstance_Afts with a corresponding timestamp.
type QualifiedNetworkInstance_Afts struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts) Val(t testing.TB) *NetworkInstance_Afts {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts sample.
func (q *QualifiedNetworkInstance_Afts) SetVal(v *NetworkInstance_Afts) *QualifiedNetworkInstance_Afts {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts samples.
type CollectionNetworkInstance_Afts struct {
	W    *NetworkInstance_AftsWatcher
	Data []*QualifiedNetworkInstance_Afts
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts) Await(t testing.TB) []*QualifiedNetworkInstance_Afts {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_AftsWatcher observes a stream of *NetworkInstance_Afts samples.
type NetworkInstance_AftsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_AftsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_Ipv4Entry is a *NetworkInstance_Afts_Ipv4Entry with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_Ipv4Entry struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_Ipv4Entry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_Ipv4Entry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_Ipv4Entry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_Ipv4Entry) Val(t testing.TB) *NetworkInstance_Afts_Ipv4Entry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_Ipv4Entry sample.
func (q *QualifiedNetworkInstance_Afts_Ipv4Entry) SetVal(v *NetworkInstance_Afts_Ipv4Entry) *QualifiedNetworkInstance_Afts_Ipv4Entry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_Ipv4Entry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_Ipv4Entry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_Ipv4Entry samples.
type CollectionNetworkInstance_Afts_Ipv4Entry struct {
	W    *NetworkInstance_Afts_Ipv4EntryWatcher
	Data []*QualifiedNetworkInstance_Afts_Ipv4Entry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_Ipv4Entry) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_Ipv4Entry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_Ipv4EntryWatcher observes a stream of *NetworkInstance_Afts_Ipv4Entry samples.
type NetworkInstance_Afts_Ipv4EntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_Ipv4Entry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_Ipv4EntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_Ipv4Entry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_Ipv4Entry_Counters is a *NetworkInstance_Afts_Ipv4Entry_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_Ipv4Entry_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_Ipv4Entry_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_Ipv4Entry_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_Ipv4Entry_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_Ipv4Entry_Counters) Val(t testing.TB) *NetworkInstance_Afts_Ipv4Entry_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_Ipv4Entry_Counters sample.
func (q *QualifiedNetworkInstance_Afts_Ipv4Entry_Counters) SetVal(v *NetworkInstance_Afts_Ipv4Entry_Counters) *QualifiedNetworkInstance_Afts_Ipv4Entry_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_Ipv4Entry_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_Ipv4Entry_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_Ipv4Entry_Counters samples.
type CollectionNetworkInstance_Afts_Ipv4Entry_Counters struct {
	W    *NetworkInstance_Afts_Ipv4Entry_CountersWatcher
	Data []*QualifiedNetworkInstance_Afts_Ipv4Entry_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_Ipv4Entry_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_Ipv4Entry_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_Ipv4Entry_CountersWatcher observes a stream of *NetworkInstance_Afts_Ipv4Entry_Counters samples.
type NetworkInstance_Afts_Ipv4Entry_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_Ipv4Entry_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_Ipv4Entry_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_Ipv4Entry_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_Ipv6Entry is a *NetworkInstance_Afts_Ipv6Entry with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_Ipv6Entry struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_Ipv6Entry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_Ipv6Entry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_Ipv6Entry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_Ipv6Entry) Val(t testing.TB) *NetworkInstance_Afts_Ipv6Entry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_Ipv6Entry sample.
func (q *QualifiedNetworkInstance_Afts_Ipv6Entry) SetVal(v *NetworkInstance_Afts_Ipv6Entry) *QualifiedNetworkInstance_Afts_Ipv6Entry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_Ipv6Entry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_Ipv6Entry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_Ipv6Entry samples.
type CollectionNetworkInstance_Afts_Ipv6Entry struct {
	W    *NetworkInstance_Afts_Ipv6EntryWatcher
	Data []*QualifiedNetworkInstance_Afts_Ipv6Entry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_Ipv6Entry) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_Ipv6Entry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_Ipv6EntryWatcher observes a stream of *NetworkInstance_Afts_Ipv6Entry samples.
type NetworkInstance_Afts_Ipv6EntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_Ipv6Entry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_Ipv6EntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_Ipv6Entry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_Ipv6Entry_Counters is a *NetworkInstance_Afts_Ipv6Entry_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_Ipv6Entry_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_Ipv6Entry_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_Ipv6Entry_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_Ipv6Entry_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_Ipv6Entry_Counters) Val(t testing.TB) *NetworkInstance_Afts_Ipv6Entry_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_Ipv6Entry_Counters sample.
func (q *QualifiedNetworkInstance_Afts_Ipv6Entry_Counters) SetVal(v *NetworkInstance_Afts_Ipv6Entry_Counters) *QualifiedNetworkInstance_Afts_Ipv6Entry_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_Ipv6Entry_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_Ipv6Entry_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_Ipv6Entry_Counters samples.
type CollectionNetworkInstance_Afts_Ipv6Entry_Counters struct {
	W    *NetworkInstance_Afts_Ipv6Entry_CountersWatcher
	Data []*QualifiedNetworkInstance_Afts_Ipv6Entry_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_Ipv6Entry_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_Ipv6Entry_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_Ipv6Entry_CountersWatcher observes a stream of *NetworkInstance_Afts_Ipv6Entry_Counters samples.
type NetworkInstance_Afts_Ipv6Entry_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_Ipv6Entry_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_Ipv6Entry_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_Ipv6Entry_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_LabelEntry is a *NetworkInstance_Afts_LabelEntry with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_LabelEntry struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_LabelEntry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_LabelEntry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_LabelEntry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_LabelEntry) Val(t testing.TB) *NetworkInstance_Afts_LabelEntry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_LabelEntry sample.
func (q *QualifiedNetworkInstance_Afts_LabelEntry) SetVal(v *NetworkInstance_Afts_LabelEntry) *QualifiedNetworkInstance_Afts_LabelEntry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_LabelEntry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_LabelEntry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_LabelEntry samples.
type CollectionNetworkInstance_Afts_LabelEntry struct {
	W    *NetworkInstance_Afts_LabelEntryWatcher
	Data []*QualifiedNetworkInstance_Afts_LabelEntry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_LabelEntry) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_LabelEntry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_LabelEntryWatcher observes a stream of *NetworkInstance_Afts_LabelEntry samples.
type NetworkInstance_Afts_LabelEntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_LabelEntry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_LabelEntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_LabelEntry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_LabelEntry_Counters is a *NetworkInstance_Afts_LabelEntry_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_LabelEntry_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_LabelEntry_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_LabelEntry_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_LabelEntry_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_Counters) Val(t testing.TB) *NetworkInstance_Afts_LabelEntry_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_LabelEntry_Counters sample.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_Counters) SetVal(v *NetworkInstance_Afts_LabelEntry_Counters) *QualifiedNetworkInstance_Afts_LabelEntry_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_LabelEntry_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_LabelEntry_Counters samples.
type CollectionNetworkInstance_Afts_LabelEntry_Counters struct {
	W    *NetworkInstance_Afts_LabelEntry_CountersWatcher
	Data []*QualifiedNetworkInstance_Afts_LabelEntry_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_LabelEntry_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_LabelEntry_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_LabelEntry_CountersWatcher observes a stream of *NetworkInstance_Afts_LabelEntry_Counters samples.
type NetworkInstance_Afts_LabelEntry_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_LabelEntry_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_LabelEntry_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_LabelEntry_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_MacEntry is a *NetworkInstance_Afts_MacEntry with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_MacEntry struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_MacEntry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_MacEntry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_MacEntry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_MacEntry) Val(t testing.TB) *NetworkInstance_Afts_MacEntry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_MacEntry sample.
func (q *QualifiedNetworkInstance_Afts_MacEntry) SetVal(v *NetworkInstance_Afts_MacEntry) *QualifiedNetworkInstance_Afts_MacEntry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_MacEntry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_MacEntry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_MacEntry samples.
type CollectionNetworkInstance_Afts_MacEntry struct {
	W    *NetworkInstance_Afts_MacEntryWatcher
	Data []*QualifiedNetworkInstance_Afts_MacEntry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_MacEntry) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_MacEntry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_MacEntryWatcher observes a stream of *NetworkInstance_Afts_MacEntry samples.
type NetworkInstance_Afts_MacEntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_MacEntry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_MacEntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_MacEntry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_MacEntry_Counters is a *NetworkInstance_Afts_MacEntry_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_MacEntry_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_MacEntry_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_MacEntry_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_MacEntry_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_MacEntry_Counters) Val(t testing.TB) *NetworkInstance_Afts_MacEntry_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_MacEntry_Counters sample.
func (q *QualifiedNetworkInstance_Afts_MacEntry_Counters) SetVal(v *NetworkInstance_Afts_MacEntry_Counters) *QualifiedNetworkInstance_Afts_MacEntry_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_MacEntry_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_MacEntry_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_MacEntry_Counters samples.
type CollectionNetworkInstance_Afts_MacEntry_Counters struct {
	W    *NetworkInstance_Afts_MacEntry_CountersWatcher
	Data []*QualifiedNetworkInstance_Afts_MacEntry_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_MacEntry_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_MacEntry_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_MacEntry_CountersWatcher observes a stream of *NetworkInstance_Afts_MacEntry_Counters samples.
type NetworkInstance_Afts_MacEntry_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_MacEntry_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_MacEntry_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_MacEntry_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHop is a *NetworkInstance_Afts_NextHop with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHop struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHop) Val(t testing.TB) *NetworkInstance_Afts_NextHop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHop sample.
func (q *QualifiedNetworkInstance_Afts_NextHop) SetVal(v *NetworkInstance_Afts_NextHop) *QualifiedNetworkInstance_Afts_NextHop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHop samples.
type CollectionNetworkInstance_Afts_NextHop struct {
	W    *NetworkInstance_Afts_NextHopWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHop) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHopWatcher observes a stream of *NetworkInstance_Afts_NextHop samples.
type NetworkInstance_Afts_NextHopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHopGroup is a *NetworkInstance_Afts_NextHopGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHopGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHopGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHopGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHopGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup) Val(t testing.TB) *NetworkInstance_Afts_NextHopGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHopGroup sample.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup) SetVal(v *NetworkInstance_Afts_NextHopGroup) *QualifiedNetworkInstance_Afts_NextHopGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHopGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHopGroup samples.
type CollectionNetworkInstance_Afts_NextHopGroup struct {
	W    *NetworkInstance_Afts_NextHopGroupWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHopGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHopGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHopGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHopGroupWatcher observes a stream of *NetworkInstance_Afts_NextHopGroup samples.
type NetworkInstance_Afts_NextHopGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHopGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHopGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHopGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHopGroup_Condition is a *NetworkInstance_Afts_NextHopGroup_Condition with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHopGroup_Condition struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHopGroup_Condition // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHopGroup_Condition sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition) Val(t testing.TB) *NetworkInstance_Afts_NextHopGroup_Condition {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHopGroup_Condition sample.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition) SetVal(v *NetworkInstance_Afts_NextHopGroup_Condition) *QualifiedNetworkInstance_Afts_NextHopGroup_Condition {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHopGroup_Condition is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHopGroup_Condition samples.
type CollectionNetworkInstance_Afts_NextHopGroup_Condition struct {
	W    *NetworkInstance_Afts_NextHopGroup_ConditionWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHopGroup_Condition
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHopGroup_Condition) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHopGroup_Condition {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHopGroup_ConditionWatcher observes a stream of *NetworkInstance_Afts_NextHopGroup_Condition samples.
type NetworkInstance_Afts_NextHopGroup_ConditionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHopGroup_Condition
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHopGroup_ConditionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHopGroup_Condition, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface is a *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface) Val(t testing.TB) *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface sample.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface) SetVal(v *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface) *QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHopGroup_Condition_InputInterface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface samples.
type CollectionNetworkInstance_Afts_NextHopGroup_Condition_InputInterface struct {
	W    *NetworkInstance_Afts_NextHopGroup_Condition_InputInterfaceWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHopGroup_Condition_InputInterface) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHopGroup_Condition_InputInterfaceWatcher observes a stream of *NetworkInstance_Afts_NextHopGroup_Condition_InputInterface samples.
type NetworkInstance_Afts_NextHopGroup_Condition_InputInterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHopGroup_Condition_InputInterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHopGroup_Condition_InputInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHopGroup_NextHop is a *NetworkInstance_Afts_NextHopGroup_NextHop with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHopGroup_NextHop struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHopGroup_NextHop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHopGroup_NextHop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHopGroup_NextHop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_NextHop) Val(t testing.TB) *NetworkInstance_Afts_NextHopGroup_NextHop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHopGroup_NextHop sample.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_NextHop) SetVal(v *NetworkInstance_Afts_NextHopGroup_NextHop) *QualifiedNetworkInstance_Afts_NextHopGroup_NextHop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHopGroup_NextHop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHopGroup_NextHop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHopGroup_NextHop samples.
type CollectionNetworkInstance_Afts_NextHopGroup_NextHop struct {
	W    *NetworkInstance_Afts_NextHopGroup_NextHopWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHopGroup_NextHop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHopGroup_NextHop) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHopGroup_NextHop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHopGroup_NextHopWatcher observes a stream of *NetworkInstance_Afts_NextHopGroup_NextHop samples.
type NetworkInstance_Afts_NextHopGroup_NextHopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHopGroup_NextHop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHopGroup_NextHopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHopGroup_NextHop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHop_Counters is a *NetworkInstance_Afts_NextHop_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHop_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHop_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHop_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHop_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHop_Counters) Val(t testing.TB) *NetworkInstance_Afts_NextHop_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHop_Counters sample.
func (q *QualifiedNetworkInstance_Afts_NextHop_Counters) SetVal(v *NetworkInstance_Afts_NextHop_Counters) *QualifiedNetworkInstance_Afts_NextHop_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHop_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHop_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHop_Counters samples.
type CollectionNetworkInstance_Afts_NextHop_Counters struct {
	W    *NetworkInstance_Afts_NextHop_CountersWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHop_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHop_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHop_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHop_CountersWatcher observes a stream of *NetworkInstance_Afts_NextHop_Counters samples.
type NetworkInstance_Afts_NextHop_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHop_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHop_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHop_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHop_InterfaceRef is a *NetworkInstance_Afts_NextHop_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHop_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHop_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHop_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHop_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHop_InterfaceRef) Val(t testing.TB) *NetworkInstance_Afts_NextHop_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHop_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Afts_NextHop_InterfaceRef) SetVal(v *NetworkInstance_Afts_NextHop_InterfaceRef) *QualifiedNetworkInstance_Afts_NextHop_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHop_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHop_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHop_InterfaceRef samples.
type CollectionNetworkInstance_Afts_NextHop_InterfaceRef struct {
	W    *NetworkInstance_Afts_NextHop_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHop_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHop_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHop_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHop_InterfaceRefWatcher observes a stream of *NetworkInstance_Afts_NextHop_InterfaceRef samples.
type NetworkInstance_Afts_NextHop_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHop_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHop_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHop_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHop_IpInIp is a *NetworkInstance_Afts_NextHop_IpInIp with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHop_IpInIp struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_NextHop_IpInIp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHop_IpInIp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_NextHop_IpInIp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHop_IpInIp) Val(t testing.TB) *NetworkInstance_Afts_NextHop_IpInIp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_NextHop_IpInIp sample.
func (q *QualifiedNetworkInstance_Afts_NextHop_IpInIp) SetVal(v *NetworkInstance_Afts_NextHop_IpInIp) *QualifiedNetworkInstance_Afts_NextHop_IpInIp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHop_IpInIp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHop_IpInIp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_NextHop_IpInIp samples.
type CollectionNetworkInstance_Afts_NextHop_IpInIp struct {
	W    *NetworkInstance_Afts_NextHop_IpInIpWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHop_IpInIp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHop_IpInIp) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHop_IpInIp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHop_IpInIpWatcher observes a stream of *NetworkInstance_Afts_NextHop_IpInIp samples.
type NetworkInstance_Afts_NextHop_IpInIpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHop_IpInIp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHop_IpInIpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHop_IpInIp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_PolicyForwardingEntry is a *NetworkInstance_Afts_PolicyForwardingEntry with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_PolicyForwardingEntry struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_PolicyForwardingEntry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_PolicyForwardingEntry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry) Val(t testing.TB) *NetworkInstance_Afts_PolicyForwardingEntry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_PolicyForwardingEntry sample.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry) SetVal(v *NetworkInstance_Afts_PolicyForwardingEntry) *QualifiedNetworkInstance_Afts_PolicyForwardingEntry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_PolicyForwardingEntry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_PolicyForwardingEntry samples.
type CollectionNetworkInstance_Afts_PolicyForwardingEntry struct {
	W    *NetworkInstance_Afts_PolicyForwardingEntryWatcher
	Data []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_PolicyForwardingEntry) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_PolicyForwardingEntryWatcher observes a stream of *NetworkInstance_Afts_PolicyForwardingEntry samples.
type NetworkInstance_Afts_PolicyForwardingEntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_PolicyForwardingEntry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_PolicyForwardingEntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_PolicyForwardingEntry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters is a *NetworkInstance_Afts_PolicyForwardingEntry_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_PolicyForwardingEntry_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_PolicyForwardingEntry_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters) Val(t testing.TB) *NetworkInstance_Afts_PolicyForwardingEntry_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_PolicyForwardingEntry_Counters sample.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters) SetVal(v *NetworkInstance_Afts_PolicyForwardingEntry_Counters) *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_PolicyForwardingEntry_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_PolicyForwardingEntry_Counters samples.
type CollectionNetworkInstance_Afts_PolicyForwardingEntry_Counters struct {
	W    *NetworkInstance_Afts_PolicyForwardingEntry_CountersWatcher
	Data []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_PolicyForwardingEntry_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_PolicyForwardingEntry_CountersWatcher observes a stream of *NetworkInstance_Afts_PolicyForwardingEntry_Counters samples.
type NetworkInstance_Afts_PolicyForwardingEntry_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_PolicyForwardingEntry_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_StateSynced is a *NetworkInstance_Afts_StateSynced with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_StateSynced struct {
	*genutil.Metadata
	val     *NetworkInstance_Afts_StateSynced // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_StateSynced) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Afts_StateSynced sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_StateSynced) Val(t testing.TB) *NetworkInstance_Afts_StateSynced {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Afts_StateSynced sample.
func (q *QualifiedNetworkInstance_Afts_StateSynced) SetVal(v *NetworkInstance_Afts_StateSynced) *QualifiedNetworkInstance_Afts_StateSynced {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_StateSynced) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_StateSynced is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Afts_StateSynced samples.
type CollectionNetworkInstance_Afts_StateSynced struct {
	W    *NetworkInstance_Afts_StateSyncedWatcher
	Data []*QualifiedNetworkInstance_Afts_StateSynced
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_StateSynced) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_StateSynced {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_StateSyncedWatcher observes a stream of *NetworkInstance_Afts_StateSynced samples.
type NetworkInstance_Afts_StateSyncedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_StateSynced
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_StateSyncedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_StateSynced, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint is a *NetworkInstance_ConnectionPoint with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint struct {
	*genutil.Metadata
	val     *NetworkInstance_ConnectionPoint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_ConnectionPoint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint) Val(t testing.TB) *NetworkInstance_ConnectionPoint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_ConnectionPoint sample.
func (q *QualifiedNetworkInstance_ConnectionPoint) SetVal(v *NetworkInstance_ConnectionPoint) *QualifiedNetworkInstance_ConnectionPoint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_ConnectionPoint samples.
type CollectionNetworkInstance_ConnectionPoint struct {
	W    *NetworkInstance_ConnectionPointWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPointWatcher observes a stream of *NetworkInstance_ConnectionPoint samples.
type NetworkInstance_ConnectionPointWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPointWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint_Endpoint is a *NetworkInstance_ConnectionPoint_Endpoint with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint_Endpoint struct {
	*genutil.Metadata
	val     *NetworkInstance_ConnectionPoint_Endpoint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_ConnectionPoint_Endpoint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint) Val(t testing.TB) *NetworkInstance_ConnectionPoint_Endpoint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_ConnectionPoint_Endpoint sample.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint) SetVal(v *NetworkInstance_ConnectionPoint_Endpoint) *QualifiedNetworkInstance_ConnectionPoint_Endpoint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint_Endpoint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_ConnectionPoint_Endpoint samples.
type CollectionNetworkInstance_ConnectionPoint_Endpoint struct {
	W    *NetworkInstance_ConnectionPoint_EndpointWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint_Endpoint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint_Endpoint) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint_Endpoint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPoint_EndpointWatcher observes a stream of *NetworkInstance_ConnectionPoint_Endpoint samples.
type NetworkInstance_ConnectionPoint_EndpointWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint_Endpoint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPoint_EndpointWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint_Endpoint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local is a *NetworkInstance_ConnectionPoint_Endpoint_Local with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local struct {
	*genutil.Metadata
	val     *NetworkInstance_ConnectionPoint_Endpoint_Local // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_ConnectionPoint_Endpoint_Local sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local) Val(t testing.TB) *NetworkInstance_ConnectionPoint_Endpoint_Local {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_ConnectionPoint_Endpoint_Local sample.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local) SetVal(v *NetworkInstance_ConnectionPoint_Endpoint_Local) *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint_Endpoint_Local is a telemetry Collection whose Await method returns a slice of *NetworkInstance_ConnectionPoint_Endpoint_Local samples.
type CollectionNetworkInstance_ConnectionPoint_Endpoint_Local struct {
	W    *NetworkInstance_ConnectionPoint_Endpoint_LocalWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint_Endpoint_Local) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPoint_Endpoint_LocalWatcher observes a stream of *NetworkInstance_ConnectionPoint_Endpoint_Local samples.
type NetworkInstance_ConnectionPoint_Endpoint_LocalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPoint_Endpoint_LocalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Local, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote is a *NetworkInstance_ConnectionPoint_Endpoint_Remote with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote struct {
	*genutil.Metadata
	val     *NetworkInstance_ConnectionPoint_Endpoint_Remote // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_ConnectionPoint_Endpoint_Remote sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote) Val(t testing.TB) *NetworkInstance_ConnectionPoint_Endpoint_Remote {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_ConnectionPoint_Endpoint_Remote sample.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote) SetVal(v *NetworkInstance_ConnectionPoint_Endpoint_Remote) *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint_Endpoint_Remote is a telemetry Collection whose Await method returns a slice of *NetworkInstance_ConnectionPoint_Endpoint_Remote samples.
type CollectionNetworkInstance_ConnectionPoint_Endpoint_Remote struct {
	W    *NetworkInstance_ConnectionPoint_Endpoint_RemoteWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint_Endpoint_Remote) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPoint_Endpoint_RemoteWatcher observes a stream of *NetworkInstance_ConnectionPoint_Endpoint_Remote samples.
type NetworkInstance_ConnectionPoint_Endpoint_RemoteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPoint_Endpoint_RemoteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Remote, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan is a *NetworkInstance_ConnectionPoint_Endpoint_Vxlan with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan struct {
	*genutil.Metadata
	val     *NetworkInstance_ConnectionPoint_Endpoint_Vxlan // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_ConnectionPoint_Endpoint_Vxlan sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan) Val(t testing.TB) *NetworkInstance_ConnectionPoint_Endpoint_Vxlan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_ConnectionPoint_Endpoint_Vxlan sample.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan) SetVal(v *NetworkInstance_ConnectionPoint_Endpoint_Vxlan) *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan is a telemetry Collection whose Await method returns a slice of *NetworkInstance_ConnectionPoint_Endpoint_Vxlan samples.
type CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan struct {
	W    *NetworkInstance_ConnectionPoint_Endpoint_VxlanWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPoint_Endpoint_VxlanWatcher observes a stream of *NetworkInstance_ConnectionPoint_Endpoint_Vxlan samples.
type NetworkInstance_ConnectionPoint_Endpoint_VxlanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPoint_Endpoint_VxlanWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer is a *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer struct {
	*genutil.Metadata
	val     *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer) Val(t testing.TB) *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer sample.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer) SetVal(v *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer) *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer is a telemetry Collection whose Await method returns a slice of *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer samples.
type CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer struct {
	W    *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeerWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeerWatcher observes a stream of *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer samples.
type NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeerWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointPeer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni is a *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni struct {
	*genutil.Metadata
	val     *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni) Val(t testing.TB) *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni sample.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni) SetVal(v *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni) *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni is a telemetry Collection whose Await method returns a slice of *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni samples.
type CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni struct {
	W    *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVniWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVniWatcher observes a stream of *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni samples.
type NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVniWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVniWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Encapsulation is a *NetworkInstance_Encapsulation with a corresponding timestamp.
type QualifiedNetworkInstance_Encapsulation struct {
	*genutil.Metadata
	val     *NetworkInstance_Encapsulation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Encapsulation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Encapsulation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Encapsulation) Val(t testing.TB) *NetworkInstance_Encapsulation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Encapsulation sample.
func (q *QualifiedNetworkInstance_Encapsulation) SetVal(v *NetworkInstance_Encapsulation) *QualifiedNetworkInstance_Encapsulation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Encapsulation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Encapsulation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Encapsulation samples.
type CollectionNetworkInstance_Encapsulation struct {
	W    *NetworkInstance_EncapsulationWatcher
	Data []*QualifiedNetworkInstance_Encapsulation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Encapsulation) Await(t testing.TB) []*QualifiedNetworkInstance_Encapsulation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_EncapsulationWatcher observes a stream of *NetworkInstance_Encapsulation samples.
type NetworkInstance_EncapsulationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Encapsulation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_EncapsulationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Encapsulation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn is a *NetworkInstance_Evpn with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn) Val(t testing.TB) *NetworkInstance_Evpn {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn sample.
func (q *QualifiedNetworkInstance_Evpn) SetVal(v *NetworkInstance_Evpn) *QualifiedNetworkInstance_Evpn {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn samples.
type CollectionNetworkInstance_Evpn struct {
	W    *NetworkInstance_EvpnWatcher
	Data []*QualifiedNetworkInstance_Evpn
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_EvpnWatcher observes a stream of *NetworkInstance_Evpn samples.
type NetworkInstance_EvpnWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_EvpnWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EthernetSegment is a *NetworkInstance_Evpn_EthernetSegment with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EthernetSegment struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EthernetSegment // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EthernetSegment) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EthernetSegment sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment) Val(t testing.TB) *NetworkInstance_Evpn_EthernetSegment {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EthernetSegment sample.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment) SetVal(v *NetworkInstance_Evpn_EthernetSegment) *QualifiedNetworkInstance_Evpn_EthernetSegment {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EthernetSegment is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EthernetSegment samples.
type CollectionNetworkInstance_Evpn_EthernetSegment struct {
	W    *NetworkInstance_Evpn_EthernetSegmentWatcher
	Data []*QualifiedNetworkInstance_Evpn_EthernetSegment
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EthernetSegment) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EthernetSegment {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EthernetSegmentWatcher observes a stream of *NetworkInstance_Evpn_EthernetSegment samples.
type NetworkInstance_Evpn_EthernetSegmentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EthernetSegment
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EthernetSegmentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EthernetSegment, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection is a *NetworkInstance_Evpn_EthernetSegment_DfElection with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EthernetSegment_DfElection // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EthernetSegment_DfElection sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection) Val(t testing.TB) *NetworkInstance_Evpn_EthernetSegment_DfElection {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EthernetSegment_DfElection sample.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection) SetVal(v *NetworkInstance_Evpn_EthernetSegment_DfElection) *QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EthernetSegment_DfElection is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EthernetSegment_DfElection samples.
type CollectionNetworkInstance_Evpn_EthernetSegment_DfElection struct {
	W    *NetworkInstance_Evpn_EthernetSegment_DfElectionWatcher
	Data []*QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EthernetSegment_DfElection) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EthernetSegment_DfElectionWatcher observes a stream of *NetworkInstance_Evpn_EthernetSegment_DfElection samples.
type NetworkInstance_Evpn_EthernetSegment_DfElectionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EthernetSegment_DfElectionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EthernetSegment_DfElection, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance is a *NetworkInstance_Evpn_EvpnInstance with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EvpnInstance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EvpnInstance sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance) Val(t testing.TB) *NetworkInstance_Evpn_EvpnInstance {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EvpnInstance sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance) SetVal(v *NetworkInstance_Evpn_EvpnInstance) *QualifiedNetworkInstance_Evpn_EvpnInstance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EvpnInstance samples.
type CollectionNetworkInstance_Evpn_EvpnInstance struct {
	W    *NetworkInstance_Evpn_EvpnInstanceWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstanceWatcher observes a stream of *NetworkInstance_Evpn_EvpnInstance samples.
type NetworkInstance_Evpn_EvpnInstanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent is a *NetworkInstance_Evpn_EvpnInstance_BComponent with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EvpnInstance_BComponent // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EvpnInstance_BComponent sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent) Val(t testing.TB) *NetworkInstance_Evpn_EvpnInstance_BComponent {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EvpnInstance_BComponent sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent) SetVal(v *NetworkInstance_Evpn_EvpnInstance_BComponent) *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_BComponent is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EvpnInstance_BComponent samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_BComponent struct {
	W    *NetworkInstance_Evpn_EvpnInstance_BComponentWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_BComponent) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_BComponentWatcher observes a stream of *NetworkInstance_Evpn_EvpnInstance_BComponent samples.
type NetworkInstance_Evpn_EvpnInstance_BComponentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_BComponentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent is a *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent) Val(t testing.TB) *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent) SetVal(v *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent) *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent struct {
	W    *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponentWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_BComponent_IComponentWatcher observes a stream of *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponent samples.
type NetworkInstance_Evpn_EvpnInstance_BComponent_IComponentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_BComponent_IComponentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_BComponent_IComponent, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy is a *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy) Val(t testing.TB) *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy) SetVal(v *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy) *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy struct {
	W    *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicyWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_ImportExportPolicyWatcher observes a stream of *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy samples.
type NetworkInstance_Evpn_EvpnInstance_ImportExportPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan is a *NetworkInstance_Evpn_EvpnInstance_Vxlan with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EvpnInstance_Vxlan // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EvpnInstance_Vxlan sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan) Val(t testing.TB) *NetworkInstance_Evpn_EvpnInstance_Vxlan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EvpnInstance_Vxlan sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan) SetVal(v *NetworkInstance_Evpn_EvpnInstance_Vxlan) *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_Vxlan is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EvpnInstance_Vxlan samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_Vxlan struct {
	W    *NetworkInstance_Evpn_EvpnInstance_VxlanWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_Vxlan) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_VxlanWatcher observes a stream of *NetworkInstance_Evpn_EvpnInstance_Vxlan samples.
type NetworkInstance_Evpn_EvpnInstance_VxlanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_VxlanWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface is a *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface struct {
	*genutil.Metadata
	val     *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface) Val(t testing.TB) *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface) SetVal(v *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface) *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface struct {
	W    *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterfaceWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterfaceWatcher observes a stream of *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface samples.
type NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_Vxlan_AnycastSourceInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb is a *NetworkInstance_Fdb with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb) Val(t testing.TB) *NetworkInstance_Fdb {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb sample.
func (q *QualifiedNetworkInstance_Fdb) SetVal(v *NetworkInstance_Fdb) *QualifiedNetworkInstance_Fdb {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb samples.
type CollectionNetworkInstance_Fdb struct {
	W    *NetworkInstance_FdbWatcher
	Data []*QualifiedNetworkInstance_Fdb
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_FdbWatcher observes a stream of *NetworkInstance_Fdb samples.
type NetworkInstance_FdbWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_FdbWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_ArpProxy is a *NetworkInstance_Fdb_ArpProxy with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_ArpProxy struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_ArpProxy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_ArpProxy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_ArpProxy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_ArpProxy) Val(t testing.TB) *NetworkInstance_Fdb_ArpProxy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_ArpProxy sample.
func (q *QualifiedNetworkInstance_Fdb_ArpProxy) SetVal(v *NetworkInstance_Fdb_ArpProxy) *QualifiedNetworkInstance_Fdb_ArpProxy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_ArpProxy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_ArpProxy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_ArpProxy samples.
type CollectionNetworkInstance_Fdb_ArpProxy struct {
	W    *NetworkInstance_Fdb_ArpProxyWatcher
	Data []*QualifiedNetworkInstance_Fdb_ArpProxy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_ArpProxy) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_ArpProxy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_ArpProxyWatcher observes a stream of *NetworkInstance_Fdb_ArpProxy samples.
type NetworkInstance_Fdb_ArpProxyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_ArpProxy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_ArpProxyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_ArpProxy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib is a *NetworkInstance_Fdb_L2Rib with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib) SetVal(v *NetworkInstance_Fdb_L2Rib) *QualifiedNetworkInstance_Fdb_L2Rib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib samples.
type CollectionNetworkInstance_Fdb_L2Rib struct {
	W    *NetworkInstance_Fdb_L2RibWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2RibWatcher observes a stream of *NetworkInstance_Fdb_L2Rib samples.
type NetworkInstance_Fdb_L2RibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2RibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable is a *NetworkInstance_Fdb_L2Rib_MacIpTable with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacIpTable // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacIpTable {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable) SetVal(v *NetworkInstance_Fdb_L2Rib_MacIpTable) *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacIpTable is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacIpTable samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacIpTable struct {
	W    *NetworkInstance_Fdb_L2Rib_MacIpTableWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacIpTable) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacIpTableWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacIpTable samples.
type NetworkInstance_Fdb_L2Rib_MacIpTableWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacIpTableWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry is a *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry) SetVal(v *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry) *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_Entry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_Entry struct {
	W    *NetworkInstance_Fdb_L2Rib_MacIpTable_EntryWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_Entry) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacIpTable_EntryWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry samples.
type NetworkInstance_Fdb_L2Rib_MacIpTable_EntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacIpTable_EntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer is a *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer) SetVal(v *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer) *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer struct {
	W    *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_ProducerWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_ProducerWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer samples.
type NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_ProducerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacIpTable_Entry_ProducerWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_Entry_Producer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop is a *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop) SetVal(v *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop) *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop struct {
	W    *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHopWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacIpTable_NextHopWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHop samples.
type NetworkInstance_Fdb_L2Rib_MacIpTable_NextHopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacIpTable_NextHopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacIpTable_NextHop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacTable is a *NetworkInstance_Fdb_L2Rib_MacTable with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacTable struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacTable // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacTable sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacTable {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacTable sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable) SetVal(v *NetworkInstance_Fdb_L2Rib_MacTable) *QualifiedNetworkInstance_Fdb_L2Rib_MacTable {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacTable is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacTable samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacTable struct {
	W    *NetworkInstance_Fdb_L2Rib_MacTableWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacTable) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacTableWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacTable samples.
type NetworkInstance_Fdb_L2Rib_MacTableWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacTable
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacTableWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacTable, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry is a *NetworkInstance_Fdb_L2Rib_MacTable_Entry with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacTable_Entry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacTable_Entry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacTable_Entry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacTable_Entry sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry) SetVal(v *NetworkInstance_Fdb_L2Rib_MacTable_Entry) *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacTable_Entry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacTable_Entry samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacTable_Entry struct {
	W    *NetworkInstance_Fdb_L2Rib_MacTable_EntryWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacTable_Entry) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacTable_EntryWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacTable_Entry samples.
type NetworkInstance_Fdb_L2Rib_MacTable_EntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacTable_EntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer is a *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer) SetVal(v *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer) *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer struct {
	W    *NetworkInstance_Fdb_L2Rib_MacTable_Entry_ProducerWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacTable_Entry_ProducerWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer samples.
type NetworkInstance_Fdb_L2Rib_MacTable_Entry_ProducerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacTable_Entry_ProducerWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_Entry_Producer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop is a *NetworkInstance_Fdb_L2Rib_MacTable_NextHop with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_L2Rib_MacTable_NextHop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_L2Rib_MacTable_NextHop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop) Val(t testing.TB) *NetworkInstance_Fdb_L2Rib_MacTable_NextHop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_L2Rib_MacTable_NextHop sample.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop) SetVal(v *NetworkInstance_Fdb_L2Rib_MacTable_NextHop) *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_L2Rib_MacTable_NextHop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_L2Rib_MacTable_NextHop samples.
type CollectionNetworkInstance_Fdb_L2Rib_MacTable_NextHop struct {
	W    *NetworkInstance_Fdb_L2Rib_MacTable_NextHopWatcher
	Data []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_L2Rib_MacTable_NextHop) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_L2Rib_MacTable_NextHopWatcher observes a stream of *NetworkInstance_Fdb_L2Rib_MacTable_NextHop samples.
type NetworkInstance_Fdb_L2Rib_MacTable_NextHopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_L2Rib_MacTable_NextHopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_L2Rib_MacTable_NextHop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_MacMobility is a *NetworkInstance_Fdb_MacMobility with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_MacMobility struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_MacMobility // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_MacMobility) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_MacMobility sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_MacMobility) Val(t testing.TB) *NetworkInstance_Fdb_MacMobility {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_MacMobility sample.
func (q *QualifiedNetworkInstance_Fdb_MacMobility) SetVal(v *NetworkInstance_Fdb_MacMobility) *QualifiedNetworkInstance_Fdb_MacMobility {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_MacMobility) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_MacMobility is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_MacMobility samples.
type CollectionNetworkInstance_Fdb_MacMobility struct {
	W    *NetworkInstance_Fdb_MacMobilityWatcher
	Data []*QualifiedNetworkInstance_Fdb_MacMobility
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_MacMobility) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_MacMobility {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_MacMobilityWatcher observes a stream of *NetworkInstance_Fdb_MacMobility samples.
type NetworkInstance_Fdb_MacMobilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_MacMobility
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_MacMobilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_MacMobility, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_MacTable is a *NetworkInstance_Fdb_MacTable with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_MacTable struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_MacTable // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_MacTable) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_MacTable sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_MacTable) Val(t testing.TB) *NetworkInstance_Fdb_MacTable {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_MacTable sample.
func (q *QualifiedNetworkInstance_Fdb_MacTable) SetVal(v *NetworkInstance_Fdb_MacTable) *QualifiedNetworkInstance_Fdb_MacTable {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_MacTable) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_MacTable is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_MacTable samples.
type CollectionNetworkInstance_Fdb_MacTable struct {
	W    *NetworkInstance_Fdb_MacTableWatcher
	Data []*QualifiedNetworkInstance_Fdb_MacTable
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_MacTable) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_MacTable {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_MacTableWatcher observes a stream of *NetworkInstance_Fdb_MacTable samples.
type NetworkInstance_Fdb_MacTableWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_MacTable
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_MacTableWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_MacTable, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_MacTable_Entry is a *NetworkInstance_Fdb_MacTable_Entry with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_MacTable_Entry struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_MacTable_Entry // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_MacTable_Entry sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry) Val(t testing.TB) *NetworkInstance_Fdb_MacTable_Entry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_MacTable_Entry sample.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry) SetVal(v *NetworkInstance_Fdb_MacTable_Entry) *QualifiedNetworkInstance_Fdb_MacTable_Entry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_MacTable_Entry is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_MacTable_Entry samples.
type CollectionNetworkInstance_Fdb_MacTable_Entry struct {
	W    *NetworkInstance_Fdb_MacTable_EntryWatcher
	Data []*QualifiedNetworkInstance_Fdb_MacTable_Entry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_MacTable_Entry) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_MacTable_Entry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_MacTable_EntryWatcher observes a stream of *NetworkInstance_Fdb_MacTable_Entry samples.
type NetworkInstance_Fdb_MacTable_EntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_MacTable_Entry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_MacTable_EntryWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_MacTable_Entry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface is a *NetworkInstance_Fdb_MacTable_Entry_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_MacTable_Entry_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_MacTable_Entry_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface) Val(t testing.TB) *NetworkInstance_Fdb_MacTable_Entry_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_MacTable_Entry_Interface sample.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface) SetVal(v *NetworkInstance_Fdb_MacTable_Entry_Interface) *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_MacTable_Entry_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_MacTable_Entry_Interface samples.
type CollectionNetworkInstance_Fdb_MacTable_Entry_Interface struct {
	W    *NetworkInstance_Fdb_MacTable_Entry_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_MacTable_Entry_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_MacTable_Entry_InterfaceWatcher observes a stream of *NetworkInstance_Fdb_MacTable_Entry_Interface samples.
type NetworkInstance_Fdb_MacTable_Entry_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_MacTable_Entry_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef is a *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef) SetVal(v *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef) *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef struct {
	W    *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef samples.
type NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_MacTable_Entry_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Fdb_NdProxy is a *NetworkInstance_Fdb_NdProxy with a corresponding timestamp.
type QualifiedNetworkInstance_Fdb_NdProxy struct {
	*genutil.Metadata
	val     *NetworkInstance_Fdb_NdProxy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Fdb_NdProxy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Fdb_NdProxy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Fdb_NdProxy) Val(t testing.TB) *NetworkInstance_Fdb_NdProxy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Fdb_NdProxy sample.
func (q *QualifiedNetworkInstance_Fdb_NdProxy) SetVal(v *NetworkInstance_Fdb_NdProxy) *QualifiedNetworkInstance_Fdb_NdProxy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Fdb_NdProxy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Fdb_NdProxy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Fdb_NdProxy samples.
type CollectionNetworkInstance_Fdb_NdProxy struct {
	W    *NetworkInstance_Fdb_NdProxyWatcher
	Data []*QualifiedNetworkInstance_Fdb_NdProxy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Fdb_NdProxy) Await(t testing.TB) []*QualifiedNetworkInstance_Fdb_NdProxy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Fdb_NdProxyWatcher observes a stream of *NetworkInstance_Fdb_NdProxy samples.
type NetworkInstance_Fdb_NdProxyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Fdb_NdProxy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Fdb_NdProxyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Fdb_NdProxy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_InterInstancePolicies is a *NetworkInstance_InterInstancePolicies with a corresponding timestamp.
type QualifiedNetworkInstance_InterInstancePolicies struct {
	*genutil.Metadata
	val     *NetworkInstance_InterInstancePolicies // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_InterInstancePolicies) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_InterInstancePolicies sample, erroring out if not present.
func (q *QualifiedNetworkInstance_InterInstancePolicies) Val(t testing.TB) *NetworkInstance_InterInstancePolicies {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_InterInstancePolicies sample.
func (q *QualifiedNetworkInstance_InterInstancePolicies) SetVal(v *NetworkInstance_InterInstancePolicies) *QualifiedNetworkInstance_InterInstancePolicies {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_InterInstancePolicies) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_InterInstancePolicies is a telemetry Collection whose Await method returns a slice of *NetworkInstance_InterInstancePolicies samples.
type CollectionNetworkInstance_InterInstancePolicies struct {
	W    *NetworkInstance_InterInstancePoliciesWatcher
	Data []*QualifiedNetworkInstance_InterInstancePolicies
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_InterInstancePolicies) Await(t testing.TB) []*QualifiedNetworkInstance_InterInstancePolicies {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_InterInstancePoliciesWatcher observes a stream of *NetworkInstance_InterInstancePolicies samples.
type NetworkInstance_InterInstancePoliciesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_InterInstancePolicies
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_InterInstancePoliciesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_InterInstancePolicies, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy is a *NetworkInstance_InterInstancePolicies_ApplyPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_InterInstancePolicies_ApplyPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_InterInstancePolicies_ApplyPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy) Val(t testing.TB) *NetworkInstance_InterInstancePolicies_ApplyPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_InterInstancePolicies_ApplyPolicy sample.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy) SetVal(v *NetworkInstance_InterInstancePolicies_ApplyPolicy) *QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_InterInstancePolicies_ApplyPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_InterInstancePolicies_ApplyPolicy samples.
type CollectionNetworkInstance_InterInstancePolicies_ApplyPolicy struct {
	W    *NetworkInstance_InterInstancePolicies_ApplyPolicyWatcher
	Data []*QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_InterInstancePolicies_ApplyPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_InterInstancePolicies_ApplyPolicyWatcher observes a stream of *NetworkInstance_InterInstancePolicies_ApplyPolicy samples.
type NetworkInstance_InterInstancePolicies_ApplyPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_InterInstancePolicies_ApplyPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_InterInstancePolicies_ApplyPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy is a *NetworkInstance_InterInstancePolicies_ImportExportPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_InterInstancePolicies_ImportExportPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_InterInstancePolicies_ImportExportPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy) Val(t testing.TB) *NetworkInstance_InterInstancePolicies_ImportExportPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_InterInstancePolicies_ImportExportPolicy sample.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy) SetVal(v *NetworkInstance_InterInstancePolicies_ImportExportPolicy) *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_InterInstancePolicies_ImportExportPolicy samples.
type CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy struct {
	W    *NetworkInstance_InterInstancePolicies_ImportExportPolicyWatcher
	Data []*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_InterInstancePolicies_ImportExportPolicyWatcher observes a stream of *NetworkInstance_InterInstancePolicies_ImportExportPolicy samples.
type NetworkInstance_InterInstancePolicies_ImportExportPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_InterInstancePolicies_ImportExportPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Interface is a *NetworkInstance_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Interface) Val(t testing.TB) *NetworkInstance_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Interface sample.
func (q *QualifiedNetworkInstance_Interface) SetVal(v *NetworkInstance_Interface) *QualifiedNetworkInstance_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Interface samples.
type CollectionNetworkInstance_Interface struct {
	W    *NetworkInstance_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_InterfaceWatcher observes a stream of *NetworkInstance_Interface samples.
type NetworkInstance_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls is a *NetworkInstance_Mpls with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls) Val(t testing.TB) *NetworkInstance_Mpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls sample.
func (q *QualifiedNetworkInstance_Mpls) SetVal(v *NetworkInstance_Mpls) *QualifiedNetworkInstance_Mpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls samples.
type CollectionNetworkInstance_Mpls struct {
	W    *NetworkInstance_MplsWatcher
	Data []*QualifiedNetworkInstance_Mpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_MplsWatcher observes a stream of *NetworkInstance_Mpls samples.
type NetworkInstance_MplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_MplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Global is a *NetworkInstance_Mpls_Global with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Global struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Global // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Global) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Global sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Global) Val(t testing.TB) *NetworkInstance_Mpls_Global {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Global sample.
func (q *QualifiedNetworkInstance_Mpls_Global) SetVal(v *NetworkInstance_Mpls_Global) *QualifiedNetworkInstance_Mpls_Global {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Global) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Global is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Global samples.
type CollectionNetworkInstance_Mpls_Global struct {
	W    *NetworkInstance_Mpls_GlobalWatcher
	Data []*QualifiedNetworkInstance_Mpls_Global
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Global) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Global {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_GlobalWatcher observes a stream of *NetworkInstance_Mpls_Global samples.
type NetworkInstance_Mpls_GlobalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Global
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_GlobalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Global, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Global_Interface is a *NetworkInstance_Mpls_Global_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Global_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Global_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Global_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Global_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Global_Interface) Val(t testing.TB) *NetworkInstance_Mpls_Global_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Global_Interface sample.
func (q *QualifiedNetworkInstance_Mpls_Global_Interface) SetVal(v *NetworkInstance_Mpls_Global_Interface) *QualifiedNetworkInstance_Mpls_Global_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Global_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Global_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Global_Interface samples.
type CollectionNetworkInstance_Mpls_Global_Interface struct {
	W    *NetworkInstance_Mpls_Global_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Mpls_Global_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Global_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Global_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Global_InterfaceWatcher observes a stream of *NetworkInstance_Mpls_Global_Interface samples.
type NetworkInstance_Mpls_Global_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Global_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Global_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Global_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef is a *NetworkInstance_Mpls_Global_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Global_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Global_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Mpls_Global_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Global_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef) SetVal(v *NetworkInstance_Mpls_Global_Interface_InterfaceRef) *QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Global_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Global_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Mpls_Global_Interface_InterfaceRef struct {
	W    *NetworkInstance_Mpls_Global_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Global_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Global_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Mpls_Global_Interface_InterfaceRef samples.
type NetworkInstance_Mpls_Global_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Global_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Global_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock is a *NetworkInstance_Mpls_Global_ReservedLabelBlock with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Global_ReservedLabelBlock // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Global_ReservedLabelBlock sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock) Val(t testing.TB) *NetworkInstance_Mpls_Global_ReservedLabelBlock {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Global_ReservedLabelBlock sample.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock) SetVal(v *NetworkInstance_Mpls_Global_ReservedLabelBlock) *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Global_ReservedLabelBlock samples.
type CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock struct {
	W    *NetworkInstance_Mpls_Global_ReservedLabelBlockWatcher
	Data []*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Global_ReservedLabelBlockWatcher observes a stream of *NetworkInstance_Mpls_Global_ReservedLabelBlock samples.
type NetworkInstance_Mpls_Global_ReservedLabelBlockWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Global_ReservedLabelBlockWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Interface is a *NetworkInstance_Mpls_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Interface) Val(t testing.TB) *NetworkInstance_Mpls_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Interface sample.
func (q *QualifiedNetworkInstance_Mpls_Interface) SetVal(v *NetworkInstance_Mpls_Interface) *QualifiedNetworkInstance_Mpls_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Interface samples.
type CollectionNetworkInstance_Mpls_Interface struct {
	W    *NetworkInstance_Mpls_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Mpls_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_InterfaceWatcher observes a stream of *NetworkInstance_Mpls_Interface samples.
type NetworkInstance_Mpls_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth is a *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth) Val(t testing.TB) *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth sample.
func (q *QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth) SetVal(v *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth) *QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Interface_IgpFloodingBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth samples.
type CollectionNetworkInstance_Mpls_Interface_IgpFloodingBandwidth struct {
	W    *NetworkInstance_Mpls_Interface_IgpFloodingBandwidthWatcher
	Data []*QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Interface_IgpFloodingBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Interface_IgpFloodingBandwidthWatcher observes a stream of *NetworkInstance_Mpls_Interface_IgpFloodingBandwidth samples.
type NetworkInstance_Mpls_Interface_IgpFloodingBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Interface_IgpFloodingBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Interface_IgpFloodingBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Interface_InterfaceRef is a *NetworkInstance_Mpls_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Mpls_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Mpls_Interface_InterfaceRef) SetVal(v *NetworkInstance_Mpls_Interface_InterfaceRef) *QualifiedNetworkInstance_Mpls_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Mpls_Interface_InterfaceRef struct {
	W    *NetworkInstance_Mpls_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Mpls_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Mpls_Interface_InterfaceRef samples.
type NetworkInstance_Mpls_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps is a *NetworkInstance_Mpls_Lsps with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps) Val(t testing.TB) *NetworkInstance_Mpls_Lsps {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps) SetVal(v *NetworkInstance_Mpls_Lsps) *QualifiedNetworkInstance_Mpls_Lsps {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps samples.
type CollectionNetworkInstance_Mpls_Lsps struct {
	W    *NetworkInstance_Mpls_LspsWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_LspsWatcher observes a stream of *NetworkInstance_Mpls_Lsps samples.
type NetworkInstance_Mpls_LspsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_LspsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath is a *NetworkInstance_Mpls_Lsps_ConstrainedPath with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPathWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPathWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPathWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPathWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObjectWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObjectWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObjectWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObjectWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_NamedExplicitPath_ExplicitRouteObject, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_TunnelWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_TunnelWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_TunnelWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_TunnelWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_BandwidthWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_BandwidthWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_BandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_BandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidthWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidthWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_OverflowWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_OverflowWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_OverflowWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_OverflowWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_UnderflowWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_UnderflowWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_UnderflowWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_UnderflowWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_CountersWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_CountersWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributesWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributesWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPathWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPathWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroupsWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroupsWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroupsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroupsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_AdminGroups, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPathWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPathWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_CandidateSecondaryPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraintWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraintWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPath_PathMetricBoundConstraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPathWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPathWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroupsWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroupsWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroupsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroupsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_AdminGroups, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint is a *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint) SetVal(v *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint) *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint samples.
type CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint struct {
	W    *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraintWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraintWatcher observes a stream of *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint samples.
type NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPath_PathMetricBoundConstraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp is a *NetworkInstance_Mpls_Lsps_StaticLsp with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_StaticLsp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_StaticLsp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_StaticLsp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_StaticLsp sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp) SetVal(v *NetworkInstance_Mpls_Lsps_StaticLsp) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_StaticLsp samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLspWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLspWatcher observes a stream of *NetworkInstance_Mpls_Lsps_StaticLsp samples.
type NetworkInstance_Mpls_Lsps_StaticLspWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLspWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress is a *NetworkInstance_Mpls_Lsps_StaticLsp_Egress with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_StaticLsp_Egress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_StaticLsp_Egress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_StaticLsp_Egress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_StaticLsp_Egress sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress) SetVal(v *NetworkInstance_Mpls_Lsps_StaticLsp_Egress) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_StaticLsp_Egress samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_EgressWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_EgressWatcher observes a stream of *NetworkInstance_Mpls_Lsps_StaticLsp_Egress samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_EgressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_EgressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress is a *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress) SetVal(v *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_IngressWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_IngressWatcher observes a stream of *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_IngressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_IngressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit is a *NetworkInstance_Mpls_Lsps_StaticLsp_Transit with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_StaticLsp_Transit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_StaticLsp_Transit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_StaticLsp_Transit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_StaticLsp_Transit sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit) SetVal(v *NetworkInstance_Mpls_Lsps_StaticLsp_Transit) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_StaticLsp_Transit samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_TransitWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_TransitWatcher observes a stream of *NetworkInstance_Mpls_Lsps_StaticLsp_Transit samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_TransitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_TransitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath is a *NetworkInstance_Mpls_Lsps_UnconstrainedPath with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_UnconstrainedPath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_UnconstrainedPath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_UnconstrainedPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_UnconstrainedPath sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath) SetVal(v *NetworkInstance_Mpls_Lsps_UnconstrainedPath) *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_UnconstrainedPath samples.
type CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath struct {
	W    *NetworkInstance_Mpls_Lsps_UnconstrainedPathWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_UnconstrainedPathWatcher observes a stream of *NetworkInstance_Mpls_Lsps_UnconstrainedPath samples.
type NetworkInstance_Mpls_Lsps_UnconstrainedPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_UnconstrainedPathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol is a *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) SetVal(v *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol samples.
type CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol struct {
	W    *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocolWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocolWatcher observes a stream of *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol samples.
type NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocolWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocolWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp is a *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) Val(t testing.TB) *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) SetVal(v *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp samples.
type CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp struct {
	W    *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_LdpWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_LdpWatcher observes a stream of *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp samples.
type NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_LdpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_LdpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols is a *NetworkInstance_Mpls_SignalingProtocols with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols) SetVal(v *NetworkInstance_Mpls_SignalingProtocols) *QualifiedNetworkInstance_Mpls_SignalingProtocols {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols struct {
	W    *NetworkInstance_Mpls_SignalingProtocolsWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocolsWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols samples.
type NetworkInstance_Mpls_SignalingProtocolsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocolsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp is a *NetworkInstance_Mpls_SignalingProtocols_Ldp with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_LdpWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_LdpWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp samples.
type NetworkInstance_Mpls_SignalingProtocols_LdpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_LdpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_GlobalWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_GlobalWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_GlobalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_GlobalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_AuthenticationWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_AuthenticationWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_AuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_AuthenticationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_Authentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestartWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Global_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributesWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributesWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_InterfaceWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamilyWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamilyWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamilyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamilyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_AddressFamily, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_CountersWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_CountersWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_InterfaceAttributes_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_NeighborWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_NeighborWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_AuthenticationWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_AuthenticationWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_AuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_AuthenticationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_Authentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacencyWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacencyWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacencyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacencyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtimeWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtimeWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtimeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtimeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_HelloHoldtime, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRefWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Neighbor_HelloAdjacency_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_TargetedWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_TargetedWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_TargetedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_TargetedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamilyWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamilyWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamilyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamilyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target is a *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target) *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_TargetWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_TargetWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target samples.
type NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_TargetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_TargetWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_Ldp_Targeted_AddressFamily_Target, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTeWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTeWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_GlobalWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_GlobalWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_GlobalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_GlobalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_CountersWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_CountersWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_ErrorsWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_ErrorsWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_ErrorsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_ErrorsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Counters_Errors, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestartWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_HellosWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_HellosWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_HellosWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_HellosWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_Hellos, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemptionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemptionWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemptionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemptionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_InterfaceWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_AuthenticationWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_AuthenticationWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_AuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_AuthenticationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Authentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservationWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservationWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_CountersWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_CountersWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_ErrorsWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_ErrorsWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_ErrorsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_ErrorsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Counters_Errors, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_HellosWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_HellosWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_HellosWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_HellosWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Hellos, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_ProtectionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_ProtectionWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_ProtectionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_ProtectionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Protection, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_SubscriptionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_SubscriptionWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_SubscriptionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_SubscriptionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_Subscription, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_NeighborWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_NeighborWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_SessionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_SessionWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_SessionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_SessionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObjectWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObjectWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObjectWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObjectWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObjectWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObjectWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObjectWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObjectWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec is a *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspecWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspecWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspecWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspecWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_SenderTspec, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting is a *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRoutingWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRoutingWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRoutingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRoutingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter is a *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounterWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounterWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounterWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounterWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface is a *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_InterfaceWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef is a *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter is a *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounterWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounterWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounterWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounterWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass is a *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass struct {
	*genutil.Metadata
	val     *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass) Val(t testing.TB) *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass) SetVal(v *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClassWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClassWatcher observes a stream of *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClassWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClassWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_ForwardingClass, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
