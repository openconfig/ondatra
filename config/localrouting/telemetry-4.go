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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_NextHopPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_NextHop_NextHop_Union {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_NextHop_NextHopPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_NextHopPath) Get(t testing.TB) oc.LocalRoutes_Static_NextHop_NextHop_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_NextHopPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_NextHop_NextHop_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_NextHop_NextHop_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_NextHopPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_NextHopPathAny) Get(t testing.TB) []oc.LocalRoutes_Static_NextHop_NextHop_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.LocalRoutes_Static_NextHop_NextHop_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop.
func (n *LocalRoutes_Static_NextHop_NextHopPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop in the given batch object.
func (n *LocalRoutes_Static_NextHop_NextHopPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop.
func (n *LocalRoutes_Static_NextHop_NextHopPath) Replace(t testing.TB, val oc.LocalRoutes_Static_NextHop_NextHop_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop in the given batch object.
func (n *LocalRoutes_Static_NextHop_NextHopPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.LocalRoutes_Static_NextHop_NextHop_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop.
func (n *LocalRoutes_Static_NextHop_NextHopPath) Update(t testing.TB, val oc.LocalRoutes_Static_NextHop_NextHop_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/next-hop in the given batch object.
func (n *LocalRoutes_Static_NextHop_NextHopPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.LocalRoutes_Static_NextHop_NextHop_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertLocalRoutes_Static_NextHop_NextHopPath extracts the value of the leaf NextHop from its parent oc.LocalRoutes_Static_NextHop
// and combines the update with an existing Metadata to return a *oc.QualifiedLocalRoutes_Static_NextHop_NextHop_Union.
func convertLocalRoutes_Static_NextHop_NextHopPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop) *oc.QualifiedLocalRoutes_Static_NextHop_NextHop_Union {
	t.Helper()
	qv := &oc.QualifiedLocalRoutes_Static_NextHop_NextHop_Union{
		Metadata: md,
	}
	val := parent.NextHop
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_RecursePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_NextHop_RecursePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetRecurse())
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_RecursePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_RecursePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_RecursePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_RecursePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse.
func (n *LocalRoutes_Static_NextHop_RecursePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse in the given batch object.
func (n *LocalRoutes_Static_NextHop_RecursePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse.
func (n *LocalRoutes_Static_NextHop_RecursePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse in the given batch object.
func (n *LocalRoutes_Static_NextHop_RecursePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse.
func (n *LocalRoutes_Static_NextHop_RecursePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/recurse in the given batch object.
func (n *LocalRoutes_Static_NextHop_RecursePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_NextHop_RecursePath extracts the value of the leaf Recurse from its parent oc.LocalRoutes_Static_NextHop
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLocalRoutes_Static_NextHop_RecursePath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Recurse
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/config/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_PrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_PrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/config/prefix with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_PrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/config/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
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
		qv := convertLocalRoutes_Static_PrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/config/prefix with a ONCE subscription.
func (n *LocalRoutes_Static_PrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/prefix.
func (n *LocalRoutes_Static_PrefixPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/config/prefix in the given batch object.
func (n *LocalRoutes_Static_PrefixPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/prefix.
func (n *LocalRoutes_Static_PrefixPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/config/prefix in the given batch object.
func (n *LocalRoutes_Static_PrefixPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/prefix.
func (n *LocalRoutes_Static_PrefixPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/config/prefix in the given batch object.
func (n *LocalRoutes_Static_PrefixPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_PrefixPath extracts the value of the leaf Prefix from its parent oc.LocalRoutes_Static
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Static_PrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_SetTagPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_SetTag_Union {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_SetTagPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_SetTagPath) Get(t testing.TB) oc.LocalRoutes_Static_SetTag_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_SetTagPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_SetTag_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_SetTag_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_SetTagPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag with a ONCE subscription.
func (n *LocalRoutes_Static_SetTagPathAny) Get(t testing.TB) []oc.LocalRoutes_Static_SetTag_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.LocalRoutes_Static_SetTag_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag.
func (n *LocalRoutes_Static_SetTagPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag in the given batch object.
func (n *LocalRoutes_Static_SetTagPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag.
func (n *LocalRoutes_Static_SetTagPath) Replace(t testing.TB, val oc.LocalRoutes_Static_SetTag_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag in the given batch object.
func (n *LocalRoutes_Static_SetTagPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.LocalRoutes_Static_SetTag_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag.
func (n *LocalRoutes_Static_SetTagPath) Update(t testing.TB, val oc.LocalRoutes_Static_SetTag_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/config/set-tag in the given batch object.
func (n *LocalRoutes_Static_SetTagPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.LocalRoutes_Static_SetTag_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertLocalRoutes_Static_SetTagPath extracts the value of the leaf SetTag from its parent oc.LocalRoutes_Static
// and combines the update with an existing Metadata to return a *oc.QualifiedLocalRoutes_Static_SetTag_Union.
func convertLocalRoutes_Static_SetTagPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static) *oc.QualifiedLocalRoutes_Static_SetTag_Union {
	t.Helper()
	qv := &oc.QualifiedLocalRoutes_Static_SetTag_Union{
		Metadata: md,
	}
	val := parent.SetTag
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
