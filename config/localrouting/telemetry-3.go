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

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, false, true)
	if ok {
		return (&oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Get(t testing.TB) *oc.LocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedLocalRoutes_Static_NextHop_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPathAny) Get(t testing.TB) []*oc.LocalRoutes_Static_NextHop_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.LocalRoutes_Static_NextHop_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Replace(t testing.TB, val *oc.LocalRoutes_Static_NextHop_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static_NextHop_InterfaceRef) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) Update(t testing.TB, val *oc.LocalRoutes_Static_NextHop_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.LocalRoutes_Static_NextHop_InterfaceRef) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/interface in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.LocalRoutes_Static_NextHop_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertLocalRoutes_Static_NextHop_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/interface-ref/config/subinterface in the given batch object.
func (n *LocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.LocalRoutes_Static_NextHop_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLocalRoutes_Static_NextHop_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *LocalRoutes_Static_NextHop_MetricPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.LocalRoutes_Static_NextHop{}
	md, ok := oc.Lookup(t, n, "LocalRoutes_Static_NextHop", goStruct, true, true)
	if ok {
		return convertLocalRoutes_Static_NextHop_MetricPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *LocalRoutes_Static_NextHop_MetricPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *LocalRoutes_Static_NextHop_MetricPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.LocalRoutes_Static_NextHop{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "LocalRoutes_Static_NextHop", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertLocalRoutes_Static_NextHop_MetricPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric with a ONCE subscription.
func (n *LocalRoutes_Static_NextHop_MetricPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric.
func (n *LocalRoutes_Static_NextHop_MetricPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric in the given batch object.
func (n *LocalRoutes_Static_NextHop_MetricPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric.
func (n *LocalRoutes_Static_NextHop_MetricPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric in the given batch object.
func (n *LocalRoutes_Static_NextHop_MetricPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric.
func (n *LocalRoutes_Static_NextHop_MetricPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-local-routing/local-routes/static-routes/static/next-hops/next-hop/config/metric in the given batch object.
func (n *LocalRoutes_Static_NextHop_MetricPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertLocalRoutes_Static_NextHop_MetricPath extracts the value of the leaf Metric from its parent oc.LocalRoutes_Static_NextHop
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertLocalRoutes_Static_NextHop_MetricPath(t testing.TB, md *genutil.Metadata, parent *oc.LocalRoutes_Static_NextHop) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Metric
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
