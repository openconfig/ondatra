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

// QualifiedE_EndpointVni_VniType is a E_EndpointVni_VniType with a corresponding timestamp.
type QualifiedE_EndpointVni_VniType struct {
	*genutil.Metadata
	val     E_EndpointVni_VniType // val is the sample value.
	present bool
}

func (q *QualifiedE_EndpointVni_VniType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_EndpointVni_VniType sample, erroring out if not present.
func (q *QualifiedE_EndpointVni_VniType) Val(t testing.TB) E_EndpointVni_VniType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_EndpointVni_VniType sample.
func (q *QualifiedE_EndpointVni_VniType) SetVal(v E_EndpointVni_VniType) *QualifiedE_EndpointVni_VniType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_EndpointVni_VniType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_EndpointVni_VniType is a telemetry Collection whose Await method returns a slice of E_EndpointVni_VniType samples.
type CollectionE_EndpointVni_VniType struct {
	W    *E_EndpointVni_VniTypeWatcher
	Data []*QualifiedE_EndpointVni_VniType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_EndpointVni_VniType) Await(t testing.TB) []*QualifiedE_EndpointVni_VniType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_EndpointVni_VniTypeWatcher observes a stream of E_EndpointVni_VniType samples.
type E_EndpointVni_VniTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_EndpointVni_VniType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_EndpointVni_VniTypeWatcher) Await(t testing.TB) (*QualifiedE_EndpointVni_VniType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Entry_EntryType is a E_Entry_EntryType with a corresponding timestamp.
type QualifiedE_Entry_EntryType struct {
	*genutil.Metadata
	val     E_Entry_EntryType // val is the sample value.
	present bool
}

func (q *QualifiedE_Entry_EntryType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Entry_EntryType sample, erroring out if not present.
func (q *QualifiedE_Entry_EntryType) Val(t testing.TB) E_Entry_EntryType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Entry_EntryType sample.
func (q *QualifiedE_Entry_EntryType) SetVal(v E_Entry_EntryType) *QualifiedE_Entry_EntryType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Entry_EntryType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Entry_EntryType is a telemetry Collection whose Await method returns a slice of E_Entry_EntryType samples.
type CollectionE_Entry_EntryType struct {
	W    *E_Entry_EntryTypeWatcher
	Data []*QualifiedE_Entry_EntryType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Entry_EntryType) Await(t testing.TB) []*QualifiedE_Entry_EntryType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Entry_EntryTypeWatcher observes a stream of E_Entry_EntryType samples.
type E_Entry_EntryTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Entry_EntryType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Entry_EntryTypeWatcher) Await(t testing.TB) (*QualifiedE_Entry_EntryType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Ethernet_DuplexMode is a E_Ethernet_DuplexMode with a corresponding timestamp.
type QualifiedE_Ethernet_DuplexMode struct {
	*genutil.Metadata
	val     E_Ethernet_DuplexMode // val is the sample value.
	present bool
}

func (q *QualifiedE_Ethernet_DuplexMode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Ethernet_DuplexMode sample, erroring out if not present.
func (q *QualifiedE_Ethernet_DuplexMode) Val(t testing.TB) E_Ethernet_DuplexMode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Ethernet_DuplexMode sample.
func (q *QualifiedE_Ethernet_DuplexMode) SetVal(v E_Ethernet_DuplexMode) *QualifiedE_Ethernet_DuplexMode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Ethernet_DuplexMode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Ethernet_DuplexMode is a telemetry Collection whose Await method returns a slice of E_Ethernet_DuplexMode samples.
type CollectionE_Ethernet_DuplexMode struct {
	W    *E_Ethernet_DuplexModeWatcher
	Data []*QualifiedE_Ethernet_DuplexMode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Ethernet_DuplexMode) Await(t testing.TB) []*QualifiedE_Ethernet_DuplexMode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Ethernet_DuplexModeWatcher observes a stream of E_Ethernet_DuplexMode samples.
type E_Ethernet_DuplexModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Ethernet_DuplexMode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Ethernet_DuplexModeWatcher) Await(t testing.TB) (*QualifiedE_Ethernet_DuplexMode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Ethernet_NegotiatedDuplexMode is a E_Ethernet_NegotiatedDuplexMode with a corresponding timestamp.
type QualifiedE_Ethernet_NegotiatedDuplexMode struct {
	*genutil.Metadata
	val     E_Ethernet_NegotiatedDuplexMode // val is the sample value.
	present bool
}

func (q *QualifiedE_Ethernet_NegotiatedDuplexMode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Ethernet_NegotiatedDuplexMode sample, erroring out if not present.
func (q *QualifiedE_Ethernet_NegotiatedDuplexMode) Val(t testing.TB) E_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Ethernet_NegotiatedDuplexMode sample.
func (q *QualifiedE_Ethernet_NegotiatedDuplexMode) SetVal(v E_Ethernet_NegotiatedDuplexMode) *QualifiedE_Ethernet_NegotiatedDuplexMode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Ethernet_NegotiatedDuplexMode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Ethernet_NegotiatedDuplexMode is a telemetry Collection whose Await method returns a slice of E_Ethernet_NegotiatedDuplexMode samples.
type CollectionE_Ethernet_NegotiatedDuplexMode struct {
	W    *E_Ethernet_NegotiatedDuplexModeWatcher
	Data []*QualifiedE_Ethernet_NegotiatedDuplexMode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Ethernet_NegotiatedDuplexMode) Await(t testing.TB) []*QualifiedE_Ethernet_NegotiatedDuplexMode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Ethernet_NegotiatedDuplexModeWatcher observes a stream of E_Ethernet_NegotiatedDuplexMode samples.
type E_Ethernet_NegotiatedDuplexModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Ethernet_NegotiatedDuplexMode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Ethernet_NegotiatedDuplexModeWatcher) Await(t testing.TB) (*QualifiedE_Ethernet_NegotiatedDuplexMode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Event_Record is a E_Event_Record with a corresponding timestamp.
type QualifiedE_Event_Record struct {
	*genutil.Metadata
	val     E_Event_Record // val is the sample value.
	present bool
}

func (q *QualifiedE_Event_Record) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Event_Record sample, erroring out if not present.
func (q *QualifiedE_Event_Record) Val(t testing.TB) E_Event_Record {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Event_Record sample.
func (q *QualifiedE_Event_Record) SetVal(v E_Event_Record) *QualifiedE_Event_Record {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Event_Record) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Event_Record is a telemetry Collection whose Await method returns a slice of E_Event_Record samples.
type CollectionE_Event_Record struct {
	W    *E_Event_RecordWatcher
	Data []*QualifiedE_Event_Record
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Event_Record) Await(t testing.TB) []*QualifiedE_Event_Record {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Event_RecordWatcher observes a stream of E_Event_Record samples.
type E_Event_RecordWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Event_Record
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Event_RecordWatcher) Await(t testing.TB) (*QualifiedE_Event_Record, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_EvpnInstance_ReplicationMode is a E_EvpnInstance_ReplicationMode with a corresponding timestamp.
type QualifiedE_EvpnInstance_ReplicationMode struct {
	*genutil.Metadata
	val     E_EvpnInstance_ReplicationMode // val is the sample value.
	present bool
}

func (q *QualifiedE_EvpnInstance_ReplicationMode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_EvpnInstance_ReplicationMode sample, erroring out if not present.
func (q *QualifiedE_EvpnInstance_ReplicationMode) Val(t testing.TB) E_EvpnInstance_ReplicationMode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_EvpnInstance_ReplicationMode sample.
func (q *QualifiedE_EvpnInstance_ReplicationMode) SetVal(v E_EvpnInstance_ReplicationMode) *QualifiedE_EvpnInstance_ReplicationMode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_EvpnInstance_ReplicationMode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_EvpnInstance_ReplicationMode is a telemetry Collection whose Await method returns a slice of E_EvpnInstance_ReplicationMode samples.
type CollectionE_EvpnInstance_ReplicationMode struct {
	W    *E_EvpnInstance_ReplicationModeWatcher
	Data []*QualifiedE_EvpnInstance_ReplicationMode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_EvpnInstance_ReplicationMode) Await(t testing.TB) []*QualifiedE_EvpnInstance_ReplicationMode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_EvpnInstance_ReplicationModeWatcher observes a stream of E_EvpnInstance_ReplicationMode samples.
type E_EvpnInstance_ReplicationModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_EvpnInstance_ReplicationMode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_EvpnInstance_ReplicationModeWatcher) Await(t testing.TB) (*QualifiedE_EvpnInstance_ReplicationMode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE is a E_EvpnTypes_EVPN_REDUNDANCY_MODE with a corresponding timestamp.
type QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE struct {
	*genutil.Metadata
	val     E_EvpnTypes_EVPN_REDUNDANCY_MODE // val is the sample value.
	present bool
}

func (q *QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_EvpnTypes_EVPN_REDUNDANCY_MODE sample, erroring out if not present.
func (q *QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE) Val(t testing.TB) E_EvpnTypes_EVPN_REDUNDANCY_MODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_EvpnTypes_EVPN_REDUNDANCY_MODE sample.
func (q *QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE) SetVal(v E_EvpnTypes_EVPN_REDUNDANCY_MODE) *QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_EvpnTypes_EVPN_REDUNDANCY_MODE is a telemetry Collection whose Await method returns a slice of E_EvpnTypes_EVPN_REDUNDANCY_MODE samples.
type CollectionE_EvpnTypes_EVPN_REDUNDANCY_MODE struct {
	W    *E_EvpnTypes_EVPN_REDUNDANCY_MODEWatcher
	Data []*QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_EvpnTypes_EVPN_REDUNDANCY_MODE) Await(t testing.TB) []*QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_EvpnTypes_EVPN_REDUNDANCY_MODEWatcher observes a stream of E_EvpnTypes_EVPN_REDUNDANCY_MODE samples.
type E_EvpnTypes_EVPN_REDUNDANCY_MODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_EvpnTypes_EVPN_REDUNDANCY_MODEWatcher) Await(t testing.TB) (*QualifiedE_EvpnTypes_EVPN_REDUNDANCY_MODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_EvpnTypes_EVPN_TYPE is a E_EvpnTypes_EVPN_TYPE with a corresponding timestamp.
type QualifiedE_EvpnTypes_EVPN_TYPE struct {
	*genutil.Metadata
	val     E_EvpnTypes_EVPN_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_EvpnTypes_EVPN_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_EvpnTypes_EVPN_TYPE sample, erroring out if not present.
func (q *QualifiedE_EvpnTypes_EVPN_TYPE) Val(t testing.TB) E_EvpnTypes_EVPN_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_EvpnTypes_EVPN_TYPE sample.
func (q *QualifiedE_EvpnTypes_EVPN_TYPE) SetVal(v E_EvpnTypes_EVPN_TYPE) *QualifiedE_EvpnTypes_EVPN_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_EvpnTypes_EVPN_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_EvpnTypes_EVPN_TYPE is a telemetry Collection whose Await method returns a slice of E_EvpnTypes_EVPN_TYPE samples.
type CollectionE_EvpnTypes_EVPN_TYPE struct {
	W    *E_EvpnTypes_EVPN_TYPEWatcher
	Data []*QualifiedE_EvpnTypes_EVPN_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_EvpnTypes_EVPN_TYPE) Await(t testing.TB) []*QualifiedE_EvpnTypes_EVPN_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_EvpnTypes_EVPN_TYPEWatcher observes a stream of E_EvpnTypes_EVPN_TYPE samples.
type E_EvpnTypes_EVPN_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_EvpnTypes_EVPN_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_EvpnTypes_EVPN_TYPEWatcher) Await(t testing.TB) (*QualifiedE_EvpnTypes_EVPN_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_EvpnTypes_EsiType is a E_EvpnTypes_EsiType with a corresponding timestamp.
type QualifiedE_EvpnTypes_EsiType struct {
	*genutil.Metadata
	val     E_EvpnTypes_EsiType // val is the sample value.
	present bool
}

func (q *QualifiedE_EvpnTypes_EsiType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_EvpnTypes_EsiType sample, erroring out if not present.
func (q *QualifiedE_EvpnTypes_EsiType) Val(t testing.TB) E_EvpnTypes_EsiType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_EvpnTypes_EsiType sample.
func (q *QualifiedE_EvpnTypes_EsiType) SetVal(v E_EvpnTypes_EsiType) *QualifiedE_EvpnTypes_EsiType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_EvpnTypes_EsiType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_EvpnTypes_EsiType is a telemetry Collection whose Await method returns a slice of E_EvpnTypes_EsiType samples.
type CollectionE_EvpnTypes_EsiType struct {
	W    *E_EvpnTypes_EsiTypeWatcher
	Data []*QualifiedE_EvpnTypes_EsiType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_EvpnTypes_EsiType) Await(t testing.TB) []*QualifiedE_EvpnTypes_EsiType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_EvpnTypes_EsiTypeWatcher observes a stream of E_EvpnTypes_EsiType samples.
type E_EvpnTypes_EsiTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_EvpnTypes_EsiType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_EvpnTypes_EsiTypeWatcher) Await(t testing.TB) (*QualifiedE_EvpnTypes_EsiType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_EvpnTypes_LearningMode is a E_EvpnTypes_LearningMode with a corresponding timestamp.
type QualifiedE_EvpnTypes_LearningMode struct {
	*genutil.Metadata
	val     E_EvpnTypes_LearningMode // val is the sample value.
	present bool
}

func (q *QualifiedE_EvpnTypes_LearningMode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_EvpnTypes_LearningMode sample, erroring out if not present.
func (q *QualifiedE_EvpnTypes_LearningMode) Val(t testing.TB) E_EvpnTypes_LearningMode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_EvpnTypes_LearningMode sample.
func (q *QualifiedE_EvpnTypes_LearningMode) SetVal(v E_EvpnTypes_LearningMode) *QualifiedE_EvpnTypes_LearningMode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_EvpnTypes_LearningMode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_EvpnTypes_LearningMode is a telemetry Collection whose Await method returns a slice of E_EvpnTypes_LearningMode samples.
type CollectionE_EvpnTypes_LearningMode struct {
	W    *E_EvpnTypes_LearningModeWatcher
	Data []*QualifiedE_EvpnTypes_LearningMode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_EvpnTypes_LearningMode) Await(t testing.TB) []*QualifiedE_EvpnTypes_LearningMode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_EvpnTypes_LearningModeWatcher observes a stream of E_EvpnTypes_LearningMode samples.
type E_EvpnTypes_LearningModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_EvpnTypes_LearningMode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_EvpnTypes_LearningModeWatcher) Await(t testing.TB) (*QualifiedE_EvpnTypes_LearningMode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_ExplicitRouteObject_Type is a E_ExplicitRouteObject_Type with a corresponding timestamp.
type QualifiedE_ExplicitRouteObject_Type struct {
	*genutil.Metadata
	val     E_ExplicitRouteObject_Type // val is the sample value.
	present bool
}

func (q *QualifiedE_ExplicitRouteObject_Type) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_ExplicitRouteObject_Type sample, erroring out if not present.
func (q *QualifiedE_ExplicitRouteObject_Type) Val(t testing.TB) E_ExplicitRouteObject_Type {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_ExplicitRouteObject_Type sample.
func (q *QualifiedE_ExplicitRouteObject_Type) SetVal(v E_ExplicitRouteObject_Type) *QualifiedE_ExplicitRouteObject_Type {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_ExplicitRouteObject_Type) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_ExplicitRouteObject_Type is a telemetry Collection whose Await method returns a slice of E_ExplicitRouteObject_Type samples.
type CollectionE_ExplicitRouteObject_Type struct {
	W    *E_ExplicitRouteObject_TypeWatcher
	Data []*QualifiedE_ExplicitRouteObject_Type
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_ExplicitRouteObject_Type) Await(t testing.TB) []*QualifiedE_ExplicitRouteObject_Type {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_ExplicitRouteObject_TypeWatcher observes a stream of E_ExplicitRouteObject_Type samples.
type E_ExplicitRouteObject_TypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_ExplicitRouteObject_Type
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_ExplicitRouteObject_TypeWatcher) Await(t testing.TB) (*QualifiedE_ExplicitRouteObject_Type, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_ExtendedPrefix_AddressFamily is a E_ExtendedPrefix_AddressFamily with a corresponding timestamp.
type QualifiedE_ExtendedPrefix_AddressFamily struct {
	*genutil.Metadata
	val     E_ExtendedPrefix_AddressFamily // val is the sample value.
	present bool
}

func (q *QualifiedE_ExtendedPrefix_AddressFamily) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_ExtendedPrefix_AddressFamily sample, erroring out if not present.
func (q *QualifiedE_ExtendedPrefix_AddressFamily) Val(t testing.TB) E_ExtendedPrefix_AddressFamily {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_ExtendedPrefix_AddressFamily sample.
func (q *QualifiedE_ExtendedPrefix_AddressFamily) SetVal(v E_ExtendedPrefix_AddressFamily) *QualifiedE_ExtendedPrefix_AddressFamily {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_ExtendedPrefix_AddressFamily) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_ExtendedPrefix_AddressFamily is a telemetry Collection whose Await method returns a slice of E_ExtendedPrefix_AddressFamily samples.
type CollectionE_ExtendedPrefix_AddressFamily struct {
	W    *E_ExtendedPrefix_AddressFamilyWatcher
	Data []*QualifiedE_ExtendedPrefix_AddressFamily
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_ExtendedPrefix_AddressFamily) Await(t testing.TB) []*QualifiedE_ExtendedPrefix_AddressFamily {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_ExtendedPrefix_AddressFamilyWatcher observes a stream of E_ExtendedPrefix_AddressFamily samples.
type E_ExtendedPrefix_AddressFamilyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_ExtendedPrefix_AddressFamily
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_ExtendedPrefix_AddressFamilyWatcher) Await(t testing.TB) (*QualifiedE_ExtendedPrefix_AddressFamily, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_ExtendedPrefix_RouteType is a E_ExtendedPrefix_RouteType with a corresponding timestamp.
type QualifiedE_ExtendedPrefix_RouteType struct {
	*genutil.Metadata
	val     E_ExtendedPrefix_RouteType // val is the sample value.
	present bool
}

func (q *QualifiedE_ExtendedPrefix_RouteType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_ExtendedPrefix_RouteType sample, erroring out if not present.
func (q *QualifiedE_ExtendedPrefix_RouteType) Val(t testing.TB) E_ExtendedPrefix_RouteType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_ExtendedPrefix_RouteType sample.
func (q *QualifiedE_ExtendedPrefix_RouteType) SetVal(v E_ExtendedPrefix_RouteType) *QualifiedE_ExtendedPrefix_RouteType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_ExtendedPrefix_RouteType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_ExtendedPrefix_RouteType is a telemetry Collection whose Await method returns a slice of E_ExtendedPrefix_RouteType samples.
type CollectionE_ExtendedPrefix_RouteType struct {
	W    *E_ExtendedPrefix_RouteTypeWatcher
	Data []*QualifiedE_ExtendedPrefix_RouteType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_ExtendedPrefix_RouteType) Await(t testing.TB) []*QualifiedE_ExtendedPrefix_RouteType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_ExtendedPrefix_RouteTypeWatcher observes a stream of E_ExtendedPrefix_RouteType samples.
type E_ExtendedPrefix_RouteTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_ExtendedPrefix_RouteType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_ExtendedPrefix_RouteTypeWatcher) Await(t testing.TB) (*QualifiedE_ExtendedPrefix_RouteType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Global_SummaryRouteCostMode is a E_Global_SummaryRouteCostMode with a corresponding timestamp.
type QualifiedE_Global_SummaryRouteCostMode struct {
	*genutil.Metadata
	val     E_Global_SummaryRouteCostMode // val is the sample value.
	present bool
}

func (q *QualifiedE_Global_SummaryRouteCostMode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Global_SummaryRouteCostMode sample, erroring out if not present.
func (q *QualifiedE_Global_SummaryRouteCostMode) Val(t testing.TB) E_Global_SummaryRouteCostMode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Global_SummaryRouteCostMode sample.
func (q *QualifiedE_Global_SummaryRouteCostMode) SetVal(v E_Global_SummaryRouteCostMode) *QualifiedE_Global_SummaryRouteCostMode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Global_SummaryRouteCostMode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Global_SummaryRouteCostMode is a telemetry Collection whose Await method returns a slice of E_Global_SummaryRouteCostMode samples.
type CollectionE_Global_SummaryRouteCostMode struct {
	W    *E_Global_SummaryRouteCostModeWatcher
	Data []*QualifiedE_Global_SummaryRouteCostMode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Global_SummaryRouteCostMode) Await(t testing.TB) []*QualifiedE_Global_SummaryRouteCostMode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Global_SummaryRouteCostModeWatcher observes a stream of E_Global_SummaryRouteCostMode samples.
type E_Global_SummaryRouteCostModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Global_SummaryRouteCostMode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Global_SummaryRouteCostModeWatcher) Await(t testing.TB) (*QualifiedE_Global_SummaryRouteCostMode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_GracefulRestart_Mode is a E_GracefulRestart_Mode with a corresponding timestamp.
type QualifiedE_GracefulRestart_Mode struct {
	*genutil.Metadata
	val     E_GracefulRestart_Mode // val is the sample value.
	present bool
}

func (q *QualifiedE_GracefulRestart_Mode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_GracefulRestart_Mode sample, erroring out if not present.
func (q *QualifiedE_GracefulRestart_Mode) Val(t testing.TB) E_GracefulRestart_Mode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_GracefulRestart_Mode sample.
func (q *QualifiedE_GracefulRestart_Mode) SetVal(v E_GracefulRestart_Mode) *QualifiedE_GracefulRestart_Mode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_GracefulRestart_Mode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_GracefulRestart_Mode is a telemetry Collection whose Await method returns a slice of E_GracefulRestart_Mode samples.
type CollectionE_GracefulRestart_Mode struct {
	W    *E_GracefulRestart_ModeWatcher
	Data []*QualifiedE_GracefulRestart_Mode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_GracefulRestart_Mode) Await(t testing.TB) []*QualifiedE_GracefulRestart_Mode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_GracefulRestart_ModeWatcher observes a stream of E_GracefulRestart_Mode samples.
type E_GracefulRestart_ModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_GracefulRestart_Mode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_GracefulRestart_ModeWatcher) Await(t testing.TB) (*QualifiedE_GracefulRestart_Mode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IETFInterfaces_InterfaceType is a E_IETFInterfaces_InterfaceType with a corresponding timestamp.
type QualifiedE_IETFInterfaces_InterfaceType struct {
	*genutil.Metadata
	val     E_IETFInterfaces_InterfaceType // val is the sample value.
	present bool
}

func (q *QualifiedE_IETFInterfaces_InterfaceType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IETFInterfaces_InterfaceType sample, erroring out if not present.
func (q *QualifiedE_IETFInterfaces_InterfaceType) Val(t testing.TB) E_IETFInterfaces_InterfaceType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IETFInterfaces_InterfaceType sample.
func (q *QualifiedE_IETFInterfaces_InterfaceType) SetVal(v E_IETFInterfaces_InterfaceType) *QualifiedE_IETFInterfaces_InterfaceType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IETFInterfaces_InterfaceType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IETFInterfaces_InterfaceType is a telemetry Collection whose Await method returns a slice of E_IETFInterfaces_InterfaceType samples.
type CollectionE_IETFInterfaces_InterfaceType struct {
	W    *E_IETFInterfaces_InterfaceTypeWatcher
	Data []*QualifiedE_IETFInterfaces_InterfaceType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IETFInterfaces_InterfaceType) Await(t testing.TB) []*QualifiedE_IETFInterfaces_InterfaceType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IETFInterfaces_InterfaceTypeWatcher observes a stream of E_IETFInterfaces_InterfaceType samples.
type E_IETFInterfaces_InterfaceTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IETFInterfaces_InterfaceType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IETFInterfaces_InterfaceTypeWatcher) Await(t testing.TB) (*QualifiedE_IETFInterfaces_InterfaceType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IfAggregate_AggregationType is a E_IfAggregate_AggregationType with a corresponding timestamp.
type QualifiedE_IfAggregate_AggregationType struct {
	*genutil.Metadata
	val     E_IfAggregate_AggregationType // val is the sample value.
	present bool
}

func (q *QualifiedE_IfAggregate_AggregationType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IfAggregate_AggregationType sample, erroring out if not present.
func (q *QualifiedE_IfAggregate_AggregationType) Val(t testing.TB) E_IfAggregate_AggregationType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IfAggregate_AggregationType sample.
func (q *QualifiedE_IfAggregate_AggregationType) SetVal(v E_IfAggregate_AggregationType) *QualifiedE_IfAggregate_AggregationType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IfAggregate_AggregationType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IfAggregate_AggregationType is a telemetry Collection whose Await method returns a slice of E_IfAggregate_AggregationType samples.
type CollectionE_IfAggregate_AggregationType struct {
	W    *E_IfAggregate_AggregationTypeWatcher
	Data []*QualifiedE_IfAggregate_AggregationType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IfAggregate_AggregationType) Await(t testing.TB) []*QualifiedE_IfAggregate_AggregationType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IfAggregate_AggregationTypeWatcher observes a stream of E_IfAggregate_AggregationType samples.
type E_IfAggregate_AggregationTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IfAggregate_AggregationType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IfAggregate_AggregationTypeWatcher) Await(t testing.TB) (*QualifiedE_IfAggregate_AggregationType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IfEthernet_ETHERNET_SPEED is a E_IfEthernet_ETHERNET_SPEED with a corresponding timestamp.
type QualifiedE_IfEthernet_ETHERNET_SPEED struct {
	*genutil.Metadata
	val     E_IfEthernet_ETHERNET_SPEED // val is the sample value.
	present bool
}

func (q *QualifiedE_IfEthernet_ETHERNET_SPEED) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IfEthernet_ETHERNET_SPEED sample, erroring out if not present.
func (q *QualifiedE_IfEthernet_ETHERNET_SPEED) Val(t testing.TB) E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IfEthernet_ETHERNET_SPEED sample.
func (q *QualifiedE_IfEthernet_ETHERNET_SPEED) SetVal(v E_IfEthernet_ETHERNET_SPEED) *QualifiedE_IfEthernet_ETHERNET_SPEED {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IfEthernet_ETHERNET_SPEED) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IfEthernet_ETHERNET_SPEED is a telemetry Collection whose Await method returns a slice of E_IfEthernet_ETHERNET_SPEED samples.
type CollectionE_IfEthernet_ETHERNET_SPEED struct {
	W    *E_IfEthernet_ETHERNET_SPEEDWatcher
	Data []*QualifiedE_IfEthernet_ETHERNET_SPEED
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IfEthernet_ETHERNET_SPEED) Await(t testing.TB) []*QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IfEthernet_ETHERNET_SPEEDWatcher observes a stream of E_IfEthernet_ETHERNET_SPEED samples.
type E_IfEthernet_ETHERNET_SPEEDWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IfEthernet_ETHERNET_SPEED
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IfEthernet_ETHERNET_SPEEDWatcher) Await(t testing.TB) (*QualifiedE_IfEthernet_ETHERNET_SPEED, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IfEthernet_INTERFACE_FEC is a E_IfEthernet_INTERFACE_FEC with a corresponding timestamp.
type QualifiedE_IfEthernet_INTERFACE_FEC struct {
	*genutil.Metadata
	val     E_IfEthernet_INTERFACE_FEC // val is the sample value.
	present bool
}

func (q *QualifiedE_IfEthernet_INTERFACE_FEC) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IfEthernet_INTERFACE_FEC sample, erroring out if not present.
func (q *QualifiedE_IfEthernet_INTERFACE_FEC) Val(t testing.TB) E_IfEthernet_INTERFACE_FEC {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IfEthernet_INTERFACE_FEC sample.
func (q *QualifiedE_IfEthernet_INTERFACE_FEC) SetVal(v E_IfEthernet_INTERFACE_FEC) *QualifiedE_IfEthernet_INTERFACE_FEC {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IfEthernet_INTERFACE_FEC) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IfEthernet_INTERFACE_FEC is a telemetry Collection whose Await method returns a slice of E_IfEthernet_INTERFACE_FEC samples.
type CollectionE_IfEthernet_INTERFACE_FEC struct {
	W    *E_IfEthernet_INTERFACE_FECWatcher
	Data []*QualifiedE_IfEthernet_INTERFACE_FEC
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IfEthernet_INTERFACE_FEC) Await(t testing.TB) []*QualifiedE_IfEthernet_INTERFACE_FEC {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IfEthernet_INTERFACE_FECWatcher observes a stream of E_IfEthernet_INTERFACE_FEC samples.
type E_IfEthernet_INTERFACE_FECWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IfEthernet_INTERFACE_FEC
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IfEthernet_INTERFACE_FECWatcher) Await(t testing.TB) (*QualifiedE_IfEthernet_INTERFACE_FEC, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IfIp_IpAddressOrigin is a E_IfIp_IpAddressOrigin with a corresponding timestamp.
type QualifiedE_IfIp_IpAddressOrigin struct {
	*genutil.Metadata
	val     E_IfIp_IpAddressOrigin // val is the sample value.
	present bool
}

func (q *QualifiedE_IfIp_IpAddressOrigin) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IfIp_IpAddressOrigin sample, erroring out if not present.
func (q *QualifiedE_IfIp_IpAddressOrigin) Val(t testing.TB) E_IfIp_IpAddressOrigin {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IfIp_IpAddressOrigin sample.
func (q *QualifiedE_IfIp_IpAddressOrigin) SetVal(v E_IfIp_IpAddressOrigin) *QualifiedE_IfIp_IpAddressOrigin {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IfIp_IpAddressOrigin) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IfIp_IpAddressOrigin is a telemetry Collection whose Await method returns a slice of E_IfIp_IpAddressOrigin samples.
type CollectionE_IfIp_IpAddressOrigin struct {
	W    *E_IfIp_IpAddressOriginWatcher
	Data []*QualifiedE_IfIp_IpAddressOrigin
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IfIp_IpAddressOrigin) Await(t testing.TB) []*QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IfIp_IpAddressOriginWatcher observes a stream of E_IfIp_IpAddressOrigin samples.
type E_IfIp_IpAddressOriginWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IfIp_IpAddressOrigin
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IfIp_IpAddressOriginWatcher) Await(t testing.TB) (*QualifiedE_IfIp_IpAddressOrigin, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IfIp_NeighborOrigin is a E_IfIp_NeighborOrigin with a corresponding timestamp.
type QualifiedE_IfIp_NeighborOrigin struct {
	*genutil.Metadata
	val     E_IfIp_NeighborOrigin // val is the sample value.
	present bool
}

func (q *QualifiedE_IfIp_NeighborOrigin) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IfIp_NeighborOrigin sample, erroring out if not present.
func (q *QualifiedE_IfIp_NeighborOrigin) Val(t testing.TB) E_IfIp_NeighborOrigin {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IfIp_NeighborOrigin sample.
func (q *QualifiedE_IfIp_NeighborOrigin) SetVal(v E_IfIp_NeighborOrigin) *QualifiedE_IfIp_NeighborOrigin {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IfIp_NeighborOrigin) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IfIp_NeighborOrigin is a telemetry Collection whose Await method returns a slice of E_IfIp_NeighborOrigin samples.
type CollectionE_IfIp_NeighborOrigin struct {
	W    *E_IfIp_NeighborOriginWatcher
	Data []*QualifiedE_IfIp_NeighborOrigin
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IfIp_NeighborOrigin) Await(t testing.TB) []*QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IfIp_NeighborOriginWatcher observes a stream of E_IfIp_NeighborOrigin samples.
type E_IfIp_NeighborOriginWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IfIp_NeighborOrigin
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IfIp_NeighborOriginWatcher) Await(t testing.TB) (*QualifiedE_IfIp_NeighborOrigin, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IgpFloodingBandwidth_ThresholdSpecification is a E_IgpFloodingBandwidth_ThresholdSpecification with a corresponding timestamp.
type QualifiedE_IgpFloodingBandwidth_ThresholdSpecification struct {
	*genutil.Metadata
	val     E_IgpFloodingBandwidth_ThresholdSpecification // val is the sample value.
	present bool
}

func (q *QualifiedE_IgpFloodingBandwidth_ThresholdSpecification) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IgpFloodingBandwidth_ThresholdSpecification sample, erroring out if not present.
func (q *QualifiedE_IgpFloodingBandwidth_ThresholdSpecification) Val(t testing.TB) E_IgpFloodingBandwidth_ThresholdSpecification {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IgpFloodingBandwidth_ThresholdSpecification sample.
func (q *QualifiedE_IgpFloodingBandwidth_ThresholdSpecification) SetVal(v E_IgpFloodingBandwidth_ThresholdSpecification) *QualifiedE_IgpFloodingBandwidth_ThresholdSpecification {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IgpFloodingBandwidth_ThresholdSpecification) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IgpFloodingBandwidth_ThresholdSpecification is a telemetry Collection whose Await method returns a slice of E_IgpFloodingBandwidth_ThresholdSpecification samples.
type CollectionE_IgpFloodingBandwidth_ThresholdSpecification struct {
	W    *E_IgpFloodingBandwidth_ThresholdSpecificationWatcher
	Data []*QualifiedE_IgpFloodingBandwidth_ThresholdSpecification
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IgpFloodingBandwidth_ThresholdSpecification) Await(t testing.TB) []*QualifiedE_IgpFloodingBandwidth_ThresholdSpecification {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IgpFloodingBandwidth_ThresholdSpecificationWatcher observes a stream of E_IgpFloodingBandwidth_ThresholdSpecification samples.
type E_IgpFloodingBandwidth_ThresholdSpecificationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IgpFloodingBandwidth_ThresholdSpecification
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IgpFloodingBandwidth_ThresholdSpecificationWatcher) Await(t testing.TB) (*QualifiedE_IgpFloodingBandwidth_ThresholdSpecification, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IgpFloodingBandwidth_ThresholdType is a E_IgpFloodingBandwidth_ThresholdType with a corresponding timestamp.
type QualifiedE_IgpFloodingBandwidth_ThresholdType struct {
	*genutil.Metadata
	val     E_IgpFloodingBandwidth_ThresholdType // val is the sample value.
	present bool
}

func (q *QualifiedE_IgpFloodingBandwidth_ThresholdType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IgpFloodingBandwidth_ThresholdType sample, erroring out if not present.
func (q *QualifiedE_IgpFloodingBandwidth_ThresholdType) Val(t testing.TB) E_IgpFloodingBandwidth_ThresholdType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IgpFloodingBandwidth_ThresholdType sample.
func (q *QualifiedE_IgpFloodingBandwidth_ThresholdType) SetVal(v E_IgpFloodingBandwidth_ThresholdType) *QualifiedE_IgpFloodingBandwidth_ThresholdType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IgpFloodingBandwidth_ThresholdType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IgpFloodingBandwidth_ThresholdType is a telemetry Collection whose Await method returns a slice of E_IgpFloodingBandwidth_ThresholdType samples.
type CollectionE_IgpFloodingBandwidth_ThresholdType struct {
	W    *E_IgpFloodingBandwidth_ThresholdTypeWatcher
	Data []*QualifiedE_IgpFloodingBandwidth_ThresholdType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IgpFloodingBandwidth_ThresholdType) Await(t testing.TB) []*QualifiedE_IgpFloodingBandwidth_ThresholdType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IgpFloodingBandwidth_ThresholdTypeWatcher observes a stream of E_IgpFloodingBandwidth_ThresholdType samples.
type E_IgpFloodingBandwidth_ThresholdTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IgpFloodingBandwidth_ThresholdType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IgpFloodingBandwidth_ThresholdTypeWatcher) Await(t testing.TB) (*QualifiedE_IgpFloodingBandwidth_ThresholdType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Input_Classifier_Type is a E_Input_Classifier_Type with a corresponding timestamp.
type QualifiedE_Input_Classifier_Type struct {
	*genutil.Metadata
	val     E_Input_Classifier_Type // val is the sample value.
	present bool
}

func (q *QualifiedE_Input_Classifier_Type) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Input_Classifier_Type sample, erroring out if not present.
func (q *QualifiedE_Input_Classifier_Type) Val(t testing.TB) E_Input_Classifier_Type {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Input_Classifier_Type sample.
func (q *QualifiedE_Input_Classifier_Type) SetVal(v E_Input_Classifier_Type) *QualifiedE_Input_Classifier_Type {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Input_Classifier_Type) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Input_Classifier_Type is a telemetry Collection whose Await method returns a slice of E_Input_Classifier_Type samples.
type CollectionE_Input_Classifier_Type struct {
	W    *E_Input_Classifier_TypeWatcher
	Data []*QualifiedE_Input_Classifier_Type
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Input_Classifier_Type) Await(t testing.TB) []*QualifiedE_Input_Classifier_Type {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Input_Classifier_TypeWatcher observes a stream of E_Input_Classifier_Type samples.
type E_Input_Classifier_TypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Input_Classifier_Type
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Input_Classifier_TypeWatcher) Await(t testing.TB) (*QualifiedE_Input_Classifier_Type, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Input_InputType is a E_Input_InputType with a corresponding timestamp.
type QualifiedE_Input_InputType struct {
	*genutil.Metadata
	val     E_Input_InputType // val is the sample value.
	present bool
}

func (q *QualifiedE_Input_InputType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Input_InputType sample, erroring out if not present.
func (q *QualifiedE_Input_InputType) Val(t testing.TB) E_Input_InputType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Input_InputType sample.
func (q *QualifiedE_Input_InputType) SetVal(v E_Input_InputType) *QualifiedE_Input_InputType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Input_InputType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Input_InputType is a telemetry Collection whose Await method returns a slice of E_Input_InputType samples.
type CollectionE_Input_InputType struct {
	W    *E_Input_InputTypeWatcher
	Data []*QualifiedE_Input_InputType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Input_InputType) Await(t testing.TB) []*QualifiedE_Input_InputType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Input_InputTypeWatcher observes a stream of E_Input_InputType samples.
type E_Input_InputTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Input_InputType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Input_InputTypeWatcher) Await(t testing.TB) (*QualifiedE_Input_InputType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Interface_AdminStatus is a E_Interface_AdminStatus with a corresponding timestamp.
type QualifiedE_Interface_AdminStatus struct {
	*genutil.Metadata
	val     E_Interface_AdminStatus // val is the sample value.
	present bool
}

func (q *QualifiedE_Interface_AdminStatus) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Interface_AdminStatus sample, erroring out if not present.
func (q *QualifiedE_Interface_AdminStatus) Val(t testing.TB) E_Interface_AdminStatus {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Interface_AdminStatus sample.
func (q *QualifiedE_Interface_AdminStatus) SetVal(v E_Interface_AdminStatus) *QualifiedE_Interface_AdminStatus {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Interface_AdminStatus) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Interface_AdminStatus is a telemetry Collection whose Await method returns a slice of E_Interface_AdminStatus samples.
type CollectionE_Interface_AdminStatus struct {
	W    *E_Interface_AdminStatusWatcher
	Data []*QualifiedE_Interface_AdminStatus
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Interface_AdminStatus) Await(t testing.TB) []*QualifiedE_Interface_AdminStatus {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Interface_AdminStatusWatcher observes a stream of E_Interface_AdminStatus samples.
type E_Interface_AdminStatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Interface_AdminStatus
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Interface_AdminStatusWatcher) Await(t testing.TB) (*QualifiedE_Interface_AdminStatus, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Interface_IrbAnycastGateway is a E_Interface_IrbAnycastGateway with a corresponding timestamp.
type QualifiedE_Interface_IrbAnycastGateway struct {
	*genutil.Metadata
	val     E_Interface_IrbAnycastGateway // val is the sample value.
	present bool
}

func (q *QualifiedE_Interface_IrbAnycastGateway) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Interface_IrbAnycastGateway sample, erroring out if not present.
func (q *QualifiedE_Interface_IrbAnycastGateway) Val(t testing.TB) E_Interface_IrbAnycastGateway {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Interface_IrbAnycastGateway sample.
func (q *QualifiedE_Interface_IrbAnycastGateway) SetVal(v E_Interface_IrbAnycastGateway) *QualifiedE_Interface_IrbAnycastGateway {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Interface_IrbAnycastGateway) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Interface_IrbAnycastGateway is a telemetry Collection whose Await method returns a slice of E_Interface_IrbAnycastGateway samples.
type CollectionE_Interface_IrbAnycastGateway struct {
	W    *E_Interface_IrbAnycastGatewayWatcher
	Data []*QualifiedE_Interface_IrbAnycastGateway
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Interface_IrbAnycastGateway) Await(t testing.TB) []*QualifiedE_Interface_IrbAnycastGateway {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Interface_IrbAnycastGatewayWatcher observes a stream of E_Interface_IrbAnycastGateway samples.
type E_Interface_IrbAnycastGatewayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Interface_IrbAnycastGateway
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Interface_IrbAnycastGatewayWatcher) Await(t testing.TB) (*QualifiedE_Interface_IrbAnycastGateway, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Interface_OperStatus is a E_Interface_OperStatus with a corresponding timestamp.
type QualifiedE_Interface_OperStatus struct {
	*genutil.Metadata
	val     E_Interface_OperStatus // val is the sample value.
	present bool
}

func (q *QualifiedE_Interface_OperStatus) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Interface_OperStatus sample, erroring out if not present.
func (q *QualifiedE_Interface_OperStatus) Val(t testing.TB) E_Interface_OperStatus {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Interface_OperStatus sample.
func (q *QualifiedE_Interface_OperStatus) SetVal(v E_Interface_OperStatus) *QualifiedE_Interface_OperStatus {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Interface_OperStatus) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Interface_OperStatus is a telemetry Collection whose Await method returns a slice of E_Interface_OperStatus samples.
type CollectionE_Interface_OperStatus struct {
	W    *E_Interface_OperStatusWatcher
	Data []*QualifiedE_Interface_OperStatus
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Interface_OperStatus) Await(t testing.TB) []*QualifiedE_Interface_OperStatus {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Interface_OperStatusWatcher observes a stream of E_Interface_OperStatus samples.
type E_Interface_OperStatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Interface_OperStatus
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Interface_OperStatusWatcher) Await(t testing.TB) (*QualifiedE_Interface_OperStatus, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE is a E_IsisLsdbTypes_ISIS_SUBTLV_TYPE with a corresponding timestamp.
type QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE struct {
	*genutil.Metadata
	val     E_IsisLsdbTypes_ISIS_SUBTLV_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisLsdbTypes_ISIS_SUBTLV_TYPE sample, erroring out if not present.
func (q *QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE) Val(t testing.TB) E_IsisLsdbTypes_ISIS_SUBTLV_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisLsdbTypes_ISIS_SUBTLV_TYPE sample.
func (q *QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE) SetVal(v E_IsisLsdbTypes_ISIS_SUBTLV_TYPE) *QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisLsdbTypes_ISIS_SUBTLV_TYPE is a telemetry Collection whose Await method returns a slice of E_IsisLsdbTypes_ISIS_SUBTLV_TYPE samples.
type CollectionE_IsisLsdbTypes_ISIS_SUBTLV_TYPE struct {
	W    *E_IsisLsdbTypes_ISIS_SUBTLV_TYPEWatcher
	Data []*QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisLsdbTypes_ISIS_SUBTLV_TYPE) Await(t testing.TB) []*QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisLsdbTypes_ISIS_SUBTLV_TYPEWatcher observes a stream of E_IsisLsdbTypes_ISIS_SUBTLV_TYPE samples.
type E_IsisLsdbTypes_ISIS_SUBTLV_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisLsdbTypes_ISIS_SUBTLV_TYPEWatcher) Await(t testing.TB) (*QualifiedE_IsisLsdbTypes_ISIS_SUBTLV_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE is a E_IsisLsdbTypes_ISIS_TLV_TYPE with a corresponding timestamp.
type QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE struct {
	*genutil.Metadata
	val     E_IsisLsdbTypes_ISIS_TLV_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisLsdbTypes_ISIS_TLV_TYPE sample, erroring out if not present.
func (q *QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE) Val(t testing.TB) E_IsisLsdbTypes_ISIS_TLV_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisLsdbTypes_ISIS_TLV_TYPE sample.
func (q *QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE) SetVal(v E_IsisLsdbTypes_ISIS_TLV_TYPE) *QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisLsdbTypes_ISIS_TLV_TYPE is a telemetry Collection whose Await method returns a slice of E_IsisLsdbTypes_ISIS_TLV_TYPE samples.
type CollectionE_IsisLsdbTypes_ISIS_TLV_TYPE struct {
	W    *E_IsisLsdbTypes_ISIS_TLV_TYPEWatcher
	Data []*QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisLsdbTypes_ISIS_TLV_TYPE) Await(t testing.TB) []*QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisLsdbTypes_ISIS_TLV_TYPEWatcher observes a stream of E_IsisLsdbTypes_ISIS_TLV_TYPE samples.
type E_IsisLsdbTypes_ISIS_TLV_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisLsdbTypes_ISIS_TLV_TYPEWatcher) Await(t testing.TB) (*QualifiedE_IsisLsdbTypes_ISIS_TLV_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_AFI_TYPE is a E_IsisTypes_AFI_TYPE with a corresponding timestamp.
type QualifiedE_IsisTypes_AFI_TYPE struct {
	*genutil.Metadata
	val     E_IsisTypes_AFI_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_AFI_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_AFI_TYPE sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_AFI_TYPE) Val(t testing.TB) E_IsisTypes_AFI_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_AFI_TYPE sample.
func (q *QualifiedE_IsisTypes_AFI_TYPE) SetVal(v E_IsisTypes_AFI_TYPE) *QualifiedE_IsisTypes_AFI_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_AFI_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_AFI_TYPE is a telemetry Collection whose Await method returns a slice of E_IsisTypes_AFI_TYPE samples.
type CollectionE_IsisTypes_AFI_TYPE struct {
	W    *E_IsisTypes_AFI_TYPEWatcher
	Data []*QualifiedE_IsisTypes_AFI_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_AFI_TYPE) Await(t testing.TB) []*QualifiedE_IsisTypes_AFI_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_AFI_TYPEWatcher observes a stream of E_IsisTypes_AFI_TYPE samples.
type E_IsisTypes_AFI_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_AFI_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_AFI_TYPEWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_AFI_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_AUTH_MODE is a E_IsisTypes_AUTH_MODE with a corresponding timestamp.
type QualifiedE_IsisTypes_AUTH_MODE struct {
	*genutil.Metadata
	val     E_IsisTypes_AUTH_MODE // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_AUTH_MODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_AUTH_MODE sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_AUTH_MODE) Val(t testing.TB) E_IsisTypes_AUTH_MODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_AUTH_MODE sample.
func (q *QualifiedE_IsisTypes_AUTH_MODE) SetVal(v E_IsisTypes_AUTH_MODE) *QualifiedE_IsisTypes_AUTH_MODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_AUTH_MODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_AUTH_MODE is a telemetry Collection whose Await method returns a slice of E_IsisTypes_AUTH_MODE samples.
type CollectionE_IsisTypes_AUTH_MODE struct {
	W    *E_IsisTypes_AUTH_MODEWatcher
	Data []*QualifiedE_IsisTypes_AUTH_MODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_AUTH_MODE) Await(t testing.TB) []*QualifiedE_IsisTypes_AUTH_MODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_AUTH_MODEWatcher observes a stream of E_IsisTypes_AUTH_MODE samples.
type E_IsisTypes_AUTH_MODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_AUTH_MODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_AUTH_MODEWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_AUTH_MODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_AdaptiveTimerType is a E_IsisTypes_AdaptiveTimerType with a corresponding timestamp.
type QualifiedE_IsisTypes_AdaptiveTimerType struct {
	*genutil.Metadata
	val     E_IsisTypes_AdaptiveTimerType // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_AdaptiveTimerType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_AdaptiveTimerType sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_AdaptiveTimerType) Val(t testing.TB) E_IsisTypes_AdaptiveTimerType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_AdaptiveTimerType sample.
func (q *QualifiedE_IsisTypes_AdaptiveTimerType) SetVal(v E_IsisTypes_AdaptiveTimerType) *QualifiedE_IsisTypes_AdaptiveTimerType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_AdaptiveTimerType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_AdaptiveTimerType is a telemetry Collection whose Await method returns a slice of E_IsisTypes_AdaptiveTimerType samples.
type CollectionE_IsisTypes_AdaptiveTimerType struct {
	W    *E_IsisTypes_AdaptiveTimerTypeWatcher
	Data []*QualifiedE_IsisTypes_AdaptiveTimerType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_AdaptiveTimerType) Await(t testing.TB) []*QualifiedE_IsisTypes_AdaptiveTimerType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_AdaptiveTimerTypeWatcher observes a stream of E_IsisTypes_AdaptiveTimerType samples.
type E_IsisTypes_AdaptiveTimerTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_AdaptiveTimerType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_AdaptiveTimerTypeWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_AdaptiveTimerType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_CircuitType is a E_IsisTypes_CircuitType with a corresponding timestamp.
type QualifiedE_IsisTypes_CircuitType struct {
	*genutil.Metadata
	val     E_IsisTypes_CircuitType // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_CircuitType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_CircuitType sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_CircuitType) Val(t testing.TB) E_IsisTypes_CircuitType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_CircuitType sample.
func (q *QualifiedE_IsisTypes_CircuitType) SetVal(v E_IsisTypes_CircuitType) *QualifiedE_IsisTypes_CircuitType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_CircuitType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_CircuitType is a telemetry Collection whose Await method returns a slice of E_IsisTypes_CircuitType samples.
type CollectionE_IsisTypes_CircuitType struct {
	W    *E_IsisTypes_CircuitTypeWatcher
	Data []*QualifiedE_IsisTypes_CircuitType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_CircuitType) Await(t testing.TB) []*QualifiedE_IsisTypes_CircuitType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_CircuitTypeWatcher observes a stream of E_IsisTypes_CircuitType samples.
type E_IsisTypes_CircuitTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_CircuitType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_CircuitTypeWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_CircuitType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_HelloPaddingType is a E_IsisTypes_HelloPaddingType with a corresponding timestamp.
type QualifiedE_IsisTypes_HelloPaddingType struct {
	*genutil.Metadata
	val     E_IsisTypes_HelloPaddingType // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_HelloPaddingType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_HelloPaddingType sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_HelloPaddingType) Val(t testing.TB) E_IsisTypes_HelloPaddingType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_HelloPaddingType sample.
func (q *QualifiedE_IsisTypes_HelloPaddingType) SetVal(v E_IsisTypes_HelloPaddingType) *QualifiedE_IsisTypes_HelloPaddingType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_HelloPaddingType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_HelloPaddingType is a telemetry Collection whose Await method returns a slice of E_IsisTypes_HelloPaddingType samples.
type CollectionE_IsisTypes_HelloPaddingType struct {
	W    *E_IsisTypes_HelloPaddingTypeWatcher
	Data []*QualifiedE_IsisTypes_HelloPaddingType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_HelloPaddingType) Await(t testing.TB) []*QualifiedE_IsisTypes_HelloPaddingType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_HelloPaddingTypeWatcher observes a stream of E_IsisTypes_HelloPaddingType samples.
type E_IsisTypes_HelloPaddingTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_HelloPaddingType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_HelloPaddingTypeWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_HelloPaddingType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_IsisInterfaceAdjState is a E_IsisTypes_IsisInterfaceAdjState with a corresponding timestamp.
type QualifiedE_IsisTypes_IsisInterfaceAdjState struct {
	*genutil.Metadata
	val     E_IsisTypes_IsisInterfaceAdjState // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_IsisInterfaceAdjState) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_IsisInterfaceAdjState sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_IsisInterfaceAdjState) Val(t testing.TB) E_IsisTypes_IsisInterfaceAdjState {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_IsisInterfaceAdjState sample.
func (q *QualifiedE_IsisTypes_IsisInterfaceAdjState) SetVal(v E_IsisTypes_IsisInterfaceAdjState) *QualifiedE_IsisTypes_IsisInterfaceAdjState {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_IsisInterfaceAdjState) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_IsisInterfaceAdjState is a telemetry Collection whose Await method returns a slice of E_IsisTypes_IsisInterfaceAdjState samples.
type CollectionE_IsisTypes_IsisInterfaceAdjState struct {
	W    *E_IsisTypes_IsisInterfaceAdjStateWatcher
	Data []*QualifiedE_IsisTypes_IsisInterfaceAdjState
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_IsisInterfaceAdjState) Await(t testing.TB) []*QualifiedE_IsisTypes_IsisInterfaceAdjState {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_IsisInterfaceAdjStateWatcher observes a stream of E_IsisTypes_IsisInterfaceAdjState samples.
type E_IsisTypes_IsisInterfaceAdjStateWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_IsisInterfaceAdjState
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_IsisInterfaceAdjStateWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_IsisInterfaceAdjState, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_LevelType is a E_IsisTypes_LevelType with a corresponding timestamp.
type QualifiedE_IsisTypes_LevelType struct {
	*genutil.Metadata
	val     E_IsisTypes_LevelType // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_LevelType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_LevelType sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_LevelType) Val(t testing.TB) E_IsisTypes_LevelType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_LevelType sample.
func (q *QualifiedE_IsisTypes_LevelType) SetVal(v E_IsisTypes_LevelType) *QualifiedE_IsisTypes_LevelType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_LevelType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_LevelType is a telemetry Collection whose Await method returns a slice of E_IsisTypes_LevelType samples.
type CollectionE_IsisTypes_LevelType struct {
	W    *E_IsisTypes_LevelTypeWatcher
	Data []*QualifiedE_IsisTypes_LevelType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_LevelType) Await(t testing.TB) []*QualifiedE_IsisTypes_LevelType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_LevelTypeWatcher observes a stream of E_IsisTypes_LevelType samples.
type E_IsisTypes_LevelTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_LevelType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_LevelTypeWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_LevelType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_MetricStyle is a E_IsisTypes_MetricStyle with a corresponding timestamp.
type QualifiedE_IsisTypes_MetricStyle struct {
	*genutil.Metadata
	val     E_IsisTypes_MetricStyle // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_MetricStyle) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_MetricStyle sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_MetricStyle) Val(t testing.TB) E_IsisTypes_MetricStyle {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_MetricStyle sample.
func (q *QualifiedE_IsisTypes_MetricStyle) SetVal(v E_IsisTypes_MetricStyle) *QualifiedE_IsisTypes_MetricStyle {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_MetricStyle) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_MetricStyle is a telemetry Collection whose Await method returns a slice of E_IsisTypes_MetricStyle samples.
type CollectionE_IsisTypes_MetricStyle struct {
	W    *E_IsisTypes_MetricStyleWatcher
	Data []*QualifiedE_IsisTypes_MetricStyle
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_MetricStyle) Await(t testing.TB) []*QualifiedE_IsisTypes_MetricStyle {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_MetricStyleWatcher observes a stream of E_IsisTypes_MetricStyle samples.
type E_IsisTypes_MetricStyleWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_MetricStyle
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_MetricStyleWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_MetricStyle, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE is a E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE with a corresponding timestamp.
type QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE struct {
	*genutil.Metadata
	val     E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE) Val(t testing.TB) E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE sample.
func (q *QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE) SetVal(v E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE) *QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE is a telemetry Collection whose Await method returns a slice of E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE samples.
type CollectionE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE struct {
	W    *E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPEWatcher
	Data []*QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE) Await(t testing.TB) []*QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPEWatcher observes a stream of E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE samples.
type E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPEWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_OVERLOAD_RESET_TRIGGER_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_SAFI_TYPE is a E_IsisTypes_SAFI_TYPE with a corresponding timestamp.
type QualifiedE_IsisTypes_SAFI_TYPE struct {
	*genutil.Metadata
	val     E_IsisTypes_SAFI_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_SAFI_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_IsisTypes_SAFI_TYPE sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_SAFI_TYPE) Val(t testing.TB) E_IsisTypes_SAFI_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_IsisTypes_SAFI_TYPE sample.
func (q *QualifiedE_IsisTypes_SAFI_TYPE) SetVal(v E_IsisTypes_SAFI_TYPE) *QualifiedE_IsisTypes_SAFI_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_SAFI_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_SAFI_TYPE is a telemetry Collection whose Await method returns a slice of E_IsisTypes_SAFI_TYPE samples.
type CollectionE_IsisTypes_SAFI_TYPE struct {
	W    *E_IsisTypes_SAFI_TYPEWatcher
	Data []*QualifiedE_IsisTypes_SAFI_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_SAFI_TYPE) Await(t testing.TB) []*QualifiedE_IsisTypes_SAFI_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_SAFI_TYPEWatcher observes a stream of E_IsisTypes_SAFI_TYPE samples.
type E_IsisTypes_SAFI_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_SAFI_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_SAFI_TYPEWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_SAFI_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_KeychainTypes_AUTH_TYPE is a E_KeychainTypes_AUTH_TYPE with a corresponding timestamp.
type QualifiedE_KeychainTypes_AUTH_TYPE struct {
	*genutil.Metadata
	val     E_KeychainTypes_AUTH_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_KeychainTypes_AUTH_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_KeychainTypes_AUTH_TYPE sample, erroring out if not present.
func (q *QualifiedE_KeychainTypes_AUTH_TYPE) Val(t testing.TB) E_KeychainTypes_AUTH_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_KeychainTypes_AUTH_TYPE sample.
func (q *QualifiedE_KeychainTypes_AUTH_TYPE) SetVal(v E_KeychainTypes_AUTH_TYPE) *QualifiedE_KeychainTypes_AUTH_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_KeychainTypes_AUTH_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_KeychainTypes_AUTH_TYPE is a telemetry Collection whose Await method returns a slice of E_KeychainTypes_AUTH_TYPE samples.
type CollectionE_KeychainTypes_AUTH_TYPE struct {
	W    *E_KeychainTypes_AUTH_TYPEWatcher
	Data []*QualifiedE_KeychainTypes_AUTH_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_KeychainTypes_AUTH_TYPE) Await(t testing.TB) []*QualifiedE_KeychainTypes_AUTH_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_KeychainTypes_AUTH_TYPEWatcher observes a stream of E_KeychainTypes_AUTH_TYPE samples.
type E_KeychainTypes_AUTH_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_KeychainTypes_AUTH_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_KeychainTypes_AUTH_TYPEWatcher) Await(t testing.TB) (*QualifiedE_KeychainTypes_AUTH_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_KeychainTypes_CRYPTO_TYPE is a E_KeychainTypes_CRYPTO_TYPE with a corresponding timestamp.
type QualifiedE_KeychainTypes_CRYPTO_TYPE struct {
	*genutil.Metadata
	val     E_KeychainTypes_CRYPTO_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_KeychainTypes_CRYPTO_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_KeychainTypes_CRYPTO_TYPE sample, erroring out if not present.
func (q *QualifiedE_KeychainTypes_CRYPTO_TYPE) Val(t testing.TB) E_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_KeychainTypes_CRYPTO_TYPE sample.
func (q *QualifiedE_KeychainTypes_CRYPTO_TYPE) SetVal(v E_KeychainTypes_CRYPTO_TYPE) *QualifiedE_KeychainTypes_CRYPTO_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_KeychainTypes_CRYPTO_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_KeychainTypes_CRYPTO_TYPE is a telemetry Collection whose Await method returns a slice of E_KeychainTypes_CRYPTO_TYPE samples.
type CollectionE_KeychainTypes_CRYPTO_TYPE struct {
	W    *E_KeychainTypes_CRYPTO_TYPEWatcher
	Data []*QualifiedE_KeychainTypes_CRYPTO_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_KeychainTypes_CRYPTO_TYPE) Await(t testing.TB) []*QualifiedE_KeychainTypes_CRYPTO_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_KeychainTypes_CRYPTO_TYPEWatcher observes a stream of E_KeychainTypes_CRYPTO_TYPE samples.
type E_KeychainTypes_CRYPTO_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_KeychainTypes_CRYPTO_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_KeychainTypes_CRYPTO_TYPEWatcher) Await(t testing.TB) (*QualifiedE_KeychainTypes_CRYPTO_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lacp_LacpActivityType is a E_Lacp_LacpActivityType with a corresponding timestamp.
type QualifiedE_Lacp_LacpActivityType struct {
	*genutil.Metadata
	val     E_Lacp_LacpActivityType // val is the sample value.
	present bool
}

func (q *QualifiedE_Lacp_LacpActivityType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lacp_LacpActivityType sample, erroring out if not present.
func (q *QualifiedE_Lacp_LacpActivityType) Val(t testing.TB) E_Lacp_LacpActivityType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Lacp_LacpActivityType sample.
func (q *QualifiedE_Lacp_LacpActivityType) SetVal(v E_Lacp_LacpActivityType) *QualifiedE_Lacp_LacpActivityType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lacp_LacpActivityType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lacp_LacpActivityType is a telemetry Collection whose Await method returns a slice of E_Lacp_LacpActivityType samples.
type CollectionE_Lacp_LacpActivityType struct {
	W    *E_Lacp_LacpActivityTypeWatcher
	Data []*QualifiedE_Lacp_LacpActivityType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lacp_LacpActivityType) Await(t testing.TB) []*QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lacp_LacpActivityTypeWatcher observes a stream of E_Lacp_LacpActivityType samples.
type E_Lacp_LacpActivityTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lacp_LacpActivityType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lacp_LacpActivityTypeWatcher) Await(t testing.TB) (*QualifiedE_Lacp_LacpActivityType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lacp_LacpPeriodType is a E_Lacp_LacpPeriodType with a corresponding timestamp.
type QualifiedE_Lacp_LacpPeriodType struct {
	*genutil.Metadata
	val     E_Lacp_LacpPeriodType // val is the sample value.
	present bool
}

func (q *QualifiedE_Lacp_LacpPeriodType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lacp_LacpPeriodType sample, erroring out if not present.
func (q *QualifiedE_Lacp_LacpPeriodType) Val(t testing.TB) E_Lacp_LacpPeriodType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Lacp_LacpPeriodType sample.
func (q *QualifiedE_Lacp_LacpPeriodType) SetVal(v E_Lacp_LacpPeriodType) *QualifiedE_Lacp_LacpPeriodType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lacp_LacpPeriodType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lacp_LacpPeriodType is a telemetry Collection whose Await method returns a slice of E_Lacp_LacpPeriodType samples.
type CollectionE_Lacp_LacpPeriodType struct {
	W    *E_Lacp_LacpPeriodTypeWatcher
	Data []*QualifiedE_Lacp_LacpPeriodType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lacp_LacpPeriodType) Await(t testing.TB) []*QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lacp_LacpPeriodTypeWatcher observes a stream of E_Lacp_LacpPeriodType samples.
type E_Lacp_LacpPeriodTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lacp_LacpPeriodType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lacp_LacpPeriodTypeWatcher) Await(t testing.TB) (*QualifiedE_Lacp_LacpPeriodType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lacp_LacpSynchronizationType is a E_Lacp_LacpSynchronizationType with a corresponding timestamp.
type QualifiedE_Lacp_LacpSynchronizationType struct {
	*genutil.Metadata
	val     E_Lacp_LacpSynchronizationType // val is the sample value.
	present bool
}

func (q *QualifiedE_Lacp_LacpSynchronizationType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lacp_LacpSynchronizationType sample, erroring out if not present.
func (q *QualifiedE_Lacp_LacpSynchronizationType) Val(t testing.TB) E_Lacp_LacpSynchronizationType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Lacp_LacpSynchronizationType sample.
func (q *QualifiedE_Lacp_LacpSynchronizationType) SetVal(v E_Lacp_LacpSynchronizationType) *QualifiedE_Lacp_LacpSynchronizationType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lacp_LacpSynchronizationType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lacp_LacpSynchronizationType is a telemetry Collection whose Await method returns a slice of E_Lacp_LacpSynchronizationType samples.
type CollectionE_Lacp_LacpSynchronizationType struct {
	W    *E_Lacp_LacpSynchronizationTypeWatcher
	Data []*QualifiedE_Lacp_LacpSynchronizationType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lacp_LacpSynchronizationType) Await(t testing.TB) []*QualifiedE_Lacp_LacpSynchronizationType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lacp_LacpSynchronizationTypeWatcher observes a stream of E_Lacp_LacpSynchronizationType samples.
type E_Lacp_LacpSynchronizationTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lacp_LacpSynchronizationType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lacp_LacpSynchronizationTypeWatcher) Await(t testing.TB) (*QualifiedE_Lacp_LacpSynchronizationType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lacp_LacpTimeoutType is a E_Lacp_LacpTimeoutType with a corresponding timestamp.
type QualifiedE_Lacp_LacpTimeoutType struct {
	*genutil.Metadata
	val     E_Lacp_LacpTimeoutType // val is the sample value.
	present bool
}

func (q *QualifiedE_Lacp_LacpTimeoutType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lacp_LacpTimeoutType sample, erroring out if not present.
func (q *QualifiedE_Lacp_LacpTimeoutType) Val(t testing.TB) E_Lacp_LacpTimeoutType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Lacp_LacpTimeoutType sample.
func (q *QualifiedE_Lacp_LacpTimeoutType) SetVal(v E_Lacp_LacpTimeoutType) *QualifiedE_Lacp_LacpTimeoutType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lacp_LacpTimeoutType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lacp_LacpTimeoutType is a telemetry Collection whose Await method returns a slice of E_Lacp_LacpTimeoutType samples.
type CollectionE_Lacp_LacpTimeoutType struct {
	W    *E_Lacp_LacpTimeoutTypeWatcher
	Data []*QualifiedE_Lacp_LacpTimeoutType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lacp_LacpTimeoutType) Await(t testing.TB) []*QualifiedE_Lacp_LacpTimeoutType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lacp_LacpTimeoutTypeWatcher observes a stream of E_Lacp_LacpTimeoutType samples.
type E_Lacp_LacpTimeoutTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lacp_LacpTimeoutType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lacp_LacpTimeoutTypeWatcher) Await(t testing.TB) (*QualifiedE_Lacp_LacpTimeoutType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LldpTypes_ChassisIdType is a E_LldpTypes_ChassisIdType with a corresponding timestamp.
type QualifiedE_LldpTypes_ChassisIdType struct {
	*genutil.Metadata
	val     E_LldpTypes_ChassisIdType // val is the sample value.
	present bool
}

func (q *QualifiedE_LldpTypes_ChassisIdType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_LldpTypes_ChassisIdType sample, erroring out if not present.
func (q *QualifiedE_LldpTypes_ChassisIdType) Val(t testing.TB) E_LldpTypes_ChassisIdType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_LldpTypes_ChassisIdType sample.
func (q *QualifiedE_LldpTypes_ChassisIdType) SetVal(v E_LldpTypes_ChassisIdType) *QualifiedE_LldpTypes_ChassisIdType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LldpTypes_ChassisIdType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LldpTypes_ChassisIdType is a telemetry Collection whose Await method returns a slice of E_LldpTypes_ChassisIdType samples.
type CollectionE_LldpTypes_ChassisIdType struct {
	W    *E_LldpTypes_ChassisIdTypeWatcher
	Data []*QualifiedE_LldpTypes_ChassisIdType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LldpTypes_ChassisIdType) Await(t testing.TB) []*QualifiedE_LldpTypes_ChassisIdType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LldpTypes_ChassisIdTypeWatcher observes a stream of E_LldpTypes_ChassisIdType samples.
type E_LldpTypes_ChassisIdTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LldpTypes_ChassisIdType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LldpTypes_ChassisIdTypeWatcher) Await(t testing.TB) (*QualifiedE_LldpTypes_ChassisIdType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY is a E_LldpTypes_LLDP_SYSTEM_CAPABILITY with a corresponding timestamp.
type QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY struct {
	*genutil.Metadata
	val     E_LldpTypes_LLDP_SYSTEM_CAPABILITY // val is the sample value.
	present bool
}

func (q *QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_LldpTypes_LLDP_SYSTEM_CAPABILITY sample, erroring out if not present.
func (q *QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) Val(t testing.TB) E_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_LldpTypes_LLDP_SYSTEM_CAPABILITY sample.
func (q *QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) SetVal(v E_LldpTypes_LLDP_SYSTEM_CAPABILITY) *QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LldpTypes_LLDP_SYSTEM_CAPABILITY is a telemetry Collection whose Await method returns a slice of E_LldpTypes_LLDP_SYSTEM_CAPABILITY samples.
type CollectionE_LldpTypes_LLDP_SYSTEM_CAPABILITY struct {
	W    *E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher
	Data []*QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LldpTypes_LLDP_SYSTEM_CAPABILITY) Await(t testing.TB) []*QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher observes a stream of E_LldpTypes_LLDP_SYSTEM_CAPABILITY samples.
type E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LldpTypes_LLDP_SYSTEM_CAPABILITYWatcher) Await(t testing.TB) (*QualifiedE_LldpTypes_LLDP_SYSTEM_CAPABILITY, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LldpTypes_PortIdType is a E_LldpTypes_PortIdType with a corresponding timestamp.
type QualifiedE_LldpTypes_PortIdType struct {
	*genutil.Metadata
	val     E_LldpTypes_PortIdType // val is the sample value.
	present bool
}

func (q *QualifiedE_LldpTypes_PortIdType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_LldpTypes_PortIdType sample, erroring out if not present.
func (q *QualifiedE_LldpTypes_PortIdType) Val(t testing.TB) E_LldpTypes_PortIdType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_LldpTypes_PortIdType sample.
func (q *QualifiedE_LldpTypes_PortIdType) SetVal(v E_LldpTypes_PortIdType) *QualifiedE_LldpTypes_PortIdType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LldpTypes_PortIdType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LldpTypes_PortIdType is a telemetry Collection whose Await method returns a slice of E_LldpTypes_PortIdType samples.
type CollectionE_LldpTypes_PortIdType struct {
	W    *E_LldpTypes_PortIdTypeWatcher
	Data []*QualifiedE_LldpTypes_PortIdType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LldpTypes_PortIdType) Await(t testing.TB) []*QualifiedE_LldpTypes_PortIdType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LldpTypes_PortIdTypeWatcher observes a stream of E_LldpTypes_PortIdType samples.
type E_LldpTypes_PortIdTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LldpTypes_PortIdType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LldpTypes_PortIdTypeWatcher) Await(t testing.TB) (*QualifiedE_LldpTypes_PortIdType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LsaGeneration_TimerType is a E_LsaGeneration_TimerType with a corresponding timestamp.
type QualifiedE_LsaGeneration_TimerType struct {
	*genutil.Metadata
	val     E_LsaGeneration_TimerType // val is the sample value.
	present bool
}

func (q *QualifiedE_LsaGeneration_TimerType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_LsaGeneration_TimerType sample, erroring out if not present.
func (q *QualifiedE_LsaGeneration_TimerType) Val(t testing.TB) E_LsaGeneration_TimerType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_LsaGeneration_TimerType sample.
func (q *QualifiedE_LsaGeneration_TimerType) SetVal(v E_LsaGeneration_TimerType) *QualifiedE_LsaGeneration_TimerType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LsaGeneration_TimerType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LsaGeneration_TimerType is a telemetry Collection whose Await method returns a slice of E_LsaGeneration_TimerType samples.
type CollectionE_LsaGeneration_TimerType struct {
	W    *E_LsaGeneration_TimerTypeWatcher
	Data []*QualifiedE_LsaGeneration_TimerType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LsaGeneration_TimerType) Await(t testing.TB) []*QualifiedE_LsaGeneration_TimerType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LsaGeneration_TimerTypeWatcher observes a stream of E_LsaGeneration_TimerType samples.
type E_LsaGeneration_TimerTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LsaGeneration_TimerType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LsaGeneration_TimerTypeWatcher) Await(t testing.TB) (*QualifiedE_LsaGeneration_TimerType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lsp_PduType is a E_Lsp_PduType with a corresponding timestamp.
type QualifiedE_Lsp_PduType struct {
	*genutil.Metadata
	val     E_Lsp_PduType // val is the sample value.
	present bool
}

func (q *QualifiedE_Lsp_PduType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lsp_PduType sample, erroring out if not present.
func (q *QualifiedE_Lsp_PduType) Val(t testing.TB) E_Lsp_PduType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Lsp_PduType sample.
func (q *QualifiedE_Lsp_PduType) SetVal(v E_Lsp_PduType) *QualifiedE_Lsp_PduType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lsp_PduType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lsp_PduType is a telemetry Collection whose Await method returns a slice of E_Lsp_PduType samples.
type CollectionE_Lsp_PduType struct {
	W    *E_Lsp_PduTypeWatcher
	Data []*QualifiedE_Lsp_PduType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lsp_PduType) Await(t testing.TB) []*QualifiedE_Lsp_PduType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lsp_PduTypeWatcher observes a stream of E_Lsp_PduType samples.
type E_Lsp_PduTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lsp_PduType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lsp_PduTypeWatcher) Await(t testing.TB) (*QualifiedE_Lsp_PduType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Messages_DEBUG_SERVICE is a E_Messages_DEBUG_SERVICE with a corresponding timestamp.
type QualifiedE_Messages_DEBUG_SERVICE struct {
	*genutil.Metadata
	val     E_Messages_DEBUG_SERVICE // val is the sample value.
	present bool
}

func (q *QualifiedE_Messages_DEBUG_SERVICE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Messages_DEBUG_SERVICE sample, erroring out if not present.
func (q *QualifiedE_Messages_DEBUG_SERVICE) Val(t testing.TB) E_Messages_DEBUG_SERVICE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Messages_DEBUG_SERVICE sample.
func (q *QualifiedE_Messages_DEBUG_SERVICE) SetVal(v E_Messages_DEBUG_SERVICE) *QualifiedE_Messages_DEBUG_SERVICE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Messages_DEBUG_SERVICE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Messages_DEBUG_SERVICE is a telemetry Collection whose Await method returns a slice of E_Messages_DEBUG_SERVICE samples.
type CollectionE_Messages_DEBUG_SERVICE struct {
	W    *E_Messages_DEBUG_SERVICEWatcher
	Data []*QualifiedE_Messages_DEBUG_SERVICE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Messages_DEBUG_SERVICE) Await(t testing.TB) []*QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Messages_DEBUG_SERVICEWatcher observes a stream of E_Messages_DEBUG_SERVICE samples.
type E_Messages_DEBUG_SERVICEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Messages_DEBUG_SERVICE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Messages_DEBUG_SERVICEWatcher) Await(t testing.TB) (*QualifiedE_Messages_DEBUG_SERVICE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsLdp_MplsLdpAdjacencyType is a E_MplsLdp_MplsLdpAdjacencyType with a corresponding timestamp.
type QualifiedE_MplsLdp_MplsLdpAdjacencyType struct {
	*genutil.Metadata
	val     E_MplsLdp_MplsLdpAdjacencyType // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsLdp_MplsLdpAdjacencyType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsLdp_MplsLdpAdjacencyType sample, erroring out if not present.
func (q *QualifiedE_MplsLdp_MplsLdpAdjacencyType) Val(t testing.TB) E_MplsLdp_MplsLdpAdjacencyType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsLdp_MplsLdpAdjacencyType sample.
func (q *QualifiedE_MplsLdp_MplsLdpAdjacencyType) SetVal(v E_MplsLdp_MplsLdpAdjacencyType) *QualifiedE_MplsLdp_MplsLdpAdjacencyType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsLdp_MplsLdpAdjacencyType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsLdp_MplsLdpAdjacencyType is a telemetry Collection whose Await method returns a slice of E_MplsLdp_MplsLdpAdjacencyType samples.
type CollectionE_MplsLdp_MplsLdpAdjacencyType struct {
	W    *E_MplsLdp_MplsLdpAdjacencyTypeWatcher
	Data []*QualifiedE_MplsLdp_MplsLdpAdjacencyType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsLdp_MplsLdpAdjacencyType) Await(t testing.TB) []*QualifiedE_MplsLdp_MplsLdpAdjacencyType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsLdp_MplsLdpAdjacencyTypeWatcher observes a stream of E_MplsLdp_MplsLdpAdjacencyType samples.
type E_MplsLdp_MplsLdpAdjacencyTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsLdp_MplsLdpAdjacencyType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsLdp_MplsLdpAdjacencyTypeWatcher) Await(t testing.TB) (*QualifiedE_MplsLdp_MplsLdpAdjacencyType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsLdp_MplsLdpAfi is a E_MplsLdp_MplsLdpAfi with a corresponding timestamp.
type QualifiedE_MplsLdp_MplsLdpAfi struct {
	*genutil.Metadata
	val     E_MplsLdp_MplsLdpAfi // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsLdp_MplsLdpAfi) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsLdp_MplsLdpAfi sample, erroring out if not present.
func (q *QualifiedE_MplsLdp_MplsLdpAfi) Val(t testing.TB) E_MplsLdp_MplsLdpAfi {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsLdp_MplsLdpAfi sample.
func (q *QualifiedE_MplsLdp_MplsLdpAfi) SetVal(v E_MplsLdp_MplsLdpAfi) *QualifiedE_MplsLdp_MplsLdpAfi {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsLdp_MplsLdpAfi) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsLdp_MplsLdpAfi is a telemetry Collection whose Await method returns a slice of E_MplsLdp_MplsLdpAfi samples.
type CollectionE_MplsLdp_MplsLdpAfi struct {
	W    *E_MplsLdp_MplsLdpAfiWatcher
	Data []*QualifiedE_MplsLdp_MplsLdpAfi
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsLdp_MplsLdpAfi) Await(t testing.TB) []*QualifiedE_MplsLdp_MplsLdpAfi {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsLdp_MplsLdpAfiWatcher observes a stream of E_MplsLdp_MplsLdpAfi samples.
type E_MplsLdp_MplsLdpAfiWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsLdp_MplsLdpAfi
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsLdp_MplsLdpAfiWatcher) Await(t testing.TB) (*QualifiedE_MplsLdp_MplsLdpAfi, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsLdp_Neighbor_SessionState is a E_MplsLdp_Neighbor_SessionState with a corresponding timestamp.
type QualifiedE_MplsLdp_Neighbor_SessionState struct {
	*genutil.Metadata
	val     E_MplsLdp_Neighbor_SessionState // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsLdp_Neighbor_SessionState) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsLdp_Neighbor_SessionState sample, erroring out if not present.
func (q *QualifiedE_MplsLdp_Neighbor_SessionState) Val(t testing.TB) E_MplsLdp_Neighbor_SessionState {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsLdp_Neighbor_SessionState sample.
func (q *QualifiedE_MplsLdp_Neighbor_SessionState) SetVal(v E_MplsLdp_Neighbor_SessionState) *QualifiedE_MplsLdp_Neighbor_SessionState {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsLdp_Neighbor_SessionState) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsLdp_Neighbor_SessionState is a telemetry Collection whose Await method returns a slice of E_MplsLdp_Neighbor_SessionState samples.
type CollectionE_MplsLdp_Neighbor_SessionState struct {
	W    *E_MplsLdp_Neighbor_SessionStateWatcher
	Data []*QualifiedE_MplsLdp_Neighbor_SessionState
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsLdp_Neighbor_SessionState) Await(t testing.TB) []*QualifiedE_MplsLdp_Neighbor_SessionState {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsLdp_Neighbor_SessionStateWatcher observes a stream of E_MplsLdp_Neighbor_SessionState samples.
type E_MplsLdp_Neighbor_SessionStateWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsLdp_Neighbor_SessionState
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsLdp_Neighbor_SessionStateWatcher) Await(t testing.TB) (*QualifiedE_MplsLdp_Neighbor_SessionState, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_LSP_METRIC_TYPE is a E_MplsTypes_LSP_METRIC_TYPE with a corresponding timestamp.
type QualifiedE_MplsTypes_LSP_METRIC_TYPE struct {
	*genutil.Metadata
	val     E_MplsTypes_LSP_METRIC_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_LSP_METRIC_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_LSP_METRIC_TYPE sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_LSP_METRIC_TYPE) Val(t testing.TB) E_MplsTypes_LSP_METRIC_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_LSP_METRIC_TYPE sample.
func (q *QualifiedE_MplsTypes_LSP_METRIC_TYPE) SetVal(v E_MplsTypes_LSP_METRIC_TYPE) *QualifiedE_MplsTypes_LSP_METRIC_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_LSP_METRIC_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_LSP_METRIC_TYPE is a telemetry Collection whose Await method returns a slice of E_MplsTypes_LSP_METRIC_TYPE samples.
type CollectionE_MplsTypes_LSP_METRIC_TYPE struct {
	W    *E_MplsTypes_LSP_METRIC_TYPEWatcher
	Data []*QualifiedE_MplsTypes_LSP_METRIC_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_LSP_METRIC_TYPE) Await(t testing.TB) []*QualifiedE_MplsTypes_LSP_METRIC_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_LSP_METRIC_TYPEWatcher observes a stream of E_MplsTypes_LSP_METRIC_TYPE samples.
type E_MplsTypes_LSP_METRIC_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_LSP_METRIC_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_LSP_METRIC_TYPEWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_LSP_METRIC_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_LSP_OPER_STATUS is a E_MplsTypes_LSP_OPER_STATUS with a corresponding timestamp.
type QualifiedE_MplsTypes_LSP_OPER_STATUS struct {
	*genutil.Metadata
	val     E_MplsTypes_LSP_OPER_STATUS // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_LSP_OPER_STATUS) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_LSP_OPER_STATUS sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_LSP_OPER_STATUS) Val(t testing.TB) E_MplsTypes_LSP_OPER_STATUS {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_LSP_OPER_STATUS sample.
func (q *QualifiedE_MplsTypes_LSP_OPER_STATUS) SetVal(v E_MplsTypes_LSP_OPER_STATUS) *QualifiedE_MplsTypes_LSP_OPER_STATUS {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_LSP_OPER_STATUS) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_LSP_OPER_STATUS is a telemetry Collection whose Await method returns a slice of E_MplsTypes_LSP_OPER_STATUS samples.
type CollectionE_MplsTypes_LSP_OPER_STATUS struct {
	W    *E_MplsTypes_LSP_OPER_STATUSWatcher
	Data []*QualifiedE_MplsTypes_LSP_OPER_STATUS
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_LSP_OPER_STATUS) Await(t testing.TB) []*QualifiedE_MplsTypes_LSP_OPER_STATUS {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_LSP_OPER_STATUSWatcher observes a stream of E_MplsTypes_LSP_OPER_STATUS samples.
type E_MplsTypes_LSP_OPER_STATUSWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_LSP_OPER_STATUS
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_LSP_OPER_STATUSWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_LSP_OPER_STATUS, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_LSP_ROLE is a E_MplsTypes_LSP_ROLE with a corresponding timestamp.
type QualifiedE_MplsTypes_LSP_ROLE struct {
	*genutil.Metadata
	val     E_MplsTypes_LSP_ROLE // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_LSP_ROLE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_LSP_ROLE sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_LSP_ROLE) Val(t testing.TB) E_MplsTypes_LSP_ROLE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_LSP_ROLE sample.
func (q *QualifiedE_MplsTypes_LSP_ROLE) SetVal(v E_MplsTypes_LSP_ROLE) *QualifiedE_MplsTypes_LSP_ROLE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_LSP_ROLE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_LSP_ROLE is a telemetry Collection whose Await method returns a slice of E_MplsTypes_LSP_ROLE samples.
type CollectionE_MplsTypes_LSP_ROLE struct {
	W    *E_MplsTypes_LSP_ROLEWatcher
	Data []*QualifiedE_MplsTypes_LSP_ROLE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_LSP_ROLE) Await(t testing.TB) []*QualifiedE_MplsTypes_LSP_ROLE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_LSP_ROLEWatcher observes a stream of E_MplsTypes_LSP_ROLE samples.
type E_MplsTypes_LSP_ROLEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_LSP_ROLE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_LSP_ROLEWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_LSP_ROLE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_NULL_LABEL_TYPE is a E_MplsTypes_NULL_LABEL_TYPE with a corresponding timestamp.
type QualifiedE_MplsTypes_NULL_LABEL_TYPE struct {
	*genutil.Metadata
	val     E_MplsTypes_NULL_LABEL_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_NULL_LABEL_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_NULL_LABEL_TYPE sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_NULL_LABEL_TYPE) Val(t testing.TB) E_MplsTypes_NULL_LABEL_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_NULL_LABEL_TYPE sample.
func (q *QualifiedE_MplsTypes_NULL_LABEL_TYPE) SetVal(v E_MplsTypes_NULL_LABEL_TYPE) *QualifiedE_MplsTypes_NULL_LABEL_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_NULL_LABEL_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_NULL_LABEL_TYPE is a telemetry Collection whose Await method returns a slice of E_MplsTypes_NULL_LABEL_TYPE samples.
type CollectionE_MplsTypes_NULL_LABEL_TYPE struct {
	W    *E_MplsTypes_NULL_LABEL_TYPEWatcher
	Data []*QualifiedE_MplsTypes_NULL_LABEL_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_NULL_LABEL_TYPE) Await(t testing.TB) []*QualifiedE_MplsTypes_NULL_LABEL_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_NULL_LABEL_TYPEWatcher observes a stream of E_MplsTypes_NULL_LABEL_TYPE samples.
type E_MplsTypes_NULL_LABEL_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_NULL_LABEL_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_NULL_LABEL_TYPEWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_NULL_LABEL_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD is a E_MplsTypes_PATH_COMPUTATION_METHOD with a corresponding timestamp.
type QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD struct {
	*genutil.Metadata
	val     E_MplsTypes_PATH_COMPUTATION_METHOD // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_PATH_COMPUTATION_METHOD sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD) Val(t testing.TB) E_MplsTypes_PATH_COMPUTATION_METHOD {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_PATH_COMPUTATION_METHOD sample.
func (q *QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD) SetVal(v E_MplsTypes_PATH_COMPUTATION_METHOD) *QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_PATH_COMPUTATION_METHOD is a telemetry Collection whose Await method returns a slice of E_MplsTypes_PATH_COMPUTATION_METHOD samples.
type CollectionE_MplsTypes_PATH_COMPUTATION_METHOD struct {
	W    *E_MplsTypes_PATH_COMPUTATION_METHODWatcher
	Data []*QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_PATH_COMPUTATION_METHOD) Await(t testing.TB) []*QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_PATH_COMPUTATION_METHODWatcher observes a stream of E_MplsTypes_PATH_COMPUTATION_METHOD samples.
type E_MplsTypes_PATH_COMPUTATION_METHODWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_PATH_COMPUTATION_METHODWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_PATH_COMPUTATION_METHOD, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_PATH_METRIC_TYPE is a E_MplsTypes_PATH_METRIC_TYPE with a corresponding timestamp.
type QualifiedE_MplsTypes_PATH_METRIC_TYPE struct {
	*genutil.Metadata
	val     E_MplsTypes_PATH_METRIC_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_PATH_METRIC_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_PATH_METRIC_TYPE sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_PATH_METRIC_TYPE) Val(t testing.TB) E_MplsTypes_PATH_METRIC_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_PATH_METRIC_TYPE sample.
func (q *QualifiedE_MplsTypes_PATH_METRIC_TYPE) SetVal(v E_MplsTypes_PATH_METRIC_TYPE) *QualifiedE_MplsTypes_PATH_METRIC_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_PATH_METRIC_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_PATH_METRIC_TYPE is a telemetry Collection whose Await method returns a slice of E_MplsTypes_PATH_METRIC_TYPE samples.
type CollectionE_MplsTypes_PATH_METRIC_TYPE struct {
	W    *E_MplsTypes_PATH_METRIC_TYPEWatcher
	Data []*QualifiedE_MplsTypes_PATH_METRIC_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_PATH_METRIC_TYPE) Await(t testing.TB) []*QualifiedE_MplsTypes_PATH_METRIC_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_PATH_METRIC_TYPEWatcher observes a stream of E_MplsTypes_PATH_METRIC_TYPE samples.
type E_MplsTypes_PATH_METRIC_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_PATH_METRIC_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_PATH_METRIC_TYPEWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_PATH_METRIC_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL is a E_MplsTypes_PATH_SETUP_PROTOCOL with a corresponding timestamp.
type QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL struct {
	*genutil.Metadata
	val     E_MplsTypes_PATH_SETUP_PROTOCOL // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_PATH_SETUP_PROTOCOL sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL) Val(t testing.TB) E_MplsTypes_PATH_SETUP_PROTOCOL {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_PATH_SETUP_PROTOCOL sample.
func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL) SetVal(v E_MplsTypes_PATH_SETUP_PROTOCOL) *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_PATH_SETUP_PROTOCOL is a telemetry Collection whose Await method returns a slice of E_MplsTypes_PATH_SETUP_PROTOCOL samples.
type CollectionE_MplsTypes_PATH_SETUP_PROTOCOL struct {
	W    *E_MplsTypes_PATH_SETUP_PROTOCOLWatcher
	Data []*QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_PATH_SETUP_PROTOCOL) Await(t testing.TB) []*QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_PATH_SETUP_PROTOCOLWatcher observes a stream of E_MplsTypes_PATH_SETUP_PROTOCOL samples.
type E_MplsTypes_PATH_SETUP_PROTOCOLWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_PATH_SETUP_PROTOCOLWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_PATH_SETUP_PROTOCOL, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_PROTECTION_TYPE is a E_MplsTypes_PROTECTION_TYPE with a corresponding timestamp.
type QualifiedE_MplsTypes_PROTECTION_TYPE struct {
	*genutil.Metadata
	val     E_MplsTypes_PROTECTION_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_PROTECTION_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_PROTECTION_TYPE sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_PROTECTION_TYPE) Val(t testing.TB) E_MplsTypes_PROTECTION_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_PROTECTION_TYPE sample.
func (q *QualifiedE_MplsTypes_PROTECTION_TYPE) SetVal(v E_MplsTypes_PROTECTION_TYPE) *QualifiedE_MplsTypes_PROTECTION_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_PROTECTION_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_PROTECTION_TYPE is a telemetry Collection whose Await method returns a slice of E_MplsTypes_PROTECTION_TYPE samples.
type CollectionE_MplsTypes_PROTECTION_TYPE struct {
	W    *E_MplsTypes_PROTECTION_TYPEWatcher
	Data []*QualifiedE_MplsTypes_PROTECTION_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_PROTECTION_TYPE) Await(t testing.TB) []*QualifiedE_MplsTypes_PROTECTION_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_PROTECTION_TYPEWatcher observes a stream of E_MplsTypes_PROTECTION_TYPE samples.
type E_MplsTypes_PROTECTION_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_PROTECTION_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_PROTECTION_TYPEWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_PROTECTION_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION is a E_MplsTypes_PSEUDOWIRE_ENCAPSULATION with a corresponding timestamp.
type QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION struct {
	*genutil.Metadata
	val     E_MplsTypes_PSEUDOWIRE_ENCAPSULATION // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_PSEUDOWIRE_ENCAPSULATION sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION) Val(t testing.TB) E_MplsTypes_PSEUDOWIRE_ENCAPSULATION {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_PSEUDOWIRE_ENCAPSULATION sample.
func (q *QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION) SetVal(v E_MplsTypes_PSEUDOWIRE_ENCAPSULATION) *QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_PSEUDOWIRE_ENCAPSULATION is a telemetry Collection whose Await method returns a slice of E_MplsTypes_PSEUDOWIRE_ENCAPSULATION samples.
type CollectionE_MplsTypes_PSEUDOWIRE_ENCAPSULATION struct {
	W    *E_MplsTypes_PSEUDOWIRE_ENCAPSULATIONWatcher
	Data []*QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_PSEUDOWIRE_ENCAPSULATION) Await(t testing.TB) []*QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_PSEUDOWIRE_ENCAPSULATIONWatcher observes a stream of E_MplsTypes_PSEUDOWIRE_ENCAPSULATION samples.
type E_MplsTypes_PSEUDOWIRE_ENCAPSULATIONWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_PSEUDOWIRE_ENCAPSULATIONWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_PSEUDOWIRE_ENCAPSULATION, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_RSVP_AUTH_TYPE is a E_MplsTypes_RSVP_AUTH_TYPE with a corresponding timestamp.
type QualifiedE_MplsTypes_RSVP_AUTH_TYPE struct {
	*genutil.Metadata
	val     E_MplsTypes_RSVP_AUTH_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_RSVP_AUTH_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_RSVP_AUTH_TYPE sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_RSVP_AUTH_TYPE) Val(t testing.TB) E_MplsTypes_RSVP_AUTH_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_RSVP_AUTH_TYPE sample.
func (q *QualifiedE_MplsTypes_RSVP_AUTH_TYPE) SetVal(v E_MplsTypes_RSVP_AUTH_TYPE) *QualifiedE_MplsTypes_RSVP_AUTH_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_RSVP_AUTH_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_RSVP_AUTH_TYPE is a telemetry Collection whose Await method returns a slice of E_MplsTypes_RSVP_AUTH_TYPE samples.
type CollectionE_MplsTypes_RSVP_AUTH_TYPE struct {
	W    *E_MplsTypes_RSVP_AUTH_TYPEWatcher
	Data []*QualifiedE_MplsTypes_RSVP_AUTH_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_RSVP_AUTH_TYPE) Await(t testing.TB) []*QualifiedE_MplsTypes_RSVP_AUTH_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_RSVP_AUTH_TYPEWatcher observes a stream of E_MplsTypes_RSVP_AUTH_TYPE samples.
type E_MplsTypes_RSVP_AUTH_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_RSVP_AUTH_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_RSVP_AUTH_TYPEWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_RSVP_AUTH_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS is a E_MplsTypes_TUNNEL_ADMIN_STATUS with a corresponding timestamp.
type QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS struct {
	*genutil.Metadata
	val     E_MplsTypes_TUNNEL_ADMIN_STATUS // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_TUNNEL_ADMIN_STATUS sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS) Val(t testing.TB) E_MplsTypes_TUNNEL_ADMIN_STATUS {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_TUNNEL_ADMIN_STATUS sample.
func (q *QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS) SetVal(v E_MplsTypes_TUNNEL_ADMIN_STATUS) *QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_TUNNEL_ADMIN_STATUS is a telemetry Collection whose Await method returns a slice of E_MplsTypes_TUNNEL_ADMIN_STATUS samples.
type CollectionE_MplsTypes_TUNNEL_ADMIN_STATUS struct {
	W    *E_MplsTypes_TUNNEL_ADMIN_STATUSWatcher
	Data []*QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_TUNNEL_ADMIN_STATUS) Await(t testing.TB) []*QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_TUNNEL_ADMIN_STATUSWatcher observes a stream of E_MplsTypes_TUNNEL_ADMIN_STATUS samples.
type E_MplsTypes_TUNNEL_ADMIN_STATUSWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_TUNNEL_ADMIN_STATUSWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_TUNNEL_ADMIN_STATUS, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_TUNNEL_TYPE is a E_MplsTypes_TUNNEL_TYPE with a corresponding timestamp.
type QualifiedE_MplsTypes_TUNNEL_TYPE struct {
	*genutil.Metadata
	val     E_MplsTypes_TUNNEL_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_TUNNEL_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_MplsTypes_TUNNEL_TYPE sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_TUNNEL_TYPE) Val(t testing.TB) E_MplsTypes_TUNNEL_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_MplsTypes_TUNNEL_TYPE sample.
func (q *QualifiedE_MplsTypes_TUNNEL_TYPE) SetVal(v E_MplsTypes_TUNNEL_TYPE) *QualifiedE_MplsTypes_TUNNEL_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_TUNNEL_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_TUNNEL_TYPE is a telemetry Collection whose Await method returns a slice of E_MplsTypes_TUNNEL_TYPE samples.
type CollectionE_MplsTypes_TUNNEL_TYPE struct {
	W    *E_MplsTypes_TUNNEL_TYPEWatcher
	Data []*QualifiedE_MplsTypes_TUNNEL_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_TUNNEL_TYPE) Await(t testing.TB) []*QualifiedE_MplsTypes_TUNNEL_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_TUNNEL_TYPEWatcher observes a stream of E_MplsTypes_TUNNEL_TYPE samples.
type E_MplsTypes_TUNNEL_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_TUNNEL_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_TUNNEL_TYPEWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_TUNNEL_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Mpls_CspfTieBreaking is a E_Mpls_CspfTieBreaking with a corresponding timestamp.
type QualifiedE_Mpls_CspfTieBreaking struct {
	*genutil.Metadata
	val     E_Mpls_CspfTieBreaking // val is the sample value.
	present bool
}

func (q *QualifiedE_Mpls_CspfTieBreaking) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Mpls_CspfTieBreaking sample, erroring out if not present.
func (q *QualifiedE_Mpls_CspfTieBreaking) Val(t testing.TB) E_Mpls_CspfTieBreaking {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Mpls_CspfTieBreaking sample.
func (q *QualifiedE_Mpls_CspfTieBreaking) SetVal(v E_Mpls_CspfTieBreaking) *QualifiedE_Mpls_CspfTieBreaking {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Mpls_CspfTieBreaking) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Mpls_CspfTieBreaking is a telemetry Collection whose Await method returns a slice of E_Mpls_CspfTieBreaking samples.
type CollectionE_Mpls_CspfTieBreaking struct {
	W    *E_Mpls_CspfTieBreakingWatcher
	Data []*QualifiedE_Mpls_CspfTieBreaking
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Mpls_CspfTieBreaking) Await(t testing.TB) []*QualifiedE_Mpls_CspfTieBreaking {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Mpls_CspfTieBreakingWatcher observes a stream of E_Mpls_CspfTieBreaking samples.
type E_Mpls_CspfTieBreakingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Mpls_CspfTieBreaking
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Mpls_CspfTieBreakingWatcher) Await(t testing.TB) (*QualifiedE_Mpls_CspfTieBreaking, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Mpls_MplsHopType is a E_Mpls_MplsHopType with a corresponding timestamp.
type QualifiedE_Mpls_MplsHopType struct {
	*genutil.Metadata
	val     E_Mpls_MplsHopType // val is the sample value.
	present bool
}

func (q *QualifiedE_Mpls_MplsHopType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Mpls_MplsHopType sample, erroring out if not present.
func (q *QualifiedE_Mpls_MplsHopType) Val(t testing.TB) E_Mpls_MplsHopType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Mpls_MplsHopType sample.
func (q *QualifiedE_Mpls_MplsHopType) SetVal(v E_Mpls_MplsHopType) *QualifiedE_Mpls_MplsHopType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Mpls_MplsHopType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Mpls_MplsHopType is a telemetry Collection whose Await method returns a slice of E_Mpls_MplsHopType samples.
type CollectionE_Mpls_MplsHopType struct {
	W    *E_Mpls_MplsHopTypeWatcher
	Data []*QualifiedE_Mpls_MplsHopType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Mpls_MplsHopType) Await(t testing.TB) []*QualifiedE_Mpls_MplsHopType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Mpls_MplsHopTypeWatcher observes a stream of E_Mpls_MplsHopType samples.
type E_Mpls_MplsHopTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Mpls_MplsHopType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Mpls_MplsHopTypeWatcher) Await(t testing.TB) (*QualifiedE_Mpls_MplsHopType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Mpls_MplsSrlgFloodingType is a E_Mpls_MplsSrlgFloodingType with a corresponding timestamp.
type QualifiedE_Mpls_MplsSrlgFloodingType struct {
	*genutil.Metadata
	val     E_Mpls_MplsSrlgFloodingType // val is the sample value.
	present bool
}

func (q *QualifiedE_Mpls_MplsSrlgFloodingType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Mpls_MplsSrlgFloodingType sample, erroring out if not present.
func (q *QualifiedE_Mpls_MplsSrlgFloodingType) Val(t testing.TB) E_Mpls_MplsSrlgFloodingType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Mpls_MplsSrlgFloodingType sample.
func (q *QualifiedE_Mpls_MplsSrlgFloodingType) SetVal(v E_Mpls_MplsSrlgFloodingType) *QualifiedE_Mpls_MplsSrlgFloodingType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Mpls_MplsSrlgFloodingType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Mpls_MplsSrlgFloodingType is a telemetry Collection whose Await method returns a slice of E_Mpls_MplsSrlgFloodingType samples.
type CollectionE_Mpls_MplsSrlgFloodingType struct {
	W    *E_Mpls_MplsSrlgFloodingTypeWatcher
	Data []*QualifiedE_Mpls_MplsSrlgFloodingType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Mpls_MplsSrlgFloodingType) Await(t testing.TB) []*QualifiedE_Mpls_MplsSrlgFloodingType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Mpls_MplsSrlgFloodingTypeWatcher observes a stream of E_Mpls_MplsSrlgFloodingType samples.
type E_Mpls_MplsSrlgFloodingTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Mpls_MplsSrlgFloodingType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Mpls_MplsSrlgFloodingTypeWatcher) Await(t testing.TB) (*QualifiedE_Mpls_MplsSrlgFloodingType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Mpls_TeBandwidthType is a E_Mpls_TeBandwidthType with a corresponding timestamp.
type QualifiedE_Mpls_TeBandwidthType struct {
	*genutil.Metadata
	val     E_Mpls_TeBandwidthType // val is the sample value.
	present bool
}

func (q *QualifiedE_Mpls_TeBandwidthType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Mpls_TeBandwidthType sample, erroring out if not present.
func (q *QualifiedE_Mpls_TeBandwidthType) Val(t testing.TB) E_Mpls_TeBandwidthType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Mpls_TeBandwidthType sample.
func (q *QualifiedE_Mpls_TeBandwidthType) SetVal(v E_Mpls_TeBandwidthType) *QualifiedE_Mpls_TeBandwidthType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Mpls_TeBandwidthType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Mpls_TeBandwidthType is a telemetry Collection whose Await method returns a slice of E_Mpls_TeBandwidthType samples.
type CollectionE_Mpls_TeBandwidthType struct {
	W    *E_Mpls_TeBandwidthTypeWatcher
	Data []*QualifiedE_Mpls_TeBandwidthType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Mpls_TeBandwidthType) Await(t testing.TB) []*QualifiedE_Mpls_TeBandwidthType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Mpls_TeBandwidthTypeWatcher observes a stream of E_Mpls_TeBandwidthType samples.
type E_Mpls_TeBandwidthTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Mpls_TeBandwidthType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Mpls_TeBandwidthTypeWatcher) Await(t testing.TB) (*QualifiedE_Mpls_TeBandwidthType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_NamedExplicitPath_SidSelectionMode is a E_NamedExplicitPath_SidSelectionMode with a corresponding timestamp.
type QualifiedE_NamedExplicitPath_SidSelectionMode struct {
	*genutil.Metadata
	val     E_NamedExplicitPath_SidSelectionMode // val is the sample value.
	present bool
}

func (q *QualifiedE_NamedExplicitPath_SidSelectionMode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_NamedExplicitPath_SidSelectionMode sample, erroring out if not present.
func (q *QualifiedE_NamedExplicitPath_SidSelectionMode) Val(t testing.TB) E_NamedExplicitPath_SidSelectionMode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_NamedExplicitPath_SidSelectionMode sample.
func (q *QualifiedE_NamedExplicitPath_SidSelectionMode) SetVal(v E_NamedExplicitPath_SidSelectionMode) *QualifiedE_NamedExplicitPath_SidSelectionMode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_NamedExplicitPath_SidSelectionMode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_NamedExplicitPath_SidSelectionMode is a telemetry Collection whose Await method returns a slice of E_NamedExplicitPath_SidSelectionMode samples.
type CollectionE_NamedExplicitPath_SidSelectionMode struct {
	W    *E_NamedExplicitPath_SidSelectionModeWatcher
	Data []*QualifiedE_NamedExplicitPath_SidSelectionMode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_NamedExplicitPath_SidSelectionMode) Await(t testing.TB) []*QualifiedE_NamedExplicitPath_SidSelectionMode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_NamedExplicitPath_SidSelectionModeWatcher observes a stream of E_NamedExplicitPath_SidSelectionMode samples.
type E_NamedExplicitPath_SidSelectionModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_NamedExplicitPath_SidSelectionMode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_NamedExplicitPath_SidSelectionModeWatcher) Await(t testing.TB) (*QualifiedE_NamedExplicitPath_SidSelectionMode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Neighbor_NeighborState is a E_Neighbor_NeighborState with a corresponding timestamp.
type QualifiedE_Neighbor_NeighborState struct {
	*genutil.Metadata
	val     E_Neighbor_NeighborState // val is the sample value.
	present bool
}

func (q *QualifiedE_Neighbor_NeighborState) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Neighbor_NeighborState sample, erroring out if not present.
func (q *QualifiedE_Neighbor_NeighborState) Val(t testing.TB) E_Neighbor_NeighborState {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Neighbor_NeighborState sample.
func (q *QualifiedE_Neighbor_NeighborState) SetVal(v E_Neighbor_NeighborState) *QualifiedE_Neighbor_NeighborState {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Neighbor_NeighborState) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Neighbor_NeighborState is a telemetry Collection whose Await method returns a slice of E_Neighbor_NeighborState samples.
type CollectionE_Neighbor_NeighborState struct {
	W    *E_Neighbor_NeighborStateWatcher
	Data []*QualifiedE_Neighbor_NeighborState
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Neighbor_NeighborState) Await(t testing.TB) []*QualifiedE_Neighbor_NeighborState {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Neighbor_NeighborStateWatcher observes a stream of E_Neighbor_NeighborState samples.
type E_Neighbor_NeighborStateWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Neighbor_NeighborState
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Neighbor_NeighborStateWatcher) Await(t testing.TB) (*QualifiedE_Neighbor_NeighborState, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Neighbor_NeighborStatus is a E_Neighbor_NeighborStatus with a corresponding timestamp.
type QualifiedE_Neighbor_NeighborStatus struct {
	*genutil.Metadata
	val     E_Neighbor_NeighborStatus // val is the sample value.
	present bool
}

func (q *QualifiedE_Neighbor_NeighborStatus) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Neighbor_NeighborStatus sample, erroring out if not present.
func (q *QualifiedE_Neighbor_NeighborStatus) Val(t testing.TB) E_Neighbor_NeighborStatus {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Neighbor_NeighborStatus sample.
func (q *QualifiedE_Neighbor_NeighborStatus) SetVal(v E_Neighbor_NeighborStatus) *QualifiedE_Neighbor_NeighborStatus {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Neighbor_NeighborStatus) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Neighbor_NeighborStatus is a telemetry Collection whose Await method returns a slice of E_Neighbor_NeighborStatus samples.
type CollectionE_Neighbor_NeighborStatus struct {
	W    *E_Neighbor_NeighborStatusWatcher
	Data []*QualifiedE_Neighbor_NeighborStatus
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Neighbor_NeighborStatus) Await(t testing.TB) []*QualifiedE_Neighbor_NeighborStatus {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Neighbor_NeighborStatusWatcher observes a stream of E_Neighbor_NeighborStatus samples.
type E_Neighbor_NeighborStatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Neighbor_NeighborStatus
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Neighbor_NeighborStatusWatcher) Await(t testing.TB) (*QualifiedE_Neighbor_NeighborStatus, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_NetworkInstanceTypes_ENCAPSULATION is a E_NetworkInstanceTypes_ENCAPSULATION with a corresponding timestamp.
type QualifiedE_NetworkInstanceTypes_ENCAPSULATION struct {
	*genutil.Metadata
	val     E_NetworkInstanceTypes_ENCAPSULATION // val is the sample value.
	present bool
}

func (q *QualifiedE_NetworkInstanceTypes_ENCAPSULATION) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_NetworkInstanceTypes_ENCAPSULATION sample, erroring out if not present.
func (q *QualifiedE_NetworkInstanceTypes_ENCAPSULATION) Val(t testing.TB) E_NetworkInstanceTypes_ENCAPSULATION {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_NetworkInstanceTypes_ENCAPSULATION sample.
func (q *QualifiedE_NetworkInstanceTypes_ENCAPSULATION) SetVal(v E_NetworkInstanceTypes_ENCAPSULATION) *QualifiedE_NetworkInstanceTypes_ENCAPSULATION {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_NetworkInstanceTypes_ENCAPSULATION) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_NetworkInstanceTypes_ENCAPSULATION is a telemetry Collection whose Await method returns a slice of E_NetworkInstanceTypes_ENCAPSULATION samples.
type CollectionE_NetworkInstanceTypes_ENCAPSULATION struct {
	W    *E_NetworkInstanceTypes_ENCAPSULATIONWatcher
	Data []*QualifiedE_NetworkInstanceTypes_ENCAPSULATION
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_NetworkInstanceTypes_ENCAPSULATION) Await(t testing.TB) []*QualifiedE_NetworkInstanceTypes_ENCAPSULATION {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_NetworkInstanceTypes_ENCAPSULATIONWatcher observes a stream of E_NetworkInstanceTypes_ENCAPSULATION samples.
type E_NetworkInstanceTypes_ENCAPSULATIONWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_NetworkInstanceTypes_ENCAPSULATION
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_NetworkInstanceTypes_ENCAPSULATIONWatcher) Await(t testing.TB) (*QualifiedE_NetworkInstanceTypes_ENCAPSULATION, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE is a E_NetworkInstanceTypes_ENDPOINT_TYPE with a corresponding timestamp.
type QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE struct {
	*genutil.Metadata
	val     E_NetworkInstanceTypes_ENDPOINT_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_NetworkInstanceTypes_ENDPOINT_TYPE sample, erroring out if not present.
func (q *QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE) Val(t testing.TB) E_NetworkInstanceTypes_ENDPOINT_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_NetworkInstanceTypes_ENDPOINT_TYPE sample.
func (q *QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE) SetVal(v E_NetworkInstanceTypes_ENDPOINT_TYPE) *QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_NetworkInstanceTypes_ENDPOINT_TYPE is a telemetry Collection whose Await method returns a slice of E_NetworkInstanceTypes_ENDPOINT_TYPE samples.
type CollectionE_NetworkInstanceTypes_ENDPOINT_TYPE struct {
	W    *E_NetworkInstanceTypes_ENDPOINT_TYPEWatcher
	Data []*QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_NetworkInstanceTypes_ENDPOINT_TYPE) Await(t testing.TB) []*QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_NetworkInstanceTypes_ENDPOINT_TYPEWatcher observes a stream of E_NetworkInstanceTypes_ENDPOINT_TYPE samples.
type E_NetworkInstanceTypes_ENDPOINT_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_NetworkInstanceTypes_ENDPOINT_TYPEWatcher) Await(t testing.TB) (*QualifiedE_NetworkInstanceTypes_ENDPOINT_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE is a E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE with a corresponding timestamp.
type QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE struct {
	*genutil.Metadata
	val     E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE // val is the sample value.
	present bool
}

func (q *QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE sample, erroring out if not present.
func (q *QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE) Val(t testing.TB) E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE sample.
func (q *QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE) SetVal(v E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE) *QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE is a telemetry Collection whose Await method returns a slice of E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE samples.
type CollectionE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE struct {
	W    *E_NetworkInstanceTypes_LABEL_ALLOCATION_MODEWatcher
	Data []*QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE) Await(t testing.TB) []*QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_NetworkInstanceTypes_LABEL_ALLOCATION_MODEWatcher observes a stream of E_NetworkInstanceTypes_LABEL_ALLOCATION_MODE samples.
type E_NetworkInstanceTypes_LABEL_ALLOCATION_MODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_NetworkInstanceTypes_LABEL_ALLOCATION_MODEWatcher) Await(t testing.TB) (*QualifiedE_NetworkInstanceTypes_LABEL_ALLOCATION_MODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE is a E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE with a corresponding timestamp.
type QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE struct {
	*genutil.Metadata
	val     E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE sample, erroring out if not present.
func (q *QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE) Val(t testing.TB) E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE sample.
func (q *QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE) SetVal(v E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE) *QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE is a telemetry Collection whose Await method returns a slice of E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE samples.
type CollectionE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE struct {
	W    *E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPEWatcher
	Data []*QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE) Await(t testing.TB) []*QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPEWatcher observes a stream of E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE samples.
type E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_NetworkInstanceTypes_NETWORK_INSTANCE_TYPEWatcher) Await(t testing.TB) (*QualifiedE_NetworkInstanceTypes_NETWORK_INSTANCE_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OpaqueLsa_Scope is a E_OpaqueLsa_Scope with a corresponding timestamp.
type QualifiedE_OpaqueLsa_Scope struct {
	*genutil.Metadata
	val     E_OpaqueLsa_Scope // val is the sample value.
	present bool
}

func (q *QualifiedE_OpaqueLsa_Scope) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OpaqueLsa_Scope sample, erroring out if not present.
func (q *QualifiedE_OpaqueLsa_Scope) Val(t testing.TB) E_OpaqueLsa_Scope {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OpaqueLsa_Scope sample.
func (q *QualifiedE_OpaqueLsa_Scope) SetVal(v E_OpaqueLsa_Scope) *QualifiedE_OpaqueLsa_Scope {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OpaqueLsa_Scope) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OpaqueLsa_Scope is a telemetry Collection whose Await method returns a slice of E_OpaqueLsa_Scope samples.
type CollectionE_OpaqueLsa_Scope struct {
	W    *E_OpaqueLsa_ScopeWatcher
	Data []*QualifiedE_OpaqueLsa_Scope
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OpaqueLsa_Scope) Await(t testing.TB) []*QualifiedE_OpaqueLsa_Scope {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OpaqueLsa_ScopeWatcher observes a stream of E_OpaqueLsa_Scope samples.
type E_OpaqueLsa_ScopeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OpaqueLsa_Scope
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OpaqueLsa_ScopeWatcher) Await(t testing.TB) (*QualifiedE_OpaqueLsa_Scope, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES is a E_OspfTypes_GRACE_LSA_TLV_TYPES with a corresponding timestamp.
type QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES struct {
	*genutil.Metadata
	val     E_OspfTypes_GRACE_LSA_TLV_TYPES // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_GRACE_LSA_TLV_TYPES sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES) Val(t testing.TB) E_OspfTypes_GRACE_LSA_TLV_TYPES {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_GRACE_LSA_TLV_TYPES sample.
func (q *QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES) SetVal(v E_OspfTypes_GRACE_LSA_TLV_TYPES) *QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_GRACE_LSA_TLV_TYPES is a telemetry Collection whose Await method returns a slice of E_OspfTypes_GRACE_LSA_TLV_TYPES samples.
type CollectionE_OspfTypes_GRACE_LSA_TLV_TYPES struct {
	W    *E_OspfTypes_GRACE_LSA_TLV_TYPESWatcher
	Data []*QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_GRACE_LSA_TLV_TYPES) Await(t testing.TB) []*QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_GRACE_LSA_TLV_TYPESWatcher observes a stream of E_OspfTypes_GRACE_LSA_TLV_TYPES samples.
type E_OspfTypes_GRACE_LSA_TLV_TYPESWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_GRACE_LSA_TLV_TYPESWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_GRACE_LSA_TLV_TYPES, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE is a E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE) Val(t testing.TB) E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE) SetVal(v E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE) *QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE samples.
type CollectionE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE struct {
	W    *E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPEWatcher observes a stream of E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE samples.
type E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPFV2_EXTENDED_LINK_SUBTLV_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE is a E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE) Val(t testing.TB) E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE) SetVal(v E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE) *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE samples.
type CollectionE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE struct {
	W    *E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPEWatcher observes a stream of E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE samples.
type E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SID_LABEL_BINDING_SUBTLV_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE is a E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE) Val(t testing.TB) E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE) SetVal(v E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE) *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE samples.
type CollectionE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE struct {
	W    *E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPEWatcher observes a stream of E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE samples.
type E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPFV2_EXTENDED_PREFIX_SUBTLV_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE is a E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE) Val(t testing.TB) E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE) SetVal(v E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE) *QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE samples.
type CollectionE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE struct {
	W    *E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPEWatcher observes a stream of E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE samples.
type E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPFV2_EXTPREFIX_BINDING_ERO_PATH_SEGMENT_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE is a E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE) Val(t testing.TB) E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE) SetVal(v E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE) *QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE samples.
type CollectionE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE struct {
	W    *E_OspfTypes_OSPFV2_ROUTER_LINK_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPFV2_ROUTER_LINK_TYPEWatcher observes a stream of E_OspfTypes_OSPFV2_ROUTER_LINK_TYPE samples.
type E_OspfTypes_OSPFV2_ROUTER_LINK_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPFV2_ROUTER_LINK_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPFV2_ROUTER_LINK_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPF_LSA_TYPE is a E_OspfTypes_OSPF_LSA_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPF_LSA_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPF_LSA_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPF_LSA_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPF_LSA_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPF_LSA_TYPE) Val(t testing.TB) E_OspfTypes_OSPF_LSA_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPF_LSA_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPF_LSA_TYPE) SetVal(v E_OspfTypes_OSPF_LSA_TYPE) *QualifiedE_OspfTypes_OSPF_LSA_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPF_LSA_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPF_LSA_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPF_LSA_TYPE samples.
type CollectionE_OspfTypes_OSPF_LSA_TYPE struct {
	W    *E_OspfTypes_OSPF_LSA_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPF_LSA_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPF_LSA_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPF_LSA_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPF_LSA_TYPEWatcher observes a stream of E_OspfTypes_OSPF_LSA_TYPE samples.
type E_OspfTypes_OSPF_LSA_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPF_LSA_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPF_LSA_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPF_LSA_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE is a E_OspfTypes_OSPF_NEIGHBOR_STATE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPF_NEIGHBOR_STATE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPF_NEIGHBOR_STATE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE) Val(t testing.TB) E_OspfTypes_OSPF_NEIGHBOR_STATE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPF_NEIGHBOR_STATE sample.
func (q *QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE) SetVal(v E_OspfTypes_OSPF_NEIGHBOR_STATE) *QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPF_NEIGHBOR_STATE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPF_NEIGHBOR_STATE samples.
type CollectionE_OspfTypes_OSPF_NEIGHBOR_STATE struct {
	W    *E_OspfTypes_OSPF_NEIGHBOR_STATEWatcher
	Data []*QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPF_NEIGHBOR_STATE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPF_NEIGHBOR_STATEWatcher observes a stream of E_OspfTypes_OSPF_NEIGHBOR_STATE samples.
type E_OspfTypes_OSPF_NEIGHBOR_STATEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPF_NEIGHBOR_STATEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPF_NEIGHBOR_STATE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPF_NETWORK_TYPE is a E_OspfTypes_OSPF_NETWORK_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPF_NETWORK_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPF_NETWORK_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPF_NETWORK_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPF_NETWORK_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPF_NETWORK_TYPE) Val(t testing.TB) E_OspfTypes_OSPF_NETWORK_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPF_NETWORK_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPF_NETWORK_TYPE) SetVal(v E_OspfTypes_OSPF_NETWORK_TYPE) *QualifiedE_OspfTypes_OSPF_NETWORK_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPF_NETWORK_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPF_NETWORK_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPF_NETWORK_TYPE samples.
type CollectionE_OspfTypes_OSPF_NETWORK_TYPE struct {
	W    *E_OspfTypes_OSPF_NETWORK_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPF_NETWORK_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPF_NETWORK_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPF_NETWORK_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPF_NETWORK_TYPEWatcher observes a stream of E_OspfTypes_OSPF_NETWORK_TYPE samples.
type E_OspfTypes_OSPF_NETWORK_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPF_NETWORK_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPF_NETWORK_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPF_NETWORK_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE is a E_OspfTypes_OSPF_OPAQUE_LSA_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPF_OPAQUE_LSA_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPF_OPAQUE_LSA_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE) Val(t testing.TB) E_OspfTypes_OSPF_OPAQUE_LSA_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPF_OPAQUE_LSA_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE) SetVal(v E_OspfTypes_OSPF_OPAQUE_LSA_TYPE) *QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPF_OPAQUE_LSA_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPF_OPAQUE_LSA_TYPE samples.
type CollectionE_OspfTypes_OSPF_OPAQUE_LSA_TYPE struct {
	W    *E_OspfTypes_OSPF_OPAQUE_LSA_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPF_OPAQUE_LSA_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPF_OPAQUE_LSA_TYPEWatcher observes a stream of E_OspfTypes_OSPF_OPAQUE_LSA_TYPE samples.
type E_OspfTypes_OSPF_OPAQUE_LSA_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPF_OPAQUE_LSA_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPF_OPAQUE_LSA_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE is a E_OspfTypes_OSPF_TE_LSA_TLV_TYPE with a corresponding timestamp.
type QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE struct {
	*genutil.Metadata
	val     E_OspfTypes_OSPF_TE_LSA_TLV_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_OSPF_TE_LSA_TLV_TYPE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE) Val(t testing.TB) E_OspfTypes_OSPF_TE_LSA_TLV_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_OSPF_TE_LSA_TLV_TYPE sample.
func (q *QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE) SetVal(v E_OspfTypes_OSPF_TE_LSA_TLV_TYPE) *QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_OSPF_TE_LSA_TLV_TYPE is a telemetry Collection whose Await method returns a slice of E_OspfTypes_OSPF_TE_LSA_TLV_TYPE samples.
type CollectionE_OspfTypes_OSPF_TE_LSA_TLV_TYPE struct {
	W    *E_OspfTypes_OSPF_TE_LSA_TLV_TYPEWatcher
	Data []*QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_OSPF_TE_LSA_TLV_TYPE) Await(t testing.TB) []*QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_OSPF_TE_LSA_TLV_TYPEWatcher observes a stream of E_OspfTypes_OSPF_TE_LSA_TLV_TYPE samples.
type E_OspfTypes_OSPF_TE_LSA_TLV_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_OSPF_TE_LSA_TLV_TYPEWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_OSPF_TE_LSA_TLV_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_ROUTER_LSA_TYPES is a E_OspfTypes_ROUTER_LSA_TYPES with a corresponding timestamp.
type QualifiedE_OspfTypes_ROUTER_LSA_TYPES struct {
	*genutil.Metadata
	val     E_OspfTypes_ROUTER_LSA_TYPES // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_ROUTER_LSA_TYPES) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_ROUTER_LSA_TYPES sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_ROUTER_LSA_TYPES) Val(t testing.TB) E_OspfTypes_ROUTER_LSA_TYPES {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_ROUTER_LSA_TYPES sample.
func (q *QualifiedE_OspfTypes_ROUTER_LSA_TYPES) SetVal(v E_OspfTypes_ROUTER_LSA_TYPES) *QualifiedE_OspfTypes_ROUTER_LSA_TYPES {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_ROUTER_LSA_TYPES) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_ROUTER_LSA_TYPES is a telemetry Collection whose Await method returns a slice of E_OspfTypes_ROUTER_LSA_TYPES samples.
type CollectionE_OspfTypes_ROUTER_LSA_TYPES struct {
	W    *E_OspfTypes_ROUTER_LSA_TYPESWatcher
	Data []*QualifiedE_OspfTypes_ROUTER_LSA_TYPES
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_ROUTER_LSA_TYPES) Await(t testing.TB) []*QualifiedE_OspfTypes_ROUTER_LSA_TYPES {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_ROUTER_LSA_TYPESWatcher observes a stream of E_OspfTypes_ROUTER_LSA_TYPES samples.
type E_OspfTypes_ROUTER_LSA_TYPESWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_ROUTER_LSA_TYPES
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_ROUTER_LSA_TYPESWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_ROUTER_LSA_TYPES, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_SrSidType is a E_OspfTypes_SrSidType with a corresponding timestamp.
type QualifiedE_OspfTypes_SrSidType struct {
	*genutil.Metadata
	val     E_OspfTypes_SrSidType // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_SrSidType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OspfTypes_SrSidType sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_SrSidType) Val(t testing.TB) E_OspfTypes_SrSidType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OspfTypes_SrSidType sample.
func (q *QualifiedE_OspfTypes_SrSidType) SetVal(v E_OspfTypes_SrSidType) *QualifiedE_OspfTypes_SrSidType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_SrSidType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_SrSidType is a telemetry Collection whose Await method returns a slice of E_OspfTypes_SrSidType samples.
type CollectionE_OspfTypes_SrSidType struct {
	W    *E_OspfTypes_SrSidTypeWatcher
	Data []*QualifiedE_OspfTypes_SrSidType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_SrSidType) Await(t testing.TB) []*QualifiedE_OspfTypes_SrSidType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_SrSidTypeWatcher observes a stream of E_OspfTypes_SrSidType samples.
type E_OspfTypes_SrSidTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_SrSidType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_SrSidTypeWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_SrSidType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Output_OutputType is a E_Output_OutputType with a corresponding timestamp.
type QualifiedE_Output_OutputType struct {
	*genutil.Metadata
	val     E_Output_OutputType // val is the sample value.
	present bool
}

func (q *QualifiedE_Output_OutputType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Output_OutputType sample, erroring out if not present.
func (q *QualifiedE_Output_OutputType) Val(t testing.TB) E_Output_OutputType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Output_OutputType sample.
func (q *QualifiedE_Output_OutputType) SetVal(v E_Output_OutputType) *QualifiedE_Output_OutputType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Output_OutputType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Output_OutputType is a telemetry Collection whose Await method returns a slice of E_Output_OutputType samples.
type CollectionE_Output_OutputType struct {
	W    *E_Output_OutputTypeWatcher
	Data []*QualifiedE_Output_OutputType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Output_OutputType) Await(t testing.TB) []*QualifiedE_Output_OutputType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Output_OutputTypeWatcher observes a stream of E_Output_OutputType samples.
type E_Output_OutputTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Output_OutputType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Output_OutputTypeWatcher) Await(t testing.TB) (*QualifiedE_Output_OutputType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Pcep_LspControlType is a E_Pcep_LspControlType with a corresponding timestamp.
type QualifiedE_Pcep_LspControlType struct {
	*genutil.Metadata
	val     E_Pcep_LspControlType // val is the sample value.
	present bool
}

func (q *QualifiedE_Pcep_LspControlType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Pcep_LspControlType sample, erroring out if not present.
func (q *QualifiedE_Pcep_LspControlType) Val(t testing.TB) E_Pcep_LspControlType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Pcep_LspControlType sample.
func (q *QualifiedE_Pcep_LspControlType) SetVal(v E_Pcep_LspControlType) *QualifiedE_Pcep_LspControlType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Pcep_LspControlType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Pcep_LspControlType is a telemetry Collection whose Await method returns a slice of E_Pcep_LspControlType samples.
type CollectionE_Pcep_LspControlType struct {
	W    *E_Pcep_LspControlTypeWatcher
	Data []*QualifiedE_Pcep_LspControlType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Pcep_LspControlType) Await(t testing.TB) []*QualifiedE_Pcep_LspControlType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Pcep_LspControlTypeWatcher observes a stream of E_Pcep_LspControlType samples.
type E_Pcep_LspControlTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Pcep_LspControlType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Pcep_LspControlTypeWatcher) Await(t testing.TB) (*QualifiedE_Pcep_LspControlType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Pcep_PceModeType is a E_Pcep_PceModeType with a corresponding timestamp.
type QualifiedE_Pcep_PceModeType struct {
	*genutil.Metadata
	val     E_Pcep_PceModeType // val is the sample value.
	present bool
}

func (q *QualifiedE_Pcep_PceModeType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Pcep_PceModeType sample, erroring out if not present.
func (q *QualifiedE_Pcep_PceModeType) Val(t testing.TB) E_Pcep_PceModeType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Pcep_PceModeType sample.
func (q *QualifiedE_Pcep_PceModeType) SetVal(v E_Pcep_PceModeType) *QualifiedE_Pcep_PceModeType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Pcep_PceModeType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Pcep_PceModeType is a telemetry Collection whose Await method returns a slice of E_Pcep_PceModeType samples.
type CollectionE_Pcep_PceModeType struct {
	W    *E_Pcep_PceModeTypeWatcher
	Data []*QualifiedE_Pcep_PceModeType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Pcep_PceModeType) Await(t testing.TB) []*QualifiedE_Pcep_PceModeType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Pcep_PceModeTypeWatcher observes a stream of E_Pcep_PceModeType samples.
type E_Pcep_PceModeTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Pcep_PceModeType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Pcep_PceModeTypeWatcher) Await(t testing.TB) (*QualifiedE_Pcep_PceModeType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PimTypes_PIM_MODE is a E_PimTypes_PIM_MODE with a corresponding timestamp.
type QualifiedE_PimTypes_PIM_MODE struct {
	*genutil.Metadata
	val     E_PimTypes_PIM_MODE // val is the sample value.
	present bool
}

func (q *QualifiedE_PimTypes_PIM_MODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PimTypes_PIM_MODE sample, erroring out if not present.
func (q *QualifiedE_PimTypes_PIM_MODE) Val(t testing.TB) E_PimTypes_PIM_MODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PimTypes_PIM_MODE sample.
func (q *QualifiedE_PimTypes_PIM_MODE) SetVal(v E_PimTypes_PIM_MODE) *QualifiedE_PimTypes_PIM_MODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PimTypes_PIM_MODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PimTypes_PIM_MODE is a telemetry Collection whose Await method returns a slice of E_PimTypes_PIM_MODE samples.
type CollectionE_PimTypes_PIM_MODE struct {
	W    *E_PimTypes_PIM_MODEWatcher
	Data []*QualifiedE_PimTypes_PIM_MODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PimTypes_PIM_MODE) Await(t testing.TB) []*QualifiedE_PimTypes_PIM_MODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PimTypes_PIM_MODEWatcher observes a stream of E_PimTypes_PIM_MODE samples.
type E_PimTypes_PIM_MODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PimTypes_PIM_MODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PimTypes_PIM_MODEWatcher) Await(t testing.TB) (*QualifiedE_PimTypes_PIM_MODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE is a E_PlatformSoftware_SOFTWARE_MODULE_TYPE with a corresponding timestamp.
type QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE struct {
	*genutil.Metadata
	val     E_PlatformSoftware_SOFTWARE_MODULE_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PlatformSoftware_SOFTWARE_MODULE_TYPE sample, erroring out if not present.
func (q *QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) Val(t testing.TB) E_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PlatformSoftware_SOFTWARE_MODULE_TYPE sample.
func (q *QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) SetVal(v E_PlatformSoftware_SOFTWARE_MODULE_TYPE) *QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PlatformSoftware_SOFTWARE_MODULE_TYPE is a telemetry Collection whose Await method returns a slice of E_PlatformSoftware_SOFTWARE_MODULE_TYPE samples.
type CollectionE_PlatformSoftware_SOFTWARE_MODULE_TYPE struct {
	W    *E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher
	Data []*QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PlatformSoftware_SOFTWARE_MODULE_TYPE) Await(t testing.TB) []*QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher observes a stream of E_PlatformSoftware_SOFTWARE_MODULE_TYPE samples.
type E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PlatformSoftware_SOFTWARE_MODULE_TYPEWatcher) Await(t testing.TB) (*QualifiedE_PlatformSoftware_SOFTWARE_MODULE_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS is a E_PlatformTypes_COMPONENT_OPER_STATUS with a corresponding timestamp.
type QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS struct {
	*genutil.Metadata
	val     E_PlatformTypes_COMPONENT_OPER_STATUS // val is the sample value.
	present bool
}

func (q *QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PlatformTypes_COMPONENT_OPER_STATUS sample, erroring out if not present.
func (q *QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) Val(t testing.TB) E_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PlatformTypes_COMPONENT_OPER_STATUS sample.
func (q *QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) SetVal(v E_PlatformTypes_COMPONENT_OPER_STATUS) *QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PlatformTypes_COMPONENT_OPER_STATUS is a telemetry Collection whose Await method returns a slice of E_PlatformTypes_COMPONENT_OPER_STATUS samples.
type CollectionE_PlatformTypes_COMPONENT_OPER_STATUS struct {
	W    *E_PlatformTypes_COMPONENT_OPER_STATUSWatcher
	Data []*QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PlatformTypes_COMPONENT_OPER_STATUS) Await(t testing.TB) []*QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PlatformTypes_COMPONENT_OPER_STATUSWatcher observes a stream of E_PlatformTypes_COMPONENT_OPER_STATUS samples.
type E_PlatformTypes_COMPONENT_OPER_STATUSWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PlatformTypes_COMPONENT_OPER_STATUSWatcher) Await(t testing.TB) (*QualifiedE_PlatformTypes_COMPONENT_OPER_STATUS, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON is a E_PlatformTypes_COMPONENT_REBOOT_REASON with a corresponding timestamp.
type QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON struct {
	*genutil.Metadata
	val     E_PlatformTypes_COMPONENT_REBOOT_REASON // val is the sample value.
	present bool
}

func (q *QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PlatformTypes_COMPONENT_REBOOT_REASON sample, erroring out if not present.
func (q *QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) Val(t testing.TB) E_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PlatformTypes_COMPONENT_REBOOT_REASON sample.
func (q *QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) SetVal(v E_PlatformTypes_COMPONENT_REBOOT_REASON) *QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON is a telemetry Collection whose Await method returns a slice of E_PlatformTypes_COMPONENT_REBOOT_REASON samples.
type CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON struct {
	W    *E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher
	Data []*QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PlatformTypes_COMPONENT_REBOOT_REASON) Await(t testing.TB) []*QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher observes a stream of E_PlatformTypes_COMPONENT_REBOOT_REASON samples.
type E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PlatformTypes_COMPONENT_REBOOT_REASONWatcher) Await(t testing.TB) (*QualifiedE_PlatformTypes_COMPONENT_REBOOT_REASON, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PlatformTypes_ComponentRedundantRole is a E_PlatformTypes_ComponentRedundantRole with a corresponding timestamp.
type QualifiedE_PlatformTypes_ComponentRedundantRole struct {
	*genutil.Metadata
	val     E_PlatformTypes_ComponentRedundantRole // val is the sample value.
	present bool
}

func (q *QualifiedE_PlatformTypes_ComponentRedundantRole) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PlatformTypes_ComponentRedundantRole sample, erroring out if not present.
func (q *QualifiedE_PlatformTypes_ComponentRedundantRole) Val(t testing.TB) E_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PlatformTypes_ComponentRedundantRole sample.
func (q *QualifiedE_PlatformTypes_ComponentRedundantRole) SetVal(v E_PlatformTypes_ComponentRedundantRole) *QualifiedE_PlatformTypes_ComponentRedundantRole {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PlatformTypes_ComponentRedundantRole) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PlatformTypes_ComponentRedundantRole is a telemetry Collection whose Await method returns a slice of E_PlatformTypes_ComponentRedundantRole samples.
type CollectionE_PlatformTypes_ComponentRedundantRole struct {
	W    *E_PlatformTypes_ComponentRedundantRoleWatcher
	Data []*QualifiedE_PlatformTypes_ComponentRedundantRole
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PlatformTypes_ComponentRedundantRole) Await(t testing.TB) []*QualifiedE_PlatformTypes_ComponentRedundantRole {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PlatformTypes_ComponentRedundantRoleWatcher observes a stream of E_PlatformTypes_ComponentRedundantRole samples.
type E_PlatformTypes_ComponentRedundantRoleWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PlatformTypes_ComponentRedundantRole
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PlatformTypes_ComponentRedundantRoleWatcher) Await(t testing.TB) (*QualifiedE_PlatformTypes_ComponentRedundantRole, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger is a E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger with a corresponding timestamp.
type QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger struct {
	*genutil.Metadata
	val     E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger // val is the sample value.
	present bool
}

func (q *QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger sample, erroring out if not present.
func (q *QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) Val(t testing.TB) E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger sample.
func (q *QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) SetVal(v E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) *QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger is a telemetry Collection whose Await method returns a slice of E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger samples.
type CollectionE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger struct {
	W    *E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher
	Data []*QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger) Await(t testing.TB) []*QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher observes a stream of E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger samples.
type E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTriggerWatcher) Await(t testing.TB) (*QualifiedE_PlatformTypes_ComponentRedundantRoleSwitchoverReasonTrigger, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PlatformTypes_FEC_MODE_TYPE is a E_PlatformTypes_FEC_MODE_TYPE with a corresponding timestamp.
type QualifiedE_PlatformTypes_FEC_MODE_TYPE struct {
	*genutil.Metadata
	val     E_PlatformTypes_FEC_MODE_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_PlatformTypes_FEC_MODE_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PlatformTypes_FEC_MODE_TYPE sample, erroring out if not present.
func (q *QualifiedE_PlatformTypes_FEC_MODE_TYPE) Val(t testing.TB) E_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PlatformTypes_FEC_MODE_TYPE sample.
func (q *QualifiedE_PlatformTypes_FEC_MODE_TYPE) SetVal(v E_PlatformTypes_FEC_MODE_TYPE) *QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PlatformTypes_FEC_MODE_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PlatformTypes_FEC_MODE_TYPE is a telemetry Collection whose Await method returns a slice of E_PlatformTypes_FEC_MODE_TYPE samples.
type CollectionE_PlatformTypes_FEC_MODE_TYPE struct {
	W    *E_PlatformTypes_FEC_MODE_TYPEWatcher
	Data []*QualifiedE_PlatformTypes_FEC_MODE_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PlatformTypes_FEC_MODE_TYPE) Await(t testing.TB) []*QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PlatformTypes_FEC_MODE_TYPEWatcher observes a stream of E_PlatformTypes_FEC_MODE_TYPE samples.
type E_PlatformTypes_FEC_MODE_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PlatformTypes_FEC_MODE_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PlatformTypes_FEC_MODE_TYPEWatcher) Await(t testing.TB) (*QualifiedE_PlatformTypes_FEC_MODE_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PlatformTypes_FEC_STATUS_TYPE is a E_PlatformTypes_FEC_STATUS_TYPE with a corresponding timestamp.
type QualifiedE_PlatformTypes_FEC_STATUS_TYPE struct {
	*genutil.Metadata
	val     E_PlatformTypes_FEC_STATUS_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_PlatformTypes_FEC_STATUS_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PlatformTypes_FEC_STATUS_TYPE sample, erroring out if not present.
func (q *QualifiedE_PlatformTypes_FEC_STATUS_TYPE) Val(t testing.TB) E_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PlatformTypes_FEC_STATUS_TYPE sample.
func (q *QualifiedE_PlatformTypes_FEC_STATUS_TYPE) SetVal(v E_PlatformTypes_FEC_STATUS_TYPE) *QualifiedE_PlatformTypes_FEC_STATUS_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PlatformTypes_FEC_STATUS_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PlatformTypes_FEC_STATUS_TYPE is a telemetry Collection whose Await method returns a slice of E_PlatformTypes_FEC_STATUS_TYPE samples.
type CollectionE_PlatformTypes_FEC_STATUS_TYPE struct {
	W    *E_PlatformTypes_FEC_STATUS_TYPEWatcher
	Data []*QualifiedE_PlatformTypes_FEC_STATUS_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PlatformTypes_FEC_STATUS_TYPE) Await(t testing.TB) []*QualifiedE_PlatformTypes_FEC_STATUS_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PlatformTypes_FEC_STATUS_TYPEWatcher observes a stream of E_PlatformTypes_FEC_STATUS_TYPE samples.
type E_PlatformTypes_FEC_STATUS_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PlatformTypes_FEC_STATUS_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PlatformTypes_FEC_STATUS_TYPEWatcher) Await(t testing.TB) (*QualifiedE_PlatformTypes_FEC_STATUS_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON is a E_PolicyTypes_ATTRIBUTE_COMPARISON with a corresponding timestamp.
type QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON struct {
	*genutil.Metadata
	val     E_PolicyTypes_ATTRIBUTE_COMPARISON // val is the sample value.
	present bool
}

func (q *QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PolicyTypes_ATTRIBUTE_COMPARISON sample, erroring out if not present.
func (q *QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) Val(t testing.TB) E_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PolicyTypes_ATTRIBUTE_COMPARISON sample.
func (q *QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) SetVal(v E_PolicyTypes_ATTRIBUTE_COMPARISON) *QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON is a telemetry Collection whose Await method returns a slice of E_PolicyTypes_ATTRIBUTE_COMPARISON samples.
type CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON struct {
	W    *E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher
	Data []*QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PolicyTypes_ATTRIBUTE_COMPARISON) Await(t testing.TB) []*QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher observes a stream of E_PolicyTypes_ATTRIBUTE_COMPARISON samples.
type E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PolicyTypes_ATTRIBUTE_COMPARISONWatcher) Await(t testing.TB) (*QualifiedE_PolicyTypes_ATTRIBUTE_COMPARISON, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE is a E_PolicyTypes_INSTALL_PROTOCOL_TYPE with a corresponding timestamp.
type QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE struct {
	*genutil.Metadata
	val     E_PolicyTypes_INSTALL_PROTOCOL_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PolicyTypes_INSTALL_PROTOCOL_TYPE sample, erroring out if not present.
func (q *QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE) Val(t testing.TB) E_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PolicyTypes_INSTALL_PROTOCOL_TYPE sample.
func (q *QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE) SetVal(v E_PolicyTypes_INSTALL_PROTOCOL_TYPE) *QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PolicyTypes_INSTALL_PROTOCOL_TYPE is a telemetry Collection whose Await method returns a slice of E_PolicyTypes_INSTALL_PROTOCOL_TYPE samples.
type CollectionE_PolicyTypes_INSTALL_PROTOCOL_TYPE struct {
	W    *E_PolicyTypes_INSTALL_PROTOCOL_TYPEWatcher
	Data []*QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PolicyTypes_INSTALL_PROTOCOL_TYPE) Await(t testing.TB) []*QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PolicyTypes_INSTALL_PROTOCOL_TYPEWatcher observes a stream of E_PolicyTypes_INSTALL_PROTOCOL_TYPE samples.
type E_PolicyTypes_INSTALL_PROTOCOL_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PolicyTypes_INSTALL_PROTOCOL_TYPEWatcher) Await(t testing.TB) (*QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType is a E_PolicyTypes_MatchSetOptionsRestrictedType with a corresponding timestamp.
type QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType struct {
	*genutil.Metadata
	val     E_PolicyTypes_MatchSetOptionsRestrictedType // val is the sample value.
	present bool
}

func (q *QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PolicyTypes_MatchSetOptionsRestrictedType sample, erroring out if not present.
func (q *QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) Val(t testing.TB) E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PolicyTypes_MatchSetOptionsRestrictedType sample.
func (q *QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) SetVal(v E_PolicyTypes_MatchSetOptionsRestrictedType) *QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PolicyTypes_MatchSetOptionsRestrictedType is a telemetry Collection whose Await method returns a slice of E_PolicyTypes_MatchSetOptionsRestrictedType samples.
type CollectionE_PolicyTypes_MatchSetOptionsRestrictedType struct {
	W    *E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher
	Data []*QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PolicyTypes_MatchSetOptionsRestrictedType) Await(t testing.TB) []*QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher observes a stream of E_PolicyTypes_MatchSetOptionsRestrictedType samples.
type E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PolicyTypes_MatchSetOptionsRestrictedTypeWatcher) Await(t testing.TB) (*QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PolicyTypes_MatchSetOptionsType is a E_PolicyTypes_MatchSetOptionsType with a corresponding timestamp.
type QualifiedE_PolicyTypes_MatchSetOptionsType struct {
	*genutil.Metadata
	val     E_PolicyTypes_MatchSetOptionsType // val is the sample value.
	present bool
}

func (q *QualifiedE_PolicyTypes_MatchSetOptionsType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PolicyTypes_MatchSetOptionsType sample, erroring out if not present.
func (q *QualifiedE_PolicyTypes_MatchSetOptionsType) Val(t testing.TB) E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PolicyTypes_MatchSetOptionsType sample.
func (q *QualifiedE_PolicyTypes_MatchSetOptionsType) SetVal(v E_PolicyTypes_MatchSetOptionsType) *QualifiedE_PolicyTypes_MatchSetOptionsType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PolicyTypes_MatchSetOptionsType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PolicyTypes_MatchSetOptionsType is a telemetry Collection whose Await method returns a slice of E_PolicyTypes_MatchSetOptionsType samples.
type CollectionE_PolicyTypes_MatchSetOptionsType struct {
	W    *E_PolicyTypes_MatchSetOptionsTypeWatcher
	Data []*QualifiedE_PolicyTypes_MatchSetOptionsType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PolicyTypes_MatchSetOptionsType) Await(t testing.TB) []*QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PolicyTypes_MatchSetOptionsTypeWatcher observes a stream of E_PolicyTypes_MatchSetOptionsType samples.
type E_PolicyTypes_MatchSetOptionsTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PolicyTypes_MatchSetOptionsType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PolicyTypes_MatchSetOptionsTypeWatcher) Await(t testing.TB) (*QualifiedE_PolicyTypes_MatchSetOptionsType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Policy_Type is a E_Policy_Type with a corresponding timestamp.
type QualifiedE_Policy_Type struct {
	*genutil.Metadata
	val     E_Policy_Type // val is the sample value.
	present bool
}

func (q *QualifiedE_Policy_Type) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Policy_Type sample, erroring out if not present.
func (q *QualifiedE_Policy_Type) Val(t testing.TB) E_Policy_Type {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Policy_Type sample.
func (q *QualifiedE_Policy_Type) SetVal(v E_Policy_Type) *QualifiedE_Policy_Type {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Policy_Type) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Policy_Type is a telemetry Collection whose Await method returns a slice of E_Policy_Type samples.
type CollectionE_Policy_Type struct {
	W    *E_Policy_TypeWatcher
	Data []*QualifiedE_Policy_Type
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Policy_Type) Await(t testing.TB) []*QualifiedE_Policy_Type {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Policy_TypeWatcher observes a stream of E_Policy_Type samples.
type E_Policy_TypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Policy_Type
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Policy_TypeWatcher) Await(t testing.TB) (*QualifiedE_Policy_Type, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PrefixSet_Mode is a E_PrefixSet_Mode with a corresponding timestamp.
type QualifiedE_PrefixSet_Mode struct {
	*genutil.Metadata
	val     E_PrefixSet_Mode // val is the sample value.
	present bool
}

func (q *QualifiedE_PrefixSet_Mode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PrefixSet_Mode sample, erroring out if not present.
func (q *QualifiedE_PrefixSet_Mode) Val(t testing.TB) E_PrefixSet_Mode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PrefixSet_Mode sample.
func (q *QualifiedE_PrefixSet_Mode) SetVal(v E_PrefixSet_Mode) *QualifiedE_PrefixSet_Mode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PrefixSet_Mode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PrefixSet_Mode is a telemetry Collection whose Await method returns a slice of E_PrefixSet_Mode samples.
type CollectionE_PrefixSet_Mode struct {
	W    *E_PrefixSet_ModeWatcher
	Data []*QualifiedE_PrefixSet_Mode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PrefixSet_Mode) Await(t testing.TB) []*QualifiedE_PrefixSet_Mode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PrefixSet_ModeWatcher observes a stream of E_PrefixSet_Mode samples.
type E_PrefixSet_ModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PrefixSet_Mode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PrefixSet_ModeWatcher) Await(t testing.TB) (*QualifiedE_PrefixSet_Mode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PrefixSid_LabelOptions is a E_PrefixSid_LabelOptions with a corresponding timestamp.
type QualifiedE_PrefixSid_LabelOptions struct {
	*genutil.Metadata
	val     E_PrefixSid_LabelOptions // val is the sample value.
	present bool
}

func (q *QualifiedE_PrefixSid_LabelOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PrefixSid_LabelOptions sample, erroring out if not present.
func (q *QualifiedE_PrefixSid_LabelOptions) Val(t testing.TB) E_PrefixSid_LabelOptions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PrefixSid_LabelOptions sample.
func (q *QualifiedE_PrefixSid_LabelOptions) SetVal(v E_PrefixSid_LabelOptions) *QualifiedE_PrefixSid_LabelOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PrefixSid_LabelOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PrefixSid_LabelOptions is a telemetry Collection whose Await method returns a slice of E_PrefixSid_LabelOptions samples.
type CollectionE_PrefixSid_LabelOptions struct {
	W    *E_PrefixSid_LabelOptionsWatcher
	Data []*QualifiedE_PrefixSid_LabelOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PrefixSid_LabelOptions) Await(t testing.TB) []*QualifiedE_PrefixSid_LabelOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PrefixSid_LabelOptionsWatcher observes a stream of E_PrefixSid_LabelOptions samples.
type E_PrefixSid_LabelOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PrefixSid_LabelOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PrefixSid_LabelOptionsWatcher) Await(t testing.TB) (*QualifiedE_PrefixSid_LabelOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PrefixSid_SidScope is a E_PrefixSid_SidScope with a corresponding timestamp.
type QualifiedE_PrefixSid_SidScope struct {
	*genutil.Metadata
	val     E_PrefixSid_SidScope // val is the sample value.
	present bool
}

func (q *QualifiedE_PrefixSid_SidScope) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PrefixSid_SidScope sample, erroring out if not present.
func (q *QualifiedE_PrefixSid_SidScope) Val(t testing.TB) E_PrefixSid_SidScope {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PrefixSid_SidScope sample.
func (q *QualifiedE_PrefixSid_SidScope) SetVal(v E_PrefixSid_SidScope) *QualifiedE_PrefixSid_SidScope {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PrefixSid_SidScope) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PrefixSid_SidScope is a telemetry Collection whose Await method returns a slice of E_PrefixSid_SidScope samples.
type CollectionE_PrefixSid_SidScope struct {
	W    *E_PrefixSid_SidScopeWatcher
	Data []*QualifiedE_PrefixSid_SidScope
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PrefixSid_SidScope) Await(t testing.TB) []*QualifiedE_PrefixSid_SidScope {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PrefixSid_SidScopeWatcher observes a stream of E_PrefixSid_SidScope samples.
type E_PrefixSid_SidScopeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PrefixSid_SidScope
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PrefixSid_SidScopeWatcher) Await(t testing.TB) (*QualifiedE_PrefixSid_SidScope, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PrefixSid_SidValueType is a E_PrefixSid_SidValueType with a corresponding timestamp.
type QualifiedE_PrefixSid_SidValueType struct {
	*genutil.Metadata
	val     E_PrefixSid_SidValueType // val is the sample value.
	present bool
}

func (q *QualifiedE_PrefixSid_SidValueType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_PrefixSid_SidValueType sample, erroring out if not present.
func (q *QualifiedE_PrefixSid_SidValueType) Val(t testing.TB) E_PrefixSid_SidValueType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_PrefixSid_SidValueType sample.
func (q *QualifiedE_PrefixSid_SidValueType) SetVal(v E_PrefixSid_SidValueType) *QualifiedE_PrefixSid_SidValueType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PrefixSid_SidValueType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PrefixSid_SidValueType is a telemetry Collection whose Await method returns a slice of E_PrefixSid_SidValueType samples.
type CollectionE_PrefixSid_SidValueType struct {
	W    *E_PrefixSid_SidValueTypeWatcher
	Data []*QualifiedE_PrefixSid_SidValueType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PrefixSid_SidValueType) Await(t testing.TB) []*QualifiedE_PrefixSid_SidValueType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PrefixSid_SidValueTypeWatcher observes a stream of E_PrefixSid_SidValueType samples.
type E_PrefixSid_SidValueTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PrefixSid_SidValueType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PrefixSid_SidValueTypeWatcher) Await(t testing.TB) (*QualifiedE_PrefixSid_SidValueType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Producer_MobilityState is a E_Producer_MobilityState with a corresponding timestamp.
type QualifiedE_Producer_MobilityState struct {
	*genutil.Metadata
	val     E_Producer_MobilityState // val is the sample value.
	present bool
}

func (q *QualifiedE_Producer_MobilityState) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Producer_MobilityState sample, erroring out if not present.
func (q *QualifiedE_Producer_MobilityState) Val(t testing.TB) E_Producer_MobilityState {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Producer_MobilityState sample.
func (q *QualifiedE_Producer_MobilityState) SetVal(v E_Producer_MobilityState) *QualifiedE_Producer_MobilityState {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Producer_MobilityState) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Producer_MobilityState is a telemetry Collection whose Await method returns a slice of E_Producer_MobilityState samples.
type CollectionE_Producer_MobilityState struct {
	W    *E_Producer_MobilityStateWatcher
	Data []*QualifiedE_Producer_MobilityState
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Producer_MobilityState) Await(t testing.TB) []*QualifiedE_Producer_MobilityState {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Producer_MobilityStateWatcher observes a stream of E_Producer_MobilityState samples.
type E_Producer_MobilityStateWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Producer_MobilityState
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Producer_MobilityStateWatcher) Await(t testing.TB) (*QualifiedE_Producer_MobilityState, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Producer_Producer is a E_Producer_Producer with a corresponding timestamp.
type QualifiedE_Producer_Producer struct {
	*genutil.Metadata
	val     E_Producer_Producer // val is the sample value.
	present bool
}

func (q *QualifiedE_Producer_Producer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Producer_Producer sample, erroring out if not present.
func (q *QualifiedE_Producer_Producer) Val(t testing.TB) E_Producer_Producer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Producer_Producer sample.
func (q *QualifiedE_Producer_Producer) SetVal(v E_Producer_Producer) *QualifiedE_Producer_Producer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Producer_Producer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Producer_Producer is a telemetry Collection whose Await method returns a slice of E_Producer_Producer samples.
type CollectionE_Producer_Producer struct {
	W    *E_Producer_ProducerWatcher
	Data []*QualifiedE_Producer_Producer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Producer_Producer) Await(t testing.TB) []*QualifiedE_Producer_Producer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Producer_ProducerWatcher observes a stream of E_Producer_Producer samples.
type E_Producer_ProducerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Producer_Producer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Producer_ProducerWatcher) Await(t testing.TB) (*QualifiedE_Producer_Producer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_ProxyArp_Mode is a E_ProxyArp_Mode with a corresponding timestamp.
type QualifiedE_ProxyArp_Mode struct {
	*genutil.Metadata
	val     E_ProxyArp_Mode // val is the sample value.
	present bool
}

func (q *QualifiedE_ProxyArp_Mode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_ProxyArp_Mode sample, erroring out if not present.
func (q *QualifiedE_ProxyArp_Mode) Val(t testing.TB) E_ProxyArp_Mode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_ProxyArp_Mode sample.
func (q *QualifiedE_ProxyArp_Mode) SetVal(v E_ProxyArp_Mode) *QualifiedE_ProxyArp_Mode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_ProxyArp_Mode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_ProxyArp_Mode is a telemetry Collection whose Await method returns a slice of E_ProxyArp_Mode samples.
type CollectionE_ProxyArp_Mode struct {
	W    *E_ProxyArp_ModeWatcher
	Data []*QualifiedE_ProxyArp_Mode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_ProxyArp_Mode) Await(t testing.TB) []*QualifiedE_ProxyArp_Mode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_ProxyArp_ModeWatcher observes a stream of E_ProxyArp_Mode samples.
type E_ProxyArp_ModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_ProxyArp_Mode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_ProxyArp_ModeWatcher) Await(t testing.TB) (*QualifiedE_ProxyArp_Mode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_QosTypes_QOS_SCHEDULER_TYPE is a E_QosTypes_QOS_SCHEDULER_TYPE with a corresponding timestamp.
type QualifiedE_QosTypes_QOS_SCHEDULER_TYPE struct {
	*genutil.Metadata
	val     E_QosTypes_QOS_SCHEDULER_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_QosTypes_QOS_SCHEDULER_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_QosTypes_QOS_SCHEDULER_TYPE sample, erroring out if not present.
func (q *QualifiedE_QosTypes_QOS_SCHEDULER_TYPE) Val(t testing.TB) E_QosTypes_QOS_SCHEDULER_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_QosTypes_QOS_SCHEDULER_TYPE sample.
func (q *QualifiedE_QosTypes_QOS_SCHEDULER_TYPE) SetVal(v E_QosTypes_QOS_SCHEDULER_TYPE) *QualifiedE_QosTypes_QOS_SCHEDULER_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_QosTypes_QOS_SCHEDULER_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_QosTypes_QOS_SCHEDULER_TYPE is a telemetry Collection whose Await method returns a slice of E_QosTypes_QOS_SCHEDULER_TYPE samples.
type CollectionE_QosTypes_QOS_SCHEDULER_TYPE struct {
	W    *E_QosTypes_QOS_SCHEDULER_TYPEWatcher
	Data []*QualifiedE_QosTypes_QOS_SCHEDULER_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_QosTypes_QOS_SCHEDULER_TYPE) Await(t testing.TB) []*QualifiedE_QosTypes_QOS_SCHEDULER_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_QosTypes_QOS_SCHEDULER_TYPEWatcher observes a stream of E_QosTypes_QOS_SCHEDULER_TYPE samples.
type E_QosTypes_QOS_SCHEDULER_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_QosTypes_QOS_SCHEDULER_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_QosTypes_QOS_SCHEDULER_TYPEWatcher) Await(t testing.TB) (*QualifiedE_QosTypes_QOS_SCHEDULER_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_QosTypes_QueueBehavior is a E_QosTypes_QueueBehavior with a corresponding timestamp.
type QualifiedE_QosTypes_QueueBehavior struct {
	*genutil.Metadata
	val     E_QosTypes_QueueBehavior // val is the sample value.
	present bool
}

func (q *QualifiedE_QosTypes_QueueBehavior) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_QosTypes_QueueBehavior sample, erroring out if not present.
func (q *QualifiedE_QosTypes_QueueBehavior) Val(t testing.TB) E_QosTypes_QueueBehavior {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_QosTypes_QueueBehavior sample.
func (q *QualifiedE_QosTypes_QueueBehavior) SetVal(v E_QosTypes_QueueBehavior) *QualifiedE_QosTypes_QueueBehavior {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_QosTypes_QueueBehavior) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_QosTypes_QueueBehavior is a telemetry Collection whose Await method returns a slice of E_QosTypes_QueueBehavior samples.
type CollectionE_QosTypes_QueueBehavior struct {
	W    *E_QosTypes_QueueBehaviorWatcher
	Data []*QualifiedE_QosTypes_QueueBehavior
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_QosTypes_QueueBehavior) Await(t testing.TB) []*QualifiedE_QosTypes_QueueBehavior {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_QosTypes_QueueBehaviorWatcher observes a stream of E_QosTypes_QueueBehavior samples.
type E_QosTypes_QueueBehaviorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_QosTypes_QueueBehavior
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_QosTypes_QueueBehaviorWatcher) Await(t testing.TB) (*QualifiedE_QosTypes_QueueBehavior, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Qos_Classifier_Type is a E_Qos_Classifier_Type with a corresponding timestamp.
type QualifiedE_Qos_Classifier_Type struct {
	*genutil.Metadata
	val     E_Qos_Classifier_Type // val is the sample value.
	present bool
}

func (q *QualifiedE_Qos_Classifier_Type) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Qos_Classifier_Type sample, erroring out if not present.
func (q *QualifiedE_Qos_Classifier_Type) Val(t testing.TB) E_Qos_Classifier_Type {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Qos_Classifier_Type sample.
func (q *QualifiedE_Qos_Classifier_Type) SetVal(v E_Qos_Classifier_Type) *QualifiedE_Qos_Classifier_Type {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Qos_Classifier_Type) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Qos_Classifier_Type is a telemetry Collection whose Await method returns a slice of E_Qos_Classifier_Type samples.
type CollectionE_Qos_Classifier_Type struct {
	W    *E_Qos_Classifier_TypeWatcher
	Data []*QualifiedE_Qos_Classifier_Type
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Qos_Classifier_Type) Await(t testing.TB) []*QualifiedE_Qos_Classifier_Type {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Qos_Classifier_TypeWatcher observes a stream of E_Qos_Classifier_Type samples.
type E_Qos_Classifier_TypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Qos_Classifier_Type
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Qos_Classifier_TypeWatcher) Await(t testing.TB) (*QualifiedE_Qos_Classifier_Type, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE is a E_Qos_SHARED_BUFFER_LIMIT_TYPE with a corresponding timestamp.
type QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE struct {
	*genutil.Metadata
	val     E_Qos_SHARED_BUFFER_LIMIT_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Qos_SHARED_BUFFER_LIMIT_TYPE sample, erroring out if not present.
func (q *QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) Val(t testing.TB) E_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Qos_SHARED_BUFFER_LIMIT_TYPE sample.
func (q *QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) SetVal(v E_Qos_SHARED_BUFFER_LIMIT_TYPE) *QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Qos_SHARED_BUFFER_LIMIT_TYPE is a telemetry Collection whose Await method returns a slice of E_Qos_SHARED_BUFFER_LIMIT_TYPE samples.
type CollectionE_Qos_SHARED_BUFFER_LIMIT_TYPE struct {
	W    *E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher
	Data []*QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Qos_SHARED_BUFFER_LIMIT_TYPE) Await(t testing.TB) []*QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher observes a stream of E_Qos_SHARED_BUFFER_LIMIT_TYPE samples.
type E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Qos_SHARED_BUFFER_LIMIT_TYPEWatcher) Await(t testing.TB) (*QualifiedE_Qos_SHARED_BUFFER_LIMIT_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON is a E_RibBgpTypes_INVALID_ROUTE_REASON with a corresponding timestamp.
type QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON struct {
	*genutil.Metadata
	val     E_RibBgpTypes_INVALID_ROUTE_REASON // val is the sample value.
	present bool
}

func (q *QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_RibBgpTypes_INVALID_ROUTE_REASON sample, erroring out if not present.
func (q *QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON) Val(t testing.TB) E_RibBgpTypes_INVALID_ROUTE_REASON {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_RibBgpTypes_INVALID_ROUTE_REASON sample.
func (q *QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON) SetVal(v E_RibBgpTypes_INVALID_ROUTE_REASON) *QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_RibBgpTypes_INVALID_ROUTE_REASON is a telemetry Collection whose Await method returns a slice of E_RibBgpTypes_INVALID_ROUTE_REASON samples.
type CollectionE_RibBgpTypes_INVALID_ROUTE_REASON struct {
	W    *E_RibBgpTypes_INVALID_ROUTE_REASONWatcher
	Data []*QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_RibBgpTypes_INVALID_ROUTE_REASON) Await(t testing.TB) []*QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_RibBgpTypes_INVALID_ROUTE_REASONWatcher observes a stream of E_RibBgpTypes_INVALID_ROUTE_REASON samples.
type E_RibBgpTypes_INVALID_ROUTE_REASONWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_RibBgpTypes_INVALID_ROUTE_REASONWatcher) Await(t testing.TB) (*QualifiedE_RibBgpTypes_INVALID_ROUTE_REASON, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE is a E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE with a corresponding timestamp.
type QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE struct {
	*genutil.Metadata
	val     E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE sample, erroring out if not present.
func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE) Val(t testing.TB) E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE sample.
func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE) SetVal(v E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE) *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE is a telemetry Collection whose Await method returns a slice of E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE samples.
type CollectionE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE struct {
	W    *E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPEWatcher
	Data []*QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE) Await(t testing.TB) []*QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPEWatcher observes a stream of E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE samples.
type E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPEWatcher) Await(t testing.TB) (*QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_SUBTLV_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE is a E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE with a corresponding timestamp.
type QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE struct {
	*genutil.Metadata
	val     E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE sample, erroring out if not present.
func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE) Val(t testing.TB) E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE sample.
func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE) SetVal(v E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE) *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE is a telemetry Collection whose Await method returns a slice of E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE samples.
type CollectionE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE struct {
	W    *E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPEWatcher
	Data []*QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE) Await(t testing.TB) []*QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPEWatcher observes a stream of E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE samples.
type E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPEWatcher) Await(t testing.TB) (*QualifiedE_RibBgpTypes_TUNNEL_ENCAPSULATION_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_RoutingPolicy_DefaultPolicyType is a E_RoutingPolicy_DefaultPolicyType with a corresponding timestamp.
type QualifiedE_RoutingPolicy_DefaultPolicyType struct {
	*genutil.Metadata
	val     E_RoutingPolicy_DefaultPolicyType // val is the sample value.
	present bool
}

func (q *QualifiedE_RoutingPolicy_DefaultPolicyType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_RoutingPolicy_DefaultPolicyType sample, erroring out if not present.
func (q *QualifiedE_RoutingPolicy_DefaultPolicyType) Val(t testing.TB) E_RoutingPolicy_DefaultPolicyType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_RoutingPolicy_DefaultPolicyType sample.
func (q *QualifiedE_RoutingPolicy_DefaultPolicyType) SetVal(v E_RoutingPolicy_DefaultPolicyType) *QualifiedE_RoutingPolicy_DefaultPolicyType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_RoutingPolicy_DefaultPolicyType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_RoutingPolicy_DefaultPolicyType is a telemetry Collection whose Await method returns a slice of E_RoutingPolicy_DefaultPolicyType samples.
type CollectionE_RoutingPolicy_DefaultPolicyType struct {
	W    *E_RoutingPolicy_DefaultPolicyTypeWatcher
	Data []*QualifiedE_RoutingPolicy_DefaultPolicyType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_RoutingPolicy_DefaultPolicyType) Await(t testing.TB) []*QualifiedE_RoutingPolicy_DefaultPolicyType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_RoutingPolicy_DefaultPolicyTypeWatcher observes a stream of E_RoutingPolicy_DefaultPolicyType samples.
type E_RoutingPolicy_DefaultPolicyTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_RoutingPolicy_DefaultPolicyType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_RoutingPolicy_DefaultPolicyTypeWatcher) Await(t testing.TB) (*QualifiedE_RoutingPolicy_DefaultPolicyType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_RoutingPolicy_PolicyResultType is a E_RoutingPolicy_PolicyResultType with a corresponding timestamp.
type QualifiedE_RoutingPolicy_PolicyResultType struct {
	*genutil.Metadata
	val     E_RoutingPolicy_PolicyResultType // val is the sample value.
	present bool
}

func (q *QualifiedE_RoutingPolicy_PolicyResultType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_RoutingPolicy_PolicyResultType sample, erroring out if not present.
func (q *QualifiedE_RoutingPolicy_PolicyResultType) Val(t testing.TB) E_RoutingPolicy_PolicyResultType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_RoutingPolicy_PolicyResultType sample.
func (q *QualifiedE_RoutingPolicy_PolicyResultType) SetVal(v E_RoutingPolicy_PolicyResultType) *QualifiedE_RoutingPolicy_PolicyResultType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_RoutingPolicy_PolicyResultType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_RoutingPolicy_PolicyResultType is a telemetry Collection whose Await method returns a slice of E_RoutingPolicy_PolicyResultType samples.
type CollectionE_RoutingPolicy_PolicyResultType struct {
	W    *E_RoutingPolicy_PolicyResultTypeWatcher
	Data []*QualifiedE_RoutingPolicy_PolicyResultType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_RoutingPolicy_PolicyResultType) Await(t testing.TB) []*QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_RoutingPolicy_PolicyResultTypeWatcher observes a stream of E_RoutingPolicy_PolicyResultType samples.
type E_RoutingPolicy_PolicyResultTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_RoutingPolicy_PolicyResultType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_RoutingPolicy_PolicyResultTypeWatcher) Await(t testing.TB) (*QualifiedE_RoutingPolicy_PolicyResultType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Scheduler_Priority is a E_Scheduler_Priority with a corresponding timestamp.
type QualifiedE_Scheduler_Priority struct {
	*genutil.Metadata
	val     E_Scheduler_Priority // val is the sample value.
	present bool
}

func (q *QualifiedE_Scheduler_Priority) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Scheduler_Priority sample, erroring out if not present.
func (q *QualifiedE_Scheduler_Priority) Val(t testing.TB) E_Scheduler_Priority {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Scheduler_Priority sample.
func (q *QualifiedE_Scheduler_Priority) SetVal(v E_Scheduler_Priority) *QualifiedE_Scheduler_Priority {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Scheduler_Priority) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Scheduler_Priority is a telemetry Collection whose Await method returns a slice of E_Scheduler_Priority samples.
type CollectionE_Scheduler_Priority struct {
	W    *E_Scheduler_PriorityWatcher
	Data []*QualifiedE_Scheduler_Priority
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Scheduler_Priority) Await(t testing.TB) []*QualifiedE_Scheduler_Priority {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Scheduler_PriorityWatcher observes a stream of E_Scheduler_Priority samples.
type E_Scheduler_PriorityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Scheduler_Priority
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Scheduler_PriorityWatcher) Await(t testing.TB) (*QualifiedE_Scheduler_Priority, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SegmentRoutingTypes_EnlpType is a E_SegmentRoutingTypes_EnlpType with a corresponding timestamp.
type QualifiedE_SegmentRoutingTypes_EnlpType struct {
	*genutil.Metadata
	val     E_SegmentRoutingTypes_EnlpType // val is the sample value.
	present bool
}

func (q *QualifiedE_SegmentRoutingTypes_EnlpType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SegmentRoutingTypes_EnlpType sample, erroring out if not present.
func (q *QualifiedE_SegmentRoutingTypes_EnlpType) Val(t testing.TB) E_SegmentRoutingTypes_EnlpType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SegmentRoutingTypes_EnlpType sample.
func (q *QualifiedE_SegmentRoutingTypes_EnlpType) SetVal(v E_SegmentRoutingTypes_EnlpType) *QualifiedE_SegmentRoutingTypes_EnlpType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SegmentRoutingTypes_EnlpType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SegmentRoutingTypes_EnlpType is a telemetry Collection whose Await method returns a slice of E_SegmentRoutingTypes_EnlpType samples.
type CollectionE_SegmentRoutingTypes_EnlpType struct {
	W    *E_SegmentRoutingTypes_EnlpTypeWatcher
	Data []*QualifiedE_SegmentRoutingTypes_EnlpType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SegmentRoutingTypes_EnlpType) Await(t testing.TB) []*QualifiedE_SegmentRoutingTypes_EnlpType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SegmentRoutingTypes_EnlpTypeWatcher observes a stream of E_SegmentRoutingTypes_EnlpType samples.
type E_SegmentRoutingTypes_EnlpTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SegmentRoutingTypes_EnlpType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SegmentRoutingTypes_EnlpTypeWatcher) Await(t testing.TB) (*QualifiedE_SegmentRoutingTypes_EnlpType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SegmentRoutingTypes_SrDataplaneType is a E_SegmentRoutingTypes_SrDataplaneType with a corresponding timestamp.
type QualifiedE_SegmentRoutingTypes_SrDataplaneType struct {
	*genutil.Metadata
	val     E_SegmentRoutingTypes_SrDataplaneType // val is the sample value.
	present bool
}

func (q *QualifiedE_SegmentRoutingTypes_SrDataplaneType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SegmentRoutingTypes_SrDataplaneType sample, erroring out if not present.
func (q *QualifiedE_SegmentRoutingTypes_SrDataplaneType) Val(t testing.TB) E_SegmentRoutingTypes_SrDataplaneType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SegmentRoutingTypes_SrDataplaneType sample.
func (q *QualifiedE_SegmentRoutingTypes_SrDataplaneType) SetVal(v E_SegmentRoutingTypes_SrDataplaneType) *QualifiedE_SegmentRoutingTypes_SrDataplaneType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SegmentRoutingTypes_SrDataplaneType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SegmentRoutingTypes_SrDataplaneType is a telemetry Collection whose Await method returns a slice of E_SegmentRoutingTypes_SrDataplaneType samples.
type CollectionE_SegmentRoutingTypes_SrDataplaneType struct {
	W    *E_SegmentRoutingTypes_SrDataplaneTypeWatcher
	Data []*QualifiedE_SegmentRoutingTypes_SrDataplaneType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SegmentRoutingTypes_SrDataplaneType) Await(t testing.TB) []*QualifiedE_SegmentRoutingTypes_SrDataplaneType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SegmentRoutingTypes_SrDataplaneTypeWatcher observes a stream of E_SegmentRoutingTypes_SrDataplaneType samples.
type E_SegmentRoutingTypes_SrDataplaneTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SegmentRoutingTypes_SrDataplaneType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SegmentRoutingTypes_SrDataplaneTypeWatcher) Await(t testing.TB) (*QualifiedE_SegmentRoutingTypes_SrDataplaneType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason is a E_SegmentRoutingTypes_SrteInvalidSlReason with a corresponding timestamp.
type QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason struct {
	*genutil.Metadata
	val     E_SegmentRoutingTypes_SrteInvalidSlReason // val is the sample value.
	present bool
}

func (q *QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SegmentRoutingTypes_SrteInvalidSlReason sample, erroring out if not present.
func (q *QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason) Val(t testing.TB) E_SegmentRoutingTypes_SrteInvalidSlReason {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SegmentRoutingTypes_SrteInvalidSlReason sample.
func (q *QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason) SetVal(v E_SegmentRoutingTypes_SrteInvalidSlReason) *QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SegmentRoutingTypes_SrteInvalidSlReason is a telemetry Collection whose Await method returns a slice of E_SegmentRoutingTypes_SrteInvalidSlReason samples.
type CollectionE_SegmentRoutingTypes_SrteInvalidSlReason struct {
	W    *E_SegmentRoutingTypes_SrteInvalidSlReasonWatcher
	Data []*QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SegmentRoutingTypes_SrteInvalidSlReason) Await(t testing.TB) []*QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SegmentRoutingTypes_SrteInvalidSlReasonWatcher observes a stream of E_SegmentRoutingTypes_SrteInvalidSlReason samples.
type E_SegmentRoutingTypes_SrteInvalidSlReasonWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SegmentRoutingTypes_SrteInvalidSlReasonWatcher) Await(t testing.TB) (*QualifiedE_SegmentRoutingTypes_SrteInvalidSlReason, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SegmentRoutingTypes_SrteProtocolType is a E_SegmentRoutingTypes_SrteProtocolType with a corresponding timestamp.
type QualifiedE_SegmentRoutingTypes_SrteProtocolType struct {
	*genutil.Metadata
	val     E_SegmentRoutingTypes_SrteProtocolType // val is the sample value.
	present bool
}

func (q *QualifiedE_SegmentRoutingTypes_SrteProtocolType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SegmentRoutingTypes_SrteProtocolType sample, erroring out if not present.
func (q *QualifiedE_SegmentRoutingTypes_SrteProtocolType) Val(t testing.TB) E_SegmentRoutingTypes_SrteProtocolType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SegmentRoutingTypes_SrteProtocolType sample.
func (q *QualifiedE_SegmentRoutingTypes_SrteProtocolType) SetVal(v E_SegmentRoutingTypes_SrteProtocolType) *QualifiedE_SegmentRoutingTypes_SrteProtocolType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SegmentRoutingTypes_SrteProtocolType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SegmentRoutingTypes_SrteProtocolType is a telemetry Collection whose Await method returns a slice of E_SegmentRoutingTypes_SrteProtocolType samples.
type CollectionE_SegmentRoutingTypes_SrteProtocolType struct {
	W    *E_SegmentRoutingTypes_SrteProtocolTypeWatcher
	Data []*QualifiedE_SegmentRoutingTypes_SrteProtocolType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SegmentRoutingTypes_SrteProtocolType) Await(t testing.TB) []*QualifiedE_SegmentRoutingTypes_SrteProtocolType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SegmentRoutingTypes_SrteProtocolTypeWatcher observes a stream of E_SegmentRoutingTypes_SrteProtocolType samples.
type E_SegmentRoutingTypes_SrteProtocolTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SegmentRoutingTypes_SrteProtocolType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SegmentRoutingTypes_SrteProtocolTypeWatcher) Await(t testing.TB) (*QualifiedE_SegmentRoutingTypes_SrteProtocolType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Segment_Type is a E_Segment_Type with a corresponding timestamp.
type QualifiedE_Segment_Type struct {
	*genutil.Metadata
	val     E_Segment_Type // val is the sample value.
	present bool
}

func (q *QualifiedE_Segment_Type) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Segment_Type sample, erroring out if not present.
func (q *QualifiedE_Segment_Type) Val(t testing.TB) E_Segment_Type {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Segment_Type sample.
func (q *QualifiedE_Segment_Type) SetVal(v E_Segment_Type) *QualifiedE_Segment_Type {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Segment_Type) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Segment_Type is a telemetry Collection whose Await method returns a slice of E_Segment_Type samples.
type CollectionE_Segment_Type struct {
	W    *E_Segment_TypeWatcher
	Data []*QualifiedE_Segment_Type
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Segment_Type) Await(t testing.TB) []*QualifiedE_Segment_Type {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Segment_TypeWatcher observes a stream of E_Segment_Type samples.
type E_Segment_TypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Segment_Type
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Segment_TypeWatcher) Await(t testing.TB) (*QualifiedE_Segment_Type, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Server_AssociationType is a E_Server_AssociationType with a corresponding timestamp.
type QualifiedE_Server_AssociationType struct {
	*genutil.Metadata
	val     E_Server_AssociationType // val is the sample value.
	present bool
}

func (q *QualifiedE_Server_AssociationType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Server_AssociationType sample, erroring out if not present.
func (q *QualifiedE_Server_AssociationType) Val(t testing.TB) E_Server_AssociationType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Server_AssociationType sample.
func (q *QualifiedE_Server_AssociationType) SetVal(v E_Server_AssociationType) *QualifiedE_Server_AssociationType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Server_AssociationType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Server_AssociationType is a telemetry Collection whose Await method returns a slice of E_Server_AssociationType samples.
type CollectionE_Server_AssociationType struct {
	W    *E_Server_AssociationTypeWatcher
	Data []*QualifiedE_Server_AssociationType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Server_AssociationType) Await(t testing.TB) []*QualifiedE_Server_AssociationType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Server_AssociationTypeWatcher observes a stream of E_Server_AssociationType samples.
type E_Server_AssociationTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Server_AssociationType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Server_AssociationTypeWatcher) Await(t testing.TB) (*QualifiedE_Server_AssociationType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Session_Status is a E_Session_Status with a corresponding timestamp.
type QualifiedE_Session_Status struct {
	*genutil.Metadata
	val     E_Session_Status // val is the sample value.
	present bool
}

func (q *QualifiedE_Session_Status) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Session_Status sample, erroring out if not present.
func (q *QualifiedE_Session_Status) Val(t testing.TB) E_Session_Status {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Session_Status sample.
func (q *QualifiedE_Session_Status) SetVal(v E_Session_Status) *QualifiedE_Session_Status {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Session_Status) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Session_Status is a telemetry Collection whose Await method returns a slice of E_Session_Status samples.
type CollectionE_Session_Status struct {
	W    *E_Session_StatusWatcher
	Data []*QualifiedE_Session_Status
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Session_Status) Await(t testing.TB) []*QualifiedE_Session_Status {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Session_StatusWatcher observes a stream of E_Session_Status samples.
type E_Session_StatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Session_Status
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Session_StatusWatcher) Await(t testing.TB) (*QualifiedE_Session_Status, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SetCommunity_Method is a E_SetCommunity_Method with a corresponding timestamp.
type QualifiedE_SetCommunity_Method struct {
	*genutil.Metadata
	val     E_SetCommunity_Method // val is the sample value.
	present bool
}

func (q *QualifiedE_SetCommunity_Method) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SetCommunity_Method sample, erroring out if not present.
func (q *QualifiedE_SetCommunity_Method) Val(t testing.TB) E_SetCommunity_Method {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SetCommunity_Method sample.
func (q *QualifiedE_SetCommunity_Method) SetVal(v E_SetCommunity_Method) *QualifiedE_SetCommunity_Method {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SetCommunity_Method) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SetCommunity_Method is a telemetry Collection whose Await method returns a slice of E_SetCommunity_Method samples.
type CollectionE_SetCommunity_Method struct {
	W    *E_SetCommunity_MethodWatcher
	Data []*QualifiedE_SetCommunity_Method
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SetCommunity_Method) Await(t testing.TB) []*QualifiedE_SetCommunity_Method {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SetCommunity_MethodWatcher observes a stream of E_SetCommunity_Method samples.
type E_SetCommunity_MethodWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SetCommunity_Method
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SetCommunity_MethodWatcher) Await(t testing.TB) (*QualifiedE_SetCommunity_Method, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SetTag_Mode is a E_SetTag_Mode with a corresponding timestamp.
type QualifiedE_SetTag_Mode struct {
	*genutil.Metadata
	val     E_SetTag_Mode // val is the sample value.
	present bool
}

func (q *QualifiedE_SetTag_Mode) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SetTag_Mode sample, erroring out if not present.
func (q *QualifiedE_SetTag_Mode) Val(t testing.TB) E_SetTag_Mode {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SetTag_Mode sample.
func (q *QualifiedE_SetTag_Mode) SetVal(v E_SetTag_Mode) *QualifiedE_SetTag_Mode {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SetTag_Mode) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SetTag_Mode is a telemetry Collection whose Await method returns a slice of E_SetTag_Mode samples.
type CollectionE_SetTag_Mode struct {
	W    *E_SetTag_ModeWatcher
	Data []*QualifiedE_SetTag_Mode
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SetTag_Mode) Await(t testing.TB) []*QualifiedE_SetTag_Mode {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SetTag_ModeWatcher observes a stream of E_SetTag_Mode samples.
type E_SetTag_ModeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SetTag_Mode
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SetTag_ModeWatcher) Await(t testing.TB) (*QualifiedE_SetTag_Mode, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SshServer_ProtocolVersion is a E_SshServer_ProtocolVersion with a corresponding timestamp.
type QualifiedE_SshServer_ProtocolVersion struct {
	*genutil.Metadata
	val     E_SshServer_ProtocolVersion // val is the sample value.
	present bool
}

func (q *QualifiedE_SshServer_ProtocolVersion) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SshServer_ProtocolVersion sample, erroring out if not present.
func (q *QualifiedE_SshServer_ProtocolVersion) Val(t testing.TB) E_SshServer_ProtocolVersion {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SshServer_ProtocolVersion sample.
func (q *QualifiedE_SshServer_ProtocolVersion) SetVal(v E_SshServer_ProtocolVersion) *QualifiedE_SshServer_ProtocolVersion {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SshServer_ProtocolVersion) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SshServer_ProtocolVersion is a telemetry Collection whose Await method returns a slice of E_SshServer_ProtocolVersion samples.
type CollectionE_SshServer_ProtocolVersion struct {
	W    *E_SshServer_ProtocolVersionWatcher
	Data []*QualifiedE_SshServer_ProtocolVersion
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SshServer_ProtocolVersion) Await(t testing.TB) []*QualifiedE_SshServer_ProtocolVersion {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SshServer_ProtocolVersionWatcher observes a stream of E_SshServer_ProtocolVersion samples.
type E_SshServer_ProtocolVersionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SshServer_ProtocolVersion
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SshServer_ProtocolVersionWatcher) Await(t testing.TB) (*QualifiedE_SshServer_ProtocolVersion, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SubTlv_LinkType is a E_SubTlv_LinkType with a corresponding timestamp.
type QualifiedE_SubTlv_LinkType struct {
	*genutil.Metadata
	val     E_SubTlv_LinkType // val is the sample value.
	present bool
}

func (q *QualifiedE_SubTlv_LinkType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SubTlv_LinkType sample, erroring out if not present.
func (q *QualifiedE_SubTlv_LinkType) Val(t testing.TB) E_SubTlv_LinkType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SubTlv_LinkType sample.
func (q *QualifiedE_SubTlv_LinkType) SetVal(v E_SubTlv_LinkType) *QualifiedE_SubTlv_LinkType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SubTlv_LinkType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SubTlv_LinkType is a telemetry Collection whose Await method returns a slice of E_SubTlv_LinkType samples.
type CollectionE_SubTlv_LinkType struct {
	W    *E_SubTlv_LinkTypeWatcher
	Data []*QualifiedE_SubTlv_LinkType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SubTlv_LinkType) Await(t testing.TB) []*QualifiedE_SubTlv_LinkType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SubTlv_LinkTypeWatcher observes a stream of E_SubTlv_LinkType samples.
type E_SubTlv_LinkTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SubTlv_LinkType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SubTlv_LinkTypeWatcher) Await(t testing.TB) (*QualifiedE_SubTlv_LinkType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SystemLogging_SYSLOG_FACILITY is a E_SystemLogging_SYSLOG_FACILITY with a corresponding timestamp.
type QualifiedE_SystemLogging_SYSLOG_FACILITY struct {
	*genutil.Metadata
	val     E_SystemLogging_SYSLOG_FACILITY // val is the sample value.
	present bool
}

func (q *QualifiedE_SystemLogging_SYSLOG_FACILITY) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SystemLogging_SYSLOG_FACILITY sample, erroring out if not present.
func (q *QualifiedE_SystemLogging_SYSLOG_FACILITY) Val(t testing.TB) E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SystemLogging_SYSLOG_FACILITY sample.
func (q *QualifiedE_SystemLogging_SYSLOG_FACILITY) SetVal(v E_SystemLogging_SYSLOG_FACILITY) *QualifiedE_SystemLogging_SYSLOG_FACILITY {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SystemLogging_SYSLOG_FACILITY) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SystemLogging_SYSLOG_FACILITY is a telemetry Collection whose Await method returns a slice of E_SystemLogging_SYSLOG_FACILITY samples.
type CollectionE_SystemLogging_SYSLOG_FACILITY struct {
	W    *E_SystemLogging_SYSLOG_FACILITYWatcher
	Data []*QualifiedE_SystemLogging_SYSLOG_FACILITY
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SystemLogging_SYSLOG_FACILITY) Await(t testing.TB) []*QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SystemLogging_SYSLOG_FACILITYWatcher observes a stream of E_SystemLogging_SYSLOG_FACILITY samples.
type E_SystemLogging_SYSLOG_FACILITYWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SystemLogging_SYSLOG_FACILITY
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SystemLogging_SYSLOG_FACILITYWatcher) Await(t testing.TB) (*QualifiedE_SystemLogging_SYSLOG_FACILITY, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SystemLogging_SyslogSeverity is a E_SystemLogging_SyslogSeverity with a corresponding timestamp.
type QualifiedE_SystemLogging_SyslogSeverity struct {
	*genutil.Metadata
	val     E_SystemLogging_SyslogSeverity // val is the sample value.
	present bool
}

func (q *QualifiedE_SystemLogging_SyslogSeverity) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_SystemLogging_SyslogSeverity sample, erroring out if not present.
func (q *QualifiedE_SystemLogging_SyslogSeverity) Val(t testing.TB) E_SystemLogging_SyslogSeverity {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_SystemLogging_SyslogSeverity sample.
func (q *QualifiedE_SystemLogging_SyslogSeverity) SetVal(v E_SystemLogging_SyslogSeverity) *QualifiedE_SystemLogging_SyslogSeverity {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SystemLogging_SyslogSeverity) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SystemLogging_SyslogSeverity is a telemetry Collection whose Await method returns a slice of E_SystemLogging_SyslogSeverity samples.
type CollectionE_SystemLogging_SyslogSeverity struct {
	W    *E_SystemLogging_SyslogSeverityWatcher
	Data []*QualifiedE_SystemLogging_SyslogSeverity
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SystemLogging_SyslogSeverity) Await(t testing.TB) []*QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SystemLogging_SyslogSeverityWatcher observes a stream of E_SystemLogging_SyslogSeverity samples.
type E_SystemLogging_SyslogSeverityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SystemLogging_SyslogSeverity
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SystemLogging_SyslogSeverityWatcher) Await(t testing.TB) (*QualifiedE_SystemLogging_SyslogSeverity, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_System_NTP_AUTH_TYPE is a E_System_NTP_AUTH_TYPE with a corresponding timestamp.
type QualifiedE_System_NTP_AUTH_TYPE struct {
	*genutil.Metadata
	val     E_System_NTP_AUTH_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_System_NTP_AUTH_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_System_NTP_AUTH_TYPE sample, erroring out if not present.
func (q *QualifiedE_System_NTP_AUTH_TYPE) Val(t testing.TB) E_System_NTP_AUTH_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_System_NTP_AUTH_TYPE sample.
func (q *QualifiedE_System_NTP_AUTH_TYPE) SetVal(v E_System_NTP_AUTH_TYPE) *QualifiedE_System_NTP_AUTH_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_System_NTP_AUTH_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_System_NTP_AUTH_TYPE is a telemetry Collection whose Await method returns a slice of E_System_NTP_AUTH_TYPE samples.
type CollectionE_System_NTP_AUTH_TYPE struct {
	W    *E_System_NTP_AUTH_TYPEWatcher
	Data []*QualifiedE_System_NTP_AUTH_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_System_NTP_AUTH_TYPE) Await(t testing.TB) []*QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_System_NTP_AUTH_TYPEWatcher observes a stream of E_System_NTP_AUTH_TYPE samples.
type E_System_NTP_AUTH_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_System_NTP_AUTH_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_System_NTP_AUTH_TYPEWatcher) Await(t testing.TB) (*QualifiedE_System_NTP_AUTH_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Tlv_Reason is a E_Tlv_Reason with a corresponding timestamp.
type QualifiedE_Tlv_Reason struct {
	*genutil.Metadata
	val     E_Tlv_Reason // val is the sample value.
	present bool
}

func (q *QualifiedE_Tlv_Reason) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Tlv_Reason sample, erroring out if not present.
func (q *QualifiedE_Tlv_Reason) Val(t testing.TB) E_Tlv_Reason {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Tlv_Reason sample.
func (q *QualifiedE_Tlv_Reason) SetVal(v E_Tlv_Reason) *QualifiedE_Tlv_Reason {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Tlv_Reason) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Tlv_Reason is a telemetry Collection whose Await method returns a slice of E_Tlv_Reason samples.
type CollectionE_Tlv_Reason struct {
	W    *E_Tlv_ReasonWatcher
	Data []*QualifiedE_Tlv_Reason
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Tlv_Reason) Await(t testing.TB) []*QualifiedE_Tlv_Reason {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Tlv_ReasonWatcher observes a stream of E_Tlv_Reason samples.
type E_Tlv_ReasonWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Tlv_Reason
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Tlv_ReasonWatcher) Await(t testing.TB) (*QualifiedE_Tlv_Reason, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Topology_Attributes is a E_Topology_Attributes with a corresponding timestamp.
type QualifiedE_Topology_Attributes struct {
	*genutil.Metadata
	val     E_Topology_Attributes // val is the sample value.
	present bool
}

func (q *QualifiedE_Topology_Attributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Topology_Attributes sample, erroring out if not present.
func (q *QualifiedE_Topology_Attributes) Val(t testing.TB) E_Topology_Attributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Topology_Attributes sample.
func (q *QualifiedE_Topology_Attributes) SetVal(v E_Topology_Attributes) *QualifiedE_Topology_Attributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Topology_Attributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Topology_Attributes is a telemetry Collection whose Await method returns a slice of E_Topology_Attributes samples.
type CollectionE_Topology_Attributes struct {
	W    *E_Topology_AttributesWatcher
	Data []*QualifiedE_Topology_Attributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Topology_Attributes) Await(t testing.TB) []*QualifiedE_Topology_Attributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Topology_AttributesWatcher observes a stream of E_Topology_Attributes samples.
type E_Topology_AttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Topology_Attributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Topology_AttributesWatcher) Await(t testing.TB) (*QualifiedE_Topology_Attributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
