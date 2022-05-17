package lacp

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

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/config/system-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_SystemPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, true)
	if ok {
		return convertLacp_Interface_SystemPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/config/system-priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_SystemPriorityPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/config/system-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_SystemPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_SystemPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/config/system-priority with a ONCE subscription.
func (n *Lacp_Interface_SystemPriorityPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lacp/lacp/interfaces/interface/config/system-priority.
func (n *Lacp_Interface_SystemPriorityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lacp/lacp/interfaces/interface/config/system-priority in the given batch object.
func (n *Lacp_Interface_SystemPriorityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lacp/lacp/interfaces/interface/config/system-priority.
func (n *Lacp_Interface_SystemPriorityPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-lacp/lacp/interfaces/interface/config/system-priority in the given batch object.
func (n *Lacp_Interface_SystemPriorityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-lacp/lacp/interfaces/interface/config/system-priority.
func (n *Lacp_Interface_SystemPriorityPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-lacp/lacp/interfaces/interface/config/system-priority in the given batch object.
func (n *Lacp_Interface_SystemPriorityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLacp_Interface_SystemPriorityPath extracts the value of the leaf SystemPriority from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertLacp_Interface_SystemPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SystemPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/config/system-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_SystemPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Lacp{}
	md, ok := oc.Lookup(t, n, "Lacp", goStruct, true, true)
	if ok {
		return convertLacp_SystemPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/config/system-priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_SystemPriorityPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/config/system-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_SystemPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLacp_SystemPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/config/system-priority with a ONCE subscription.
func (n *Lacp_SystemPriorityPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lacp/lacp/config/system-priority.
func (n *Lacp_SystemPriorityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lacp/lacp/config/system-priority in the given batch object.
func (n *Lacp_SystemPriorityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lacp/lacp/config/system-priority.
func (n *Lacp_SystemPriorityPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-lacp/lacp/config/system-priority in the given batch object.
func (n *Lacp_SystemPriorityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-lacp/lacp/config/system-priority.
func (n *Lacp_SystemPriorityPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-lacp/lacp/config/system-priority in the given batch object.
func (n *Lacp_SystemPriorityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLacp_SystemPriorityPath extracts the value of the leaf SystemPriority from its parent oc.Lacp
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertLacp_SystemPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.SystemPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
