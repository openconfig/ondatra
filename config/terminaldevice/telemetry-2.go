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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_EthernetPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ethernet {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ethernet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_EthernetPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ethernet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_EthernetPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ethernet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ethernet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet with a ONCE subscription.
func (n *TerminalDevice_Channel_EthernetPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ethernet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ethernet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet.
func (n *TerminalDevice_Channel_EthernetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet in the given batch object.
func (n *TerminalDevice_Channel_EthernetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet.
func (n *TerminalDevice_Channel_EthernetPath) Replace(t testing.TB, val *oc.TerminalDevice_Channel_Ethernet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet in the given batch object.
func (n *TerminalDevice_Channel_EthernetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Ethernet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet.
func (n *TerminalDevice_Channel_EthernetPath) Update(t testing.TB, val *oc.TerminalDevice_Channel_Ethernet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet in the given batch object.
func (n *TerminalDevice_Channel_EthernetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Ethernet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetAlsDelay())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/als-delay in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_AlsDelayPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Ethernet_AlsDelayPath extracts the value of the leaf AlsDelay from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_Ethernet_AlsDelayPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.AlsDelay
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Lookup(t testing.TB) *oc.QualifiedE_Ethernet_ClientAls {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_Ethernet_ClientAls{
		Metadata: md,
	}).SetVal(goStruct.GetClientAls())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Get(t testing.TB) oc.E_Ethernet_ClientAls {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Ethernet_ClientAls {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Ethernet_ClientAls
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPathAny) Get(t testing.TB) []oc.E_Ethernet_ClientAls {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Ethernet_ClientAls
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Replace(t testing.TB, val oc.E_Ethernet_ClientAls) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Ethernet_ClientAls) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) Update(t testing.TB, val oc.E_Ethernet_ClientAls) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/config/client-als in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_ClientAlsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Ethernet_ClientAls) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_Ethernet_ClientAlsPath extracts the value of the leaf ClientAls from its parent oc.TerminalDevice_Channel_Ethernet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Ethernet_ClientAls.
func convertTerminalDevice_Channel_Ethernet_ClientAlsPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet) *oc.QualifiedE_Ethernet_ClientAls {
	t.Helper()
	qv := &oc.QualifiedE_Ethernet_ClientAls{
		Metadata: md,
	}
	val := parent.ClientAls
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp", goStruct, false, true)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Ethernet_Lldp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_LldpPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Ethernet_Lldp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_LldpPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Ethernet_Lldp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Ethernet_Lldp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) Replace(t testing.TB, val *oc.TerminalDevice_Channel_Ethernet_Lldp) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Ethernet_Lldp) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) Update(t testing.TB, val *oc.TerminalDevice_Channel_Ethernet_Lldp) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_LldpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Ethernet_Lldp) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/enabled in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath extracts the value of the leaf Enabled from its parent oc.TerminalDevice_Channel_Ethernet_Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertTerminalDevice_Channel_Ethernet_Lldp_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Ethernet_Lldp", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetSnooping())
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Ethernet_Lldp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Ethernet_Lldp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping with a ONCE subscription.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/ethernet/lldp/config/snooping in the given batch object.
func (n *TerminalDevice_Channel_Ethernet_Lldp_SnoopingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath extracts the value of the leaf Snooping from its parent oc.TerminalDevice_Channel_Ethernet_Lldp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertTerminalDevice_Channel_Ethernet_Lldp_SnoopingPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Ethernet_Lldp) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Snooping
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_IndexPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index with a ONCE subscription.
func (n *TerminalDevice_Channel_IndexPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index.
func (n *TerminalDevice_Channel_IndexPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index in the given batch object.
func (n *TerminalDevice_Channel_IndexPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index.
func (n *TerminalDevice_Channel_IndexPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index in the given batch object.
func (n *TerminalDevice_Channel_IndexPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index.
func (n *TerminalDevice_Channel_IndexPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/index in the given batch object.
func (n *TerminalDevice_Channel_IndexPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_IndexPath extracts the value of the leaf Index from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertTerminalDevice_Channel_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedUint32 {
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
