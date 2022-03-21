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

// QualifiedQos_QueueManagementProfile_Red_Uniform is a *Qos_QueueManagementProfile_Red_Uniform with a corresponding timestamp.
type QualifiedQos_QueueManagementProfile_Red_Uniform struct {
	*genutil.Metadata
	val     *Qos_QueueManagementProfile_Red_Uniform // val is the sample value.
	present bool
}

func (q *QualifiedQos_QueueManagementProfile_Red_Uniform) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_QueueManagementProfile_Red_Uniform sample, erroring out if not present.
func (q *QualifiedQos_QueueManagementProfile_Red_Uniform) Val(t testing.TB) *Qos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_QueueManagementProfile_Red_Uniform sample.
func (q *QualifiedQos_QueueManagementProfile_Red_Uniform) SetVal(v *Qos_QueueManagementProfile_Red_Uniform) *QualifiedQos_QueueManagementProfile_Red_Uniform {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_QueueManagementProfile_Red_Uniform) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_QueueManagementProfile_Red_Uniform is a telemetry Collection whose Await method returns a slice of *Qos_QueueManagementProfile_Red_Uniform samples.
type CollectionQos_QueueManagementProfile_Red_Uniform struct {
	W    *Qos_QueueManagementProfile_Red_UniformWatcher
	Data []*QualifiedQos_QueueManagementProfile_Red_Uniform
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_QueueManagementProfile_Red_Uniform) Await(t testing.TB) []*QualifiedQos_QueueManagementProfile_Red_Uniform {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_QueueManagementProfile_Red_UniformWatcher observes a stream of *Qos_QueueManagementProfile_Red_Uniform samples.
type Qos_QueueManagementProfile_Red_UniformWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_QueueManagementProfile_Red_Uniform
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_QueueManagementProfile_Red_UniformWatcher) Await(t testing.TB) (*QualifiedQos_QueueManagementProfile_Red_Uniform, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_QueueManagementProfile_Wred is a *Qos_QueueManagementProfile_Wred with a corresponding timestamp.
type QualifiedQos_QueueManagementProfile_Wred struct {
	*genutil.Metadata
	val     *Qos_QueueManagementProfile_Wred // val is the sample value.
	present bool
}

func (q *QualifiedQos_QueueManagementProfile_Wred) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_QueueManagementProfile_Wred sample, erroring out if not present.
func (q *QualifiedQos_QueueManagementProfile_Wred) Val(t testing.TB) *Qos_QueueManagementProfile_Wred {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_QueueManagementProfile_Wred sample.
func (q *QualifiedQos_QueueManagementProfile_Wred) SetVal(v *Qos_QueueManagementProfile_Wred) *QualifiedQos_QueueManagementProfile_Wred {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_QueueManagementProfile_Wred) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_QueueManagementProfile_Wred is a telemetry Collection whose Await method returns a slice of *Qos_QueueManagementProfile_Wred samples.
type CollectionQos_QueueManagementProfile_Wred struct {
	W    *Qos_QueueManagementProfile_WredWatcher
	Data []*QualifiedQos_QueueManagementProfile_Wred
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_QueueManagementProfile_Wred) Await(t testing.TB) []*QualifiedQos_QueueManagementProfile_Wred {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_QueueManagementProfile_WredWatcher observes a stream of *Qos_QueueManagementProfile_Wred samples.
type Qos_QueueManagementProfile_WredWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_QueueManagementProfile_Wred
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_QueueManagementProfile_WredWatcher) Await(t testing.TB) (*QualifiedQos_QueueManagementProfile_Wred, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_QueueManagementProfile_Wred_Uniform is a *Qos_QueueManagementProfile_Wred_Uniform with a corresponding timestamp.
type QualifiedQos_QueueManagementProfile_Wred_Uniform struct {
	*genutil.Metadata
	val     *Qos_QueueManagementProfile_Wred_Uniform // val is the sample value.
	present bool
}

func (q *QualifiedQos_QueueManagementProfile_Wred_Uniform) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_QueueManagementProfile_Wred_Uniform sample, erroring out if not present.
func (q *QualifiedQos_QueueManagementProfile_Wred_Uniform) Val(t testing.TB) *Qos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_QueueManagementProfile_Wred_Uniform sample.
func (q *QualifiedQos_QueueManagementProfile_Wred_Uniform) SetVal(v *Qos_QueueManagementProfile_Wred_Uniform) *QualifiedQos_QueueManagementProfile_Wred_Uniform {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_QueueManagementProfile_Wred_Uniform) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_QueueManagementProfile_Wred_Uniform is a telemetry Collection whose Await method returns a slice of *Qos_QueueManagementProfile_Wred_Uniform samples.
type CollectionQos_QueueManagementProfile_Wred_Uniform struct {
	W    *Qos_QueueManagementProfile_Wred_UniformWatcher
	Data []*QualifiedQos_QueueManagementProfile_Wred_Uniform
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_QueueManagementProfile_Wred_Uniform) Await(t testing.TB) []*QualifiedQos_QueueManagementProfile_Wred_Uniform {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_QueueManagementProfile_Wred_UniformWatcher observes a stream of *Qos_QueueManagementProfile_Wred_Uniform samples.
type Qos_QueueManagementProfile_Wred_UniformWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_QueueManagementProfile_Wred_Uniform
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_QueueManagementProfile_Wred_UniformWatcher) Await(t testing.TB) (*QualifiedQos_QueueManagementProfile_Wred_Uniform, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy is a *Qos_SchedulerPolicy with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy) Val(t testing.TB) *Qos_SchedulerPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy sample.
func (q *QualifiedQos_SchedulerPolicy) SetVal(v *Qos_SchedulerPolicy) *QualifiedQos_SchedulerPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy samples.
type CollectionQos_SchedulerPolicy struct {
	W    *Qos_SchedulerPolicyWatcher
	Data []*QualifiedQos_SchedulerPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicyWatcher observes a stream of *Qos_SchedulerPolicy samples.
type Qos_SchedulerPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicyWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler is a *Qos_SchedulerPolicy_Scheduler with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler) SetVal(v *Qos_SchedulerPolicy_Scheduler) *QualifiedQos_SchedulerPolicy_Scheduler {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler samples.
type CollectionQos_SchedulerPolicy_Scheduler struct {
	W    *Qos_SchedulerPolicy_SchedulerWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_SchedulerWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler samples.
type Qos_SchedulerPolicy_SchedulerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_SchedulerWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_Input is a *Qos_SchedulerPolicy_Scheduler_Input with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_Input struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_Input // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_Input) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_Input sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_Input) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_Input sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_Input) SetVal(v *Qos_SchedulerPolicy_Scheduler_Input) *QualifiedQos_SchedulerPolicy_Scheduler_Input {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_Input) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_Input is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_Input samples.
type CollectionQos_SchedulerPolicy_Scheduler_Input struct {
	W    *Qos_SchedulerPolicy_Scheduler_InputWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_Input
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_Input) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_Input {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_InputWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_Input samples.
type Qos_SchedulerPolicy_Scheduler_InputWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_Input
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_InputWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_Input, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor is a *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) SetVal(v *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor) *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor samples.
type CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor struct {
	W    *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor samples.
type Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_OneRateTwoColorWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction is a *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) SetVal(v *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction samples.
type CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction struct {
	W    *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction samples.
type Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformActionWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ConformAction, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction is a *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) SetVal(v *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction samples.
type CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction struct {
	W    *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction samples.
type Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedActionWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_OneRateTwoColor_ExceedAction, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_Output is a *Qos_SchedulerPolicy_Scheduler_Output with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_Output struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_Output // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_Output) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_Output sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_Output) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_Output sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_Output) SetVal(v *Qos_SchedulerPolicy_Scheduler_Output) *QualifiedQos_SchedulerPolicy_Scheduler_Output {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_Output) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_Output is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_Output samples.
type CollectionQos_SchedulerPolicy_Scheduler_Output struct {
	W    *Qos_SchedulerPolicy_Scheduler_OutputWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_Output
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_Output) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_Output {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_OutputWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_Output samples.
type Qos_SchedulerPolicy_Scheduler_OutputWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_Output
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_OutputWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_Output, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor is a *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) SetVal(v *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor) *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor samples.
type CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor struct {
	W    *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor samples.
type Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColorWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction is a *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) SetVal(v *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction samples.
type CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction struct {
	W    *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction samples.
type Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformActionWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ConformAction, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction is a *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) SetVal(v *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction samples.
type CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction struct {
	W    *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction samples.
type Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedActionWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ExceedAction, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction is a *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction with a corresponding timestamp.
type QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction struct {
	*genutil.Metadata
	val     *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction // val is the sample value.
	present bool
}

func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction sample, erroring out if not present.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) Val(t testing.TB) *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction sample.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) SetVal(v *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) IsPresent() bool {
	return q != nil && q.present
}

// CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction is a telemetry Collection whose Await method returns a slice of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction samples.
type CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction struct {
	W    *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionWatcher
	Data []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction) Await(t testing.TB) []*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionWatcher observes a stream of *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction samples.
type Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Qos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateActionWatcher) Await(t testing.TB) (*QualifiedQos_SchedulerPolicy_Scheduler_TwoRateThreeColor_ViolateAction, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy is a *RoutingPolicy with a corresponding timestamp.
type QualifiedRoutingPolicy struct {
	*genutil.Metadata
	val     *RoutingPolicy // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy sample, erroring out if not present.
func (q *QualifiedRoutingPolicy) Val(t testing.TB) *RoutingPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy sample.
func (q *QualifiedRoutingPolicy) SetVal(v *RoutingPolicy) *QualifiedRoutingPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy is a telemetry Collection whose Await method returns a slice of *RoutingPolicy samples.
type CollectionRoutingPolicy struct {
	W    *RoutingPolicyWatcher
	Data []*QualifiedRoutingPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy) Await(t testing.TB) []*QualifiedRoutingPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicyWatcher observes a stream of *RoutingPolicy samples.
type RoutingPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicyWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets is a *RoutingPolicy_DefinedSets with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets) Val(t testing.TB) *RoutingPolicy_DefinedSets {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets sample.
func (q *QualifiedRoutingPolicy_DefinedSets) SetVal(v *RoutingPolicy_DefinedSets) *QualifiedRoutingPolicy_DefinedSets {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets samples.
type CollectionRoutingPolicy_DefinedSets struct {
	W    *RoutingPolicy_DefinedSetsWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSetsWatcher observes a stream of *RoutingPolicy_DefinedSets samples.
type RoutingPolicy_DefinedSetsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSetsWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets is a *RoutingPolicy_DefinedSets_BgpDefinedSets with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_BgpDefinedSets // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) Val(t testing.TB) *RoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets sample.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) SetVal(v *RoutingPolicy_DefinedSets_BgpDefinedSets) *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_BgpDefinedSets is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_BgpDefinedSets samples.
type CollectionRoutingPolicy_DefinedSets_BgpDefinedSets struct {
	W    *RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_BgpDefinedSets) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher observes a stream of *RoutingPolicy_DefinedSets_BgpDefinedSets samples.
type RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_BgpDefinedSetsWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet is a *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) Val(t testing.TB) *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet sample.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) SetVal(v *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet samples.
type CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet struct {
	W    *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher observes a stream of *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet samples.
type RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet is a *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) Val(t testing.TB) *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet sample.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) SetVal(v *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet samples.
type CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet struct {
	W    *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher observes a stream of *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet samples.
type RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet is a *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) Val(t testing.TB) *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet sample.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) SetVal(v *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet samples.
type CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet struct {
	W    *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher observes a stream of *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet samples.
type RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_NeighborSet is a *RoutingPolicy_DefinedSets_NeighborSet with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_NeighborSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_NeighborSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_NeighborSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_NeighborSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_NeighborSet) Val(t testing.TB) *RoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_NeighborSet sample.
func (q *QualifiedRoutingPolicy_DefinedSets_NeighborSet) SetVal(v *RoutingPolicy_DefinedSets_NeighborSet) *QualifiedRoutingPolicy_DefinedSets_NeighborSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_NeighborSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_NeighborSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_NeighborSet samples.
type CollectionRoutingPolicy_DefinedSets_NeighborSet struct {
	W    *RoutingPolicy_DefinedSets_NeighborSetWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_NeighborSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_NeighborSet) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_NeighborSetWatcher observes a stream of *RoutingPolicy_DefinedSets_NeighborSet samples.
type RoutingPolicy_DefinedSets_NeighborSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_NeighborSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_NeighborSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_NeighborSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_PrefixSet is a *RoutingPolicy_DefinedSets_PrefixSet with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_PrefixSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_PrefixSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_PrefixSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet) Val(t testing.TB) *RoutingPolicy_DefinedSets_PrefixSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_PrefixSet sample.
func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet) SetVal(v *RoutingPolicy_DefinedSets_PrefixSet) *QualifiedRoutingPolicy_DefinedSets_PrefixSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_PrefixSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_PrefixSet samples.
type CollectionRoutingPolicy_DefinedSets_PrefixSet struct {
	W    *RoutingPolicy_DefinedSets_PrefixSetWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_PrefixSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_PrefixSet) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_PrefixSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_PrefixSetWatcher observes a stream of *RoutingPolicy_DefinedSets_PrefixSet samples.
type RoutingPolicy_DefinedSets_PrefixSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_PrefixSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_PrefixSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_PrefixSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix is a *RoutingPolicy_DefinedSets_PrefixSet_Prefix with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_PrefixSet_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_PrefixSet_Prefix sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) Val(t testing.TB) *RoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_PrefixSet_Prefix sample.
func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) SetVal(v *RoutingPolicy_DefinedSets_PrefixSet_Prefix) *QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_PrefixSet_Prefix is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_PrefixSet_Prefix samples.
type CollectionRoutingPolicy_DefinedSets_PrefixSet_Prefix struct {
	W    *RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_PrefixSet_Prefix) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher observes a stream of *RoutingPolicy_DefinedSets_PrefixSet_Prefix samples.
type RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_PrefixSet_PrefixWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_DefinedSets_TagSet is a *RoutingPolicy_DefinedSets_TagSet with a corresponding timestamp.
type QualifiedRoutingPolicy_DefinedSets_TagSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_DefinedSets_TagSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_DefinedSets_TagSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_DefinedSets_TagSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_DefinedSets_TagSet) Val(t testing.TB) *RoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_DefinedSets_TagSet sample.
func (q *QualifiedRoutingPolicy_DefinedSets_TagSet) SetVal(v *RoutingPolicy_DefinedSets_TagSet) *QualifiedRoutingPolicy_DefinedSets_TagSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_DefinedSets_TagSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_DefinedSets_TagSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_DefinedSets_TagSet samples.
type CollectionRoutingPolicy_DefinedSets_TagSet struct {
	W    *RoutingPolicy_DefinedSets_TagSetWatcher
	Data []*QualifiedRoutingPolicy_DefinedSets_TagSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_DefinedSets_TagSet) Await(t testing.TB) []*QualifiedRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_DefinedSets_TagSetWatcher observes a stream of *RoutingPolicy_DefinedSets_TagSet samples.
type RoutingPolicy_DefinedSets_TagSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_DefinedSets_TagSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_DefinedSets_TagSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_DefinedSets_TagSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition is a *RoutingPolicy_PolicyDefinition with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition) Val(t testing.TB) *RoutingPolicy_PolicyDefinition {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition) SetVal(v *RoutingPolicy_PolicyDefinition) *QualifiedRoutingPolicy_PolicyDefinition {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition samples.
type CollectionRoutingPolicy_PolicyDefinition struct {
	W    *RoutingPolicy_PolicyDefinitionWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinitionWatcher observes a stream of *RoutingPolicy_PolicyDefinition samples.
type RoutingPolicy_PolicyDefinitionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinitionWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement is a *RoutingPolicy_PolicyDefinition_Statement with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement) SetVal(v *RoutingPolicy_PolicyDefinition_Statement) *QualifiedRoutingPolicy_PolicyDefinition_Statement {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement struct {
	W    *RoutingPolicy_PolicyDefinition_StatementWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_StatementWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement samples.
type RoutingPolicy_PolicyDefinition_StatementWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_StatementWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions is a *RoutingPolicy_PolicyDefinition_Statement_Actions with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions samples.
type RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_ActionsWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlineWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferenceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlineWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference is a *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferenceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag is a *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline is a *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlineWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference is a *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference samples.
type RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ReferenceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Reference, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions is a *RoutingPolicy_PolicyDefinition_Statement_Conditions with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions samples.
type RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_ConditionsWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditionsWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLengthWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_AsPathLength, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCountWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_CommunityCount, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSetWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSetWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_MatchAsPathSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfaceWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet is a *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet with a corresponding timestamp.
type QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet struct {
	*genutil.Metadata
	val     *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet // val is the sample value.
	present bool
}

func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet sample, erroring out if not present.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) Val(t testing.TB) *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet sample.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) SetVal(v *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet is a telemetry Collection whose Await method returns a slice of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet samples.
type CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet struct {
	W    *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher
	Data []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) Await(t testing.TB) []*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher observes a stream of *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet samples.
type RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetWatcher) Await(t testing.TB) (*QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem is a *System with a corresponding timestamp.
type QualifiedSystem struct {
	*genutil.Metadata
	val     *System // val is the sample value.
	present bool
}

func (q *QualifiedSystem) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System sample, erroring out if not present.
func (q *QualifiedSystem) Val(t testing.TB) *System {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System sample.
func (q *QualifiedSystem) SetVal(v *System) *QualifiedSystem {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem is a telemetry Collection whose Await method returns a slice of *System samples.
type CollectionSystem struct {
	W    *SystemWatcher
	Data []*QualifiedSystem
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem) Await(t testing.TB) []*QualifiedSystem {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// SystemWatcher observes a stream of *System samples.
type SystemWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *SystemWatcher) Await(t testing.TB) (*QualifiedSystem, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa is a *System_Aaa with a corresponding timestamp.
type QualifiedSystem_Aaa struct {
	*genutil.Metadata
	val     *System_Aaa // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa sample, erroring out if not present.
func (q *QualifiedSystem_Aaa) Val(t testing.TB) *System_Aaa {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa sample.
func (q *QualifiedSystem_Aaa) SetVal(v *System_Aaa) *QualifiedSystem_Aaa {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa is a telemetry Collection whose Await method returns a slice of *System_Aaa samples.
type CollectionSystem_Aaa struct {
	W    *System_AaaWatcher
	Data []*QualifiedSystem_Aaa
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa) Await(t testing.TB) []*QualifiedSystem_Aaa {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_AaaWatcher observes a stream of *System_Aaa samples.
type System_AaaWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_AaaWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Accounting is a *System_Aaa_Accounting with a corresponding timestamp.
type QualifiedSystem_Aaa_Accounting struct {
	*genutil.Metadata
	val     *System_Aaa_Accounting // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Accounting) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_Accounting sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Accounting) Val(t testing.TB) *System_Aaa_Accounting {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_Accounting sample.
func (q *QualifiedSystem_Aaa_Accounting) SetVal(v *System_Aaa_Accounting) *QualifiedSystem_Aaa_Accounting {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Accounting) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Accounting is a telemetry Collection whose Await method returns a slice of *System_Aaa_Accounting samples.
type CollectionSystem_Aaa_Accounting struct {
	W    *System_Aaa_AccountingWatcher
	Data []*QualifiedSystem_Aaa_Accounting
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Accounting) Await(t testing.TB) []*QualifiedSystem_Aaa_Accounting {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_AccountingWatcher observes a stream of *System_Aaa_Accounting samples.
type System_Aaa_AccountingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Accounting
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_AccountingWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Accounting, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Accounting_Event is a *System_Aaa_Accounting_Event with a corresponding timestamp.
type QualifiedSystem_Aaa_Accounting_Event struct {
	*genutil.Metadata
	val     *System_Aaa_Accounting_Event // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Accounting_Event) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_Accounting_Event sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Accounting_Event) Val(t testing.TB) *System_Aaa_Accounting_Event {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_Accounting_Event sample.
func (q *QualifiedSystem_Aaa_Accounting_Event) SetVal(v *System_Aaa_Accounting_Event) *QualifiedSystem_Aaa_Accounting_Event {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Accounting_Event) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Accounting_Event is a telemetry Collection whose Await method returns a slice of *System_Aaa_Accounting_Event samples.
type CollectionSystem_Aaa_Accounting_Event struct {
	W    *System_Aaa_Accounting_EventWatcher
	Data []*QualifiedSystem_Aaa_Accounting_Event
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Accounting_Event) Await(t testing.TB) []*QualifiedSystem_Aaa_Accounting_Event {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Accounting_EventWatcher observes a stream of *System_Aaa_Accounting_Event samples.
type System_Aaa_Accounting_EventWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Accounting_Event
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Accounting_EventWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Accounting_Event, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authentication is a *System_Aaa_Authentication with a corresponding timestamp.
type QualifiedSystem_Aaa_Authentication struct {
	*genutil.Metadata
	val     *System_Aaa_Authentication // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_Authentication sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authentication) Val(t testing.TB) *System_Aaa_Authentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_Authentication sample.
func (q *QualifiedSystem_Aaa_Authentication) SetVal(v *System_Aaa_Authentication) *QualifiedSystem_Aaa_Authentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authentication is a telemetry Collection whose Await method returns a slice of *System_Aaa_Authentication samples.
type CollectionSystem_Aaa_Authentication struct {
	W    *System_Aaa_AuthenticationWatcher
	Data []*QualifiedSystem_Aaa_Authentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authentication) Await(t testing.TB) []*QualifiedSystem_Aaa_Authentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_AuthenticationWatcher observes a stream of *System_Aaa_Authentication samples.
type System_Aaa_AuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_AuthenticationWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authentication_AdminUser is a *System_Aaa_Authentication_AdminUser with a corresponding timestamp.
type QualifiedSystem_Aaa_Authentication_AdminUser struct {
	*genutil.Metadata
	val     *System_Aaa_Authentication_AdminUser // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authentication_AdminUser) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_Authentication_AdminUser sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authentication_AdminUser) Val(t testing.TB) *System_Aaa_Authentication_AdminUser {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_Authentication_AdminUser sample.
func (q *QualifiedSystem_Aaa_Authentication_AdminUser) SetVal(v *System_Aaa_Authentication_AdminUser) *QualifiedSystem_Aaa_Authentication_AdminUser {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authentication_AdminUser) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authentication_AdminUser is a telemetry Collection whose Await method returns a slice of *System_Aaa_Authentication_AdminUser samples.
type CollectionSystem_Aaa_Authentication_AdminUser struct {
	W    *System_Aaa_Authentication_AdminUserWatcher
	Data []*QualifiedSystem_Aaa_Authentication_AdminUser
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authentication_AdminUser) Await(t testing.TB) []*QualifiedSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Authentication_AdminUserWatcher observes a stream of *System_Aaa_Authentication_AdminUser samples.
type System_Aaa_Authentication_AdminUserWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authentication_AdminUser
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Authentication_AdminUserWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authentication_AdminUser, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authentication_User is a *System_Aaa_Authentication_User with a corresponding timestamp.
type QualifiedSystem_Aaa_Authentication_User struct {
	*genutil.Metadata
	val     *System_Aaa_Authentication_User // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authentication_User) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_Authentication_User sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authentication_User) Val(t testing.TB) *System_Aaa_Authentication_User {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_Authentication_User sample.
func (q *QualifiedSystem_Aaa_Authentication_User) SetVal(v *System_Aaa_Authentication_User) *QualifiedSystem_Aaa_Authentication_User {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authentication_User) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authentication_User is a telemetry Collection whose Await method returns a slice of *System_Aaa_Authentication_User samples.
type CollectionSystem_Aaa_Authentication_User struct {
	W    *System_Aaa_Authentication_UserWatcher
	Data []*QualifiedSystem_Aaa_Authentication_User
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authentication_User) Await(t testing.TB) []*QualifiedSystem_Aaa_Authentication_User {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Authentication_UserWatcher observes a stream of *System_Aaa_Authentication_User samples.
type System_Aaa_Authentication_UserWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authentication_User
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Authentication_UserWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authentication_User, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authorization is a *System_Aaa_Authorization with a corresponding timestamp.
type QualifiedSystem_Aaa_Authorization struct {
	*genutil.Metadata
	val     *System_Aaa_Authorization // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authorization) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_Authorization sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authorization) Val(t testing.TB) *System_Aaa_Authorization {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_Authorization sample.
func (q *QualifiedSystem_Aaa_Authorization) SetVal(v *System_Aaa_Authorization) *QualifiedSystem_Aaa_Authorization {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authorization) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authorization is a telemetry Collection whose Await method returns a slice of *System_Aaa_Authorization samples.
type CollectionSystem_Aaa_Authorization struct {
	W    *System_Aaa_AuthorizationWatcher
	Data []*QualifiedSystem_Aaa_Authorization
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authorization) Await(t testing.TB) []*QualifiedSystem_Aaa_Authorization {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_AuthorizationWatcher observes a stream of *System_Aaa_Authorization samples.
type System_Aaa_AuthorizationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authorization
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_AuthorizationWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authorization, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_Authorization_Event is a *System_Aaa_Authorization_Event with a corresponding timestamp.
type QualifiedSystem_Aaa_Authorization_Event struct {
	*genutil.Metadata
	val     *System_Aaa_Authorization_Event // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_Authorization_Event) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_Authorization_Event sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_Authorization_Event) Val(t testing.TB) *System_Aaa_Authorization_Event {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_Authorization_Event sample.
func (q *QualifiedSystem_Aaa_Authorization_Event) SetVal(v *System_Aaa_Authorization_Event) *QualifiedSystem_Aaa_Authorization_Event {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_Authorization_Event) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_Authorization_Event is a telemetry Collection whose Await method returns a slice of *System_Aaa_Authorization_Event samples.
type CollectionSystem_Aaa_Authorization_Event struct {
	W    *System_Aaa_Authorization_EventWatcher
	Data []*QualifiedSystem_Aaa_Authorization_Event
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_Authorization_Event) Await(t testing.TB) []*QualifiedSystem_Aaa_Authorization_Event {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_Authorization_EventWatcher observes a stream of *System_Aaa_Authorization_Event samples.
type System_Aaa_Authorization_EventWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_Authorization_Event
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_Authorization_EventWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_Authorization_Event, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_ServerGroup is a *System_Aaa_ServerGroup with a corresponding timestamp.
type QualifiedSystem_Aaa_ServerGroup struct {
	*genutil.Metadata
	val     *System_Aaa_ServerGroup // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_ServerGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_ServerGroup sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_ServerGroup) Val(t testing.TB) *System_Aaa_ServerGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_ServerGroup sample.
func (q *QualifiedSystem_Aaa_ServerGroup) SetVal(v *System_Aaa_ServerGroup) *QualifiedSystem_Aaa_ServerGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_ServerGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_ServerGroup is a telemetry Collection whose Await method returns a slice of *System_Aaa_ServerGroup samples.
type CollectionSystem_Aaa_ServerGroup struct {
	W    *System_Aaa_ServerGroupWatcher
	Data []*QualifiedSystem_Aaa_ServerGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_ServerGroup) Await(t testing.TB) []*QualifiedSystem_Aaa_ServerGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_ServerGroupWatcher observes a stream of *System_Aaa_ServerGroup samples.
type System_Aaa_ServerGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_ServerGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_ServerGroupWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_ServerGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_ServerGroup_Server is a *System_Aaa_ServerGroup_Server with a corresponding timestamp.
type QualifiedSystem_Aaa_ServerGroup_Server struct {
	*genutil.Metadata
	val     *System_Aaa_ServerGroup_Server // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_ServerGroup_Server) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_ServerGroup_Server sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_ServerGroup_Server) Val(t testing.TB) *System_Aaa_ServerGroup_Server {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_ServerGroup_Server sample.
func (q *QualifiedSystem_Aaa_ServerGroup_Server) SetVal(v *System_Aaa_ServerGroup_Server) *QualifiedSystem_Aaa_ServerGroup_Server {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_ServerGroup_Server) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_ServerGroup_Server is a telemetry Collection whose Await method returns a slice of *System_Aaa_ServerGroup_Server samples.
type CollectionSystem_Aaa_ServerGroup_Server struct {
	W    *System_Aaa_ServerGroup_ServerWatcher
	Data []*QualifiedSystem_Aaa_ServerGroup_Server
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_ServerGroup_Server) Await(t testing.TB) []*QualifiedSystem_Aaa_ServerGroup_Server {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_ServerGroup_ServerWatcher observes a stream of *System_Aaa_ServerGroup_Server samples.
type System_Aaa_ServerGroup_ServerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_ServerGroup_Server
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_ServerGroup_ServerWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_ServerGroup_Server, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_ServerGroup_Server_Radius is a *System_Aaa_ServerGroup_Server_Radius with a corresponding timestamp.
type QualifiedSystem_Aaa_ServerGroup_Server_Radius struct {
	*genutil.Metadata
	val     *System_Aaa_ServerGroup_Server_Radius // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_ServerGroup_Server_Radius sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius) Val(t testing.TB) *System_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_ServerGroup_Server_Radius sample.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius) SetVal(v *System_Aaa_ServerGroup_Server_Radius) *QualifiedSystem_Aaa_ServerGroup_Server_Radius {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_ServerGroup_Server_Radius is a telemetry Collection whose Await method returns a slice of *System_Aaa_ServerGroup_Server_Radius samples.
type CollectionSystem_Aaa_ServerGroup_Server_Radius struct {
	W    *System_Aaa_ServerGroup_Server_RadiusWatcher
	Data []*QualifiedSystem_Aaa_ServerGroup_Server_Radius
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_ServerGroup_Server_Radius) Await(t testing.TB) []*QualifiedSystem_Aaa_ServerGroup_Server_Radius {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_ServerGroup_Server_RadiusWatcher observes a stream of *System_Aaa_ServerGroup_Server_Radius samples.
type System_Aaa_ServerGroup_Server_RadiusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_ServerGroup_Server_Radius
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_ServerGroup_Server_RadiusWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_ServerGroup_Server_Radius, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters is a *System_Aaa_ServerGroup_Server_Radius_Counters with a corresponding timestamp.
type QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters struct {
	*genutil.Metadata
	val     *System_Aaa_ServerGroup_Server_Radius_Counters // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_ServerGroup_Server_Radius_Counters sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) Val(t testing.TB) *System_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_ServerGroup_Server_Radius_Counters sample.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) SetVal(v *System_Aaa_ServerGroup_Server_Radius_Counters) *QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_ServerGroup_Server_Radius_Counters is a telemetry Collection whose Await method returns a slice of *System_Aaa_ServerGroup_Server_Radius_Counters samples.
type CollectionSystem_Aaa_ServerGroup_Server_Radius_Counters struct {
	W    *System_Aaa_ServerGroup_Server_Radius_CountersWatcher
	Data []*QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_ServerGroup_Server_Radius_Counters) Await(t testing.TB) []*QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_ServerGroup_Server_Radius_CountersWatcher observes a stream of *System_Aaa_ServerGroup_Server_Radius_Counters samples.
type System_Aaa_ServerGroup_Server_Radius_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_ServerGroup_Server_Radius_CountersWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_ServerGroup_Server_Radius_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Aaa_ServerGroup_Server_Tacacs is a *System_Aaa_ServerGroup_Server_Tacacs with a corresponding timestamp.
type QualifiedSystem_Aaa_ServerGroup_Server_Tacacs struct {
	*genutil.Metadata
	val     *System_Aaa_ServerGroup_Server_Tacacs // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Aaa_ServerGroup_Server_Tacacs sample, erroring out if not present.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) Val(t testing.TB) *System_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Aaa_ServerGroup_Server_Tacacs sample.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) SetVal(v *System_Aaa_ServerGroup_Server_Tacacs) *QualifiedSystem_Aaa_ServerGroup_Server_Tacacs {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Aaa_ServerGroup_Server_Tacacs) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Aaa_ServerGroup_Server_Tacacs is a telemetry Collection whose Await method returns a slice of *System_Aaa_ServerGroup_Server_Tacacs samples.
type CollectionSystem_Aaa_ServerGroup_Server_Tacacs struct {
	W    *System_Aaa_ServerGroup_Server_TacacsWatcher
	Data []*QualifiedSystem_Aaa_ServerGroup_Server_Tacacs
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Aaa_ServerGroup_Server_Tacacs) Await(t testing.TB) []*QualifiedSystem_Aaa_ServerGroup_Server_Tacacs {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Aaa_ServerGroup_Server_TacacsWatcher observes a stream of *System_Aaa_ServerGroup_Server_Tacacs samples.
type System_Aaa_ServerGroup_Server_TacacsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Aaa_ServerGroup_Server_Tacacs
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Aaa_ServerGroup_Server_TacacsWatcher) Await(t testing.TB) (*QualifiedSystem_Aaa_ServerGroup_Server_Tacacs, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Alarm is a *System_Alarm with a corresponding timestamp.
type QualifiedSystem_Alarm struct {
	*genutil.Metadata
	val     *System_Alarm // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Alarm) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Alarm sample, erroring out if not present.
func (q *QualifiedSystem_Alarm) Val(t testing.TB) *System_Alarm {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Alarm sample.
func (q *QualifiedSystem_Alarm) SetVal(v *System_Alarm) *QualifiedSystem_Alarm {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Alarm) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Alarm is a telemetry Collection whose Await method returns a slice of *System_Alarm samples.
type CollectionSystem_Alarm struct {
	W    *System_AlarmWatcher
	Data []*QualifiedSystem_Alarm
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Alarm) Await(t testing.TB) []*QualifiedSystem_Alarm {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_AlarmWatcher observes a stream of *System_Alarm samples.
type System_AlarmWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Alarm
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_AlarmWatcher) Await(t testing.TB) (*QualifiedSystem_Alarm, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Clock is a *System_Clock with a corresponding timestamp.
type QualifiedSystem_Clock struct {
	*genutil.Metadata
	val     *System_Clock // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Clock) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Clock sample, erroring out if not present.
func (q *QualifiedSystem_Clock) Val(t testing.TB) *System_Clock {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Clock sample.
func (q *QualifiedSystem_Clock) SetVal(v *System_Clock) *QualifiedSystem_Clock {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Clock) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Clock is a telemetry Collection whose Await method returns a slice of *System_Clock samples.
type CollectionSystem_Clock struct {
	W    *System_ClockWatcher
	Data []*QualifiedSystem_Clock
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Clock) Await(t testing.TB) []*QualifiedSystem_Clock {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_ClockWatcher observes a stream of *System_Clock samples.
type System_ClockWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Clock
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_ClockWatcher) Await(t testing.TB) (*QualifiedSystem_Clock, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu is a *System_Cpu with a corresponding timestamp.
type QualifiedSystem_Cpu struct {
	*genutil.Metadata
	val     *System_Cpu // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu sample, erroring out if not present.
func (q *QualifiedSystem_Cpu) Val(t testing.TB) *System_Cpu {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu sample.
func (q *QualifiedSystem_Cpu) SetVal(v *System_Cpu) *QualifiedSystem_Cpu {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu is a telemetry Collection whose Await method returns a slice of *System_Cpu samples.
type CollectionSystem_Cpu struct {
	W    *System_CpuWatcher
	Data []*QualifiedSystem_Cpu
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu) Await(t testing.TB) []*QualifiedSystem_Cpu {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_CpuWatcher observes a stream of *System_Cpu samples.
type System_CpuWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_CpuWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_HardwareInterrupt is a *System_Cpu_HardwareInterrupt with a corresponding timestamp.
type QualifiedSystem_Cpu_HardwareInterrupt struct {
	*genutil.Metadata
	val     *System_Cpu_HardwareInterrupt // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_HardwareInterrupt) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_HardwareInterrupt sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_HardwareInterrupt) Val(t testing.TB) *System_Cpu_HardwareInterrupt {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_HardwareInterrupt sample.
func (q *QualifiedSystem_Cpu_HardwareInterrupt) SetVal(v *System_Cpu_HardwareInterrupt) *QualifiedSystem_Cpu_HardwareInterrupt {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_HardwareInterrupt) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_HardwareInterrupt is a telemetry Collection whose Await method returns a slice of *System_Cpu_HardwareInterrupt samples.
type CollectionSystem_Cpu_HardwareInterrupt struct {
	W    *System_Cpu_HardwareInterruptWatcher
	Data []*QualifiedSystem_Cpu_HardwareInterrupt
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_HardwareInterrupt) Await(t testing.TB) []*QualifiedSystem_Cpu_HardwareInterrupt {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_HardwareInterruptWatcher observes a stream of *System_Cpu_HardwareInterrupt samples.
type System_Cpu_HardwareInterruptWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_HardwareInterrupt
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_HardwareInterruptWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_HardwareInterrupt, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_Idle is a *System_Cpu_Idle with a corresponding timestamp.
type QualifiedSystem_Cpu_Idle struct {
	*genutil.Metadata
	val     *System_Cpu_Idle // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_Idle) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_Idle sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_Idle) Val(t testing.TB) *System_Cpu_Idle {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_Idle sample.
func (q *QualifiedSystem_Cpu_Idle) SetVal(v *System_Cpu_Idle) *QualifiedSystem_Cpu_Idle {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_Idle) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_Idle is a telemetry Collection whose Await method returns a slice of *System_Cpu_Idle samples.
type CollectionSystem_Cpu_Idle struct {
	W    *System_Cpu_IdleWatcher
	Data []*QualifiedSystem_Cpu_Idle
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_Idle) Await(t testing.TB) []*QualifiedSystem_Cpu_Idle {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_IdleWatcher observes a stream of *System_Cpu_Idle samples.
type System_Cpu_IdleWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_Idle
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_IdleWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_Idle, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_Kernel is a *System_Cpu_Kernel with a corresponding timestamp.
type QualifiedSystem_Cpu_Kernel struct {
	*genutil.Metadata
	val     *System_Cpu_Kernel // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_Kernel) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_Kernel sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_Kernel) Val(t testing.TB) *System_Cpu_Kernel {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_Kernel sample.
func (q *QualifiedSystem_Cpu_Kernel) SetVal(v *System_Cpu_Kernel) *QualifiedSystem_Cpu_Kernel {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_Kernel) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_Kernel is a telemetry Collection whose Await method returns a slice of *System_Cpu_Kernel samples.
type CollectionSystem_Cpu_Kernel struct {
	W    *System_Cpu_KernelWatcher
	Data []*QualifiedSystem_Cpu_Kernel
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_Kernel) Await(t testing.TB) []*QualifiedSystem_Cpu_Kernel {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_KernelWatcher observes a stream of *System_Cpu_Kernel samples.
type System_Cpu_KernelWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_Kernel
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_KernelWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_Kernel, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_Nice is a *System_Cpu_Nice with a corresponding timestamp.
type QualifiedSystem_Cpu_Nice struct {
	*genutil.Metadata
	val     *System_Cpu_Nice // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_Nice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_Nice sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_Nice) Val(t testing.TB) *System_Cpu_Nice {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_Nice sample.
func (q *QualifiedSystem_Cpu_Nice) SetVal(v *System_Cpu_Nice) *QualifiedSystem_Cpu_Nice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_Nice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_Nice is a telemetry Collection whose Await method returns a slice of *System_Cpu_Nice samples.
type CollectionSystem_Cpu_Nice struct {
	W    *System_Cpu_NiceWatcher
	Data []*QualifiedSystem_Cpu_Nice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_Nice) Await(t testing.TB) []*QualifiedSystem_Cpu_Nice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_NiceWatcher observes a stream of *System_Cpu_Nice samples.
type System_Cpu_NiceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_Nice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_NiceWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_Nice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_SoftwareInterrupt is a *System_Cpu_SoftwareInterrupt with a corresponding timestamp.
type QualifiedSystem_Cpu_SoftwareInterrupt struct {
	*genutil.Metadata
	val     *System_Cpu_SoftwareInterrupt // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_SoftwareInterrupt) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_SoftwareInterrupt sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_SoftwareInterrupt) Val(t testing.TB) *System_Cpu_SoftwareInterrupt {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_SoftwareInterrupt sample.
func (q *QualifiedSystem_Cpu_SoftwareInterrupt) SetVal(v *System_Cpu_SoftwareInterrupt) *QualifiedSystem_Cpu_SoftwareInterrupt {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_SoftwareInterrupt) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_SoftwareInterrupt is a telemetry Collection whose Await method returns a slice of *System_Cpu_SoftwareInterrupt samples.
type CollectionSystem_Cpu_SoftwareInterrupt struct {
	W    *System_Cpu_SoftwareInterruptWatcher
	Data []*QualifiedSystem_Cpu_SoftwareInterrupt
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_SoftwareInterrupt) Await(t testing.TB) []*QualifiedSystem_Cpu_SoftwareInterrupt {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_SoftwareInterruptWatcher observes a stream of *System_Cpu_SoftwareInterrupt samples.
type System_Cpu_SoftwareInterruptWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_SoftwareInterrupt
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_SoftwareInterruptWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_SoftwareInterrupt, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_Total is a *System_Cpu_Total with a corresponding timestamp.
type QualifiedSystem_Cpu_Total struct {
	*genutil.Metadata
	val     *System_Cpu_Total // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_Total) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_Total sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_Total) Val(t testing.TB) *System_Cpu_Total {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_Total sample.
func (q *QualifiedSystem_Cpu_Total) SetVal(v *System_Cpu_Total) *QualifiedSystem_Cpu_Total {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_Total) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_Total is a telemetry Collection whose Await method returns a slice of *System_Cpu_Total samples.
type CollectionSystem_Cpu_Total struct {
	W    *System_Cpu_TotalWatcher
	Data []*QualifiedSystem_Cpu_Total
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_Total) Await(t testing.TB) []*QualifiedSystem_Cpu_Total {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_TotalWatcher observes a stream of *System_Cpu_Total samples.
type System_Cpu_TotalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_Total
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_TotalWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_Total, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_User is a *System_Cpu_User with a corresponding timestamp.
type QualifiedSystem_Cpu_User struct {
	*genutil.Metadata
	val     *System_Cpu_User // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_User) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_User sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_User) Val(t testing.TB) *System_Cpu_User {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_User sample.
func (q *QualifiedSystem_Cpu_User) SetVal(v *System_Cpu_User) *QualifiedSystem_Cpu_User {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_User) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_User is a telemetry Collection whose Await method returns a slice of *System_Cpu_User samples.
type CollectionSystem_Cpu_User struct {
	W    *System_Cpu_UserWatcher
	Data []*QualifiedSystem_Cpu_User
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_User) Await(t testing.TB) []*QualifiedSystem_Cpu_User {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_UserWatcher observes a stream of *System_Cpu_User samples.
type System_Cpu_UserWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_User
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_UserWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_User, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Cpu_Wait is a *System_Cpu_Wait with a corresponding timestamp.
type QualifiedSystem_Cpu_Wait struct {
	*genutil.Metadata
	val     *System_Cpu_Wait // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Cpu_Wait) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Cpu_Wait sample, erroring out if not present.
func (q *QualifiedSystem_Cpu_Wait) Val(t testing.TB) *System_Cpu_Wait {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Cpu_Wait sample.
func (q *QualifiedSystem_Cpu_Wait) SetVal(v *System_Cpu_Wait) *QualifiedSystem_Cpu_Wait {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Cpu_Wait) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Cpu_Wait is a telemetry Collection whose Await method returns a slice of *System_Cpu_Wait samples.
type CollectionSystem_Cpu_Wait struct {
	W    *System_Cpu_WaitWatcher
	Data []*QualifiedSystem_Cpu_Wait
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Cpu_Wait) Await(t testing.TB) []*QualifiedSystem_Cpu_Wait {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Cpu_WaitWatcher observes a stream of *System_Cpu_Wait samples.
type System_Cpu_WaitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Cpu_Wait
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Cpu_WaitWatcher) Await(t testing.TB) (*QualifiedSystem_Cpu_Wait, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Dns is a *System_Dns with a corresponding timestamp.
type QualifiedSystem_Dns struct {
	*genutil.Metadata
	val     *System_Dns // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Dns) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Dns sample, erroring out if not present.
func (q *QualifiedSystem_Dns) Val(t testing.TB) *System_Dns {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Dns sample.
func (q *QualifiedSystem_Dns) SetVal(v *System_Dns) *QualifiedSystem_Dns {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Dns) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Dns is a telemetry Collection whose Await method returns a slice of *System_Dns samples.
type CollectionSystem_Dns struct {
	W    *System_DnsWatcher
	Data []*QualifiedSystem_Dns
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Dns) Await(t testing.TB) []*QualifiedSystem_Dns {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_DnsWatcher observes a stream of *System_Dns samples.
type System_DnsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Dns
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_DnsWatcher) Await(t testing.TB) (*QualifiedSystem_Dns, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Dns_HostEntry is a *System_Dns_HostEntry with a corresponding timestamp.
type QualifiedSystem_Dns_HostEntry struct {
	*genutil.Metadata
	val     *System_Dns_HostEntry // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Dns_HostEntry) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Dns_HostEntry sample, erroring out if not present.
func (q *QualifiedSystem_Dns_HostEntry) Val(t testing.TB) *System_Dns_HostEntry {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Dns_HostEntry sample.
func (q *QualifiedSystem_Dns_HostEntry) SetVal(v *System_Dns_HostEntry) *QualifiedSystem_Dns_HostEntry {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Dns_HostEntry) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Dns_HostEntry is a telemetry Collection whose Await method returns a slice of *System_Dns_HostEntry samples.
type CollectionSystem_Dns_HostEntry struct {
	W    *System_Dns_HostEntryWatcher
	Data []*QualifiedSystem_Dns_HostEntry
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Dns_HostEntry) Await(t testing.TB) []*QualifiedSystem_Dns_HostEntry {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Dns_HostEntryWatcher observes a stream of *System_Dns_HostEntry samples.
type System_Dns_HostEntryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Dns_HostEntry
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Dns_HostEntryWatcher) Await(t testing.TB) (*QualifiedSystem_Dns_HostEntry, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Dns_Server is a *System_Dns_Server with a corresponding timestamp.
type QualifiedSystem_Dns_Server struct {
	*genutil.Metadata
	val     *System_Dns_Server // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Dns_Server) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Dns_Server sample, erroring out if not present.
func (q *QualifiedSystem_Dns_Server) Val(t testing.TB) *System_Dns_Server {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Dns_Server sample.
func (q *QualifiedSystem_Dns_Server) SetVal(v *System_Dns_Server) *QualifiedSystem_Dns_Server {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Dns_Server) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Dns_Server is a telemetry Collection whose Await method returns a slice of *System_Dns_Server samples.
type CollectionSystem_Dns_Server struct {
	W    *System_Dns_ServerWatcher
	Data []*QualifiedSystem_Dns_Server
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Dns_Server) Await(t testing.TB) []*QualifiedSystem_Dns_Server {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Dns_ServerWatcher observes a stream of *System_Dns_Server samples.
type System_Dns_ServerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Dns_Server
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Dns_ServerWatcher) Await(t testing.TB) (*QualifiedSystem_Dns_Server, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_License is a *System_License with a corresponding timestamp.
type QualifiedSystem_License struct {
	*genutil.Metadata
	val     *System_License // val is the sample value.
	present bool
}

func (q *QualifiedSystem_License) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_License sample, erroring out if not present.
func (q *QualifiedSystem_License) Val(t testing.TB) *System_License {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_License sample.
func (q *QualifiedSystem_License) SetVal(v *System_License) *QualifiedSystem_License {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_License) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_License is a telemetry Collection whose Await method returns a slice of *System_License samples.
type CollectionSystem_License struct {
	W    *System_LicenseWatcher
	Data []*QualifiedSystem_License
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_License) Await(t testing.TB) []*QualifiedSystem_License {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_LicenseWatcher observes a stream of *System_License samples.
type System_LicenseWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_License
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_LicenseWatcher) Await(t testing.TB) (*QualifiedSystem_License, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_License_License is a *System_License_License with a corresponding timestamp.
type QualifiedSystem_License_License struct {
	*genutil.Metadata
	val     *System_License_License // val is the sample value.
	present bool
}

func (q *QualifiedSystem_License_License) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_License_License sample, erroring out if not present.
func (q *QualifiedSystem_License_License) Val(t testing.TB) *System_License_License {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_License_License sample.
func (q *QualifiedSystem_License_License) SetVal(v *System_License_License) *QualifiedSystem_License_License {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_License_License) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_License_License is a telemetry Collection whose Await method returns a slice of *System_License_License samples.
type CollectionSystem_License_License struct {
	W    *System_License_LicenseWatcher
	Data []*QualifiedSystem_License_License
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_License_License) Await(t testing.TB) []*QualifiedSystem_License_License {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_License_LicenseWatcher observes a stream of *System_License_License samples.
type System_License_LicenseWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_License_License
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_License_LicenseWatcher) Await(t testing.TB) (*QualifiedSystem_License_License, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Logging is a *System_Logging with a corresponding timestamp.
type QualifiedSystem_Logging struct {
	*genutil.Metadata
	val     *System_Logging // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Logging) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Logging sample, erroring out if not present.
func (q *QualifiedSystem_Logging) Val(t testing.TB) *System_Logging {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Logging sample.
func (q *QualifiedSystem_Logging) SetVal(v *System_Logging) *QualifiedSystem_Logging {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Logging) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Logging is a telemetry Collection whose Await method returns a slice of *System_Logging samples.
type CollectionSystem_Logging struct {
	W    *System_LoggingWatcher
	Data []*QualifiedSystem_Logging
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Logging) Await(t testing.TB) []*QualifiedSystem_Logging {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_LoggingWatcher observes a stream of *System_Logging samples.
type System_LoggingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Logging
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_LoggingWatcher) Await(t testing.TB) (*QualifiedSystem_Logging, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Logging_Console is a *System_Logging_Console with a corresponding timestamp.
type QualifiedSystem_Logging_Console struct {
	*genutil.Metadata
	val     *System_Logging_Console // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Logging_Console) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Logging_Console sample, erroring out if not present.
func (q *QualifiedSystem_Logging_Console) Val(t testing.TB) *System_Logging_Console {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Logging_Console sample.
func (q *QualifiedSystem_Logging_Console) SetVal(v *System_Logging_Console) *QualifiedSystem_Logging_Console {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Logging_Console) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Logging_Console is a telemetry Collection whose Await method returns a slice of *System_Logging_Console samples.
type CollectionSystem_Logging_Console struct {
	W    *System_Logging_ConsoleWatcher
	Data []*QualifiedSystem_Logging_Console
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Logging_Console) Await(t testing.TB) []*QualifiedSystem_Logging_Console {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Logging_ConsoleWatcher observes a stream of *System_Logging_Console samples.
type System_Logging_ConsoleWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Logging_Console
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Logging_ConsoleWatcher) Await(t testing.TB) (*QualifiedSystem_Logging_Console, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Logging_Console_Selector is a *System_Logging_Console_Selector with a corresponding timestamp.
type QualifiedSystem_Logging_Console_Selector struct {
	*genutil.Metadata
	val     *System_Logging_Console_Selector // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Logging_Console_Selector) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Logging_Console_Selector sample, erroring out if not present.
func (q *QualifiedSystem_Logging_Console_Selector) Val(t testing.TB) *System_Logging_Console_Selector {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Logging_Console_Selector sample.
func (q *QualifiedSystem_Logging_Console_Selector) SetVal(v *System_Logging_Console_Selector) *QualifiedSystem_Logging_Console_Selector {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Logging_Console_Selector) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Logging_Console_Selector is a telemetry Collection whose Await method returns a slice of *System_Logging_Console_Selector samples.
type CollectionSystem_Logging_Console_Selector struct {
	W    *System_Logging_Console_SelectorWatcher
	Data []*QualifiedSystem_Logging_Console_Selector
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Logging_Console_Selector) Await(t testing.TB) []*QualifiedSystem_Logging_Console_Selector {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Logging_Console_SelectorWatcher observes a stream of *System_Logging_Console_Selector samples.
type System_Logging_Console_SelectorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Logging_Console_Selector
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Logging_Console_SelectorWatcher) Await(t testing.TB) (*QualifiedSystem_Logging_Console_Selector, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Logging_RemoteServer is a *System_Logging_RemoteServer with a corresponding timestamp.
type QualifiedSystem_Logging_RemoteServer struct {
	*genutil.Metadata
	val     *System_Logging_RemoteServer // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Logging_RemoteServer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Logging_RemoteServer sample, erroring out if not present.
func (q *QualifiedSystem_Logging_RemoteServer) Val(t testing.TB) *System_Logging_RemoteServer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Logging_RemoteServer sample.
func (q *QualifiedSystem_Logging_RemoteServer) SetVal(v *System_Logging_RemoteServer) *QualifiedSystem_Logging_RemoteServer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Logging_RemoteServer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Logging_RemoteServer is a telemetry Collection whose Await method returns a slice of *System_Logging_RemoteServer samples.
type CollectionSystem_Logging_RemoteServer struct {
	W    *System_Logging_RemoteServerWatcher
	Data []*QualifiedSystem_Logging_RemoteServer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Logging_RemoteServer) Await(t testing.TB) []*QualifiedSystem_Logging_RemoteServer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Logging_RemoteServerWatcher observes a stream of *System_Logging_RemoteServer samples.
type System_Logging_RemoteServerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Logging_RemoteServer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Logging_RemoteServerWatcher) Await(t testing.TB) (*QualifiedSystem_Logging_RemoteServer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Logging_RemoteServer_Selector is a *System_Logging_RemoteServer_Selector with a corresponding timestamp.
type QualifiedSystem_Logging_RemoteServer_Selector struct {
	*genutil.Metadata
	val     *System_Logging_RemoteServer_Selector // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Logging_RemoteServer_Selector) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Logging_RemoteServer_Selector sample, erroring out if not present.
func (q *QualifiedSystem_Logging_RemoteServer_Selector) Val(t testing.TB) *System_Logging_RemoteServer_Selector {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Logging_RemoteServer_Selector sample.
func (q *QualifiedSystem_Logging_RemoteServer_Selector) SetVal(v *System_Logging_RemoteServer_Selector) *QualifiedSystem_Logging_RemoteServer_Selector {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Logging_RemoteServer_Selector) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Logging_RemoteServer_Selector is a telemetry Collection whose Await method returns a slice of *System_Logging_RemoteServer_Selector samples.
type CollectionSystem_Logging_RemoteServer_Selector struct {
	W    *System_Logging_RemoteServer_SelectorWatcher
	Data []*QualifiedSystem_Logging_RemoteServer_Selector
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Logging_RemoteServer_Selector) Await(t testing.TB) []*QualifiedSystem_Logging_RemoteServer_Selector {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Logging_RemoteServer_SelectorWatcher observes a stream of *System_Logging_RemoteServer_Selector samples.
type System_Logging_RemoteServer_SelectorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Logging_RemoteServer_Selector
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Logging_RemoteServer_SelectorWatcher) Await(t testing.TB) (*QualifiedSystem_Logging_RemoteServer_Selector, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Memory is a *System_Memory with a corresponding timestamp.
type QualifiedSystem_Memory struct {
	*genutil.Metadata
	val     *System_Memory // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Memory) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Memory sample, erroring out if not present.
func (q *QualifiedSystem_Memory) Val(t testing.TB) *System_Memory {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Memory sample.
func (q *QualifiedSystem_Memory) SetVal(v *System_Memory) *QualifiedSystem_Memory {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Memory) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Memory is a telemetry Collection whose Await method returns a slice of *System_Memory samples.
type CollectionSystem_Memory struct {
	W    *System_MemoryWatcher
	Data []*QualifiedSystem_Memory
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Memory) Await(t testing.TB) []*QualifiedSystem_Memory {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_MemoryWatcher observes a stream of *System_Memory samples.
type System_MemoryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Memory
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_MemoryWatcher) Await(t testing.TB) (*QualifiedSystem_Memory, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Memory_Counters is a *System_Memory_Counters with a corresponding timestamp.
type QualifiedSystem_Memory_Counters struct {
	*genutil.Metadata
	val     *System_Memory_Counters // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Memory_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Memory_Counters sample, erroring out if not present.
func (q *QualifiedSystem_Memory_Counters) Val(t testing.TB) *System_Memory_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Memory_Counters sample.
func (q *QualifiedSystem_Memory_Counters) SetVal(v *System_Memory_Counters) *QualifiedSystem_Memory_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Memory_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Memory_Counters is a telemetry Collection whose Await method returns a slice of *System_Memory_Counters samples.
type CollectionSystem_Memory_Counters struct {
	W    *System_Memory_CountersWatcher
	Data []*QualifiedSystem_Memory_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Memory_Counters) Await(t testing.TB) []*QualifiedSystem_Memory_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Memory_CountersWatcher observes a stream of *System_Memory_Counters samples.
type System_Memory_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Memory_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Memory_CountersWatcher) Await(t testing.TB) (*QualifiedSystem_Memory_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Messages is a *System_Messages with a corresponding timestamp.
type QualifiedSystem_Messages struct {
	*genutil.Metadata
	val     *System_Messages // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Messages) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Messages sample, erroring out if not present.
func (q *QualifiedSystem_Messages) Val(t testing.TB) *System_Messages {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Messages sample.
func (q *QualifiedSystem_Messages) SetVal(v *System_Messages) *QualifiedSystem_Messages {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Messages) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Messages is a telemetry Collection whose Await method returns a slice of *System_Messages samples.
type CollectionSystem_Messages struct {
	W    *System_MessagesWatcher
	Data []*QualifiedSystem_Messages
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Messages) Await(t testing.TB) []*QualifiedSystem_Messages {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_MessagesWatcher observes a stream of *System_Messages samples.
type System_MessagesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Messages
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_MessagesWatcher) Await(t testing.TB) (*QualifiedSystem_Messages, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Messages_DebugService is a *System_Messages_DebugService with a corresponding timestamp.
type QualifiedSystem_Messages_DebugService struct {
	*genutil.Metadata
	val     *System_Messages_DebugService // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Messages_DebugService) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Messages_DebugService sample, erroring out if not present.
func (q *QualifiedSystem_Messages_DebugService) Val(t testing.TB) *System_Messages_DebugService {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Messages_DebugService sample.
func (q *QualifiedSystem_Messages_DebugService) SetVal(v *System_Messages_DebugService) *QualifiedSystem_Messages_DebugService {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Messages_DebugService) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Messages_DebugService is a telemetry Collection whose Await method returns a slice of *System_Messages_DebugService samples.
type CollectionSystem_Messages_DebugService struct {
	W    *System_Messages_DebugServiceWatcher
	Data []*QualifiedSystem_Messages_DebugService
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Messages_DebugService) Await(t testing.TB) []*QualifiedSystem_Messages_DebugService {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Messages_DebugServiceWatcher observes a stream of *System_Messages_DebugService samples.
type System_Messages_DebugServiceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Messages_DebugService
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Messages_DebugServiceWatcher) Await(t testing.TB) (*QualifiedSystem_Messages_DebugService, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Messages_Message is a *System_Messages_Message with a corresponding timestamp.
type QualifiedSystem_Messages_Message struct {
	*genutil.Metadata
	val     *System_Messages_Message // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Messages_Message) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Messages_Message sample, erroring out if not present.
func (q *QualifiedSystem_Messages_Message) Val(t testing.TB) *System_Messages_Message {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Messages_Message sample.
func (q *QualifiedSystem_Messages_Message) SetVal(v *System_Messages_Message) *QualifiedSystem_Messages_Message {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Messages_Message) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Messages_Message is a telemetry Collection whose Await method returns a slice of *System_Messages_Message samples.
type CollectionSystem_Messages_Message struct {
	W    *System_Messages_MessageWatcher
	Data []*QualifiedSystem_Messages_Message
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Messages_Message) Await(t testing.TB) []*QualifiedSystem_Messages_Message {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Messages_MessageWatcher observes a stream of *System_Messages_Message samples.
type System_Messages_MessageWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Messages_Message
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Messages_MessageWatcher) Await(t testing.TB) (*QualifiedSystem_Messages_Message, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_MountPoint is a *System_MountPoint with a corresponding timestamp.
type QualifiedSystem_MountPoint struct {
	*genutil.Metadata
	val     *System_MountPoint // val is the sample value.
	present bool
}

func (q *QualifiedSystem_MountPoint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_MountPoint sample, erroring out if not present.
func (q *QualifiedSystem_MountPoint) Val(t testing.TB) *System_MountPoint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_MountPoint sample.
func (q *QualifiedSystem_MountPoint) SetVal(v *System_MountPoint) *QualifiedSystem_MountPoint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_MountPoint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_MountPoint is a telemetry Collection whose Await method returns a slice of *System_MountPoint samples.
type CollectionSystem_MountPoint struct {
	W    *System_MountPointWatcher
	Data []*QualifiedSystem_MountPoint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_MountPoint) Await(t testing.TB) []*QualifiedSystem_MountPoint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_MountPointWatcher observes a stream of *System_MountPoint samples.
type System_MountPointWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_MountPoint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_MountPointWatcher) Await(t testing.TB) (*QualifiedSystem_MountPoint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Ntp is a *System_Ntp with a corresponding timestamp.
type QualifiedSystem_Ntp struct {
	*genutil.Metadata
	val     *System_Ntp // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Ntp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Ntp sample, erroring out if not present.
func (q *QualifiedSystem_Ntp) Val(t testing.TB) *System_Ntp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Ntp sample.
func (q *QualifiedSystem_Ntp) SetVal(v *System_Ntp) *QualifiedSystem_Ntp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Ntp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Ntp is a telemetry Collection whose Await method returns a slice of *System_Ntp samples.
type CollectionSystem_Ntp struct {
	W    *System_NtpWatcher
	Data []*QualifiedSystem_Ntp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Ntp) Await(t testing.TB) []*QualifiedSystem_Ntp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_NtpWatcher observes a stream of *System_Ntp samples.
type System_NtpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Ntp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_NtpWatcher) Await(t testing.TB) (*QualifiedSystem_Ntp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Ntp_NtpKey is a *System_Ntp_NtpKey with a corresponding timestamp.
type QualifiedSystem_Ntp_NtpKey struct {
	*genutil.Metadata
	val     *System_Ntp_NtpKey // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Ntp_NtpKey) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Ntp_NtpKey sample, erroring out if not present.
func (q *QualifiedSystem_Ntp_NtpKey) Val(t testing.TB) *System_Ntp_NtpKey {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Ntp_NtpKey sample.
func (q *QualifiedSystem_Ntp_NtpKey) SetVal(v *System_Ntp_NtpKey) *QualifiedSystem_Ntp_NtpKey {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Ntp_NtpKey) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Ntp_NtpKey is a telemetry Collection whose Await method returns a slice of *System_Ntp_NtpKey samples.
type CollectionSystem_Ntp_NtpKey struct {
	W    *System_Ntp_NtpKeyWatcher
	Data []*QualifiedSystem_Ntp_NtpKey
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Ntp_NtpKey) Await(t testing.TB) []*QualifiedSystem_Ntp_NtpKey {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Ntp_NtpKeyWatcher observes a stream of *System_Ntp_NtpKey samples.
type System_Ntp_NtpKeyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Ntp_NtpKey
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Ntp_NtpKeyWatcher) Await(t testing.TB) (*QualifiedSystem_Ntp_NtpKey, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Ntp_Server is a *System_Ntp_Server with a corresponding timestamp.
type QualifiedSystem_Ntp_Server struct {
	*genutil.Metadata
	val     *System_Ntp_Server // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Ntp_Server) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Ntp_Server sample, erroring out if not present.
func (q *QualifiedSystem_Ntp_Server) Val(t testing.TB) *System_Ntp_Server {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Ntp_Server sample.
func (q *QualifiedSystem_Ntp_Server) SetVal(v *System_Ntp_Server) *QualifiedSystem_Ntp_Server {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Ntp_Server) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Ntp_Server is a telemetry Collection whose Await method returns a slice of *System_Ntp_Server samples.
type CollectionSystem_Ntp_Server struct {
	W    *System_Ntp_ServerWatcher
	Data []*QualifiedSystem_Ntp_Server
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Ntp_Server) Await(t testing.TB) []*QualifiedSystem_Ntp_Server {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_Ntp_ServerWatcher observes a stream of *System_Ntp_Server samples.
type System_Ntp_ServerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Ntp_Server
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_Ntp_ServerWatcher) Await(t testing.TB) (*QualifiedSystem_Ntp_Server, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_Process is a *System_Process with a corresponding timestamp.
type QualifiedSystem_Process struct {
	*genutil.Metadata
	val     *System_Process // val is the sample value.
	present bool
}

func (q *QualifiedSystem_Process) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_Process sample, erroring out if not present.
func (q *QualifiedSystem_Process) Val(t testing.TB) *System_Process {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_Process sample.
func (q *QualifiedSystem_Process) SetVal(v *System_Process) *QualifiedSystem_Process {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_Process) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_Process is a telemetry Collection whose Await method returns a slice of *System_Process samples.
type CollectionSystem_Process struct {
	W    *System_ProcessWatcher
	Data []*QualifiedSystem_Process
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_Process) Await(t testing.TB) []*QualifiedSystem_Process {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_ProcessWatcher observes a stream of *System_Process samples.
type System_ProcessWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_Process
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_ProcessWatcher) Await(t testing.TB) (*QualifiedSystem_Process, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_SshServer is a *System_SshServer with a corresponding timestamp.
type QualifiedSystem_SshServer struct {
	*genutil.Metadata
	val     *System_SshServer // val is the sample value.
	present bool
}

func (q *QualifiedSystem_SshServer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_SshServer sample, erroring out if not present.
func (q *QualifiedSystem_SshServer) Val(t testing.TB) *System_SshServer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_SshServer sample.
func (q *QualifiedSystem_SshServer) SetVal(v *System_SshServer) *QualifiedSystem_SshServer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_SshServer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_SshServer is a telemetry Collection whose Await method returns a slice of *System_SshServer samples.
type CollectionSystem_SshServer struct {
	W    *System_SshServerWatcher
	Data []*QualifiedSystem_SshServer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_SshServer) Await(t testing.TB) []*QualifiedSystem_SshServer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_SshServerWatcher observes a stream of *System_SshServer samples.
type System_SshServerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_SshServer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_SshServerWatcher) Await(t testing.TB) (*QualifiedSystem_SshServer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedSystem_TelnetServer is a *System_TelnetServer with a corresponding timestamp.
type QualifiedSystem_TelnetServer struct {
	*genutil.Metadata
	val     *System_TelnetServer // val is the sample value.
	present bool
}

func (q *QualifiedSystem_TelnetServer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *System_TelnetServer sample, erroring out if not present.
func (q *QualifiedSystem_TelnetServer) Val(t testing.TB) *System_TelnetServer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *System_TelnetServer sample.
func (q *QualifiedSystem_TelnetServer) SetVal(v *System_TelnetServer) *QualifiedSystem_TelnetServer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedSystem_TelnetServer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionSystem_TelnetServer is a telemetry Collection whose Await method returns a slice of *System_TelnetServer samples.
type CollectionSystem_TelnetServer struct {
	W    *System_TelnetServerWatcher
	Data []*QualifiedSystem_TelnetServer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionSystem_TelnetServer) Await(t testing.TB) []*QualifiedSystem_TelnetServer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// System_TelnetServerWatcher observes a stream of *System_TelnetServer samples.
type System_TelnetServerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedSystem_TelnetServer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *System_TelnetServerWatcher) Await(t testing.TB) (*QualifiedSystem_TelnetServer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union is a Acl_AclSet_AclEntry_Ipv4_Protocol_Union with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union struct {
	*genutil.Metadata
	val     Acl_AclSet_AclEntry_Ipv4_Protocol_Union // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Acl_AclSet_AclEntry_Ipv4_Protocol_Union sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) Val(t testing.TB) Acl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Acl_AclSet_AclEntry_Ipv4_Protocol_Union sample.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) SetVal(v Acl_AclSet_AclEntry_Ipv4_Protocol_Union) *QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Ipv4_Protocol_Union is a telemetry Collection whose Await method returns a slice of Acl_AclSet_AclEntry_Ipv4_Protocol_Union samples.
type CollectionAcl_AclSet_AclEntry_Ipv4_Protocol_Union struct {
	W    *Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Ipv4_Protocol_Union) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher observes a stream of Acl_AclSet_AclEntry_Ipv4_Protocol_Union samples.
type Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Ipv4_Protocol_UnionWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Ipv4_Protocol_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union is a Acl_AclSet_AclEntry_Ipv6_Protocol_Union with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union struct {
	*genutil.Metadata
	val     Acl_AclSet_AclEntry_Ipv6_Protocol_Union // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Acl_AclSet_AclEntry_Ipv6_Protocol_Union sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union) Val(t testing.TB) Acl_AclSet_AclEntry_Ipv6_Protocol_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Acl_AclSet_AclEntry_Ipv6_Protocol_Union sample.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union) SetVal(v Acl_AclSet_AclEntry_Ipv6_Protocol_Union) *QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Ipv6_Protocol_Union is a telemetry Collection whose Await method returns a slice of Acl_AclSet_AclEntry_Ipv6_Protocol_Union samples.
type CollectionAcl_AclSet_AclEntry_Ipv6_Protocol_Union struct {
	W    *Acl_AclSet_AclEntry_Ipv6_Protocol_UnionWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Ipv6_Protocol_Union) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Ipv6_Protocol_UnionWatcher observes a stream of Acl_AclSet_AclEntry_Ipv6_Protocol_Union samples.
type Acl_AclSet_AclEntry_Ipv6_Protocol_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Ipv6_Protocol_UnionWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Ipv6_Protocol_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union is a Acl_AclSet_AclEntry_L2_Ethertype_Union with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union struct {
	*genutil.Metadata
	val     Acl_AclSet_AclEntry_L2_Ethertype_Union // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Acl_AclSet_AclEntry_L2_Ethertype_Union sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) Val(t testing.TB) Acl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Acl_AclSet_AclEntry_L2_Ethertype_Union sample.
func (q *QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) SetVal(v Acl_AclSet_AclEntry_L2_Ethertype_Union) *QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_L2_Ethertype_Union is a telemetry Collection whose Await method returns a slice of Acl_AclSet_AclEntry_L2_Ethertype_Union samples.
type CollectionAcl_AclSet_AclEntry_L2_Ethertype_Union struct {
	W    *Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_L2_Ethertype_Union) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher observes a stream of Acl_AclSet_AclEntry_L2_Ethertype_Union samples.
type Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_L2_Ethertype_UnionWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_L2_Ethertype_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union is a Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union struct {
	*genutil.Metadata
	val     Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) Val(t testing.TB) Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union sample.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) SetVal(v Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union) *QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union is a telemetry Collection whose Await method returns a slice of Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union samples.
type CollectionAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union struct {
	W    *Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher observes a stream of Acl_AclSet_AclEntry_Mpls_EndLabelValue_Union samples.
type Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Mpls_EndLabelValue_UnionWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Mpls_EndLabelValue_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union is a Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union struct {
	*genutil.Metadata
	val     Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) Val(t testing.TB) Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union sample.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) SetVal(v Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union) *QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union is a telemetry Collection whose Await method returns a slice of Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union samples.
type CollectionAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union struct {
	W    *Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher observes a stream of Acl_AclSet_AclEntry_Mpls_StartLabelValue_Union samples.
type Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Mpls_StartLabelValue_UnionWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Mpls_StartLabelValue_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union is a Acl_AclSet_AclEntry_Transport_DestinationPort_Union with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union struct {
	*genutil.Metadata
	val     Acl_AclSet_AclEntry_Transport_DestinationPort_Union // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Acl_AclSet_AclEntry_Transport_DestinationPort_Union sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) Val(t testing.TB) Acl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Acl_AclSet_AclEntry_Transport_DestinationPort_Union sample.
func (q *QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) SetVal(v Acl_AclSet_AclEntry_Transport_DestinationPort_Union) *QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Transport_DestinationPort_Union is a telemetry Collection whose Await method returns a slice of Acl_AclSet_AclEntry_Transport_DestinationPort_Union samples.
type CollectionAcl_AclSet_AclEntry_Transport_DestinationPort_Union struct {
	W    *Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Transport_DestinationPort_Union) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher observes a stream of Acl_AclSet_AclEntry_Transport_DestinationPort_Union samples.
type Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Transport_DestinationPort_UnionWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Transport_DestinationPort_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union is a Acl_AclSet_AclEntry_Transport_SourcePort_Union with a corresponding timestamp.
type QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union struct {
	*genutil.Metadata
	val     Acl_AclSet_AclEntry_Transport_SourcePort_Union // val is the sample value.
	present bool
}

func (q *QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Acl_AclSet_AclEntry_Transport_SourcePort_Union sample, erroring out if not present.
func (q *QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) Val(t testing.TB) Acl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Acl_AclSet_AclEntry_Transport_SourcePort_Union sample.
func (q *QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) SetVal(v Acl_AclSet_AclEntry_Transport_SourcePort_Union) *QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionAcl_AclSet_AclEntry_Transport_SourcePort_Union is a telemetry Collection whose Await method returns a slice of Acl_AclSet_AclEntry_Transport_SourcePort_Union samples.
type CollectionAcl_AclSet_AclEntry_Transport_SourcePort_Union struct {
	W    *Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher
	Data []*QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionAcl_AclSet_AclEntry_Transport_SourcePort_Union) Await(t testing.TB) []*QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher observes a stream of Acl_AclSet_AclEntry_Transport_SourcePort_Union samples.
type Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Acl_AclSet_AclEntry_Transport_SourcePort_UnionWatcher) Await(t testing.TB) (*QualifiedAcl_AclSet_AclEntry_Transport_SourcePort_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBinary is a Binary with a corresponding timestamp.
type QualifiedBinary struct {
	*genutil.Metadata
	val     Binary // val is the sample value.
	present bool
}

func (q *QualifiedBinary) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Binary sample, erroring out if not present.
func (q *QualifiedBinary) Val(t testing.TB) Binary {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Binary sample.
func (q *QualifiedBinary) SetVal(v Binary) *QualifiedBinary {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBinary) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBinary is a telemetry Collection whose Await method returns a slice of Binary samples.
type CollectionBinary struct {
	W    *BinaryWatcher
	Data []*QualifiedBinary
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBinary) Await(t testing.TB) []*QualifiedBinary {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BinaryWatcher observes a stream of Binary samples.
type BinaryWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBinary
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BinaryWatcher) Await(t testing.TB) (*QualifiedBinary, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Property_Value_Union is a Component_Property_Value_Union with a corresponding timestamp.
type QualifiedComponent_Property_Value_Union struct {
	*genutil.Metadata
	val     Component_Property_Value_Union // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Property_Value_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Component_Property_Value_Union sample, erroring out if not present.
func (q *QualifiedComponent_Property_Value_Union) Val(t testing.TB) Component_Property_Value_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Component_Property_Value_Union sample.
func (q *QualifiedComponent_Property_Value_Union) SetVal(v Component_Property_Value_Union) *QualifiedComponent_Property_Value_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Property_Value_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Property_Value_Union is a telemetry Collection whose Await method returns a slice of Component_Property_Value_Union samples.
type CollectionComponent_Property_Value_Union struct {
	W    *Component_Property_Value_UnionWatcher
	Data []*QualifiedComponent_Property_Value_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Property_Value_Union) Await(t testing.TB) []*QualifiedComponent_Property_Value_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Property_Value_UnionWatcher observes a stream of Component_Property_Value_Union samples.
type Component_Property_Value_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Property_Value_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Property_Value_UnionWatcher) Await(t testing.TB) (*QualifiedComponent_Property_Value_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedComponent_Type_Union is a Component_Type_Union with a corresponding timestamp.
type QualifiedComponent_Type_Union struct {
	*genutil.Metadata
	val     Component_Type_Union // val is the sample value.
	present bool
}

func (q *QualifiedComponent_Type_Union) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the Component_Type_Union sample, erroring out if not present.
func (q *QualifiedComponent_Type_Union) Val(t testing.TB) Component_Type_Union {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the Component_Type_Union sample.
func (q *QualifiedComponent_Type_Union) SetVal(v Component_Type_Union) *QualifiedComponent_Type_Union {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedComponent_Type_Union) IsPresent() bool {
	return q != nil && q.present
}

// CollectionComponent_Type_Union is a telemetry Collection whose Await method returns a slice of Component_Type_Union samples.
type CollectionComponent_Type_Union struct {
	W    *Component_Type_UnionWatcher
	Data []*QualifiedComponent_Type_Union
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionComponent_Type_Union) Await(t testing.TB) []*QualifiedComponent_Type_Union {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Component_Type_UnionWatcher observes a stream of Component_Type_Union samples.
type Component_Type_UnionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedComponent_Type_Union
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Component_Type_UnionWatcher) Await(t testing.TB) (*QualifiedComponent_Type_Union, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE is a E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE with a corresponding timestamp.
type QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE struct {
	*genutil.Metadata
	val     E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE sample, erroring out if not present.
func (q *QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) Val(t testing.TB) E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE sample.
func (q *QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) SetVal(v E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) *QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE is a telemetry Collection whose Await method returns a slice of E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE samples.
type CollectionE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE struct {
	W    *E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher
	Data []*QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) Await(t testing.TB) []*QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher observes a stream of E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE samples.
type E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPEWatcher) Await(t testing.TB) (*QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE is a E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE with a corresponding timestamp.
type QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE struct {
	*genutil.Metadata
	val     E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE sample, erroring out if not present.
func (q *QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) Val(t testing.TB) E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE sample.
func (q *QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) SetVal(v E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) *QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE is a telemetry Collection whose Await method returns a slice of E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE samples.
type CollectionE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE struct {
	W    *E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher
	Data []*QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) Await(t testing.TB) []*QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher observes a stream of E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE samples.
type E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPEWatcher) Await(t testing.TB) (*QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_AaaTypes_AAA_SERVER_TYPE is a E_AaaTypes_AAA_SERVER_TYPE with a corresponding timestamp.
type QualifiedE_AaaTypes_AAA_SERVER_TYPE struct {
	*genutil.Metadata
	val     E_AaaTypes_AAA_SERVER_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_AaaTypes_AAA_SERVER_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_AaaTypes_AAA_SERVER_TYPE sample, erroring out if not present.
func (q *QualifiedE_AaaTypes_AAA_SERVER_TYPE) Val(t testing.TB) E_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_AaaTypes_AAA_SERVER_TYPE sample.
func (q *QualifiedE_AaaTypes_AAA_SERVER_TYPE) SetVal(v E_AaaTypes_AAA_SERVER_TYPE) *QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_AaaTypes_AAA_SERVER_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_AaaTypes_AAA_SERVER_TYPE is a telemetry Collection whose Await method returns a slice of E_AaaTypes_AAA_SERVER_TYPE samples.
type CollectionE_AaaTypes_AAA_SERVER_TYPE struct {
	W    *E_AaaTypes_AAA_SERVER_TYPEWatcher
	Data []*QualifiedE_AaaTypes_AAA_SERVER_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_AaaTypes_AAA_SERVER_TYPE) Await(t testing.TB) []*QualifiedE_AaaTypes_AAA_SERVER_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_AaaTypes_AAA_SERVER_TYPEWatcher observes a stream of E_AaaTypes_AAA_SERVER_TYPE samples.
type E_AaaTypes_AAA_SERVER_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_AaaTypes_AAA_SERVER_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_AaaTypes_AAA_SERVER_TYPEWatcher) Await(t testing.TB) (*QualifiedE_AaaTypes_AAA_SERVER_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Acl_ACL_COUNTER_CAPABILITY is a E_Acl_ACL_COUNTER_CAPABILITY with a corresponding timestamp.
type QualifiedE_Acl_ACL_COUNTER_CAPABILITY struct {
	*genutil.Metadata
	val     E_Acl_ACL_COUNTER_CAPABILITY // val is the sample value.
	present bool
}

func (q *QualifiedE_Acl_ACL_COUNTER_CAPABILITY) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Acl_ACL_COUNTER_CAPABILITY sample, erroring out if not present.
func (q *QualifiedE_Acl_ACL_COUNTER_CAPABILITY) Val(t testing.TB) E_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Acl_ACL_COUNTER_CAPABILITY sample.
func (q *QualifiedE_Acl_ACL_COUNTER_CAPABILITY) SetVal(v E_Acl_ACL_COUNTER_CAPABILITY) *QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Acl_ACL_COUNTER_CAPABILITY) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Acl_ACL_COUNTER_CAPABILITY is a telemetry Collection whose Await method returns a slice of E_Acl_ACL_COUNTER_CAPABILITY samples.
type CollectionE_Acl_ACL_COUNTER_CAPABILITY struct {
	W    *E_Acl_ACL_COUNTER_CAPABILITYWatcher
	Data []*QualifiedE_Acl_ACL_COUNTER_CAPABILITY
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Acl_ACL_COUNTER_CAPABILITY) Await(t testing.TB) []*QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Acl_ACL_COUNTER_CAPABILITYWatcher observes a stream of E_Acl_ACL_COUNTER_CAPABILITY samples.
type E_Acl_ACL_COUNTER_CAPABILITYWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Acl_ACL_COUNTER_CAPABILITY
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Acl_ACL_COUNTER_CAPABILITYWatcher) Await(t testing.TB) (*QualifiedE_Acl_ACL_COUNTER_CAPABILITY, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Acl_ACL_TYPE is a E_Acl_ACL_TYPE with a corresponding timestamp.
type QualifiedE_Acl_ACL_TYPE struct {
	*genutil.Metadata
	val     E_Acl_ACL_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_Acl_ACL_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Acl_ACL_TYPE sample, erroring out if not present.
func (q *QualifiedE_Acl_ACL_TYPE) Val(t testing.TB) E_Acl_ACL_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Acl_ACL_TYPE sample.
func (q *QualifiedE_Acl_ACL_TYPE) SetVal(v E_Acl_ACL_TYPE) *QualifiedE_Acl_ACL_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Acl_ACL_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Acl_ACL_TYPE is a telemetry Collection whose Await method returns a slice of E_Acl_ACL_TYPE samples.
type CollectionE_Acl_ACL_TYPE struct {
	W    *E_Acl_ACL_TYPEWatcher
	Data []*QualifiedE_Acl_ACL_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Acl_ACL_TYPE) Await(t testing.TB) []*QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Acl_ACL_TYPEWatcher observes a stream of E_Acl_ACL_TYPE samples.
type E_Acl_ACL_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Acl_ACL_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Acl_ACL_TYPEWatcher) Await(t testing.TB) (*QualifiedE_Acl_ACL_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Acl_FORWARDING_ACTION is a E_Acl_FORWARDING_ACTION with a corresponding timestamp.
type QualifiedE_Acl_FORWARDING_ACTION struct {
	*genutil.Metadata
	val     E_Acl_FORWARDING_ACTION // val is the sample value.
	present bool
}

func (q *QualifiedE_Acl_FORWARDING_ACTION) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Acl_FORWARDING_ACTION sample, erroring out if not present.
func (q *QualifiedE_Acl_FORWARDING_ACTION) Val(t testing.TB) E_Acl_FORWARDING_ACTION {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Acl_FORWARDING_ACTION sample.
func (q *QualifiedE_Acl_FORWARDING_ACTION) SetVal(v E_Acl_FORWARDING_ACTION) *QualifiedE_Acl_FORWARDING_ACTION {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Acl_FORWARDING_ACTION) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Acl_FORWARDING_ACTION is a telemetry Collection whose Await method returns a slice of E_Acl_FORWARDING_ACTION samples.
type CollectionE_Acl_FORWARDING_ACTION struct {
	W    *E_Acl_FORWARDING_ACTIONWatcher
	Data []*QualifiedE_Acl_FORWARDING_ACTION
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Acl_FORWARDING_ACTION) Await(t testing.TB) []*QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Acl_FORWARDING_ACTIONWatcher observes a stream of E_Acl_FORWARDING_ACTION samples.
type E_Acl_FORWARDING_ACTIONWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Acl_FORWARDING_ACTION
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Acl_FORWARDING_ACTIONWatcher) Await(t testing.TB) (*QualifiedE_Acl_FORWARDING_ACTION, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Acl_LOG_ACTION is a E_Acl_LOG_ACTION with a corresponding timestamp.
type QualifiedE_Acl_LOG_ACTION struct {
	*genutil.Metadata
	val     E_Acl_LOG_ACTION // val is the sample value.
	present bool
}

func (q *QualifiedE_Acl_LOG_ACTION) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Acl_LOG_ACTION sample, erroring out if not present.
func (q *QualifiedE_Acl_LOG_ACTION) Val(t testing.TB) E_Acl_LOG_ACTION {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Acl_LOG_ACTION sample.
func (q *QualifiedE_Acl_LOG_ACTION) SetVal(v E_Acl_LOG_ACTION) *QualifiedE_Acl_LOG_ACTION {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Acl_LOG_ACTION) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Acl_LOG_ACTION is a telemetry Collection whose Await method returns a slice of E_Acl_LOG_ACTION samples.
type CollectionE_Acl_LOG_ACTION struct {
	W    *E_Acl_LOG_ACTIONWatcher
	Data []*QualifiedE_Acl_LOG_ACTION
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Acl_LOG_ACTION) Await(t testing.TB) []*QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Acl_LOG_ACTIONWatcher observes a stream of E_Acl_LOG_ACTION samples.
type E_Acl_LOG_ACTIONWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Acl_LOG_ACTION
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Acl_LOG_ACTIONWatcher) Await(t testing.TB) (*QualifiedE_Acl_LOG_ACTION, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Address_Status is a E_Address_Status with a corresponding timestamp.
type QualifiedE_Address_Status struct {
	*genutil.Metadata
	val     E_Address_Status // val is the sample value.
	present bool
}

func (q *QualifiedE_Address_Status) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Address_Status sample, erroring out if not present.
func (q *QualifiedE_Address_Status) Val(t testing.TB) E_Address_Status {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Address_Status sample.
func (q *QualifiedE_Address_Status) SetVal(v E_Address_Status) *QualifiedE_Address_Status {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Address_Status) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Address_Status is a telemetry Collection whose Await method returns a slice of E_Address_Status samples.
type CollectionE_Address_Status struct {
	W    *E_Address_StatusWatcher
	Data []*QualifiedE_Address_Status
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Address_Status) Await(t testing.TB) []*QualifiedE_Address_Status {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Address_StatusWatcher observes a stream of E_Address_Status samples.
type E_Address_StatusWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Address_Status
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Address_StatusWatcher) Await(t testing.TB) (*QualifiedE_Address_Status, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_AftTypes_EncapsulationHeaderType is a E_AftTypes_EncapsulationHeaderType with a corresponding timestamp.
type QualifiedE_AftTypes_EncapsulationHeaderType struct {
	*genutil.Metadata
	val     E_AftTypes_EncapsulationHeaderType // val is the sample value.
	present bool
}

func (q *QualifiedE_AftTypes_EncapsulationHeaderType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_AftTypes_EncapsulationHeaderType sample, erroring out if not present.
func (q *QualifiedE_AftTypes_EncapsulationHeaderType) Val(t testing.TB) E_AftTypes_EncapsulationHeaderType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_AftTypes_EncapsulationHeaderType sample.
func (q *QualifiedE_AftTypes_EncapsulationHeaderType) SetVal(v E_AftTypes_EncapsulationHeaderType) *QualifiedE_AftTypes_EncapsulationHeaderType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_AftTypes_EncapsulationHeaderType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_AftTypes_EncapsulationHeaderType is a telemetry Collection whose Await method returns a slice of E_AftTypes_EncapsulationHeaderType samples.
type CollectionE_AftTypes_EncapsulationHeaderType struct {
	W    *E_AftTypes_EncapsulationHeaderTypeWatcher
	Data []*QualifiedE_AftTypes_EncapsulationHeaderType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_AftTypes_EncapsulationHeaderType) Await(t testing.TB) []*QualifiedE_AftTypes_EncapsulationHeaderType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_AftTypes_EncapsulationHeaderTypeWatcher observes a stream of E_AftTypes_EncapsulationHeaderType samples.
type E_AftTypes_EncapsulationHeaderTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_AftTypes_EncapsulationHeaderType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_AftTypes_EncapsulationHeaderTypeWatcher) Await(t testing.TB) (*QualifiedE_AftTypes_EncapsulationHeaderType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY is a E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY with a corresponding timestamp.
type QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY struct {
	*genutil.Metadata
	val     E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY // val is the sample value.
	present bool
}

func (q *QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY sample, erroring out if not present.
func (q *QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) Val(t testing.TB) E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY sample.
func (q *QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) SetVal(v E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) *QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY is a telemetry Collection whose Await method returns a slice of E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY samples.
type CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY struct {
	W    *E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher
	Data []*QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY) Await(t testing.TB) []*QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher observes a stream of E_AlarmTypes_OPENCONFIG_ALARM_SEVERITY samples.
type E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_AlarmTypes_OPENCONFIG_ALARM_SEVERITYWatcher) Await(t testing.TB) (*QualifiedE_AlarmTypes_OPENCONFIG_ALARM_SEVERITY, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_AsExternalLsa_MetricType is a E_AsExternalLsa_MetricType with a corresponding timestamp.
type QualifiedE_AsExternalLsa_MetricType struct {
	*genutil.Metadata
	val     E_AsExternalLsa_MetricType // val is the sample value.
	present bool
}

func (q *QualifiedE_AsExternalLsa_MetricType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_AsExternalLsa_MetricType sample, erroring out if not present.
func (q *QualifiedE_AsExternalLsa_MetricType) Val(t testing.TB) E_AsExternalLsa_MetricType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_AsExternalLsa_MetricType sample.
func (q *QualifiedE_AsExternalLsa_MetricType) SetVal(v E_AsExternalLsa_MetricType) *QualifiedE_AsExternalLsa_MetricType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_AsExternalLsa_MetricType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_AsExternalLsa_MetricType is a telemetry Collection whose Await method returns a slice of E_AsExternalLsa_MetricType samples.
type CollectionE_AsExternalLsa_MetricType struct {
	W    *E_AsExternalLsa_MetricTypeWatcher
	Data []*QualifiedE_AsExternalLsa_MetricType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_AsExternalLsa_MetricType) Await(t testing.TB) []*QualifiedE_AsExternalLsa_MetricType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_AsExternalLsa_MetricTypeWatcher observes a stream of E_AsExternalLsa_MetricType samples.
type E_AsExternalLsa_MetricTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_AsExternalLsa_MetricType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_AsExternalLsa_MetricTypeWatcher) Await(t testing.TB) (*QualifiedE_AsExternalLsa_MetricType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Authentication_CryptoType is a E_Authentication_CryptoType with a corresponding timestamp.
type QualifiedE_Authentication_CryptoType struct {
	*genutil.Metadata
	val     E_Authentication_CryptoType // val is the sample value.
	present bool
}

func (q *QualifiedE_Authentication_CryptoType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Authentication_CryptoType sample, erroring out if not present.
func (q *QualifiedE_Authentication_CryptoType) Val(t testing.TB) E_Authentication_CryptoType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Authentication_CryptoType sample.
func (q *QualifiedE_Authentication_CryptoType) SetVal(v E_Authentication_CryptoType) *QualifiedE_Authentication_CryptoType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Authentication_CryptoType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Authentication_CryptoType is a telemetry Collection whose Await method returns a slice of E_Authentication_CryptoType samples.
type CollectionE_Authentication_CryptoType struct {
	W    *E_Authentication_CryptoTypeWatcher
	Data []*QualifiedE_Authentication_CryptoType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Authentication_CryptoType) Await(t testing.TB) []*QualifiedE_Authentication_CryptoType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Authentication_CryptoTypeWatcher observes a stream of E_Authentication_CryptoType samples.
type E_Authentication_CryptoTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Authentication_CryptoType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Authentication_CryptoTypeWatcher) Await(t testing.TB) (*QualifiedE_Authentication_CryptoType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpConditions_RouteType is a E_BgpConditions_RouteType with a corresponding timestamp.
type QualifiedE_BgpConditions_RouteType struct {
	*genutil.Metadata
	val     E_BgpConditions_RouteType // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpConditions_RouteType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpConditions_RouteType sample, erroring out if not present.
func (q *QualifiedE_BgpConditions_RouteType) Val(t testing.TB) E_BgpConditions_RouteType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpConditions_RouteType sample.
func (q *QualifiedE_BgpConditions_RouteType) SetVal(v E_BgpConditions_RouteType) *QualifiedE_BgpConditions_RouteType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpConditions_RouteType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpConditions_RouteType is a telemetry Collection whose Await method returns a slice of E_BgpConditions_RouteType samples.
type CollectionE_BgpConditions_RouteType struct {
	W    *E_BgpConditions_RouteTypeWatcher
	Data []*QualifiedE_BgpConditions_RouteType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpConditions_RouteType) Await(t testing.TB) []*QualifiedE_BgpConditions_RouteType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpConditions_RouteTypeWatcher observes a stream of E_BgpConditions_RouteType samples.
type E_BgpConditions_RouteTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpConditions_RouteType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpConditions_RouteTypeWatcher) Await(t testing.TB) (*QualifiedE_BgpConditions_RouteType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpPolicy_BgpSetCommunityOptionType is a E_BgpPolicy_BgpSetCommunityOptionType with a corresponding timestamp.
type QualifiedE_BgpPolicy_BgpSetCommunityOptionType struct {
	*genutil.Metadata
	val     E_BgpPolicy_BgpSetCommunityOptionType // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpPolicy_BgpSetCommunityOptionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpPolicy_BgpSetCommunityOptionType sample, erroring out if not present.
func (q *QualifiedE_BgpPolicy_BgpSetCommunityOptionType) Val(t testing.TB) E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpPolicy_BgpSetCommunityOptionType sample.
func (q *QualifiedE_BgpPolicy_BgpSetCommunityOptionType) SetVal(v E_BgpPolicy_BgpSetCommunityOptionType) *QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpPolicy_BgpSetCommunityOptionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpPolicy_BgpSetCommunityOptionType is a telemetry Collection whose Await method returns a slice of E_BgpPolicy_BgpSetCommunityOptionType samples.
type CollectionE_BgpPolicy_BgpSetCommunityOptionType struct {
	W    *E_BgpPolicy_BgpSetCommunityOptionTypeWatcher
	Data []*QualifiedE_BgpPolicy_BgpSetCommunityOptionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpPolicy_BgpSetCommunityOptionType) Await(t testing.TB) []*QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpPolicy_BgpSetCommunityOptionTypeWatcher observes a stream of E_BgpPolicy_BgpSetCommunityOptionType samples.
type E_BgpPolicy_BgpSetCommunityOptionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpPolicy_BgpSetCommunityOptionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpPolicy_BgpSetCommunityOptionTypeWatcher) Await(t testing.TB) (*QualifiedE_BgpPolicy_BgpSetCommunityOptionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_AFI_SAFI_TYPE is a E_BgpTypes_AFI_SAFI_TYPE with a corresponding timestamp.
type QualifiedE_BgpTypes_AFI_SAFI_TYPE struct {
	*genutil.Metadata
	val     E_BgpTypes_AFI_SAFI_TYPE // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_AFI_SAFI_TYPE sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPE) Val(t testing.TB) E_BgpTypes_AFI_SAFI_TYPE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_AFI_SAFI_TYPE sample.
func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPE) SetVal(v E_BgpTypes_AFI_SAFI_TYPE) *QualifiedE_BgpTypes_AFI_SAFI_TYPE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_AFI_SAFI_TYPE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_AFI_SAFI_TYPE is a telemetry Collection whose Await method returns a slice of E_BgpTypes_AFI_SAFI_TYPE samples.
type CollectionE_BgpTypes_AFI_SAFI_TYPE struct {
	W    *E_BgpTypes_AFI_SAFI_TYPEWatcher
	Data []*QualifiedE_BgpTypes_AFI_SAFI_TYPE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_AFI_SAFI_TYPE) Await(t testing.TB) []*QualifiedE_BgpTypes_AFI_SAFI_TYPE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_AFI_SAFI_TYPEWatcher observes a stream of E_BgpTypes_AFI_SAFI_TYPE samples.
type E_BgpTypes_AFI_SAFI_TYPEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_AFI_SAFI_TYPE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_AFI_SAFI_TYPEWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_AFI_SAFI_TYPE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_AsPathSegmentType is a E_BgpTypes_AsPathSegmentType with a corresponding timestamp.
type QualifiedE_BgpTypes_AsPathSegmentType struct {
	*genutil.Metadata
	val     E_BgpTypes_AsPathSegmentType // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_AsPathSegmentType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_AsPathSegmentType sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_AsPathSegmentType) Val(t testing.TB) E_BgpTypes_AsPathSegmentType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_AsPathSegmentType sample.
func (q *QualifiedE_BgpTypes_AsPathSegmentType) SetVal(v E_BgpTypes_AsPathSegmentType) *QualifiedE_BgpTypes_AsPathSegmentType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_AsPathSegmentType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_AsPathSegmentType is a telemetry Collection whose Await method returns a slice of E_BgpTypes_AsPathSegmentType samples.
type CollectionE_BgpTypes_AsPathSegmentType struct {
	W    *E_BgpTypes_AsPathSegmentTypeWatcher
	Data []*QualifiedE_BgpTypes_AsPathSegmentType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_AsPathSegmentType) Await(t testing.TB) []*QualifiedE_BgpTypes_AsPathSegmentType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_AsPathSegmentTypeWatcher observes a stream of E_BgpTypes_AsPathSegmentType samples.
type E_BgpTypes_AsPathSegmentTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_AsPathSegmentType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_AsPathSegmentTypeWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_AsPathSegmentType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_BGP_ERROR_CODE is a E_BgpTypes_BGP_ERROR_CODE with a corresponding timestamp.
type QualifiedE_BgpTypes_BGP_ERROR_CODE struct {
	*genutil.Metadata
	val     E_BgpTypes_BGP_ERROR_CODE // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_BGP_ERROR_CODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_BGP_ERROR_CODE sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_BGP_ERROR_CODE) Val(t testing.TB) E_BgpTypes_BGP_ERROR_CODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_BGP_ERROR_CODE sample.
func (q *QualifiedE_BgpTypes_BGP_ERROR_CODE) SetVal(v E_BgpTypes_BGP_ERROR_CODE) *QualifiedE_BgpTypes_BGP_ERROR_CODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_BGP_ERROR_CODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_BGP_ERROR_CODE is a telemetry Collection whose Await method returns a slice of E_BgpTypes_BGP_ERROR_CODE samples.
type CollectionE_BgpTypes_BGP_ERROR_CODE struct {
	W    *E_BgpTypes_BGP_ERROR_CODEWatcher
	Data []*QualifiedE_BgpTypes_BGP_ERROR_CODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_BGP_ERROR_CODE) Await(t testing.TB) []*QualifiedE_BgpTypes_BGP_ERROR_CODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_BGP_ERROR_CODEWatcher observes a stream of E_BgpTypes_BGP_ERROR_CODE samples.
type E_BgpTypes_BGP_ERROR_CODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_BGP_ERROR_CODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_BGP_ERROR_CODEWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_BGP_ERROR_CODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_BGP_ERROR_SUBCODE is a E_BgpTypes_BGP_ERROR_SUBCODE with a corresponding timestamp.
type QualifiedE_BgpTypes_BGP_ERROR_SUBCODE struct {
	*genutil.Metadata
	val     E_BgpTypes_BGP_ERROR_SUBCODE // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_BGP_ERROR_SUBCODE) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_BGP_ERROR_SUBCODE sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_BGP_ERROR_SUBCODE) Val(t testing.TB) E_BgpTypes_BGP_ERROR_SUBCODE {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_BGP_ERROR_SUBCODE sample.
func (q *QualifiedE_BgpTypes_BGP_ERROR_SUBCODE) SetVal(v E_BgpTypes_BGP_ERROR_SUBCODE) *QualifiedE_BgpTypes_BGP_ERROR_SUBCODE {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_BGP_ERROR_SUBCODE) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_BGP_ERROR_SUBCODE is a telemetry Collection whose Await method returns a slice of E_BgpTypes_BGP_ERROR_SUBCODE samples.
type CollectionE_BgpTypes_BGP_ERROR_SUBCODE struct {
	W    *E_BgpTypes_BGP_ERROR_SUBCODEWatcher
	Data []*QualifiedE_BgpTypes_BGP_ERROR_SUBCODE
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_BGP_ERROR_SUBCODE) Await(t testing.TB) []*QualifiedE_BgpTypes_BGP_ERROR_SUBCODE {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_BGP_ERROR_SUBCODEWatcher observes a stream of E_BgpTypes_BGP_ERROR_SUBCODE samples.
type E_BgpTypes_BGP_ERROR_SUBCODEWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_BGP_ERROR_SUBCODE
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_BGP_ERROR_SUBCODEWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_BGP_ERROR_SUBCODE, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_BgpOriginAttrType is a E_BgpTypes_BgpOriginAttrType with a corresponding timestamp.
type QualifiedE_BgpTypes_BgpOriginAttrType struct {
	*genutil.Metadata
	val     E_BgpTypes_BgpOriginAttrType // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_BgpOriginAttrType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_BgpOriginAttrType sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_BgpOriginAttrType) Val(t testing.TB) E_BgpTypes_BgpOriginAttrType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_BgpOriginAttrType sample.
func (q *QualifiedE_BgpTypes_BgpOriginAttrType) SetVal(v E_BgpTypes_BgpOriginAttrType) *QualifiedE_BgpTypes_BgpOriginAttrType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_BgpOriginAttrType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_BgpOriginAttrType is a telemetry Collection whose Await method returns a slice of E_BgpTypes_BgpOriginAttrType samples.
type CollectionE_BgpTypes_BgpOriginAttrType struct {
	W    *E_BgpTypes_BgpOriginAttrTypeWatcher
	Data []*QualifiedE_BgpTypes_BgpOriginAttrType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_BgpOriginAttrType) Await(t testing.TB) []*QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_BgpOriginAttrTypeWatcher observes a stream of E_BgpTypes_BgpOriginAttrType samples.
type E_BgpTypes_BgpOriginAttrTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_BgpOriginAttrType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_BgpOriginAttrTypeWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_BgpOriginAttrType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_CommunityType is a E_BgpTypes_CommunityType with a corresponding timestamp.
type QualifiedE_BgpTypes_CommunityType struct {
	*genutil.Metadata
	val     E_BgpTypes_CommunityType // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_CommunityType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_CommunityType sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_CommunityType) Val(t testing.TB) E_BgpTypes_CommunityType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_CommunityType sample.
func (q *QualifiedE_BgpTypes_CommunityType) SetVal(v E_BgpTypes_CommunityType) *QualifiedE_BgpTypes_CommunityType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_CommunityType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_CommunityType is a telemetry Collection whose Await method returns a slice of E_BgpTypes_CommunityType samples.
type CollectionE_BgpTypes_CommunityType struct {
	W    *E_BgpTypes_CommunityTypeWatcher
	Data []*QualifiedE_BgpTypes_CommunityType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_CommunityType) Await(t testing.TB) []*QualifiedE_BgpTypes_CommunityType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_CommunityTypeWatcher observes a stream of E_BgpTypes_CommunityType samples.
type E_BgpTypes_CommunityTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_CommunityType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_CommunityTypeWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_CommunityType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_PeerType is a E_BgpTypes_PeerType with a corresponding timestamp.
type QualifiedE_BgpTypes_PeerType struct {
	*genutil.Metadata
	val     E_BgpTypes_PeerType // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_PeerType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_PeerType sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_PeerType) Val(t testing.TB) E_BgpTypes_PeerType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_PeerType sample.
func (q *QualifiedE_BgpTypes_PeerType) SetVal(v E_BgpTypes_PeerType) *QualifiedE_BgpTypes_PeerType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_PeerType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_PeerType is a telemetry Collection whose Await method returns a slice of E_BgpTypes_PeerType samples.
type CollectionE_BgpTypes_PeerType struct {
	W    *E_BgpTypes_PeerTypeWatcher
	Data []*QualifiedE_BgpTypes_PeerType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_PeerType) Await(t testing.TB) []*QualifiedE_BgpTypes_PeerType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_PeerTypeWatcher observes a stream of E_BgpTypes_PeerType samples.
type E_BgpTypes_PeerTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_PeerType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_PeerTypeWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_PeerType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_BgpTypes_RemovePrivateAsOption is a E_BgpTypes_RemovePrivateAsOption with a corresponding timestamp.
type QualifiedE_BgpTypes_RemovePrivateAsOption struct {
	*genutil.Metadata
	val     E_BgpTypes_RemovePrivateAsOption // val is the sample value.
	present bool
}

func (q *QualifiedE_BgpTypes_RemovePrivateAsOption) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_BgpTypes_RemovePrivateAsOption sample, erroring out if not present.
func (q *QualifiedE_BgpTypes_RemovePrivateAsOption) Val(t testing.TB) E_BgpTypes_RemovePrivateAsOption {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_BgpTypes_RemovePrivateAsOption sample.
func (q *QualifiedE_BgpTypes_RemovePrivateAsOption) SetVal(v E_BgpTypes_RemovePrivateAsOption) *QualifiedE_BgpTypes_RemovePrivateAsOption {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_BgpTypes_RemovePrivateAsOption) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_BgpTypes_RemovePrivateAsOption is a telemetry Collection whose Await method returns a slice of E_BgpTypes_RemovePrivateAsOption samples.
type CollectionE_BgpTypes_RemovePrivateAsOption struct {
	W    *E_BgpTypes_RemovePrivateAsOptionWatcher
	Data []*QualifiedE_BgpTypes_RemovePrivateAsOption
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_BgpTypes_RemovePrivateAsOption) Await(t testing.TB) []*QualifiedE_BgpTypes_RemovePrivateAsOption {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_BgpTypes_RemovePrivateAsOptionWatcher observes a stream of E_BgpTypes_RemovePrivateAsOption samples.
type E_BgpTypes_RemovePrivateAsOptionWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_BgpTypes_RemovePrivateAsOption
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_BgpTypes_RemovePrivateAsOptionWatcher) Await(t testing.TB) (*QualifiedE_BgpTypes_RemovePrivateAsOption, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_Bgp_Neighbor_SessionState is a E_Bgp_Neighbor_SessionState with a corresponding timestamp.
type QualifiedE_Bgp_Neighbor_SessionState struct {
	*genutil.Metadata
	val     E_Bgp_Neighbor_SessionState // val is the sample value.
	present bool
}

func (q *QualifiedE_Bgp_Neighbor_SessionState) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_Bgp_Neighbor_SessionState sample, erroring out if not present.
func (q *QualifiedE_Bgp_Neighbor_SessionState) Val(t testing.TB) E_Bgp_Neighbor_SessionState {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_Bgp_Neighbor_SessionState sample.
func (q *QualifiedE_Bgp_Neighbor_SessionState) SetVal(v E_Bgp_Neighbor_SessionState) *QualifiedE_Bgp_Neighbor_SessionState {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_Bgp_Neighbor_SessionState) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_Bgp_Neighbor_SessionState is a telemetry Collection whose Await method returns a slice of E_Bgp_Neighbor_SessionState samples.
type CollectionE_Bgp_Neighbor_SessionState struct {
	W    *E_Bgp_Neighbor_SessionStateWatcher
	Data []*QualifiedE_Bgp_Neighbor_SessionState
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_Bgp_Neighbor_SessionState) Await(t testing.TB) []*QualifiedE_Bgp_Neighbor_SessionState {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_Bgp_Neighbor_SessionStateWatcher observes a stream of E_Bgp_Neighbor_SessionState samples.
type E_Bgp_Neighbor_SessionStateWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_Bgp_Neighbor_SessionState
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_Bgp_Neighbor_SessionStateWatcher) Await(t testing.TB) (*QualifiedE_Bgp_Neighbor_SessionState, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_DefaultMetric_Flags is a E_DefaultMetric_Flags with a corresponding timestamp.
type QualifiedE_DefaultMetric_Flags struct {
	*genutil.Metadata
	val     E_DefaultMetric_Flags // val is the sample value.
	present bool
}

func (q *QualifiedE_DefaultMetric_Flags) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_DefaultMetric_Flags sample, erroring out if not present.
func (q *QualifiedE_DefaultMetric_Flags) Val(t testing.TB) E_DefaultMetric_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_DefaultMetric_Flags sample.
func (q *QualifiedE_DefaultMetric_Flags) SetVal(v E_DefaultMetric_Flags) *QualifiedE_DefaultMetric_Flags {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_DefaultMetric_Flags) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_DefaultMetric_Flags is a telemetry Collection whose Await method returns a slice of E_DefaultMetric_Flags samples.
type CollectionE_DefaultMetric_Flags struct {
	W    *E_DefaultMetric_FlagsWatcher
	Data []*QualifiedE_DefaultMetric_Flags
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_DefaultMetric_Flags) Await(t testing.TB) []*QualifiedE_DefaultMetric_Flags {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_DefaultMetric_FlagsWatcher observes a stream of E_DefaultMetric_Flags samples.
type E_DefaultMetric_FlagsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_DefaultMetric_Flags
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_DefaultMetric_FlagsWatcher) Await(t testing.TB) (*QualifiedE_DefaultMetric_Flags, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedE_DfElection_DfElectionMethod is a E_DfElection_DfElectionMethod with a corresponding timestamp.
type QualifiedE_DfElection_DfElectionMethod struct {
	*genutil.Metadata
	val     E_DfElection_DfElectionMethod // val is the sample value.
	present bool
}

func (q *QualifiedE_DfElection_DfElectionMethod) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the E_DfElection_DfElectionMethod sample, erroring out if not present.
func (q *QualifiedE_DfElection_DfElectionMethod) Val(t testing.TB) E_DfElection_DfElectionMethod {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the E_DfElection_DfElectionMethod sample.
func (q *QualifiedE_DfElection_DfElectionMethod) SetVal(v E_DfElection_DfElectionMethod) *QualifiedE_DfElection_DfElectionMethod {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedE_DfElection_DfElectionMethod) IsPresent() bool {
	return q != nil && q.present
}

// CollectionE_DfElection_DfElectionMethod is a telemetry Collection whose Await method returns a slice of E_DfElection_DfElectionMethod samples.
type CollectionE_DfElection_DfElectionMethod struct {
	W    *E_DfElection_DfElectionMethodWatcher
	Data []*QualifiedE_DfElection_DfElectionMethod
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionE_DfElection_DfElectionMethod) Await(t testing.TB) []*QualifiedE_DfElection_DfElectionMethod {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// E_DfElection_DfElectionMethodWatcher observes a stream of E_DfElection_DfElectionMethod samples.
type E_DfElection_DfElectionMethodWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedE_DfElection_DfElectionMethod
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *E_DfElection_DfElectionMethodWatcher) Await(t testing.TB) (*QualifiedE_DfElection_DfElectionMethod, bool) {
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
