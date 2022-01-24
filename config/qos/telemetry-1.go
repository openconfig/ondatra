package qos

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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-address in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_Ipv6_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Ipv6", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/ipv6/config/source-flow-label in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath extracts the value of the leaf SourceFlowLabel from its parent oc.Qos_Classifier_Term_Conditions_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertQos_Classifier_Term_Conditions_Ipv6_SourceFlowLabelPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.SourceFlowLabel
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2Path) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_L2{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2Path) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_L2 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2PathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_L2 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_L2
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_L2{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2PathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_L2 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_L2
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2.
func (n *Qos_Classifier_Term_Conditions_L2Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2.
func (n *Qos_Classifier_Term_Conditions_L2Path) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_L2) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_L2) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2.
func (n *Qos_Classifier_Term_Conditions_L2Path) Update(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_L2) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2 in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_L2) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac-mask in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacMaskPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath extracts the value of the leaf DestinationMacMask from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_DestinationMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_DestinationMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_DestinationMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/destination-mac in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_DestinationMacPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_L2_DestinationMacPath extracts the value of the leaf DestinationMac from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_DestinationMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.DestinationMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_EthertypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_EthertypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Replace(t testing.TB, val oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) Update(t testing.TB, val oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/ethertype in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_EthertypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_L2_Ethertype_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_L2_EthertypePath extracts the value of the leaf Ethertype from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union.
func convertQos_Classifier_Term_Conditions_L2_EthertypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_L2_Ethertype_Union{
		Metadata: md,
	}
	val := parent.Ethertype
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac-mask in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacMaskPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath extracts the value of the leaf SourceMacMask from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_SourceMacMaskPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMacMask
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_L2", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_L2_SourceMacPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_L2{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_L2", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_L2_SourceMacPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/l2/config/source-mac in the given batch object.
func (n *Qos_Classifier_Term_Conditions_L2_SourceMacPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_L2_SourceMacPath extracts the value of the leaf SourceMac from its parent oc.Qos_Classifier_Term_Conditions_L2
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_Conditions_L2_SourceMacPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_L2) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.SourceMac
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Mpls{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_MplsPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Mpls {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Mpls
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls in the given batch object.
func (n *Qos_Classifier_Term_Conditions_MplsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Mpls) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls in the given batch object.
func (n *Qos_Classifier_Term_Conditions_MplsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Mpls) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls.
func (n *Qos_Classifier_Term_Conditions_MplsPath) Update(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Mpls) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls in the given batch object.
func (n *Qos_Classifier_Term_Conditions_MplsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Mpls) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Replace(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) Update(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/end-label-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_EndLabelValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath extracts the value of the leaf EndLabelValue from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union.
func convertQos_Classifier_Term_Conditions_Mpls_EndLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Mpls_EndLabelValue_Union{
		Metadata: md,
	}
	val := parent.EndLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Replace(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) Update(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/start-label-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_StartLabelValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath extracts the value of the leaf StartLabelValue from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union.
func convertQos_Classifier_Term_Conditions_Mpls_StartLabelValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Mpls_StartLabelValue_Union{
		Metadata: md,
	}
	val := parent.StartLabelValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/traffic-class in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TrafficClassPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath extracts the value of the leaf TrafficClass from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Mpls_TrafficClassPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.TrafficClass
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Mpls", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Mpls{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Mpls", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/mpls/config/ttl-value in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Mpls_TtlValuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath extracts the value of the leaf TtlValue from its parent oc.Qos_Classifier_Term_Conditions_Mpls
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_Classifier_Term_Conditions_Mpls_TtlValuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Mpls) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.TtlValue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Classifier_Term_Conditions_Transport{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Get(t testing.TB) *oc.Qos_Classifier_Term_Conditions_Transport {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_TransportPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Transport {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Transport
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Classifier_Term_Conditions_Transport{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_TransportPathAny) Get(t testing.TB) []*oc.Qos_Classifier_Term_Conditions_Transport {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Classifier_Term_Conditions_Transport
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport in the given batch object.
func (n *Qos_Classifier_Term_Conditions_TransportPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Replace(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Transport) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport in the given batch object.
func (n *Qos_Classifier_Term_Conditions_TransportPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Transport) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport.
func (n *Qos_Classifier_Term_Conditions_TransportPath) Update(t testing.TB, val *oc.Qos_Classifier_Term_Conditions_Transport) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport in the given batch object.
func (n *Qos_Classifier_Term_Conditions_TransportPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Classifier_Term_Conditions_Transport) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Replace(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) Update(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/destination-port in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_DestinationPortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Transport_DestinationPort_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath extracts the value of the leaf DestinationPort from its parent oc.Qos_Classifier_Term_Conditions_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union.
func convertQos_Classifier_Term_Conditions_Transport_DestinationPortPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Transport) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Transport_DestinationPort_Union{
		Metadata: md,
	}
	val := parent.DestinationPort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Lookup(t testing.TB) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Get(t testing.TB) oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPathAny) Get(t testing.TB) []oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Replace(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) Update(t testing.TB, val oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/source-port in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_SourcePortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.Qos_Classifier_Term_Conditions_Transport_SourcePort_Union) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Transport_SourcePortPath extracts the value of the leaf SourcePort from its parent oc.Qos_Classifier_Term_Conditions_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union.
func convertQos_Classifier_Term_Conditions_Transport_SourcePortPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Transport) *oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union {
	t.Helper()
	qv := &oc.QualifiedQos_Classifier_Term_Conditions_Transport_SourcePort_Union{
		Metadata: md,
	}
	val := parent.SourcePort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Lookup(t testing.TB) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term_Conditions_Transport", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Get(t testing.TB) []oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny) Lookup(t testing.TB) []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term_Conditions_Transport{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term_Conditions_Transport", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags with a ONCE subscription.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPathAny) Get(t testing.TB) [][]oc.E_PacketMatchTypes_TCP_FLAGS {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.E_PacketMatchTypes_TCP_FLAGS
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Replace(t testing.TB, val []oc.E_PacketMatchTypes_TCP_FLAGS) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.E_PacketMatchTypes_TCP_FLAGS) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) Update(t testing.TB, val []oc.E_PacketMatchTypes_TCP_FLAGS) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/conditions/transport/config/tcp-flags in the given batch object.
func (n *Qos_Classifier_Term_Conditions_Transport_TcpFlagsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.E_PacketMatchTypes_TCP_FLAGS) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath extracts the value of the leaf TcpFlags from its parent oc.Qos_Classifier_Term_Conditions_Transport
// and combines the update with an existing Metadata to return a *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice.
func convertQos_Classifier_Term_Conditions_Transport_TcpFlagsPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term_Conditions_Transport) *oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice {
	t.Helper()
	qv := &oc.QualifiedE_PacketMatchTypes_TCP_FLAGSSlice{
		Metadata: md,
	}
	val := parent.TcpFlags
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_Term_IdPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Classifier_Term{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier_Term", goStruct, true, true)
	if ok {
		return convertQos_Classifier_Term_IdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_Term_IdPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_Term_IdPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier_Term{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier_Term", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_Term_IdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id with a ONCE subscription.
func (n *Qos_Classifier_Term_IdPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id.
func (n *Qos_Classifier_Term_IdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id in the given batch object.
func (n *Qos_Classifier_Term_IdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id.
func (n *Qos_Classifier_Term_IdPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id in the given batch object.
func (n *Qos_Classifier_Term_IdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id.
func (n *Qos_Classifier_Term_IdPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/terms/term/config/id in the given batch object.
func (n *Qos_Classifier_Term_IdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Classifier_Term_IdPath extracts the value of the leaf Id from its parent oc.Qos_Classifier_Term
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Classifier_Term_IdPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier_Term) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/classifiers/classifier/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Classifier_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Qos_Classifier_Type {
	t.Helper()
	goStruct := &oc.Qos_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Classifier", goStruct, true, true)
	if ok {
		return convertQos_Classifier_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/classifiers/classifier/config/type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Classifier_TypePath) Get(t testing.TB) oc.E_Qos_Classifier_Type {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/classifiers/classifier/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Classifier_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Qos_Classifier_Type {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Qos_Classifier_Type
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Classifier", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Classifier_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/classifiers/classifier/config/type with a ONCE subscription.
func (n *Qos_Classifier_TypePathAny) Get(t testing.TB) []oc.E_Qos_Classifier_Type {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Qos_Classifier_Type
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/classifiers/classifier/config/type.
func (n *Qos_Classifier_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Classifier_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/classifiers/classifier/config/type.
func (n *Qos_Classifier_TypePath) Replace(t testing.TB, val oc.E_Qos_Classifier_Type) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Classifier_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Qos_Classifier_Type) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/classifiers/classifier/config/type.
func (n *Qos_Classifier_TypePath) Update(t testing.TB, val oc.E_Qos_Classifier_Type) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Classifier_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Qos_Classifier_Type) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Classifier_TypePath extracts the value of the leaf Type from its parent oc.Qos_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Qos_Classifier_Type.
func convertQos_Classifier_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Classifier) *oc.QualifiedE_Qos_Classifier_Type {
	t.Helper()
	qv := &oc.QualifiedE_Qos_Classifier_Type{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroupPath) Lookup(t testing.TB) *oc.QualifiedQos_ForwardingGroup {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_ForwardingGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroupPath) Get(t testing.TB) *oc.Qos_ForwardingGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_ForwardingGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_ForwardingGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_ForwardingGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group with a ONCE subscription.
func (n *Qos_ForwardingGroupPathAny) Get(t testing.TB) []*oc.Qos_ForwardingGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_ForwardingGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group.
func (n *Qos_ForwardingGroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/forwarding-groups/forwarding-group in the given batch object.
func (n *Qos_ForwardingGroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group.
func (n *Qos_ForwardingGroupPath) Replace(t testing.TB, val *oc.Qos_ForwardingGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/forwarding-groups/forwarding-group in the given batch object.
func (n *Qos_ForwardingGroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_ForwardingGroup) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group.
func (n *Qos_ForwardingGroupPath) Update(t testing.TB, val *oc.Qos_ForwardingGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/forwarding-groups/forwarding-group in the given batch object.
func (n *Qos_ForwardingGroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_ForwardingGroup) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, true)
	if ok {
		return convertQos_ForwardingGroup_FabricPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_FabricPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_FabricPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority with a ONCE subscription.
func (n *Qos_ForwardingGroup_FabricPriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority in the given batch object.
func (n *Qos_ForwardingGroup_FabricPriorityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority in the given batch object.
func (n *Qos_ForwardingGroup_FabricPriorityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority.
func (n *Qos_ForwardingGroup_FabricPriorityPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/fabric-priority in the given batch object.
func (n *Qos_ForwardingGroup_FabricPriorityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_ForwardingGroup_FabricPriorityPath extracts the value of the leaf FabricPriority from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertQos_ForwardingGroup_FabricPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.FabricPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, true)
	if ok {
		return convertQos_ForwardingGroup_MulticastOutputQueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_MulticastOutputQueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue with a ONCE subscription.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue in the given batch object.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue in the given batch object.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/multicast-output-queue in the given batch object.
func (n *Qos_ForwardingGroup_MulticastOutputQueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_ForwardingGroup_MulticastOutputQueuePath extracts the value of the leaf MulticastOutputQueue from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_MulticastOutputQueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MulticastOutputQueue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, true)
	if ok {
		return convertQos_ForwardingGroup_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name with a ONCE subscription.
func (n *Qos_ForwardingGroup_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name.
func (n *Qos_ForwardingGroup_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name in the given batch object.
func (n *Qos_ForwardingGroup_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name.
func (n *Qos_ForwardingGroup_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name in the given batch object.
func (n *Qos_ForwardingGroup_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name.
func (n *Qos_ForwardingGroup_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/name in the given batch object.
func (n *Qos_ForwardingGroup_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_ForwardingGroup_NamePath extracts the value of the leaf Name from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_OutputQueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, true)
	if ok {
		return convertQos_ForwardingGroup_OutputQueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_OutputQueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_OutputQueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_OutputQueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue with a ONCE subscription.
func (n *Qos_ForwardingGroup_OutputQueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue.
func (n *Qos_ForwardingGroup_OutputQueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue in the given batch object.
func (n *Qos_ForwardingGroup_OutputQueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue.
func (n *Qos_ForwardingGroup_OutputQueuePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue in the given batch object.
func (n *Qos_ForwardingGroup_OutputQueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue.
func (n *Qos_ForwardingGroup_OutputQueuePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/output-queue in the given batch object.
func (n *Qos_ForwardingGroup_OutputQueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_ForwardingGroup_OutputQueuePath extracts the value of the leaf OutputQueue from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_OutputQueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.OutputQueue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_ForwardingGroup{}
	md, ok := oc.Lookup(t, n, "Qos_ForwardingGroup", goStruct, true, true)
	if ok {
		return convertQos_ForwardingGroup_UnicastOutputQueuePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_ForwardingGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_ForwardingGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_ForwardingGroup_UnicastOutputQueuePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue with a ONCE subscription.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue in the given batch object.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue in the given batch object.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/forwarding-groups/forwarding-group/config/unicast-output-queue in the given batch object.
func (n *Qos_ForwardingGroup_UnicastOutputQueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_ForwardingGroup_UnicastOutputQueuePath extracts the value of the leaf UnicastOutputQueue from its parent oc.Qos_ForwardingGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_ForwardingGroup_UnicastOutputQueuePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_ForwardingGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.UnicastOutputQueue
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_InterfacePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface {
	t.Helper()
	goStruct := &oc.Qos_Interface{}
	md, ok := oc.Lookup(t, n, "Qos_Interface", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_InterfacePath) Get(t testing.TB) *oc.Qos_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface with a ONCE subscription.
func (n *Qos_InterfacePathAny) Get(t testing.TB) []*oc.Qos_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface.
func (n *Qos_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface in the given batch object.
func (n *Qos_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface.
func (n *Qos_InterfacePath) Replace(t testing.TB, val *oc.Qos_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface in the given batch object.
func (n *Qos_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface.
func (n *Qos_InterfacePath) Update(t testing.TB, val *oc.Qos_Interface) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface in the given batch object.
func (n *Qos_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_InputPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Input{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_InputPath) Get(t testing.TB) *oc.Qos_Interface_Input {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_InputPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input with a ONCE subscription.
func (n *Qos_Interface_InputPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input.
func (n *Qos_Interface_InputPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input in the given batch object.
func (n *Qos_Interface_InputPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input.
func (n *Qos_Interface_InputPath) Replace(t testing.TB, val *oc.Qos_Interface_Input) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input in the given batch object.
func (n *Qos_Interface_InputPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input.
func (n *Qos_Interface_InputPath) Update(t testing.TB, val *oc.Qos_Interface_Input) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input in the given batch object.
func (n *Qos_Interface_InputPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_BufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_BufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_BufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Input_BufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/config/buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_BufferAllocationProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_BufferAllocationProfilePath extracts the value of the leaf BufferAllocationProfile from its parent oc.Qos_Interface_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_BufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.BufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_ClassifierPath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_Classifier {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_ClassifierPath) Get(t testing.TB) *oc.Qos_Interface_Input_Classifier {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_ClassifierPathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_Classifier {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_Classifier
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_Classifier{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier with a ONCE subscription.
func (n *Qos_Interface_Input_ClassifierPathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_Classifier {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_Classifier
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier.
func (n *Qos_Interface_Input_ClassifierPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier in the given batch object.
func (n *Qos_Interface_Input_ClassifierPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier.
func (n *Qos_Interface_Input_ClassifierPath) Replace(t testing.TB, val *oc.Qos_Interface_Input_Classifier) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier in the given batch object.
func (n *Qos_Interface_Input_ClassifierPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_Classifier) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier.
func (n *Qos_Interface_Input_ClassifierPath) Update(t testing.TB, val *oc.Qos_Interface_Input_Classifier) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier in the given batch object.
func (n *Qos_Interface_Input_ClassifierPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_Classifier) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_Classifier_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Classifier_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name.
func (n *Qos_Interface_Input_Classifier_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Interface_Input_Classifier_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name.
func (n *Qos_Interface_Input_Classifier_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Interface_Input_Classifier_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name.
func (n *Qos_Interface_Input_Classifier_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/name in the given batch object.
func (n *Qos_Interface_Input_Classifier_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_Classifier_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_Classifier_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Classifier) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Classifier_TypePath) Lookup(t testing.TB) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Classifier{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Classifier", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_Classifier_TypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Classifier_TypePath) Get(t testing.TB) oc.E_Input_Classifier_Type {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Classifier_TypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Input_Classifier_Type
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Classifier{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Classifier", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Classifier_TypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type with a ONCE subscription.
func (n *Qos_Interface_Input_Classifier_TypePathAny) Get(t testing.TB) []oc.E_Input_Classifier_Type {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Input_Classifier_Type
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type.
func (n *Qos_Interface_Input_Classifier_TypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Interface_Input_Classifier_TypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type.
func (n *Qos_Interface_Input_Classifier_TypePath) Replace(t testing.TB, val oc.E_Input_Classifier_Type) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Interface_Input_Classifier_TypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Input_Classifier_Type) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type.
func (n *Qos_Interface_Input_Classifier_TypePath) Update(t testing.TB, val oc.E_Input_Classifier_Type) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/classifiers/classifier/config/type in the given batch object.
func (n *Qos_Interface_Input_Classifier_TypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Input_Classifier_Type) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertQos_Interface_Input_Classifier_TypePath extracts the value of the leaf Type from its parent oc.Qos_Interface_Input_Classifier
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Input_Classifier_Type.
func convertQos_Interface_Input_Classifier_TypePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Classifier) *oc.QualifiedE_Input_Classifier_Type {
	t.Helper()
	qv := &oc.QualifiedE_Input_Classifier_Type{
		Metadata: md,
	}
	val := parent.Type
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_MulticastBufferAllocationProfilePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_MulticastBufferAllocationProfilePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile with a ONCE subscription.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/config/multicast-buffer-allocation-profile in the given batch object.
func (n *Qos_Interface_Input_MulticastBufferAllocationProfilePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_MulticastBufferAllocationProfilePath extracts the value of the leaf MulticastBufferAllocationProfile from its parent oc.Qos_Interface_Input
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_MulticastBufferAllocationProfilePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MulticastBufferAllocationProfile
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_QueuePath) Lookup(t testing.TB) *oc.QualifiedQos_Interface_Input_Queue {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, false, true)
	if ok {
		return (&oc.QualifiedQos_Interface_Input_Queue{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_QueuePath) Get(t testing.TB) *oc.Qos_Interface_Input_Queue {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_QueuePathAny) Lookup(t testing.TB) []*oc.QualifiedQos_Interface_Input_Queue {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedQos_Interface_Input_Queue
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedQos_Interface_Input_Queue{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue with a ONCE subscription.
func (n *Qos_Interface_Input_QueuePathAny) Get(t testing.TB) []*oc.Qos_Interface_Input_Queue {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Qos_Interface_Input_Queue
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue.
func (n *Qos_Interface_Input_QueuePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue in the given batch object.
func (n *Qos_Interface_Input_QueuePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue.
func (n *Qos_Interface_Input_QueuePath) Replace(t testing.TB, val *oc.Qos_Interface_Input_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue in the given batch object.
func (n *Qos_Interface_Input_QueuePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_Queue) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue.
func (n *Qos_Interface_Input_QueuePath) Update(t testing.TB, val *oc.Qos_Interface_Input_Queue) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue in the given batch object.
func (n *Qos_Interface_Input_QueuePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Qos_Interface_Input_Queue) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Qos_Interface_Input_Queue_NamePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Qos_Interface_Input_Queue{}
	md, ok := oc.Lookup(t, n, "Qos_Interface_Input_Queue", goStruct, true, true)
	if ok {
		return convertQos_Interface_Input_Queue_NamePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Qos_Interface_Input_Queue_NamePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Qos_Interface_Input_Queue_NamePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Qos_Interface_Input_Queue{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Qos_Interface_Input_Queue", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertQos_Interface_Input_Queue_NamePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name with a ONCE subscription.
func (n *Qos_Interface_Input_Queue_NamePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name.
func (n *Qos_Interface_Input_Queue_NamePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Input_Queue_NamePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name.
func (n *Qos_Interface_Input_Queue_NamePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Input_Queue_NamePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name.
func (n *Qos_Interface_Input_Queue_NamePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-qos/qos/interfaces/interface/input/queues/queue/config/name in the given batch object.
func (n *Qos_Interface_Input_Queue_NamePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertQos_Interface_Input_Queue_NamePath extracts the value of the leaf Name from its parent oc.Qos_Interface_Input_Queue
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertQos_Interface_Input_Queue_NamePath(t testing.TB, md *genutil.Metadata, parent *oc.Qos_Interface_Input_Queue) *oc.QualifiedString {
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
