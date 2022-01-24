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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicyPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy {
	t.Helper()
	goStruct := &oc.RoutingPolicy{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicyPath) Get(t testing.TB) *oc.RoutingPolicy {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy with a ONCE subscription.
func (n *RoutingPolicyPathAny) Get(t testing.TB) []*oc.RoutingPolicy {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy.
func (n *RoutingPolicyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy in the given batch object.
func (n *RoutingPolicyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy.
func (n *RoutingPolicyPath) Replace(t testing.TB, val *oc.RoutingPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy in the given batch object.
func (n *RoutingPolicyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy.
func (n *RoutingPolicyPath) Update(t testing.TB, val *oc.RoutingPolicy) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy in the given batch object.
func (n *RoutingPolicyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSetsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSetsPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSetsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets with a ONCE subscription.
func (n *RoutingPolicy_DefinedSetsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets.
func (n *RoutingPolicy_DefinedSetsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets in the given batch object.
func (n *RoutingPolicy_DefinedSetsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets.
func (n *RoutingPolicy_DefinedSetsPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets in the given batch object.
func (n *RoutingPolicy_DefinedSetsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets.
func (n *RoutingPolicy_DefinedSetsPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets in the given batch object.
func (n *RoutingPolicy_DefinedSetsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSetsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath extracts the value of the leaf AsPathSetMember from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetMemberPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.AsPathSetMember
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/as-path-sets/as-path-set/config/as-path-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath extracts the value of the leaf AsPathSetName from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet_AsPathSetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_AsPathSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.AsPathSetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Get(t testing.TB) []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Replace(t testing.TB, val []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) Update(t testing.TB, val []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath extracts the value of the leaf CommunityMember from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMemberPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_UnionSlice{
		Metadata: md,
	}
	val := parent.CommunityMember
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/community-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath extracts the value of the leaf CommunitySetName from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunitySetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CommunitySetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Replace(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) Update(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/community-sets/community-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsType.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-member in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath extracts the value of the leaf ExtCommunityMember from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunityMemberPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.ExtCommunityMember
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/ext-community-set-name in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath extracts the value of the leaf ExtCommunitySetName from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_ExtCommunitySetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ExtCommunitySetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Replace(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) Update(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/bgp-defined-sets/ext-community-sets/ext-community-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsType.
func convertRoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_BgpDefinedSets_ExtCommunitySet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_NeighborSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_NeighborSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_NeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_NeighborSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_NeighborSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_NeighborSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_NeighborSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_NeighborSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_NeighborSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_NeighborSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_NeighborSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_NeighborSet_AddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_NeighborSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_NeighborSet_AddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/address in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSet_AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_NeighborSet_AddressPath extracts the value of the leaf Address from its parent oc.RoutingPolicy_DefinedSets_NeighborSet
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertRoutingPolicy_DefinedSets_NeighborSet_AddressPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_NeighborSet) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.Address
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_NeighborSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_NeighborSet_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_NeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_NeighborSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_NeighborSet_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/neighbor-sets/neighbor-set/config/name in the given batch object.
func (n *RoutingPolicy_DefinedSets_NeighborSet_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_DefinedSets_NeighborSet_NamePath extracts the value of the leaf Name from its parent oc.RoutingPolicy_DefinedSets_NeighborSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_DefinedSets_NeighborSet_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_NeighborSet) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) Get(t testing.TB) *oc.RoutingPolicy_DefinedSets_PrefixSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_DefinedSets_PrefixSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_DefinedSets_PrefixSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_DefinedSets_PrefixSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) Replace(t testing.TB, val *oc.RoutingPolicy_DefinedSets_PrefixSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_PrefixSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) Update(t testing.TB, val *oc.RoutingPolicy_DefinedSets_PrefixSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_DefinedSets_PrefixSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) Lookup(t testing.TB) *oc.QualifiedE_PrefixSet_Mode {
	t.Helper()
	goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_DefinedSets_PrefixSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_DefinedSets_PrefixSet_ModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) Get(t testing.TB) oc.E_PrefixSet_Mode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_PrefixSet_Mode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PrefixSet_Mode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_DefinedSets_PrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_DefinedSets_PrefixSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_DefinedSets_PrefixSet_ModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode with a ONCE subscription.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePathAny) Get(t testing.TB) []oc.E_PrefixSet_Mode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PrefixSet_Mode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) Replace(t testing.TB, val oc.E_PrefixSet_Mode) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PrefixSet_Mode) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) Update(t testing.TB, val oc.E_PrefixSet_Mode) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/defined-sets/prefix-sets/prefix-set/config/mode in the given batch object.
func (n *RoutingPolicy_DefinedSets_PrefixSet_ModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PrefixSet_Mode) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_DefinedSets_PrefixSet_ModePath extracts the value of the leaf Mode from its parent oc.RoutingPolicy_DefinedSets_PrefixSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PrefixSet_Mode.
func convertRoutingPolicy_DefinedSets_PrefixSet_ModePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_DefinedSets_PrefixSet) *oc.QualifiedE_PrefixSet_Mode {
	t.Helper()
	qv := &oc.QualifiedE_PrefixSet_Mode{
		Metadata: md,
	}
	val := parent.Mode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
