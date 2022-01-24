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

// Lookup fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Aggregate_PrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Aggregate{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Aggregate", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Aggregate_PrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Aggregate_PrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Aggregate_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLocalRoutes_Aggregate_PrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix with a ONCE subscription.
func (n *LocalRoutes_Aggregate_PrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix.
func (n *LocalRoutes_Aggregate_PrefixPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix in the given batch object.
func (n *LocalRoutes_Aggregate_PrefixPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix.
func (n *LocalRoutes_Aggregate_PrefixPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix in the given batch object.
func (n *LocalRoutes_Aggregate_PrefixPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix.
func (n *LocalRoutes_Aggregate_PrefixPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/prefix in the given batch object.
func (n *LocalRoutes_Aggregate_PrefixPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Aggregate_PrefixPath extracts the value of the leaf Prefix from its parent oc.LocalRoutes_Aggregate
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Aggregate_PrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Aggregate) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Prefix
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Aggregate_SetTagPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Aggregate_SetTag_Union {
	t.Helper()
	goStruct := &oc.LocalRoutes_Aggregate{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Aggregate", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Aggregate_SetTagPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Aggregate_SetTagPath) Get(t testing.TB) oc.LocalRoutes_Aggregate_SetTag_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Aggregate_SetTagPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Aggregate_SetTag_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Aggregate_SetTag_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Aggregate{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Aggregate", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Aggregate_SetTagPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag with a ONCE subscription.
func (n *LocalRoutes_Aggregate_SetTagPathAny) Get(t testing.TB) []oc.LocalRoutes_Aggregate_SetTag_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.LocalRoutes_Aggregate_SetTag_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag.
func (n *LocalRoutes_Aggregate_SetTagPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag in the given batch object.
func (n *LocalRoutes_Aggregate_SetTagPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag.
func (n *LocalRoutes_Aggregate_SetTagPath) Replace(t testing.TB, val oc.LocalRoutes_Aggregate_SetTag_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag in the given batch object.
func (n *LocalRoutes_Aggregate_SetTagPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.LocalRoutes_Aggregate_SetTag_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag.
func (n *LocalRoutes_Aggregate_SetTagPath) Update(t testing.TB, val oc.LocalRoutes_Aggregate_SetTag_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/local-aggregates/aggregate/config/set-tag in the given batch object.
func (n *LocalRoutes_Aggregate_SetTagPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.LocalRoutes_Aggregate_SetTag_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertLocalRoutes_Aggregate_SetTagPath extracts the value of the leaf SetTag from its parent oc.LocalRoutes_Aggregate
// and combines the update with an existing Metadata to return a *oc.QualifiedLocalRoutes_Aggregate_SetTag_Union.
func convertLocalRoutes_Aggregate_SetTagPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Aggregate) *oc.QualifiedLocalRoutes_Aggregate_SetTag_Union {
	t.Helper()
	qv := &oc.QualifiedLocalRoutes_Aggregate_SetTag_Union{
		Metadata: md,
	}
	val := parent.SetTag
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_StaticPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLocalRoutes_Static{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_StaticPath) Get(t testing.TB) *oc.LocalRoutes_Static {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_StaticPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Static{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static with a ONCE subscription.
func (n *LocalRoutes_StaticPathAny) Get(t testing.TB) []*oc.LocalRoutes_Static {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Static
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static.
func (n *LocalRoutes_StaticPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static in the given batch object.
func (n *LocalRoutes_StaticPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static.
func (n *LocalRoutes_StaticPath) Replace(t testing.TB, val *oc.LocalRoutes_Static) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static in the given batch object.
func (n *LocalRoutes_StaticPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static.
func (n *LocalRoutes_StaticPath) Update(t testing.TB, val *oc.LocalRoutes_Static) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static in the given batch object.
func (n *LocalRoutes_StaticPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/config/description with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_DescriptionPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_DescriptionPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/config/description with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_DescriptionPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/config/description with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_DescriptionPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_DescriptionPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/config/description with a ONCE subscription.
func (n *LocalRoutes_Static_DescriptionPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/description.
func (n *LocalRoutes_Static_DescriptionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/config/description in the given batch object.
func (n *LocalRoutes_Static_DescriptionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/description.
func (n *LocalRoutes_Static_DescriptionPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/config/description in the given batch object.
func (n *LocalRoutes_Static_DescriptionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/description.
func (n *LocalRoutes_Static_DescriptionPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/config/description in the given batch object.
func (n *LocalRoutes_Static_DescriptionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_DescriptionPath extracts the value of the leaf Description from its parent oc.LocalRoutes_Static
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Static_DescriptionPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static) *oc.QualifiedString {
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
