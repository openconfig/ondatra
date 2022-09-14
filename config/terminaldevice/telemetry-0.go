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

// Lookup fetches the value at /openconfig-terminal-device/terminal-device with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevicePath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice {
	t.Helper()
	goStruct := &oc.TerminalDevice{}
	md, ok := oc.Lookup(t, n, "TerminalDevice", goStruct, false, true)
	if ok {
		return (&oc.QualifiedTerminalDevice{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevicePath) Get(t testing.TB) *oc.TerminalDevice {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevicePathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device with a ONCE subscription.
func (n *TerminalDevicePathAny) Get(t testing.TB) []*oc.TerminalDevice {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device.
func (n *TerminalDevicePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device in the given batch object.
func (n *TerminalDevicePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device.
func (n *TerminalDevicePath) Replace(t testing.TB, val *oc.TerminalDevice) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device in the given batch object.
func (n *TerminalDevicePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device.
func (n *TerminalDevicePath) Update(t testing.TB, val *oc.TerminalDevice) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device in the given batch object.
func (n *TerminalDevicePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_ChannelPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, false, true)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_ChannelPath) Get(t testing.TB) *oc.TerminalDevice_Channel {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_ChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel with a ONCE subscription.
func (n *TerminalDevice_ChannelPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel.
func (n *TerminalDevice_ChannelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel in the given batch object.
func (n *TerminalDevice_ChannelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel.
func (n *TerminalDevice_ChannelPath) Replace(t testing.TB, val *oc.TerminalDevice_Channel) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel in the given batch object.
func (n *TerminalDevice_ChannelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel.
func (n *TerminalDevice_ChannelPath) Update(t testing.TB, val *oc.TerminalDevice_Channel) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel in the given batch object.
func (n *TerminalDevice_ChannelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_AdminStatePath) Lookup(t testing.TB) *oc.QualifiedE_TransportTypes_AdminStateType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_AdminStatePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_AdminStatePath) Get(t testing.TB) oc.E_TransportTypes_AdminStateType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_AdminStatePathAny) Lookup(t testing.TB) []*oc.QualifiedE_TransportTypes_AdminStateType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_TransportTypes_AdminStateType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_AdminStatePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state with a ONCE subscription.
func (n *TerminalDevice_Channel_AdminStatePathAny) Get(t testing.TB) []oc.E_TransportTypes_AdminStateType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_TransportTypes_AdminStateType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state.
func (n *TerminalDevice_Channel_AdminStatePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state in the given batch object.
func (n *TerminalDevice_Channel_AdminStatePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state.
func (n *TerminalDevice_Channel_AdminStatePath) Replace(t testing.TB, val oc.E_TransportTypes_AdminStateType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state in the given batch object.
func (n *TerminalDevice_Channel_AdminStatePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_AdminStateType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state.
func (n *TerminalDevice_Channel_AdminStatePath) Update(t testing.TB, val oc.E_TransportTypes_AdminStateType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/config/admin-state in the given batch object.
func (n *TerminalDevice_Channel_AdminStatePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_TransportTypes_AdminStateType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_AdminStatePath extracts the value of the leaf AdminState from its parent oc.TerminalDevice_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedE_TransportTypes_AdminStateType.
func convertTerminalDevice_Channel_AdminStatePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel) *oc.QualifiedE_TransportTypes_AdminStateType {
	t.Helper()
	qv := &oc.QualifiedE_TransportTypes_AdminStateType{
		Metadata: md,
	}
	val := parent.AdminState
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_AssignmentPath) Lookup(t testing.TB) *oc.QualifiedTerminalDevice_Channel_Assignment {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, false, true)
	if ok {
		return (&oc.QualifiedTerminalDevice_Channel_Assignment{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_AssignmentPath) Get(t testing.TB) *oc.TerminalDevice_Channel_Assignment {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_AssignmentPathAny) Lookup(t testing.TB) []*oc.QualifiedTerminalDevice_Channel_Assignment {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedTerminalDevice_Channel_Assignment
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedTerminalDevice_Channel_Assignment{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment with a ONCE subscription.
func (n *TerminalDevice_Channel_AssignmentPathAny) Get(t testing.TB) []*oc.TerminalDevice_Channel_Assignment {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.TerminalDevice_Channel_Assignment
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment.
func (n *TerminalDevice_Channel_AssignmentPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment in the given batch object.
func (n *TerminalDevice_Channel_AssignmentPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment.
func (n *TerminalDevice_Channel_AssignmentPath) Replace(t testing.TB, val *oc.TerminalDevice_Channel_Assignment) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment in the given batch object.
func (n *TerminalDevice_Channel_AssignmentPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Assignment) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment.
func (n *TerminalDevice_Channel_AssignmentPath) Update(t testing.TB, val *oc.TerminalDevice_Channel_Assignment) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment in the given batch object.
func (n *TerminalDevice_Channel_AssignmentPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.TerminalDevice_Channel_Assignment) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Lookup(t testing.TB) *oc.QualifiedFloat64 {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_AllocationPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Get(t testing.TB) float64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_AllocationPathAny) Lookup(t testing.TB) []*oc.QualifiedFloat64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedFloat64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_AllocationPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_AllocationPathAny) Get(t testing.TB) []float64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation in the given batch object.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Replace(t testing.TB, val float64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation in the given batch object.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val float64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) Update(t testing.TB, val float64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/allocation in the given batch object.
func (n *TerminalDevice_Channel_Assignment_AllocationPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val float64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Assignment_AllocationPath extracts the value of the leaf Allocation from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedFloat64.
func convertTerminalDevice_Channel_Assignment_AllocationPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedFloat64 {
	t.Helper()
	qv := &oc.QualifiedFloat64{
		Metadata: md,
	}
	val := parent.Allocation
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Lookup(t testing.TB) *oc.QualifiedE_Assignment_AssignmentType {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Get(t testing.TB) oc.E_Assignment_AssignmentType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Assignment_AssignmentType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Assignment_AssignmentType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.TerminalDevice_Channel_Assignment{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "TerminalDevice_Channel_Assignment", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePathAny) Get(t testing.TB) []oc.E_Assignment_AssignmentType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Assignment_AssignmentType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type in the given batch object.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Replace(t testing.TB, val oc.E_Assignment_AssignmentType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type in the given batch object.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Assignment_AssignmentType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) Update(t testing.TB, val oc.E_Assignment_AssignmentType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/assignment-type in the given batch object.
func (n *TerminalDevice_Channel_Assignment_AssignmentTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Assignment_AssignmentType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertTerminalDevice_Channel_Assignment_AssignmentTypePath extracts the value of the leaf AssignmentType from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Assignment_AssignmentType.
func convertTerminalDevice_Channel_Assignment_AssignmentTypePath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedE_Assignment_AssignmentType {
	t.Helper()
	qv := &oc.QualifiedE_Assignment_AssignmentType{
		Metadata: md,
	}
	val := parent.AssignmentType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.TerminalDevice_Channel_Assignment{}
	md, ok := oc.Lookup(t, n, "TerminalDevice_Channel_Assignment", goStruct, true, true)
	if ok {
		return convertTerminalDevice_Channel_Assignment_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *TerminalDevice_Channel_Assignment_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertTerminalDevice_Channel_Assignment_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description with a ONCE subscription.
func (n *TerminalDevice_Channel_Assignment_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description in the given batch object.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description in the given batch object.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-terminal-device/terminal-device/logical-channels/channel/logical-channel-assignments/assignment/config/description in the given batch object.
func (n *TerminalDevice_Channel_Assignment_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertTerminalDevice_Channel_Assignment_DescriptionPath extracts the value of the leaf Description from its parent oc.TerminalDevice_Channel_Assignment
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertTerminalDevice_Channel_Assignment_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.TerminalDevice_Channel_Assignment) *oc.QualifiedString {
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
