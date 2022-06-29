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

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy is a *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions is a *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop is a *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd is a *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling is a *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart is a *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions is a *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector is a *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers is a *NetworkInstance_Protocol_Bgp_PeerGroup_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_Timers {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_Timers) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_Timers samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Timers struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_Timers samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport is a *NetworkInstance_Protocol_Bgp_PeerGroup_Transport with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_Transport // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Transport sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_Transport {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Transport sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_Transport) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Transport is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_Transport samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Transport struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Transport) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_Transport samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib is a *NetworkInstance_Protocol_Bgp_Rib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib) *QualifiedNetworkInstance_Protocol_Bgp_Rib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib struct {
	W    *NetworkInstance_Protocol_Bgp_RibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_RibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib samples.
type NetworkInstance_Protocol_Bgp_RibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_RibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpnWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpnWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpnWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpnWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisherWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisherWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisherWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisherWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_PathWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_PathWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_PathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_PathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFiveRoute_Path_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_PathWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_PathWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_PathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_PathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeFourRoute_Path_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_PathWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_PathWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_PathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_PathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeOneRoute_Path_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_PathWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_PathWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_PathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_PathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeThreeRoute_Path_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_PathWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_PathWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_PathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_PathWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_LocRib_RouteDistinguisher_TypeTwoRoute_Path_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_L2VpnEvpn_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
