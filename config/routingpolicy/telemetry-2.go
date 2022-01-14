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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Lookup(t testing.TB) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Get(t testing.TB) oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPathAny) Get(t testing.TB) []oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpPolicy_BgpSetCommunityOptionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Replace(t testing.TB, val oc.E_BgpPolicy_BgpSetCommunityOptionType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpPolicy_BgpSetCommunityOptionType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) Update(t testing.TB, val oc.E_BgpPolicy_BgpSetCommunityOptionType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/config/options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpPolicy_BgpSetCommunityOptionType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath extracts the value of the leaf Options from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_OptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	qv := &oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType{
		Metadata: md,
	}
	val := parent.Options
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_ReferencePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-community/reference/config/community-set-ref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath extracts the value of the leaf CommunitySetRef from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference_CommunitySetRefPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Reference) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CommunitySetRef
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_InlinePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Replace(t testing.TB, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) Update(t testing.TB, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/inline/config/communities in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath extracts the value of the leaf Communities from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_CommunitiesPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Inline_Communities_UnionSlice{
		Metadata: md,
	}
	val := parent.Communities
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Lookup(t testing.TB) *oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Get(t testing.TB) oc.E_SetCommunity_Method {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SetCommunity_Method {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SetCommunity_Method
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPathAny) Get(t testing.TB) []oc.E_SetCommunity_Method {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SetCommunity_Method
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Replace(t testing.TB, val oc.E_SetCommunity_Method) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SetCommunity_Method) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) Update(t testing.TB, val oc.E_SetCommunity_Method) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/method in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SetCommunity_Method) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath extracts the value of the leaf Method from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SetCommunity_Method.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_MethodPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *oc.QualifiedE_SetCommunity_Method {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Lookup(t testing.TB) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Get(t testing.TB) oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPathAny) Get(t testing.TB) []oc.E_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpPolicy_BgpSetCommunityOptionType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Replace(t testing.TB, val oc.E_BgpPolicy_BgpSetCommunityOptionType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpPolicy_BgpSetCommunityOptionType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) Update(t testing.TB, val oc.E_BgpPolicy_BgpSetCommunityOptionType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/config/options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpPolicy_BgpSetCommunityOptionType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath extracts the value of the leaf Options from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_OptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity) *oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType {
	t.Helper()
	qv := &oc.QualifiedE_BgpPolicy_BgpSetCommunityOptionType{
		Metadata: md,
	}
	val := parent.Options
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_ReferencePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/set-ext-community/reference/config/ext-community-set-ref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath extracts the value of the leaf ExtCommunitySetRef from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference_ExtCommunitySetRefPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetExtCommunity_Reference) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.ExtCommunitySetRef
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-local-pref in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath extracts the value of the leaf SetLocalPref from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetLocalPrefPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SetLocalPref
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Get(t testing.TB) oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPathAny) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Replace(t testing.TB, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) Update(t testing.TB, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-med in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath extracts the value of the leaf SetMed from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMedPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetMed_Union{
		Metadata: md,
	}
	val := parent.SetMed
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Get(t testing.TB) oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPathAny) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Replace(t testing.TB, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) Update(t testing.TB, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-next-hop in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath extracts the value of the leaf SetNextHop from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHopPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetNextHop_Union{
		Metadata: md,
	}
	val := parent.SetNextHop
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Lookup(t testing.TB) *oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Get(t testing.TB) oc.E_BgpTypes_BgpOriginAttrType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpTypes_BgpOriginAttrType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPathAny) Get(t testing.TB) []oc.E_BgpTypes_BgpOriginAttrType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpTypes_BgpOriginAttrType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Replace(t testing.TB, val oc.E_BgpTypes_BgpOriginAttrType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpTypes_BgpOriginAttrType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) Update(t testing.TB, val oc.E_BgpTypes_BgpOriginAttrType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/bgp-actions/config/set-route-origin in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpTypes_BgpOriginAttrType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath extracts the value of the leaf SetRouteOrigin from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpTypes_BgpOriginAttrType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetRouteOriginPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions) *oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	qv := &oc.QualifiedE_BgpTypes_BgpOriginAttrType{
		Metadata: md,
	}
	val := parent.SetRouteOrigin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Lookup(t testing.TB) *oc.QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Get(t testing.TB) oc.E_RoutingPolicy_PolicyResultType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPathAny) Lookup(t testing.TB) []*oc.QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_RoutingPolicy_PolicyResultType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPathAny) Get(t testing.TB) []oc.E_RoutingPolicy_PolicyResultType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_RoutingPolicy_PolicyResultType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Replace(t testing.TB, val oc.E_RoutingPolicy_PolicyResultType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_RoutingPolicy_PolicyResultType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) Update(t testing.TB, val oc.E_RoutingPolicy_PolicyResultType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/config/policy-result in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_RoutingPolicy_PolicyResultType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath extracts the value of the leaf PolicyResult from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_RoutingPolicy_PolicyResultType.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_PolicyResultPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions) *oc.QualifiedE_RoutingPolicy_PolicyResultType {
	t.Helper()
	qv := &oc.QualifiedE_RoutingPolicy_PolicyResultType{
		Metadata: md,
	}
	val := parent.PolicyResult
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTagPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_InlinePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Get(t testing.TB) []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPathAny) Get(t testing.TB) [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Replace(t testing.TB, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) Update(t testing.TB, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/inline/config/tag in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath extracts the value of the leaf Tag from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline
// and combines the update with an existing Metadata to return a *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_TagPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice {
	t.Helper()
	qv := &oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_Inline_Tag_UnionSlice{
		Metadata: md,
	}
	val := parent.Tag
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Lookup(t testing.TB) *oc.QualifiedE_SetTag_Mode {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Get(t testing.TB) oc.E_SetTag_Mode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_SetTag_Mode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SetTag_Mode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePathAny) Get(t testing.TB) []oc.E_SetTag_Mode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SetTag_Mode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Replace(t testing.TB, val oc.E_SetTag_Mode) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SetTag_Mode) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) Update(t testing.TB, val oc.E_SetTag_Mode) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/actions/set-tag/config/mode in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SetTag_Mode) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath extracts the value of the leaf Mode from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SetTag_Mode.
func convertRoutingPolicy_PolicyDefinition_Statement_Actions_SetTag_ModePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Actions_SetTag) *oc.QualifiedE_SetTag_Mode {
	t.Helper()
	qv := &oc.QualifiedE_SetTag_Mode{
		Metadata: md,
	}
	val := parent.Mode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
