package lldp

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"

	config "github.com/openconfig/ondatra/config"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-lldp/lldp/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Lldp{}
	md, ok := oc.Lookup(t, n, "Lldp", goStruct, true, true)
	if ok {
		return convertLldp_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-lldp/lldp/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLldp_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/config/enabled with a ONCE subscription.
func (n *Lldp_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lldp/lldp/config/enabled.
func (n *Lldp_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lldp/lldp/config/enabled in the given batch object.
func (n *Lldp_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lldp/lldp/config/enabled.
func (n *Lldp_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-lldp/lldp/config/enabled in the given batch object.
func (n *Lldp_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-lldp/lldp/config/enabled.
func (n *Lldp_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-lldp/lldp/config/enabled in the given batch object.
func (n *Lldp_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLldp_EnabledPath extracts the value of the leaf Enabled from its parent oc.Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLldp_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/config/hello-timer with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_HelloTimerPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Lldp{}
	md, ok := oc.Lookup(t, n, "Lldp", goStruct, true, true)
	if ok {
		return convertLldp_HelloTimerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/config/hello-timer with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_HelloTimerPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/config/hello-timer with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_HelloTimerPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLldp_HelloTimerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/config/hello-timer with a ONCE subscription.
func (n *Lldp_HelloTimerPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lldp/lldp/config/hello-timer.
func (n *Lldp_HelloTimerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lldp/lldp/config/hello-timer in the given batch object.
func (n *Lldp_HelloTimerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lldp/lldp/config/hello-timer.
func (n *Lldp_HelloTimerPath) Replace(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-lldp/lldp/config/hello-timer in the given batch object.
func (n *Lldp_HelloTimerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-lldp/lldp/config/hello-timer.
func (n *Lldp_HelloTimerPath) Update(t testing.TB, val uint64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-lldp/lldp/config/hello-timer in the given batch object.
func (n *Lldp_HelloTimerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLldp_HelloTimerPath extracts the value of the leaf HelloTimer from its parent oc.Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertLldp_HelloTimerPath(t testing.TB, md *genutil.Metadata, parent *oc.Lldp) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.HelloTimer
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lldp_InterfacePath) Lookup(t testing.TB) *oc.QualifiedLldp_Interface {
	t.Helper()
	goStruct := &oc.Lldp_Interface{}
	md, ok := oc.Lookup(t, n, "Lldp_Interface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLldp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lldp_InterfacePath) Get(t testing.TB) *oc.Lldp_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lldp_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedLldp_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLldp_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lldp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lldp_Interface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLldp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lldp/lldp/interfaces/interface with a ONCE subscription.
func (n *Lldp_InterfacePathAny) Get(t testing.TB) []*oc.Lldp_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lldp_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lldp/lldp/interfaces/interface.
func (n *Lldp_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lldp/lldp/interfaces/interface in the given batch object.
func (n *Lldp_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lldp/lldp/interfaces/interface.
func (n *Lldp_InterfacePath) Replace(t testing.TB, val *oc.Lldp_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-lldp/lldp/interfaces/interface in the given batch object.
func (n *Lldp_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Lldp_Interface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-lldp/lldp/interfaces/interface.
func (n *Lldp_InterfacePath) Update(t testing.TB, val *oc.Lldp_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-lldp/lldp/interfaces/interface in the given batch object.
func (n *Lldp_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Lldp_Interface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
