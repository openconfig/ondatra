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

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesUpdated with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesUpdatedPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetLeavesUpdatedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesUpdated with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesUpdatedPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesUpdated with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesUpdatedPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetLeavesUpdatedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesUpdated with a ONCE subscription.
func (n *Meta_TargetLeavesUpdatedPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetLeavesUpdated.
func (n *Meta_TargetLeavesUpdatedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetLeavesUpdated in the given batch object.
func (n *Meta_TargetLeavesUpdatedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetLeavesUpdated.
func (n *Meta_TargetLeavesUpdatedPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetLeavesUpdated in the given batch object.
func (n *Meta_TargetLeavesUpdatedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetLeavesUpdated.
func (n *Meta_TargetLeavesUpdatedPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetLeavesUpdated in the given batch object.
func (n *Meta_TargetLeavesUpdatedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetLeavesUpdatedPath extracts the value of the leaf TargetLeavesUpdated from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesUpdatedPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesUpdated
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetSizePath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetSizePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetSizePath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetSizePathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetSizePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetSize with a ONCE subscription.
func (n *Meta_TargetSizePathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetSize.
func (n *Meta_TargetSizePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetSize in the given batch object.
func (n *Meta_TargetSizePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetSize.
func (n *Meta_TargetSizePath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetSize in the given batch object.
func (n *Meta_TargetSizePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetSize.
func (n *Meta_TargetSizePath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetSize in the given batch object.
func (n *Meta_TargetSizePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetSizePath extracts the value of the leaf TargetSize from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetSizePath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetSize
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_WindowPath) Lookup(t testing.TB) *oc.QualifiedMeta_Window {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, false, true)
	if ok {
		return (&oc.QualifiedMeta_Window{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_WindowPath) Get(t testing.TB) *oc.Meta_Window {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_WindowPathAny) Lookup(t testing.TB) []*oc.QualifiedMeta_Window {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedMeta_Window
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedMeta_Window{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window with a ONCE subscription.
func (n *Meta_WindowPathAny) Get(t testing.TB) []*oc.Meta_Window {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Meta_Window
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latency/window.
func (n *Meta_WindowPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latency/window in the given batch object.
func (n *Meta_WindowPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latency/window.
func (n *Meta_WindowPath) Replace(t testing.TB, val *oc.Meta_Window) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latency/window in the given batch object.
func (n *Meta_WindowPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Meta_Window) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latency/window.
func (n *Meta_WindowPath) Update(t testing.TB, val *oc.Meta_Window) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latency/window in the given batch object.
func (n *Meta_WindowPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Meta_Window) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_Window_AvgPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, true, true)
	if ok {
		return convertMeta_Window_AvgPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_Window_AvgPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_Window_AvgPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertMeta_Window_AvgPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window/avg with a ONCE subscription.
func (n *Meta_Window_AvgPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latency/window/avg.
func (n *Meta_Window_AvgPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latency/window/avg in the given batch object.
func (n *Meta_Window_AvgPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latency/window/avg.
func (n *Meta_Window_AvgPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latency/window/avg in the given batch object.
func (n *Meta_Window_AvgPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latency/window/avg.
func (n *Meta_Window_AvgPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latency/window/avg in the given batch object.
func (n *Meta_Window_AvgPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_Window_AvgPath extracts the value of the leaf Avg from its parent oc.Meta_Window
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_Window_AvgPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta_Window) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.Avg
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_Window_MaxPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, true, true)
	if ok {
		return convertMeta_Window_MaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_Window_MaxPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_Window_MaxPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertMeta_Window_MaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window/max with a ONCE subscription.
func (n *Meta_Window_MaxPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latency/window/max.
func (n *Meta_Window_MaxPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latency/window/max in the given batch object.
func (n *Meta_Window_MaxPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latency/window/max.
func (n *Meta_Window_MaxPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latency/window/max in the given batch object.
func (n *Meta_Window_MaxPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latency/window/max.
func (n *Meta_Window_MaxPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latency/window/max in the given batch object.
func (n *Meta_Window_MaxPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_Window_MaxPath extracts the value of the leaf Max from its parent oc.Meta_Window
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_Window_MaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta_Window) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.Max
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
