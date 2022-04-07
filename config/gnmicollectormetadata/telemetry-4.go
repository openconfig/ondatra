package gnmicollectormetadata

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

// Lookup fetches the value at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Meta_Window_MinPath) Lookup(t testing.TB) *oc.QualifiedInt64 {
	t.Helper()
	goStruct := &oc.Meta_Window{}
	md, ok := oc.Lookup(t, n, "Meta_Window", goStruct, true, true)
	if ok {
		return convertMeta_Window_MinPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Meta_Window_MinPath) Get(t testing.TB) int64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Meta_Window_MinPathAny) Lookup(t testing.TB) []*oc.QualifiedInt64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Meta_Window{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Meta_Window", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertMeta_Window_MinPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /gnmi-collector-metadata/meta/latency/window/min with a ONCE subscription.
func (n *Meta_Window_MinPathAny) Get(t testing.TB) []int64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /gnmi-collector-metadata/meta/latency/window/min.
func (n *Meta_Window_MinPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /gnmi-collector-metadata/meta/latency/window/min in the given batch object.
func (n *Meta_Window_MinPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /gnmi-collector-metadata/meta/latency/window/min.
func (n *Meta_Window_MinPath) Replace(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /gnmi-collector-metadata/meta/latency/window/min in the given batch object.
func (n *Meta_Window_MinPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /gnmi-collector-metadata/meta/latency/window/min.
func (n *Meta_Window_MinPath) Update(t testing.TB, val int64) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /gnmi-collector-metadata/meta/latency/window/min in the given batch object.
func (n *Meta_Window_MinPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int64) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertMeta_Window_MinPath extracts the value of the leaf Min from its parent oc.Meta_Window
// and combines the update with an existing Metadata to return a *oc.QualifiedInt64.
func convertMeta_Window_MinPath(t testing.TB, md *genutil.Metadata, parent *oc.Meta_Window) *oc.QualifiedInt64 {
	t.Helper()
	qv := &oc.QualifiedInt64{
		Metadata: md,
	}
	val := parent.Min
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
