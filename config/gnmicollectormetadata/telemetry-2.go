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

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesDeletedPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetLeavesDeletedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesDeletedPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesDeletedPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetLeavesDeletedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesDeleted with a ONCE subscription.
func (n *Meta_TargetLeavesDeletedPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetLeavesDeleted.
func (n *Meta_TargetLeavesDeletedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetLeavesDeleted in the given batch object.
func (n *Meta_TargetLeavesDeletedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetLeavesDeleted.
func (n *Meta_TargetLeavesDeletedPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetLeavesDeleted in the given batch object.
func (n *Meta_TargetLeavesDeletedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetLeavesDeleted.
func (n *Meta_TargetLeavesDeletedPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetLeavesDeleted in the given batch object.
func (n *Meta_TargetLeavesDeletedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetLeavesDeletedPath extracts the value of the leaf TargetLeavesDeleted from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesDeletedPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesDeleted
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesEmptyPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetLeavesEmptyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesEmptyPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesEmptyPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetLeavesEmptyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesEmpty with a ONCE subscription.
func (n *Meta_TargetLeavesEmptyPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetLeavesEmpty.
func (n *Meta_TargetLeavesEmptyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetLeavesEmpty in the given batch object.
func (n *Meta_TargetLeavesEmptyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetLeavesEmpty.
func (n *Meta_TargetLeavesEmptyPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetLeavesEmpty in the given batch object.
func (n *Meta_TargetLeavesEmptyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetLeavesEmpty.
func (n *Meta_TargetLeavesEmptyPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetLeavesEmpty in the given batch object.
func (n *Meta_TargetLeavesEmptyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetLeavesEmptyPath extracts the value of the leaf TargetLeavesEmpty from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesEmptyPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesEmpty
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeaves with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetLeavesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeaves with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeaves with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetLeavesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeaves with a ONCE subscription.
func (n *Meta_TargetLeavesPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetLeaves.
func (n *Meta_TargetLeavesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetLeaves in the given batch object.
func (n *Meta_TargetLeavesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetLeaves.
func (n *Meta_TargetLeavesPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetLeaves in the given batch object.
func (n *Meta_TargetLeavesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetLeaves.
func (n *Meta_TargetLeavesPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetLeaves in the given batch object.
func (n *Meta_TargetLeavesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetLeavesPath extracts the value of the leaf TargetLeaves from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeaves
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesStale with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesStalePath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetLeavesStalePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesStale with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesStalePath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesStale with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesStalePathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetLeavesStalePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesStale with a ONCE subscription.
func (n *Meta_TargetLeavesStalePathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetLeavesStale.
func (n *Meta_TargetLeavesStalePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetLeavesStale in the given batch object.
func (n *Meta_TargetLeavesStalePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetLeavesStale.
func (n *Meta_TargetLeavesStalePath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetLeavesStale in the given batch object.
func (n *Meta_TargetLeavesStalePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetLeavesStale.
func (n *Meta_TargetLeavesStalePath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetLeavesStale in the given batch object.
func (n *Meta_TargetLeavesStalePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetLeavesStalePath extracts the value of the leaf TargetLeavesStale from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesStalePath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesStale
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /gnmi-collector-metadata/meta/targetLeavesSuppressed with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_TargetLeavesSuppressedPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta{}
	md, ok := oc.Lookup(t, n, "Meta", goStruct, true, true)
	if ok {
		return convertMeta_TargetLeavesSuppressedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/targetLeavesSuppressed with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_TargetLeavesSuppressedPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/targetLeavesSuppressed with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_TargetLeavesSuppressedPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
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
		qv := convertMeta_TargetLeavesSuppressedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/targetLeavesSuppressed with a ONCE subscription.
func (n *Meta_TargetLeavesSuppressedPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/targetLeavesSuppressed.
func (n *Meta_TargetLeavesSuppressedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/targetLeavesSuppressed in the given batch object.
func (n *Meta_TargetLeavesSuppressedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/targetLeavesSuppressed.
func (n *Meta_TargetLeavesSuppressedPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/targetLeavesSuppressed in the given batch object.
func (n *Meta_TargetLeavesSuppressedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/targetLeavesSuppressed.
func (n *Meta_TargetLeavesSuppressedPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/targetLeavesSuppressed in the given batch object.
func (n *Meta_TargetLeavesSuppressedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_TargetLeavesSuppressedPath extracts the value of the leaf TargetLeavesSuppressed from its parent oc.Meta
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_TargetLeavesSuppressedPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.TargetLeavesSuppressed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
