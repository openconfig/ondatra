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

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutModePath) Lookup(t testing.TB) *oc.QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutModePath) Get(t testing.TB) *oc.Component_Port_BreakoutMode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutModePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port_BreakoutMode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port_BreakoutMode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port_BreakoutMode{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode with a ONCE subscription.
func (n *Component_Port_BreakoutModePathAny) Get(t testing.TB) []*oc.Component_Port_BreakoutMode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port_BreakoutMode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode.
func (n *Component_Port_BreakoutModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode in the given batch object.
func (n *Component_Port_BreakoutModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode.
func (n *Component_Port_BreakoutModePath) Replace(t testing.TB, val *oc.Component_Port_BreakoutMode) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode in the given batch object.
func (n *Component_Port_BreakoutModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode.
func (n *Component_Port_BreakoutModePath) Update(t testing.TB, val *oc.Component_Port_BreakoutMode) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode in the given batch object.
func (n *Component_Port_BreakoutModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_GroupPath) Lookup(t testing.TB) *oc.QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_GroupPath) Get(t testing.TB) *oc.Component_Port_BreakoutMode_Group {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_GroupPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Port_BreakoutMode_Group {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Port_BreakoutMode_Group
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Port_BreakoutMode_Group{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group with a ONCE subscription.
func (n *Component_Port_BreakoutMode_GroupPathAny) Get(t testing.TB) []*oc.Component_Port_BreakoutMode_Group {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Port_BreakoutMode_Group
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group.
func (n *Component_Port_BreakoutMode_GroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group in the given batch object.
func (n *Component_Port_BreakoutMode_GroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group.
func (n *Component_Port_BreakoutMode_GroupPath) Replace(t testing.TB, val *oc.Component_Port_BreakoutMode_Group) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group in the given batch object.
func (n *Component_Port_BreakoutMode_GroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode_Group) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group.
func (n *Component_Port_BreakoutMode_GroupPath) Update(t testing.TB, val *oc.Component_Port_BreakoutMode_Group) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group in the given batch object.
func (n *Component_Port_BreakoutMode_GroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Port_BreakoutMode_Group) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Lookup(t testing.TB) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, true)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Get(t testing.TB) oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfEthernet_ETHERNET_SPEED
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPathAny) Get(t testing.TB) []oc.E_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfEthernet_ETHERNET_SPEED
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed in the given batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Replace(t testing.TB, val oc.E_IfEthernet_ETHERNET_SPEED) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed in the given batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_ETHERNET_SPEED) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) Update(t testing.TB, val oc.E_IfEthernet_ETHERNET_SPEED) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/breakout-speed in the given batch object.
func (n *Component_Port_BreakoutMode_Group_BreakoutSpeedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_IfEthernet_ETHERNET_SPEED) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath extracts the value of the leaf BreakoutSpeed from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfEthernet_ETHERNET_SPEED.
func convertComponent_Port_BreakoutMode_Group_BreakoutSpeedPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedE_IfEthernet_ETHERNET_SPEED {
	t.Helper()
	qv := &oc.QualifiedE_IfEthernet_ETHERNET_SPEED{
		Metadata: md,
	}
	val := parent.BreakoutSpeed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, true)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_IndexPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index in the given batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index in the given batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index.
func (n *Component_Port_BreakoutMode_Group_IndexPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/index in the given batch object.
func (n *Component_Port_BreakoutMode_Group_IndexPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Port_BreakoutMode_Group_IndexPath extracts the value of the leaf Index from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Port_BreakoutMode_Group_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, true)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts in the given batch object.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts in the given batch object.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-breakouts in the given batch object.
func (n *Component_Port_BreakoutMode_Group_NumBreakoutsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath extracts the value of the leaf NumBreakouts from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Port_BreakoutMode_Group_NumBreakoutsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.NumBreakouts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Component_Port_BreakoutMode_Group{}
	md, ok := oc.Lookup(t, n, "Component_Port_BreakoutMode_Group", goStruct, true, true)
	if ok {
		return convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Port_BreakoutMode_Group{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Port_BreakoutMode_Group", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels with a ONCE subscription.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels in the given batch object.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels in the given batch object.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/port/breakout-mode/groups/group/config/num-physical-channels in the given batch object.
func (n *Component_Port_BreakoutMode_Group_NumPhysicalChannelsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath extracts the value of the leaf NumPhysicalChannels from its parent oc.Component_Port_BreakoutMode_Group
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertComponent_Port_BreakoutMode_Group_NumPhysicalChannelsPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Port_BreakoutMode_Group) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.NumPhysicalChannels
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/power-supply with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PowerSupplyPath) Lookup(t testing.TB) *oc.QualifiedComponent_PowerSupply {
	t.Helper()
	goStruct := &oc.Component_PowerSupply{}
	md, ok := oc.Lookup(t, n, "Component_PowerSupply", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_PowerSupply{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/power-supply with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PowerSupplyPath) Get(t testing.TB) *oc.Component_PowerSupply {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/power-supply with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PowerSupplyPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_PowerSupply {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_PowerSupply
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_PowerSupply{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_PowerSupply", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_PowerSupply{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/power-supply with a ONCE subscription.
func (n *Component_PowerSupplyPathAny) Get(t testing.TB) []*oc.Component_PowerSupply {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_PowerSupply
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/power-supply.
func (n *Component_PowerSupplyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/power-supply in the given batch object.
func (n *Component_PowerSupplyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/power-supply.
func (n *Component_PowerSupplyPath) Replace(t testing.TB, val *oc.Component_PowerSupply) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/power-supply in the given batch object.
func (n *Component_PowerSupplyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_PowerSupply) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/power-supply.
func (n *Component_PowerSupplyPath) Update(t testing.TB, val *oc.Component_PowerSupply) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/power-supply in the given batch object.
func (n *Component_PowerSupplyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_PowerSupply) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_PropertyPath) Lookup(t testing.TB) *oc.QualifiedComponent_Property {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Property{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_PropertyPath) Get(t testing.TB) *oc.Component_Property {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_PropertyPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Property {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Property
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Property{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property with a ONCE subscription.
func (n *Component_PropertyPathAny) Get(t testing.TB) []*oc.Component_Property {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Property
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/properties/property.
func (n *Component_PropertyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/properties/property in the given batch object.
func (n *Component_PropertyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/properties/property.
func (n *Component_PropertyPath) Replace(t testing.TB, val *oc.Component_Property) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/properties/property in the given batch object.
func (n *Component_PropertyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Property) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/properties/property.
func (n *Component_PropertyPath) Update(t testing.TB, val *oc.Component_Property) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/properties/property in the given batch object.
func (n *Component_PropertyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Property) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Property_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, true, true)
	if ok {
		return convertComponent_Property_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Property_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Property_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Property_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property/config/name with a ONCE subscription.
func (n *Component_Property_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/properties/property/config/name.
func (n *Component_Property_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/properties/property/config/name in the given batch object.
func (n *Component_Property_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/properties/property/config/name.
func (n *Component_Property_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/properties/property/config/name in the given batch object.
func (n *Component_Property_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/properties/property/config/name.
func (n *Component_Property_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/properties/property/config/name in the given batch object.
func (n *Component_Property_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Property_NamePath extracts the value of the leaf Name from its parent oc.Component_Property
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Property_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Property) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Name
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Property_ValuePath) Lookup(t testing.TB) *oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	goStruct := &oc.Component_Property{}
	md, ok := oc.Lookup(t, n, "Component_Property", goStruct, true, true)
	if ok {
		return convertComponent_Property_ValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Property_ValuePath) Get(t testing.TB) oc.Component_Property_Value_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Property_ValuePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Property_Value_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Property{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Property", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Property_ValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/properties/property/config/value with a ONCE subscription.
func (n *Component_Property_ValuePathAny) Get(t testing.TB) []oc.Component_Property_Value_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Component_Property_Value_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/properties/property/config/value.
func (n *Component_Property_ValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/properties/property/config/value in the given batch object.
func (n *Component_Property_ValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/properties/property/config/value.
func (n *Component_Property_ValuePath) Replace(t testing.TB, val oc.Component_Property_Value_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/properties/property/config/value in the given batch object.
func (n *Component_Property_ValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Component_Property_Value_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/properties/property/config/value.
func (n *Component_Property_ValuePath) Update(t testing.TB, val oc.Component_Property_Value_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/properties/property/config/value in the given batch object.
func (n *Component_Property_ValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Component_Property_Value_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertComponent_Property_ValuePath extracts the value of the leaf Value from its parent oc.Component_Property
// and combines the update with an existing Metadata to return a *oc.QualifiedComponent_Property_Value_Union.
func convertComponent_Property_ValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Property) *oc.QualifiedComponent_Property_Value_Union {
	t.Helper()
	qv := &oc.QualifiedComponent_Property_Value_Union{
		Metadata: md,
	}
	val := parent.Value
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
