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

// QualifiedE_TransportTypes_LoopbackModeType is a E_TransportTypes_LoopbackModeType with a corresponding timestamp.
type QualifiedE_TransportTypes_LoopbackModeType struct {
	*genutil.Metadata
	val     E_TransportTypes_LoopbackModeType // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_LoopbackModeType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_LoopbackModeType sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_LoopbackModeType) Val(t testing.TB) E_TransportTypes_LoopbackModeType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_LoopbackModeType sample.
func (q *QualifiedE_TransportTypes_LoopbackModeType) SetVal(v E_TransportTypes_LoopbackModeType) *QualifiedE_TransportTypes_LoopbackModeType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_LoopbackModeType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_LoopbackModeType is a telemetry Collection whose Await method returns a slice of E_TransportTypes_LoopbackModeType samples.
type CollectionE_TransportTypes_LoopbackModeType struct {
	W    *E_TransportTypes_LoopbackModeTypeWatcher
	Data []*QualifiedE_TransportTypes_LoopbackModeType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_LoopbackModeType) Await(t testing.TB) []*QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_LoopbackModeTypeWatcher observes a stream of E_TransportTypes_LoopbackModeType samples.
type E_TransportTypes_LoopbackModeTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_LoopbackModeType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_LoopbackModeTypeWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_LoopbackModeType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_TransportTypes_OTN_APPLICATION_CODE is a E_TransportTypes_OTN_APPLICATION_CODE with a corresponding timestamp.
type QualifiedE_TransportTypes_OTN_APPLICATION_CODE struct {
	*genutil.Metadata
	val     E_TransportTypes_OTN_APPLICATION_CODE // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_OTN_APPLICATION_CODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_OTN_APPLICATION_CODE sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_OTN_APPLICATION_CODE) Val(t testing.TB) E_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_OTN_APPLICATION_CODE sample.
func (q *QualifiedE_TransportTypes_OTN_APPLICATION_CODE) SetVal(v E_TransportTypes_OTN_APPLICATION_CODE) *QualifiedE_TransportTypes_OTN_APPLICATION_CODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_OTN_APPLICATION_CODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_OTN_APPLICATION_CODE is a telemetry Collection whose Await method returns a slice of E_TransportTypes_OTN_APPLICATION_CODE samples.
type CollectionE_TransportTypes_OTN_APPLICATION_CODE struct {
	W    *E_TransportTypes_OTN_APPLICATION_CODEWatcher
	Data []*QualifiedE_TransportTypes_OTN_APPLICATION_CODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_OTN_APPLICATION_CODE) Await(t testing.TB) []*QualifiedE_TransportTypes_OTN_APPLICATION_CODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_OTN_APPLICATION_CODEWatcher observes a stream of E_TransportTypes_OTN_APPLICATION_CODE samples.
type E_TransportTypes_OTN_APPLICATION_CODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_OTN_APPLICATION_CODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_OTN_APPLICATION_CODEWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_OTN_APPLICATION_CODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_TransportTypes_SONET_APPLICATION_CODE is a E_TransportTypes_SONET_APPLICATION_CODE with a corresponding timestamp.
type QualifiedE_TransportTypes_SONET_APPLICATION_CODE struct {
	*genutil.Metadata
	val     E_TransportTypes_SONET_APPLICATION_CODE // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_SONET_APPLICATION_CODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_SONET_APPLICATION_CODE sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_SONET_APPLICATION_CODE) Val(t testing.TB) E_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_SONET_APPLICATION_CODE sample.
func (q *QualifiedE_TransportTypes_SONET_APPLICATION_CODE) SetVal(v E_TransportTypes_SONET_APPLICATION_CODE) *QualifiedE_TransportTypes_SONET_APPLICATION_CODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_SONET_APPLICATION_CODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_SONET_APPLICATION_CODE is a telemetry Collection whose Await method returns a slice of E_TransportTypes_SONET_APPLICATION_CODE samples.
type CollectionE_TransportTypes_SONET_APPLICATION_CODE struct {
	W    *E_TransportTypes_SONET_APPLICATION_CODEWatcher
	Data []*QualifiedE_TransportTypes_SONET_APPLICATION_CODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_SONET_APPLICATION_CODE) Await(t testing.TB) []*QualifiedE_TransportTypes_SONET_APPLICATION_CODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_SONET_APPLICATION_CODEWatcher observes a stream of E_TransportTypes_SONET_APPLICATION_CODE samples.
type E_TransportTypes_SONET_APPLICATION_CODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_SONET_APPLICATION_CODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_SONET_APPLICATION_CODEWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_SONET_APPLICATION_CODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE is a E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE with a corresponding timestamp.
type QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE struct {
	*genutil.Metadata
	val     E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) Val(t testing.TB) E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE sample.
func (q *QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) SetVal(v E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) *QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE is a telemetry Collection whose Await method returns a slice of E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE samples.
type CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE struct {
	W    *E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher
	Data []*QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) Await(t testing.TB) []*QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher observes a stream of E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE samples.
type E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPEWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE is a E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE with a corresponding timestamp.
type QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE struct {
	*genutil.Metadata
	val     E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) Val(t testing.TB) E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE sample.
func (q *QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) SetVal(v E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) *QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE is a telemetry Collection whose Await method returns a slice of E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE samples.
type CollectionE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE struct {
	W    *E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher
	Data []*QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) Await(t testing.TB) []*QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher observes a stream of E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE samples.
type E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPEWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE is a E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE with a corresponding timestamp.
type QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE struct {
	*genutil.Metadata
	val     E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) Val(t testing.TB) E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE sample.
func (q *QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) SetVal(v E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) *QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE is a telemetry Collection whose Await method returns a slice of E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE samples.
type CollectionE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE struct {
	W    *E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher
	Data []*QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) Await(t testing.TB) []*QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher observes a stream of E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE samples.
type E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_TRIBUTARY_PROTOCOL_TYPEWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE is a E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE with a corresponding timestamp.
type QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE struct {
	*genutil.Metadata
	val     E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) Val(t testing.TB) E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE sample.
func (q *QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) SetVal(v E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) *QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE is a telemetry Collection whose Await method returns a slice of E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE samples.
type CollectionE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE struct {
	W    *E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher
	Data []*QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) Await(t testing.TB) []*QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher observes a stream of E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE samples.
type E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPEWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY is a E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY with a corresponding timestamp.
type QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY struct {
	*genutil.Metadata
	val     E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY // val is the sample value.
	present bool
}

func (q *QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY sample, erroring out if not present.
func (q *QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) Val(t testing.TB) E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY sample.
func (q *QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) SetVal(v E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) *QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY is a telemetry Collection whose Await method returns a slice of E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY samples.
type CollectionE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY struct {
	W    *E_TransportTypes_TRIBUTARY_SLOT_GRANULARITYWatcher
	Data []*QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) Await(t testing.TB) []*QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_TransportTypes_TRIBUTARY_SLOT_GRANULARITYWatcher observes a stream of E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY samples.
type E_TransportTypes_TRIBUTARY_SLOT_GRANULARITYWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_TransportTypes_TRIBUTARY_SLOT_GRANULARITYWatcher) Await(t testing.TB) (*QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Types_ADDRESS_FAMILY is a E_Types_ADDRESS_FAMILY with a corresponding timestamp.
type QualifiedE_Types_ADDRESS_FAMILY struct {
	*genutil.Metadata
	val     E_Types_ADDRESS_FAMILY // val is the sample value.
	present bool
}

func (q *QualifiedE_Types_ADDRESS_FAMILY) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Types_ADDRESS_FAMILY sample, erroring out if not present.
func (q *QualifiedE_Types_ADDRESS_FAMILY) Val(t testing.TB) E_Types_ADDRESS_FAMILY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Types_ADDRESS_FAMILY sample.
func (q *QualifiedE_Types_ADDRESS_FAMILY) SetVal(v E_Types_ADDRESS_FAMILY) *QualifiedE_Types_ADDRESS_FAMILY {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Types_ADDRESS_FAMILY) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Types_ADDRESS_FAMILY is a telemetry Collection whose Await method returns a slice of E_Types_ADDRESS_FAMILY samples.
type CollectionE_Types_ADDRESS_FAMILY struct {
	W    *E_Types_ADDRESS_FAMILYWatcher
	Data []*QualifiedE_Types_ADDRESS_FAMILY
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Types_ADDRESS_FAMILY) Await(t testing.TB) []*QualifiedE_Types_ADDRESS_FAMILY {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Types_ADDRESS_FAMILYWatcher observes a stream of E_Types_ADDRESS_FAMILY samples.
type E_Types_ADDRESS_FAMILYWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Types_ADDRESS_FAMILY
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Types_ADDRESS_FAMILYWatcher) Await(t testing.TB) (*QualifiedE_Types_ADDRESS_FAMILY, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_VlanTypes_TPID_TYPES is a E_VlanTypes_TPID_TYPES with a corresponding timestamp.
type QualifiedE_VlanTypes_TPID_TYPES struct {
	*genutil.Metadata
	val     E_VlanTypes_TPID_TYPES // val is the sample value.
	present bool
}

func (q *QualifiedE_VlanTypes_TPID_TYPES) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_VlanTypes_TPID_TYPES sample, erroring out if not present.
func (q *QualifiedE_VlanTypes_TPID_TYPES) Val(t testing.TB) E_VlanTypes_TPID_TYPES {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_VlanTypes_TPID_TYPES sample.
func (q *QualifiedE_VlanTypes_TPID_TYPES) SetVal(v E_VlanTypes_TPID_TYPES) *QualifiedE_VlanTypes_TPID_TYPES {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_VlanTypes_TPID_TYPES) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_VlanTypes_TPID_TYPES is a telemetry Collection whose Await method returns a slice of E_VlanTypes_TPID_TYPES samples.
type CollectionE_VlanTypes_TPID_TYPES struct {
	W    *E_VlanTypes_TPID_TYPESWatcher
	Data []*QualifiedE_VlanTypes_TPID_TYPES
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_VlanTypes_TPID_TYPES) Await(t testing.TB) []*QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_VlanTypes_TPID_TYPESWatcher observes a stream of E_VlanTypes_TPID_TYPES samples.
type E_VlanTypes_TPID_TYPESWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_VlanTypes_TPID_TYPES
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_VlanTypes_TPID_TYPESWatcher) Await(t testing.TB) (*QualifiedE_VlanTypes_TPID_TYPES, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_VlanTypes_VlanModeType is a E_VlanTypes_VlanModeType with a corresponding timestamp.
type QualifiedE_VlanTypes_VlanModeType struct {
	*genutil.Metadata
	val     E_VlanTypes_VlanModeType // val is the sample value.
	present bool
}

func (q *QualifiedE_VlanTypes_VlanModeType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_VlanTypes_VlanModeType sample, erroring out if not present.
func (q *QualifiedE_VlanTypes_VlanModeType) Val(t testing.TB) E_VlanTypes_VlanModeType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_VlanTypes_VlanModeType sample.
func (q *QualifiedE_VlanTypes_VlanModeType) SetVal(v E_VlanTypes_VlanModeType) *QualifiedE_VlanTypes_VlanModeType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_VlanTypes_VlanModeType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_VlanTypes_VlanModeType is a telemetry Collection whose Await method returns a slice of E_VlanTypes_VlanModeType samples.
type CollectionE_VlanTypes_VlanModeType struct {
	W    *E_VlanTypes_VlanModeTypeWatcher
	Data []*QualifiedE_VlanTypes_VlanModeType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_VlanTypes_VlanModeType) Await(t testing.TB) []*QualifiedE_VlanTypes_VlanModeType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_VlanTypes_VlanModeTypeWatcher observes a stream of E_VlanTypes_VlanModeType samples.
type E_VlanTypes_VlanModeTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_VlanTypes_VlanModeType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_VlanTypes_VlanModeTypeWatcher) Await(t testing.TB) (*QualifiedE_VlanTypes_VlanModeType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_VlanTypes_VlanStackAction is a E_VlanTypes_VlanStackAction with a corresponding timestamp.
type QualifiedE_VlanTypes_VlanStackAction struct {
	*genutil.Metadata
	val     E_VlanTypes_VlanStackAction // val is the sample value.
	present bool
}

func (q *QualifiedE_VlanTypes_VlanStackAction) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_VlanTypes_VlanStackAction sample, erroring out if not present.
func (q *QualifiedE_VlanTypes_VlanStackAction) Val(t testing.TB) E_VlanTypes_VlanStackAction {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_VlanTypes_VlanStackAction sample.
func (q *QualifiedE_VlanTypes_VlanStackAction) SetVal(v E_VlanTypes_VlanStackAction) *QualifiedE_VlanTypes_VlanStackAction {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_VlanTypes_VlanStackAction) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_VlanTypes_VlanStackAction is a telemetry Collection whose Await method returns a slice of E_VlanTypes_VlanStackAction samples.
type CollectionE_VlanTypes_VlanStackAction struct {
	W    *E_VlanTypes_VlanStackActionWatcher
	Data []*QualifiedE_VlanTypes_VlanStackAction
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_VlanTypes_VlanStackAction) Await(t testing.TB) []*QualifiedE_VlanTypes_VlanStackAction {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_VlanTypes_VlanStackActionWatcher observes a stream of E_VlanTypes_VlanStackAction samples.
type E_VlanTypes_VlanStackActionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_VlanTypes_VlanStackAction
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_VlanTypes_VlanStackActionWatcher) Await(t testing.TB) (*QualifiedE_VlanTypes_VlanStackAction, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Vlan_Status is a E_Vlan_Status with a corresponding timestamp.
type QualifiedE_Vlan_Status struct {
	*genutil.Metadata
	val     E_Vlan_Status // val is the sample value.
	present bool
}

func (q *QualifiedE_Vlan_Status) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Vlan_Status sample, erroring out if not present.
func (q *QualifiedE_Vlan_Status) Val(t testing.TB) E_Vlan_Status {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Vlan_Status sample.
func (q *QualifiedE_Vlan_Status) SetVal(v E_Vlan_Status) *QualifiedE_Vlan_Status {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Vlan_Status) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Vlan_Status is a telemetry Collection whose Await method returns a slice of E_Vlan_Status samples.
type CollectionE_Vlan_Status struct {
	W    *E_Vlan_StatusWatcher
	Data []*QualifiedE_Vlan_Status
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Vlan_Status) Await(t testing.TB) []*QualifiedE_Vlan_Status {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Vlan_StatusWatcher observes a stream of E_Vlan_Status samples.
type E_Vlan_StatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Vlan_Status
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Vlan_StatusWatcher) Await(t testing.TB) (*QualifiedE_Vlan_Status, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_IngressTracking_MplsLabel_Union is a Flow_IngressTracking_MplsLabel_Union with a corresponding timestamp.
type QualifiedFlow_IngressTracking_MplsLabel_Union struct {
	*genutil.Metadata
	val     Flow_IngressTracking_MplsLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedFlow_IngressTracking_MplsLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Flow_IngressTracking_MplsLabel_Union sample, erroring out if not present.
func (q *QualifiedFlow_IngressTracking_MplsLabel_Union) Val(t testing.TB) Flow_IngressTracking_MplsLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Flow_IngressTracking_MplsLabel_Union sample.
func (q *QualifiedFlow_IngressTracking_MplsLabel_Union) SetVal(v Flow_IngressTracking_MplsLabel_Union) *QualifiedFlow_IngressTracking_MplsLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_IngressTracking_MplsLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_IngressTracking_MplsLabel_Union is a telemetry Collection whose Await method returns a slice of Flow_IngressTracking_MplsLabel_Union samples.
type CollectionFlow_IngressTracking_MplsLabel_Union struct {
	W    *Flow_IngressTracking_MplsLabel_UnionWatcher
	Data []*QualifiedFlow_IngressTracking_MplsLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_IngressTracking_MplsLabel_Union) Await(t testing.TB) []*QualifiedFlow_IngressTracking_MplsLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_IngressTracking_MplsLabel_UnionWatcher observes a stream of Flow_IngressTracking_MplsLabel_Union samples.
type Flow_IngressTracking_MplsLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_IngressTracking_MplsLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_IngressTracking_MplsLabel_UnionWatcher) Await(t testing.TB) (*QualifiedFlow_IngressTracking_MplsLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_MplsLabel_Union is a Flow_MplsLabel_Union with a corresponding timestamp.
type QualifiedFlow_MplsLabel_Union struct {
	*genutil.Metadata
	val     Flow_MplsLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedFlow_MplsLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Flow_MplsLabel_Union sample, erroring out if not present.
func (q *QualifiedFlow_MplsLabel_Union) Val(t testing.TB) Flow_MplsLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Flow_MplsLabel_Union sample.
func (q *QualifiedFlow_MplsLabel_Union) SetVal(v Flow_MplsLabel_Union) *QualifiedFlow_MplsLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_MplsLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_MplsLabel_Union is a telemetry Collection whose Await method returns a slice of Flow_MplsLabel_Union samples.
type CollectionFlow_MplsLabel_Union struct {
	W    *Flow_MplsLabel_UnionWatcher
	Data []*QualifiedFlow_MplsLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_MplsLabel_Union) Await(t testing.TB) []*QualifiedFlow_MplsLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_MplsLabel_UnionWatcher observes a stream of Flow_MplsLabel_Union samples.
type Flow_MplsLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_MplsLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_MplsLabel_UnionWatcher) Await(t testing.TB) (*QualifiedFlow_MplsLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_RoutedVlan_Vlan_Union is a Interface_RoutedVlan_Vlan_Union with a corresponding timestamp.
type QualifiedInterface_RoutedVlan_Vlan_Union struct {
	*genutil.Metadata
	val     Interface_RoutedVlan_Vlan_Union // val is the sample value.
	present bool
}

func (q *QualifiedInterface_RoutedVlan_Vlan_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Interface_RoutedVlan_Vlan_Union sample, erroring out if not present.
func (q *QualifiedInterface_RoutedVlan_Vlan_Union) Val(t testing.TB) Interface_RoutedVlan_Vlan_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Interface_RoutedVlan_Vlan_Union sample.
func (q *QualifiedInterface_RoutedVlan_Vlan_Union) SetVal(v Interface_RoutedVlan_Vlan_Union) *QualifiedInterface_RoutedVlan_Vlan_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_RoutedVlan_Vlan_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_RoutedVlan_Vlan_Union is a telemetry Collection whose Await method returns a slice of Interface_RoutedVlan_Vlan_Union samples.
type CollectionInterface_RoutedVlan_Vlan_Union struct {
	W    *Interface_RoutedVlan_Vlan_UnionWatcher
	Data []*QualifiedInterface_RoutedVlan_Vlan_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_RoutedVlan_Vlan_Union) Await(t testing.TB) []*QualifiedInterface_RoutedVlan_Vlan_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_RoutedVlan_Vlan_UnionWatcher observes a stream of Interface_RoutedVlan_Vlan_Union samples.
type Interface_RoutedVlan_Vlan_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_RoutedVlan_Vlan_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_RoutedVlan_Vlan_UnionWatcher) Await(t testing.TB) (*QualifiedInterface_RoutedVlan_Vlan_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Subinterface_Vlan_VlanId_Union is a Interface_Subinterface_Vlan_VlanId_Union with a corresponding timestamp.
type QualifiedInterface_Subinterface_Vlan_VlanId_Union struct {
	*genutil.Metadata
	val     Interface_Subinterface_Vlan_VlanId_Union // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Subinterface_Vlan_VlanId_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Interface_Subinterface_Vlan_VlanId_Union sample, erroring out if not present.
func (q *QualifiedInterface_Subinterface_Vlan_VlanId_Union) Val(t testing.TB) Interface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Interface_Subinterface_Vlan_VlanId_Union sample.
func (q *QualifiedInterface_Subinterface_Vlan_VlanId_Union) SetVal(v Interface_Subinterface_Vlan_VlanId_Union) *QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Subinterface_Vlan_VlanId_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Subinterface_Vlan_VlanId_Union is a telemetry Collection whose Await method returns a slice of Interface_Subinterface_Vlan_VlanId_Union samples.
type CollectionInterface_Subinterface_Vlan_VlanId_Union struct {
	W    *Interface_Subinterface_Vlan_VlanId_UnionWatcher
	Data []*QualifiedInterface_Subinterface_Vlan_VlanId_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Subinterface_Vlan_VlanId_Union) Await(t testing.TB) []*QualifiedInterface_Subinterface_Vlan_VlanId_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Subinterface_Vlan_VlanId_UnionWatcher observes a stream of Interface_Subinterface_Vlan_VlanId_Union samples.
type Interface_Subinterface_Vlan_VlanId_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Subinterface_Vlan_VlanId_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Subinterface_Vlan_VlanId_UnionWatcher) Await(t testing.TB) (*QualifiedInterface_Subinterface_Vlan_VlanId_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedKeychain_Tolerance_Union is a Keychain_Tolerance_Union with a corresponding timestamp.
type QualifiedKeychain_Tolerance_Union struct {
	*genutil.Metadata
	val     Keychain_Tolerance_Union // val is the sample value.
	present bool
}

func (q *QualifiedKeychain_Tolerance_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Keychain_Tolerance_Union sample, erroring out if not present.
func (q *QualifiedKeychain_Tolerance_Union) Val(t testing.TB) Keychain_Tolerance_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Keychain_Tolerance_Union sample.
func (q *QualifiedKeychain_Tolerance_Union) SetVal(v Keychain_Tolerance_Union) *QualifiedKeychain_Tolerance_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedKeychain_Tolerance_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionKeychain_Tolerance_Union is a telemetry Collection whose Await method returns a slice of Keychain_Tolerance_Union samples.
type CollectionKeychain_Tolerance_Union struct {
	W    *Keychain_Tolerance_UnionWatcher
	Data []*QualifiedKeychain_Tolerance_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionKeychain_Tolerance_Union) Await(t testing.TB) []*QualifiedKeychain_Tolerance_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Keychain_Tolerance_UnionWatcher observes a stream of Keychain_Tolerance_Union samples.
type Keychain_Tolerance_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedKeychain_Tolerance_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Keychain_Tolerance_UnionWatcher) Await(t testing.TB) (*QualifiedKeychain_Tolerance_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_LabelEntry_Label_Union is a NetworkInstance_Afts_LabelEntry_Label_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_LabelEntry_Label_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Afts_LabelEntry_Label_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_LabelEntry_Label_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Afts_LabelEntry_Label_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_Label_Union) Val(t testing.TB) NetworkInstance_Afts_LabelEntry_Label_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Afts_LabelEntry_Label_Union sample.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_Label_Union) SetVal(v NetworkInstance_Afts_LabelEntry_Label_Union) *QualifiedNetworkInstance_Afts_LabelEntry_Label_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_Label_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_LabelEntry_Label_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Afts_LabelEntry_Label_Union samples.
type CollectionNetworkInstance_Afts_LabelEntry_Label_Union struct {
	W    *NetworkInstance_Afts_LabelEntry_Label_UnionWatcher
	Data []*QualifiedNetworkInstance_Afts_LabelEntry_Label_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_LabelEntry_Label_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_LabelEntry_Label_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_LabelEntry_Label_UnionWatcher observes a stream of NetworkInstance_Afts_LabelEntry_Label_Union samples.
type NetworkInstance_Afts_LabelEntry_Label_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_LabelEntry_Label_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_LabelEntry_Label_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_LabelEntry_Label_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union is a NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union) Val(t testing.TB) NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union sample.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union) SetVal(v NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union) *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union samples.
type CollectionNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union struct {
	W    *NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_UnionWatcher
	Data []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_UnionWatcher observes a stream of NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union samples.
type NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_IpProtocol_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union is a NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union) Val(t testing.TB) NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union sample.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union) SetVal(v NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union) *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union samples.
type CollectionNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union struct {
	W    *NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_UnionWatcher observes a stream of NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union samples.
type NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_PolicyForwardingEntry_MplsLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union is a NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union with a corresponding timestamp.
type QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union struct {
	*genutil.Metadata
	val     NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union) Val(t testing.TB) NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union sample.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union) SetVal(v NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union) *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union samples.
type CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union struct {
	W    *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_UnionWatcher
	Data []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union) Await(t testing.TB) []*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_UnionWatcher observes a stream of NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union samples.
type NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_ConnectionPoint_Endpoint_Vxlan_EndpointVni_MultidestinationTraffic_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union is a NetworkInstance_Evpn_EthernetSegment_Esi_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Evpn_EthernetSegment_Esi_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Evpn_EthernetSegment_Esi_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union) Val(t testing.TB) NetworkInstance_Evpn_EthernetSegment_Esi_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Evpn_EthernetSegment_Esi_Union sample.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union) SetVal(v NetworkInstance_Evpn_EthernetSegment_Esi_Union) *QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EthernetSegment_Esi_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Evpn_EthernetSegment_Esi_Union samples.
type CollectionNetworkInstance_Evpn_EthernetSegment_Esi_Union struct {
	W    *NetworkInstance_Evpn_EthernetSegment_Esi_UnionWatcher
	Data []*QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EthernetSegment_Esi_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EthernetSegment_Esi_UnionWatcher observes a stream of NetworkInstance_Evpn_EthernetSegment_Esi_Union samples.
type NetworkInstance_Evpn_EthernetSegment_Esi_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EthernetSegment_Esi_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EthernetSegment_Esi_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union is a NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union) Val(t testing.TB) NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union) SetVal(v NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union) *QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union struct {
	W    *NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_UnionWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_UnionWatcher observes a stream of NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union samples.
type NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_RouteDistinguisher_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union is a NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union) Val(t testing.TB) NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union) SetVal(v NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union) *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union samples.
type CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union struct {
	W    *NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_UnionWatcher observes a stream of NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union samples.
type NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_LowerBound_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union is a NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union) Val(t testing.TB) NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union) SetVal(v NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union) *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union samples.
type CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union struct {
	W    *NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_UnionWatcher observes a stream of NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union samples.
type NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Global_ReservedLabelBlock_UpperBound_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union is a NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union) SetVal(v NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_IncomingLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union is a NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union) SetVal(v NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Egress_PushLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union is a NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union) SetVal(v NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_IncomingLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union is a NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union) SetVal(v NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Ingress_PushLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union is a NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union) SetVal(v NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_IncomingLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union is a NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union) SetVal(v NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union) *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union samples.
type CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union struct {
	W    *NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union samples.
type NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_Lsps_StaticLsp_Transit_PushLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union is a NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union) Val(t testing.TB) NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union) SetVal(v NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_UnionWatcher observes a stream of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Interface_BandwidthReservation_Priority_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union is a NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union) Val(t testing.TB) NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union) SetVal(v NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_UnionWatcher observes a stream of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_ExplicitRouteObject_Label_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union is a NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union) Val(t testing.TB) NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union) SetVal(v NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_UnionWatcher observes a stream of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelIn_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union is a NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union) Val(t testing.TB) NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union) SetVal(v NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_UnionWatcher observes a stream of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_LabelOut_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union is a NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union) SetVal(v NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union) *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union samples.
type NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session_RecordRouteObject_ReportedLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union is a NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union) SetVal(v NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounter_MplsLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union is a NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union) Val(t testing.TB) NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union sample.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union) SetVal(v NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union) *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union samples.
type CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union struct {
	W    *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_UnionWatcher
	Data []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_UnionWatcher observes a stream of NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union samples.
type NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Mpls_SignalingProtocols_SegmentRouting_Interface_SidCounter_MplsLabel_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union is a NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union struct {
	*genutil.Metadata
	val     NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union) Val(t testing.TB) NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union) SetVal(v NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_UnionWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_UnionWatcher observes a stream of NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv4_Protocol_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union is a NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union struct {
	*genutil.Metadata
	val     NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union) Val(t testing.TB) NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union) SetVal(v NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_UnionWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_UnionWatcher observes a stream of NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Ipv6_Protocol_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union is a NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union struct {
	*genutil.Metadata
	val     NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union) Val(t testing.TB) NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union) SetVal(v NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_UnionWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_UnionWatcher observes a stream of NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_L2_Ethertype_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union is a NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union struct {
	*genutil.Metadata
	val     NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union) Val(t testing.TB) NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union) SetVal(v NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_UnionWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_UnionWatcher observes a stream of NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_DestinationPort_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union is a NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union with a corresponding timestamp.
type QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union struct {
	*genutil.Metadata
	val     NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union) Val(t testing.TB) NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union sample.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union) SetVal(v NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union) *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union samples.
type CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union struct {
	W    *NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_UnionWatcher
	Data []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union) Await(t testing.TB) []*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_UnionWatcher observes a stream of NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union samples.
type NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_PolicyForwarding_Policy_Rule_Transport_SourcePort_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union is a NetworkInstance_Protocol_Aggregate_SetTag_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Aggregate_SetTag_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Aggregate_SetTag_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union) Val(t testing.TB) NetworkInstance_Protocol_Aggregate_SetTag_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Aggregate_SetTag_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union) SetVal(v NetworkInstance_Protocol_Aggregate_SetTag_Union) *QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Aggregate_SetTag_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Aggregate_SetTag_Union samples.
type CollectionNetworkInstance_Protocol_Aggregate_SetTag_Union struct {
	W    *NetworkInstance_Protocol_Aggregate_SetTag_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Aggregate_SetTag_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Aggregate_SetTag_UnionWatcher observes a stream of NetworkInstance_Protocol_Aggregate_SetTag_Union samples.
type NetworkInstance_Protocol_Aggregate_SetTag_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Aggregate_SetTag_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Aggregate_SetTag_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union is a NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union) Val(t testing.TB) NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union) SetVal(v NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union) *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union struct {
	W    *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_UnionWatcher observes a stream of NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union samples.
type NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Neighbor_RouteReflector_RouteReflectorClusterId_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union is a NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union) Val(t testing.TB) NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union) SetVal(v NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_UnionWatcher observes a stream of NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector_RouteReflectorClusterId_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union is a NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union) Val(t testing.TB) NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union) SetVal(v NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_UnionWatcher observes a stream of NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_Origin_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union is a NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union) Val(t testing.TB) NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union) SetVal(v NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_UnionWatcher observes a stream of NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_Origin_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union is a NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union) Val(t testing.TB) NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union) SetVal(v NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_UnionWatcher observes a stream of NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_BindingSid_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union is a NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union) Val(t testing.TB) NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union) SetVal(v NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_UnionWatcher observes a stream of NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment_Sid_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union is a NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union) Val(t testing.TB) NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union) SetVal(v NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_UnionWatcher observes a stream of NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union samples.
type NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_AllocatedDynamicLocal_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union is a NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union) Val(t testing.TB) NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union) SetVal(v NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_UnionWatcher observes a stream of NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union samples.
type NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid_SidId_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union is a NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union) Val(t testing.TB) NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union) SetVal(v NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_UnionWatcher observes a stream of NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union samples.
type NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid_SidId_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union is a NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union) Val(t testing.TB) NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union) SetVal(v NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_UnionWatcher observes a stream of NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor_Label_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union is a NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Identifier_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Identifier_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Identifier_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Identifier_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Identifier_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Identifier_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union is a NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_Identifier_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union is a NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_ExtendedLink_LinkData_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union is a NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_SegmentRoutingSidLabelRange_Tlv_Type_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union is a NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_RouterInformation_Tlv_Type_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union is a NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_Link_SubTlv_Type_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union is a NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_OpaqueLsa_TrafficEngineering_Tlv_NodeAttribute_SubTlv_Type_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union is a NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union) Val(t testing.TB) NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union) SetVal(v NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union) *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union samples.
type CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union struct {
	W    *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_UnionWatcher observes a stream of NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union samples.
type NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Ospfv2_Area_Lsdb_LsaType_Lsa_RouterLsa_LinkData_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union is a NetworkInstance_Protocol_Static_NextHop_NextHop_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Static_NextHop_NextHop_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Static_NextHop_NextHop_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union) Val(t testing.TB) NetworkInstance_Protocol_Static_NextHop_NextHop_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Static_NextHop_NextHop_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union) SetVal(v NetworkInstance_Protocol_Static_NextHop_NextHop_Union) *QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Static_NextHop_NextHop_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Static_NextHop_NextHop_Union samples.
type CollectionNetworkInstance_Protocol_Static_NextHop_NextHop_Union struct {
	W    *NetworkInstance_Protocol_Static_NextHop_NextHop_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Static_NextHop_NextHop_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Static_NextHop_NextHop_UnionWatcher observes a stream of NetworkInstance_Protocol_Static_NextHop_NextHop_Union samples.
type NetworkInstance_Protocol_Static_NextHop_NextHop_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Static_NextHop_NextHop_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Static_NextHop_NextHop_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Static_SetTag_Union is a NetworkInstance_Protocol_Static_SetTag_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Static_SetTag_Union struct {
	*genutil.Metadata
	val     NetworkInstance_Protocol_Static_SetTag_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Static_SetTag_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_Protocol_Static_SetTag_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Static_SetTag_Union) Val(t testing.TB) NetworkInstance_Protocol_Static_SetTag_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_Protocol_Static_SetTag_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Static_SetTag_Union) SetVal(v NetworkInstance_Protocol_Static_SetTag_Union) *QualifiedNetworkInstance_Protocol_Static_SetTag_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Static_SetTag_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Static_SetTag_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_Protocol_Static_SetTag_Union samples.
type CollectionNetworkInstance_Protocol_Static_SetTag_Union struct {
	W    *NetworkInstance_Protocol_Static_SetTag_UnionWatcher
	Data []*QualifiedNetworkInstance_Protocol_Static_SetTag_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Static_SetTag_Union) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Static_SetTag_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Static_SetTag_UnionWatcher observes a stream of NetworkInstance_Protocol_Static_SetTag_Union samples.
type NetworkInstance_Protocol_Static_SetTag_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Static_SetTag_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Static_SetTag_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Static_SetTag_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union is a NetworkInstance_SegmentRouting_TePolicy_Bsid_Union with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union struct {
	*genutil.Metadata
	val     NetworkInstance_SegmentRouting_TePolicy_Bsid_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_SegmentRouting_TePolicy_Bsid_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union) Val(t testing.TB) NetworkInstance_SegmentRouting_TePolicy_Bsid_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_SegmentRouting_TePolicy_Bsid_Union sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union) SetVal(v NetworkInstance_SegmentRouting_TePolicy_Bsid_Union) *QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_Bsid_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_SegmentRouting_TePolicy_Bsid_Union samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_Bsid_Union struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_Bsid_UnionWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_Bsid_Union) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_Bsid_UnionWatcher observes a stream of NetworkInstance_SegmentRouting_TePolicy_Bsid_Union samples.
type NetworkInstance_SegmentRouting_TePolicy_Bsid_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_Bsid_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_Bsid_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union is a NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union struct {
	*genutil.Metadata
	val     NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union) Val(t testing.TB) NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union) SetVal(v NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union is a telemetry Collection whose Await method returns a slice of NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_UnionWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_UnionWatcher observes a stream of NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_UnionWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_Sid_Value_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union is a Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union struct {
	*genutil.Metadata
	val     Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union) Val(t testing.TB) Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union) SetVal(v Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union) *QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Ipv4_Protocol_Union is a telemetry Collection whose Await method returns a slice of Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union samples.
type CollectionQos_Classifier_Term_Conditions_Ipv4_Protocol_Union struct {
	W    *Qos_Classifier_Term_Conditions_Ipv4_Protocol_UnionWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Ipv4_Protocol_Union) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Ipv4_Protocol_UnionWatcher observes a stream of Qos_Classifier_Term_Conditions_Ipv4_Protocol_Union samples.
type Qos_Classifier_Term_Conditions_Ipv4_Protocol_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Ipv4_Protocol_UnionWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Ipv4_Protocol_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union is a Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union struct {
	*genutil.Metadata
	val     Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) Val(t testing.TB) Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) SetVal(v Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union) *QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Ipv6_Protocol_Union is a telemetry Collection whose Await method returns a slice of Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union samples.
type CollectionQos_Classifier_Term_Conditions_Ipv6_Protocol_Union struct {
	W    *Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Ipv6_Protocol_Union) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher observes a stream of Qos_Classifier_Term_Conditions_Ipv6_Protocol_Union samples.
type Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Ipv6_Protocol_UnionWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Ipv6_Protocol_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union is a Qos_Classifier_Term_Conditions_L2_Ethertype_Union with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union struct {
	*genutil.Metadata
	val     Qos_Classifier_Term_Conditions_L2_Ethertype_Union // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Qos_Classifier_Term_Conditions_L2_Ethertype_Union sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) Val(t testing.TB) Qos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Qos_Classifier_Term_Conditions_L2_Ethertype_Union sample.
func (q *QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) SetVal(v Qos_Classifier_Term_Conditions_L2_Ethertype_Union) *QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_L2_Ethertype_Union is a telemetry Collection whose Await method returns a slice of Qos_Classifier_Term_Conditions_L2_Ethertype_Union samples.
type CollectionQos_Classifier_Term_Conditions_L2_Ethertype_Union struct {
	W    *Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_L2_Ethertype_Union) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher observes a stream of Qos_Classifier_Term_Conditions_L2_Ethertype_Union samples.
type Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_L2_Ethertype_UnionWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union is a Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union struct {
	*genutil.Metadata
	val     Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) Val(t testing.TB) Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) SetVal(v Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) *QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union is a telemetry Collection whose Await method returns a slice of Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union samples.
type CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union struct {
	W    *Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher observes a stream of Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union samples.
type Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_UnionWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union is a Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union struct {
	*genutil.Metadata
	val     Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) Val(t testing.TB) Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) SetVal(v Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) *QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union is a telemetry Collection whose Await method returns a slice of Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union samples.
type CollectionQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union struct {
	W    *Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher observes a stream of Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union samples.
type Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_UnionWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union is a Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union struct {
	*genutil.Metadata
	val     Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) Val(t testing.TB) Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) SetVal(v Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union) *QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Transport_DestinationPort_Union is a telemetry Collection whose Await method returns a slice of Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union samples.
type CollectionQos_Classifier_Term_Conditions_Transport_DestinationPort_Union struct {
	W    *Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Transport_DestinationPort_Union) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher observes a stream of Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union samples.
type Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Transport_DestinationPort_UnionWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union is a Qos_Classifier_Term_Conditions_Transport_SourcePort_Union with a corresponding timestamp.
type QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union struct {
	*genutil.Metadata
	val     Qos_Classifier_Term_Conditions_Transport_SourcePort_Union // val is the sample value.
	present bool
}

func (q *QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Qos_Classifier_Term_Conditions_Transport_SourcePort_Union sample, erroring out if not present.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) Val(t testing.TB) Qos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Qos_Classifier_Term_Conditions_Transport_SourcePort_Union sample.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) SetVal(v Qos_Classifier_Term_Conditions_Transport_SourcePort_Union) *QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_Classifier_Term_Conditions_Transport_SourcePort_Union is a telemetry Collection whose Await method returns a slice of Qos_Classifier_Term_Conditions_Transport_SourcePort_Union samples.
type CollectionQos_Classifier_Term_Conditions_Transport_SourcePort_Union struct {
	W    *Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher
	Data []*QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_Classifier_Term_Conditions_Transport_SourcePort_Union) Await(t testing.TB) []*QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher observes a stream of Qos_Classifier_Term_Conditions_Transport_SourcePort_Union samples.
type Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_Classifier_Term_Conditions_Transport_SourcePort_UnionWatcher) Await(t testing.TB) (*QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union is a RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union struct {
	*genutil.Metadata
	val     RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) Val(t testing.TB) RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) SetVal(v RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union is a telemetry Collection whose Await method returns a slice of RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher observes a stream of RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_UnionWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union is a RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union struct {
	*genutil.Metadata
	val     RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) Val(t testing.TB) RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) SetVal(v RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union is a telemetry Collection whose Await method returns a slice of RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher observes a stream of RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_UnionWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authentication_User_Role_Union is a System_Aaa_Authentication_User_Role_Union with a corresponding timestamp.
type QualifiedSystem_Aaa_Authentication_User_Role_Union struct {
	*genutil.Metadata
	val     System_Aaa_Authentication_User_Role_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authentication_User_Role_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the System_Aaa_Authentication_User_Role_Union sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authentication_User_Role_Union) Val(t testing.TB) System_Aaa_Authentication_User_Role_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the System_Aaa_Authentication_User_Role_Union sample.
func (q *QualifiedSystem_Aaa_Authentication_User_Role_Union) SetVal(v System_Aaa_Authentication_User_Role_Union) *QualifiedSystem_Aaa_Authentication_User_Role_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authentication_User_Role_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authentication_User_Role_Union is a telemetry Collection whose Await method returns a slice of System_Aaa_Authentication_User_Role_Union samples.
type CollectionSystem_Aaa_Authentication_User_Role_Union struct {
	W    *System_Aaa_Authentication_User_Role_UnionWatcher
	Data []*QualifiedSystem_Aaa_Authentication_User_Role_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authentication_User_Role_Union) Await(t testing.TB) []*QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Authentication_User_Role_UnionWatcher observes a stream of System_Aaa_Authentication_User_Role_Union samples.
type System_Aaa_Authentication_User_Role_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authentication_User_Role_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Authentication_User_Role_UnionWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authentication_User_Role_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Alarm_TypeId_Union is a System_Alarm_TypeId_Union with a corresponding timestamp.
type QualifiedSystem_Alarm_TypeId_Union struct {
	*genutil.Metadata
	val     System_Alarm_TypeId_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Alarm_TypeId_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the System_Alarm_TypeId_Union sample, erroring out if not present.
func (q *QualifiedSystem_Alarm_TypeId_Union) Val(t testing.TB) System_Alarm_TypeId_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the System_Alarm_TypeId_Union sample.
func (q *QualifiedSystem_Alarm_TypeId_Union) SetVal(v System_Alarm_TypeId_Union) *QualifiedSystem_Alarm_TypeId_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Alarm_TypeId_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Alarm_TypeId_Union is a telemetry Collection whose Await method returns a slice of System_Alarm_TypeId_Union samples.
type CollectionSystem_Alarm_TypeId_Union struct {
	W    *System_Alarm_TypeId_UnionWatcher
	Data []*QualifiedSystem_Alarm_TypeId_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Alarm_TypeId_Union) Await(t testing.TB) []*QualifiedSystem_Alarm_TypeId_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Alarm_TypeId_UnionWatcher observes a stream of System_Alarm_TypeId_Union samples.
type System_Alarm_TypeId_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Alarm_TypeId_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Alarm_TypeId_UnionWatcher) Await(t testing.TB) (*QualifiedSystem_Alarm_TypeId_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_Index_Union is a System_Cpu_Index_Union with a corresponding timestamp.
type QualifiedSystem_Cpu_Index_Union struct {
	*genutil.Metadata
	val     System_Cpu_Index_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_Index_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the System_Cpu_Index_Union sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_Index_Union) Val(t testing.TB) System_Cpu_Index_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the System_Cpu_Index_Union sample.
func (q *QualifiedSystem_Cpu_Index_Union) SetVal(v System_Cpu_Index_Union) *QualifiedSystem_Cpu_Index_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_Index_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_Index_Union is a telemetry Collection whose Await method returns a slice of System_Cpu_Index_Union samples.
type CollectionSystem_Cpu_Index_Union struct {
	W    *System_Cpu_Index_UnionWatcher
	Data []*QualifiedSystem_Cpu_Index_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_Index_Union) Await(t testing.TB) []*QualifiedSystem_Cpu_Index_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_Index_UnionWatcher observes a stream of System_Cpu_Index_Union samples.
type System_Cpu_Index_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_Index_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_Index_UnionWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_Index_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_License_License_LicenseData_Union is a System_License_License_LicenseData_Union with a corresponding timestamp.
type QualifiedSystem_License_License_LicenseData_Union struct {
	*genutil.Metadata
	val     System_License_License_LicenseData_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_License_License_LicenseData_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the System_License_License_LicenseData_Union sample, erroring out if not present.
func (q *QualifiedSystem_License_License_LicenseData_Union) Val(t testing.TB) System_License_License_LicenseData_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the System_License_License_LicenseData_Union sample.
func (q *QualifiedSystem_License_License_LicenseData_Union) SetVal(v System_License_License_LicenseData_Union) *QualifiedSystem_License_License_LicenseData_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_License_License_LicenseData_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_License_License_LicenseData_Union is a telemetry Collection whose Await method returns a slice of System_License_License_LicenseData_Union samples.
type CollectionSystem_License_License_LicenseData_Union struct {
	W    *System_License_License_LicenseData_UnionWatcher
	Data []*QualifiedSystem_License_License_LicenseData_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_License_License_LicenseData_Union) Await(t testing.TB) []*QualifiedSystem_License_License_LicenseData_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_License_License_LicenseData_UnionWatcher observes a stream of System_License_License_LicenseData_Union samples.
type System_License_License_LicenseData_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_License_License_LicenseData_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_License_License_LicenseData_UnionWatcher) Await(t testing.TB) (*QualifiedSystem_License_License_LicenseData_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_AdjacencySid_FlagsSlice is a []E_AdjacencySid_Flags with a corresponding timestamp.
type QualifiedE_AdjacencySid_FlagsSlice struct {
	*genutil.Metadata
	val     []E_AdjacencySid_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_AdjacencySid_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_AdjacencySid_Flags sample, erroring out if not present.
func (q *QualifiedE_AdjacencySid_FlagsSlice) Val(t testing.TB) []E_AdjacencySid_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_AdjacencySid_Flags sample.
func (q *QualifiedE_AdjacencySid_FlagsSlice) SetVal(v []E_AdjacencySid_Flags) *QualifiedE_AdjacencySid_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_AdjacencySid_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_AdjacencySid_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_AdjacencySid_Flags samples.
type CollectionE_AdjacencySid_FlagsSlice struct {
	W    *E_AdjacencySid_FlagsSliceWatcher
	Data []*QualifiedE_AdjacencySid_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_AdjacencySid_FlagsSlice) Await(t testing.TB) []*QualifiedE_AdjacencySid_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_AdjacencySid_FlagsSliceWatcher observes a stream of []E_AdjacencySid_Flags samples.
type E_AdjacencySid_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_AdjacencySid_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_AdjacencySid_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_AdjacencySid_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Adjacency_NlpidSlice is a []E_Adjacency_Nlpid with a corresponding timestamp.
type QualifiedE_Adjacency_NlpidSlice struct {
	*genutil.Metadata
	val     []E_Adjacency_Nlpid // val is the sample value.
	present bool
}

func (q *QualifiedE_Adjacency_NlpidSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Adjacency_Nlpid sample, erroring out if not present.
func (q *QualifiedE_Adjacency_NlpidSlice) Val(t testing.TB) []E_Adjacency_Nlpid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Adjacency_Nlpid sample.
func (q *QualifiedE_Adjacency_NlpidSlice) SetVal(v []E_Adjacency_Nlpid) *QualifiedE_Adjacency_NlpidSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Adjacency_NlpidSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Adjacency_NlpidSlice is a telemetry Collection whose Await method returns a slice of []E_Adjacency_Nlpid samples.
type CollectionE_Adjacency_NlpidSlice struct {
	W    *E_Adjacency_NlpidSliceWatcher
	Data []*QualifiedE_Adjacency_NlpidSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Adjacency_NlpidSlice) Await(t testing.TB) []*QualifiedE_Adjacency_NlpidSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Adjacency_NlpidSliceWatcher observes a stream of []E_Adjacency_Nlpid samples.
type E_Adjacency_NlpidSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Adjacency_NlpidSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Adjacency_NlpidSliceWatcher) Await(t testing.TB) (*QualifiedE_Adjacency_NlpidSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_AFI_SAFI_TYPESlice is a []E_BgpTypes_AFI_SAFI_TYPE with a corresponding timestamp.
type QualifiedE_BgpTypes_AFI_SAFI_TYPESlice struct {
	*genutil.Metadata
	val     []E_BgpTypes_AFI_SAFI_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_BgpTypes_AFI_SAFI_TYPE sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) Val(t testing.TB) []E_BgpTypes_AFI_SAFI_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_BgpTypes_AFI_SAFI_TYPE sample.
func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) SetVal(v []E_BgpTypes_AFI_SAFI_TYPE) *QualifiedE_BgpTypes_AFI_SAFI_TYPESlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPESlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_AFI_SAFI_TYPESlice is a telemetry Collection whose Await method returns a slice of []E_BgpTypes_AFI_SAFI_TYPE samples.
type CollectionE_BgpTypes_AFI_SAFI_TYPESlice struct {
	W    *E_BgpTypes_AFI_SAFI_TYPESliceWatcher
	Data []*QualifiedE_BgpTypes_AFI_SAFI_TYPESlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_AFI_SAFI_TYPESlice) Await(t testing.TB) []*QualifiedE_BgpTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_AFI_SAFI_TYPESliceWatcher observes a stream of []E_BgpTypes_AFI_SAFI_TYPE samples.
type E_BgpTypes_AFI_SAFI_TYPESliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_AFI_SAFI_TYPESlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_AFI_SAFI_TYPESliceWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_AFI_SAFI_TYPESlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_BGP_CAPABILITYSlice is a []E_BgpTypes_BGP_CAPABILITY with a corresponding timestamp.
type QualifiedE_BgpTypes_BGP_CAPABILITYSlice struct {
	*genutil.Metadata
	val     []E_BgpTypes_BGP_CAPABILITY // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_BGP_CAPABILITYSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_BgpTypes_BGP_CAPABILITY sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_BGP_CAPABILITYSlice) Val(t testing.TB) []E_BgpTypes_BGP_CAPABILITY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_BgpTypes_BGP_CAPABILITY sample.
func (q *QualifiedE_BgpTypes_BGP_CAPABILITYSlice) SetVal(v []E_BgpTypes_BGP_CAPABILITY) *QualifiedE_BgpTypes_BGP_CAPABILITYSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_BGP_CAPABILITYSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_BGP_CAPABILITYSlice is a telemetry Collection whose Await method returns a slice of []E_BgpTypes_BGP_CAPABILITY samples.
type CollectionE_BgpTypes_BGP_CAPABILITYSlice struct {
	W    *E_BgpTypes_BGP_CAPABILITYSliceWatcher
	Data []*QualifiedE_BgpTypes_BGP_CAPABILITYSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_BGP_CAPABILITYSlice) Await(t testing.TB) []*QualifiedE_BgpTypes_BGP_CAPABILITYSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_BGP_CAPABILITYSliceWatcher observes a stream of []E_BgpTypes_BGP_CAPABILITY samples.
type E_BgpTypes_BGP_CAPABILITYSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_BGP_CAPABILITYSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_BGP_CAPABILITYSliceWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_BGP_CAPABILITYSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Capability_FlagsSlice is a []E_Capability_Flags with a corresponding timestamp.
type QualifiedE_Capability_FlagsSlice struct {
	*genutil.Metadata
	val     []E_Capability_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_Capability_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Capability_Flags sample, erroring out if not present.
func (q *QualifiedE_Capability_FlagsSlice) Val(t testing.TB) []E_Capability_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Capability_Flags sample.
func (q *QualifiedE_Capability_FlagsSlice) SetVal(v []E_Capability_Flags) *QualifiedE_Capability_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Capability_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Capability_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_Capability_Flags samples.
type CollectionE_Capability_FlagsSlice struct {
	W    *E_Capability_FlagsSliceWatcher
	Data []*QualifiedE_Capability_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Capability_FlagsSlice) Await(t testing.TB) []*QualifiedE_Capability_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Capability_FlagsSliceWatcher observes a stream of []E_Capability_Flags samples.
type E_Capability_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Capability_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Capability_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_Capability_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Flags_FlagsSlice is a []E_Flags_Flags with a corresponding timestamp.
type QualifiedE_Flags_FlagsSlice struct {
	*genutil.Metadata
	val     []E_Flags_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_Flags_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Flags_Flags sample, erroring out if not present.
func (q *QualifiedE_Flags_FlagsSlice) Val(t testing.TB) []E_Flags_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Flags_Flags sample.
func (q *QualifiedE_Flags_FlagsSlice) SetVal(v []E_Flags_Flags) *QualifiedE_Flags_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Flags_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Flags_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_Flags_Flags samples.
type CollectionE_Flags_FlagsSlice struct {
	W    *E_Flags_FlagsSliceWatcher
	Data []*QualifiedE_Flags_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Flags_FlagsSlice) Await(t testing.TB) []*QualifiedE_Flags_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Flags_FlagsSliceWatcher observes a stream of []E_Flags_Flags samples.
type E_Flags_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Flags_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Flags_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_Flags_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Ipv4Srlg_FlagsSlice is a []E_Ipv4Srlg_Flags with a corresponding timestamp.
type QualifiedE_Ipv4Srlg_FlagsSlice struct {
	*genutil.Metadata
	val     []E_Ipv4Srlg_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_Ipv4Srlg_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Ipv4Srlg_Flags sample, erroring out if not present.
func (q *QualifiedE_Ipv4Srlg_FlagsSlice) Val(t testing.TB) []E_Ipv4Srlg_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Ipv4Srlg_Flags sample.
func (q *QualifiedE_Ipv4Srlg_FlagsSlice) SetVal(v []E_Ipv4Srlg_Flags) *QualifiedE_Ipv4Srlg_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Ipv4Srlg_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Ipv4Srlg_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_Ipv4Srlg_Flags samples.
type CollectionE_Ipv4Srlg_FlagsSlice struct {
	W    *E_Ipv4Srlg_FlagsSliceWatcher
	Data []*QualifiedE_Ipv4Srlg_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Ipv4Srlg_FlagsSlice) Await(t testing.TB) []*QualifiedE_Ipv4Srlg_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Ipv4Srlg_FlagsSliceWatcher observes a stream of []E_Ipv4Srlg_Flags samples.
type E_Ipv4Srlg_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Ipv4Srlg_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Ipv4Srlg_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_Ipv4Srlg_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Ipv6Srlg_FlagsSlice is a []E_Ipv6Srlg_Flags with a corresponding timestamp.
type QualifiedE_Ipv6Srlg_FlagsSlice struct {
	*genutil.Metadata
	val     []E_Ipv6Srlg_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_Ipv6Srlg_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Ipv6Srlg_Flags sample, erroring out if not present.
func (q *QualifiedE_Ipv6Srlg_FlagsSlice) Val(t testing.TB) []E_Ipv6Srlg_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Ipv6Srlg_Flags sample.
func (q *QualifiedE_Ipv6Srlg_FlagsSlice) SetVal(v []E_Ipv6Srlg_Flags) *QualifiedE_Ipv6Srlg_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Ipv6Srlg_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Ipv6Srlg_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_Ipv6Srlg_Flags samples.
type CollectionE_Ipv6Srlg_FlagsSlice struct {
	W    *E_Ipv6Srlg_FlagsSliceWatcher
	Data []*QualifiedE_Ipv6Srlg_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Ipv6Srlg_FlagsSlice) Await(t testing.TB) []*QualifiedE_Ipv6Srlg_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Ipv6Srlg_FlagsSliceWatcher observes a stream of []E_Ipv6Srlg_Flags samples.
type E_Ipv6Srlg_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Ipv6Srlg_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Ipv6Srlg_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_Ipv6Srlg_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_IsisTypes_AFI_SAFI_TYPESlice is a []E_IsisTypes_AFI_SAFI_TYPE with a corresponding timestamp.
type QualifiedE_IsisTypes_AFI_SAFI_TYPESlice struct {
	*genutil.Metadata
	val     []E_IsisTypes_AFI_SAFI_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_IsisTypes_AFI_SAFI_TYPESlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_IsisTypes_AFI_SAFI_TYPE sample, erroring out if not present.
func (q *QualifiedE_IsisTypes_AFI_SAFI_TYPESlice) Val(t testing.TB) []E_IsisTypes_AFI_SAFI_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_IsisTypes_AFI_SAFI_TYPE sample.
func (q *QualifiedE_IsisTypes_AFI_SAFI_TYPESlice) SetVal(v []E_IsisTypes_AFI_SAFI_TYPE) *QualifiedE_IsisTypes_AFI_SAFI_TYPESlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_IsisTypes_AFI_SAFI_TYPESlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_IsisTypes_AFI_SAFI_TYPESlice is a telemetry Collection whose Await method returns a slice of []E_IsisTypes_AFI_SAFI_TYPE samples.
type CollectionE_IsisTypes_AFI_SAFI_TYPESlice struct {
	W    *E_IsisTypes_AFI_SAFI_TYPESliceWatcher
	Data []*QualifiedE_IsisTypes_AFI_SAFI_TYPESlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_IsisTypes_AFI_SAFI_TYPESlice) Await(t testing.TB) []*QualifiedE_IsisTypes_AFI_SAFI_TYPESlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_IsisTypes_AFI_SAFI_TYPESliceWatcher observes a stream of []E_IsisTypes_AFI_SAFI_TYPE samples.
type E_IsisTypes_AFI_SAFI_TYPESliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_IsisTypes_AFI_SAFI_TYPESlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_IsisTypes_AFI_SAFI_TYPESliceWatcher) Await(t testing.TB) (*QualifiedE_IsisTypes_AFI_SAFI_TYPESlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Isis_IsisMetricFlagsSlice is a []E_Isis_IsisMetricFlags with a corresponding timestamp.
type QualifiedE_Isis_IsisMetricFlagsSlice struct {
	*genutil.Metadata
	val     []E_Isis_IsisMetricFlags // val is the sample value.
	present bool
}

func (q *QualifiedE_Isis_IsisMetricFlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Isis_IsisMetricFlags sample, erroring out if not present.
func (q *QualifiedE_Isis_IsisMetricFlagsSlice) Val(t testing.TB) []E_Isis_IsisMetricFlags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Isis_IsisMetricFlags sample.
func (q *QualifiedE_Isis_IsisMetricFlagsSlice) SetVal(v []E_Isis_IsisMetricFlags) *QualifiedE_Isis_IsisMetricFlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Isis_IsisMetricFlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Isis_IsisMetricFlagsSlice is a telemetry Collection whose Await method returns a slice of []E_Isis_IsisMetricFlags samples.
type CollectionE_Isis_IsisMetricFlagsSlice struct {
	W    *E_Isis_IsisMetricFlagsSliceWatcher
	Data []*QualifiedE_Isis_IsisMetricFlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Isis_IsisMetricFlagsSlice) Await(t testing.TB) []*QualifiedE_Isis_IsisMetricFlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Isis_IsisMetricFlagsSliceWatcher observes a stream of []E_Isis_IsisMetricFlags samples.
type E_Isis_IsisMetricFlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Isis_IsisMetricFlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Isis_IsisMetricFlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_Isis_IsisMetricFlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LanAdjacencySid_FlagsSlice is a []E_LanAdjacencySid_Flags with a corresponding timestamp.
type QualifiedE_LanAdjacencySid_FlagsSlice struct {
	*genutil.Metadata
	val     []E_LanAdjacencySid_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_LanAdjacencySid_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_LanAdjacencySid_Flags sample, erroring out if not present.
func (q *QualifiedE_LanAdjacencySid_FlagsSlice) Val(t testing.TB) []E_LanAdjacencySid_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_LanAdjacencySid_Flags sample.
func (q *QualifiedE_LanAdjacencySid_FlagsSlice) SetVal(v []E_LanAdjacencySid_Flags) *QualifiedE_LanAdjacencySid_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LanAdjacencySid_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LanAdjacencySid_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_LanAdjacencySid_Flags samples.
type CollectionE_LanAdjacencySid_FlagsSlice struct {
	W    *E_LanAdjacencySid_FlagsSliceWatcher
	Data []*QualifiedE_LanAdjacencySid_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LanAdjacencySid_FlagsSlice) Await(t testing.TB) []*QualifiedE_LanAdjacencySid_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LanAdjacencySid_FlagsSliceWatcher observes a stream of []E_LanAdjacencySid_Flags samples.
type E_LanAdjacencySid_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LanAdjacencySid_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LanAdjacencySid_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_LanAdjacencySid_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LinkAttributes_LocalProtectionSlice is a []E_LinkAttributes_LocalProtection with a corresponding timestamp.
type QualifiedE_LinkAttributes_LocalProtectionSlice struct {
	*genutil.Metadata
	val     []E_LinkAttributes_LocalProtection // val is the sample value.
	present bool
}

func (q *QualifiedE_LinkAttributes_LocalProtectionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_LinkAttributes_LocalProtection sample, erroring out if not present.
func (q *QualifiedE_LinkAttributes_LocalProtectionSlice) Val(t testing.TB) []E_LinkAttributes_LocalProtection {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_LinkAttributes_LocalProtection sample.
func (q *QualifiedE_LinkAttributes_LocalProtectionSlice) SetVal(v []E_LinkAttributes_LocalProtection) *QualifiedE_LinkAttributes_LocalProtectionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LinkAttributes_LocalProtectionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LinkAttributes_LocalProtectionSlice is a telemetry Collection whose Await method returns a slice of []E_LinkAttributes_LocalProtection samples.
type CollectionE_LinkAttributes_LocalProtectionSlice struct {
	W    *E_LinkAttributes_LocalProtectionSliceWatcher
	Data []*QualifiedE_LinkAttributes_LocalProtectionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LinkAttributes_LocalProtectionSlice) Await(t testing.TB) []*QualifiedE_LinkAttributes_LocalProtectionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LinkAttributes_LocalProtectionSliceWatcher observes a stream of []E_LinkAttributes_LocalProtection samples.
type E_LinkAttributes_LocalProtectionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LinkAttributes_LocalProtectionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LinkAttributes_LocalProtectionSliceWatcher) Await(t testing.TB) (*QualifiedE_LinkAttributes_LocalProtectionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LinkProtectionType_TypeSlice is a []E_LinkProtectionType_Type with a corresponding timestamp.
type QualifiedE_LinkProtectionType_TypeSlice struct {
	*genutil.Metadata
	val     []E_LinkProtectionType_Type // val is the sample value.
	present bool
}

func (q *QualifiedE_LinkProtectionType_TypeSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_LinkProtectionType_Type sample, erroring out if not present.
func (q *QualifiedE_LinkProtectionType_TypeSlice) Val(t testing.TB) []E_LinkProtectionType_Type {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_LinkProtectionType_Type sample.
func (q *QualifiedE_LinkProtectionType_TypeSlice) SetVal(v []E_LinkProtectionType_Type) *QualifiedE_LinkProtectionType_TypeSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LinkProtectionType_TypeSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LinkProtectionType_TypeSlice is a telemetry Collection whose Await method returns a slice of []E_LinkProtectionType_Type samples.
type CollectionE_LinkProtectionType_TypeSlice struct {
	W    *E_LinkProtectionType_TypeSliceWatcher
	Data []*QualifiedE_LinkProtectionType_TypeSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LinkProtectionType_TypeSlice) Await(t testing.TB) []*QualifiedE_LinkProtectionType_TypeSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LinkProtectionType_TypeSliceWatcher observes a stream of []E_LinkProtectionType_Type samples.
type E_LinkProtectionType_TypeSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LinkProtectionType_TypeSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LinkProtectionType_TypeSliceWatcher) Await(t testing.TB) (*QualifiedE_LinkProtectionType_TypeSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_LldpTypes_LLDP_TLVSlice is a []E_LldpTypes_LLDP_TLV with a corresponding timestamp.
type QualifiedE_LldpTypes_LLDP_TLVSlice struct {
	*genutil.Metadata
	val     []E_LldpTypes_LLDP_TLV // val is the sample value.
	present bool
}

func (q *QualifiedE_LldpTypes_LLDP_TLVSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_LldpTypes_LLDP_TLV sample, erroring out if not present.
func (q *QualifiedE_LldpTypes_LLDP_TLVSlice) Val(t testing.TB) []E_LldpTypes_LLDP_TLV {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_LldpTypes_LLDP_TLV sample.
func (q *QualifiedE_LldpTypes_LLDP_TLVSlice) SetVal(v []E_LldpTypes_LLDP_TLV) *QualifiedE_LldpTypes_LLDP_TLVSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_LldpTypes_LLDP_TLVSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_LldpTypes_LLDP_TLVSlice is a telemetry Collection whose Await method returns a slice of []E_LldpTypes_LLDP_TLV samples.
type CollectionE_LldpTypes_LLDP_TLVSlice struct {
	W    *E_LldpTypes_LLDP_TLVSliceWatcher
	Data []*QualifiedE_LldpTypes_LLDP_TLVSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_LldpTypes_LLDP_TLVSlice) Await(t testing.TB) []*QualifiedE_LldpTypes_LLDP_TLVSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_LldpTypes_LLDP_TLVSliceWatcher observes a stream of []E_LldpTypes_LLDP_TLV samples.
type E_LldpTypes_LLDP_TLVSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_LldpTypes_LLDP_TLVSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_LldpTypes_LLDP_TLVSliceWatcher) Await(t testing.TB) (*QualifiedE_LldpTypes_LLDP_TLVSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lsp_FlagsSlice is a []E_Lsp_Flags with a corresponding timestamp.
type QualifiedE_Lsp_FlagsSlice struct {
	*genutil.Metadata
	val     []E_Lsp_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_Lsp_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Lsp_Flags sample, erroring out if not present.
func (q *QualifiedE_Lsp_FlagsSlice) Val(t testing.TB) []E_Lsp_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Lsp_Flags sample.
func (q *QualifiedE_Lsp_FlagsSlice) SetVal(v []E_Lsp_Flags) *QualifiedE_Lsp_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lsp_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lsp_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_Lsp_Flags samples.
type CollectionE_Lsp_FlagsSlice struct {
	W    *E_Lsp_FlagsSliceWatcher
	Data []*QualifiedE_Lsp_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lsp_FlagsSlice) Await(t testing.TB) []*QualifiedE_Lsp_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lsp_FlagsSliceWatcher observes a stream of []E_Lsp_Flags samples.
type E_Lsp_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lsp_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lsp_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_Lsp_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice is a []E_MplsTypes_PATH_SETUP_PROTOCOL with a corresponding timestamp.
type QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice struct {
	*genutil.Metadata
	val     []E_MplsTypes_PATH_SETUP_PROTOCOL // val is the sample value.
	present bool
}

func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_MplsTypes_PATH_SETUP_PROTOCOL sample, erroring out if not present.
func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice) Val(t testing.TB) []E_MplsTypes_PATH_SETUP_PROTOCOL {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_MplsTypes_PATH_SETUP_PROTOCOL sample.
func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice) SetVal(v []E_MplsTypes_PATH_SETUP_PROTOCOL) *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_MplsTypes_PATH_SETUP_PROTOCOLSlice is a telemetry Collection whose Await method returns a slice of []E_MplsTypes_PATH_SETUP_PROTOCOL samples.
type CollectionE_MplsTypes_PATH_SETUP_PROTOCOLSlice struct {
	W    *E_MplsTypes_PATH_SETUP_PROTOCOLSliceWatcher
	Data []*QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_MplsTypes_PATH_SETUP_PROTOCOLSlice) Await(t testing.TB) []*QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_MplsTypes_PATH_SETUP_PROTOCOLSliceWatcher observes a stream of []E_MplsTypes_PATH_SETUP_PROTOCOL samples.
type E_MplsTypes_PATH_SETUP_PROTOCOLSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_MplsTypes_PATH_SETUP_PROTOCOLSliceWatcher) Await(t testing.TB) (*QualifiedE_MplsTypes_PATH_SETUP_PROTOCOLSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Nlpid_NlpidSlice is a []E_Nlpid_Nlpid with a corresponding timestamp.
type QualifiedE_Nlpid_NlpidSlice struct {
	*genutil.Metadata
	val     []E_Nlpid_Nlpid // val is the sample value.
	present bool
}

func (q *QualifiedE_Nlpid_NlpidSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Nlpid_Nlpid sample, erroring out if not present.
func (q *QualifiedE_Nlpid_NlpidSlice) Val(t testing.TB) []E_Nlpid_Nlpid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Nlpid_Nlpid sample.
func (q *QualifiedE_Nlpid_NlpidSlice) SetVal(v []E_Nlpid_Nlpid) *QualifiedE_Nlpid_NlpidSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Nlpid_NlpidSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Nlpid_NlpidSlice is a telemetry Collection whose Await method returns a slice of []E_Nlpid_Nlpid samples.
type CollectionE_Nlpid_NlpidSlice struct {
	W    *E_Nlpid_NlpidSliceWatcher
	Data []*QualifiedE_Nlpid_NlpidSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Nlpid_NlpidSlice) Await(t testing.TB) []*QualifiedE_Nlpid_NlpidSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Nlpid_NlpidSliceWatcher observes a stream of []E_Nlpid_Nlpid samples.
type E_Nlpid_NlpidSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Nlpid_NlpidSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Nlpid_NlpidSliceWatcher) Await(t testing.TB) (*QualifiedE_Nlpid_NlpidSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice is a []E_OspfTypes_MAX_METRIC_INCLUDE with a corresponding timestamp.
type QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice struct {
	*genutil.Metadata
	val     []E_OspfTypes_MAX_METRIC_INCLUDE // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_OspfTypes_MAX_METRIC_INCLUDE sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice) Val(t testing.TB) []E_OspfTypes_MAX_METRIC_INCLUDE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_OspfTypes_MAX_METRIC_INCLUDE sample.
func (q *QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice) SetVal(v []E_OspfTypes_MAX_METRIC_INCLUDE) *QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_MAX_METRIC_INCLUDESlice is a telemetry Collection whose Await method returns a slice of []E_OspfTypes_MAX_METRIC_INCLUDE samples.
type CollectionE_OspfTypes_MAX_METRIC_INCLUDESlice struct {
	W    *E_OspfTypes_MAX_METRIC_INCLUDESliceWatcher
	Data []*QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_MAX_METRIC_INCLUDESlice) Await(t testing.TB) []*QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_MAX_METRIC_INCLUDESliceWatcher observes a stream of []E_OspfTypes_MAX_METRIC_INCLUDE samples.
type E_OspfTypes_MAX_METRIC_INCLUDESliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_MAX_METRIC_INCLUDESliceWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_MAX_METRIC_INCLUDESlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice is a []E_OspfTypes_MAX_METRIC_TRIGGER with a corresponding timestamp.
type QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice struct {
	*genutil.Metadata
	val     []E_OspfTypes_MAX_METRIC_TRIGGER // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_OspfTypes_MAX_METRIC_TRIGGER sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice) Val(t testing.TB) []E_OspfTypes_MAX_METRIC_TRIGGER {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_OspfTypes_MAX_METRIC_TRIGGER sample.
func (q *QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice) SetVal(v []E_OspfTypes_MAX_METRIC_TRIGGER) *QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_MAX_METRIC_TRIGGERSlice is a telemetry Collection whose Await method returns a slice of []E_OspfTypes_MAX_METRIC_TRIGGER samples.
type CollectionE_OspfTypes_MAX_METRIC_TRIGGERSlice struct {
	W    *E_OspfTypes_MAX_METRIC_TRIGGERSliceWatcher
	Data []*QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_MAX_METRIC_TRIGGERSlice) Await(t testing.TB) []*QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_MAX_METRIC_TRIGGERSliceWatcher observes a stream of []E_OspfTypes_MAX_METRIC_TRIGGER samples.
type E_OspfTypes_MAX_METRIC_TRIGGERSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_MAX_METRIC_TRIGGERSliceWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_MAX_METRIC_TRIGGERSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OspfTypes_SR_ALGORITHMSlice is a []E_OspfTypes_SR_ALGORITHM with a corresponding timestamp.
type QualifiedE_OspfTypes_SR_ALGORITHMSlice struct {
	*genutil.Metadata
	val     []E_OspfTypes_SR_ALGORITHM // val is the sample value.
	present bool
}

func (q *QualifiedE_OspfTypes_SR_ALGORITHMSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_OspfTypes_SR_ALGORITHM sample, erroring out if not present.
func (q *QualifiedE_OspfTypes_SR_ALGORITHMSlice) Val(t testing.TB) []E_OspfTypes_SR_ALGORITHM {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_OspfTypes_SR_ALGORITHM sample.
func (q *QualifiedE_OspfTypes_SR_ALGORITHMSlice) SetVal(v []E_OspfTypes_SR_ALGORITHM) *QualifiedE_OspfTypes_SR_ALGORITHMSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OspfTypes_SR_ALGORITHMSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OspfTypes_SR_ALGORITHMSlice is a telemetry Collection whose Await method returns a slice of []E_OspfTypes_SR_ALGORITHM samples.
type CollectionE_OspfTypes_SR_ALGORITHMSlice struct {
	W    *E_OspfTypes_SR_ALGORITHMSliceWatcher
	Data []*QualifiedE_OspfTypes_SR_ALGORITHMSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OspfTypes_SR_ALGORITHMSlice) Await(t testing.TB) []*QualifiedE_OspfTypes_SR_ALGORITHMSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OspfTypes_SR_ALGORITHMSliceWatcher observes a stream of []E_OspfTypes_SR_ALGORITHM samples.
type E_OspfTypes_SR_ALGORITHMSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OspfTypes_SR_ALGORITHMSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OspfTypes_SR_ALGORITHMSliceWatcher) Await(t testing.TB) (*QualifiedE_OspfTypes_SR_ALGORITHMSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PacketMatchTypes_TCP_FLAGSSlice is a []E_PacketMatchTypes_TCP_FLAGS with a corresponding timestamp.
type QualifiedE_PacketMatchTypes_TCP_FLAGSSlice struct {
	*genutil.Metadata
	val     []E_PacketMatchTypes_TCP_FLAGS // val is the sample value.
	present bool
}

func (q *QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_PacketMatchTypes_TCP_FLAGS sample, erroring out if not present.
func (q *QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) Val(t testing.TB) []E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_PacketMatchTypes_TCP_FLAGS sample.
func (q *QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) SetVal(v []E_PacketMatchTypes_TCP_FLAGS) *QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PacketMatchTypes_TCP_FLAGSSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PacketMatchTypes_TCP_FLAGSSlice is a telemetry Collection whose Await method returns a slice of []E_PacketMatchTypes_TCP_FLAGS samples.
type CollectionE_PacketMatchTypes_TCP_FLAGSSlice struct {
	W    *E_PacketMatchTypes_TCP_FLAGSSliceWatcher
	Data []*QualifiedE_PacketMatchTypes_TCP_FLAGSSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PacketMatchTypes_TCP_FLAGSSlice) Await(t testing.TB) []*QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PacketMatchTypes_TCP_FLAGSSliceWatcher observes a stream of []E_PacketMatchTypes_TCP_FLAGS samples.
type E_PacketMatchTypes_TCP_FLAGSSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PacketMatchTypes_TCP_FLAGSSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PacketMatchTypes_TCP_FLAGSSliceWatcher) Await(t testing.TB) (*QualifiedE_PacketMatchTypes_TCP_FLAGSSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_PrefixSid_FlagsSlice is a []E_PrefixSid_Flags with a corresponding timestamp.
type QualifiedE_PrefixSid_FlagsSlice struct {
	*genutil.Metadata
	val     []E_PrefixSid_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_PrefixSid_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_PrefixSid_Flags sample, erroring out if not present.
func (q *QualifiedE_PrefixSid_FlagsSlice) Val(t testing.TB) []E_PrefixSid_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_PrefixSid_Flags sample.
func (q *QualifiedE_PrefixSid_FlagsSlice) SetVal(v []E_PrefixSid_Flags) *QualifiedE_PrefixSid_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_PrefixSid_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_PrefixSid_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_PrefixSid_Flags samples.
type CollectionE_PrefixSid_FlagsSlice struct {
	W    *E_PrefixSid_FlagsSliceWatcher
	Data []*QualifiedE_PrefixSid_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_PrefixSid_FlagsSlice) Await(t testing.TB) []*QualifiedE_PrefixSid_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_PrefixSid_FlagsSliceWatcher observes a stream of []E_PrefixSid_Flags samples.
type E_PrefixSid_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_PrefixSid_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_PrefixSid_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_PrefixSid_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice is a []E_SegmentRoutingAlgorithms_Algorithm with a corresponding timestamp.
type QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice struct {
	*genutil.Metadata
	val     []E_SegmentRoutingAlgorithms_Algorithm // val is the sample value.
	present bool
}

func (q *QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_SegmentRoutingAlgorithms_Algorithm sample, erroring out if not present.
func (q *QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice) Val(t testing.TB) []E_SegmentRoutingAlgorithms_Algorithm {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_SegmentRoutingAlgorithms_Algorithm sample.
func (q *QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice) SetVal(v []E_SegmentRoutingAlgorithms_Algorithm) *QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SegmentRoutingAlgorithms_AlgorithmSlice is a telemetry Collection whose Await method returns a slice of []E_SegmentRoutingAlgorithms_Algorithm samples.
type CollectionE_SegmentRoutingAlgorithms_AlgorithmSlice struct {
	W    *E_SegmentRoutingAlgorithms_AlgorithmSliceWatcher
	Data []*QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SegmentRoutingAlgorithms_AlgorithmSlice) Await(t testing.TB) []*QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SegmentRoutingAlgorithms_AlgorithmSliceWatcher observes a stream of []E_SegmentRoutingAlgorithms_Algorithm samples.
type E_SegmentRoutingAlgorithms_AlgorithmSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SegmentRoutingAlgorithms_AlgorithmSliceWatcher) Await(t testing.TB) (*QualifiedE_SegmentRoutingAlgorithms_AlgorithmSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SegmentRoutingCapability_FlagsSlice is a []E_SegmentRoutingCapability_Flags with a corresponding timestamp.
type QualifiedE_SegmentRoutingCapability_FlagsSlice struct {
	*genutil.Metadata
	val     []E_SegmentRoutingCapability_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_SegmentRoutingCapability_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_SegmentRoutingCapability_Flags sample, erroring out if not present.
func (q *QualifiedE_SegmentRoutingCapability_FlagsSlice) Val(t testing.TB) []E_SegmentRoutingCapability_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_SegmentRoutingCapability_Flags sample.
func (q *QualifiedE_SegmentRoutingCapability_FlagsSlice) SetVal(v []E_SegmentRoutingCapability_Flags) *QualifiedE_SegmentRoutingCapability_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SegmentRoutingCapability_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SegmentRoutingCapability_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_SegmentRoutingCapability_Flags samples.
type CollectionE_SegmentRoutingCapability_FlagsSlice struct {
	W    *E_SegmentRoutingCapability_FlagsSliceWatcher
	Data []*QualifiedE_SegmentRoutingCapability_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SegmentRoutingCapability_FlagsSlice) Await(t testing.TB) []*QualifiedE_SegmentRoutingCapability_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SegmentRoutingCapability_FlagsSliceWatcher observes a stream of []E_SegmentRoutingCapability_Flags samples.
type E_SegmentRoutingCapability_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SegmentRoutingCapability_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SegmentRoutingCapability_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_SegmentRoutingCapability_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_SystemGrpc_GRPC_SERVICESlice is a []E_SystemGrpc_GRPC_SERVICE with a corresponding timestamp.
type QualifiedE_SystemGrpc_GRPC_SERVICESlice struct {
	*genutil.Metadata
	val     []E_SystemGrpc_GRPC_SERVICE // val is the sample value.
	present bool
}

func (q *QualifiedE_SystemGrpc_GRPC_SERVICESlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_SystemGrpc_GRPC_SERVICE sample, erroring out if not present.
func (q *QualifiedE_SystemGrpc_GRPC_SERVICESlice) Val(t testing.TB) []E_SystemGrpc_GRPC_SERVICE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_SystemGrpc_GRPC_SERVICE sample.
func (q *QualifiedE_SystemGrpc_GRPC_SERVICESlice) SetVal(v []E_SystemGrpc_GRPC_SERVICE) *QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_SystemGrpc_GRPC_SERVICESlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_SystemGrpc_GRPC_SERVICESlice is a telemetry Collection whose Await method returns a slice of []E_SystemGrpc_GRPC_SERVICE samples.
type CollectionE_SystemGrpc_GRPC_SERVICESlice struct {
	W    *E_SystemGrpc_GRPC_SERVICESliceWatcher
	Data []*QualifiedE_SystemGrpc_GRPC_SERVICESlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_SystemGrpc_GRPC_SERVICESlice) Await(t testing.TB) []*QualifiedE_SystemGrpc_GRPC_SERVICESlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_SystemGrpc_GRPC_SERVICESliceWatcher observes a stream of []E_SystemGrpc_GRPC_SERVICE samples.
type E_SystemGrpc_GRPC_SERVICESliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_SystemGrpc_GRPC_SERVICESlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_SystemGrpc_GRPC_SERVICESliceWatcher) Await(t testing.TB) (*QualifiedE_SystemGrpc_GRPC_SERVICESlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Types_ADDRESS_FAMILYSlice is a []E_Types_ADDRESS_FAMILY with a corresponding timestamp.
type QualifiedE_Types_ADDRESS_FAMILYSlice struct {
	*genutil.Metadata
	val     []E_Types_ADDRESS_FAMILY // val is the sample value.
	present bool
}

func (q *QualifiedE_Types_ADDRESS_FAMILYSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Types_ADDRESS_FAMILY sample, erroring out if not present.
func (q *QualifiedE_Types_ADDRESS_FAMILYSlice) Val(t testing.TB) []E_Types_ADDRESS_FAMILY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Types_ADDRESS_FAMILY sample.
func (q *QualifiedE_Types_ADDRESS_FAMILYSlice) SetVal(v []E_Types_ADDRESS_FAMILY) *QualifiedE_Types_ADDRESS_FAMILYSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Types_ADDRESS_FAMILYSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Types_ADDRESS_FAMILYSlice is a telemetry Collection whose Await method returns a slice of []E_Types_ADDRESS_FAMILY samples.
type CollectionE_Types_ADDRESS_FAMILYSlice struct {
	W    *E_Types_ADDRESS_FAMILYSliceWatcher
	Data []*QualifiedE_Types_ADDRESS_FAMILYSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Types_ADDRESS_FAMILYSlice) Await(t testing.TB) []*QualifiedE_Types_ADDRESS_FAMILYSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Types_ADDRESS_FAMILYSliceWatcher observes a stream of []E_Types_ADDRESS_FAMILY samples.
type E_Types_ADDRESS_FAMILYSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Types_ADDRESS_FAMILYSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Types_ADDRESS_FAMILYSliceWatcher) Await(t testing.TB) (*QualifiedE_Types_ADDRESS_FAMILYSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice is a []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union with a corresponding timestamp.
type QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice struct {
	*genutil.Metadata
	val     []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union sample, erroring out if not present.
func (q *QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice) Val(t testing.TB) []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union sample.
func (q *QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice) SetVal(v []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union) *QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice is a telemetry Collection whose Await method returns a slice of []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union samples.
type CollectionInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice struct {
	W    *Interface_Aggregation_SwitchedVlan_TrunkVlans_UnionSliceWatcher
	Data []*QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice) Await(t testing.TB) []*QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Aggregation_SwitchedVlan_TrunkVlans_UnionSliceWatcher observes a stream of []Interface_Aggregation_SwitchedVlan_TrunkVlans_Union samples.
type Interface_Aggregation_SwitchedVlan_TrunkVlans_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Aggregation_SwitchedVlan_TrunkVlans_UnionSliceWatcher) Await(t testing.TB) (*QualifiedInterface_Aggregation_SwitchedVlan_TrunkVlans_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice is a []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union with a corresponding timestamp.
type QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice struct {
	*genutil.Metadata
	val     []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union sample, erroring out if not present.
func (q *QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice) Val(t testing.TB) []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union sample.
func (q *QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice) SetVal(v []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union) *QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice is a telemetry Collection whose Await method returns a slice of []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union samples.
type CollectionInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice struct {
	W    *Interface_Ethernet_SwitchedVlan_TrunkVlans_UnionSliceWatcher
	Data []*QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice) Await(t testing.TB) []*QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Ethernet_SwitchedVlan_TrunkVlans_UnionSliceWatcher observes a stream of []Interface_Ethernet_SwitchedVlan_TrunkVlans_Union samples.
type Interface_Ethernet_SwitchedVlan_TrunkVlans_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Ethernet_SwitchedVlan_TrunkVlans_UnionSliceWatcher) Await(t testing.TB) (*QualifiedInterface_Ethernet_SwitchedVlan_TrunkVlans_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice is a []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice) Val(t testing.TB) []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union sample.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice) SetVal(v []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union) *QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union samples.
type CollectionNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice struct {
	W    *NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSliceWatcher observes a stream of []NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_Union samples.
type NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_LabelEntry_PoppedMplsLabelStack_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice is a []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice) Val(t testing.TB) []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union sample.
func (q *QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice) SetVal(v []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union) *QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union samples.
type CollectionNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice struct {
	W    *NetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSliceWatcher observes a stream of []NetworkInstance_Afts_NextHop_PushedMplsLabelStack_Union samples.
type NetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Afts_NextHop_PushedMplsLabelStack_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice is a []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice) Val(t testing.TB) []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice) SetVal(v []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union) *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice struct {
	W    *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher observes a stream of []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_Union samples.
type NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ExportRouteTarget_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice is a []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice) Val(t testing.TB) []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union sample.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice) SetVal(v []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union) *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union samples.
type CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice struct {
	W    *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher observes a stream of []NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_Union samples.
type NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Evpn_EvpnInstance_ImportExportPolicy_ImportRouteTarget_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice is a []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union with a corresponding timestamp.
type QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice) Val(t testing.TB) []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union sample.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice) SetVal(v []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union) *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union samples.
type CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice struct {
	W    *NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher observes a stream of []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_Union samples.
type NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ExportRouteTarget_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice is a []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union with a corresponding timestamp.
type QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice) Val(t testing.TB) []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union sample.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice) SetVal(v []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union) *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union samples.
type CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice struct {
	W    *NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher observes a stream of []NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_Union samples.
type NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_InterInstancePolicies_ImportExportPolicy_ImportRouteTarget_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice is a []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice) Val(t testing.TB) []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice) SetVal(v []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union) *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSliceWatcher observes a stream of []NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union samples.
type NetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_Community_Community_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice is a []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice) Val(t testing.TB) []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice) SetVal(v []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union) *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSliceWatcher observes a stream of []NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union samples.
type NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice is a []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union with a corresponding timestamp.
type QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice struct {
	*genutil.Metadata
	val     []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union sample, erroring out if not present.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice) Val(t testing.TB) []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union sample.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice) SetVal(v []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union) *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice is a telemetry Collection whose Await method returns a slice of []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union samples.
type CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice struct {
	W    *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSliceWatcher
	Data []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice) Await(t testing.TB) []*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSliceWatcher observes a stream of []NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_Union samples.
type NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSliceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_SegmentRouting_TePolicy_CandidatePath_SegmentList_NextHop_PushedMplsLabelStack_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice is a []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice struct {
	*genutil.Metadata
	val     []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) Val(t testing.TB) []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union sample.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) SetVal(v []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union) *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice is a telemetry Collection whose Await method returns a slice of []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union samples.
type CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice struct {
	W    *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher observes a stream of []RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union samples.
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSliceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice is a []RoutingPolicy_DefinedSets_TagSet_TagValue_Union with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice struct {
	*genutil.Metadata
	val     []RoutingPolicy_DefinedSets_TagSet_TagValue_Union // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []RoutingPolicy_DefinedSets_TagSet_TagValue_Union sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) Val(t testing.TB) []RoutingPolicy_DefinedSets_TagSet_TagValue_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []RoutingPolicy_DefinedSets_TagSet_TagValue_Union sample.
func (q *QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) SetVal(v []RoutingPolicy_DefinedSets_TagSet_TagValue_Union) *QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice is a telemetry Collection whose Await method returns a slice of []RoutingPolicy_DefinedSets_TagSet_TagValue_Union samples.
type CollectionRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice struct {
	W    *RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher observes a stream of []RoutingPolicy_DefinedSets_TagSet_TagValue_Union samples.
type RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_TagSet_TagValue_UnionSliceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice is a []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice struct {
	*genutil.Metadata
	val     []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) Val(t testing.TB) []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) SetVal(v []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice is a telemetry Collection whose Await method returns a slice of []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher observes a stream of []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSliceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice is a []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice struct {
	*genutil.Metadata
	val     []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) Val(t testing.TB) []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) SetVal(v []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice is a telemetry Collection whose Await method returns a slice of []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher observes a stream of []RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSliceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice is a []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice struct {
	*genutil.Metadata
	val     []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) Val(t testing.TB) []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) SetVal(v []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice is a telemetry Collection whose Await method returns a slice of []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher observes a stream of []RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSliceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice is a []System_Aaa_Accounting_AccountingMethod_Union with a corresponding timestamp.
type QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice struct {
	*genutil.Metadata
	val     []System_Aaa_Accounting_AccountingMethod_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []System_Aaa_Accounting_AccountingMethod_Union sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) Val(t testing.TB) []System_Aaa_Accounting_AccountingMethod_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []System_Aaa_Accounting_AccountingMethod_Union sample.
func (q *QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) SetVal(v []System_Aaa_Accounting_AccountingMethod_Union) *QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Accounting_AccountingMethod_UnionSlice is a telemetry Collection whose Await method returns a slice of []System_Aaa_Accounting_AccountingMethod_Union samples.
type CollectionSystem_Aaa_Accounting_AccountingMethod_UnionSlice struct {
	W    *System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher
	Data []*QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Accounting_AccountingMethod_UnionSlice) Await(t testing.TB) []*QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher observes a stream of []System_Aaa_Accounting_AccountingMethod_Union samples.
type System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Accounting_AccountingMethod_UnionSliceWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice is a []System_Aaa_Authentication_AuthenticationMethod_Union with a corresponding timestamp.
type QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice struct {
	*genutil.Metadata
	val     []System_Aaa_Authentication_AuthenticationMethod_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []System_Aaa_Authentication_AuthenticationMethod_Union sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) Val(t testing.TB) []System_Aaa_Authentication_AuthenticationMethod_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []System_Aaa_Authentication_AuthenticationMethod_Union sample.
func (q *QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) SetVal(v []System_Aaa_Authentication_AuthenticationMethod_Union) *QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice is a telemetry Collection whose Await method returns a slice of []System_Aaa_Authentication_AuthenticationMethod_Union samples.
type CollectionSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice struct {
	W    *System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher
	Data []*QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice) Await(t testing.TB) []*QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher observes a stream of []System_Aaa_Authentication_AuthenticationMethod_Union samples.
type System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Authentication_AuthenticationMethod_UnionSliceWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice is a []System_Aaa_Authorization_AuthorizationMethod_Union with a corresponding timestamp.
type QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice struct {
	*genutil.Metadata
	val     []System_Aaa_Authorization_AuthorizationMethod_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []System_Aaa_Authorization_AuthorizationMethod_Union sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) Val(t testing.TB) []System_Aaa_Authorization_AuthorizationMethod_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []System_Aaa_Authorization_AuthorizationMethod_Union sample.
func (q *QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) SetVal(v []System_Aaa_Authorization_AuthorizationMethod_Union) *QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice is a telemetry Collection whose Await method returns a slice of []System_Aaa_Authorization_AuthorizationMethod_Union samples.
type CollectionSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice struct {
	W    *System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher
	Data []*QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice) Await(t testing.TB) []*QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher observes a stream of []System_Aaa_Authorization_AuthorizationMethod_Union samples.
type System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Authorization_AuthorizationMethod_UnionSliceWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice is a []System_GrpcServer_ListenAddresses_Union with a corresponding timestamp.
type QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice struct {
	*genutil.Metadata
	val     []System_GrpcServer_ListenAddresses_Union // val is the sample value.
	present bool
}

func (q *QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []System_GrpcServer_ListenAddresses_Union sample, erroring out if not present.
func (q *QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) Val(t testing.TB) []System_GrpcServer_ListenAddresses_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []System_GrpcServer_ListenAddresses_Union sample.
func (q *QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) SetVal(v []System_GrpcServer_ListenAddresses_Union) *QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_GrpcServer_ListenAddresses_UnionSlice is a telemetry Collection whose Await method returns a slice of []System_GrpcServer_ListenAddresses_Union samples.
type CollectionSystem_GrpcServer_ListenAddresses_UnionSlice struct {
	W    *System_GrpcServer_ListenAddresses_UnionSliceWatcher
	Data []*QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_GrpcServer_ListenAddresses_UnionSlice) Await(t testing.TB) []*QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_GrpcServer_ListenAddresses_UnionSliceWatcher observes a stream of []System_GrpcServer_ListenAddresses_Union samples.
type System_GrpcServer_ListenAddresses_UnionSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_GrpcServer_ListenAddresses_UnionSliceWatcher) Await(t testing.TB) (*QualifiedSystem_GrpcServer_ListenAddresses_UnionSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedStringSlice is a []string with a corresponding timestamp.
type QualifiedStringSlice struct {
	*genutil.Metadata
	val     []string // val is the sample value.
	present bool
}

func (q *QualifiedStringSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []string sample, erroring out if not present.
func (q *QualifiedStringSlice) Val(t testing.TB) []string {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []string sample.
func (q *QualifiedStringSlice) SetVal(v []string) *QualifiedStringSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedStringSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionStringSlice is a telemetry Collection whose Await method returns a slice of []string samples.
type CollectionStringSlice struct {
	W    *StringSliceWatcher
	Data []*QualifiedStringSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionStringSlice) Await(t testing.TB) []*QualifiedStringSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// StringSliceWatcher observes a stream of []string samples.
type StringSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedStringSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *StringSliceWatcher) Await(t testing.TB) (*QualifiedStringSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedUint16Slice is a []uint16 with a corresponding timestamp.
type QualifiedUint16Slice struct {
	*genutil.Metadata
	val     []uint16 // val is the sample value.
	present bool
}

func (q *QualifiedUint16Slice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []uint16 sample, erroring out if not present.
func (q *QualifiedUint16Slice) Val(t testing.TB) []uint16 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []uint16 sample.
func (q *QualifiedUint16Slice) SetVal(v []uint16) *QualifiedUint16Slice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint16Slice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint16Slice is a telemetry Collection whose Await method returns a slice of []uint16 samples.
type CollectionUint16Slice struct {
	W    *Uint16SliceWatcher
	Data []*QualifiedUint16Slice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint16Slice) Await(t testing.TB) []*QualifiedUint16Slice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint16SliceWatcher observes a stream of []uint16 samples.
type Uint16SliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint16Slice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint16SliceWatcher) Await(t testing.TB) (*QualifiedUint16Slice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedUint32Slice is a []uint32 with a corresponding timestamp.
type QualifiedUint32Slice struct {
	*genutil.Metadata
	val     []uint32 // val is the sample value.
	present bool
}

func (q *QualifiedUint32Slice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []uint32 sample, erroring out if not present.
func (q *QualifiedUint32Slice) Val(t testing.TB) []uint32 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []uint32 sample.
func (q *QualifiedUint32Slice) SetVal(v []uint32) *QualifiedUint32Slice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint32Slice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint32Slice is a telemetry Collection whose Await method returns a slice of []uint32 samples.
type CollectionUint32Slice struct {
	W    *Uint32SliceWatcher
	Data []*QualifiedUint32Slice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint32Slice) Await(t testing.TB) []*QualifiedUint32Slice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint32SliceWatcher observes a stream of []uint32 samples.
type Uint32SliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint32Slice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint32SliceWatcher) Await(t testing.TB) (*QualifiedUint32Slice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedUint64Slice is a []uint64 with a corresponding timestamp.
type QualifiedUint64Slice struct {
	*genutil.Metadata
	val     []uint64 // val is the sample value.
	present bool
}

func (q *QualifiedUint64Slice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []uint64 sample, erroring out if not present.
func (q *QualifiedUint64Slice) Val(t testing.TB) []uint64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []uint64 sample.
func (q *QualifiedUint64Slice) SetVal(v []uint64) *QualifiedUint64Slice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint64Slice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint64Slice is a telemetry Collection whose Await method returns a slice of []uint64 samples.
type CollectionUint64Slice struct {
	W    *Uint64SliceWatcher
	Data []*QualifiedUint64Slice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint64Slice) Await(t testing.TB) []*QualifiedUint64Slice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint64SliceWatcher observes a stream of []uint64 samples.
type Uint64SliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint64Slice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint64SliceWatcher) Await(t testing.TB) (*QualifiedUint64Slice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedUint8Slice is a []uint8 with a corresponding timestamp.
type QualifiedUint8Slice struct {
	*genutil.Metadata
	val     []uint8 // val is the sample value.
	present bool
}

func (q *QualifiedUint8Slice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []uint8 sample, erroring out if not present.
func (q *QualifiedUint8Slice) Val(t testing.TB) []uint8 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []uint8 sample.
func (q *QualifiedUint8Slice) SetVal(v []uint8) *QualifiedUint8Slice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint8Slice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint8Slice is a telemetry Collection whose Await method returns a slice of []uint8 samples.
type CollectionUint8Slice struct {
	W    *Uint8SliceWatcher
	Data []*QualifiedUint8Slice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint8Slice) Await(t testing.TB) []*QualifiedUint8Slice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint8SliceWatcher observes a stream of []uint8 samples.
type Uint8SliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint8Slice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint8SliceWatcher) Await(t testing.TB) (*QualifiedUint8Slice, bool) {
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

// QualifiedFloat64 is a float64 with a corresponding timestamp.
type QualifiedFloat64 struct {
	*genutil.Metadata
	val     float64 // val is the sample value.
	present bool
}

func (q *QualifiedFloat64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the float64 sample, erroring out if not present.
func (q *QualifiedFloat64) Val(t testing.TB) float64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the float64 sample.
func (q *QualifiedFloat64) SetVal(v float64) *QualifiedFloat64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFloat64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFloat64 is a telemetry Collection whose Await method returns a slice of float64 samples.
type CollectionFloat64 struct {
	W    *Float64Watcher
	Data []*QualifiedFloat64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFloat64) Await(t testing.TB) []*QualifiedFloat64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Float64Watcher observes a stream of float64 samples.
type Float64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFloat64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Float64Watcher) Await(t testing.TB) (*QualifiedFloat64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInt32 is a int32 with a corresponding timestamp.
type QualifiedInt32 struct {
	*genutil.Metadata
	val     int32 // val is the sample value.
	present bool
}

func (q *QualifiedInt32) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the int32 sample, erroring out if not present.
func (q *QualifiedInt32) Val(t testing.TB) int32 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the int32 sample.
func (q *QualifiedInt32) SetVal(v int32) *QualifiedInt32 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInt32) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInt32 is a telemetry Collection whose Await method returns a slice of int32 samples.
type CollectionInt32 struct {
	W    *Int32Watcher
	Data []*QualifiedInt32
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInt32) Await(t testing.TB) []*QualifiedInt32 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Int32Watcher observes a stream of int32 samples.
type Int32Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInt32
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Int32Watcher) Await(t testing.TB) (*QualifiedInt32, bool) {
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

// QualifiedUint16 is a uint16 with a corresponding timestamp.
type QualifiedUint16 struct {
	*genutil.Metadata
	val     uint16 // val is the sample value.
	present bool
}

func (q *QualifiedUint16) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the uint16 sample, erroring out if not present.
func (q *QualifiedUint16) Val(t testing.TB) uint16 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the uint16 sample.
func (q *QualifiedUint16) SetVal(v uint16) *QualifiedUint16 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint16) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint16 is a telemetry Collection whose Await method returns a slice of uint16 samples.
type CollectionUint16 struct {
	W    *Uint16Watcher
	Data []*QualifiedUint16
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint16) Await(t testing.TB) []*QualifiedUint16 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint16Watcher observes a stream of uint16 samples.
type Uint16Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint16
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint16Watcher) Await(t testing.TB) (*QualifiedUint16, bool) {
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

// QualifiedUint8 is a uint8 with a corresponding timestamp.
type QualifiedUint8 struct {
	*genutil.Metadata
	val     uint8 // val is the sample value.
	present bool
}

func (q *QualifiedUint8) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the uint8 sample, erroring out if not present.
func (q *QualifiedUint8) Val(t testing.TB) uint8 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the uint8 sample.
func (q *QualifiedUint8) SetVal(v uint8) *QualifiedUint8 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedUint8) IsPresent() bool {
	return q != nil && q.present
}

// CollectionUint8 is a telemetry Collection whose Await method returns a slice of uint8 samples.
type CollectionUint8 struct {
	W    *Uint8Watcher
	Data []*QualifiedUint8
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionUint8) Await(t testing.TB) []*QualifiedUint8 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Uint8Watcher observes a stream of uint8 samples.
type Uint8Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedUint8
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Uint8Watcher) Await(t testing.TB) (*QualifiedUint8, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
