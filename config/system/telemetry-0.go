package system

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

// Lookup fetches the value at /openconfig-system/system with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *SystemPath) Lookup(t testing.TB) *oc.QualifiedSystem {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *SystemPath) Get(t testing.TB) *oc.System {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *SystemPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system with a ONCE subscription.
func (n *SystemPathAny) Get(t testing.TB) []*oc.System {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system.
func (n *SystemPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system in the given batch object.
func (n *SystemPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system.
func (n *SystemPath) Replace(t testing.TB, val *oc.System) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system in the given batch object.
func (n *SystemPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system.
func (n *SystemPath) Update(t testing.TB, val *oc.System) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system in the given batch object.
func (n *SystemPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_AaaPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa {
	t.Helper()
	goStruct := &oc.System_Aaa{}
	md, ok := oc.Lookup(t, n, "System_Aaa", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_AaaPath) Get(t testing.TB) *oc.System_Aaa {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_AaaPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa with a ONCE subscription.
func (n *System_AaaPathAny) Get(t testing.TB) []*oc.System_Aaa {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa.
func (n *System_AaaPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa in the given batch object.
func (n *System_AaaPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa.
func (n *System_AaaPath) Replace(t testing.TB, val *oc.System_Aaa) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa in the given batch object.
func (n *System_AaaPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa.
func (n *System_AaaPath) Update(t testing.TB, val *oc.System_Aaa) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa in the given batch object.
func (n *System_AaaPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_AccountingPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Accounting {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Accounting{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_AccountingPath) Get(t testing.TB) *oc.System_Aaa_Accounting {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_AccountingPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Accounting {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Accounting
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Accounting{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting with a ONCE subscription.
func (n *System_Aaa_AccountingPathAny) Get(t testing.TB) []*oc.System_Aaa_Accounting {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Accounting
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/accounting.
func (n *System_Aaa_AccountingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/accounting in the given batch object.
func (n *System_Aaa_AccountingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/accounting.
func (n *System_Aaa_AccountingPath) Replace(t testing.TB, val *oc.System_Aaa_Accounting) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/accounting in the given batch object.
func (n *System_Aaa_AccountingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Accounting) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/accounting.
func (n *System_Aaa_AccountingPath) Update(t testing.TB, val *oc.System_Aaa_Accounting) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/accounting in the given batch object.
func (n *System_Aaa_AccountingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Accounting) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/config/accounting-method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_AccountingMethodPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Accounting_AccountingMethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/config/accounting-method with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_AccountingMethodPath) Get(t testing.TB) []oc.System_Aaa_Accounting_AccountingMethod_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/config/accounting-method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Accounting_AccountingMethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/config/accounting-method with a ONCE subscription.
func (n *System_Aaa_Accounting_AccountingMethodPathAny) Get(t testing.TB) [][]oc.System_Aaa_Accounting_AccountingMethod_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_Aaa_Accounting_AccountingMethod_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/accounting/config/accounting-method.
func (n *System_Aaa_Accounting_AccountingMethodPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/accounting/config/accounting-method in the given batch object.
func (n *System_Aaa_Accounting_AccountingMethodPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/accounting/config/accounting-method.
func (n *System_Aaa_Accounting_AccountingMethodPath) Replace(t testing.TB, val []oc.System_Aaa_Accounting_AccountingMethod_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/accounting/config/accounting-method in the given batch object.
func (n *System_Aaa_Accounting_AccountingMethodPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.System_Aaa_Accounting_AccountingMethod_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/accounting/config/accounting-method.
func (n *System_Aaa_Accounting_AccountingMethodPath) Update(t testing.TB, val []oc.System_Aaa_Accounting_AccountingMethod_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/accounting/config/accounting-method in the given batch object.
func (n *System_Aaa_Accounting_AccountingMethodPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.System_Aaa_Accounting_AccountingMethod_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_Accounting_AccountingMethodPath extracts the value of the leaf AccountingMethod from its parent oc.System_Aaa_Accounting
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice.
func convertSystem_Aaa_Accounting_AccountingMethodPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Accounting) *oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Accounting_AccountingMethod_UnionSlice{
		Metadata: md,
	}
	val := parent.AccountingMethod
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_EventPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Accounting_Event {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting_Event", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Accounting_Event{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_EventPath) Get(t testing.TB) *oc.System_Aaa_Accounting_Event {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_EventPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Accounting_Event {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Accounting_Event
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting_Event", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Accounting_Event{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/events/event with a ONCE subscription.
func (n *System_Aaa_Accounting_EventPathAny) Get(t testing.TB) []*oc.System_Aaa_Accounting_Event {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Accounting_Event
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/accounting/events/event.
func (n *System_Aaa_Accounting_EventPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/accounting/events/event in the given batch object.
func (n *System_Aaa_Accounting_EventPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/accounting/events/event.
func (n *System_Aaa_Accounting_EventPath) Replace(t testing.TB, val *oc.System_Aaa_Accounting_Event) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/accounting/events/event in the given batch object.
func (n *System_Aaa_Accounting_EventPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Accounting_Event) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/accounting/events/event.
func (n *System_Aaa_Accounting_EventPath) Update(t testing.TB, val *oc.System_Aaa_Accounting_Event) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/accounting/events/event in the given batch object.
func (n *System_Aaa_Accounting_EventPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Accounting_Event) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/events/event/config/event-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_Event_EventTypePath) Lookup(t testing.TB) *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting_Event", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Accounting_Event_EventTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/events/event/config/event-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_Event_EventTypePath) Get(t testing.TB) oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/events/event/config/event-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting_Event", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Accounting_Event_EventTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/events/event/config/event-type with a ONCE subscription.
func (n *System_Aaa_Accounting_Event_EventTypePathAny) Get(t testing.TB) []oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/accounting/events/event/config/event-type.
func (n *System_Aaa_Accounting_Event_EventTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/accounting/events/event/config/event-type in the given batch object.
func (n *System_Aaa_Accounting_Event_EventTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/accounting/events/event/config/event-type.
func (n *System_Aaa_Accounting_Event_EventTypePath) Replace(t testing.TB, val oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/accounting/events/event/config/event-type in the given batch object.
func (n *System_Aaa_Accounting_Event_EventTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/accounting/events/event/config/event-type.
func (n *System_Aaa_Accounting_Event_EventTypePath) Update(t testing.TB, val oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/accounting/events/event/config/event-type in the given batch object.
func (n *System_Aaa_Accounting_Event_EventTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_Accounting_Event_EventTypePath extracts the value of the leaf EventType from its parent oc.System_Aaa_Accounting_Event
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE.
func convertSystem_Aaa_Accounting_Event_EventTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Accounting_Event) *oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_AaaTypes_AAA_ACCOUNTING_EVENT_TYPE{
		Metadata: md,
	}
	val := parent.EventType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/accounting/events/event/config/record with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Accounting_Event_RecordPath) Lookup(t testing.TB) *oc.QualifiedE_Event_Record {
	t.Helper()
	goStruct := &oc.System_Aaa_Accounting_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Accounting_Event", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Accounting_Event_RecordPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/accounting/events/event/config/record with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Accounting_Event_RecordPath) Get(t testing.TB) oc.E_Event_Record {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/accounting/events/event/config/record with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Accounting_Event_RecordPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Event_Record {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Event_Record
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Accounting_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Accounting_Event", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Accounting_Event_RecordPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/accounting/events/event/config/record with a ONCE subscription.
func (n *System_Aaa_Accounting_Event_RecordPathAny) Get(t testing.TB) []oc.E_Event_Record {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Event_Record
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/accounting/events/event/config/record.
func (n *System_Aaa_Accounting_Event_RecordPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/accounting/events/event/config/record in the given batch object.
func (n *System_Aaa_Accounting_Event_RecordPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/accounting/events/event/config/record.
func (n *System_Aaa_Accounting_Event_RecordPath) Replace(t testing.TB, val oc.E_Event_Record) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/accounting/events/event/config/record in the given batch object.
func (n *System_Aaa_Accounting_Event_RecordPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Event_Record) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/accounting/events/event/config/record.
func (n *System_Aaa_Accounting_Event_RecordPath) Update(t testing.TB, val oc.E_Event_Record) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/accounting/events/event/config/record in the given batch object.
func (n *System_Aaa_Accounting_Event_RecordPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Event_Record) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_Accounting_Event_RecordPath extracts the value of the leaf Record from its parent oc.System_Aaa_Accounting_Event
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Event_Record.
func convertSystem_Aaa_Accounting_Event_RecordPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Accounting_Event) *oc.QualifiedE_Event_Record {
	t.Helper()
	qv := &oc.QualifiedE_Event_Record{
		Metadata: md,
	}
	val := parent.Record
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_AuthenticationPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authentication{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_AuthenticationPath) Get(t testing.TB) *oc.System_Aaa_Authentication {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_AuthenticationPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authentication{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication with a ONCE subscription.
func (n *System_Aaa_AuthenticationPathAny) Get(t testing.TB) []*oc.System_Aaa_Authentication {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authentication
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication.
func (n *System_Aaa_AuthenticationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication in the given batch object.
func (n *System_Aaa_AuthenticationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication.
func (n *System_Aaa_AuthenticationPath) Replace(t testing.TB, val *oc.System_Aaa_Authentication) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication in the given batch object.
func (n *System_Aaa_AuthenticationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authentication) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication.
func (n *System_Aaa_AuthenticationPath) Update(t testing.TB, val *oc.System_Aaa_Authentication) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication in the given batch object.
func (n *System_Aaa_AuthenticationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authentication) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AdminUserPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_AdminUser{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_AdminUser", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AdminUserPath) Get(t testing.TB) *oc.System_Aaa_Authentication_AdminUser {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AdminUserPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_AdminUser {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_AdminUser
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_AdminUser{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authentication_AdminUser{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/admin-user with a ONCE subscription.
func (n *System_Aaa_Authentication_AdminUserPathAny) Get(t testing.TB) []*oc.System_Aaa_Authentication_AdminUser {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authentication_AdminUser
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/admin-user.
func (n *System_Aaa_Authentication_AdminUserPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/admin-user in the given batch object.
func (n *System_Aaa_Authentication_AdminUserPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/admin-user.
func (n *System_Aaa_Authentication_AdminUserPath) Replace(t testing.TB, val *oc.System_Aaa_Authentication_AdminUser) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/admin-user in the given batch object.
func (n *System_Aaa_Authentication_AdminUserPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authentication_AdminUser) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/admin-user.
func (n *System_Aaa_Authentication_AdminUserPath) Update(t testing.TB, val *oc.System_Aaa_Authentication_AdminUser) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/admin-user in the given batch object.
func (n *System_Aaa_Authentication_AdminUserPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authentication_AdminUser) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_AdminUser{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_AdminUser", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_AdminUser{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed with a ONCE subscription.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed in the given batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed in the given batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password-hashed in the given batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordHashedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath extracts the value of the leaf AdminPasswordHashed from its parent oc.System_Aaa_Authentication_AdminUser
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_AdminUser_AdminPasswordHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_AdminUser) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AdminPasswordHashed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_AdminUser{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_AdminUser", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_AdminUser{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_AdminUser", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password with a ONCE subscription.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password in the given batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password in the given batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/admin-user/config/admin-password in the given batch object.
func (n *System_Aaa_Authentication_AdminUser_AdminPasswordPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath extracts the value of the leaf AdminPassword from its parent oc.System_Aaa_Authentication_AdminUser
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_AdminUser_AdminPasswordPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_AdminUser) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AdminPassword
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/config/authentication-method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_AuthenticationMethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/config/authentication-method with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Get(t testing.TB) []oc.System_Aaa_Authentication_AuthenticationMethod_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/config/authentication-method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_AuthenticationMethodPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_AuthenticationMethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/config/authentication-method with a ONCE subscription.
func (n *System_Aaa_Authentication_AuthenticationMethodPathAny) Get(t testing.TB) [][]oc.System_Aaa_Authentication_AuthenticationMethod_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_Aaa_Authentication_AuthenticationMethod_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/config/authentication-method.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/config/authentication-method in the given batch object.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/config/authentication-method.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Replace(t testing.TB, val []oc.System_Aaa_Authentication_AuthenticationMethod_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/config/authentication-method in the given batch object.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.System_Aaa_Authentication_AuthenticationMethod_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/config/authentication-method.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) Update(t testing.TB, val []oc.System_Aaa_Authentication_AuthenticationMethod_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/config/authentication-method in the given batch object.
func (n *System_Aaa_Authentication_AuthenticationMethodPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.System_Aaa_Authentication_AuthenticationMethod_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_Authentication_AuthenticationMethodPath extracts the value of the leaf AuthenticationMethod from its parent oc.System_Aaa_Authentication
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice.
func convertSystem_Aaa_Authentication_AuthenticationMethodPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication) *oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Authentication_AuthenticationMethod_UnionSlice{
		Metadata: md,
	}
	val := parent.AuthenticationMethod
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_UserPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_User {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authentication_User{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_UserPath) Get(t testing.TB) *oc.System_Aaa_Authentication_User {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_UserPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_User {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_User
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authentication_User{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user with a ONCE subscription.
func (n *System_Aaa_Authentication_UserPathAny) Get(t testing.TB) []*oc.System_Aaa_Authentication_User {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authentication_User
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/users/user.
func (n *System_Aaa_Authentication_UserPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/users/user in the given batch object.
func (n *System_Aaa_Authentication_UserPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/users/user.
func (n *System_Aaa_Authentication_UserPath) Replace(t testing.TB, val *oc.System_Aaa_Authentication_User) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/users/user in the given batch object.
func (n *System_Aaa_Authentication_UserPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authentication_User) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/users/user.
func (n *System_Aaa_Authentication_UserPath) Update(t testing.TB, val *oc.System_Aaa_Authentication_User) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/users/user in the given batch object.
func (n *System_Aaa_Authentication_UserPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authentication_User) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_User_PasswordHashedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_PasswordHashedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_PasswordHashedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed with a ONCE subscription.
func (n *System_Aaa_Authentication_User_PasswordHashedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed in the given batch object.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed in the given batch object.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/users/user/config/password-hashed in the given batch object.
func (n *System_Aaa_Authentication_User_PasswordHashedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_Authentication_User_PasswordHashedPath extracts the value of the leaf PasswordHashed from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_PasswordHashedPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PasswordHashed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/password with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_PasswordPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_User_PasswordPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/password with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_PasswordPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/password with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_PasswordPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_PasswordPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/password with a ONCE subscription.
func (n *System_Aaa_Authentication_User_PasswordPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/users/user/config/password.
func (n *System_Aaa_Authentication_User_PasswordPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/users/user/config/password in the given batch object.
func (n *System_Aaa_Authentication_User_PasswordPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/users/user/config/password.
func (n *System_Aaa_Authentication_User_PasswordPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/users/user/config/password in the given batch object.
func (n *System_Aaa_Authentication_User_PasswordPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/users/user/config/password.
func (n *System_Aaa_Authentication_User_PasswordPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/users/user/config/password in the given batch object.
func (n *System_Aaa_Authentication_User_PasswordPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_Authentication_User_PasswordPath extracts the value of the leaf Password from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_PasswordPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Password
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/role with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_RolePath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_User_RolePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/role with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_RolePath) Get(t testing.TB) oc.System_Aaa_Authentication_User_Role_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/role with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_RolePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authentication_User_Role_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_RolePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/role with a ONCE subscription.
func (n *System_Aaa_Authentication_User_RolePathAny) Get(t testing.TB) []oc.System_Aaa_Authentication_User_Role_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.System_Aaa_Authentication_User_Role_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/users/user/config/role.
func (n *System_Aaa_Authentication_User_RolePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/users/user/config/role in the given batch object.
func (n *System_Aaa_Authentication_User_RolePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/users/user/config/role.
func (n *System_Aaa_Authentication_User_RolePath) Replace(t testing.TB, val oc.System_Aaa_Authentication_User_Role_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/users/user/config/role in the given batch object.
func (n *System_Aaa_Authentication_User_RolePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.System_Aaa_Authentication_User_Role_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/users/user/config/role.
func (n *System_Aaa_Authentication_User_RolePath) Update(t testing.TB, val oc.System_Aaa_Authentication_User_Role_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/users/user/config/role in the given batch object.
func (n *System_Aaa_Authentication_User_RolePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.System_Aaa_Authentication_User_Role_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_Authentication_User_RolePath extracts the value of the leaf Role from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union.
func convertSystem_Aaa_Authentication_User_RolePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedSystem_Aaa_Authentication_User_Role_Union {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Authentication_User_Role_Union{
		Metadata: md,
	}
	val := parent.Role
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_SshKeyPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_User_SshKeyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_SshKeyPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_SshKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_SshKeyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key with a ONCE subscription.
func (n *System_Aaa_Authentication_User_SshKeyPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key.
func (n *System_Aaa_Authentication_User_SshKeyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key in the given batch object.
func (n *System_Aaa_Authentication_User_SshKeyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key.
func (n *System_Aaa_Authentication_User_SshKeyPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key in the given batch object.
func (n *System_Aaa_Authentication_User_SshKeyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key.
func (n *System_Aaa_Authentication_User_SshKeyPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/users/user/config/ssh-key in the given batch object.
func (n *System_Aaa_Authentication_User_SshKeyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_Authentication_User_SshKeyPath extracts the value of the leaf SshKey from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_SshKeyPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SshKey
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/username with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authentication_User_UsernamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Aaa_Authentication_User{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authentication_User", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authentication_User_UsernamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authentication/users/user/config/username with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authentication_User_UsernamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/username with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authentication_User_UsernamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authentication_User{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authentication_User", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authentication_User_UsernamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authentication/users/user/config/username with a ONCE subscription.
func (n *System_Aaa_Authentication_User_UsernamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authentication/users/user/config/username.
func (n *System_Aaa_Authentication_User_UsernamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authentication/users/user/config/username in the given batch object.
func (n *System_Aaa_Authentication_User_UsernamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authentication/users/user/config/username.
func (n *System_Aaa_Authentication_User_UsernamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authentication/users/user/config/username in the given batch object.
func (n *System_Aaa_Authentication_User_UsernamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authentication/users/user/config/username.
func (n *System_Aaa_Authentication_User_UsernamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authentication/users/user/config/username in the given batch object.
func (n *System_Aaa_Authentication_User_UsernamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Aaa_Authentication_User_UsernamePath extracts the value of the leaf Username from its parent oc.System_Aaa_Authentication_User
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Aaa_Authentication_User_UsernamePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authentication_User) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Username
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_AuthorizationPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authorization {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authorization{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_AuthorizationPath) Get(t testing.TB) *oc.System_Aaa_Authorization {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_AuthorizationPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authorization {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authorization
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authorization{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization with a ONCE subscription.
func (n *System_Aaa_AuthorizationPathAny) Get(t testing.TB) []*oc.System_Aaa_Authorization {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authorization
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authorization.
func (n *System_Aaa_AuthorizationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authorization in the given batch object.
func (n *System_Aaa_AuthorizationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authorization.
func (n *System_Aaa_AuthorizationPath) Replace(t testing.TB, val *oc.System_Aaa_Authorization) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authorization in the given batch object.
func (n *System_Aaa_AuthorizationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authorization) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authorization.
func (n *System_Aaa_AuthorizationPath) Update(t testing.TB, val *oc.System_Aaa_Authorization) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authorization in the given batch object.
func (n *System_Aaa_AuthorizationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authorization) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization/config/authorization-method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authorization_AuthorizationMethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization/config/authorization-method with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Get(t testing.TB) []oc.System_Aaa_Authorization_AuthorizationMethod_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization/config/authorization-method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authorization_AuthorizationMethodPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authorization_AuthorizationMethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization/config/authorization-method with a ONCE subscription.
func (n *System_Aaa_Authorization_AuthorizationMethodPathAny) Get(t testing.TB) [][]oc.System_Aaa_Authorization_AuthorizationMethod_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.System_Aaa_Authorization_AuthorizationMethod_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authorization/config/authorization-method.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authorization/config/authorization-method in the given batch object.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authorization/config/authorization-method.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Replace(t testing.TB, val []oc.System_Aaa_Authorization_AuthorizationMethod_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authorization/config/authorization-method in the given batch object.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.System_Aaa_Authorization_AuthorizationMethod_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authorization/config/authorization-method.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) Update(t testing.TB, val []oc.System_Aaa_Authorization_AuthorizationMethod_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authorization/config/authorization-method in the given batch object.
func (n *System_Aaa_Authorization_AuthorizationMethodPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.System_Aaa_Authorization_AuthorizationMethod_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_Authorization_AuthorizationMethodPath extracts the value of the leaf AuthorizationMethod from its parent oc.System_Aaa_Authorization
// and combines the update with an existing Metadata to return a *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice.
func convertSystem_Aaa_Authorization_AuthorizationMethodPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authorization) *oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedSystem_Aaa_Authorization_AuthorizationMethod_UnionSlice{
		Metadata: md,
	}
	val := parent.AuthorizationMethod
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authorization_EventPath) Lookup(t testing.TB) *oc.QualifiedSystem_Aaa_Authorization_Event {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization_Event", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Aaa_Authorization_Event{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authorization_EventPath) Get(t testing.TB) *oc.System_Aaa_Authorization_Event {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authorization_EventPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Aaa_Authorization_Event {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Aaa_Authorization_Event
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization_Event", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Aaa_Authorization_Event{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization/events/event with a ONCE subscription.
func (n *System_Aaa_Authorization_EventPathAny) Get(t testing.TB) []*oc.System_Aaa_Authorization_Event {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Aaa_Authorization_Event
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authorization/events/event.
func (n *System_Aaa_Authorization_EventPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authorization/events/event in the given batch object.
func (n *System_Aaa_Authorization_EventPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authorization/events/event.
func (n *System_Aaa_Authorization_EventPath) Replace(t testing.TB, val *oc.System_Aaa_Authorization_Event) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authorization/events/event in the given batch object.
func (n *System_Aaa_Authorization_EventPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authorization_Event) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authorization/events/event.
func (n *System_Aaa_Authorization_EventPath) Update(t testing.TB, val *oc.System_Aaa_Authorization_Event) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authorization/events/event in the given batch object.
func (n *System_Aaa_Authorization_EventPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Aaa_Authorization_Event) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/aaa/authorization/events/event/config/event-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Aaa_Authorization_Event_EventTypePath) Lookup(t testing.TB) *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	goStruct := &oc.System_Aaa_Authorization_Event{}
	md, ok := oc.Lookup(t, n, "System_Aaa_Authorization_Event", goStruct, true, true)
	if ok {
		return convertSystem_Aaa_Authorization_Event_EventTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/aaa/authorization/events/event/config/event-type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Aaa_Authorization_Event_EventTypePath) Get(t testing.TB) oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/aaa/authorization/events/event/config/event-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Aaa_Authorization_Event_EventTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Aaa_Authorization_Event{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Aaa_Authorization_Event", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Aaa_Authorization_Event_EventTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/aaa/authorization/events/event/config/event-type with a ONCE subscription.
func (n *System_Aaa_Authorization_Event_EventTypePathAny) Get(t testing.TB) []oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/aaa/authorization/events/event/config/event-type.
func (n *System_Aaa_Authorization_Event_EventTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/aaa/authorization/events/event/config/event-type in the given batch object.
func (n *System_Aaa_Authorization_Event_EventTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/aaa/authorization/events/event/config/event-type.
func (n *System_Aaa_Authorization_Event_EventTypePath) Replace(t testing.TB, val oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/aaa/authorization/events/event/config/event-type in the given batch object.
func (n *System_Aaa_Authorization_Event_EventTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/aaa/authorization/events/event/config/event-type.
func (n *System_Aaa_Authorization_Event_EventTypePath) Update(t testing.TB, val oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/aaa/authorization/events/event/config/event-type in the given batch object.
func (n *System_Aaa_Authorization_Event_EventTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Aaa_Authorization_Event_EventTypePath extracts the value of the leaf EventType from its parent oc.System_Aaa_Authorization_Event
// and combines the update with an existing Metadata to return a *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE.
func convertSystem_Aaa_Authorization_Event_EventTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Aaa_Authorization_Event) *oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_AaaTypes_AAA_AUTHORIZATION_EVENT_TYPE{
		Metadata: md,
	}
	val := parent.EventType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
