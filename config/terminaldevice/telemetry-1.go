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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_IndexPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index in the given batch object.
func (n *TerminalDevice_Channel_Assignment_IndexPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index in the given batch object.
func (n *TerminalDevice_Channel_Assignment_IndexPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index.
func (n *TerminalDevice_Channel_Assignment_IndexPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/index in the given batch object.
func (n *TerminalDevice_Channel_Assignment_IndexPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Assignment_IndexPath extracts the value of the leaf Index from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_Assignment_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel in the given batch object.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel in the given batch object.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/logical-channel in the given batch object.
func (n *TerminalDevice_Channel_Assignment_LogicalChannelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Assignment_LogicalChannelPath extracts the value of the leaf LogicalChannel from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_Assignment_LogicalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.LogicalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_MappingPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Get(t testing.TB) oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_MappingPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_MappingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_MappingPathAny) Get(t testing.TB) []oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping in the given batch object.
func (n *TerminalDevice_Channel_Assignment_MappingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Replace(t testing.TB, val oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping in the given batch object.
func (n *TerminalDevice_Channel_Assignment_MappingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping.
func (n *TerminalDevice_Channel_Assignment_MappingPath) Update(t testing.TB, val oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/mapping in the given batch object.
func (n *TerminalDevice_Channel_Assignment_MappingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_FRAME_MAPPING_PROTOCOL) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_Assignment_MappingPath extracts the value of the leaf Mapping from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL.
func convertTerminalDevice_Channel_Assignment_MappingPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_FRAME_MAPPING_PROTOCOL{
		Metadata: md,
	}
	val := parent.Mapping
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel in the given batch object.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel in the given batch object.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/optical-channel in the given batch object.
func (n *TerminalDevice_Channel_Assignment_OpticalChannelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Assignment_OpticalChannelPath extracts the value of the leaf OpticalChannel from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Assignment_OpticalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OpticalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Lookup(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny) Lookup(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPathAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index in the given batch object.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Replace(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index in the given batch object.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) Update(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/tributary-slot-index in the given batch object.
func (n *TerminalDevice_Channel_Assignment_TributarySlotIndexPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath extracts the value of the leaf TributarySlotIndex from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedInt32.
func convertTerminalDevice_Channel_Assignment_TributarySlotIndexPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedInt32 {
	t.Helper()
	qv := &oc.QualifiedInt32{
		Metadata: md,
	}
	val := parent.TributarySlotIndex
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_ClientMappingModePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_ClientMappingModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_ClientMappingModePath) Get(t testing.TB) oc.E_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_ClientMappingModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_ClientMappingModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode with a ONCE subscription.
func (n *TerminalDevice_Channel_ClientMappingModePathAny) Get(t testing.TB) []oc.E_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_CLIENT_MAPPING_MODE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode.
func (n *TerminalDevice_Channel_ClientMappingModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode in the given batch object.
func (n *TerminalDevice_Channel_ClientMappingModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode.
func (n *TerminalDevice_Channel_ClientMappingModePath) Replace(t testing.TB, val oc.E_TransportTypes_CLIENT_MAPPING_MODE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode in the given batch object.
func (n *TerminalDevice_Channel_ClientMappingModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_CLIENT_MAPPING_MODE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode.
func (n *TerminalDevice_Channel_ClientMappingModePath) Update(t testing.TB, val oc.E_TransportTypes_CLIENT_MAPPING_MODE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/client-mapping-mode in the given batch object.
func (n *TerminalDevice_Channel_ClientMappingModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_CLIENT_MAPPING_MODE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_ClientMappingModePath extracts the value of the leaf ClientMappingMode from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE.
func convertTerminalDevice_Channel_ClientMappingModePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_CLIENT_MAPPING_MODE{
		Metadata: md,
	}
	val := parent.ClientMappingMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description with a ONCE subscription.
func (n *TerminalDevice_Channel_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description.
func (n *TerminalDevice_Channel_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description in the given batch object.
func (n *TerminalDevice_Channel_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description.
func (n *TerminalDevice_Channel_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description in the given batch object.
func (n *TerminalDevice_Channel_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description.
func (n *TerminalDevice_Channel_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/description in the given batch object.
func (n *TerminalDevice_Channel_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_DescriptionPath extracts the value of the leaf Description from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedString {
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
