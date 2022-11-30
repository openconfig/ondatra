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
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ygot/ygot"
)

var (
	translateFunctions = map[string]func(ixweb.StatTable, []string, []string) (*oc.Root, error){
		portStatsCaption:              translatePortStats,
		portCPUStatsCaption:           translatePortCPUStats,
		ixweb.TrafficItemStatsCaption: translateTrafficItemStats,
		flowStatsCaption:              translateFlowStats,
		ixweb.EgressStatsCaption:      translateEgressStats,
	}
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
func translate(st *Stats) (ygot.GoStruct, error) {
	root := &oc.Root{}
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
			root = n.(*oc.Root)
		}
	}

	if log.V(3) {
		log.Infof("Merged translated device: %s", jsonDebug(root))
	}
	return root, nil
}

func translatePortCPUStats(in ixweb.StatTable, _, _ []string) (*oc.Root, error) {
	portCPURows, err := parsePortCPUStats(in)
	if err != nil {
		return nil, err
	}
	log.V(3).Infof("Parsed Port CPU Stats: %v", portCPURows)

	d := &oc.Root{}
	for _, row := range portCPURows {
		port := d.GetOrCreateComponent(row.PortName)
		cpu := d.GetOrCreateComponent(fmt.Sprintf("%s_CPU", row.PortName))

		port.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
		port.NewSubcomponent(*cpu.Name)

		cpu.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU
		cpu.Parent = port.Name

		if row.TotalMemory != nil && row.FreeMemory != nil {
			port.Memory = &oc.Component_Memory{
				Available: ygot.Uint64(*row.FreeMemory),
				Utilized:  ygot.Uint64(*row.TotalMemory - *row.FreeMemory),
			}
		}

		if row.CPULoad != nil {
			cpu.GetOrCreateCpu().GetOrCreateUtilization().Instant = ygot.Uint8(uint8(*row.CPULoad))
		}
	}

	return d, nil
}

func translatePortStats(in ixweb.StatTable, _, _ []string) (*oc.Root, error) {
	portRows, err := parsePortStats(in)
	if err != nil {
		return nil, err
	}
	log.V(3).Infof("Parsed Port Stats: %v", portRows)

	d := &oc.Root{}
	for _, row := range portRows {
		i := d.GetOrCreateInterface(row.PortName)
		i.Counters = &oc.Interface_Counters{
			InOctets:  row.BytesRx,
			InPkts:    row.FramesRx,
			OutOctets: row.BytesTx,
			OutPkts:   row.FramesTx,
		}
		i.GetOrCreateEthernet().GetOrCreateCounters().InCrcErrors = row.CRCErrs
		i.InRate = pfloat32Bytes(row.RxRate)
		i.OutRate = pfloat32Bytes(row.TxRate)

		i.Type = oc.IETFInterfaces_InterfaceType_ethernetCsmacd
		switch row.LinkState {
		case "":
			// No error when empty.
		case "Link Up":
			i.OperStatus = oc.Interface_OperStatus_UP
		case "Link Down", "No PCS Lock":
			i.OperStatus = oc.Interface_OperStatus_DOWN
		default:
			return nil, fmt.Errorf("statistics row %q has an unmappable port link state %q", row.PortName, row.LinkState)
		}

		// TODO(team): Map different speed interfaces - need to determine what the possible Ixia values are.
		switch row.LineSpeed {
		case "":
			// No error when empty.
		case "10GE LAN":
			i.GetOrCreateEthernet().PortSpeed = oc.IfEthernet_ETHERNET_SPEED_SPEED_10GB
		case "100GE":
			i.GetOrCreateEthernet().PortSpeed = oc.IfEthernet_ETHERNET_SPEED_SPEED_100GB
		case "400GE":
			i.GetOrCreateEthernet().PortSpeed = oc.IfEthernet_ETHERNET_SPEED_SPEED_400GB
		default:
			return nil, fmt.Errorf("statistics row %q has an unmappable port link speed %q", row.PortName, row.LineSpeed)
		}
	}

	return d, nil
}

func translateTrafficItemStats(in ixweb.StatTable, _, _ []string) (*oc.Root, error) {
	itemRows, err := parseFlowStats(in)
	if err != nil {
		return nil, err
	}
	log.V(3).Infof("Parsed Item Stats: %v", itemRows)

	d := &oc.Root{}
	for _, row := range itemRows {
		f := d.GetOrCreateFlow(row.TrafficItem)
		f.Counters = &oc.Flow_Counters{
			InOctets: row.RxBytes,
			InPkts:   row.RxFrames,
			OutPkts:  row.TxFrames,
		}
		f.LossPct = pfloat32Bytes(row.LossPct)
		f.InRate = pfloat32Bytes(row.RxRate)
		f.InFrameRate = pfloat32Bytes(row.RxFrameRate)
		f.OutRate = pfloat32Bytes(row.TxRate)
		f.OutFrameRate = pfloat32Bytes(row.TxFrameRate)
		f.ConvergenceTime = row.ConvergenceTime
		f.FirstPacketLatency = row.FirstPacketTime
	}
	return d, nil
}

func translateFlowStats(in ixweb.StatTable, _, _ []string) (*oc.Root, error) {
	flowRows, err := parseFlowStats(in)
	if err != nil {
		return nil, err
	}
	log.V(3).Infof("Parsed Flow Stats: %v", flowRows)

	d := &oc.Root{}
	for _, row := range flowRows {
		f := d.GetOrCreateFlow(row.TrafficItem)
		it := ingressTrackingFromFlowStats(f, row)
		it.Counters = &oc.Flow_IngressTracking_Counters{
			InOctets: row.RxBytes,
			InPkts:   row.RxFrames,
			OutPkts:  row.TxFrames,
		}
		it.LossPct = pfloat32Bytes(row.LossPct)
		it.InRate = pfloat32Bytes(row.RxRate)
		it.InFrameRate = pfloat32Bytes(row.RxFrameRate)
		it.OutRate = pfloat32Bytes(row.TxRate)
		it.OutFrameRate = pfloat32Bytes(row.TxFrameRate)
	}
	return d, nil
}

func translateEgressStats(in ixweb.StatTable, itFlows, etFlows []string) (*oc.Root, error) {
	egressRows, err := parseEgressStats(in)
	if err != nil {
		return nil, err
	}
	log.V(3).Infof("Parsed Egress Stats: %v", egressRows)

	d := &oc.Root{}
	var f *oc.Flow
	var it *oc.Flow_IngressTracking
	for i, row := range egressRows {
		// When there is only a single egress tracking flow, Ixia will omit the
		// "Traffic Item" column entirely, so populate it with that flow name.
		if i == 0 && len(etFlows) == 1 {
			row.TrafficItem = etFlows[0]
		}

		// If the traffic item value is set, we are starting a new flow row.
		if row.TrafficItem != "" {
			f = d.GetOrCreateFlow(row.TrafficItem)
			if isIngressTracked(row.TrafficItem, itFlows) {
				it = ingressTrackingFromFlowStats(f, row.flowRow)
				it.Filter = &row.filter
			} else {
				f.Filter = &row.filter
			}
			continue
		}

		// If we aren't starting a new flow row but the filter starts with "Custom:",
		// then we are starting a new ingress-tracked value within the flow.
		// NOTE: This only works because we only support "custom" egress tracking.
		// If that changes, we will need to do something more sophisticated here.
		if it != nil && strings.HasPrefix(row.filter, "Custom:") {
			it = ingressTrackingFromFlowStats(f, row.flowRow)
			it.Filter = &row.filter
			continue
		}

		if it == nil {
			ef := f.GetOrCreateEgressTracking(row.filter)
			ef.Counters = &oc.Flow_EgressTracking_Counters{
				InOctets: row.RxBytes,
				InPkts:   row.RxFrames,
				OutPkts:  row.TxFrames,
			}
			ef.LossPct = pfloat32Bytes(row.LossPct)
			ef.InRate = pfloat32Bytes(row.RxRate)
			ef.InFrameRate = pfloat32Bytes(row.RxFrameRate)
			ef.OutRate = pfloat32Bytes(row.TxRate)
			ef.OutFrameRate = pfloat32Bytes(row.TxFrameRate)
		} else {
			ef := it.GetOrCreateEgressTracking(row.filter)
			ef.Counters = &oc.Flow_IngressTracking_EgressTracking_Counters{
				InOctets: row.RxBytes,
				InPkts:   row.RxFrames,
				OutPkts:  row.TxFrames,
			}
			ef.LossPct = pfloat32Bytes(row.LossPct)
			ef.InRate = pfloat32Bytes(row.RxRate)
			ef.InFrameRate = pfloat32Bytes(row.RxFrameRate)
			ef.OutRate = pfloat32Bytes(row.TxRate)
			ef.OutFrameRate = pfloat32Bytes(row.TxFrameRate)
		}
	}

	return d, nil
}

func ingressTrackingFromFlowStats(flow *oc.Flow, row *flowRow) *oc.Flow_IngressTracking {
	return flow.GetOrCreateIngressTracking(
		row.TxPort,
		row.RxPort,
		mplsLabelFromUint(row.MPLSLabel),
		row.SrcIPv4,
		row.DstIPv4,
		row.SrcIPv6,
		row.DstIPv6,
		vlanIDFromUint(row.VLANID),
	)
}

func mplsLabelFromUint(label *uint64) oc.Flow_IngressTracking_MplsLabel_Union {
	if label == nil {
		return oc.IngressTracking_MplsLabel_NO_LABEL
	}
	labelNum := *label
	switch labelNum {
	case 0:
		return oc.IngressTracking_MplsLabel_IPV4_EXPLICIT_NULL
	case 1:
		return oc.IngressTracking_MplsLabel_ROUTER_ALERT
	case 2:
		return oc.IngressTracking_MplsLabel_IPV6_EXPLICIT_NULL
	case 3:
		return oc.IngressTracking_MplsLabel_IMPLICIT_NULL
	case 7:
		return oc.IngressTracking_MplsLabel_ENTROPY_LABEL_INDICATOR
	}
	const maxUint32 = uint64(^uint32(0))
	if labelNum > maxUint32 {
		return oc.IngressTracking_MplsLabel_UNSET
	}
	return oc.UnionUint32(labelNum)
}

func vlanIDFromUint(id *uint64) uint16 {
	if id == nil {
		return 0
	}
	return uint16(*id)
}

func pfloat32Bytes(f *float32) oc.Binary {
	if f == nil {
		return nil
	}
	return float32Bytes(*f)
}

func float32Bytes(f float32) oc.Binary {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(f))
	return oc.Binary(b)
}

// jsonDebug renders a device struct to JSON such that it can be logged.
// Since it is used only in debugging, it logs an error if the device struct
// cannot be rendered to JSON, and returns an error string to the user.
func jsonDebug(d *oc.Root) string {
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
