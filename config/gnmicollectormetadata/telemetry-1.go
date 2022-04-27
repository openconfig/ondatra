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

// Lookup fetches the value at /gnmi-collector-metadata/meta/latencyMax with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_LatencyMaxPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_LatencyMaxPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latencyMax with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_LatencyMaxPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latencyMax with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_LatencyMaxPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_LatencyMaxPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latencyMax with a ONCE subscription.
func (n *Meta_LatencyMaxPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latencyMax.
func (n *Meta_LatencyMaxPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latencyMax in the given batch object.
func (n *Meta_LatencyMaxPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latencyMax.
func (n *Meta_LatencyMaxPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latencyMax in the given batch object.
func (n *Meta_LatencyMaxPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latencyMax.
func (n *Meta_LatencyMaxPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latencyMax in the given batch object.
func (n *Meta_LatencyMaxPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_LatencyMaxPath extracts the value of the leaf LatencyMax from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_LatencyMaxPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.LatencyMax
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latencyMin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_LatencyMinPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_LatencyMinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latencyMin with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_LatencyMinPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latencyMin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_LatencyMinPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_LatencyMinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latencyMin with a ONCE subscription.
func (n *Meta_LatencyMinPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latencyMin.
func (n *Meta_LatencyMinPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latencyMin in the given batch object.
func (n *Meta_LatencyMinPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latencyMin.
func (n *Meta_LatencyMinPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latencyMin in the given batch object.
func (n *Meta_LatencyMinPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latencyMin.
func (n *Meta_LatencyMinPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latencyMin in the given batch object.
func (n *Meta_LatencyMinPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_LatencyMinPath extracts the value of the leaf LatencyMin from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_LatencyMinPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.LatencyMin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/latestTimestamp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_LatestTimestampPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_LatestTimestampPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latestTimestamp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_LatestTimestampPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latestTimestamp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_LatestTimestampPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_LatestTimestampPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latestTimestamp with a ONCE subscription.
func (n *Meta_LatestTimestampPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latestTimestamp.
func (n *Meta_LatestTimestampPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latestTimestamp in the given batch object.
func (n *Meta_LatestTimestampPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latestTimestamp.
func (n *Meta_LatestTimestampPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latestTimestamp in the given batch object.
func (n *Meta_LatestTimestampPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latestTimestamp.
func (n *Meta_LatestTimestampPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latestTimestamp in the given batch object.
func (n *Meta_LatestTimestampPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_LatestTimestampPath extracts the value of the leaf LatestTimestamp from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_LatestTimestampPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.LatestTimestamp
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/sync with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_SyncPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_SyncPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/sync with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_SyncPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/sync with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_SyncPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
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
		qv := convertMeta_SyncPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/sync with a ONCE subscription.
func (n *Meta_SyncPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/sync.
func (n *Meta_SyncPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/sync in the given batch object.
func (n *Meta_SyncPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/sync.
func (n *Meta_SyncPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/sync in the given batch object.
func (n *Meta_SyncPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/sync.
func (n *Meta_SyncPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/sync in the given batch object.
func (n *Meta_SyncPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_SyncPath extracts the value of the leaf Sync from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertMeta_SyncPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Sync
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesAddedPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetLeavesAddedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesAddedPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesAddedPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetLeavesAddedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesAdded with a ONCE subscription.
func (n *Meta_TargetLeavesAddedPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetLeavesAdded.
func (n *Meta_TargetLeavesAddedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetLeavesAdded in the given batch object.
func (n *Meta_TargetLeavesAddedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetLeavesAdded.
func (n *Meta_TargetLeavesAddedPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetLeavesAdded in the given batch object.
func (n *Meta_TargetLeavesAddedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetLeavesAdded.
func (n *Meta_TargetLeavesAddedPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetLeavesAdded in the given batch object.
func (n *Meta_TargetLeavesAddedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetLeavesAddedPath extracts the value of the leaf TargetLeavesAdded from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesAddedPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesAdded
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
