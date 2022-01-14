package gnmicollectormetadata

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

// Lookup fetches the value at /gnmi-collector-metadata/meta with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *MetaPath) Lookup(t testing.TB) *oc.QualifiedMeta {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, false, true)
	if ok {
		return (&oc.QualifiedMeta{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *MetaPath) Get(t testing.TB) *oc.Meta {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *MetaPathAny) Lookup(t testing.TB) []*oc.QualifiedMeta {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedMeta
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedMeta{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta with a ONCE subscription.
func (n *MetaPathAny) Get(t testing.TB) []*oc.Meta {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Meta
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta.
func (n *MetaPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta in the given batch object.
func (n *MetaPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta.
func (n *MetaPath) Replace(t testing.TB, val *oc.Meta) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta in the given batch object.
func (n *MetaPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Meta) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta.
func (n *MetaPath) Update(t testing.TB, val *oc.Meta) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta in the given batch object.
func (n *MetaPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Meta) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/connectError with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_ConnectErrorPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_ConnectErrorPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/connectError with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_ConnectErrorPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/connectError with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_ConnectErrorPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertMeta_ConnectErrorPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/connectError with a ONCE subscription.
func (n *Meta_ConnectErrorPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/connectError.
func (n *Meta_ConnectErrorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/connectError in the given batch object.
func (n *Meta_ConnectErrorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/connectError.
func (n *Meta_ConnectErrorPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/connectError in the given batch object.
func (n *Meta_ConnectErrorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/connectError.
func (n *Meta_ConnectErrorPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/connectError in the given batch object.
func (n *Meta_ConnectErrorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_ConnectErrorPath extracts the value of the leaf ConnectError from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertMeta_ConnectErrorPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ConnectError
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/connectedAddress with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_ConnectedAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_ConnectedAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/connectedAddress with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_ConnectedAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/connectedAddress with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_ConnectedAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertMeta_ConnectedAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/connectedAddress with a ONCE subscription.
func (n *Meta_ConnectedAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/connectedAddress.
func (n *Meta_ConnectedAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/connectedAddress in the given batch object.
func (n *Meta_ConnectedAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/connectedAddress.
func (n *Meta_ConnectedAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/connectedAddress in the given batch object.
func (n *Meta_ConnectedAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/connectedAddress.
func (n *Meta_ConnectedAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/connectedAddress in the given batch object.
func (n *Meta_ConnectedAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_ConnectedAddressPath extracts the value of the leaf ConnectedAddress from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertMeta_ConnectedAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ConnectedAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/connected with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_ConnectedPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_ConnectedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/connected with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_ConnectedPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/connected with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_ConnectedPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertMeta_ConnectedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/connected with a ONCE subscription.
func (n *Meta_ConnectedPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/connected.
func (n *Meta_ConnectedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/connected in the given batch object.
func (n *Meta_ConnectedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/connected.
func (n *Meta_ConnectedPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/connected in the given batch object.
func (n *Meta_ConnectedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/connected.
func (n *Meta_ConnectedPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/connected in the given batch object.
func (n *Meta_ConnectedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_ConnectedPath extracts the value of the leaf Connected from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertMeta_ConnectedPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Connected
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latencyAvg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_LatencyAvgPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_LatencyAvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latencyAvg with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_LatencyAvgPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latencyAvg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_LatencyAvgPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertMeta_LatencyAvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latencyAvg with a ONCE subscription.
func (n *Meta_LatencyAvgPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latencyAvg.
func (n *Meta_LatencyAvgPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latencyAvg in the given batch object.
func (n *Meta_LatencyAvgPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latencyAvg.
func (n *Meta_LatencyAvgPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latencyAvg in the given batch object.
func (n *Meta_LatencyAvgPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latencyAvg.
func (n *Meta_LatencyAvgPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latencyAvg in the given batch object.
func (n *Meta_LatencyAvgPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_LatencyAvgPath extracts the value of the leaf LatencyAvg from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_LatencyAvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.LatencyAvg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
