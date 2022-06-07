package platform

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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_Channel_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description with a ONCE subscription.
func (n *Component_Transceiver_Channel_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description.
func (n *Component_Transceiver_Channel_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description in the given batch object.
func (n *Component_Transceiver_Channel_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description.
func (n *Component_Transceiver_Channel_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description in the given batch object.
func (n *Component_Transceiver_Channel_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description.
func (n *Component_Transceiver_Channel_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/description in the given batch object.
func (n *Component_Transceiver_Channel_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Transceiver_Channel_DescriptionPath extracts the value of the leaf Description from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_Channel_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_Channel_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_IndexPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index with a ONCE subscription.
func (n *Component_Transceiver_Channel_IndexPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index.
func (n *Component_Transceiver_Channel_IndexPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index in the given batch object.
func (n *Component_Transceiver_Channel_IndexPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index.
func (n *Component_Transceiver_Channel_IndexPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index in the given batch object.
func (n *Component_Transceiver_Channel_IndexPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index.
func (n *Component_Transceiver_Channel_IndexPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/index in the given batch object.
func (n *Component_Transceiver_Channel_IndexPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Transceiver_Channel_IndexPath extracts the value of the leaf Index from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertComponent_Transceiver_Channel_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_Channel_TargetOutputPowerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_TargetOutputPowerPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_TargetOutputPowerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power with a ONCE subscription.
func (n *Component_Transceiver_Channel_TargetOutputPowerPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power in the given batch object.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) Replace(t testing.TB, val float64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power in the given batch object.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val float64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) Update(t testing.TB, val float64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/target-output-power in the given batch object.
func (n *Component_Transceiver_Channel_TargetOutputPowerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val float64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Transceiver_Channel_TargetOutputPowerPath extracts the value of the leaf TargetOutputPower from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertComponent_Transceiver_Channel_TargetOutputPowerPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.TargetOutputPower
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_TxLaserPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_Channel_TxLaserPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_TxLaserPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_TxLaserPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_Channel_TxLaserPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser with a ONCE subscription.
func (n *Component_Transceiver_Channel_TxLaserPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser.
func (n *Component_Transceiver_Channel_TxLaserPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser in the given batch object.
func (n *Component_Transceiver_Channel_TxLaserPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser.
func (n *Component_Transceiver_Channel_TxLaserPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser in the given batch object.
func (n *Component_Transceiver_Channel_TxLaserPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser.
func (n *Component_Transceiver_Channel_TxLaserPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/tx-laser in the given batch object.
func (n *Component_Transceiver_Channel_TxLaserPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Transceiver_Channel_TxLaserPath extracts the value of the leaf TxLaser from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertComponent_Transceiver_Channel_TxLaserPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.TxLaser
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_EnabledPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/config/enabled with a ONCE subscription.
func (n *Component_Transceiver_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/config/enabled.
func (n *Component_Transceiver_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/config/enabled in the given batch object.
func (n *Component_Transceiver_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/config/enabled.
func (n *Component_Transceiver_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/config/enabled in the given batch object.
func (n *Component_Transceiver_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/config/enabled.
func (n *Component_Transceiver_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/config/enabled in the given batch object.
func (n *Component_Transceiver_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Transceiver_EnabledPath extracts the value of the leaf Enabled from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertComponent_Transceiver_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_EthernetPmdPreconfPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Get(t testing.TB) oc.E_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_EthernetPmdPreconfPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_EthernetPmdPreconfPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf with a ONCE subscription.
func (n *Component_Transceiver_EthernetPmdPreconfPathAny) Get(t testing.TB) []oc.E_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_ETHERNET_PMD_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf in the given batch object.
func (n *Component_Transceiver_EthernetPmdPreconfPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Replace(t testing.TB, val oc.E_TransportTypes_ETHERNET_PMD_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf in the given batch object.
func (n *Component_Transceiver_EthernetPmdPreconfPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_ETHERNET_PMD_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf.
func (n *Component_Transceiver_EthernetPmdPreconfPath) Update(t testing.TB, val oc.E_TransportTypes_ETHERNET_PMD_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/config/ethernet-pmd-preconf in the given batch object.
func (n *Component_Transceiver_EthernetPmdPreconfPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_ETHERNET_PMD_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertComponent_Transceiver_EthernetPmdPreconfPath extracts the value of the leaf EthernetPmdPreconf from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE.
func convertComponent_Transceiver_EthernetPmdPreconfPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_ETHERNET_PMD_TYPE{
		Metadata: md,
	}
	val := parent.EthernetPmdPreconf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/config/fec-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FecModePath) Lookup(t testing.TB) *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_FecModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/config/fec-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FecModePath) Get(t testing.TB) oc.E_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/config/fec-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FecModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FecModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/config/fec-mode with a ONCE subscription.
func (n *Component_Transceiver_FecModePathAny) Get(t testing.TB) []oc.E_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PlatformTypes_FEC_MODE_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/config/fec-mode.
func (n *Component_Transceiver_FecModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/config/fec-mode in the given batch object.
func (n *Component_Transceiver_FecModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/config/fec-mode.
func (n *Component_Transceiver_FecModePath) Replace(t testing.TB, val oc.E_PlatformTypes_FEC_MODE_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/config/fec-mode in the given batch object.
func (n *Component_Transceiver_FecModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PlatformTypes_FEC_MODE_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/config/fec-mode.
func (n *Component_Transceiver_FecModePath) Update(t testing.TB, val oc.E_PlatformTypes_FEC_MODE_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/config/fec-mode in the given batch object.
func (n *Component_Transceiver_FecModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PlatformTypes_FEC_MODE_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertComponent_Transceiver_FecModePath extracts the value of the leaf FecMode from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE.
func convertComponent_Transceiver_FecModePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_PlatformTypes_FEC_MODE_TYPE{
		Metadata: md,
	}
	val := parent.FecMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/config/form-factor-preconf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_FormFactorPreconfPath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_FormFactorPreconfPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/config/form-factor-preconf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_FormFactorPreconfPath) Get(t testing.TB) oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/config/form-factor-preconf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_FormFactorPreconfPathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_FormFactorPreconfPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/config/form-factor-preconf with a ONCE subscription.
func (n *Component_Transceiver_FormFactorPreconfPathAny) Get(t testing.TB) []oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/config/form-factor-preconf.
func (n *Component_Transceiver_FormFactorPreconfPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/config/form-factor-preconf in the given batch object.
func (n *Component_Transceiver_FormFactorPreconfPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/config/form-factor-preconf.
func (n *Component_Transceiver_FormFactorPreconfPath) Replace(t testing.TB, val oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/config/form-factor-preconf in the given batch object.
func (n *Component_Transceiver_FormFactorPreconfPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/config/form-factor-preconf.
func (n *Component_Transceiver_FormFactorPreconfPath) Update(t testing.TB, val oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/config/form-factor-preconf in the given batch object.
func (n *Component_Transceiver_FormFactorPreconfPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertComponent_Transceiver_FormFactorPreconfPath extracts the value of the leaf FormFactorPreconf from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE.
func convertComponent_Transceiver_FormFactorPreconfPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRANSCEIVER_FORM_FACTOR_TYPE{
		Metadata: md,
	}
	val := parent.FormFactorPreconf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/config/module-functional-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_ModuleFunctionalTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/config/module-functional-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Get(t testing.TB) oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/config/module-functional-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_ModuleFunctionalTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Transceiver_ModuleFunctionalTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/config/module-functional-type with a ONCE subscription.
func (n *Component_Transceiver_ModuleFunctionalTypePathAny) Get(t testing.TB) []oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/config/module-functional-type.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/config/module-functional-type in the given batch object.
func (n *Component_Transceiver_ModuleFunctionalTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/config/module-functional-type.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Replace(t testing.TB, val oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/config/module-functional-type in the given batch object.
func (n *Component_Transceiver_ModuleFunctionalTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/config/module-functional-type.
func (n *Component_Transceiver_ModuleFunctionalTypePath) Update(t testing.TB, val oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/config/module-functional-type in the given batch object.
func (n *Component_Transceiver_ModuleFunctionalTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertComponent_Transceiver_ModuleFunctionalTypePath extracts the value of the leaf ModuleFunctionalType from its parent oc.Component_Transceiver
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE.
func convertComponent_Transceiver_ModuleFunctionalTypePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver) *oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_TRANSCEIVER_MODULE_FUNCTIONAL_TYPE{
		Metadata: md,
	}
	val := parent.ModuleFunctionalType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
