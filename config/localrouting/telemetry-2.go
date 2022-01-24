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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHopPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_NextHop {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLocalRoutes_Static_NextHop{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHopPath) Get(t testing.TB) *oc.LocalRoutes_Static_NextHop {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHopPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_NextHop {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_NextHop
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Static_NextHop{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop with a ONCE subscription.
func (n *LocalRoutes_Static_NextHopPathAny) Get(t testing.TB) []*oc.LocalRoutes_Static_NextHop {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Static_NextHop
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop.
func (n *LocalRoutes_Static_NextHopPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop in the given batch object.
func (n *LocalRoutes_Static_NextHopPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop.
func (n *LocalRoutes_Static_NextHopPath) Replace(t testing.TB, val *oc.LocalRoutes_Static_NextHop) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop in the given batch object.
func (n *LocalRoutes_Static_NextHopPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static_NextHop) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop.
func (n *LocalRoutes_Static_NextHopPath) Update(t testing.TB, val *oc.LocalRoutes_Static_NextHop) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop in the given batch object.
func (n *LocalRoutes_Static_NextHopPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static_NextHop) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_EnableBfd{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_EnableBfd", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Get(t testing.TB) *oc.LocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_EnableBfdPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_EnableBfd{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_EnableBfd", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Static_NextHop_EnableBfd{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_EnableBfdPathAny) Get(t testing.TB) []*oc.LocalRoutes_Static_NextHop_EnableBfd {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Static_NextHop_EnableBfd
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd in the given batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Replace(t testing.TB, val *oc.LocalRoutes_Static_NextHop_EnableBfd) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd in the given batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static_NextHop_EnableBfd) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) Update(t testing.TB, val *oc.LocalRoutes_Static_NextHop_EnableBfd) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd in the given batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static_NextHop_EnableBfd) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_EnableBfd{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_EnableBfd", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_NextHop_EnableBfd_EnabledPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_EnableBfd{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_EnableBfd", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_EnableBfd_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled in the given batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled in the given batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/enable-bfd/config/enabled in the given batch object.
func (n *LocalRoutes_Static_NextHop_EnableBfd_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_NextHop_EnableBfd_EnabledPath extracts the value of the leaf Enabled from its parent oc.LocalRoutes_Static_NextHop_EnableBfd
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertLocalRoutes_Static_NextHop_EnableBfd_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop_EnableBfd) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_IndexPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_NextHop_IndexPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_IndexPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_IndexPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_IndexPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_IndexPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index.
func (n *LocalRoutes_Static_NextHop_IndexPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index in the given batch object.
func (n *LocalRoutes_Static_NextHop_IndexPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index.
func (n *LocalRoutes_Static_NextHop_IndexPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index in the given batch object.
func (n *LocalRoutes_Static_NextHop_IndexPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index.
func (n *LocalRoutes_Static_NextHop_IndexPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/index in the given batch object.
func (n *LocalRoutes_Static_NextHop_IndexPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_NextHop_IndexPath extracts the value of the leaf Index from its parent oc.LocalRoutes_Static_NextHop
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Static_NextHop_IndexPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Index
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
