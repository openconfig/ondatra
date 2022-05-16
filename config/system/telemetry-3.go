package system

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

// Lookup fetches the value at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_Console_SelectorPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_Console_Selector {
	t.Helper()
	goStruct := &oc.System_Logging_Console_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console_Selector", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Logging_Console_Selector{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_Console_SelectorPath) Get(t testing.TB) *oc.System_Logging_Console_Selector {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_Console_SelectorPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_Console_Selector {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_Console_Selector
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console_Selector", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_Console_Selector{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console/selectors/selector with a ONCE subscription.
func (n *System_Logging_Console_SelectorPathAny) Get(t testing.TB) []*oc.System_Logging_Console_Selector {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_Console_Selector
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/console/selectors/selector.
func (n *System_Logging_Console_SelectorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/console/selectors/selector in the given batch object.
func (n *System_Logging_Console_SelectorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/console/selectors/selector.
func (n *System_Logging_Console_SelectorPath) Replace(t testing.TB, val *oc.System_Logging_Console_Selector) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/console/selectors/selector in the given batch object.
func (n *System_Logging_Console_SelectorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_Console_Selector) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/console/selectors/selector.
func (n *System_Logging_Console_SelectorPath) Update(t testing.TB, val *oc.System_Logging_Console_Selector) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/console/selectors/selector in the given batch object.
func (n *System_Logging_Console_SelectorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_Console_Selector) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/logging/console/selectors/selector/config/facility with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_Console_Selector_FacilityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	goStruct := &oc.System_Logging_Console_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console_Selector", goStruct, true, true)
	if ok {
		return convertSystem_Logging_Console_Selector_FacilityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console/selectors/selector/config/facility with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_Console_Selector_FacilityPath) Get(t testing.TB) oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console/selectors/selector/config/facility with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_Console_Selector_FacilityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console_Selector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_Console_Selector_FacilityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console/selectors/selector/config/facility with a ONCE subscription.
func (n *System_Logging_Console_Selector_FacilityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SYSLOG_FACILITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/console/selectors/selector/config/facility.
func (n *System_Logging_Console_Selector_FacilityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/console/selectors/selector/config/facility in the given batch object.
func (n *System_Logging_Console_Selector_FacilityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/console/selectors/selector/config/facility.
func (n *System_Logging_Console_Selector_FacilityPath) Replace(t testing.TB, val oc.E_SystemLogging_SYSLOG_FACILITY) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/console/selectors/selector/config/facility in the given batch object.
func (n *System_Logging_Console_Selector_FacilityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SYSLOG_FACILITY) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/console/selectors/selector/config/facility.
func (n *System_Logging_Console_Selector_FacilityPath) Update(t testing.TB, val oc.E_SystemLogging_SYSLOG_FACILITY) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/console/selectors/selector/config/facility in the given batch object.
func (n *System_Logging_Console_Selector_FacilityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SYSLOG_FACILITY) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Logging_Console_Selector_FacilityPath extracts the value of the leaf Facility from its parent oc.System_Logging_Console_Selector
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY.
func convertSystem_Logging_Console_Selector_FacilityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_Console_Selector) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SYSLOG_FACILITY{
		Metadata: md,
	}
	val := parent.Facility
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/console/selectors/selector/config/severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_Console_Selector_SeverityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	goStruct := &oc.System_Logging_Console_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_Console_Selector", goStruct, true, true)
	if ok {
		return convertSystem_Logging_Console_Selector_SeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/console/selectors/selector/config/severity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_Console_Selector_SeverityPath) Get(t testing.TB) oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/console/selectors/selector/config/severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_Console_Selector_SeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SyslogSeverity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_Console_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_Console_Selector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_Console_Selector_SeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/console/selectors/selector/config/severity with a ONCE subscription.
func (n *System_Logging_Console_Selector_SeverityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SyslogSeverity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/console/selectors/selector/config/severity.
func (n *System_Logging_Console_Selector_SeverityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/console/selectors/selector/config/severity in the given batch object.
func (n *System_Logging_Console_Selector_SeverityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/console/selectors/selector/config/severity.
func (n *System_Logging_Console_Selector_SeverityPath) Replace(t testing.TB, val oc.E_SystemLogging_SyslogSeverity) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/console/selectors/selector/config/severity in the given batch object.
func (n *System_Logging_Console_Selector_SeverityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SyslogSeverity) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/console/selectors/selector/config/severity.
func (n *System_Logging_Console_Selector_SeverityPath) Update(t testing.TB, val oc.E_SystemLogging_SyslogSeverity) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/console/selectors/selector/config/severity in the given batch object.
func (n *System_Logging_Console_Selector_SeverityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SyslogSeverity) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Logging_Console_Selector_SeverityPath extracts the value of the leaf Severity from its parent oc.System_Logging_Console_Selector
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SyslogSeverity.
func convertSystem_Logging_Console_Selector_SeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_Console_Selector) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SyslogSeverity{
		Metadata: md,
	}
	val := parent.Severity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServerPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_RemoteServer {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Logging_RemoteServer{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServerPath) Get(t testing.TB) *oc.System_Logging_RemoteServer {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServerPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_RemoteServer {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_RemoteServer
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_RemoteServer{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server with a ONCE subscription.
func (n *System_Logging_RemoteServerPathAny) Get(t testing.TB) []*oc.System_Logging_RemoteServer {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_RemoteServer
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/remote-servers/remote-server.
func (n *System_Logging_RemoteServerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/remote-servers/remote-server in the given batch object.
func (n *System_Logging_RemoteServerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/remote-servers/remote-server.
func (n *System_Logging_RemoteServerPath) Replace(t testing.TB, val *oc.System_Logging_RemoteServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/remote-servers/remote-server in the given batch object.
func (n *System_Logging_RemoteServerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_RemoteServer) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/remote-servers/remote-server.
func (n *System_Logging_RemoteServerPath) Update(t testing.TB, val *oc.System_Logging_RemoteServer) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/remote-servers/remote-server in the given batch object.
func (n *System_Logging_RemoteServerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_RemoteServer) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/config/host with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_HostPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer", goStruct, true, true)
	if ok {
		return convertSystem_Logging_RemoteServer_HostPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/config/host with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_HostPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/config/host with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_HostPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_HostPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/config/host with a ONCE subscription.
func (n *System_Logging_RemoteServer_HostPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/host.
func (n *System_Logging_RemoteServer_HostPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/remote-servers/remote-server/config/host in the given batch object.
func (n *System_Logging_RemoteServer_HostPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/host.
func (n *System_Logging_RemoteServer_HostPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/remote-servers/remote-server/config/host in the given batch object.
func (n *System_Logging_RemoteServer_HostPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/host.
func (n *System_Logging_RemoteServer_HostPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/remote-servers/remote-server/config/host in the given batch object.
func (n *System_Logging_RemoteServer_HostPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Logging_RemoteServer_HostPath extracts the value of the leaf Host from its parent oc.System_Logging_RemoteServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Logging_RemoteServer_HostPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Host
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_RemotePortPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer", goStruct, true, true)
	if ok {
		return convertSystem_Logging_RemoteServer_RemotePortPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetRemotePort())
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_RemotePortPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_RemotePortPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_RemotePortPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port with a ONCE subscription.
func (n *System_Logging_RemoteServer_RemotePortPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port.
func (n *System_Logging_RemoteServer_RemotePortPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port in the given batch object.
func (n *System_Logging_RemoteServer_RemotePortPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port.
func (n *System_Logging_RemoteServer_RemotePortPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port in the given batch object.
func (n *System_Logging_RemoteServer_RemotePortPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port.
func (n *System_Logging_RemoteServer_RemotePortPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/remote-servers/remote-server/config/remote-port in the given batch object.
func (n *System_Logging_RemoteServer_RemotePortPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Logging_RemoteServer_RemotePortPath extracts the value of the leaf RemotePort from its parent oc.System_Logging_RemoteServer
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Logging_RemoteServer_RemotePortPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.RemotePort
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_SelectorPath) Lookup(t testing.TB) *oc.QualifiedSystem_Logging_RemoteServer_Selector {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer_Selector", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Logging_RemoteServer_Selector{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_SelectorPath) Get(t testing.TB) *oc.System_Logging_RemoteServer_Selector {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_SelectorPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Logging_RemoteServer_Selector {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Logging_RemoteServer_Selector
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer_Selector", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Logging_RemoteServer_Selector{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector with a ONCE subscription.
func (n *System_Logging_RemoteServer_SelectorPathAny) Get(t testing.TB) []*oc.System_Logging_RemoteServer_Selector {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Logging_RemoteServer_Selector
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector.
func (n *System_Logging_RemoteServer_SelectorPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector in the given batch object.
func (n *System_Logging_RemoteServer_SelectorPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector.
func (n *System_Logging_RemoteServer_SelectorPath) Replace(t testing.TB, val *oc.System_Logging_RemoteServer_Selector) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector in the given batch object.
func (n *System_Logging_RemoteServer_SelectorPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_RemoteServer_Selector) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector.
func (n *System_Logging_RemoteServer_SelectorPath) Update(t testing.TB, val *oc.System_Logging_RemoteServer_Selector) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector in the given batch object.
func (n *System_Logging_RemoteServer_SelectorPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Logging_RemoteServer_Selector) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer_Selector", goStruct, true, true)
	if ok {
		return convertSystem_Logging_RemoteServer_Selector_FacilityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Get(t testing.TB) oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_Selector_FacilityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SYSLOG_FACILITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer_Selector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_Selector_FacilityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility with a ONCE subscription.
func (n *System_Logging_RemoteServer_Selector_FacilityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SYSLOG_FACILITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility in the given batch object.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Replace(t testing.TB, val oc.E_SystemLogging_SYSLOG_FACILITY) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility in the given batch object.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SYSLOG_FACILITY) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) Update(t testing.TB, val oc.E_SystemLogging_SYSLOG_FACILITY) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/facility in the given batch object.
func (n *System_Logging_RemoteServer_Selector_FacilityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SYSLOG_FACILITY) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Logging_RemoteServer_Selector_FacilityPath extracts the value of the leaf Facility from its parent oc.System_Logging_RemoteServer_Selector
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY.
func convertSystem_Logging_RemoteServer_Selector_FacilityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer_Selector) *oc.QualifiedE_SystemLogging_SYSLOG_FACILITY {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SYSLOG_FACILITY{
		Metadata: md,
	}
	val := parent.Facility
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer_Selector{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer_Selector", goStruct, true, true)
	if ok {
		return convertSystem_Logging_RemoteServer_Selector_SeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) Get(t testing.TB) oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_Selector_SeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SyslogSeverity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer_Selector{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer_Selector", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_Selector_SeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity with a ONCE subscription.
func (n *System_Logging_RemoteServer_Selector_SeverityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SyslogSeverity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity in the given batch object.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) Replace(t testing.TB, val oc.E_SystemLogging_SyslogSeverity) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity in the given batch object.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SyslogSeverity) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) Update(t testing.TB, val oc.E_SystemLogging_SyslogSeverity) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/remote-servers/remote-server/selectors/selector/config/severity in the given batch object.
func (n *System_Logging_RemoteServer_Selector_SeverityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SyslogSeverity) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Logging_RemoteServer_Selector_SeverityPath extracts the value of the leaf Severity from its parent oc.System_Logging_RemoteServer_Selector
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SyslogSeverity.
func convertSystem_Logging_RemoteServer_Selector_SeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer_Selector) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SyslogSeverity{
		Metadata: md,
	}
	val := parent.Severity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Logging_RemoteServer_SourceAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System_Logging_RemoteServer{}
	md, ok := oc.Lookup(t, n, "System_Logging_RemoteServer", goStruct, true, true)
	if ok {
		return convertSystem_Logging_RemoteServer_SourceAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Logging_RemoteServer_SourceAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Logging_RemoteServer_SourceAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Logging_RemoteServer{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Logging_RemoteServer", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Logging_RemoteServer_SourceAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address with a ONCE subscription.
func (n *System_Logging_RemoteServer_SourceAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address.
func (n *System_Logging_RemoteServer_SourceAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address in the given batch object.
func (n *System_Logging_RemoteServer_SourceAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address.
func (n *System_Logging_RemoteServer_SourceAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address in the given batch object.
func (n *System_Logging_RemoteServer_SourceAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address.
func (n *System_Logging_RemoteServer_SourceAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/logging/remote-servers/remote-server/config/source-address in the given batch object.
func (n *System_Logging_RemoteServer_SourceAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Logging_RemoteServer_SourceAddressPath extracts the value of the leaf SourceAddress from its parent oc.System_Logging_RemoteServer
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_Logging_RemoteServer_SourceAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Logging_RemoteServer) *oc.QualifiedString {
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

// Lookup fetches the value at /openconfig-system/system/config/login-banner with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_LoginBannerPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, true)
	if ok {
		return convertSystem_LoginBannerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/config/login-banner with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_LoginBannerPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/config/login-banner with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_LoginBannerPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_LoginBannerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/config/login-banner with a ONCE subscription.
func (n *System_LoginBannerPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/config/login-banner.
func (n *System_LoginBannerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/config/login-banner in the given batch object.
func (n *System_LoginBannerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/config/login-banner.
func (n *System_LoginBannerPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/config/login-banner in the given batch object.
func (n *System_LoginBannerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/config/login-banner.
func (n *System_LoginBannerPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/config/login-banner in the given batch object.
func (n *System_LoginBannerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_LoginBannerPath extracts the value of the leaf LoginBanner from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_LoginBannerPath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LoginBanner
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/memory with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MemoryPath) Lookup(t testing.TB) *oc.QualifiedSystem_Memory {
	t.Helper()
	goStruct := &oc.System_Memory{}
	md, ok := oc.Lookup(t, n, "System_Memory", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Memory{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/memory with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MemoryPath) Get(t testing.TB) *oc.System_Memory {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/memory with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MemoryPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Memory {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Memory
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Memory{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Memory", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Memory{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/memory with a ONCE subscription.
func (n *System_MemoryPathAny) Get(t testing.TB) []*oc.System_Memory {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Memory
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/memory.
func (n *System_MemoryPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/memory in the given batch object.
func (n *System_MemoryPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/memory.
func (n *System_MemoryPath) Replace(t testing.TB, val *oc.System_Memory) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/memory in the given batch object.
func (n *System_MemoryPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Memory) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/memory.
func (n *System_MemoryPath) Update(t testing.TB, val *oc.System_Memory) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/memory in the given batch object.
func (n *System_MemoryPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Memory) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/messages with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MessagesPath) Lookup(t testing.TB) *oc.QualifiedSystem_Messages {
	t.Helper()
	goStruct := &oc.System_Messages{}
	md, ok := oc.Lookup(t, n, "System_Messages", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Messages{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MessagesPath) Get(t testing.TB) *oc.System_Messages {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MessagesPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Messages {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Messages
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Messages{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages with a ONCE subscription.
func (n *System_MessagesPathAny) Get(t testing.TB) []*oc.System_Messages {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Messages
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/messages.
func (n *System_MessagesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/messages in the given batch object.
func (n *System_MessagesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/messages.
func (n *System_MessagesPath) Replace(t testing.TB, val *oc.System_Messages) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/messages in the given batch object.
func (n *System_MessagesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Messages) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/messages.
func (n *System_MessagesPath) Update(t testing.TB, val *oc.System_Messages) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/messages in the given batch object.
func (n *System_MessagesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Messages) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_DebugServicePath) Lookup(t testing.TB) *oc.QualifiedSystem_Messages_DebugService {
	t.Helper()
	goStruct := &oc.System_Messages_DebugService{}
	md, ok := oc.Lookup(t, n, "System_Messages_DebugService", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Messages_DebugService{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_DebugServicePath) Get(t testing.TB) *oc.System_Messages_DebugService {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_DebugServicePathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Messages_DebugService {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Messages_DebugService
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_DebugService{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_DebugService", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Messages_DebugService{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/debug-entries/debug-service with a ONCE subscription.
func (n *System_Messages_DebugServicePathAny) Get(t testing.TB) []*oc.System_Messages_DebugService {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Messages_DebugService
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/messages/debug-entries/debug-service.
func (n *System_Messages_DebugServicePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/messages/debug-entries/debug-service in the given batch object.
func (n *System_Messages_DebugServicePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/messages/debug-entries/debug-service.
func (n *System_Messages_DebugServicePath) Replace(t testing.TB, val *oc.System_Messages_DebugService) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/messages/debug-entries/debug-service in the given batch object.
func (n *System_Messages_DebugServicePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Messages_DebugService) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/messages/debug-entries/debug-service.
func (n *System_Messages_DebugServicePath) Update(t testing.TB, val *oc.System_Messages_DebugService) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/messages/debug-entries/debug-service in the given batch object.
func (n *System_Messages_DebugServicePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Messages_DebugService) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_DebugService_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Messages_DebugService{}
	md, ok := oc.Lookup(t, n, "System_Messages_DebugService", goStruct, true, true)
	if ok {
		return convertSystem_Messages_DebugService_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_DebugService_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_DebugService_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_DebugService{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_DebugService", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_DebugService_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled with a ONCE subscription.
func (n *System_Messages_DebugService_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled.
func (n *System_Messages_DebugService_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled in the given batch object.
func (n *System_Messages_DebugService_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled.
func (n *System_Messages_DebugService_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled in the given batch object.
func (n *System_Messages_DebugService_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled.
func (n *System_Messages_DebugService_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/messages/debug-entries/debug-service/config/enabled in the given batch object.
func (n *System_Messages_DebugService_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Messages_DebugService_EnabledPath extracts the value of the leaf Enabled from its parent oc.System_Messages_DebugService
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Messages_DebugService_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_DebugService) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/config/service with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_DebugService_ServicePath) Lookup(t testing.TB) *oc.QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	goStruct := &oc.System_Messages_DebugService{}
	md, ok := oc.Lookup(t, n, "System_Messages_DebugService", goStruct, true, true)
	if ok {
		return convertSystem_Messages_DebugService_ServicePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/debug-entries/debug-service/config/service with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_DebugService_ServicePath) Get(t testing.TB) oc.E_Messages_DEBUG_SERVICE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/config/service with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_DebugService_ServicePathAny) Lookup(t testing.TB) []*oc.QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Messages_DEBUG_SERVICE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages_DebugService{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages_DebugService", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_DebugService_ServicePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/debug-entries/debug-service/config/service with a ONCE subscription.
func (n *System_Messages_DebugService_ServicePathAny) Get(t testing.TB) []oc.E_Messages_DEBUG_SERVICE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Messages_DEBUG_SERVICE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/messages/debug-entries/debug-service/config/service.
func (n *System_Messages_DebugService_ServicePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/messages/debug-entries/debug-service/config/service in the given batch object.
func (n *System_Messages_DebugService_ServicePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/messages/debug-entries/debug-service/config/service.
func (n *System_Messages_DebugService_ServicePath) Replace(t testing.TB, val oc.E_Messages_DEBUG_SERVICE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/messages/debug-entries/debug-service/config/service in the given batch object.
func (n *System_Messages_DebugService_ServicePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_Messages_DEBUG_SERVICE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/messages/debug-entries/debug-service/config/service.
func (n *System_Messages_DebugService_ServicePath) Update(t testing.TB, val oc.E_Messages_DEBUG_SERVICE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/messages/debug-entries/debug-service/config/service in the given batch object.
func (n *System_Messages_DebugService_ServicePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_Messages_DEBUG_SERVICE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Messages_DebugService_ServicePath extracts the value of the leaf Service from its parent oc.System_Messages_DebugService
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Messages_DEBUG_SERVICE.
func convertSystem_Messages_DebugService_ServicePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages_DebugService) *oc.QualifiedE_Messages_DEBUG_SERVICE {
	t.Helper()
	qv := &oc.QualifiedE_Messages_DEBUG_SERVICE{
		Metadata: md,
	}
	val := parent.Service
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/messages/config/severity with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Messages_SeverityPath) Lookup(t testing.TB) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	goStruct := &oc.System_Messages{}
	md, ok := oc.Lookup(t, n, "System_Messages", goStruct, true, true)
	if ok {
		return convertSystem_Messages_SeverityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/messages/config/severity with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Messages_SeverityPath) Get(t testing.TB) oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/messages/config/severity with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Messages_SeverityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_SystemLogging_SyslogSeverity
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Messages{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Messages", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Messages_SeverityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/messages/config/severity with a ONCE subscription.
func (n *System_Messages_SeverityPathAny) Get(t testing.TB) []oc.E_SystemLogging_SyslogSeverity {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_SystemLogging_SyslogSeverity
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/messages/config/severity.
func (n *System_Messages_SeverityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/messages/config/severity in the given batch object.
func (n *System_Messages_SeverityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/messages/config/severity.
func (n *System_Messages_SeverityPath) Replace(t testing.TB, val oc.E_SystemLogging_SyslogSeverity) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/messages/config/severity in the given batch object.
func (n *System_Messages_SeverityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SyslogSeverity) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/messages/config/severity.
func (n *System_Messages_SeverityPath) Update(t testing.TB, val oc.E_SystemLogging_SyslogSeverity) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/messages/config/severity in the given batch object.
func (n *System_Messages_SeverityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_SystemLogging_SyslogSeverity) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Messages_SeverityPath extracts the value of the leaf Severity from its parent oc.System_Messages
// and combines the update with an existing Metadata to return a *oc.QualifiedE_SystemLogging_SyslogSeverity.
func convertSystem_Messages_SeverityPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Messages) *oc.QualifiedE_SystemLogging_SyslogSeverity {
	t.Helper()
	qv := &oc.QualifiedE_SystemLogging_SyslogSeverity{
		Metadata: md,
	}
	val := parent.Severity
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/config/motd-banner with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_MotdBannerPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.System{}
	md, ok := oc.Lookup(t, n, "System", goStruct, true, true)
	if ok {
		return convertSystem_MotdBannerPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/config/motd-banner with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_MotdBannerPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/config/motd-banner with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_MotdBannerPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_MotdBannerPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/config/motd-banner with a ONCE subscription.
func (n *System_MotdBannerPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/config/motd-banner.
func (n *System_MotdBannerPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/config/motd-banner in the given batch object.
func (n *System_MotdBannerPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/config/motd-banner.
func (n *System_MotdBannerPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/config/motd-banner in the given batch object.
func (n *System_MotdBannerPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/config/motd-banner.
func (n *System_MotdBannerPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/config/motd-banner in the given batch object.
func (n *System_MotdBannerPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_MotdBannerPath extracts the value of the leaf MotdBanner from its parent oc.System
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertSystem_MotdBannerPath(t testing.TB, md *genutil.Metadata, parent *oc.System) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.MotdBanner
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_NtpPath) Lookup(t testing.TB) *oc.QualifiedSystem_Ntp {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Ntp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_NtpPath) Get(t testing.TB) *oc.System_Ntp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_NtpPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Ntp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Ntp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Ntp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp with a ONCE subscription.
func (n *System_NtpPathAny) Get(t testing.TB) []*oc.System_Ntp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Ntp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp.
func (n *System_NtpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp in the given batch object.
func (n *System_NtpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp.
func (n *System_NtpPath) Replace(t testing.TB, val *oc.System_Ntp) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp in the given batch object.
func (n *System_NtpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Ntp) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ntp.
func (n *System_NtpPath) Update(t testing.TB, val *oc.System_Ntp) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp in the given batch object.
func (n *System_NtpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Ntp) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/ntp/config/enable-ntp-auth with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_EnableNtpAuthPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_EnableNtpAuthPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnableNtpAuth())
}

// Get fetches the value at /openconfig-system/system/ntp/config/enable-ntp-auth with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_EnableNtpAuthPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/config/enable-ntp-auth with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_EnableNtpAuthPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_EnableNtpAuthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/config/enable-ntp-auth with a ONCE subscription.
func (n *System_Ntp_EnableNtpAuthPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/config/enable-ntp-auth.
func (n *System_Ntp_EnableNtpAuthPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/config/enable-ntp-auth in the given batch object.
func (n *System_Ntp_EnableNtpAuthPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/config/enable-ntp-auth.
func (n *System_Ntp_EnableNtpAuthPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/config/enable-ntp-auth in the given batch object.
func (n *System_Ntp_EnableNtpAuthPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/config/enable-ntp-auth.
func (n *System_Ntp_EnableNtpAuthPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/config/enable-ntp-auth in the given batch object.
func (n *System_Ntp_EnableNtpAuthPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_EnableNtpAuthPath extracts the value of the leaf EnableNtpAuth from its parent oc.System_Ntp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Ntp_EnableNtpAuthPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.EnableNtpAuth
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.System_Ntp{}
	md, ok := oc.Lookup(t, n, "System_Ntp", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-system/system/ntp/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/config/enabled with a ONCE subscription.
func (n *System_Ntp_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/config/enabled.
func (n *System_Ntp_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/config/enabled in the given batch object.
func (n *System_Ntp_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/config/enabled.
func (n *System_Ntp_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/config/enabled in the given batch object.
func (n *System_Ntp_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/config/enabled.
func (n *System_Ntp_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/config/enabled in the given batch object.
func (n *System_Ntp_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_EnabledPath extracts the value of the leaf Enabled from its parent oc.System_Ntp
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertSystem_Ntp_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp) *oc.QualifiedBool {
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

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKeyPath) Lookup(t testing.TB) *oc.QualifiedSystem_Ntp_NtpKey {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, false, true)
	if ok {
		return (&oc.QualifiedSystem_Ntp_NtpKey{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKeyPath) Get(t testing.TB) *oc.System_Ntp_NtpKey {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKeyPathAny) Lookup(t testing.TB) []*oc.QualifiedSystem_Ntp_NtpKey {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSystem_Ntp_NtpKey
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSystem_Ntp_NtpKey{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key with a ONCE subscription.
func (n *System_Ntp_NtpKeyPathAny) Get(t testing.TB) []*oc.System_Ntp_NtpKey {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.System_Ntp_NtpKey
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key.
func (n *System_Ntp_NtpKeyPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/ntp-keys/ntp-key in the given batch object.
func (n *System_Ntp_NtpKeyPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key.
func (n *System_Ntp_NtpKeyPath) Replace(t testing.TB, val *oc.System_Ntp_NtpKey) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/ntp-keys/ntp-key in the given batch object.
func (n *System_Ntp_NtpKeyPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.System_Ntp_NtpKey) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key.
func (n *System_Ntp_NtpKeyPath) Update(t testing.TB, val *oc.System_Ntp_NtpKey) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/ntp-keys/ntp-key in the given batch object.
func (n *System_Ntp_NtpKeyPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.System_Ntp_NtpKey) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id.
func (n *System_Ntp_NtpKey_KeyIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id in the given batch object.
func (n *System_Ntp_NtpKey_KeyIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id.
func (n *System_Ntp_NtpKey_KeyIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id in the given batch object.
func (n *System_Ntp_NtpKey_KeyIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id.
func (n *System_Ntp_NtpKey_KeyIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-id in the given batch object.
func (n *System_Ntp_NtpKey_KeyIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertSystem_Ntp_NtpKey_KeyIdPath extracts the value of the leaf KeyId from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertSystem_Ntp_NtpKey_KeyIdPath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.KeyId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *System_Ntp_NtpKey_KeyTypePath) Lookup(t testing.TB) *oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	goStruct := &oc.System_Ntp_NtpKey{}
	md, ok := oc.Lookup(t, n, "System_Ntp_NtpKey", goStruct, true, true)
	if ok {
		return convertSystem_Ntp_NtpKey_KeyTypePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *System_Ntp_NtpKey_KeyTypePath) Get(t testing.TB) oc.E_System_NTP_AUTH_TYPE {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Lookup(t testing.TB) []*oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_System_NTP_AUTH_TYPE
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.System_Ntp_NtpKey{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "System_Ntp_NtpKey", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertSystem_Ntp_NtpKey_KeyTypePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type with a ONCE subscription.
func (n *System_Ntp_NtpKey_KeyTypePathAny) Get(t testing.TB) []oc.E_System_NTP_AUTH_TYPE {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_System_NTP_AUTH_TYPE
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type.
func (n *System_Ntp_NtpKey_KeyTypePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type in the given batch object.
func (n *System_Ntp_NtpKey_KeyTypePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type.
func (n *System_Ntp_NtpKey_KeyTypePath) Replace(t testing.TB, val oc.E_System_NTP_AUTH_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type in the given batch object.
func (n *System_Ntp_NtpKey_KeyTypePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_System_NTP_AUTH_TYPE) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type.
func (n *System_Ntp_NtpKey_KeyTypePath) Update(t testing.TB, val oc.E_System_NTP_AUTH_TYPE) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-system/system/ntp/ntp-keys/ntp-key/config/key-type in the given batch object.
func (n *System_Ntp_NtpKey_KeyTypePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_System_NTP_AUTH_TYPE) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertSystem_Ntp_NtpKey_KeyTypePath extracts the value of the leaf KeyType from its parent oc.System_Ntp_NtpKey
// and combines the update with an existing Metadata to return a *oc.QualifiedE_System_NTP_AUTH_TYPE.
func convertSystem_Ntp_NtpKey_KeyTypePath(t testing.TB, md *genutil.Metadata, parent *oc.System_Ntp_NtpKey) *oc.QualifiedE_System_NTP_AUTH_TYPE {
	t.Helper()
	qv := &oc.QualifiedE_System_NTP_AUTH_TYPE{
		Metadata: md,
	}
	val := parent.KeyType
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}
