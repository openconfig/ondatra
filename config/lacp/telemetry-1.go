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

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/config/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_IntervalPath) Lookup(t testing.TB) *oc.QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, true)
	if ok {
		return convertLacp_Interface_IntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Lacp_LacpPeriodType{
		Metadata: md,
	}).SetVal(goStruct.GetInterval())
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/config/interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_IntervalPath) Get(t testing.TB) oc.E_Lacp_LacpPeriodType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/config/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lacp_LacpPeriodType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/config/interval with a ONCE subscription.
func (n *Lacp_Interface_IntervalPathAny) Get(t testing.TB) []oc.E_Lacp_LacpPeriodType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lacp_LacpPeriodType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lacp/lacp/interfaces/interface/config/interval.
func (n *Lacp_Interface_IntervalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lacp/lacp/interfaces/interface/config/interval in the given batch object.
func (n *Lacp_Interface_IntervalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lacp/lacp/interfaces/interface/config/interval.
func (n *Lacp_Interface_IntervalPath) Replace(t testing.TB, val oc.E_Lacp_LacpPeriodType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-lacp/lacp/interfaces/interface/config/interval in the given batch object.
func (n *Lacp_Interface_IntervalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Lacp_LacpPeriodType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-lacp/lacp/interfaces/interface/config/interval.
func (n *Lacp_Interface_IntervalPath) Update(t testing.TB, val oc.E_Lacp_LacpPeriodType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-lacp/lacp/interfaces/interface/config/interval in the given batch object.
func (n *Lacp_Interface_IntervalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Lacp_LacpPeriodType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertLacp_Interface_IntervalPath extracts the value of the leaf Interval from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lacp_LacpPeriodType.
func convertLacp_Interface_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedE_Lacp_LacpPeriodType {
	t.Helper()
	qv := &oc.QualifiedE_Lacp_LacpPeriodType{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Lacp_Interface_LacpModePath) Lookup(t testing.TB) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	goStruct := &oc.Lacp_Interface{}
	md, ok := oc.Lookup(t, n, "Lacp_Interface", goStruct, true, true)
	if ok {
		return convertLacp_Interface_LacpModePath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Lacp_LacpActivityType{
		Metadata: md,
	}).SetVal(goStruct.GetLacpMode())
}

// Get fetches the value at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Lacp_Interface_LacpModePath) Get(t testing.TB) oc.E_Lacp_LacpActivityType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Lacp_Interface_LacpModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Lacp_LacpActivityType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Lacp_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Lacp_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLacp_Interface_LacpModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode with a ONCE subscription.
func (n *Lacp_Interface_LacpModePathAny) Get(t testing.TB) []oc.E_Lacp_LacpActivityType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Lacp_LacpActivityType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode.
func (n *Lacp_Interface_LacpModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode in the given batch object.
func (n *Lacp_Interface_LacpModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode.
func (n *Lacp_Interface_LacpModePath) Replace(t testing.TB, val oc.E_Lacp_LacpActivityType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode in the given batch object.
func (n *Lacp_Interface_LacpModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Lacp_LacpActivityType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode.
func (n *Lacp_Interface_LacpModePath) Update(t testing.TB, val oc.E_Lacp_LacpActivityType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-lacp/lacp/interfaces/interface/config/lacp-mode in the given batch object.
func (n *Lacp_Interface_LacpModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Lacp_LacpActivityType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertLacp_Interface_LacpModePath extracts the value of the leaf LacpMode from its parent oc.Lacp_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Lacp_LacpActivityType.
func convertLacp_Interface_LacpModePath(t testing.TB, md *genutil.Metadata, parent *oc.Lacp_Interface) *oc.QualifiedE_Lacp_LacpActivityType {
	t.Helper()
	qv := &oc.QualifiedE_Lacp_LacpActivityType{
		Metadata: md,
	}
	val := parent.LacpMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
