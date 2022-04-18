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

// QualifiedBgpPeer is a *BgpPeer with a corresponding timestamp.
type QualifiedBgpPeer struct {
	*genutil.Metadata
	val     *BgpPeer // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer sample, erroring out if not present.
func (q *QualifiedBgpPeer) Val(t testing.TB) *BgpPeer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer sample.
func (q *QualifiedBgpPeer) SetVal(v *BgpPeer) *QualifiedBgpPeer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer is a telemetry Collection whose Await method returns a slice of *BgpPeer samples.
type CollectionBgpPeer struct {
	W    *BgpPeerWatcher
	Data []*QualifiedBgpPeer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer) Await(t testing.TB) []*QualifiedBgpPeer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeerWatcher observes a stream of *BgpPeer samples.
type BgpPeerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeerWatcher) Await(t testing.TB) (*QualifiedBgpPeer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_Counters is a *BgpPeer_Counters with a corresponding timestamp.
type QualifiedBgpPeer_Counters struct {
	*genutil.Metadata
	val     *BgpPeer_Counters // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_Counters sample, erroring out if not present.
func (q *QualifiedBgpPeer_Counters) Val(t testing.TB) *BgpPeer_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_Counters sample.
func (q *QualifiedBgpPeer_Counters) SetVal(v *BgpPeer_Counters) *QualifiedBgpPeer_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_Counters is a telemetry Collection whose Await method returns a slice of *BgpPeer_Counters samples.
type CollectionBgpPeer_Counters struct {
	W    *BgpPeer_CountersWatcher
	Data []*QualifiedBgpPeer_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_Counters) Await(t testing.TB) []*QualifiedBgpPeer_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_CountersWatcher observes a stream of *BgpPeer_Counters samples.
type BgpPeer_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_CountersWatcher) Await(t testing.TB) (*QualifiedBgpPeer_Counters, bool) {
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

// QualifiedInterface_Ipv4Neighbor is a *Interface_Ipv4Neighbor with a corresponding timestamp.
type QualifiedInterface_Ipv4Neighbor struct {
	*genutil.Metadata
	val     *Interface_Ipv4Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ipv4Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Ipv4Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_Ipv4Neighbor) Val(t testing.TB) *Interface_Ipv4Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Ipv4Neighbor sample.
func (q *QualifiedInterface_Ipv4Neighbor) SetVal(v *Interface_Ipv4Neighbor) *QualifiedInterface_Ipv4Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ipv4Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ipv4Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_Ipv4Neighbor samples.
type CollectionInterface_Ipv4Neighbor struct {
	W    *Interface_Ipv4NeighborWatcher
	Data []*QualifiedInterface_Ipv4Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ipv4Neighbor) Await(t testing.TB) []*QualifiedInterface_Ipv4Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Ipv4NeighborWatcher observes a stream of *Interface_Ipv4Neighbor samples.
type Interface_Ipv4NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ipv4Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Ipv4NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_Ipv4Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Ipv6Neighbor is a *Interface_Ipv6Neighbor with a corresponding timestamp.
type QualifiedInterface_Ipv6Neighbor struct {
	*genutil.Metadata
	val     *Interface_Ipv6Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ipv6Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Ipv6Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_Ipv6Neighbor) Val(t testing.TB) *Interface_Ipv6Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Ipv6Neighbor sample.
func (q *QualifiedInterface_Ipv6Neighbor) SetVal(v *Interface_Ipv6Neighbor) *QualifiedInterface_Ipv6Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ipv6Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ipv6Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_Ipv6Neighbor samples.
type CollectionInterface_Ipv6Neighbor struct {
	W    *Interface_Ipv6NeighborWatcher
	Data []*QualifiedInterface_Ipv6Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ipv6Neighbor) Await(t testing.TB) []*QualifiedInterface_Ipv6Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Ipv6NeighborWatcher observes a stream of *Interface_Ipv6Neighbor samples.
type Interface_Ipv6NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ipv6Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Ipv6NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_Ipv6Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
