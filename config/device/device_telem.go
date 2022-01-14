package device

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"testing"

	config "github.com/openconfig/ondatra/config"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at / with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *DevicePath) Lookup(t testing.TB) *oc.QualifiedDevice {
	t.Helper()
	goStruct := &oc.Device{}
	md, ok := oc.Lookup(t, n, "Device", goStruct, false, true)
	if ok {
		return (&oc.QualifiedDevice{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at / with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *DevicePath) Get(t testing.TB) *oc.Device {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Delete deletes the configuration at /.
func (n *DevicePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at / in the given batch object.
func (n *DevicePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /.
func (n *DevicePath) Replace(t testing.TB, val *oc.Device) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at / in the given batch object.
func (n *DevicePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Device) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /.
func (n *DevicePath) Update(t testing.TB, val *oc.Device) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at / in the given batch object.
func (n *DevicePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Device) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
