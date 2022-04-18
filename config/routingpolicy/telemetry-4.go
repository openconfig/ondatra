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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) Lookup(t testing.TB) *oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) Get(t testing.TB) oc.E_BgpTypes_BgpOriginAttrType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpTypes_BgpOriginAttrType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPathAny) Get(t testing.TB) []oc.E_BgpTypes_BgpOriginAttrType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpTypes_BgpOriginAttrType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) Replace(t testing.TB, val oc.E_BgpTypes_BgpOriginAttrType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpTypes_BgpOriginAttrType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) Update(t testing.TB, val oc.E_BgpTypes_BgpOriginAttrType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/origin-eq in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpTypes_BgpOriginAttrType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath extracts the value of the leaf OriginEq from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpTypes_BgpOriginAttrType.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_OriginEqPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) *oc.QualifiedE_BgpTypes_BgpOriginAttrType {
	t.Helper()
	qv := &oc.QualifiedE_BgpTypes_BgpOriginAttrType{
		Metadata: md,
	}
	val := parent.OriginEq
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) Lookup(t testing.TB) *oc.QualifiedE_BgpConditions_RouteType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) Get(t testing.TB) oc.E_BgpConditions_RouteType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_BgpConditions_RouteType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_BgpConditions_RouteType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePathAny) Get(t testing.TB) []oc.E_BgpConditions_RouteType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_BgpConditions_RouteType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) Replace(t testing.TB, val oc.E_BgpConditions_RouteType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpConditions_RouteType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) Update(t testing.TB, val oc.E_BgpConditions_RouteType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/bgp-conditions/config/route-type in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_BgpConditions_RouteType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath extracts the value of the leaf RouteType from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_BgpConditions_RouteType.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions_RouteTypePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_BgpConditions) *oc.QualifiedE_BgpConditions_RouteType {
	t.Helper()
	qv := &oc.QualifiedE_BgpConditions_RouteType{
		Metadata: md,
	}
	val := parent.RouteType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/call-policy in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath extracts the value of the leaf CallPolicy from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_CallPolicyPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.CallPolicy
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) Get(t testing.TB) oc.E_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPathAny) Get(t testing.TB) []oc.E_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_INSTALL_PROTOCOL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) Replace(t testing.TB, val oc.E_PolicyTypes_INSTALL_PROTOCOL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_INSTALL_PROTOCOL_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) Update(t testing.TB, val oc.E_PolicyTypes_INSTALL_PROTOCOL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/config/install-protocol-eq in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_INSTALL_PROTOCOL_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath extracts the value of the leaf InstallProtocolEq from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_InstallProtocolEqPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions) *oc.QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_INSTALL_PROTOCOL_TYPE{
		Metadata: md,
	}
	val := parent.InstallProtocolEq
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/interface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath extracts the value of the leaf Interface from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-interface/config/subinterface in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchInterface) *oc.QualifiedUint32 {
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

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsRestrictedType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Replace(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) Update(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-neighbor-set/config/neighbor-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath extracts the value of the leaf NeighborSet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet_NeighborSetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchNeighborSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.NeighborSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsRestrictedType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Replace(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) Update(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-prefix-set/config/prefix-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath extracts the value of the leaf PrefixSet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet_PrefixSetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchPrefixSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.PrefixSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Lookup(t testing.TB) *oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Get(t testing.TB) *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny) Lookup(t testing.TB) []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPathAny) Get(t testing.TB) []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Replace(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) Update(t testing.TB, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) Lookup(t testing.TB) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath(t, md, goStruct)
	}
	return (&oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}).SetVal(goStruct.GetMatchSetOptions())
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) Get(t testing.TB) oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPathAny) Get(t testing.TB) []oc.E_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_PolicyTypes_MatchSetOptionsRestrictedType
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) Replace(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) Update(t testing.TB, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/match-set-options in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_PolicyTypes_MatchSetOptionsRestrictedType) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath extracts the value of the leaf MatchSetOptions from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_MatchSetOptionsPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) *oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType {
	t.Helper()
	qv := &oc.QualifiedE_PolicyTypes_MatchSetOptionsRestrictedType{
		Metadata: md,
	}
	val := parent.MatchSetOptions
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/conditions/match-tag-set/config/tag-set in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath extracts the value of the leaf TagSet from its parent oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet_TagSetPath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement_Conditions_MatchTagSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.TagSet
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement{}
	md, ok := oc.Lookup(t, n, "RoutingPolicy_PolicyDefinition_Statement", goStruct, true, true)
	if ok {
		return convertRoutingPolicy_PolicyDefinition_Statement_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.RoutingPolicy_PolicyDefinition_Statement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "RoutingPolicy_PolicyDefinition_Statement", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertRoutingPolicy_PolicyDefinition_Statement_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name with a ONCE subscription.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-routing-policy/routing-policy/policy-definitions/policy-definition/statements/statement/config/name in the given batch object.
func (n *RoutingPolicy_PolicyDefinition_Statement_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertRoutingPolicy_PolicyDefinition_Statement_NamePath extracts the value of the leaf Name from its parent oc.RoutingPolicy_PolicyDefinition_Statement
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertRoutingPolicy_PolicyDefinition_Statement_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.RoutingPolicy_PolicyDefinition_Statement) *oc.QualifiedString {
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
