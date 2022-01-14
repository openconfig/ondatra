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

// QualifiedNetworkInstance_Protocol_Isis_Global_Timers is a *NetworkInstance_Protocol_Isis_Global_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Global_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Global_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Global_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Global_Timers {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Global_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers) SetVal(v *NetworkInstance_Protocol_Isis_Global_Timers) *QualifiedNetworkInstance_Protocol_Isis_Global_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Global_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Global_Timers samples.
type CollectionNetworkInstance_Protocol_Isis_Global_Timers struct {
	W    *NetworkInstance_Protocol_Isis_Global_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Global_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Global_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Global_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Global_TimersWatcher observes a stream of *NetworkInstance_Protocol_Isis_Global_Timers samples.
type NetworkInstance_Protocol_Isis_Global_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Global_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Global_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Global_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration is a *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration) SetVal(v *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration) *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration samples.
type CollectionNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration struct {
	W    *NetworkInstance_Protocol_Isis_Global_Timers_LspGenerationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Global_Timers_LspGenerationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Global_Timers_LspGeneration samples.
type NetworkInstance_Protocol_Isis_Global_Timers_LspGenerationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Global_Timers_LspGenerationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Global_Timers_LspGeneration, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf is a *NetworkInstance_Protocol_Isis_Global_Timers_Spf with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Global_Timers_Spf // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Global_Timers_Spf sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Global_Timers_Spf {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Global_Timers_Spf sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf) SetVal(v *NetworkInstance_Protocol_Isis_Global_Timers_Spf) *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Global_Timers_Spf is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Global_Timers_Spf samples.
type CollectionNetworkInstance_Protocol_Isis_Global_Timers_Spf struct {
	W    *NetworkInstance_Protocol_Isis_Global_Timers_SpfWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Global_Timers_Spf) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Global_Timers_SpfWatcher observes a stream of *NetworkInstance_Protocol_Isis_Global_Timers_Spf samples.
type NetworkInstance_Protocol_Isis_Global_Timers_SpfWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Global_Timers_SpfWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Global_Timers_Spf, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Global_Transport is a *NetworkInstance_Protocol_Isis_Global_Transport with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Global_Transport struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Global_Transport // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Transport) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Global_Transport sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Transport) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Global_Transport {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Global_Transport sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Transport) SetVal(v *NetworkInstance_Protocol_Isis_Global_Transport) *QualifiedNetworkInstance_Protocol_Isis_Global_Transport {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Global_Transport) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Global_Transport is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Global_Transport samples.
type CollectionNetworkInstance_Protocol_Isis_Global_Transport struct {
	W    *NetworkInstance_Protocol_Isis_Global_TransportWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Global_Transport
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Global_Transport) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Global_Transport {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Global_TransportWatcher observes a stream of *NetworkInstance_Protocol_Isis_Global_Transport samples.
type NetworkInstance_Protocol_Isis_Global_TransportWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Global_Transport
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Global_TransportWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Global_Transport, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface is a *NetworkInstance_Protocol_Isis_Interface with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface) SetVal(v *NetworkInstance_Protocol_Isis_Interface) *QualifiedNetworkInstance_Protocol_Isis_Interface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface samples.
type CollectionNetworkInstance_Protocol_Isis_Interface struct {
	W    *NetworkInstance_Protocol_Isis_InterfaceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_InterfaceWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface samples.
type NetworkInstance_Protocol_Isis_InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_InterfaceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Af is a *NetworkInstance_Protocol_Isis_Interface_Af with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Af struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Af // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Af) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Af sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Af) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Af {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Af sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Af) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Af) *QualifiedNetworkInstance_Protocol_Isis_Interface_Af {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Af) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Af is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Af samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Af struct {
	W    *NetworkInstance_Protocol_Isis_Interface_AfWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Af
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Af) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Af {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_AfWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Af samples.
type NetworkInstance_Protocol_Isis_Interface_AfWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Af
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_AfWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Af, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication is a *NetworkInstance_Protocol_Isis_Interface_Authentication with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Authentication // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Authentication sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Authentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Authentication sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Authentication) *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Authentication is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Authentication samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Authentication struct {
	W    *NetworkInstance_Protocol_Isis_Interface_AuthenticationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Authentication) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_AuthenticationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Authentication samples.
type NetworkInstance_Protocol_Isis_Interface_AuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_AuthenticationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key is a *NetworkInstance_Protocol_Isis_Interface_Authentication_Key with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Authentication_Key // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Authentication_Key sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Authentication_Key {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Authentication_Key sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Authentication_Key) *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Authentication_Key is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Authentication_Key samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Authentication_Key struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Authentication_KeyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Authentication_Key) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Authentication_KeyWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Authentication_Key samples.
type NetworkInstance_Protocol_Isis_Interface_Authentication_KeyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Authentication_KeyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Key, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain is a *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain) *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Authentication_KeychainWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Authentication_KeychainWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Authentication_Keychain samples.
type NetworkInstance_Protocol_Isis_Interface_Authentication_KeychainWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Authentication_KeychainWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Authentication_Keychain, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd is a *NetworkInstance_Protocol_Isis_Interface_Bfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Bfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Bfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Bfd {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Bfd sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Bfd) *QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Bfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Bfd samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Bfd struct {
	W    *NetworkInstance_Protocol_Isis_Interface_BfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Bfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_BfdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Bfd samples.
type NetworkInstance_Protocol_Isis_Interface_BfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_BfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Bfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters is a *NetworkInstance_Protocol_Isis_Interface_CircuitCounters with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_CircuitCounters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_CircuitCounters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_CircuitCounters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_CircuitCounters sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters) SetVal(v *NetworkInstance_Protocol_Isis_Interface_CircuitCounters) *QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_CircuitCounters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_CircuitCounters samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_CircuitCounters struct {
	W    *NetworkInstance_Protocol_Isis_Interface_CircuitCountersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_CircuitCounters) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_CircuitCountersWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_CircuitCounters samples.
type NetworkInstance_Protocol_Isis_Interface_CircuitCountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_CircuitCountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_CircuitCounters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd is a *NetworkInstance_Protocol_Isis_Interface_EnableBfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_EnableBfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_EnableBfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_EnableBfd {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_EnableBfd sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd) SetVal(v *NetworkInstance_Protocol_Isis_Interface_EnableBfd) *QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_EnableBfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_EnableBfd samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_EnableBfd struct {
	W    *NetworkInstance_Protocol_Isis_Interface_EnableBfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_EnableBfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_EnableBfdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_EnableBfd samples.
type NetworkInstance_Protocol_Isis_Interface_EnableBfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_EnableBfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_EnableBfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef is a *NetworkInstance_Protocol_Isis_Interface_InterfaceRef with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_InterfaceRef // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_InterfaceRef sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_InterfaceRef {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_InterfaceRef sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef) SetVal(v *NetworkInstance_Protocol_Isis_Interface_InterfaceRef) *QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_InterfaceRef is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_InterfaceRef samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_InterfaceRef struct {
	W    *NetworkInstance_Protocol_Isis_Interface_InterfaceRefWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_InterfaceRef) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_InterfaceRefWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_InterfaceRef samples.
type NetworkInstance_Protocol_Isis_Interface_InterfaceRefWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_InterfaceRefWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_InterfaceRef, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level is a *NetworkInstance_Protocol_Isis_Interface_Level with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level struct {
	W    *NetworkInstance_Protocol_Isis_Interface_LevelWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_LevelWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level samples.
type NetworkInstance_Protocol_Isis_Interface_LevelWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_LevelWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency is a *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Adjacency is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Adjacency struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_AdjacencyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Adjacency) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_AdjacencyWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_Adjacency samples.
type NetworkInstance_Protocol_Isis_Interface_Level_AdjacencyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_AdjacencyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Adjacency, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af is a *NetworkInstance_Protocol_Isis_Interface_Level_Af with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_Af // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_Af {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_Af) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_Af samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_AfWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_AfWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_Af samples.
type NetworkInstance_Protocol_Isis_Interface_Level_AfWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_AfWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting is a *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRoutingWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRoutingWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting samples.
type NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRoutingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRoutingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid is a *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid samples.
type NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_AdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid is a *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid samples.
type NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Af_SegmentRouting_PrefixSid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication is a *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthenticationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthenticationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication samples.
type NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthenticationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key is a *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeyWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key samples.
type NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Key, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain is a *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeychainWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeychainWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain samples.
type NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeychainWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_KeychainWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_HelloAuthentication_Keychain, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCountersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCountersWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCountersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_CsnpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_CsnpWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_CsnpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_CsnpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Csnp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_EshWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_EshWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_EshWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_EshWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Esh, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IihWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IihWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IihWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IihWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Iih, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IshWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IshWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IshWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_IshWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Ish, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_LspWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_LspWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_LspWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_LspWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Lsp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_PsnpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_PsnpWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_PsnpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_PsnpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Psnp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown is a *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_UnknownWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_UnknownWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown samples.
type NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_UnknownWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_UnknownWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_PacketCounters_Unknown, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers is a *NetworkInstance_Protocol_Isis_Interface_Level_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Level_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Level_Timers {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Level_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Level_Timers) *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Level_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Level_Timers samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Level_Timers struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Level_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Level_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Level_TimersWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Level_Timers samples.
type NetworkInstance_Protocol_Isis_Interface_Level_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Level_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Level_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls is a *NetworkInstance_Protocol_Isis_Interface_Mpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Mpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Mpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Mpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Mpls sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Mpls) *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Mpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Mpls samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Mpls struct {
	W    *NetworkInstance_Protocol_Isis_Interface_MplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Mpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_MplsWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Mpls samples.
type NetworkInstance_Protocol_Isis_Interface_MplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_MplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync is a *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync) *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync struct {
	W    *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSyncWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSyncWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync samples.
type NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSyncWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSyncWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Mpls_IgpLdpSync, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Interface_Timers is a *NetworkInstance_Protocol_Isis_Interface_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Interface_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Interface_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Interface_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Interface_Timers {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Interface_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Timers) SetVal(v *NetworkInstance_Protocol_Isis_Interface_Timers) *QualifiedNetworkInstance_Protocol_Isis_Interface_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Interface_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Interface_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Interface_Timers samples.
type CollectionNetworkInstance_Protocol_Isis_Interface_Timers struct {
	W    *NetworkInstance_Protocol_Isis_Interface_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Interface_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Interface_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Interface_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Interface_TimersWatcher observes a stream of *NetworkInstance_Protocol_Isis_Interface_Timers samples.
type NetworkInstance_Protocol_Isis_Interface_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Interface_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Interface_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Interface_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level is a *NetworkInstance_Protocol_Isis_Level with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level) SetVal(v *NetworkInstance_Protocol_Isis_Level) *QualifiedNetworkInstance_Protocol_Isis_Level {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level samples.
type CollectionNetworkInstance_Protocol_Isis_Level struct {
	W    *NetworkInstance_Protocol_Isis_LevelWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_LevelWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level samples.
type NetworkInstance_Protocol_Isis_LevelWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_LevelWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Authentication is a *NetworkInstance_Protocol_Isis_Level_Authentication with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Authentication struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Authentication // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Authentication sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Authentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Authentication sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication) SetVal(v *NetworkInstance_Protocol_Isis_Level_Authentication) *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Authentication is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Authentication samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Authentication struct {
	W    *NetworkInstance_Protocol_Isis_Level_AuthenticationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Authentication) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_AuthenticationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Authentication samples.
type NetworkInstance_Protocol_Isis_Level_AuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_AuthenticationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key is a *NetworkInstance_Protocol_Isis_Level_Authentication_Key with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Authentication_Key // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Authentication_Key sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Authentication_Key {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Authentication_Key sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key) SetVal(v *NetworkInstance_Protocol_Isis_Level_Authentication_Key) *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Authentication_Key is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Authentication_Key samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Authentication_Key struct {
	W    *NetworkInstance_Protocol_Isis_Level_Authentication_KeyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Authentication_Key) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Authentication_KeyWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Authentication_Key samples.
type NetworkInstance_Protocol_Isis_Level_Authentication_KeyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Authentication_KeyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Key, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain is a *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain) SetVal(v *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain) *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Authentication_Keychain is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Authentication_Keychain struct {
	W    *NetworkInstance_Protocol_Isis_Level_Authentication_KeychainWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Authentication_Keychain) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Authentication_KeychainWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Authentication_Keychain samples.
type NetworkInstance_Protocol_Isis_Level_Authentication_KeychainWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Authentication_KeychainWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Authentication_Keychain, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp is a *NetworkInstance_Protocol_Isis_Level_Lsp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp struct {
	W    *NetworkInstance_Protocol_Isis_Level_LspWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_LspWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp samples.
type NetworkInstance_Protocol_Isis_Level_LspWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_LspWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_TlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_TlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_TlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_TlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AreaAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AuthenticationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AuthenticationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AuthenticationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_AuthenticationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Authentication, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_CapabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_CapabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_CapabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_CapabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithmsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithmsWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithmsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithmsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingAlgorithms, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptorWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptorWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptorWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_Subtlv_SegmentRoutingCapability_SrgbDescriptor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Capability_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4ReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4ReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4ReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4ReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_PrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_PrefixWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_PrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_FlagsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_FlagsWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_FlagsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_FlagsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Flags, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv4SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Ipv6SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_PrefixSid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_TagWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_TagWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_TagWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_TagWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64Watcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64Watcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64 samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_Subtlv_Tag64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIpv4Reachability_Prefix_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_InstanceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_InstanceWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_InstanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_InstanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_AvailableBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_ConstraintWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_BandwidthConstraint_Constraint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ExtendedAdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv4NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6InterfaceAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddressWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_Ipv6NeighborAddress, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LanAdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributesWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkDelayVariation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLossWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLossWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLossWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLossWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkLoss, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionTypeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_LinkProtectionType, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MaxReservableLinkBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelayWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_MinMaxLinkDelay, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_ResidualBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriorityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriorityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriorityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriorityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_SetupPriority, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_TeDefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLspWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UnconstrainedLsp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidthWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_Subtlv_UtilizedBandwidth, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_ExtendedIsReachability_Neighbor_Instance_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_HostnameWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_HostnameWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_HostnameWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_HostnameWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Hostname, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_InstanceId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_PrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_PrefixWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_PrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_DelayMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ErrorMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4ExternalReachability_Prefix_ExpenseMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddressesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddressesWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddressesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddressesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InterfaceAddresses, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_PrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_PrefixWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_PrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_DelayMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ErrorMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4InternalReachability_Prefix_ExpenseMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4SrlgWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4Srlg, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv4TeRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddressesWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6InterfaceAddresses, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6ReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_PrefixWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_FlagsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Flags, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv4SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Ipv6SourceRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_PrefixSid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_TagWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64 samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_Subtlv_Tag64, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Reachability_Prefix_UndefinedSubtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6SrlgWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6Srlg, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_Ipv6TeRouterId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasIdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsAliasId, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachabilityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DefaultMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_DelayMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ErrorMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetricWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsReachability_Neighbor_ExpenseMetric, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_InstanceWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySidWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdjacencySid, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup is a *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) Val(t testing.TB) *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup sample.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) SetVal(v *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup samples.
type CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup struct {
	W    *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher
	Data []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher observes a stream of *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup samples.
type NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroupWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Isis_Level_Lsp_Tlv_IsisNeighborAttribute_Neighbor_Instance_Subtlv_AdminGroup, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
