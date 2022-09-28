// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ate

import (
	"fmt"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

type trafficType string

const (
	ethTraffic  trafficType = "ethernetVlan"
	ipv4Traffic trafficType = "ipv4"
	ipv6Traffic trafficType = "ipv6"
	rawTraffic  trafficType = "raw"
)

var imixPreset = map[opb.FrameSize_ImixPreset]string{
	opb.FrameSize_IMIX_CISCO:         `"cisco"`,
	opb.FrameSize_IMIX_DEFAULT:       `"imix"`,
	opb.FrameSize_IMIX_IPSEC:         `"ipSecImix"`,
	opb.FrameSize_IMIX_IPV6:          `"ipV6Imix"`,
	opb.FrameSize_IMIX_RPR_QUADMODAL: `"rprQuar"`,
	opb.FrameSize_IMIX_RPR_TRIMODAL:  `"rprTri"`,
	opb.FrameSize_IMIX_STANDARD:      `"standardImix"`,
	opb.FrameSize_IMIX_TCP:           `"tcpImix"`,
	opb.FrameSize_IMIX_TOLLY:         `"tolly"`,
}

type headers struct {
	eth  *opb.EthernetHeader
	ipv4 *opb.Ipv4Header
	ipv6 *opb.Ipv6Header
}

func (h *headers) String() string {
	return fmt.Sprintf("headers%+v", *h)
}

func (ix *ixATE) addTraffic(flows []*opb.Flow) error {
	ix.cfg.Traffic = &ixconfig.Traffic{UseRfc5952: ixconfig.Bool(true)}
	for _, f := range flows {
		if err := ix.addTrafficItem(f); err != nil {
			return err
		}
	}
	return nil
}

func (ix *ixATE) addTrafficItem(f *opb.Flow) error {
	hdrs, err := resolveHeaders(f.GetHeaders())
	if err != nil {
		return fmt.Errorf("bad header spec for flow %q: %w", f.GetName(), err)
	}
	srcEPs := f.GetSrcEndpoints()
	dstEPs := f.GetDstEndpoints()
	if len(srcEPs) == 0 || len(dstEPs) == 0 {
		return fmt.Errorf("flow %q needs both source and destination endpoints defined", f.GetName())
	}
	// Check that reference interfaces exist. All subsequent functions assume that they do.
	for _, ep := range srcEPs {
		if _, ok := ix.intfs[ep.GetInterfaceName()]; !ok {
			return fmt.Errorf("no interface configured for source endpoint %v", ep)
		}
	}
	for _, ep := range dstEPs {
		if _, ok := ix.intfs[ep.GetInterfaceName()]; !ok {
			return fmt.Errorf("no interface configured for dest endpoint %v", ep)
		}
	}

	trafType, err := resolveTrafficType(hdrs, srcEPs, dstEPs)
	if err != nil {
		return fmt.Errorf("could not determine traffic type for flow %q: %w", f.GetName(), err)
	}

	epFn := devOrGeneratedEPs
	if trafType == rawTraffic {
		inferAddresses(hdrs, srcEPs, dstEPs, ix.intfs)
		epFn = portOrLagEPs
	}
	srcs, err := epFn(srcEPs, ix.intfs, true)
	if err != nil {
		return fmt.Errorf("could not find source endpoint paths for flow %q: %w", f.GetName(), err)
	}
	dsts, err := epFn(dstEPs, ix.intfs, false)
	if err != nil {
		return fmt.Errorf("could not find dest endpoint paths for flow %q: %w", f.GetName(), err)
	}

	fr, _, err := frameRate(f.GetFrameRate())
	if err != nil {
		return fmt.Errorf("could not compute frame rate for flow %q: %w", f.GetName(), err)
	}
	fs, _, err := frameSize(f.GetFrameSize())
	if err != nil {
		return fmt.Errorf("could not compute frame size for flow %q: %w", f.GetName(), err)
	}

	tc, err := transmissionControl(f.GetTransmission())
	if err != nil {
		return fmt.Errorf("could not compute transmission control for flow %q: %w", f.GetName(), err)
	}
	var crc *string
	if hdrs.eth.GetBadCrc() {
		crc = ixconfig.String("badCrc")
	}
	epSet := &ixconfig.TrafficTrafficItemEndpointSet{
		Name: ixconfig.String(f.GetName()),
		// Following two values are explicitly set to reduce warnings from config push.
		FullyMeshedEndpoints: []string{},
		TrafficGroups:        []string{},
	}
	epSet.SetSourcesRefs(srcs)
	epSet.SetDestinationsRefs(dsts)
	ti := &ixconfig.TrafficTrafficItem{
		Name:        ixconfig.String(f.GetName()),
		TrafficType: ixconfig.String(string(trafType)),
		EndpointSet: []*ixconfig.TrafficTrafficItemEndpointSet{epSet},
		RouteMesh:   ixconfig.String("oneToOne"),
		ConfigElement: []*ixconfig.TrafficTrafficItemConfigElement{{
			FrameRate:           fr,
			FrameSize:           fs,
			TransmissionControl: tc,
			Crc:                 crc,
		}},
	}
	if trafType == rawTraffic {
		ti.RawTrafficRxPortsBehavior = ixconfig.String("loadBalanced")
	}

	ingressTracking, trackFlow, err := ingressTrackingCfg(f, trafType)
	if err != nil {
		return fmt.Errorf("could not configure ingress tracking for flow %q: %w", f.GetName(), err)
	}
	if trackFlow {
		ix.ingressTrackFlows = append(ix.ingressTrackFlows, f.GetName())
	}
	ti.Tracking = []*ixconfig.TrafficTrafficItemTracking{ingressTracking}

	if et := f.GetEgressTracking(); et != nil && et.GetEnabled() {
		const minOffset, maxOffset = 32, 262128
		if offset := et.GetOffset(); offset < minOffset || offset > maxOffset {
			return fmt.Errorf("egress tracking offset %v not in allowed range [%v, %v]", offset, minOffset, maxOffset)
		}
		const minWidth, maxWidth = 1, 22
		if width := et.GetWidth(); width < minWidth || width > maxWidth {
			return fmt.Errorf("egress tracking width %v not in allowed range [%v, %v]", width, minWidth, maxWidth)
		}
		ix.egressTrackFlowCounts[f.GetName()] = int(et.GetCount())
		ti.EgressEnabled = ixconfig.Bool(true)
		ti.EgressTracking = []*ixconfig.TrafficTrafficItemEgressTracking{{
			Encapsulation:    ixconfig.String("Any: Use Custom Settings"),
			CustomOffsetBits: ixconfig.NumberUint32(et.GetOffset()),
			CustomWidthBits:  ixconfig.NumberUint32(et.GetWidth()),
		}}
	}

	var hasSrcVLAN bool
	for _, ep := range srcEPs {
		hasSrcVLAN = hasSrcVLAN || ix.intfs[ep.GetInterfaceName()].hasVLAN
	}
	stacks := make([]*ixconfig.TrafficTrafficItemConfigElementStack, 0)
	for _, hdr := range f.GetHeaders() {
		s, err := headerStacks(hdr, len(stacks), hasSrcVLAN)
		if err != nil {
			return fmt.Errorf("could not add header for flow %q: %w", f.GetName(), err)
		}
		stacks = append(stacks, s...)
	}
	ti.ConfigElement[0].Stack = stacks

	ix.flowToTrafficItem[f.GetName()] = ti
	ix.cfg.Traffic.TrafficItem = append(ix.cfg.Traffic.TrafficItem, ti)
	return nil
}

func resolveHeaders(flowHdrs []*opb.Header) (*headers, error) {
	if len(flowHdrs) == 0 {
		return nil, fmt.Errorf("must specify at least one header")
	}
	hdrs := &headers{eth: flowHdrs[0].GetEth()}
	if hdrs.eth == nil {
		return nil, fmt.Errorf("first header must be an ethernet header, got: %v", flowHdrs)
	}
	if len(flowHdrs) > 1 {
		h := flowHdrs[1]
		if (h.GetMpls() != nil || h.GetMacsec() != nil) && len(flowHdrs) > 2 {
			h = flowHdrs[2]
		}
		hdrs.ipv4 = h.GetIpv4()
		hdrs.ipv6 = h.GetIpv6()
	}
	return hdrs, nil
}

func resolveTrafficType(hdrs *headers, srcEPs, dstEPs []*opb.Flow_Endpoint) (trafficType, error) {
	ethAddrsSet := hdrs.eth.GetVlanId() != 0 || hdrs.eth.GetSrcAddr() != nil || hdrs.eth.GetDstAddr() != nil

	// For IP traffic with no header addresses set, use an IP traffic flow.
	if !ethAddrsSet {
		if hdrs.ipv4 != nil && hdrs.ipv4.GetSrcAddr() == nil && hdrs.ipv4.GetDstAddr() == nil {
			return ipv4Traffic, nil
		}
		if hdrs.ipv6 != nil && hdrs.ipv6.GetSrcAddr() == nil && hdrs.ipv6.GetDstAddr() == nil {
			return ipv6Traffic, nil
		}
	}

	if hasRSVPEP(srcEPs) || hasRSVPEP(dstEPs) {
		return "", fmt.Errorf("cannot use RSVP endpoint for non-IP traffic")
	}

	// Traffic with a Network endpoint cannot not have any addresses set.
	if hasNetworkEP(srcEPs) || hasNetworkEP(dstEPs) {
		if ethAddrsSet {
			return "", fmt.Errorf("addresses of the initial Ethernet header should not be set when using generated endpoints")
		}
		// IP traffic types would have already been detected above, so must be Ethernet here.
		return ethTraffic, nil
	}

	// For all other flows, use raw traffic so that we can infer the gateway MAC.
	return rawTraffic, nil
}

func hasNetworkEP(eps []*opb.Flow_Endpoint) bool {
	for _, ep := range eps {
		switch ep.GetGenerated().(type) {
		case *opb.Flow_Endpoint_NetworkName:
			return true
		}
	}
	return false
}

func hasRSVPEP(eps []*opb.Flow_Endpoint) bool {
	for _, ep := range eps {
		switch ep.GetGenerated().(type) {
		case *opb.Flow_Endpoint_RsvpName:
			return true
		}
	}
	return false
}

func inferAddresses(hdrs *headers, srcEPs, dstEPs []*opb.Flow_Endpoint, intfs map[string]*intf) {
	if hdrs.eth.GetSrcAddr() == nil {
		hdrs.eth.SrcAddr = inferAddr(srcEPs, intfs, func(i *intf) string { return i.ethMac })
	}
	if hdrs.eth.GetDstAddr() == nil {
		gatewayMac := inferAddr(srcEPs, intfs, func(i *intf) string { return i.resolvedIpv4Mac })
		if gatewayMac == nil {
			gatewayMac = inferAddr(srcEPs, intfs, func(i *intf) string { return i.resolvedIpv6Mac })
		}
		hdrs.eth.DstAddr = gatewayMac
	}
	if hdrs.ipv4 != nil {
		if hdrs.ipv4.GetSrcAddr() == nil {
			hdrs.ipv4.SrcAddr = inferAddr(srcEPs, intfs, ipv4Addr)
		}
		if hdrs.ipv4.GetDstAddr() == nil {
			hdrs.ipv4.DstAddr = inferAddr(dstEPs, intfs, ipv4Addr)
		}
	}
	if hdrs.ipv6 != nil {
		if hdrs.ipv6.GetSrcAddr() == nil {
			hdrs.ipv6.SrcAddr = inferAddr(srcEPs, intfs, ipv6Addr)
		}
		if hdrs.ipv6.GetDstAddr() == nil {
			hdrs.ipv6.DstAddr = inferAddr(dstEPs, intfs, ipv6Addr)
		}
	}
}

func inferAddr(eps []*opb.Flow_Endpoint, intfs map[string]*intf, addrFn func(*intf) string) *opb.AddressRange {
	for _, ep := range eps {
		intf, _ := intfs[ep.GetInterfaceName()]
		a := addrFn(intf)
		if a != "" {
			return &opb.AddressRange{Min: a, Max: a, Count: 1}
		}
	}
	// Give up trying to infer it.
	return nil
}

func ipv4Addr(intf *intf) string {
	if intf.ipv4 != nil {
		if ip := intf.ipv4.Address.SingleValue.Value; ip != nil {
			return *ip
		}
	}
	return ""
}

func ipv6Addr(intf *intf) string {
	if intf.ipv6 != nil {
		if ip := intf.ipv6.Address.SingleValue.Value; ip != nil {
			return *ip
		}
	}
	return ""
}

func portOrLagEPs(eps []*opb.Flow_Endpoint, intfs map[string]*intf, _ bool) ([]ixconfig.IxiaCfgNode, error) {
	visited := make(map[ixconfig.IxiaCfgNode]bool)
	var nodes []ixconfig.IxiaCfgNode
	for _, ep := range eps {
		if vl := intfs[ep.GetInterfaceName()].link; !visited[vl] {
			var n ixconfig.IxiaCfgNode
			switch v := vl.(type) {
			case *ixconfig.Vport:
				n = v.Protocols
			case *ixconfig.Lag:
				n = v
			default:
				return nil, fmt.Errorf("configured link on interface %q is not a Vport or Lag", ep.GetInterfaceName())
			}
			nodes = append(nodes, n)
			visited[vl] = true
		}
	}
	return nodes, nil
}

func devOrGeneratedEPs(eps []*opb.Flow_Endpoint, intfs map[string]*intf, isSrcEP bool) ([]ixconfig.IxiaCfgNode, error) {
	visited := make(map[ixconfig.IxiaCfgNode]bool)
	var nodes []ixconfig.IxiaCfgNode
	for _, ep := range eps {
		var n ixconfig.IxiaCfgNode
		intf, _ := intfs[ep.GetInterfaceName()]
		switch ept := ep.GetGenerated().(type) {
		case nil:
			n = intf.deviceGroup
		case *opb.Flow_Endpoint_NetworkName:
			netg, ok := intf.netToNetworkGroup[ep.GetNetworkName()]
			if !ok {
				return nil, fmt.Errorf("no network group associated with endpoint %v", ep)
			}
			n = netg
		case *opb.Flow_Endpoint_RsvpName:
			rsvp, ok := intf.rsvpLSPs[ep.GetRsvpName()]
			if !ok {
				return nil, fmt.Errorf("no RSVP config associated with endpoint %v", ep)
			}
			// LSPs are unidirectional, so need to distinguish between src/dest endpoints.
			n = rsvp.RsvpP2PEgressLsps
			if isSrcEP {
				n = rsvp.RsvpP2PIngressLsps
			}
		default:
			return nil, fmt.Errorf("unrecognized endpoint type %T", ept)
		}
		if !visited[n] {
			nodes = append(nodes, n)
			visited[n] = true
		}
	}
	return nodes, nil
}

func frameRate(fr *opb.FrameRate) (*ixconfig.TrafficTrafficItemConfigElementFrameRate, map[string]any, error) {
	if fr == nil || fr.Type == nil {
		return nil, nil, nil
	}

	switch frt := fr.Type.(type) {
	case *opb.FrameRate_Percent:
		return &ixconfig.TrafficTrafficItemConfigElementFrameRate{
			Type_: ixconfig.String("percentLineRate"),
			Rate:  ixconfig.NumberFloat64(fr.GetPercent()),
		}, map[string]any{"rate": fr.GetPercent()}, nil
	case *opb.FrameRate_Bps:
		return &ixconfig.TrafficTrafficItemConfigElementFrameRate{
				Type_:            ixconfig.String("bitsPerSecond"),
				BitRateUnitsType: ixconfig.String("bitsPerSec"),
				Rate:             ixconfig.NumberUint64(fr.GetBps()),
			}, map[string]any{
				"bitRateUnitsType": "bitsPerSec",
				"rate":             fr.GetBps(),
			}, nil
	case *opb.FrameRate_Fps:
		return &ixconfig.TrafficTrafficItemConfigElementFrameRate{
			Type_: ixconfig.String("framesPerSecond"),
			Rate:  ixconfig.NumberUint64(fr.GetFps()),
		}, map[string]any{"rate": fr.GetFps()}, nil
	default:
		return nil, nil, fmt.Errorf("unrecognized FrameRate type %T", frt)
	}
}

func frameSize(fs *opb.FrameSize) (*ixconfig.TrafficTrafficItemConfigElementFrameSize, map[string]any, error) {
	if fs == nil || fs.Type == nil {
		return nil, nil, nil
	}

	tfs := &ixconfig.TrafficTrafficItemConfigElementFrameSize{
		// Set the following values to reduce warnings on traffic config import.
		WeightedPairs: []float32{},
		QuadGaussian:  []float32{},
	}
	fsMap := map[string]any{}

	switch fst := fs.Type.(type) {
	case *opb.FrameSize_ImixPreset_:
		preset, ok := imixPreset[fs.GetImixPreset()]
		if !ok {
			return nil, nil, fmt.Errorf("invalid preset %v specified for frame size", fs.GetImixPreset())
		}
		tfs.Type_ = ixconfig.String("presetDistribution")
		tfs.PresetDistribution = ixconfig.String(preset)
		fsMap["presetDistribution"] = preset
	case *opb.FrameSize_Fixed:
		tfs.Type_ = ixconfig.String("fixed")
		tfs.FixedSize = ixconfig.NumberUint32(fs.GetFixed())
		fsMap["fixedSize"] = fs.GetFixed()
	case *opb.FrameSize_Random_:
		tfs.Type_ = ixconfig.String("random")
		tfs.RandomMin = ixconfig.NumberUint32(fs.GetRandom().GetMin())
		tfs.RandomMax = ixconfig.NumberUint32(fs.GetRandom().GetMax())
		fsMap["randomMin"] = fs.GetRandom().GetMin()
		fsMap["randomMax"] = fs.GetRandom().GetMax()
	case *opb.FrameSize_ImixCustom_:
		tfs.Type_ = ixconfig.String("weightedPairs")
		for _, entry := range fs.GetImixCustom().GetEntries() {
			tfs.WeightedPairs = append(tfs.WeightedPairs, float32(entry.Size), float32(entry.Weight))
		}
		fsMap["weightedPairs"] = tfs.WeightedPairs
	default:
		return nil, nil, fmt.Errorf("unrecognized FrameSize type %T", fst)
	}
	return tfs, fsMap, nil
}

func transmissionControl(tc *opb.Transmission) (*ixconfig.TrafficTrafficItemConfigElementTransmissionControl, error) {
	if tc == nil {
		return nil, nil
	}
	switch tcp := tc.GetPattern(); tcp {
	case opb.Transmission_CONTINUOUS:
		if tc.GetPacketsPerBurst() != 0 {
			return nil, fmt.Errorf("burst packet count should not be set for continuous transmissions")
		}
		if tc.GetInterburstGap() != nil {
			return nil, fmt.Errorf("burst gap should not be set for continuous transmissions")
		}
		return &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:       ixconfig.String("continuous"),
			MinGapBytes: ixconfig.NumberUint32(tc.GetMinGapBytes()),
		}, nil
	case opb.Transmission_BURST:
		if tc.GetPacketsPerBurst() == 0 {
			return nil, fmt.Errorf("burst packet count must be a positive value for burst transmissions")
		}
		if tc.GetInterburstGap() == nil {
			return nil, fmt.Errorf("burst gap must be set for burst transmissions")
		}
		var burstGap uint32
		var burstGapUnits string
		switch ibg := tc.GetInterburstGap().(type) {
		case *opb.Transmission_Bytes:
			burstGap = ibg.Bytes
			burstGapUnits = "bytes"
		case *opb.Transmission_Nanoseconds:
			burstGap = ibg.Nanoseconds
			burstGapUnits = "nanoseconds"
		default:
			return nil, fmt.Errorf("unrecognized burst gap type %T", ibg)
		}
		return &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:               ixconfig.String("custom"),
			MinGapBytes:         ixconfig.NumberUint32(tc.GetMinGapBytes()),
			BurstPacketCount:    ixconfig.NumberUint32(tc.GetPacketsPerBurst()),
			EnableInterBurstGap: ixconfig.Bool(true),
			InterBurstGap:       ixconfig.NumberUint32(burstGap),
			InterBurstGapUnits:  ixconfig.String(burstGapUnits),
		}, nil
	case opb.Transmission_FIXED_FRAME_COUNT:
		if tc.GetPacketsPerBurst() != 0 {
			return nil, fmt.Errorf("burst packet count should not be set for fixed packet count transmissions")
		}
		if tc.GetInterburstGap() != nil {
			return nil, fmt.Errorf("burst gap should not be set for fixed packet count transmissions")
		}
		return &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:       ixconfig.String("fixedFrameCount"),
			MinGapBytes: ixconfig.NumberUint32(tc.GetMinGapBytes()),
			FrameCount:  ixconfig.NumberUint32(tc.GetFrameCount()),
		}, nil
	case opb.Transmission_FIXED_DURATION:
		if tc.GetPacketsPerBurst() != 0 {
			return nil, fmt.Errorf("burst packet count should not be set for fixed duration transmissions")
		}
		if tc.GetInterburstGap() != nil {
			return nil, fmt.Errorf("burst gap should not be set for fixed duration transmissions")
		}
		return &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:       ixconfig.String("fixedDuration"),
			MinGapBytes: ixconfig.NumberUint32(tc.GetMinGapBytes()),
			Duration:    ixconfig.NumberUint32(tc.GetDurationSecs()),
		}, nil
	default:
		return nil, fmt.Errorf("unrecognized transmission pattern %v", tcp)
	}
}

func ingressTrackingCfg(f *opb.Flow, trafType trafficType) (*ixconfig.TrafficTrafficItemTracking, bool, error) {
	tracking := &ixconfig.TrafficTrafficItemTracking{
		TrackBy: []string{"trackingenabled0"},
		// Set this value to reduce warnings on traffic config push.
		Values: []string{},
	}

	filter := f.GetIngressTrackingFilters()

	if f.GetConvergenceTracking() {
		if filter != nil {
			return nil, false, fmt.Errorf("Ixia does not support both ingress and convergence tracking tracking for the same flow")
		}
		if trafType == rawTraffic {
			return nil, false, fmt.Errorf("Ixia does not support convergence tracking for raw traffic flows")
		}
		tracking.TrackBy = append(tracking.TrackBy, "destEndpoint0", "destSessionDescription0")
	}

	if filter == nil {
		return tracking, false, nil
	}

	if trafType == rawTraffic && (filter.GetDstEndpoint() || filter.GetSrcEndpoint()) {
		return nil, false, fmt.Errorf("Ixia does not support ingress tracking by src/dst endpoint for raw traffic flows")
	}

	if filter.GetMplsLabel() {
		tracking.TrackBy = append(tracking.TrackBy, "mplsMplsLabelValue0")
	}
	if filter.GetSrcEndpoint() {
		tracking.TrackBy = append(tracking.TrackBy, "sourceEndpoint0")
	}
	if filter.GetDstEndpoint() {
		tracking.TrackBy = append(tracking.TrackBy, "destEndpoint0")
	}
	if filter.GetSrcIpv4() {
		tracking.TrackBy = append(tracking.TrackBy, "ipv4SourceIp0")
	}
	if filter.GetDstIpv4() {
		tracking.TrackBy = append(tracking.TrackBy, "ipv4DestIp0")
	}
	if filter.GetSrcIpv6() {
		tracking.TrackBy = append(tracking.TrackBy, "ipv6SourceIp0")
	}
	if filter.GetDstIpv6() {
		tracking.TrackBy = append(tracking.TrackBy, "ipv6DestIp0")
	}
	if filter.GetVlanId() {
		tracking.TrackBy = append(tracking.TrackBy, "vlanVlanId0")
	}
	return tracking, len(tracking.TrackBy) > 1, nil
}
