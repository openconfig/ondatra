// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixgnmi

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strings"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ondatra/internal/ixweb"
	"github.com/openconfig/ondatra/telemetry"
)

var (
	translateFunctions = map[string]func(ixweb.StatTable, []string, []string) (*telemetry.Device, error){
		portStatsCaption:         translatePortStats,
		portCPUStatsCaption:      translatePortCPUStats,
		flowStatsCaption:         translateFlowStats,
		ixweb.EgressStatsCaption: translateEgressStats,
	}
	// debugLog enables logging of debug information in the library. It is expensive
	// so should only be enabled for development purposes.
	debugLog = false
)

// translate converts Ixia statistics tables to an OpenConfig schema.
// The form of the input table is assumed to be:
//
//	{
//		StatisticsViewName1: [
//			{
//				Caption1: Value1,
//				Caption2: Value2
//			},
//			{
//				Caption1: Value1,
//				Caption2: Value2,
//			},
//		],
//		StatisticsViewName2: [
//			{
//				...
//			}
//		],
//	}
//
// where each item in the slice under a statistics view represents a row of data in that view.
//
// The itFlows and etFlows parameters indicates which flows have ingress
// tracking and egress tracking enabled, respectively.
func translate(st *Stats) (ygot.ValidatedGoStruct, error) {
	root := &telemetry.Device{}
	for caption, table := range st.Tables {
		if fn, ok := translateFunctions[caption]; ok {
			d, err := fn(table, st.IngressTrackFlows, st.EgressTrackFlows)
			if err != nil {
				log.Errorf("Got error: %v while translating %s", err, caption)
				continue
			}
			n, err := ygot.MergeStructs(root, d)
			if err != nil {
				log.Errorf("Got error: %v while merging %s", err, caption)
				continue
			}
			root = n.(*telemetry.Device)
		}
	}

	if debugLog {
		log.Infof("merged processed device, %s", jsonDebug(root))
	}

	return root, nil
}

// translatePortCPUStats maps the Ixia Port CPU Statistics to an OpenConfig
// device structure.
func translatePortCPUStats(in ixweb.StatTable, _, _ []string) (*telemetry.Device, error) {
	portCPURows, err := parsePortCPUStats(in)
	if err != nil {
		return nil, err
	}

	d := &telemetry.Device{}
	for _, row := range portCPURows {
		portName, err := shortPortName(row.portName)
		if err != nil {
			return nil, err
		}
		port := d.GetOrCreateComponent(portName)
		cpu := d.GetOrCreateComponent(fmt.Sprintf("%s_CPU", portName))

		port.Type = telemetry.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
		port.NewSubcomponent(*cpu.Name)

		cpu.Type = telemetry.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU
		cpu.Parent = port.Name

		if row.totalMemory != nil && row.freeMemory != nil {
			port.Memory = &telemetry.Component_Memory{
				Available: ygot.Uint64(*row.freeMemory),
				Utilized:  ygot.Uint64(*row.totalMemory - *row.freeMemory),
			}
		}

		if row.cpuLoad != nil {
			cpu.GetOrCreateCpu().GetOrCreateUtilization().Instant = ygot.Uint8(uint8(*row.cpuLoad))
		}
	}

	if debugLog {
		log.Infof("port CPU device, %s", jsonDebug(d))
	}

	return d, nil
}

// translatePortStats translates the Ixia "Port Stats" view to OpenConfig, returning
// the translated statistics as a ygot-generated struct.
func translatePortStats(in ixweb.StatTable, _, _ []string) (*telemetry.Device, error) {
	portRows, err := parsePortStats(in)
	if err != nil {
		return nil, err
	}

	d := &telemetry.Device{}
	for _, row := range portRows {
		portName, err := shortPortName(row.portName)
		if err != nil {
			return nil, err
		}
		i := d.GetOrCreateInterface(portName)
		i.Counters = &telemetry.Interface_Counters{
			InOctets:  row.bytesRx,
			InPkts:    row.framesRx,
			OutOctets: row.bytesTx,
			OutPkts:   row.framesTx,
		}
		i.GetOrCreateEthernet().GetOrCreateCounters().InCrcErrors = row.crcErrs
		i.InRate = pfloat32Bytes(row.rxRate)
		i.OutRate = pfloat32Bytes(row.txRate)

		i.Type = telemetry.IETFInterfaces_InterfaceType_ethernetCsmacd
		switch row.linkState {
		case "":
			// No error when empty.
		case "Link Up":
			i.OperStatus = telemetry.Interface_OperStatus_UP
		case "Link Down", "No PCS Lock":
			i.OperStatus = telemetry.Interface_OperStatus_DOWN
		default:
			return nil, errors.Errorf("statistics row %q has an unmappable port link state %q", row.portName, row.linkState)
		}

		// TODO: Map different speed interfaces - need to determine what the possible Ixia values are.
		switch row.lineSpeed {
		case "":
			// No error when empty.
		case "10GE LAN":
			i.GetOrCreateEthernet().PortSpeed = telemetry.IfEthernet_ETHERNET_SPEED_SPEED_10GB
		case "100GE":
			i.GetOrCreateEthernet().PortSpeed = telemetry.IfEthernet_ETHERNET_SPEED_SPEED_100GB
		case "400GE":
			i.GetOrCreateEthernet().PortSpeed = telemetry.IfEthernet_ETHERNET_SPEED_SPEED_400GB
		default:
			return nil, errors.Errorf("statistics row %q has an unmappable port link speed %q", row.portName, row.lineSpeed)
		}
	}

	return d, nil
}

// translateFlowStats translates the Flow Statistics view in the
// supplied ixweb.StatTable to a telemetry.Device ygot-generated object.
func translateFlowStats(in ixweb.StatTable, itFlows, _ []string) (*telemetry.Device, error) {
	flowRows, err := parseFlowStats(in)
	if err != nil {
		return nil, err
	}

	d := &telemetry.Device{}
	for _, row := range flowRows {
		f := d.GetOrCreateFlow(row.trafficItem)
		if isIngressTracked(row.trafficItem, itFlows) {
			it := ingressTrackingFromFlowStats(f, row)
			it.Counters = &telemetry.Flow_IngressTracking_Counters{
				InOctets: row.rxBytes,
				InPkts:   row.rxFrames,
				OutPkts:  row.txFrames,
			}
			it.LossPct = pfloat32Bytes(row.lossPct)
			it.InRate = pfloat32Bytes(row.rxRate)
			it.InFrameRate = pfloat32Bytes(row.rxFrameRate)
			it.OutRate = pfloat32Bytes(row.txRate)
			it.OutFrameRate = pfloat32Bytes(row.txFrameRate)
		} else {
			f.Counters = &telemetry.Flow_Counters{
				InOctets: row.rxBytes,
				InPkts:   row.rxFrames,
				OutPkts:  row.txFrames,
			}
			f.LossPct = pfloat32Bytes(row.lossPct)
			f.InRate = pfloat32Bytes(row.rxRate)
			f.InFrameRate = pfloat32Bytes(row.rxFrameRate)
			f.OutRate = pfloat32Bytes(row.txRate)
			f.OutFrameRate = pfloat32Bytes(row.txFrameRate)
		}
	}
	return d, nil
}

// translateEgressStats translates the Custom Egress Stats view in the
// supplied ixweb.StatTable to a telemetry.Device ygot-generated object.
func translateEgressStats(in ixweb.StatTable, itFlows, etFlows []string) (*telemetry.Device, error) {
	egressRows, err := parseEgressStats(in)
	if err != nil {
		return nil, err
	}

	d := &telemetry.Device{}
	var f *telemetry.Flow
	var it *telemetry.Flow_IngressTracking
	for i, row := range egressRows {
		// When there is only a single egress tracking flow, Ixia will omit the
		// "Traffic Item" column entirely, so populate it with that flow name.
		if i == 0 && len(etFlows) == 1 {
			row.trafficItem = etFlows[0]
		}

		// Setting the traffic item key, if present.
		if row.trafficItem != "" {
			f = d.GetOrCreateFlow(row.trafficItem)
			if isIngressTracked(row.trafficItem, itFlows) {
				it = ingressTrackingFromFlowStats(f, row.flowRow)
				it.Filter = &row.filter
			} else {
				f.Filter = &row.filter
			}
			continue
		}

		if it == nil {
			ef := f.GetOrCreateEgressTracking(row.filter)
			ef.Counters = &telemetry.Flow_EgressTracking_Counters{
				InOctets: row.rxBytes,
				InPkts:   row.rxFrames,
				OutPkts:  row.txFrames,
			}
			ef.LossPct = pfloat32Bytes(row.lossPct)
			ef.InRate = pfloat32Bytes(row.rxRate)
			ef.InFrameRate = pfloat32Bytes(row.rxFrameRate)
			ef.OutRate = pfloat32Bytes(row.txRate)
			ef.OutFrameRate = pfloat32Bytes(row.txFrameRate)
		} else {
			ef := it.GetOrCreateEgressTracking(row.filter)
			ef.Counters = &telemetry.Flow_IngressTracking_EgressTracking_Counters{
				InOctets: row.rxBytes,
				InPkts:   row.rxFrames,
				OutPkts:  row.txFrames,
			}
			ef.LossPct = pfloat32Bytes(row.lossPct)
			ef.InRate = pfloat32Bytes(row.rxRate)
			ef.InFrameRate = pfloat32Bytes(row.rxFrameRate)
			ef.OutRate = pfloat32Bytes(row.txRate)
			ef.OutFrameRate = pfloat32Bytes(row.txFrameRate)
		}
	}

	return d, nil
}

func ingressTrackingFromFlowStats(flow *telemetry.Flow, row *flowRow) *telemetry.Flow_IngressTracking {
	return flow.GetOrCreateIngressTracking(
		row.rxPort,
		row.txPort,
		mplsLabelFromUint(row.mplsLabel),
		row.srcIPv4,
		row.dstIPv4,
		row.srcIPv6,
		row.dstIPv6,
	)
}

func mplsLabelFromUint(label *uint64) telemetry.Flow_IngressTracking_MplsLabel_Union {
	if label == nil {
		return telemetry.MplsTypes_MplsLabel_Enum_NO_LABEL
	}
	labelNum := *label
	switch labelNum {
	case 0:
		return telemetry.MplsTypes_MplsLabel_Enum_IPV4_EXPLICIT_NULL
	case 1:
		return telemetry.MplsTypes_MplsLabel_Enum_ROUTER_ALERT
	case 2:
		return telemetry.MplsTypes_MplsLabel_Enum_IPV6_EXPLICIT_NULL
	case 3:
		return telemetry.MplsTypes_MplsLabel_Enum_IMPLICIT_NULL
	case 7:
		return telemetry.MplsTypes_MplsLabel_Enum_ENTROPY_LABEL_INDICATOR
	}
	const maxUint32 = uint64(^uint32(0))
	if labelNum > maxUint32 {
		return telemetry.MplsTypes_MplsLabel_Enum_UNSET
	}
	return telemetry.UnionUint32(labelNum)
}

// Strips off the Ixia name from the port name.
func shortPortName(fullPortName string) (string, error) {
	parts := strings.SplitAfterN(fullPortName, "/", 2)
	if len(parts) < 2 || len(parts[1]) == 0 {
		return "", errors.Errorf("invalid port name: got %q, want [ixia_name]/[port_name]", fullPortName)
	}
	return parts[1], nil
}

func pfloat32Bytes(f *float32) telemetry.Binary {
	if f == nil {
		return nil
	}
	return float32Bytes(*f)
}

func float32Bytes(f float32) telemetry.Binary {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(f))
	return telemetry.Binary(b)
}

// jsonDebug renders a device struct to JSON such that it can be logged.
// Since it is used only in debugging, it logs an error if the device struct
// cannot be rendered to JSON, and returns an error string to the user.
func jsonDebug(d *telemetry.Device) string {
	js, err := ygot.ConstructIETFJSON(d, nil)
	if err != nil {
		log.Errorf("cannot render device to JSON during debugging, %v", err)
		return "(unrenderable)"
	}

	j, err := json.MarshalIndent(js, "", "  ")
	if err != nil {
		log.Errorf("cannot marshal JSON for device, %v", err)
		return "(unmarshallable)"
	}
	return string(j)
}

func isIngressTracked(flow string, itFlows []string) bool {
	for _, f := range itFlows {
		if f == flow {
			return true
		}
	}
	return false
}
