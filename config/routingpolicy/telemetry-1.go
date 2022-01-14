package routingpolicy

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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_PrefixSet_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_PrefixSet_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_PrefixSet_NamePath extracts the value of the leaf Name from its parent oc.RoutingPolicy_DefinedSets_PrefixSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_PrefixSet_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_PrefixSet) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet_Prefix{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_PrefixPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/ip-prefix in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath extracts the value of the leaf IpPrefix from its parent oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_IpPrefixPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.IpPrefix
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet_Prefix", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/prefixes/prefix/config/masklength-range in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath extracts the value of the leaf MasklengthRange from its parent oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_PrefixSet_Prefix_MasklengthRangePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_PrefixSet_Prefix) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MasklengthRange
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_TagSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_TagSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_TagSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_TagSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_TagSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_TagSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_TagSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set.
func (n *RoutingPolicy_DefinedSets_TagSetPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_TagSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_TagSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_TagSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_TagSet_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_TagSet_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_TagSet_NamePath extracts the value of the leaf Name from its parent oc.RoutingPolicy_DefinedSets_TagSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_TagSet_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_TagSet) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_TagSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Get(t testing.TB) []oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_TagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_TagSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePathAny) Get(t testing.TB) [][]oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Replace(t testing.TB, val []oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) Update(t testing.TB, val []oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/tag-sets/tag-set/config/tag-value in the given batch object.
func (n *RoutingPolicy_DefinedSets_TagSet_TagValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_DefinedSets_TagSet_TagValue_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_TagSet_TagValuePath extracts the value of the leaf TagValue from its parent oc.RoutingPolicy_DefinedSets_TagSet
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice.
func convertRoutingPolicy_DefinedSets_TagSet_TagValuePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_TagSet) *oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_DefinedSets_TagSet_TagValue_UnionSlice{
		Metadata: md,
	}
	val := parent.TagValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinitionPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinitionPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinitionPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinitionPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition.
func (n *RoutingPolicy_PolicyDefinitionPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition in the given batch object.
func (n *RoutingPolicy_PolicyDefinitionPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition.
func (n *RoutingPolicy_PolicyDefinitionPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition in the given batch object.
func (n *RoutingPolicy_PolicyDefinitionPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition.
func (n *RoutingPolicy_PolicyDefinitionPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition in the given batch object.
func (n *RoutingPolicy_PolicyDefinitionPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name.
func (n *RoutingPolicy_PolicyDefinition_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/config/name in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_NamePath extracts the value of the leaf Name from its parent oc.RoutingPolicy_PolicyDefinition
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_StatementPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_StatementPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_StatementPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_ActionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrependPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/asn in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath extracts the value of the leaf Asn from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_AsnPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Asn
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-as-path-prepend/config/repeat-n in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath extracts the value of the leaf RepeatN from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend_RepeatNPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetAsPathPrepend) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.RepeatN
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_InlinePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Replace(t testing.TB, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) Update(t testing.TB, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/inline/config/communities in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath extracts the value of the leaf Communities from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_CommunitiesPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_UnionSlice{
		Metadata: md,
	}
	val := parent.Communities
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Lookup(t testing.TB) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Get(t testing.TB) oc.E_SetCommunity_Method {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SetCommunity_Method
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPathAny) Get(t testing.TB) []oc.E_SetCommunity_Method {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SetCommunity_Method
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Replace(t testing.TB, val oc.E_SetCommunity_Method) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SetCommunity_Method) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) Update(t testing.TB, val oc.E_SetCommunity_Method) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/method in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SetCommunity_Method) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath extracts the value of the leaf Method from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SetCommunity_Method.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_MethodPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	qv := &oc.QualifiedE_SetCommunity_Method{
		Metadata: md,
	}
	val := parent.Method
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
