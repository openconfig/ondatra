package lacp

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

// Lookup fetches the value at /openconfig-lacp/lacp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LacpPath) Lookup(t testing.TB) *oc.QualifiedLacp {
	t.Helper()
	goStruct := &oc.Lacp{}
	md, ok := oc.Lookup(t, n, "Lacp", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LacpPath) Get(t testing.TB) *oc.Lacp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LacpPathAny) Lookup(t testing.TB) []*oc.QualifiedLacp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp with a ONCE subscription.
func (n *LacpPathAny) Get(t testing.TB) []*oc.Lacp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lacp/lacp.
func (n *LacpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lacp/lacp in the given batch object.
func (n *LacpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lacp/lacp.
func (n *LacpPath) Replace(t testing.TB, val *oc.Lacp) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-lacp/lacp in the given batch object.
func (n *LacpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Lacp) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-lacp/lacp.
func (n *LacpPath) Update(t testing.TB, val *oc.Lacp) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-lacp/lacp in the given batch object.
func (n *LacpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Lacp) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_InterfacePath) Lookup(t testing.TB) *oc.QualifiedLacp_Interface {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLacp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_InterfacePath) Get(t testing.TB) *oc.Lacp_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedLacp_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLacp_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLacp_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface with a ONCE subscription.
func (n *Lacp_InterfacePathAny) Get(t testing.TB) []*oc.Lacp_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Lacp_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lacp/lacp/interfaces/interface.
func (n *Lacp_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lacp/lacp/interfaces/interface in the given batch object.
func (n *Lacp_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lacp/lacp/interfaces/interface.
func (n *Lacp_InterfacePath) Replace(t testing.TB, val *oc.Lacp_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-lacp/lacp/interfaces/interface in the given batch object.
func (n *Lacp_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Lacp_Interface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-lacp/lacp/interfaces/interface.
func (n *Lacp_InterfacePath) Update(t testing.TB, val *oc.Lacp_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-lacp/lacp/interfaces/interface in the given batch object.
func (n *Lacp_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Lacp_Interface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
