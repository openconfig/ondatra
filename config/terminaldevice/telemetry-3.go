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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_IngressPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ingress {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, false, true)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ingress{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_IngressPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ingress {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_IngressPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ingress {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ingress
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ingress{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress with a ONCE subscription.
func (n *TerminalDevice_Channel_IngressPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ingress {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ingress
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress.
func (n *TerminalDevice_Channel_IngressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress in the given batch object.
func (n *TerminalDevice_Channel_IngressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress.
func (n *TerminalDevice_Channel_IngressPath) Replace(t testing.TB, val *oc.TerminalDevice_Channel_Ingress) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress in the given batch object.
func (n *TerminalDevice_Channel_IngressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Ingress) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress.
func (n *TerminalDevice_Channel_IngressPath) Update(t testing.TB, val *oc.TerminalDevice_Channel_Ingress) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress in the given batch object.
func (n *TerminalDevice_Channel_IngressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Ingress) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Ingress_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ingress_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ingress_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface with a ONCE subscription.
func (n *TerminalDevice_Channel_Ingress_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface in the given batch object.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface in the given batch object.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/interface in the given batch object.
func (n *TerminalDevice_Channel_Ingress_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Ingress_InterfacePath extracts the value of the leaf Interface from its parent oc.TerminalDevice_Channel_Ingress
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ingress_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ingress) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Lookup(t testing.TB) *oc.QualifiedUint16Slice {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Get(t testing.TB) []uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel with a ONCE subscription.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPathAny) Get(t testing.TB) [][]uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel in the given batch object.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Replace(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel in the given batch object.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) Update(t testing.TB, val []uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/physical-channel in the given batch object.
func (n *TerminalDevice_Channel_Ingress_PhysicalChannelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []uint16) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_Ingress_PhysicalChannelPath extracts the value of the leaf PhysicalChannel from its parent oc.TerminalDevice_Channel_Ingress
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16Slice.
func convertTerminalDevice_Channel_Ingress_PhysicalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ingress) *oc.QualifiedUint16Slice {
	t.Helper()
	qv := &oc.QualifiedUint16Slice{
		Metadata: md,
	}
	val := parent.PhysicalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ingress{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ingress", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Ingress_TransceiverPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ingress_TransceiverPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ingress{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ingress", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ingress_TransceiverPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver with a ONCE subscription.
func (n *TerminalDevice_Channel_Ingress_TransceiverPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver in the given batch object.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver in the given batch object.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ingress/config/transceiver in the given batch object.
func (n *TerminalDevice_Channel_Ingress_TransceiverPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Ingress_TransceiverPath extracts the value of the leaf Transceiver from its parent oc.TerminalDevice_Channel_Ingress
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Ingress_TransceiverPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ingress) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Transceiver
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_LogicalChannelTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Get(t testing.TB) oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_LogicalChannelTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_LogicalChannelTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type with a ONCE subscription.
func (n *TerminalDevice_Channel_LogicalChannelTypePathAny) Get(t testing.TB) []oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type in the given batch object.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Replace(t testing.TB, val oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type in the given batch object.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) Update(t testing.TB, val oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/logical-channel-type in the given batch object.
func (n *TerminalDevice_Channel_LogicalChannelTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_LogicalChannelTypePath extracts the value of the leaf LogicalChannelType from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE.
func convertTerminalDevice_Channel_LogicalChannelTypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_LOGICAL_ELEMENT_PROTOCOL_TYPE{
		Metadata: md,
	}
	val := parent.LogicalChannelType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_LoopbackModePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_LoopbackModePath(t, md, goStruct)
	}
	return (&oc.QualifiedE_TransportTypes_LoopbackModeType{
		Metadata: md,
	}).SetVal(goStruct.GetLoopbackMode())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_LoopbackModePath) Get(t testing.TB) oc.E_TransportTypes_LoopbackModeType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_LoopbackModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_LoopbackModeType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_LoopbackModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode with a ONCE subscription.
func (n *TerminalDevice_Channel_LoopbackModePathAny) Get(t testing.TB) []oc.E_TransportTypes_LoopbackModeType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_LoopbackModeType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode.
func (n *TerminalDevice_Channel_LoopbackModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode in the given batch object.
func (n *TerminalDevice_Channel_LoopbackModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode.
func (n *TerminalDevice_Channel_LoopbackModePath) Replace(t testing.TB, val oc.E_TransportTypes_LoopbackModeType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode in the given batch object.
func (n *TerminalDevice_Channel_LoopbackModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_LoopbackModeType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode.
func (n *TerminalDevice_Channel_LoopbackModePath) Update(t testing.TB, val oc.E_TransportTypes_LoopbackModeType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/loopback-mode in the given batch object.
func (n *TerminalDevice_Channel_LoopbackModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_LoopbackModeType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_LoopbackModePath extracts the value of the leaf LoopbackMode from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_LoopbackModeType.
func convertTerminalDevice_Channel_LoopbackModePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_LoopbackModeType {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_LoopbackModeType{
		Metadata: md,
	}
	val := parent.LoopbackMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_OtnPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Otn {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Otn{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Otn", goStruct, false, true)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Otn{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_OtnPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Otn {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_OtnPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Otn {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Otn
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Otn{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Otn", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Otn{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn with a ONCE subscription.
func (n *TerminalDevice_Channel_OtnPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Otn {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Otn
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn.
func (n *TerminalDevice_Channel_OtnPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn in the given batch object.
func (n *TerminalDevice_Channel_OtnPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn.
func (n *TerminalDevice_Channel_OtnPath) Replace(t testing.TB, val *oc.TerminalDevice_Channel_Otn) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn in the given batch object.
func (n *TerminalDevice_Channel_OtnPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Otn) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn.
func (n *TerminalDevice_Channel_OtnPath) Update(t testing.TB, val *oc.TerminalDevice_Channel_Otn) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/otn in the given batch object.
func (n *TerminalDevice_Channel_OtnPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Otn) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
