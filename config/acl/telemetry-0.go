package acl

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

// Lookup fetches the value at /openconfig-acl/acl with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *AclPath) Lookup(t testing.TB) *oc.QualifiedAcl {
	t.Helper()
	goStruct := &oc.Acl{}
	md, ok := oc.Lookup(t, n, "Acl", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *AclPath) Get(t testing.TB) *oc.Acl {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *AclPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl with a ONCE subscription.
func (n *AclPathAny) Get(t testing.TB) []*oc.Acl {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl.
func (n *AclPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl in the given batch object.
func (n *AclPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl.
func (n *AclPath) Replace(t testing.TB, val *oc.Acl) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl in the given batch object.
func (n *AclPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl.
func (n *AclPath) Update(t testing.TB, val *oc.Acl) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl in the given batch object.
func (n *AclPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSetPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet {
	t.Helper()
	goStruct := &oc.Acl_AclSet{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSetPath) Get(t testing.TB) *oc.Acl_AclSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSetPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set with a ONCE subscription.
func (n *Acl_AclSetPathAny) Get(t testing.TB) []*oc.Acl_AclSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set.
func (n *Acl_AclSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set in the given batch object.
func (n *Acl_AclSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set.
func (n *Acl_AclSetPath) Replace(t testing.TB, val *oc.Acl_AclSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set in the given batch object.
func (n *Acl_AclSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set.
func (n *Acl_AclSetPath) Update(t testing.TB, val *oc.Acl_AclSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set in the given batch object.
func (n *Acl_AclSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntryPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntryPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntryPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry with a ONCE subscription.
func (n *Acl_AclSet_AclEntryPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry.
func (n *Acl_AclSet_AclEntryPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry in the given batch object.
func (n *Acl_AclSet_AclEntryPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry.
func (n *Acl_AclSet_AclEntryPath) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry in the given batch object.
func (n *Acl_AclSet_AclEntryPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry.
func (n *Acl_AclSet_AclEntryPath) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry in the given batch object.
func (n *Acl_AclSet_AclEntryPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_ActionsPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Actions {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Actions", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Actions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_ActionsPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Actions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_ActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Actions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Actions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Actions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_ActionsPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Actions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Actions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions.
func (n *Acl_AclSet_AclEntry_ActionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions in the given batch object.
func (n *Acl_AclSet_AclEntry_ActionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions.
func (n *Acl_AclSet_AclEntry_ActionsPath) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_Actions) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions in the given batch object.
func (n *Acl_AclSet_AclEntry_ActionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Actions) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions.
func (n *Acl_AclSet_AclEntry_ActionsPath) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_Actions) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions in the given batch object.
func (n *Acl_AclSet_AclEntry_ActionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Actions) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Lookup(t testing.TB) *oc.QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Actions", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Get(t testing.TB) oc.E_Acl_FORWARDING_ACTION {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_FORWARDING_ACTION
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPathAny) Get(t testing.TB) []oc.E_Acl_FORWARDING_ACTION {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_FORWARDING_ACTION
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action in the given batch object.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Replace(t testing.TB, val oc.E_Acl_FORWARDING_ACTION) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action in the given batch object.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_FORWARDING_ACTION) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) Update(t testing.TB, val oc.E_Acl_FORWARDING_ACTION) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/forwarding-action in the given batch object.
func (n *Acl_AclSet_AclEntry_Actions_ForwardingActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_FORWARDING_ACTION) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath extracts the value of the leaf ForwardingAction from its parent oc.Acl_AclSet_AclEntry_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_FORWARDING_ACTION.
func convertAcl_AclSet_AclEntry_Actions_ForwardingActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Actions) *oc.QualifiedE_Acl_FORWARDING_ACTION {
	t.Helper()
	qv := &oc.QualifiedE_Acl_FORWARDING_ACTION{
		Metadata: md,
	}
	val := parent.ForwardingAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Lookup(t testing.TB) *oc.QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Actions", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_Actions_LogActionPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Acl_LOG_ACTION{
		Metadata: md,
	}).SetVal(goStruct.GetLogAction())
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Get(t testing.TB) oc.E_Acl_LOG_ACTION {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_LOG_ACTION
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Actions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_Actions_LogActionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPathAny) Get(t testing.TB) []oc.E_Acl_LOG_ACTION {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_LOG_ACTION
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action in the given batch object.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Replace(t testing.TB, val oc.E_Acl_LOG_ACTION) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action in the given batch object.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_LOG_ACTION) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) Update(t testing.TB, val oc.E_Acl_LOG_ACTION) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/actions/config/log-action in the given batch object.
func (n *Acl_AclSet_AclEntry_Actions_LogActionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_LOG_ACTION) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_AclSet_AclEntry_Actions_LogActionPath extracts the value of the leaf LogAction from its parent oc.Acl_AclSet_AclEntry_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_LOG_ACTION.
func convertAcl_AclSet_AclEntry_Actions_LogActionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_Actions) *oc.QualifiedE_Acl_LOG_ACTION {
	t.Helper()
	qv := &oc.QualifiedE_Acl_LOG_ACTION{
		Metadata: md,
	}
	val := parent.LogAction
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description in the given batch object.
func (n *Acl_AclSet_AclEntry_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description in the given batch object.
func (n *Acl_AclSet_AclEntry_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description.
func (n *Acl_AclSet_AclEntry_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/config/description in the given batch object.
func (n *Acl_AclSet_AclEntry_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_DescriptionPath extracts the value of the leaf Description from its parent oc.Acl_AclSet_AclEntry
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Description
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_InputInterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterfacePathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_InputInterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_InputInterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_InputInterface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_InputInterface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_InputInterface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_InputInterface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_InputInterface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/interface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, true, true)
	if ok {
		return convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_InputInterface_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/input-interface/interface-ref/config/subinterface in the given batch object.
func (n *Acl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_AclSet_AclEntry_InputInterface_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_AclSet_AclEntry_InputInterface_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Lookup(t testing.TB) *oc.QualifiedAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
	md, ok := oc.Lookup(t, n, "Acl_AclSet_AclEntry_Ipv4", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_AclSet_AclEntry_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Get(t testing.TB) *oc.Acl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_AclSet_AclEntry_Ipv4PathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_AclSet_AclEntry_Ipv4
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_AclSet_AclEntry_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_AclSet_AclEntry_Ipv4", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_AclSet_AclEntry_Ipv4{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 with a ONCE subscription.
func (n *Acl_AclSet_AclEntry_Ipv4PathAny) Get(t testing.TB) []*oc.Acl_AclSet_AclEntry_Ipv4 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_AclSet_AclEntry_Ipv4
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Replace(t testing.TB, val *oc.Acl_AclSet_AclEntry_Ipv4) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Ipv4) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4.
func (n *Acl_AclSet_AclEntry_Ipv4Path) Update(t testing.TB, val *oc.Acl_AclSet_AclEntry_Ipv4) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/acl-sets/acl-set/acl-entries/acl-entry/ipv4 in the given batch object.
func (n *Acl_AclSet_AclEntry_Ipv4Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_AclSet_AclEntry_Ipv4) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
