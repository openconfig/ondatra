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

// QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering is a *NetworkInstance_Protocol_Isis_Level_TrafficEngineering with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_TrafficEngineering // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_TrafficEngineering sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_TrafficEngineering {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_TrafficEngineering sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering) SetVal(v *NetworkInstance_Protocol_Isis_Level_TrafficEngineering) *QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_TrafficEngineering is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_TrafficEngineering samples.
type CollectionNetworkInstance_Protocol_Isis_Level_TrafficEngineering struct {
	W    *NetworkInstance_Protocol_Isis_Level_TrafficEngineeringWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_TrafficEngineering) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_TrafficEngineeringWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_TrafficEngineering samples.
type NetworkInstance_Protocol_Isis_Level_TrafficEngineeringWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_TrafficEngineeringWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_TrafficEngineering, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2 is a *NetworkInstance_Protocol_Ospfv2 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2 sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2) SetVal(v *NetworkInstance_Protocol_Ospfv2) *QualifiedNetworkInstance_Protocol_Ospfv2 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2 samples.
type CollectionNetworkInstance_Protocol_Ospfv2 struct {
	W    *NetworkInstance_Protocol_Ospfv2Watcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2Watcher observes a stream of *NetworkInstance_Protocol_Ospfv2 samples.
type NetworkInstance_Protocol_Ospfv2Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area is a *NetworkInstance_Protocol_Ospfv2_Area with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area) *QualifiedNetworkInstance_Protocol_Ospfv2_Area {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area struct {
	W    *NetworkInstance_Protocol_Ospfv2_AreaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_AreaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area samples.
type NetworkInstance_Protocol_Ospfv2_AreaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_AreaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface is a *NetworkInstance_Protocol_Ospfv2_Area_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_InterfaceWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface samples.
type NetworkInstance_Protocol_Ospfv2_Area_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd is a *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfdWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd samples.
type NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_EnableBfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef is a *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef samples.
type NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter is a *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilterWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilterWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter samples.
type NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilterWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilterWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_LsaFilter, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls is a *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Interface_MplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Interface_MplsWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls samples.
type NetworkInstance_Protocol_Ospfv2_Area_Interface_MplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Interface_MplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync is a *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSyncWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSyncWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync samples.
type NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSyncWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSyncWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Mpls_IgpLdpSync, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor is a *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Interface_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Interface_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor samples.
type NetworkInstance_Protocol_Ospfv2_Area_Interface_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Interface_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers is a *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Interface_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Interface_TimersWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Interface_Timers samples.
type NetworkInstance_Protocol_Ospfv2_Area_Interface_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Interface_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Interface_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_LsdbWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_LsdbWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb samples.
type NetworkInstance_Protocol_Ospfv2_Area_LsdbWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_LsdbWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaTypeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaTypeWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaTypeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_LsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_LsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_LsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_LsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfServiceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfServiceWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfServiceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfServiceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_AsExternalLsa_TypeOfService, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NetworkLsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfServiceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfServiceWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfServiceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfServiceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_NssaExternalLsa_TypeOfService, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLinkWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLinkWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLinkWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLinkWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_TlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_AdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_Tlv_UnknownTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefixWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_TlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRangeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRangeWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRangeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRangeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_ExtendedPrefixRange, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSidWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_PrefixSid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBindingWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBindingWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBindingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBindingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_TlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetricWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPathWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPathWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_SegmentWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_SegmentWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_SegmentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_SegmentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4SegmentWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4SegmentWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4SegmentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4SegmentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_Ipv4Segment, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHopWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHopWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_EroPath_Segment_UnnumberedHop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBindingWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBindingWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBindingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBindingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_SidLabelBinding_Tlv_SidLabelBinding, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedPrefix_Tlv_UnknownTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_TlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_GraceLsa_Tlv_UnknownTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformationWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_TlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilitiesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilitiesWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilitiesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilitiesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_InformationalCapabilities, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTagsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTagsWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTagsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTagsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_NodeAdministrativeTags, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithmWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithmWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithmWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithmWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingAlgorithm, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRangeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRangeWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRangeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRangeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_TlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabelWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabelWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabelWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabelWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_SidLabel, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_UnknownTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_UnknownTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineeringWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineeringWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineeringWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineeringWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_TlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_LinkWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_LinkWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_LinkWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_LinkWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_AdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnknownSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_UnreservedBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttributeWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_UnknownSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddressWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_RouterAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_UnknownTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlvWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_UnknownTlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfServiceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfServiceWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfServiceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfServiceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_TypeOfService, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsaWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsaWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsaWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService is a *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfServiceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfServiceWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfServiceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfServiceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_SummaryLsa_TypeOfService, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls is a *NetworkInstance_Protocol_Ospfv2_Area_Mpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_Mpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_Mpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_Mpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_Mpls sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_Mpls) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Mpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_Mpls samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Mpls struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_MplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Mpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_MplsWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_Mpls samples.
type NetworkInstance_Protocol_Ospfv2_Area_MplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_MplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Mpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink is a *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink) SetVal(v *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_VirtualLink is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_VirtualLink struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_VirtualLinkWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_VirtualLink) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_VirtualLinkWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Area_VirtualLink samples.
type NetworkInstance_Protocol_Ospfv2_Area_VirtualLinkWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_VirtualLinkWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_VirtualLink, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global is a *NetworkInstance_Protocol_Ospfv2_Global with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global) *QualifiedNetworkInstance_Protocol_Ospfv2_Global {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global struct {
	W    *NetworkInstance_Protocol_Ospfv2_GlobalWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_GlobalWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global samples.
type NetworkInstance_Protocol_Ospfv2_GlobalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_GlobalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart is a *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestart samples.
type NetworkInstance_Protocol_Ospfv2_Global_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy is a *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicyWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy samples.
type NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_InterAreaPropagationPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls is a *NetworkInstance_Protocol_Ospfv2_Global_Mpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_Mpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_Mpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_Mpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_Mpls sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_Mpls) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_Mpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_Mpls samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_Mpls struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_MplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_Mpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_MplsWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_Mpls samples.
type NetworkInstance_Protocol_Ospfv2_Global_MplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_MplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync is a *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSyncWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSyncWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync samples.
type NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSyncWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSyncWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Mpls_IgpLdpSync, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers is a *NetworkInstance_Protocol_Ospfv2_Global_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_Timers {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_Timers) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_Timers samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_TimersWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_Timers samples.
type NetworkInstance_Protocol_Ospfv2_Global_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration is a *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGenerationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGenerationWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration samples.
type NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGenerationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGenerationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_LsaGeneration, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric is a *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetricWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric samples.
type NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_MaxMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf is a *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf) Val(t testing.TB) *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf) SetVal(v *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf) *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf struct {
	W    *NetworkInstance_Protocol_Ospfv2_Global_Timers_SpfWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Global_Timers_SpfWatcher observes a stream of *NetworkInstance_Protocol_Ospfv2_Global_Timers_Spf samples.
type NetworkInstance_Protocol_Ospfv2_Global_Timers_SpfWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Global_Timers_SpfWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Global_Timers_Spf, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim is a *NetworkInstance_Protocol_Pim with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim) Val(t testing.TB) *NetworkInstance_Protocol_Pim {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim sample.
func (q *QualifiedNetworkInstance_Protocol_Pim) SetVal(v *NetworkInstance_Protocol_Pim) *QualifiedNetworkInstance_Protocol_Pim {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim samples.
type CollectionNetworkInstance_Protocol_Pim struct {
	W    *NetworkInstance_Protocol_PimWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_PimWatcher observes a stream of *NetworkInstance_Protocol_Pim samples.
type NetworkInstance_Protocol_PimWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_PimWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Global is a *NetworkInstance_Protocol_Pim_Global with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Global struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Global // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Global) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Global sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Global {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Global sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global) SetVal(v *NetworkInstance_Protocol_Pim_Global) *QualifiedNetworkInstance_Protocol_Pim_Global {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Global is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Global samples.
type CollectionNetworkInstance_Protocol_Pim_Global struct {
	W    *NetworkInstance_Protocol_Pim_GlobalWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Global
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Global) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Global {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_GlobalWatcher observes a stream of *NetworkInstance_Protocol_Pim_Global samples.
type NetworkInstance_Protocol_Pim_GlobalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Global
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_GlobalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Global, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Global_Counters is a *NetworkInstance_Protocol_Pim_Global_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Global_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Global_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Global_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Counters) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Global_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Global_Counters sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Counters) SetVal(v *NetworkInstance_Protocol_Pim_Global_Counters) *QualifiedNetworkInstance_Protocol_Pim_Global_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Global_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Global_Counters samples.
type CollectionNetworkInstance_Protocol_Pim_Global_Counters struct {
	W    *NetworkInstance_Protocol_Pim_Global_CountersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Global_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Global_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Global_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Global_CountersWatcher observes a stream of *NetworkInstance_Protocol_Pim_Global_Counters samples.
type NetworkInstance_Protocol_Pim_Global_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Global_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Global_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Global_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint is a *NetworkInstance_Protocol_Pim_Global_RendezvousPoint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Global_RendezvousPoint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Global_RendezvousPoint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Global_RendezvousPoint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Global_RendezvousPoint sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint) SetVal(v *NetworkInstance_Protocol_Pim_Global_RendezvousPoint) *QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Global_RendezvousPoint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Global_RendezvousPoint samples.
type CollectionNetworkInstance_Protocol_Pim_Global_RendezvousPoint struct {
	W    *NetworkInstance_Protocol_Pim_Global_RendezvousPointWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Global_RendezvousPoint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Global_RendezvousPointWatcher observes a stream of *NetworkInstance_Protocol_Pim_Global_RendezvousPoint samples.
type NetworkInstance_Protocol_Pim_Global_RendezvousPointWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Global_RendezvousPointWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Global_RendezvousPoint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Global_Source is a *NetworkInstance_Protocol_Pim_Global_Source with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Global_Source struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Global_Source // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Source) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Global_Source sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Source) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Global_Source {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Global_Source sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Source) SetVal(v *NetworkInstance_Protocol_Pim_Global_Source) *QualifiedNetworkInstance_Protocol_Pim_Global_Source {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Source) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Global_Source is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Global_Source samples.
type CollectionNetworkInstance_Protocol_Pim_Global_Source struct {
	W    *NetworkInstance_Protocol_Pim_Global_SourceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Global_Source
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Global_Source) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Global_Source {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Global_SourceWatcher observes a stream of *NetworkInstance_Protocol_Pim_Global_Source samples.
type NetworkInstance_Protocol_Pim_Global_SourceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Global_Source
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Global_SourceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Global_Source, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Global_Ssm is a *NetworkInstance_Protocol_Pim_Global_Ssm with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Global_Ssm struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Global_Ssm // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Ssm) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Global_Ssm sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Ssm) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Global_Ssm {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Global_Ssm sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Ssm) SetVal(v *NetworkInstance_Protocol_Pim_Global_Ssm) *QualifiedNetworkInstance_Protocol_Pim_Global_Ssm {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Global_Ssm) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Global_Ssm is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Global_Ssm samples.
type CollectionNetworkInstance_Protocol_Pim_Global_Ssm struct {
	W    *NetworkInstance_Protocol_Pim_Global_SsmWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Global_Ssm
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Global_Ssm) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Global_Ssm {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Global_SsmWatcher observes a stream of *NetworkInstance_Protocol_Pim_Global_Ssm samples.
type NetworkInstance_Protocol_Pim_Global_SsmWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Global_Ssm
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Global_SsmWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Global_Ssm, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Interface is a *NetworkInstance_Protocol_Pim_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Interface sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface) SetVal(v *NetworkInstance_Protocol_Pim_Interface) *QualifiedNetworkInstance_Protocol_Pim_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Interface samples.
type CollectionNetworkInstance_Protocol_Pim_Interface struct {
	W    *NetworkInstance_Protocol_Pim_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_InterfaceWatcher observes a stream of *NetworkInstance_Protocol_Pim_Interface samples.
type NetworkInstance_Protocol_Pim_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Interface_Counters is a *NetworkInstance_Protocol_Pim_Interface_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Interface_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Interface_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Interface_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Counters) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Interface_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Interface_Counters sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Counters) SetVal(v *NetworkInstance_Protocol_Pim_Interface_Counters) *QualifiedNetworkInstance_Protocol_Pim_Interface_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Interface_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Interface_Counters samples.
type CollectionNetworkInstance_Protocol_Pim_Interface_Counters struct {
	W    *NetworkInstance_Protocol_Pim_Interface_CountersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Interface_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Interface_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Interface_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Interface_CountersWatcher observes a stream of *NetworkInstance_Protocol_Pim_Interface_Counters samples.
type NetworkInstance_Protocol_Pim_Interface_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Interface_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Interface_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Interface_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd is a *NetworkInstance_Protocol_Pim_Interface_EnableBfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Interface_EnableBfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Interface_EnableBfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Interface_EnableBfd {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Interface_EnableBfd sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd) SetVal(v *NetworkInstance_Protocol_Pim_Interface_EnableBfd) *QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Interface_EnableBfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Interface_EnableBfd samples.
type CollectionNetworkInstance_Protocol_Pim_Interface_EnableBfd struct {
	W    *NetworkInstance_Protocol_Pim_Interface_EnableBfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Interface_EnableBfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Interface_EnableBfdWatcher observes a stream of *NetworkInstance_Protocol_Pim_Interface_EnableBfd samples.
type NetworkInstance_Protocol_Pim_Interface_EnableBfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Interface_EnableBfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Interface_EnableBfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef is a *NetworkInstance_Protocol_Pim_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef) SetVal(v *NetworkInstance_Protocol_Pim_Interface_InterfaceRef) *QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Protocol_Pim_Interface_InterfaceRef struct {
	W    *NetworkInstance_Protocol_Pim_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Protocol_Pim_Interface_InterfaceRef samples.
type NetworkInstance_Protocol_Pim_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor is a *NetworkInstance_Protocol_Pim_Interface_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Pim_Interface_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Pim_Interface_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Pim_Interface_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Pim_Interface_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor) SetVal(v *NetworkInstance_Protocol_Pim_Interface_Neighbor) *QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Pim_Interface_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Pim_Interface_Neighbor samples.
type CollectionNetworkInstance_Protocol_Pim_Interface_Neighbor struct {
	W    *NetworkInstance_Protocol_Pim_Interface_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Pim_Interface_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Pim_Interface_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Pim_Interface_Neighbor samples.
type NetworkInstance_Protocol_Pim_Interface_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Pim_Interface_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Pim_Interface_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Static is a *NetworkInstance_Protocol_Static with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Static struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Static // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Static) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Static sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Static) Val(t testing.TB) *NetworkInstance_Protocol_Static {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Static sample.
func (q *QualifiedNetworkInstance_Protocol_Static) SetVal(v *NetworkInstance_Protocol_Static) *QualifiedNetworkInstance_Protocol_Static {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Static) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Static is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Static samples.
type CollectionNetworkInstance_Protocol_Static struct {
	W    *NetworkInstance_Protocol_StaticWatcher
	Data []*QualifiedNetworkInstance_Protocol_Static
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Static) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Static {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_StaticWatcher observes a stream of *NetworkInstance_Protocol_Static samples.
type NetworkInstance_Protocol_StaticWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Static
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_StaticWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Static, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Static_NextHop is a *NetworkInstance_Protocol_Static_NextHop with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Static_NextHop struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Static_NextHop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Static_NextHop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Static_NextHop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop) Val(t testing.TB) *NetworkInstance_Protocol_Static_NextHop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Static_NextHop sample.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop) SetVal(v *NetworkInstance_Protocol_Static_NextHop) *QualifiedNetworkInstance_Protocol_Static_NextHop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Static_NextHop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Static_NextHop samples.
type CollectionNetworkInstance_Protocol_Static_NextHop struct {
	W    *NetworkInstance_Protocol_Static_NextHopWatcher
	Data []*QualifiedNetworkInstance_Protocol_Static_NextHop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Static_NextHop) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Static_NextHop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Static_NextHopWatcher observes a stream of *NetworkInstance_Protocol_Static_NextHop samples.
type NetworkInstance_Protocol_Static_NextHopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Static_NextHop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Static_NextHopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Static_NextHop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd is a *NetworkInstance_Protocol_Static_NextHop_EnableBfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Static_NextHop_EnableBfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Static_NextHop_EnableBfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd) Val(t testing.TB) *NetworkInstance_Protocol_Static_NextHop_EnableBfd {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Static_NextHop_EnableBfd sample.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd) SetVal(v *NetworkInstance_Protocol_Static_NextHop_EnableBfd) *QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Static_NextHop_EnableBfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Static_NextHop_EnableBfd samples.
type CollectionNetworkInstance_Protocol_Static_NextHop_EnableBfd struct {
	W    *NetworkInstance_Protocol_Static_NextHop_EnableBfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Static_NextHop_EnableBfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Static_NextHop_EnableBfdWatcher observes a stream of *NetworkInstance_Protocol_Static_NextHop_EnableBfd samples.
type NetworkInstance_Protocol_Static_NextHop_EnableBfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Static_NextHop_EnableBfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Static_NextHop_EnableBfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef is a *NetworkInstance_Protocol_Static_NextHop_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Static_NextHop_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Static_NextHop_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef) Val(t testing.TB) *NetworkInstance_Protocol_Static_NextHop_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Static_NextHop_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef) SetVal(v *NetworkInstance_Protocol_Static_NextHop_InterfaceRef) *QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Static_NextHop_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Static_NextHop_InterfaceRef samples.
type CollectionNetworkInstance_Protocol_Static_NextHop_InterfaceRef struct {
	W    *NetworkInstance_Protocol_Static_NextHop_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Static_NextHop_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Static_NextHop_InterfaceRefWatcher observes a stream of *NetworkInstance_Protocol_Static_NextHop_InterfaceRef samples.
type NetworkInstance_Protocol_Static_NextHop_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Static_NextHop_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Static_NextHop_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_RouteLimit is a *NetworkInstance_RouteLimit with a corresponding timestamp.
type QualifiedNetworkInstance_RouteLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_RouteLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_RouteLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_RouteLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_RouteLimit) Val(t testing.TB) *NetworkInstance_RouteLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_RouteLimit sample.
func (q *QualifiedNetworkInstance_RouteLimit) SetVal(v *NetworkInstance_RouteLimit) *QualifiedNetworkInstance_RouteLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_RouteLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_RouteLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_RouteLimit samples.
type CollectionNetworkInstance_RouteLimit struct {
	W    *NetworkInstance_RouteLimitWatcher
	Data []*QualifiedNetworkInstance_RouteLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_RouteLimit) Await(t testing.TB) []*QualifiedNetworkInstance_RouteLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_RouteLimitWatcher observes a stream of *NetworkInstance_RouteLimit samples.
type NetworkInstance_RouteLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_RouteLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_RouteLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_RouteLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting is a *NetworkInstance_SegmentRouting with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting) Val(t testing.TB) *NetworkInstance_SegmentRouting {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting sample.
func (q *QualifiedNetworkInstance_SegmentRouting) SetVal(v *NetworkInstance_SegmentRouting) *QualifiedNetworkInstance_SegmentRouting {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting samples.
type CollectionNetworkInstance_SegmentRouting struct {
	W    *NetworkInstance_SegmentRoutingWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRoutingWatcher observes a stream of *NetworkInstance_SegmentRouting samples.
type NetworkInstance_SegmentRoutingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRoutingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_Srgb is a *NetworkInstance_SegmentRouting_Srgb with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_Srgb struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_Srgb // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_Srgb) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_Srgb sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_Srgb) Val(t testing.TB) *NetworkInstance_SegmentRouting_Srgb {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_Srgb sample.
func (q *QualifiedNetworkInstance_SegmentRouting_Srgb) SetVal(v *NetworkInstance_SegmentRouting_Srgb) *QualifiedNetworkInstance_SegmentRouting_Srgb {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_Srgb) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_Srgb is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_Srgb samples.
type CollectionNetworkInstance_SegmentRouting_Srgb struct {
	W    *NetworkInstance_SegmentRouting_SrgbWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_Srgb
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_Srgb) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_Srgb {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_SrgbWatcher observes a stream of *NetworkInstance_SegmentRouting_Srgb samples.
type NetworkInstance_SegmentRouting_SrgbWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_Srgb
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_SrgbWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_Srgb, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_Srlb is a *NetworkInstance_SegmentRouting_Srlb with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_Srlb struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_Srlb // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_Srlb) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_Srlb sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_Srlb) Val(t testing.TB) *NetworkInstance_SegmentRouting_Srlb {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_Srlb sample.
func (q *QualifiedNetworkInstance_SegmentRouting_Srlb) SetVal(v *NetworkInstance_SegmentRouting_Srlb) *QualifiedNetworkInstance_SegmentRouting_Srlb {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_Srlb) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_Srlb is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_Srlb samples.
type CollectionNetworkInstance_SegmentRouting_Srlb struct {
	W    *NetworkInstance_SegmentRouting_SrlbWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_Srlb
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_Srlb) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_Srlb {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_SrlbWatcher observes a stream of *NetworkInstance_SegmentRouting_Srlb samples.
type NetworkInstance_SegmentRouting_SrlbWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_Srlb
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_SrlbWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_Srlb, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy is a *NetworkInstance_SegmentRouting_TePolicy with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy) SetVal(v *NetworkInstance_SegmentRouting_TePolicy) *QualifiedNetworkInstance_SegmentRouting_TePolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy struct {
	W    *NetworkInstance_SegmentRouting_TePolicyWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicyWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy samples.
type NetworkInstance_SegmentRouting_TePolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath is a *NetworkInstance_SegmentRouting_TePolicy_CandidatePath with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_CandidatePath // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_CandidatePath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_CandidatePath) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePathWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePathWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList is a *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentListWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentListWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentListWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentListWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters is a *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_CountersWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_CountersWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop is a *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHopWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHopWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters is a *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_CountersWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_CountersWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef is a *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRefWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid is a *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_SidWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_SidWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_SidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_SidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters is a *NetworkInstance_SegmentRouting_TePolicy_Counters with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters struct {
	*genutil.Metadata
	val     *NetworkInstance_SegmentRouting_TePolicy_Counters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_SegmentRouting_TePolicy_Counters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters) Val(t testing.TB) *NetworkInstance_SegmentRouting_TePolicy_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_SegmentRouting_TePolicy_Counters sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters) SetVal(v *NetworkInstance_SegmentRouting_TePolicy_Counters) *QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_Counters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_SegmentRouting_TePolicy_Counters samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_Counters struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CountersWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_Counters) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CountersWatcher observes a stream of *NetworkInstance_SegmentRouting_TePolicy_Counters samples.
type NetworkInstance_SegmentRouting_TePolicy_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Table is a *NetworkInstance_Table with a corresponding timestamp.
type QualifiedNetworkInstance_Table struct {
	*genutil.Metadata
	val     *NetworkInstance_Table // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Table) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Table sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Table) Val(t testing.TB) *NetworkInstance_Table {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Table sample.
func (q *QualifiedNetworkInstance_Table) SetVal(v *NetworkInstance_Table) *QualifiedNetworkInstance_Table {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Table) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Table is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Table samples.
type CollectionNetworkInstance_Table struct {
	W    *NetworkInstance_TableWatcher
	Data []*QualifiedNetworkInstance_Table
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Table) Await(t testing.TB) []*QualifiedNetworkInstance_Table {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_TableWatcher observes a stream of *NetworkInstance_Table samples.
type NetworkInstance_TableWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Table
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_TableWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Table, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_TableConnection is a *NetworkInstance_TableConnection with a corresponding timestamp.
type QualifiedNetworkInstance_TableConnection struct {
	*genutil.Metadata
	val     *NetworkInstance_TableConnection // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_TableConnection) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_TableConnection sample, erroring out if not present.
func (q *QualifiedNetworkInstance_TableConnection) Val(t testing.TB) *NetworkInstance_TableConnection {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_TableConnection sample.
func (q *QualifiedNetworkInstance_TableConnection) SetVal(v *NetworkInstance_TableConnection) *QualifiedNetworkInstance_TableConnection {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_TableConnection) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_TableConnection is a telemetry Collection whose Await method returns a slice of *NetworkInstance_TableConnection samples.
type CollectionNetworkInstance_TableConnection struct {
	W    *NetworkInstance_TableConnectionWatcher
	Data []*QualifiedNetworkInstance_TableConnection
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_TableConnection) Await(t testing.TB) []*QualifiedNetworkInstance_TableConnection {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_TableConnectionWatcher observes a stream of *NetworkInstance_TableConnection samples.
type NetworkInstance_TableConnectionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_TableConnection
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_TableConnectionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_TableConnection, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Vlan is a *NetworkInstance_Vlan with a corresponding timestamp.
type QualifiedNetworkInstance_Vlan struct {
	*genutil.Metadata
	val     *NetworkInstance_Vlan // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Vlan) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Vlan sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Vlan) Val(t testing.TB) *NetworkInstance_Vlan {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Vlan sample.
func (q *QualifiedNetworkInstance_Vlan) SetVal(v *NetworkInstance_Vlan) *QualifiedNetworkInstance_Vlan {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Vlan) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Vlan is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Vlan samples.
type CollectionNetworkInstance_Vlan struct {
	W    *NetworkInstance_VlanWatcher
	Data []*QualifiedNetworkInstance_Vlan
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Vlan) Await(t testing.TB) []*QualifiedNetworkInstance_Vlan {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_VlanWatcher observes a stream of *NetworkInstance_Vlan samples.
type NetworkInstance_VlanWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Vlan
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_VlanWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Vlan, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Vlan_Member is a *NetworkInstance_Vlan_Member with a corresponding timestamp.
type QualifiedNetworkInstance_Vlan_Member struct {
	*genutil.Metadata
	val     *NetworkInstance_Vlan_Member // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Vlan_Member) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Vlan_Member sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Vlan_Member) Val(t testing.TB) *NetworkInstance_Vlan_Member {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Vlan_Member sample.
func (q *QualifiedNetworkInstance_Vlan_Member) SetVal(v *NetworkInstance_Vlan_Member) *QualifiedNetworkInstance_Vlan_Member {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Vlan_Member) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Vlan_Member is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Vlan_Member samples.
type CollectionNetworkInstance_Vlan_Member struct {
	W    *NetworkInstance_Vlan_MemberWatcher
	Data []*QualifiedNetworkInstance_Vlan_Member
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Vlan_Member) Await(t testing.TB) []*QualifiedNetworkInstance_Vlan_Member {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Vlan_MemberWatcher observes a stream of *NetworkInstance_Vlan_Member samples.
type NetworkInstance_Vlan_MemberWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Vlan_Member
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Vlan_MemberWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Vlan_Member, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos is a *Qos with a corresponding timestamp.
type QualifiedQos struct {
	*genutil.Metadata
	val     *Qos // val is the sample value.
	present bool
}

func (q *QualifiedQos) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos sample, erroring out if not present.
func (q *QualifiedQos) Val(t testing.TB) *Qos {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos sample.
func (q *QualifiedQos) SetVal(v *Qos) *QualifiedQos {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos is a telemetry Collection whose Await method returns a slice of *Qos samples.
type CollectionQos struct {
	W    *QosWatcher
	Data []*QualifiedQos
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos) Await(t testing.TB) []*QualifiedQos {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// QosWatcher observes a stream of *Qos samples.
type QosWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *QosWatcher) Await(t testing.TB) (*QualifiedQos, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_BufferAllocationProfile is a *Qos_BufferAllocationProfile with a corresponding timestamp.
type QualifiedQos_BufferAllocationProfile struct {
	*genutil.Metadata
	val     *Qos_BufferAllocationProfile // val is the sample value.
	present bool
}

func (q *QualifiedQos_BufferAllocationProfile) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_BufferAllocationProfile sample, erroring out if not present.
func (q *QualifiedQos_BufferAllocationProfile) Val(t testing.TB) *Qos_BufferAllocationProfile {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_BufferAllocationProfile sample.
func (q *QualifiedQos_BufferAllocationProfile) SetVal(v *Qos_BufferAllocationProfile) *QualifiedQos_BufferAllocationProfile {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_BufferAllocationProfile) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_BufferAllocationProfile is a telemetry Collection whose Await method returns a slice of *Qos_BufferAllocationProfile samples.
type CollectionQos_BufferAllocationProfile struct {
	W    *Qos_BufferAllocationProfileWatcher
	Data []*QualifiedQos_BufferAllocationProfile
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_BufferAllocationProfile) Await(t testing.TB) []*QualifiedQos_BufferAllocationProfile {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_BufferAllocationProfileWatcher observes a stream of *Qos_BufferAllocationProfile samples.
type Qos_BufferAllocationProfileWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_BufferAllocationProfile
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_BufferAllocationProfileWatcher) Await(t testing.TB) (*QualifiedQos_BufferAllocationProfile, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_BufferAllocationProfile_Queue is a *Qos_BufferAllocationProfile_Queue with a corresponding timestamp.
type QualifiedQos_BufferAllocationProfile_Queue struct {
	*genutil.Metadata
	val     *Qos_BufferAllocationProfile_Queue // val is the sample value.
	present bool
}

func (q *QualifiedQos_BufferAllocationProfile_Queue) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_BufferAllocationProfile_Queue sample, erroring out if not present.
func (q *QualifiedQos_BufferAllocationProfile_Queue) Val(t testing.TB) *Qos_BufferAllocationProfile_Queue {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_BufferAllocationProfile_Queue sample.
func (q *QualifiedQos_BufferAllocationProfile_Queue) SetVal(v *Qos_BufferAllocationProfile_Queue) *QualifiedQos_BufferAllocationProfile_Queue {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_BufferAllocationProfile_Queue) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_BufferAllocationProfile_Queue is a telemetry Collection whose Await method returns a slice of *Qos_BufferAllocationProfile_Queue samples.
type CollectionQos_BufferAllocationProfile_Queue struct {
	W    *Qos_BufferAllocationProfile_QueueWatcher
	Data []*QualifiedQos_BufferAllocationProfile_Queue
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_BufferAllocationProfile_Queue) Await(t testing.TB) []*QualifiedQos_BufferAllocationProfile_Queue {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_BufferAllocationProfile_QueueWatcher observes a stream of *Qos_BufferAllocationProfile_Queue samples.
type Qos_BufferAllocationProfile_QueueWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_BufferAllocationProfile_Queue
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_BufferAllocationProfile_QueueWatcher) Await(t testing.TB) (*QualifiedQos_BufferAllocationProfile_Queue, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier is a *Qos_Classifier with a corresponding timestamp.
type QualifiedQos_Classifier struct {
	*genutil.Metadata
	val     *Qos_Classifier // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier sample, erroring out if not present.
func (q *QualifiedQos_Classifier) Val(t testing.TB) *Qos_Classifier {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier sample.
func (q *QualifiedQos_Classifier) SetVal(v *Qos_Classifier) *QualifiedQos_Classifier {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier is a telemetry Collection whose Await method returns a slice of *Qos_Classifier samples.
type CollectionQos_Classifier struct {
	W    *Qos_ClassifierWatcher
	Data []*QualifiedQos_Classifier
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier) Await(t testing.TB) []*QualifiedQos_Classifier {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_ClassifierWatcher observes a stream of *Qos_Classifier samples.
type Qos_ClassifierWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_ClassifierWatcher) Await(t testing.TB) (*QualifiedQos_Classifier, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term is a *Qos_Classifier_Term with a corresponding timestamp.
type QualifiedQos_Classifier_Term struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term) Val(t testing.TB) *Qos_Classifier_Term {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term sample.
func (q *QualifiedQos_Classifier_Term) SetVal(v *Qos_Classifier_Term) *QualifiedQos_Classifier_Term {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term samples.
type CollectionQos_Classifier_Term struct {
	W    *Qos_Classifier_TermWatcher
	Data []*QualifiedQos_Classifier_Term
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term) Await(t testing.TB) []*QualifiedQos_Classifier_Term {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_TermWatcher observes a stream of *Qos_Classifier_Term samples.
type Qos_Classifier_TermWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_TermWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Actions is a *Qos_Classifier_Term_Actions with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Actions struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Actions // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Actions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Actions sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Actions) Val(t testing.TB) *Qos_Classifier_Term_Actions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Actions sample.
func (q *QualifiedQos_Classifier_Term_Actions) SetVal(v *Qos_Classifier_Term_Actions) *QualifiedQos_Classifier_Term_Actions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Actions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Actions is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Actions samples.
type CollectionQos_Classifier_Term_Actions struct {
	W    *Qos_Classifier_Term_ActionsWatcher
	Data []*QualifiedQos_Classifier_Term_Actions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Actions) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Actions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_ActionsWatcher observes a stream of *Qos_Classifier_Term_Actions samples.
type Qos_Classifier_Term_ActionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Actions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_ActionsWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Actions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Actions_Remark is a *Qos_Classifier_Term_Actions_Remark with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Actions_Remark struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Actions_Remark // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Actions_Remark) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Actions_Remark sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Actions_Remark) Val(t testing.TB) *Qos_Classifier_Term_Actions_Remark {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Actions_Remark sample.
func (q *QualifiedQos_Classifier_Term_Actions_Remark) SetVal(v *Qos_Classifier_Term_Actions_Remark) *QualifiedQos_Classifier_Term_Actions_Remark {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Actions_Remark) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Actions_Remark is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Actions_Remark samples.
type CollectionQos_Classifier_Term_Actions_Remark struct {
	W    *Qos_Classifier_Term_Actions_RemarkWatcher
	Data []*QualifiedQos_Classifier_Term_Actions_Remark
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Actions_Remark) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Actions_Remark {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Actions_RemarkWatcher observes a stream of *Qos_Classifier_Term_Actions_Remark samples.
type Qos_Classifier_Term_Actions_RemarkWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Actions_Remark
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Actions_RemarkWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Actions_Remark, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions is a *Qos_Classifier_Term_Conditions with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Conditions // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Conditions sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions) Val(t testing.TB) *Qos_Classifier_Term_Conditions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Conditions sample.
func (q *QualifiedQos_Classifier_Term_Conditions) SetVal(v *Qos_Classifier_Term_Conditions) *QualifiedQos_Classifier_Term_Conditions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Conditions samples.
type CollectionQos_Classifier_Term_Conditions struct {
	W    *Qos_Classifier_Term_ConditionsWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_ConditionsWatcher observes a stream of *Qos_Classifier_Term_Conditions samples.
type Qos_Classifier_Term_ConditionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_ConditionsWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Ipv4 is a *Qos_Classifier_Term_Conditions_Ipv4 with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Ipv4 struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Conditions_Ipv4 // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Conditions_Ipv4 sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4) Val(t testing.TB) *Qos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Conditions_Ipv4 sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4) SetVal(v *Qos_Classifier_Term_Conditions_Ipv4) *QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Ipv4 is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Conditions_Ipv4 samples.
type CollectionQos_Classifier_Term_Conditions_Ipv4 struct {
	W    *Qos_Classifier_Term_Conditions_Ipv4Watcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Ipv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Ipv4) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Ipv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Ipv4Watcher observes a stream of *Qos_Classifier_Term_Conditions_Ipv4 samples.
type Qos_Classifier_Term_Conditions_Ipv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Ipv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Ipv4Watcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Ipv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Ipv6 is a *Qos_Classifier_Term_Conditions_Ipv6 with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Ipv6 struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Conditions_Ipv6 // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Conditions_Ipv6 sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6) Val(t testing.TB) *Qos_Classifier_Term_Conditions_Ipv6 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Conditions_Ipv6 sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6) SetVal(v *Qos_Classifier_Term_Conditions_Ipv6) *QualifiedQos_Classifier_Term_Conditions_Ipv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Ipv6 is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Conditions_Ipv6 samples.
type CollectionQos_Classifier_Term_Conditions_Ipv6 struct {
	W    *Qos_Classifier_Term_Conditions_Ipv6Watcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Ipv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Ipv6) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Ipv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Ipv6Watcher observes a stream of *Qos_Classifier_Term_Conditions_Ipv6 samples.
type Qos_Classifier_Term_Conditions_Ipv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Ipv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Ipv6Watcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Ipv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_L2 is a *Qos_Classifier_Term_Conditions_L2 with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_L2 struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Conditions_L2 // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_L2) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Conditions_L2 sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_L2) Val(t testing.TB) *Qos_Classifier_Term_Conditions_L2 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Conditions_L2 sample.
func (q *QualifiedQos_Classifier_Term_Conditions_L2) SetVal(v *Qos_Classifier_Term_Conditions_L2) *QualifiedQos_Classifier_Term_Conditions_L2 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_L2) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_L2 is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Conditions_L2 samples.
type CollectionQos_Classifier_Term_Conditions_L2 struct {
	W    *Qos_Classifier_Term_Conditions_L2Watcher
	Data []*QualifiedQos_Classifier_Term_Conditions_L2
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_L2) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_L2Watcher observes a stream of *Qos_Classifier_Term_Conditions_L2 samples.
type Qos_Classifier_Term_Conditions_L2Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_L2
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_L2Watcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_L2, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Mpls is a *Qos_Classifier_Term_Conditions_Mpls with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Mpls struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Conditions_Mpls // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Mpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Conditions_Mpls sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls) Val(t testing.TB) *Qos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Conditions_Mpls sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls) SetVal(v *Qos_Classifier_Term_Conditions_Mpls) *QualifiedQos_Classifier_Term_Conditions_Mpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Mpls is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Conditions_Mpls samples.
type CollectionQos_Classifier_Term_Conditions_Mpls struct {
	W    *Qos_Classifier_Term_Conditions_MplsWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Mpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Mpls) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_MplsWatcher observes a stream of *Qos_Classifier_Term_Conditions_Mpls samples.
type Qos_Classifier_Term_Conditions_MplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Mpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_MplsWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Mpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Transport is a *Qos_Classifier_Term_Conditions_Transport with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Transport struct {
	*genutil.Metadata
	val     *Qos_Classifier_Term_Conditions_Transport // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Transport) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Classifier_Term_Conditions_Transport sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport) Val(t testing.TB) *Qos_Classifier_Term_Conditions_Transport {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Classifier_Term_Conditions_Transport sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport) SetVal(v *Qos_Classifier_Term_Conditions_Transport) *QualifiedQos_Classifier_Term_Conditions_Transport {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Transport is a telemetry Collection whose Await method returns a slice of *Qos_Classifier_Term_Conditions_Transport samples.
type CollectionQos_Classifier_Term_Conditions_Transport struct {
	W    *Qos_Classifier_Term_Conditions_TransportWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Transport
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Transport) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_TransportWatcher observes a stream of *Qos_Classifier_Term_Conditions_Transport samples.
type Qos_Classifier_Term_Conditions_TransportWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Transport
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_TransportWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Transport, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_ForwardingGroup is a *Qos_ForwardingGroup with a corresponding timestamp.
type QualifiedQos_ForwardingGroup struct {
	*genutil.Metadata
	val     *Qos_ForwardingGroup // val is the sample value.
	present bool
}

func (q *QualifiedQos_ForwardingGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_ForwardingGroup sample, erroring out if not present.
func (q *QualifiedQos_ForwardingGroup) Val(t testing.TB) *Qos_ForwardingGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_ForwardingGroup sample.
func (q *QualifiedQos_ForwardingGroup) SetVal(v *Qos_ForwardingGroup) *QualifiedQos_ForwardingGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_ForwardingGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_ForwardingGroup is a telemetry Collection whose Await method returns a slice of *Qos_ForwardingGroup samples.
type CollectionQos_ForwardingGroup struct {
	W    *Qos_ForwardingGroupWatcher
	Data []*QualifiedQos_ForwardingGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_ForwardingGroup) Await(t testing.TB) []*QualifiedQos_ForwardingGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_ForwardingGroupWatcher observes a stream of *Qos_ForwardingGroup samples.
type Qos_ForwardingGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_ForwardingGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_ForwardingGroupWatcher) Await(t testing.TB) (*QualifiedQos_ForwardingGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface is a *Qos_Interface with a corresponding timestamp.
type QualifiedQos_Interface struct {
	*genutil.Metadata
	val     *Qos_Interface // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface sample, erroring out if not present.
func (q *QualifiedQos_Interface) Val(t testing.TB) *Qos_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface sample.
func (q *QualifiedQos_Interface) SetVal(v *Qos_Interface) *QualifiedQos_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface is a telemetry Collection whose Await method returns a slice of *Qos_Interface samples.
type CollectionQos_Interface struct {
	W    *Qos_InterfaceWatcher
	Data []*QualifiedQos_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface) Await(t testing.TB) []*QualifiedQos_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_InterfaceWatcher observes a stream of *Qos_Interface samples.
type Qos_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_InterfaceWatcher) Await(t testing.TB) (*QualifiedQos_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input is a *Qos_Interface_Input with a corresponding timestamp.
type QualifiedQos_Interface_Input struct {
	*genutil.Metadata
	val     *Qos_Interface_Input // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input) Val(t testing.TB) *Qos_Interface_Input {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input sample.
func (q *QualifiedQos_Interface_Input) SetVal(v *Qos_Interface_Input) *QualifiedQos_Interface_Input {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input samples.
type CollectionQos_Interface_Input struct {
	W    *Qos_Interface_InputWatcher
	Data []*QualifiedQos_Interface_Input
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input) Await(t testing.TB) []*QualifiedQos_Interface_Input {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_InputWatcher observes a stream of *Qos_Interface_Input samples.
type Qos_Interface_InputWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_InputWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input_Classifier is a *Qos_Interface_Input_Classifier with a corresponding timestamp.
type QualifiedQos_Interface_Input_Classifier struct {
	*genutil.Metadata
	val     *Qos_Interface_Input_Classifier // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input_Classifier) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input_Classifier sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input_Classifier) Val(t testing.TB) *Qos_Interface_Input_Classifier {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input_Classifier sample.
func (q *QualifiedQos_Interface_Input_Classifier) SetVal(v *Qos_Interface_Input_Classifier) *QualifiedQos_Interface_Input_Classifier {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input_Classifier) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input_Classifier is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input_Classifier samples.
type CollectionQos_Interface_Input_Classifier struct {
	W    *Qos_Interface_Input_ClassifierWatcher
	Data []*QualifiedQos_Interface_Input_Classifier
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input_Classifier) Await(t testing.TB) []*QualifiedQos_Interface_Input_Classifier {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Input_ClassifierWatcher observes a stream of *Qos_Interface_Input_Classifier samples.
type Qos_Interface_Input_ClassifierWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input_Classifier
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Input_ClassifierWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input_Classifier, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input_Classifier_Term is a *Qos_Interface_Input_Classifier_Term with a corresponding timestamp.
type QualifiedQos_Interface_Input_Classifier_Term struct {
	*genutil.Metadata
	val     *Qos_Interface_Input_Classifier_Term // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input_Classifier_Term) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input_Classifier_Term sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input_Classifier_Term) Val(t testing.TB) *Qos_Interface_Input_Classifier_Term {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input_Classifier_Term sample.
func (q *QualifiedQos_Interface_Input_Classifier_Term) SetVal(v *Qos_Interface_Input_Classifier_Term) *QualifiedQos_Interface_Input_Classifier_Term {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input_Classifier_Term) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input_Classifier_Term is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input_Classifier_Term samples.
type CollectionQos_Interface_Input_Classifier_Term struct {
	W    *Qos_Interface_Input_Classifier_TermWatcher
	Data []*QualifiedQos_Interface_Input_Classifier_Term
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input_Classifier_Term) Await(t testing.TB) []*QualifiedQos_Interface_Input_Classifier_Term {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Input_Classifier_TermWatcher observes a stream of *Qos_Interface_Input_Classifier_Term samples.
type Qos_Interface_Input_Classifier_TermWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input_Classifier_Term
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Input_Classifier_TermWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input_Classifier_Term, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input_Queue is a *Qos_Interface_Input_Queue with a corresponding timestamp.
type QualifiedQos_Interface_Input_Queue struct {
	*genutil.Metadata
	val     *Qos_Interface_Input_Queue // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input_Queue) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input_Queue sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input_Queue) Val(t testing.TB) *Qos_Interface_Input_Queue {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input_Queue sample.
func (q *QualifiedQos_Interface_Input_Queue) SetVal(v *Qos_Interface_Input_Queue) *QualifiedQos_Interface_Input_Queue {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input_Queue) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input_Queue is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input_Queue samples.
type CollectionQos_Interface_Input_Queue struct {
	W    *Qos_Interface_Input_QueueWatcher
	Data []*QualifiedQos_Interface_Input_Queue
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input_Queue) Await(t testing.TB) []*QualifiedQos_Interface_Input_Queue {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Input_QueueWatcher observes a stream of *Qos_Interface_Input_Queue samples.
type Qos_Interface_Input_QueueWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input_Queue
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Input_QueueWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input_Queue, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input_SchedulerPolicy is a *Qos_Interface_Input_SchedulerPolicy with a corresponding timestamp.
type QualifiedQos_Interface_Input_SchedulerPolicy struct {
	*genutil.Metadata
	val     *Qos_Interface_Input_SchedulerPolicy // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input_SchedulerPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input_SchedulerPolicy sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input_SchedulerPolicy) Val(t testing.TB) *Qos_Interface_Input_SchedulerPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input_SchedulerPolicy sample.
func (q *QualifiedQos_Interface_Input_SchedulerPolicy) SetVal(v *Qos_Interface_Input_SchedulerPolicy) *QualifiedQos_Interface_Input_SchedulerPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input_SchedulerPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input_SchedulerPolicy is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input_SchedulerPolicy samples.
type CollectionQos_Interface_Input_SchedulerPolicy struct {
	W    *Qos_Interface_Input_SchedulerPolicyWatcher
	Data []*QualifiedQos_Interface_Input_SchedulerPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input_SchedulerPolicy) Await(t testing.TB) []*QualifiedQos_Interface_Input_SchedulerPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Input_SchedulerPolicyWatcher observes a stream of *Qos_Interface_Input_SchedulerPolicy samples.
type Qos_Interface_Input_SchedulerPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input_SchedulerPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Input_SchedulerPolicyWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input_SchedulerPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler is a *Qos_Interface_Input_SchedulerPolicy_Scheduler with a corresponding timestamp.
type QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler struct {
	*genutil.Metadata
	val     *Qos_Interface_Input_SchedulerPolicy_Scheduler // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input_SchedulerPolicy_Scheduler sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) Val(t testing.TB) *Qos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input_SchedulerPolicy_Scheduler sample.
func (q *QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) SetVal(v *Qos_Interface_Input_SchedulerPolicy_Scheduler) *QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input_SchedulerPolicy_Scheduler is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input_SchedulerPolicy_Scheduler samples.
type CollectionQos_Interface_Input_SchedulerPolicy_Scheduler struct {
	W    *Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher
	Data []*QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input_SchedulerPolicy_Scheduler) Await(t testing.TB) []*QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher observes a stream of *Qos_Interface_Input_SchedulerPolicy_Scheduler samples.
type Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Input_SchedulerPolicy_SchedulerWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input_SchedulerPolicy_Scheduler, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input_VoqInterface is a *Qos_Interface_Input_VoqInterface with a corresponding timestamp.
type QualifiedQos_Interface_Input_VoqInterface struct {
	*genutil.Metadata
	val     *Qos_Interface_Input_VoqInterface // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input_VoqInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input_VoqInterface sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input_VoqInterface) Val(t testing.TB) *Qos_Interface_Input_VoqInterface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input_VoqInterface sample.
func (q *QualifiedQos_Interface_Input_VoqInterface) SetVal(v *Qos_Interface_Input_VoqInterface) *QualifiedQos_Interface_Input_VoqInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input_VoqInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input_VoqInterface is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input_VoqInterface samples.
type CollectionQos_Interface_Input_VoqInterface struct {
	W    *Qos_Interface_Input_VoqInterfaceWatcher
	Data []*QualifiedQos_Interface_Input_VoqInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input_VoqInterface) Await(t testing.TB) []*QualifiedQos_Interface_Input_VoqInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Input_VoqInterfaceWatcher observes a stream of *Qos_Interface_Input_VoqInterface samples.
type Qos_Interface_Input_VoqInterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input_VoqInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Input_VoqInterfaceWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input_VoqInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Input_VoqInterface_Queue is a *Qos_Interface_Input_VoqInterface_Queue with a corresponding timestamp.
type QualifiedQos_Interface_Input_VoqInterface_Queue struct {
	*genutil.Metadata
	val     *Qos_Interface_Input_VoqInterface_Queue // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Input_VoqInterface_Queue) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Input_VoqInterface_Queue sample, erroring out if not present.
func (q *QualifiedQos_Interface_Input_VoqInterface_Queue) Val(t testing.TB) *Qos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Input_VoqInterface_Queue sample.
func (q *QualifiedQos_Interface_Input_VoqInterface_Queue) SetVal(v *Qos_Interface_Input_VoqInterface_Queue) *QualifiedQos_Interface_Input_VoqInterface_Queue {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Input_VoqInterface_Queue) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Input_VoqInterface_Queue is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Input_VoqInterface_Queue samples.
type CollectionQos_Interface_Input_VoqInterface_Queue struct {
	W    *Qos_Interface_Input_VoqInterface_QueueWatcher
	Data []*QualifiedQos_Interface_Input_VoqInterface_Queue
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Input_VoqInterface_Queue) Await(t testing.TB) []*QualifiedQos_Interface_Input_VoqInterface_Queue {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Input_VoqInterface_QueueWatcher observes a stream of *Qos_Interface_Input_VoqInterface_Queue samples.
type Qos_Interface_Input_VoqInterface_QueueWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Input_VoqInterface_Queue
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Input_VoqInterface_QueueWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Input_VoqInterface_Queue, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_InterfaceRef is a *Qos_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedQos_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *Qos_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedQos_Interface_InterfaceRef) Val(t testing.TB) *Qos_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_InterfaceRef sample.
func (q *QualifiedQos_Interface_InterfaceRef) SetVal(v *Qos_Interface_InterfaceRef) *QualifiedQos_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Qos_Interface_InterfaceRef samples.
type CollectionQos_Interface_InterfaceRef struct {
	W    *Qos_Interface_InterfaceRefWatcher
	Data []*QualifiedQos_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedQos_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_InterfaceRefWatcher observes a stream of *Qos_Interface_InterfaceRef samples.
type Qos_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedQos_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Output is a *Qos_Interface_Output with a corresponding timestamp.
type QualifiedQos_Interface_Output struct {
	*genutil.Metadata
	val     *Qos_Interface_Output // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Output) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Output sample, erroring out if not present.
func (q *QualifiedQos_Interface_Output) Val(t testing.TB) *Qos_Interface_Output {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Output sample.
func (q *QualifiedQos_Interface_Output) SetVal(v *Qos_Interface_Output) *QualifiedQos_Interface_Output {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Output) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Output is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Output samples.
type CollectionQos_Interface_Output struct {
	W    *Qos_Interface_OutputWatcher
	Data []*QualifiedQos_Interface_Output
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Output) Await(t testing.TB) []*QualifiedQos_Interface_Output {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_OutputWatcher observes a stream of *Qos_Interface_Output samples.
type Qos_Interface_OutputWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Output
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_OutputWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Output, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Output_Classifier is a *Qos_Interface_Output_Classifier with a corresponding timestamp.
type QualifiedQos_Interface_Output_Classifier struct {
	*genutil.Metadata
	val     *Qos_Interface_Output_Classifier // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Output_Classifier) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Output_Classifier sample, erroring out if not present.
func (q *QualifiedQos_Interface_Output_Classifier) Val(t testing.TB) *Qos_Interface_Output_Classifier {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Output_Classifier sample.
func (q *QualifiedQos_Interface_Output_Classifier) SetVal(v *Qos_Interface_Output_Classifier) *QualifiedQos_Interface_Output_Classifier {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Output_Classifier) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Output_Classifier is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Output_Classifier samples.
type CollectionQos_Interface_Output_Classifier struct {
	W    *Qos_Interface_Output_ClassifierWatcher
	Data []*QualifiedQos_Interface_Output_Classifier
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Output_Classifier) Await(t testing.TB) []*QualifiedQos_Interface_Output_Classifier {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Output_ClassifierWatcher observes a stream of *Qos_Interface_Output_Classifier samples.
type Qos_Interface_Output_ClassifierWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Output_Classifier
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Output_ClassifierWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Output_Classifier, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Output_Classifier_Term is a *Qos_Interface_Output_Classifier_Term with a corresponding timestamp.
type QualifiedQos_Interface_Output_Classifier_Term struct {
	*genutil.Metadata
	val     *Qos_Interface_Output_Classifier_Term // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Output_Classifier_Term) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Output_Classifier_Term sample, erroring out if not present.
func (q *QualifiedQos_Interface_Output_Classifier_Term) Val(t testing.TB) *Qos_Interface_Output_Classifier_Term {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Output_Classifier_Term sample.
func (q *QualifiedQos_Interface_Output_Classifier_Term) SetVal(v *Qos_Interface_Output_Classifier_Term) *QualifiedQos_Interface_Output_Classifier_Term {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Output_Classifier_Term) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Output_Classifier_Term is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Output_Classifier_Term samples.
type CollectionQos_Interface_Output_Classifier_Term struct {
	W    *Qos_Interface_Output_Classifier_TermWatcher
	Data []*QualifiedQos_Interface_Output_Classifier_Term
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Output_Classifier_Term) Await(t testing.TB) []*QualifiedQos_Interface_Output_Classifier_Term {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Output_Classifier_TermWatcher observes a stream of *Qos_Interface_Output_Classifier_Term samples.
type Qos_Interface_Output_Classifier_TermWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Output_Classifier_Term
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Output_Classifier_TermWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Output_Classifier_Term, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Output_InterfaceRef is a *Qos_Interface_Output_InterfaceRef with a corresponding timestamp.
type QualifiedQos_Interface_Output_InterfaceRef struct {
	*genutil.Metadata
	val     *Qos_Interface_Output_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Output_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Output_InterfaceRef sample, erroring out if not present.
func (q *QualifiedQos_Interface_Output_InterfaceRef) Val(t testing.TB) *Qos_Interface_Output_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Output_InterfaceRef sample.
func (q *QualifiedQos_Interface_Output_InterfaceRef) SetVal(v *Qos_Interface_Output_InterfaceRef) *QualifiedQos_Interface_Output_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Output_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Output_InterfaceRef is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Output_InterfaceRef samples.
type CollectionQos_Interface_Output_InterfaceRef struct {
	W    *Qos_Interface_Output_InterfaceRefWatcher
	Data []*QualifiedQos_Interface_Output_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Output_InterfaceRef) Await(t testing.TB) []*QualifiedQos_Interface_Output_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Output_InterfaceRefWatcher observes a stream of *Qos_Interface_Output_InterfaceRef samples.
type Qos_Interface_Output_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Output_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Output_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Output_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Interface_Output_Queue is a *Qos_Interface_Output_Queue with a corresponding timestamp.
type QualifiedQos_Interface_Output_Queue struct {
	*genutil.Metadata
	val     *Qos_Interface_Output_Queue // val is the sample value.
	present bool
}

func (q *QualifiedQos_Interface_Output_Queue) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_Interface_Output_Queue sample, erroring out if not present.
func (q *QualifiedQos_Interface_Output_Queue) Val(t testing.TB) *Qos_Interface_Output_Queue {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_Interface_Output_Queue sample.
func (q *QualifiedQos_Interface_Output_Queue) SetVal(v *Qos_Interface_Output_Queue) *QualifiedQos_Interface_Output_Queue {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Interface_Output_Queue) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Interface_Output_Queue is a telemetry Collection whose Await method returns a slice of *Qos_Interface_Output_Queue samples.
type CollectionQos_Interface_Output_Queue struct {
	W    *Qos_Interface_Output_QueueWatcher
	Data []*QualifiedQos_Interface_Output_Queue
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Interface_Output_Queue) Await(t testing.TB) []*QualifiedQos_Interface_Output_Queue {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Interface_Output_QueueWatcher observes a stream of *Qos_Interface_Output_Queue samples.
type Qos_Interface_Output_QueueWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Interface_Output_Queue
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Interface_Output_QueueWatcher) Await(t testing.TB) (*QualifiedQos_Interface_Output_Queue, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
