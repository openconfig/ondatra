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

// QualifiedE_Ipv6Reachability_Prefix_RedistributionType is a E_Ipv6Reachability_Prefix_RedistributionType with a corresponding timestamp.
type QualifiedE_Ipv6Reachability_Prefix_RedistributionType struct {
	*genutil.Metadata
	val     E_Ipv6Reachability_Prefix_RedistributionType // val is the sample value.
	present bool
}

func (q *QualifiedE_Ipv6Reachability_Prefix_RedistributionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Ipv6Reachability_Prefix_RedistributionType sample, erroring out if not present.
func (q *QualifiedE_Ipv6Reachability_Prefix_RedistributionType) Val(t testing.TB) E_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Ipv6Reachability_Prefix_RedistributionType sample.
func (q *QualifiedE_Ipv6Reachability_Prefix_RedistributionType) SetVal(v E_Ipv6Reachability_Prefix_RedistributionType) *QualifiedE_Ipv6Reachability_Prefix_RedistributionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Ipv6Reachability_Prefix_RedistributionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Ipv6Reachability_Prefix_RedistributionType is a telemetry Collection whose Await method returns a slice of E_Ipv6Reachability_Prefix_RedistributionType samples.
type CollectionE_Ipv6Reachability_Prefix_RedistributionType struct {
	W    *E_Ipv6Reachability_Prefix_RedistributionTypeWatcher
	Data []*QualifiedE_Ipv6Reachability_Prefix_RedistributionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Ipv6Reachability_Prefix_RedistributionType) Await(t testing.TB) []*QualifiedE_Ipv6Reachability_Prefix_RedistributionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Ipv6Reachability_Prefix_RedistributionTypeWatcher observes a stream of E_Ipv6Reachability_Prefix_RedistributionType samples.
type E_Ipv6Reachability_Prefix_RedistributionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Ipv6Reachability_Prefix_RedistributionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Ipv6Reachability_Prefix_RedistributionTypeWatcher) Await(t testing.TB) (*QualifiedE_Ipv6Reachability_Prefix_RedistributionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lag_OperStatus is a E_Lag_OperStatus with a corresponding timestamp.
type QualifiedE_Lag_OperStatus struct {
	*genutil.Metadata
	val     E_Lag_OperStatus // val is the sample value.
	present bool
}

func (q *QualifiedE_Lag_OperStatus) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lag_OperStatus sample, erroring out if not present.
func (q *QualifiedE_Lag_OperStatus) Val(t testing.TB) E_Lag_OperStatus {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Lag_OperStatus sample.
func (q *QualifiedE_Lag_OperStatus) SetVal(v E_Lag_OperStatus) *QualifiedE_Lag_OperStatus {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lag_OperStatus) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lag_OperStatus is a telemetry Collection whose Await method returns a slice of E_Lag_OperStatus samples.
type CollectionE_Lag_OperStatus struct {
	W    *E_Lag_OperStatusWatcher
	Data []*QualifiedE_Lag_OperStatus
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lag_OperStatus) Await(t testing.TB) []*QualifiedE_Lag_OperStatus {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lag_OperStatusWatcher observes a stream of E_Lag_OperStatus samples.
type E_Lag_OperStatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lag_OperStatus
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lag_OperStatusWatcher) Await(t testing.TB) (*QualifiedE_Lag_OperStatus, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Lsps_PduType is a E_Lsps_PduType with a corresponding timestamp.
type QualifiedE_Lsps_PduType struct {
	*genutil.Metadata
	val     E_Lsps_PduType // val is the sample value.
	present bool
}

func (q *QualifiedE_Lsps_PduType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Lsps_PduType sample, erroring out if not present.
func (q *QualifiedE_Lsps_PduType) Val(t testing.TB) E_Lsps_PduType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Lsps_PduType sample.
func (q *QualifiedE_Lsps_PduType) SetVal(v E_Lsps_PduType) *QualifiedE_Lsps_PduType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lsps_PduType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lsps_PduType is a telemetry Collection whose Await method returns a slice of E_Lsps_PduType samples.
type CollectionE_Lsps_PduType struct {
	W    *E_Lsps_PduTypeWatcher
	Data []*QualifiedE_Lsps_PduType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lsps_PduType) Await(t testing.TB) []*QualifiedE_Lsps_PduType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lsps_PduTypeWatcher observes a stream of E_Lsps_PduType samples.
type E_Lsps_PduTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lsps_PduType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lsps_PduTypeWatcher) Await(t testing.TB) (*QualifiedE_Lsps_PduType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType is a E_OpenTrafficGeneratorLacp_LacpActivityType with a corresponding timestamp.
type QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType struct {
	*genutil.Metadata
	val     E_OpenTrafficGeneratorLacp_LacpActivityType // val is the sample value.
	present bool
}

func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OpenTrafficGeneratorLacp_LacpActivityType sample, erroring out if not present.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) Val(t testing.TB) E_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OpenTrafficGeneratorLacp_LacpActivityType sample.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) SetVal(v E_OpenTrafficGeneratorLacp_LacpActivityType) *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OpenTrafficGeneratorLacp_LacpActivityType is a telemetry Collection whose Await method returns a slice of E_OpenTrafficGeneratorLacp_LacpActivityType samples.
type CollectionE_OpenTrafficGeneratorLacp_LacpActivityType struct {
	W    *E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher
	Data []*QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OpenTrafficGeneratorLacp_LacpActivityType) Await(t testing.TB) []*QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher observes a stream of E_OpenTrafficGeneratorLacp_LacpActivityType samples.
type E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OpenTrafficGeneratorLacp_LacpActivityTypeWatcher) Await(t testing.TB) (*QualifiedE_OpenTrafficGeneratorLacp_LacpActivityType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType is a E_OpenTrafficGeneratorLacp_LacpSynchronizationType with a corresponding timestamp.
type QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType struct {
	*genutil.Metadata
	val     E_OpenTrafficGeneratorLacp_LacpSynchronizationType // val is the sample value.
	present bool
}

func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OpenTrafficGeneratorLacp_LacpSynchronizationType sample, erroring out if not present.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) Val(t testing.TB) E_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OpenTrafficGeneratorLacp_LacpSynchronizationType sample.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) SetVal(v E_OpenTrafficGeneratorLacp_LacpSynchronizationType) *QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OpenTrafficGeneratorLacp_LacpSynchronizationType is a telemetry Collection whose Await method returns a slice of E_OpenTrafficGeneratorLacp_LacpSynchronizationType samples.
type CollectionE_OpenTrafficGeneratorLacp_LacpSynchronizationType struct {
	W    *E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher
	Data []*QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OpenTrafficGeneratorLacp_LacpSynchronizationType) Await(t testing.TB) []*QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher observes a stream of E_OpenTrafficGeneratorLacp_LacpSynchronizationType samples.
type E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OpenTrafficGeneratorLacp_LacpSynchronizationTypeWatcher) Await(t testing.TB) (*QualifiedE_OpenTrafficGeneratorLacp_LacpSynchronizationType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType is a E_OpenTrafficGeneratorLacp_LacpTimeoutType with a corresponding timestamp.
type QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType struct {
	*genutil.Metadata
	val     E_OpenTrafficGeneratorLacp_LacpTimeoutType // val is the sample value.
	present bool
}

func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_OpenTrafficGeneratorLacp_LacpTimeoutType sample, erroring out if not present.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) Val(t testing.TB) E_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_OpenTrafficGeneratorLacp_LacpTimeoutType sample.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) SetVal(v E_OpenTrafficGeneratorLacp_LacpTimeoutType) *QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_OpenTrafficGeneratorLacp_LacpTimeoutType is a telemetry Collection whose Await method returns a slice of E_OpenTrafficGeneratorLacp_LacpTimeoutType samples.
type CollectionE_OpenTrafficGeneratorLacp_LacpTimeoutType struct {
	W    *E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher
	Data []*QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_OpenTrafficGeneratorLacp_LacpTimeoutType) Await(t testing.TB) []*QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher observes a stream of E_OpenTrafficGeneratorLacp_LacpTimeoutType samples.
type E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_OpenTrafficGeneratorLacp_LacpTimeoutTypeWatcher) Await(t testing.TB) (*QualifiedE_OpenTrafficGeneratorLacp_LacpTimeoutType, bool) {
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

// QualifiedE_State_CommunityType is a E_State_CommunityType with a corresponding timestamp.
type QualifiedE_State_CommunityType struct {
	*genutil.Metadata
	val     E_State_CommunityType // val is the sample value.
	present bool
}

func (q *QualifiedE_State_CommunityType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_State_CommunityType sample, erroring out if not present.
func (q *QualifiedE_State_CommunityType) Val(t testing.TB) E_State_CommunityType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_State_CommunityType sample.
func (q *QualifiedE_State_CommunityType) SetVal(v E_State_CommunityType) *QualifiedE_State_CommunityType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_State_CommunityType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_State_CommunityType is a telemetry Collection whose Await method returns a slice of E_State_CommunityType samples.
type CollectionE_State_CommunityType struct {
	W    *E_State_CommunityTypeWatcher
	Data []*QualifiedE_State_CommunityType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_State_CommunityType) Await(t testing.TB) []*QualifiedE_State_CommunityType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_State_CommunityTypeWatcher observes a stream of E_State_CommunityType samples.
type E_State_CommunityTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_State_CommunityType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_State_CommunityTypeWatcher) Await(t testing.TB) (*QualifiedE_State_CommunityType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_State_SegmentType is a E_State_SegmentType with a corresponding timestamp.
type QualifiedE_State_SegmentType struct {
	*genutil.Metadata
	val     E_State_SegmentType // val is the sample value.
	present bool
}

func (q *QualifiedE_State_SegmentType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_State_SegmentType sample, erroring out if not present.
func (q *QualifiedE_State_SegmentType) Val(t testing.TB) E_State_SegmentType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_State_SegmentType sample.
func (q *QualifiedE_State_SegmentType) SetVal(v E_State_SegmentType) *QualifiedE_State_SegmentType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_State_SegmentType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_State_SegmentType is a telemetry Collection whose Await method returns a slice of E_State_SegmentType samples.
type CollectionE_State_SegmentType struct {
	W    *E_State_SegmentTypeWatcher
	Data []*QualifiedE_State_SegmentType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_State_SegmentType) Await(t testing.TB) []*QualifiedE_State_SegmentType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_State_SegmentTypeWatcher observes a stream of E_State_SegmentType samples.
type E_State_SegmentTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_State_SegmentType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_State_SegmentTypeWatcher) Await(t testing.TB) (*QualifiedE_State_SegmentType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

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

// QualifiedE_Lsps_FlagsSlice is a []E_Lsps_Flags with a corresponding timestamp.
type QualifiedE_Lsps_FlagsSlice struct {
	*genutil.Metadata
	val     []E_Lsps_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_Lsps_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_Lsps_Flags sample, erroring out if not present.
func (q *QualifiedE_Lsps_FlagsSlice) Val(t testing.TB) []E_Lsps_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_Lsps_Flags sample.
func (q *QualifiedE_Lsps_FlagsSlice) SetVal(v []E_Lsps_Flags) *QualifiedE_Lsps_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Lsps_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Lsps_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_Lsps_Flags samples.
type CollectionE_Lsps_FlagsSlice struct {
	W    *E_Lsps_FlagsSliceWatcher
	Data []*QualifiedE_Lsps_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Lsps_FlagsSlice) Await(t testing.TB) []*QualifiedE_Lsps_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Lsps_FlagsSliceWatcher observes a stream of []E_Lsps_Flags samples.
type E_Lsps_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Lsps_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Lsps_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_Lsps_FlagsSlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_State_FlagsSlice is a []E_State_Flags with a corresponding timestamp.
type QualifiedE_State_FlagsSlice struct {
	*genutil.Metadata
	val     []E_State_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_State_FlagsSlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []E_State_Flags sample, erroring out if not present.
func (q *QualifiedE_State_FlagsSlice) Val(t testing.TB) []E_State_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []E_State_Flags sample.
func (q *QualifiedE_State_FlagsSlice) SetVal(v []E_State_Flags) *QualifiedE_State_FlagsSlice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_State_FlagsSlice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_State_FlagsSlice is a telemetry Collection whose Await method returns a slice of []E_State_Flags samples.
type CollectionE_State_FlagsSlice struct {
	W    *E_State_FlagsSliceWatcher
	Data []*QualifiedE_State_FlagsSlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_State_FlagsSlice) Await(t testing.TB) []*QualifiedE_State_FlagsSlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_State_FlagsSliceWatcher observes a stream of []E_State_Flags samples.
type E_State_FlagsSliceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_State_FlagsSlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_State_FlagsSliceWatcher) Await(t testing.TB) (*QualifiedE_State_FlagsSlice, bool) {
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
