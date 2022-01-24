package localrouting

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

// Lookup fetches the value at /openconfig-local-routing/local-routes with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutesPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes {
	t.Helper()
	goStruct := &oc.LocalRoutes{}
	md, ok := oc.Lookup(t, n, "LocalRoutes", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLocalRoutes{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutesPath) Get(t testing.TB) *oc.LocalRoutes {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutesPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes with a ONCE subscription.
func (n *LocalRoutesPathAny) Get(t testing.TB) []*oc.LocalRoutes {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes.
func (n *LocalRoutesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes in the given batch object.
func (n *LocalRoutesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes.
func (n *LocalRoutesPath) Replace(t testing.TB, val *oc.LocalRoutes) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes in the given batch object.
func (n *LocalRoutesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes.
func (n *LocalRoutesPath) Update(t testing.TB, val *oc.LocalRoutes) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes in the given batch object.
func (n *LocalRoutesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_AggregatePath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Aggregate {
	t.Helper()
	goStruct := &oc.LocalRoutes_Aggregate{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Aggregate", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLocalRoutes_Aggregate{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_AggregatePath) Get(t testing.TB) *oc.LocalRoutes_Aggregate {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_AggregatePathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Aggregate {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Aggregate
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Aggregate{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Aggregate", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Aggregate{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate with a ONCE subscription.
func (n *LocalRoutes_AggregatePathAny) Get(t testing.TB) []*oc.LocalRoutes_Aggregate {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Aggregate
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate.
func (n *LocalRoutes_AggregatePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate in the given batch object.
func (n *LocalRoutes_AggregatePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate.
func (n *LocalRoutes_AggregatePath) Replace(t testing.TB, val *oc.LocalRoutes_Aggregate) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate in the given batch object.
func (n *LocalRoutes_AggregatePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Aggregate) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate.
func (n *LocalRoutes_AggregatePath) Update(t testing.TB, val *oc.LocalRoutes_Aggregate) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate in the given batch object.
func (n *LocalRoutes_AggregatePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Aggregate) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Aggregate_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Aggregate{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Aggregate", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Aggregate_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Aggregate_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Aggregate_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Aggregate{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Aggregate", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Aggregate_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description with a ONCE subscription.
func (n *LocalRoutes_Aggregate_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description.
func (n *LocalRoutes_Aggregate_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description in the given batch object.
func (n *LocalRoutes_Aggregate_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description.
func (n *LocalRoutes_Aggregate_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description in the given batch object.
func (n *LocalRoutes_Aggregate_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description.
func (n *LocalRoutes_Aggregate_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/description in the given batch object.
func (n *LocalRoutes_Aggregate_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Aggregate_DescriptionPath extracts the value of the leaf Description from its parent oc.LocalRoutes_Aggregate
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Aggregate_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Aggregate) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Aggregate_DiscardPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.LocalRoutes_Aggregate{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Aggregate", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Aggregate_DiscardPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDiscard())
}

// Get fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Aggregate_DiscardPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Aggregate_DiscardPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Aggregate{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Aggregate", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Aggregate_DiscardPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard with a ONCE subscription.
func (n *LocalRoutes_Aggregate_DiscardPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard.
func (n *LocalRoutes_Aggregate_DiscardPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard in the given batch object.
func (n *LocalRoutes_Aggregate_DiscardPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard.
func (n *LocalRoutes_Aggregate_DiscardPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard in the given batch object.
func (n *LocalRoutes_Aggregate_DiscardPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard.
func (n *LocalRoutes_Aggregate_DiscardPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/discard in the given batch object.
func (n *LocalRoutes_Aggregate_DiscardPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Aggregate_DiscardPath extracts the value of the leaf Discard from its parent oc.LocalRoutes_Aggregate
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLocalRoutes_Aggregate_DiscardPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Aggregate) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Discard
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
