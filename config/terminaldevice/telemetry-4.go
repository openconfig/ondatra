package terminaldevice

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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Otn_TributarySlotGranularityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) Get(t testing.TB) oc.E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_TributarySlotGranularityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity in the given batch object.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) Replace(t testing.TB, val oc.E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity in the given batch object.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) Update(t testing.TB, val oc.E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tributary-slot-granularity in the given batch object.
func (n *TerminalDevice_Channel_Otn_TributarySlotGranularityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRIBUTARY_SLOT_GRANULARITY) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_Otn_TributarySlotGranularityPath extracts the value of the leaf TributarySlotGranularity from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY.
func convertTerminalDevice_Channel_Otn_TributarySlotGranularityPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRIBUTARY_SLOT_GRANULARITY{
		Metadata: md,
	}
	val := parent.TributarySlotGranularity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Otn_TtiMsgAutoPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_TtiMsgAutoPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-auto in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgAutoPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Otn_TtiMsgAutoPath extracts the value of the leaf TtiMsgAuto from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertTerminalDevice_Channel_Otn_TtiMsgAutoPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.TtiMsgAuto
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Otn_TtiMsgExpectedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_TtiMsgExpectedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-expected in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgExpectedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Otn_TtiMsgExpectedPath extracts the value of the leaf TtiMsgExpected from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Otn_TtiMsgExpectedPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TtiMsgExpected
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit with a ONCE subscription.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn/config/tti-msg-transmit in the given batch object.
func (n *TerminalDevice_Channel_Otn_TtiMsgTransmitPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath extracts the value of the leaf TtiMsgTransmit from its parent oc.TerminalDevice_Channel_Otn
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Otn_TtiMsgTransmitPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Otn) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TtiMsgTransmit
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_RateClassPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_RateClassPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_RateClassPath) Get(t testing.TB) oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_RateClassPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_RateClassPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class with a ONCE subscription.
func (n *TerminalDevice_Channel_RateClassPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class.
func (n *TerminalDevice_Channel_RateClassPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class in the given batch object.
func (n *TerminalDevice_Channel_RateClassPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class.
func (n *TerminalDevice_Channel_RateClassPath) Replace(t testing.TB, val oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class in the given batch object.
func (n *TerminalDevice_Channel_RateClassPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class.
func (n *TerminalDevice_Channel_RateClassPath) Update(t testing.TB, val oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/rate-class in the given batch object.
func (n *TerminalDevice_Channel_RateClassPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_RateClassPath extracts the value of the leaf RateClass from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE.
func convertTerminalDevice_Channel_RateClassPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRIBUTARY_RATE_CLASS_TYPE{
		Metadata: md,
	}
	val := parent.RateClass
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_TestSignalPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_TestSignalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_TestSignalPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_TestSignalPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_TestSignalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal with a ONCE subscription.
func (n *TerminalDevice_Channel_TestSignalPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal.
func (n *TerminalDevice_Channel_TestSignalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal in the given batch object.
func (n *TerminalDevice_Channel_TestSignalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal.
func (n *TerminalDevice_Channel_TestSignalPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal in the given batch object.
func (n *TerminalDevice_Channel_TestSignalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal.
func (n *TerminalDevice_Channel_TestSignalPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/test-signal in the given batch object.
func (n *TerminalDevice_Channel_TestSignalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_TestSignalPath extracts the value of the leaf TestSignal from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertTerminalDevice_Channel_TestSignalPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.TestSignal
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_TribProtocolPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_TribProtocolPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_TribProtocolPath) Get(t testing.TB) oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_TribProtocolPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_TribProtocolPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol with a ONCE subscription.
func (n *TerminalDevice_Channel_TribProtocolPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol.
func (n *TerminalDevice_Channel_TribProtocolPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol in the given batch object.
func (n *TerminalDevice_Channel_TribProtocolPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol.
func (n *TerminalDevice_Channel_TribProtocolPath) Replace(t testing.TB, val oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol in the given batch object.
func (n *TerminalDevice_Channel_TribProtocolPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol.
func (n *TerminalDevice_Channel_TribProtocolPath) Update(t testing.TB, val oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/trib-protocol in the given batch object.
func (n *TerminalDevice_Channel_TribProtocolPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRIBUTARY_PROTOCOL_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_TribProtocolPath extracts the value of the leaf TribProtocol from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE.
func convertTerminalDevice_Channel_TribProtocolPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRIBUTARY_PROTOCOL_TYPE{
		Metadata: md,
	}
	val := parent.TribProtocol
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
