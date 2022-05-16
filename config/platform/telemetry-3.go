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

// Lookup fetches the value at /openconfig-platform/components/component/storage with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_StoragePath) Lookup(t testing.TB) *oc.QualifiedComponent_Storage {
	t.Helper()
	goStruct := &oc.Component_Storage{}
	md, ok := oc.Lookup(t, n, "Component_Storage", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Storage{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/storage with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_StoragePath) Get(t testing.TB) *oc.Component_Storage {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/storage with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_StoragePathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Storage {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Storage
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Storage{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Storage", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Storage{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/storage with a ONCE subscription.
func (n *Component_StoragePathAny) Get(t testing.TB) []*oc.Component_Storage {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Storage
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/storage.
func (n *Component_StoragePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/storage in the given batch object.
func (n *Component_StoragePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/storage.
func (n *Component_StoragePath) Replace(t testing.TB, val *oc.Component_Storage) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/storage in the given batch object.
func (n *Component_StoragePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Storage) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/storage.
func (n *Component_StoragePath) Update(t testing.TB, val *oc.Component_Storage) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/storage in the given batch object.
func (n *Component_StoragePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Storage) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_SubcomponentPath) Lookup(t testing.TB) *oc.QualifiedComponent_Subcomponent {
	t.Helper()
	goStruct := &oc.Component_Subcomponent{}
	md, ok := oc.Lookup(t, n, "Component_Subcomponent", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Subcomponent{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_SubcomponentPath) Get(t testing.TB) *oc.Component_Subcomponent {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_SubcomponentPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Subcomponent {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Subcomponent
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Subcomponent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Subcomponent", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Subcomponent{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent with a ONCE subscription.
func (n *Component_SubcomponentPathAny) Get(t testing.TB) []*oc.Component_Subcomponent {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Subcomponent
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/subcomponents/subcomponent.
func (n *Component_SubcomponentPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/subcomponents/subcomponent in the given batch object.
func (n *Component_SubcomponentPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/subcomponents/subcomponent.
func (n *Component_SubcomponentPath) Replace(t testing.TB, val *oc.Component_Subcomponent) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/subcomponents/subcomponent in the given batch object.
func (n *Component_SubcomponentPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Subcomponent) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/subcomponents/subcomponent.
func (n *Component_SubcomponentPath) Update(t testing.TB, val *oc.Component_Subcomponent) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/subcomponents/subcomponent in the given batch object.
func (n *Component_SubcomponentPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Subcomponent) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Subcomponent_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Subcomponent{}
	md, ok := oc.Lookup(t, n, "Component_Subcomponent", goStruct, true, true)
	if ok {
		return convertComponent_Subcomponent_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Subcomponent_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Subcomponent_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Subcomponent{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Subcomponent", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertComponent_Subcomponent_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/subcomponents/subcomponent/config/name with a ONCE subscription.
func (n *Component_Subcomponent_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/subcomponents/subcomponent/config/name.
func (n *Component_Subcomponent_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/subcomponents/subcomponent/config/name in the given batch object.
func (n *Component_Subcomponent_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/subcomponents/subcomponent/config/name.
func (n *Component_Subcomponent_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/subcomponents/subcomponent/config/name in the given batch object.
func (n *Component_Subcomponent_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/subcomponents/subcomponent/config/name.
func (n *Component_Subcomponent_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/subcomponents/subcomponent/config/name in the given batch object.
func (n *Component_Subcomponent_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Subcomponent_NamePath extracts the value of the leaf Name from its parent oc.Component_Subcomponent
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Subcomponent_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Subcomponent) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-platform/components/component/transceiver with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_TransceiverPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver {
	t.Helper()
	goStruct := &oc.Component_Transceiver{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_TransceiverPath) Get(t testing.TB) *oc.Component_Transceiver {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_TransceiverPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver with a ONCE subscription.
func (n *Component_TransceiverPathAny) Get(t testing.TB) []*oc.Component_Transceiver {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver.
func (n *Component_TransceiverPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver in the given batch object.
func (n *Component_TransceiverPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver.
func (n *Component_TransceiverPath) Replace(t testing.TB, val *oc.Component_Transceiver) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver in the given batch object.
func (n *Component_TransceiverPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Transceiver) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver.
func (n *Component_TransceiverPath) Update(t testing.TB, val *oc.Component_Transceiver) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver in the given batch object.
func (n *Component_TransceiverPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Transceiver) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_ChannelPath) Lookup(t testing.TB) *oc.QualifiedComponent_Transceiver_Channel {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, false, true)
	if ok {
		return (&oc.QualifiedComponent_Transceiver_Channel{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_ChannelPath) Get(t testing.TB) *oc.Component_Transceiver_Channel {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_ChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedComponent_Transceiver_Channel {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedComponent_Transceiver_Channel
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Component_Transceiver_Channel{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Component_Transceiver_Channel", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedComponent_Transceiver_Channel{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel with a ONCE subscription.
func (n *Component_Transceiver_ChannelPathAny) Get(t testing.TB) []*oc.Component_Transceiver_Channel {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Component_Transceiver_Channel
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel.
func (n *Component_Transceiver_ChannelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/physical-channels/channel in the given batch object.
func (n *Component_Transceiver_ChannelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel.
func (n *Component_Transceiver_ChannelPath) Replace(t testing.TB, val *oc.Component_Transceiver_Channel) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/physical-channels/channel in the given batch object.
func (n *Component_Transceiver_ChannelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Transceiver_Channel) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel.
func (n *Component_Transceiver_ChannelPath) Update(t testing.TB, val *oc.Component_Transceiver_Channel) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/physical-channels/channel in the given batch object.
func (n *Component_Transceiver_ChannelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Component_Transceiver_Channel) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Component_Transceiver_Channel{}
	md, ok := oc.Lookup(t, n, "Component_Transceiver_Channel", goStruct, true, true)
	if ok {
		return convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel with a ONCE subscription.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel in the given batch object.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel in the given batch object.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-platform/components/component/transceiver/physical-channels/channel/config/associated-optical-channel in the given batch object.
func (n *Component_Transceiver_Channel_AssociatedOpticalChannelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath extracts the value of the leaf AssociatedOpticalChannel from its parent oc.Component_Transceiver_Channel
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertComponent_Transceiver_Channel_AssociatedOpticalChannelPath(t testing.TB, md *genutil.Metadata, parent *oc.Component_Transceiver_Channel) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AssociatedOpticalChannel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

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
