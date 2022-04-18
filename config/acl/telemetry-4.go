package acl

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

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, true, true)
	if ok {
		return convertAcl_Interface_EgressAclSet_SetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_SetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_SetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name in the given batch object.
func (n *Acl_Interface_EgressAclSet_SetNamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name in the given batch object.
func (n *Acl_Interface_EgressAclSet_SetNamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name.
func (n *Acl_Interface_EgressAclSet_SetNamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/set-name in the given batch object.
func (n *Acl_Interface_EgressAclSet_SetNamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_Interface_EgressAclSet_SetNamePath extracts the value of the leaf SetName from its parent oc.Acl_Interface_EgressAclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_Interface_EgressAclSet_SetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, true, true)
	if ok {
		return convertAcl_Interface_EgressAclSet_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_TypePath) Get(t testing.TB) oc.E_Acl_ACL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_ACL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_EgressAclSet_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_TypePathAny) Get(t testing.TB) []oc.E_Acl_ACL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_ACL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type.
func (n *Acl_Interface_EgressAclSet_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type in the given batch object.
func (n *Acl_Interface_EgressAclSet_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type.
func (n *Acl_Interface_EgressAclSet_TypePath) Replace(t testing.TB, val oc.E_Acl_ACL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type in the given batch object.
func (n *Acl_Interface_EgressAclSet_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_ACL_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type.
func (n *Acl_Interface_EgressAclSet_TypePath) Update(t testing.TB, val oc.E_Acl_ACL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/config/type in the given batch object.
func (n *Acl_Interface_EgressAclSet_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_ACL_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_Interface_EgressAclSet_TypePath extracts the value of the leaf Type from its parent oc.Acl_Interface_EgressAclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_ACL_TYPE.
func convertAcl_Interface_EgressAclSet_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_EgressAclSet) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_Acl_ACL_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/config/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_Interface{}
	md, ok := oc.Lookup(t, n, "Acl_Interface", goStruct, true, true)
	if ok {
		return convertAcl_Interface_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/config/id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/config/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/config/id with a ONCE subscription.
func (n *Acl_Interface_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/config/id.
func (n *Acl_Interface_IdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/config/id in the given batch object.
func (n *Acl_Interface_IdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/config/id.
func (n *Acl_Interface_IdPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/config/id in the given batch object.
func (n *Acl_Interface_IdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/config/id.
func (n *Acl_Interface_IdPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/config/id in the given batch object.
func (n *Acl_Interface_IdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_Interface_IdPath extracts the value of the leaf Id from its parent oc.Acl_Interface
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_Interface_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Id
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_IngressAclSetPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_IngressAclSet {
	t.Helper()
	goStruct := &oc.Acl_Interface_IngressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_IngressAclSet", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_Interface_IngressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_IngressAclSetPath) Get(t testing.TB) *oc.Acl_Interface_IngressAclSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_IngressAclSetPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_IngressAclSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_IngressAclSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_IngressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_IngressAclSet", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_IngressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set with a ONCE subscription.
func (n *Acl_Interface_IngressAclSetPathAny) Get(t testing.TB) []*oc.Acl_Interface_IngressAclSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_IngressAclSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set.
func (n *Acl_Interface_IngressAclSetPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set in the given batch object.
func (n *Acl_Interface_IngressAclSetPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set.
func (n *Acl_Interface_IngressAclSetPath) Replace(t testing.TB, val *oc.Acl_Interface_IngressAclSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set in the given batch object.
func (n *Acl_Interface_IngressAclSetPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface_IngressAclSet) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set.
func (n *Acl_Interface_IngressAclSetPath) Update(t testing.TB, val *oc.Acl_Interface_IngressAclSet) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set in the given batch object.
func (n *Acl_Interface_IngressAclSetPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface_IngressAclSet) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_IngressAclSet_SetNamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_Interface_IngressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_IngressAclSet", goStruct, true, true)
	if ok {
		return convertAcl_Interface_IngressAclSet_SetNamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_IngressAclSet_SetNamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_IngressAclSet_SetNamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_IngressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_IngressAclSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_IngressAclSet_SetNamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name with a ONCE subscription.
func (n *Acl_Interface_IngressAclSet_SetNamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name.
func (n *Acl_Interface_IngressAclSet_SetNamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name in the given batch object.
func (n *Acl_Interface_IngressAclSet_SetNamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name.
func (n *Acl_Interface_IngressAclSet_SetNamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name in the given batch object.
func (n *Acl_Interface_IngressAclSet_SetNamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name.
func (n *Acl_Interface_IngressAclSet_SetNamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/set-name in the given batch object.
func (n *Acl_Interface_IngressAclSet_SetNamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_Interface_IngressAclSet_SetNamePath extracts the value of the leaf SetName from its parent oc.Acl_Interface_IngressAclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_Interface_IngressAclSet_SetNamePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_IngressAclSet) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SetName
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_IngressAclSet_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	goStruct := &oc.Acl_Interface_IngressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_IngressAclSet", goStruct, true, true)
	if ok {
		return convertAcl_Interface_IngressAclSet_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_IngressAclSet_TypePath) Get(t testing.TB) oc.E_Acl_ACL_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_IngressAclSet_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_ACL_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_IngressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_IngressAclSet", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_IngressAclSet_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type with a ONCE subscription.
func (n *Acl_Interface_IngressAclSet_TypePathAny) Get(t testing.TB) []oc.E_Acl_ACL_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_ACL_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type.
func (n *Acl_Interface_IngressAclSet_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type in the given batch object.
func (n *Acl_Interface_IngressAclSet_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type.
func (n *Acl_Interface_IngressAclSet_TypePath) Replace(t testing.TB, val oc.E_Acl_ACL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type in the given batch object.
func (n *Acl_Interface_IngressAclSet_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_ACL_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type.
func (n *Acl_Interface_IngressAclSet_TypePath) Update(t testing.TB, val oc.E_Acl_ACL_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/ingress-acl-sets/ingress-acl-set/config/type in the given batch object.
func (n *Acl_Interface_IngressAclSet_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Acl_ACL_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertAcl_Interface_IngressAclSet_TypePath extracts the value of the leaf Type from its parent oc.Acl_Interface_IngressAclSet
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_ACL_TYPE.
func convertAcl_Interface_IngressAclSet_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_IngressAclSet) *oc.QualifiedE_Acl_ACL_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_Acl_ACL_TYPE{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_InterfaceRef {
	t.Helper()
	goStruct := &oc.Acl_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_InterfaceRef", goStruct, false, true)
	if ok {
		return (&oc.QualifiedAcl_Interface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_InterfaceRefPath) Get(t testing.TB) *oc.Acl_Interface_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_InterfaceRef", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/interface-ref with a ONCE subscription.
func (n *Acl_Interface_InterfaceRefPathAny) Get(t testing.TB) []*oc.Acl_Interface_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref.
func (n *Acl_Interface_InterfaceRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/interface-ref in the given batch object.
func (n *Acl_Interface_InterfaceRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref.
func (n *Acl_Interface_InterfaceRefPath) Replace(t testing.TB, val *oc.Acl_Interface_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/interface-ref in the given batch object.
func (n *Acl_Interface_InterfaceRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface_InterfaceRef) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref.
func (n *Acl_Interface_InterfaceRefPath) Update(t testing.TB, val *oc.Acl_Interface_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/interface-ref in the given batch object.
func (n *Acl_Interface_InterfaceRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Acl_Interface_InterfaceRef) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Acl_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_InterfaceRef", goStruct, true, true)
	if ok {
		return convertAcl_Interface_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface with a ONCE subscription.
func (n *Acl_Interface_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface.
func (n *Acl_Interface_InterfaceRef_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface in the given batch object.
func (n *Acl_Interface_InterfaceRef_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface.
func (n *Acl_Interface_InterfaceRef_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface in the given batch object.
func (n *Acl_Interface_InterfaceRef_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface.
func (n *Acl_Interface_InterfaceRef_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/interface-ref/config/interface in the given batch object.
func (n *Acl_Interface_InterfaceRef_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_Interface_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Acl_Interface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertAcl_Interface_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_InterfaceRef) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Acl_Interface_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_InterfaceRef", goStruct, true, true)
	if ok {
		return convertAcl_Interface_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertAcl_Interface_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface with a ONCE subscription.
func (n *Acl_Interface_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface in the given batch object.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface in the given batch object.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-acl/acl/interfaces/interface/interface-ref/config/subinterface in the given batch object.
func (n *Acl_Interface_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertAcl_Interface_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Acl_Interface_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertAcl_Interface_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Acl_Interface_InterfaceRef) *oc.QualifiedUint32 {
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
