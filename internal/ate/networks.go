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
	"net"
	"strings"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	originToStr = map[opb.BgpAttributes_Origin]string{
		opb.BgpAttributes_ORIGIN_IGP:        "igp",
		opb.BgpAttributes_ORIGIN_EGP:        "egp",
		opb.BgpAttributes_ORIGIN_INCOMPLETE: "incomplete",
	}
	coBitsToStr = map[opb.BgpAttributes_ExtendedCommunity_Color_CoBits]string{
		opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00: "00",
		opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_01: "01",
		opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_10: "10",
		opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_11: "11",
	}
	asPathSegTypeToStr = map[opb.BgpAttributes_AsPathSegment_Type]string{
		opb.BgpAttributes_AsPathSegment_TYPE_AS_SET:               "asset",
		opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ:               "asseq",
		opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ_CONFEDERATION: "asseqconfederation",
		opb.BgpAttributes_AsPathSegment_TYPE_AS_SET_CONFEDERATION: "assetconfederation",
	}
)

// addNetworks adds IxNetwork network groups for the given interface config.
func (ix *ixATE) addNetworks(ifc *opb.InterfaceConfig) error {
	intf := ix.intfs[ifc.GetName()]
	intf.netToNetworkGroup = make(map[string]*ixconfig.TopologyNetworkGroup)
	intf.netToRouteTables = make(map[string]*routeTables)
	dg := intf.deviceGroup

	var isisL3 *ixconfig.TopologyIsisL3
	if isisL3s := intf.deviceGroup.Ethernet[0].IsisL3; len(isisL3s) > 0 {
		isisL3 = isisL3s[0]
	}
	hasBgpCfg := false
	for _, netCfg := range ifc.GetNetworks() {
		if netCfg.GetIsis() != nil && isisL3 == nil {
			return fmt.Errorf("cannot add IS-IS attributes for network group %q without IS-IS configured", netCfg.GetName())
		}
		if netCfg.GetBgpAttributes() != nil {
			hasBgpCfg = true
		}
	}

	for _, netCfg := range ifc.GetNetworks() {
		ng := &ixconfig.TopologyNetworkGroup{Name: ixconfig.String(netCfg.GetName())}
		if eth := netCfg.GetEth(); eth != nil {
			ng.MacPools = []*ixconfig.TopologyMacPools{{
				Mac:                  ixconfig.MultivalueStr(eth.GetMacAddress()),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(eth.GetCount()),
				EnableVlans:          ixconfig.MultivalueBool(eth.GetVlanId() != 0),
				Vlan: []*ixconfig.TopologyVlan{{
					VlanId: ixconfig.MultivalueUint32(eth.GetVlanId()),
				}},
			}}
		}

		var err error
		var v4Pools []*ixconfig.TopologyIpv4PrefixPools
		var v6Pools []*ixconfig.TopologyIpv6PrefixPools

		hasBGPv4Peer, hasBGPv6Peer := false, false
		if (intf.ipv4 != nil && len(intf.ipv4.BgpIpv4Peer) > 0) || (intf.ipv4Loopback != nil && len(intf.ipv4Loopback.BgpIpv4Peer) > 0) {
			hasBGPv4Peer = true
		}
		if (intf.ipv6 != nil && len(intf.ipv6.BgpIpv6Peer) > 0) || (intf.ipv6Loopback != nil && len(intf.ipv6Loopback.BgpIpv6Peer) > 0) {
			hasBGPv6Peer = true
		}
		if netCfg.GetImportedBgpRoutes() != nil {
			v4Pools, v6Pools, err = importedBGPRoutePools(intf.netToRouteTables, netCfg, hasBGPv4Peer, hasBGPv6Peer)
			if err != nil {
				return err
			}
		} else {
			v4Pools, err = ipv4Pools(netCfg, isisL3, hasBgpCfg)
			if err != nil {
				return err
			}

			v6Pools, err = ipv6Pools(netCfg, isisL3, hasBgpCfg)
			if err != nil {
				return err
			}
		}
		ng.Ipv4PrefixPools = v4Pools
		ng.Ipv6PrefixPools = v6Pools

		intf.netToNetworkGroup[netCfg.GetName()] = ng
		dg.NetworkGroup = append(dg.NetworkGroup, ng)
	}
	return nil
}

func routeOriginStr(origin opb.IPReachability_RouteOrigin) (string, error) {
	switch origin {
	case opb.IPReachability_ROUTE_ORIGIN_UNSPECIFIED:
		return "", fmt.Errorf("route origin not specified")
	case opb.IPReachability_INTERNAL:
		return "internal", nil
	case opb.IPReachability_EXTERNAL:
		return "external", nil
	default:
		return "", fmt.Errorf("unrecognized route origin %s", origin)
	}
}

func isisRouteProp(ipr *opb.IPReachability) (*ixconfig.TopologyIsisL3RouteProperty, error) {
	origin, err := routeOriginStr(ipr.GetRouteOrigin())
	if err != nil {
		return nil, err
	}
	return &ixconfig.TopologyIsisL3RouteProperty{
		Metric:                 ixconfig.MultivalueUint32(ipr.GetMetric()),
		Algorithm:              ixconfig.MultivalueUint32(ipr.GetAlgorithm()),
		RouteOrigin:            ixconfig.MultivalueStr(origin),
		Active:                 ixconfig.MultivalueBool(ipr.GetActive()),
		ConfigureSIDIndexLabel: ixconfig.MultivalueBool(ipr.GetEnableSidIndexLabel()),
		SIDIndexLabel:          ixconfig.MultivalueUint32(ipr.GetSidIndexLabel()),
		RFlag:                  ixconfig.MultivalueBool(ipr.GetFlagReadvertise()),
		NFlag:                  ixconfig.MultivalueBool(ipr.GetFlagNodeSid()),
		PFlag:                  ixconfig.MultivalueBool(ipr.GetFlagNoPhp()),
		EFlag:                  ixconfig.MultivalueBool(ipr.GetFlagExplicitNull()),
		VFlag:                  ixconfig.MultivalueBool(ipr.GetFlagValue()),
		LFlag:                  ixconfig.MultivalueBool(ipr.GetFlagLocal()),
		IsisL3PrefixesSrSid: &ixconfig.TopologyIsisL3PrefixesSrSid{
			Active: ixconfig.MultivalueBool(ipr.GetEnableSidIndexLabel()),
		},
	}, nil
}

func bgpComms(communities *opb.BgpCommunities) ([]*ixconfig.TopologyBgpCommunitiesList, error) {
	var comms []*ixconfig.TopologyBgpCommunitiesList
	if communities.GetNoExport() {
		comms = append(comms, &ixconfig.TopologyBgpCommunitiesList{
			Type_: ixconfig.MultivalueStr("noexport"),
		})
	}
	if communities.GetNoAdvertise() {
		comms = append(comms, &ixconfig.TopologyBgpCommunitiesList{
			Type_: ixconfig.MultivalueStr("noadvertised"),
		})
	}
	if communities.GetNoExportSubconfed() {
		comms = append(comms, &ixconfig.TopologyBgpCommunitiesList{
			Type_: ixconfig.MultivalueStr("noexport_subconfed"),
		})
	}
	if communities.GetLlgrStale() {
		comms = append(comms, &ixconfig.TopologyBgpCommunitiesList{
			Type_: ixconfig.MultivalueStr("llgr_stale"),
		})
	}
	if communities.GetNoLlgr() {
		comms = append(comms, &ixconfig.TopologyBgpCommunitiesList{
			Type_: ixconfig.MultivalueStr("no_llgr"),
		})
	}
	for _, comm := range communities.GetPrivateCommunities() {
		commSplit := strings.SplitN(comm, ":", 2)
		if len(commSplit) != 2 {
			return nil, fmt.Errorf("invalid format for BGP community %q", comm)
		}
		comms = append(comms, &ixconfig.TopologyBgpCommunitiesList{
			AsNumber:      ixconfig.MultivalueStr(commSplit[0]),
			LastTwoOctets: ixconfig.MultivalueStr(commSplit[1]),
			Type_:         ixconfig.MultivalueStr("manual"),
		})
	}
	return comms, nil
}

func bgpExtComms(extCommPBs []*opb.BgpAttributes_ExtendedCommunity) ([]*ixconfig.TopologyBgpExtendedCommunitiesList, error) {
	var extComms []*ixconfig.TopologyBgpExtendedCommunitiesList
	for _, commPB := range extCommPBs {
		extComm := &ixconfig.TopologyBgpExtendedCommunitiesList{}
		switch commPB.Type.(type) {
		case *opb.BgpAttributes_ExtendedCommunity_Color_:
			color := commPB.GetColor()
			coBits, ok := coBitsToStr[color.GetCoBits()]
			if !ok {
				return nil, fmt.Errorf("invalid extended communited color bits value %s", color.GetCoBits())
			}
			extComm.Type_ = ixconfig.MultivalueStr("opaque")
			extComm.SubType = ixconfig.MultivalueStr("color")
			extComm.ColorCOBits = ixconfig.MultivalueStr(coBits)
			extComm.ColorReservedBits = ixconfig.MultivalueUint32(color.GetReservedBits())
			extComm.ColorValue = ixconfig.MultivalueUint32(color.GetValue())
		default:
			return nil, fmt.Errorf("unrecognized extended community type %s", commPB.GetType())
		}
		extComms = append(extComms, extComm)
	}
	return extComms, nil
}

func bgpASPathSegments(asPathSegPBs []*opb.BgpAttributes_AsPathSegment) ([]*ixconfig.TopologyBgpAsPathSegmentList, error) {
	var asPathSegs []*ixconfig.TopologyBgpAsPathSegmentList
	for _, asPathSegPB := range asPathSegPBs {
		asPathSegType, ok := asPathSegTypeToStr[asPathSegPB.GetType()]
		if !ok {
			return nil, fmt.Errorf("invalid AS path segment type %s", asPathSegPB.GetType())
		}
		asPathSeg := &ixconfig.TopologyBgpAsPathSegmentList{
			EnableASPathSegment:       ixconfig.MultivalueTrue(),
			SegmentType:               ixconfig.MultivalueStr(asPathSegType),
			NumberOfAsNumberInSegment: ixconfig.NumberInt(len(asPathSegPB.GetAsns())),
		}
		for _, asn := range asPathSegPB.GetAsns() {
			asPathSeg.BgpAsNumberList = append(asPathSeg.BgpAsNumberList, &ixconfig.TopologyBgpAsNumberList{
				EnableASNumber: ixconfig.MultivalueTrue(),
				AsNumber:       ixconfig.MultivalueUint32(asn),
			})
		}
		asPathSegs = append(asPathSegs, asPathSeg)
	}
	return asPathSegs, nil
}

func originatorID(origID *opb.StringIncRange) (string, string, error) {
	startIP, isV6 := parseIP(origID.GetStart())
	if startIP == nil || isV6 {
		return "", "", fmt.Errorf("originator ID start IP %q is not a valid IPv4 address", origID.GetStart())
	}
	stepIP, isV6 := parseIP(origID.GetStep())
	if stepIP == nil || isV6 {
		return "", "", fmt.Errorf("originator ID step %q is not a valid IPv4 address", origID.GetStep())
	}
	return origID.GetStart(), origID.GetStep(), nil
}

func bgpV4RouteProp(bgp *opb.BgpAttributes) (*ixconfig.TopologyBgpIpRouteProperty, error) {
	brp := &ixconfig.TopologyBgpIpRouteProperty{
		Active:                ixconfig.MultivalueBool(bgp.GetActive()),
		EnableNextHop:         ixconfig.MultivalueTrue(),
		EnableOrigin:          ixconfig.MultivalueTrue(),
		EnableLocalPreference: ixconfig.MultivalueTrue(),
		LocalPreference:       ixconfig.MultivalueUint32(bgp.GetLocalPreference()),
		NoOfLargeCommunities:  ixconfig.NumberUint32(0),
	}

	if nh := bgp.GetNextHopAddress(); nh != "" {
		brp.NextHopType = ixconfig.MultivalueStr("manual")
		if _, isIPv6 := parseIP(nh); isIPv6 {
			brp.Ipv6NextHop = ixconfig.MultivalueStr(nh)
			brp.AdvertiseNexthopAsV4 = ixconfig.MultivalueFalse()
			brp.NextHopIPType = ixconfig.MultivalueStr("ipv6")
		} else {
			brp.Ipv4NextHop = ixconfig.MultivalueStr(nh)
			brp.AdvertiseNexthopAsV4 = ixconfig.MultivalueTrue()
			brp.NextHopIPType = ixconfig.MultivalueStr("ipv4")
		}
	} else {
		brp.NextHopType = ixconfig.MultivalueStr("sameaslocalip")
	}

	origin, ok := originToStr[bgp.GetOrigin()]
	if !ok {
		return nil, fmt.Errorf("invalid BGP route origin %s", bgp.GetOrigin())
	}
	brp.Origin = ixconfig.MultivalueStr(origin)

	comms, err := bgpComms(bgp.GetCommunities())
	if err != nil {
		return nil, err
	}
	brp.EnableCommunity = ixconfig.MultivalueBool(len(comms) != 0)
	brp.NoOfCommunities = ixconfig.NumberInt(len(comms))
	brp.BgpCommunitiesList = comms

	extComms, err := bgpExtComms(bgp.GetExtendedCommunities())
	if err != nil {
		return nil, err
	}
	brp.EnableExtendedCommunity = ixconfig.MultivalueBool(len(extComms) != 0)
	brp.NoOfExternalCommunities = ixconfig.NumberInt(len(extComms))
	brp.BgpExtendedCommunitiesList = extComms

	asnSetMode, ok := asnSetModeToStr[bgp.GetAsnSetMode()]
	if !ok {
		return nil, fmt.Errorf("invalid BGP ASN set mode %s", bgp.GetAsnSetMode())
	}
	brp.AsSetMode = ixconfig.MultivalueStr(asnSetMode)

	asPathSegs, err := bgpASPathSegments(bgp.GetAsPathSegments())
	if err != nil {
		return nil, err
	}
	brp.EnableAsPathSegments = ixconfig.MultivalueBool(len(asPathSegs) != 0)
	brp.NoOfASPathSegmentsPerRouteRange = ixconfig.NumberInt(len(asPathSegs))
	brp.BgpAsPathSegmentList = asPathSegs

	if origID := bgp.GetOriginatorId(); origID != nil {
		brp.EnableOriginatorId = ixconfig.MultivalueTrue()
		start, step, err := originatorID(origID)
		if err != nil {
			return nil, err
		}
		brp.OriginatorId = ixconfig.MultivalueStrIncCounter(start, step)
	}

	if clusterIDs := bgp.GetClusterIds(); len(clusterIDs) > 0 {
		brp.EnableCluster = ixconfig.MultivalueTrue()
		brp.NoOfClusters = ixconfig.NumberInt(len(clusterIDs))
		for _, ci := range clusterIDs {
			cIP, isV6 := parseIP(ci)
			if cIP == nil || isV6 {
				return nil, fmt.Errorf("cluster ID %q is not a valid IPv4 address", ci)
			}
			brp.BgpClusterIdList = append(brp.BgpClusterIdList, &ixconfig.TopologyBgpClusterIdList{ClusterId: ixconfig.MultivalueStr(ci)})
		}
	}
	return brp, nil
}

func bgpV6RouteProp(bgp *opb.BgpAttributes) (*ixconfig.TopologyBgpV6IpRouteProperty, error) {
	brp := &ixconfig.TopologyBgpV6IpRouteProperty{
		Active:                ixconfig.MultivalueBool(bgp.GetActive()),
		EnableNextHop:         ixconfig.MultivalueTrue(),
		EnableOrigin:          ixconfig.MultivalueTrue(),
		EnableLocalPreference: ixconfig.MultivalueTrue(),
		LocalPreference:       ixconfig.MultivalueUint32(bgp.GetLocalPreference()),
		NoOfLargeCommunities:  ixconfig.NumberUint32(0),
	}

	if nh := bgp.GetNextHopAddress(); nh != "" {
		brp.NextHopType = ixconfig.MultivalueStr("manual")
		if _, isIPv6 := parseIP(nh); isIPv6 {
			brp.Ipv6NextHop = ixconfig.MultivalueStr(nh)
			brp.AdvertiseNexthopAsV4 = ixconfig.MultivalueFalse()
			brp.NextHopIPType = ixconfig.MultivalueStr("ipv6")
		} else {
			brp.Ipv4NextHop = ixconfig.MultivalueStr(nh)
			brp.AdvertiseNexthopAsV4 = ixconfig.MultivalueTrue()
			brp.NextHopIPType = ixconfig.MultivalueStr("ipv4")
		}
	} else {
		brp.NextHopType = ixconfig.MultivalueStr("sameaslocalip")
	}

	origin, ok := originToStr[bgp.GetOrigin()]
	if !ok {
		return nil, fmt.Errorf("invalid BGP route origin %s", bgp.GetOrigin())
	}
	brp.Origin = ixconfig.MultivalueStr(origin)

	comms, err := bgpComms(bgp.GetCommunities())
	if err != nil {
		return nil, err
	}
	brp.EnableCommunity = ixconfig.MultivalueBool(len(comms) != 0)
	brp.NoOfCommunities = ixconfig.NumberInt(len(comms))
	brp.BgpCommunitiesList = comms

	extComms, err := bgpExtComms(bgp.GetExtendedCommunities())
	if err != nil {
		return nil, err
	}
	brp.EnableExtendedCommunity = ixconfig.MultivalueBool(len(extComms) != 0)
	brp.NoOfExternalCommunities = ixconfig.NumberInt(len(extComms))
	brp.BgpExtendedCommunitiesList = extComms

	asnSetMode, ok := asnSetModeToStr[bgp.GetAsnSetMode()]
	if !ok {
		return nil, fmt.Errorf("invalid BGP ASN set mode %s", bgp.GetAsnSetMode())
	}
	brp.AsSetMode = ixconfig.MultivalueStr(asnSetMode)

	asPathSegs, err := bgpASPathSegments(bgp.GetAsPathSegments())
	if err != nil {
		return nil, err
	}
	brp.EnableAsPathSegments = ixconfig.MultivalueBool(len(asPathSegs) != 0)
	brp.NoOfASPathSegmentsPerRouteRange = ixconfig.NumberInt(len(asPathSegs))
	brp.BgpAsPathSegmentList = asPathSegs

	if origID := bgp.GetOriginatorId(); origID != nil {
		brp.EnableOriginatorId = ixconfig.MultivalueTrue()
		start, step, err := originatorID(origID)
		if err != nil {
			return nil, err
		}
		brp.OriginatorId = ixconfig.MultivalueStrIncCounter(start, step)
	}

	if clusterIDs := bgp.GetClusterIds(); len(clusterIDs) > 0 {
		brp.EnableCluster = ixconfig.MultivalueTrue()
		brp.NoOfClusters = ixconfig.NumberInt(len(clusterIDs))
		for _, ci := range clusterIDs {
			cIP, isV6 := parseIP(ci)
			if cIP == nil || isV6 {
				return nil, fmt.Errorf("cluster ID %q is not a valid IPv4 address", ci)
			}
			brp.BgpClusterIdList = append(brp.BgpClusterIdList, &ixconfig.TopologyBgpClusterIdList{ClusterId: ixconfig.MultivalueStr(ci)})
		}
	}
	return brp, nil
}

func routeProps(bgp *opb.BgpAttributes, v6Routes bool) ([]*ixconfig.TopologyBgpIpRouteProperty, []*ixconfig.TopologyBgpV6IpRouteProperty, error) {
	var exportV4, exportV6 bool
	switch bgp.GetAdvertisementProtocol() {
	case opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_UNSPECIFIED:
		return nil, nil, fmt.Errorf("BGP advertisement protocol not set for routes")
	case opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V4:
		exportV4 = true
	case opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V6:
		exportV6 = true
	case opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V4_AND_V6:
		exportV4 = true
		exportV6 = true
	case opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_SAME_AS_ROUTE:
		exportV4 = !v6Routes
		exportV6 = v6Routes
	default:
		return nil, nil, fmt.Errorf("unrecognized advertisement protocol specification: %v", bgp.GetAdvertisementProtocol())
	}
	var brps []*ixconfig.TopologyBgpIpRouteProperty
	var brpsV6 []*ixconfig.TopologyBgpV6IpRouteProperty
	if exportV4 {
		brp, err := bgpV4RouteProp(bgp)
		if err != nil {
			return nil, nil, err
		}
		brps = []*ixconfig.TopologyBgpIpRouteProperty{brp}
	}
	if exportV6 {
		brpV6, err := bgpV6RouteProp(bgp)
		if err != nil {
			return nil, nil, err
		}
		brpsV6 = []*ixconfig.TopologyBgpV6IpRouteProperty{brpV6}
	}
	return brps, brpsV6, nil
}

func ipv4Pools(netCfg *opb.Network, isisL3 *ixconfig.TopologyIsisL3, hasBgpCfg bool) ([]*ixconfig.TopologyIpv4PrefixPools, error) {
	isis := netCfg.GetIsis()
	bgp := netCfg.GetBgpAttributes()
	ipv4 := netCfg.GetIpv4()
	if ipv4 == nil {
		return nil, nil
	}
	if ipv4.GetAddressCidr() == "" {
		return nil, fmt.Errorf("need address defined for IP V4 network group %q", netCfg.GetName())
	}
	ip, netw, err := net.ParseCIDR(ipv4.GetAddressCidr())
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as an IP address: %w", ipv4.GetAddressCidr(), err)
	}
	mask, _ := netw.Mask.Size()

	if isis != nil && bgp != nil {
		return nil, fmt.Errorf("cannot set both IS-IS and BGP attributes for IPv4 network group %q (use multiple network groups if needed)", netCfg.GetName())
	}

	var connector *ixconfig.TopologyConnector
	var irps []*ixconfig.TopologyIsisL3RouteProperty
	if isis != nil {
		irp, err := isisRouteProp(isis)
		if err != nil {
			return nil, err
		}
		irps = []*ixconfig.TopologyIsisL3RouteProperty{irp}
		connector = &ixconfig.TopologyConnector{}
		connector.SetConnectedToRef(isisL3)
	} else if isisL3 != nil {
		// IS-IS route config needs to be present if configured on this interface.
		// Otherwise, Ixnetwork will autocreate an active config.
		irps = []*ixconfig.TopologyIsisL3RouteProperty{{
			Name:   ixconfig.String(fmt.Sprintf("%s IS-IS Inactive", netCfg.GetName())),
			Active: ixconfig.MultivalueFalse(),
		}}
	}

	var brps []*ixconfig.TopologyBgpIpRouteProperty
	var brpsV6 []*ixconfig.TopologyBgpV6IpRouteProperty
	if bgp != nil {
		brps, brpsV6, err = routeProps(bgp, false)
		if err != nil {
			return nil, err
		}
	} else if hasBgpCfg {
		// BGP route config needs to be present if configured on _any_ network for this interface.
		// Otherwise, Ixnetwork will autocreate an active config.
		brps = []*ixconfig.TopologyBgpIpRouteProperty{{
			Name:   ixconfig.String(fmt.Sprintf("%s BGP Inactive", netCfg.GetName())),
			Active: ixconfig.MultivalueFalse(),
		}}
	}

	return []*ixconfig.TopologyIpv4PrefixPools{{
		Connector:            connector,
		NetworkAddress:       ixconfig.MultivalueStr(ip.String()),
		PrefixLength:         ixconfig.MultivalueUint32(uint32(mask)),
		NumberOfAddressesAsy: ixconfig.MultivalueUint32(ipv4.GetCount()),
		IsisL3RouteProperty:  irps,
		BgpIPRouteProperty:   brps,
		BgpV6IPRouteProperty: brpsV6,
	}}, nil
}

func ipv6Pools(netCfg *opb.Network, isisL3 *ixconfig.TopologyIsisL3, hasBgpCfg bool) ([]*ixconfig.TopologyIpv6PrefixPools, error) {
	isis := netCfg.GetIsis()
	bgp := netCfg.GetBgpAttributes()
	ipv6 := netCfg.GetIpv6()
	if ipv6 == nil {
		return nil, nil
	}
	if ipv6.GetAddressCidr() == "" {
		return nil, fmt.Errorf("need address defined for IP V6 network group %q", netCfg.GetName())
	}
	ip, netw, err := net.ParseCIDR(ipv6.GetAddressCidr())
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as an IP address: %w", ipv6.GetAddressCidr(), err)
	}
	mask, _ := netw.Mask.Size()

	if isis != nil && bgp != nil {
		return nil, fmt.Errorf("cannot set both IS-IS and BGP attributes for IPv6 network group %q (use multiple network groups if needed)", netCfg.GetName())
	}

	var connector *ixconfig.TopologyConnector
	var irps []*ixconfig.TopologyIsisL3RouteProperty
	if isis != nil {
		irp, err := isisRouteProp(isis)
		if err != nil {
			return nil, err
		}
		irps = []*ixconfig.TopologyIsisL3RouteProperty{irp}
		connector = &ixconfig.TopologyConnector{}
		connector.SetConnectedToRef(isisL3)
	} else if isisL3 != nil {
		// IS-IS route config needs to be present if configured on this interface.
		// Otherwise, Ixnetwork will autocreate an active config.
		irps = []*ixconfig.TopologyIsisL3RouteProperty{{
			Name:   ixconfig.String(fmt.Sprintf("%s IS-IS Inactive", netCfg.GetName())),
			Active: ixconfig.MultivalueFalse(),
		}}
	}

	var brps []*ixconfig.TopologyBgpIpRouteProperty
	var brpsV6 []*ixconfig.TopologyBgpV6IpRouteProperty
	if bgp != nil {
		brps, brpsV6, err = routeProps(bgp, true)
		if err != nil {
			return nil, err
		}
	} else if hasBgpCfg {
		// BGP route config needs to be present if configured on _any_ network for this interface.
		// Otherwise, Ixnetwork will autocreate an active config.
		brpsV6 = []*ixconfig.TopologyBgpV6IpRouteProperty{{
			Name:   ixconfig.String(fmt.Sprintf("%s BGP V6 Inactive", netCfg.GetName())),
			Active: ixconfig.MultivalueFalse(),
		}}
	}

	return []*ixconfig.TopologyIpv6PrefixPools{{
		Connector:            connector,
		NetworkAddress:       ixconfig.MultivalueStr(ip.String()),
		PrefixLength:         ixconfig.MultivalueUint32(uint32(mask)),
		NumberOfAddressesAsy: ixconfig.MultivalueUint32(ipv6.GetCount()),
		IsisL3RouteProperty:  irps,
		BgpIPRouteProperty:   brps,
		BgpV6IPRouteProperty: brpsV6,
	}}, nil
}

func importedBGPRoutePools(netToRouteTables map[string]*routeTables, netCfg *opb.Network, hasBGPv4Peer, hasBGPv6Peer bool) ([]*ixconfig.TopologyIpv4PrefixPools, []*ixconfig.TopologyIpv6PrefixPools, error) {
	imported := netCfg.GetImportedBgpRoutes()
	if netCfg.GetIsis() != nil || netCfg.GetBgpAttributes() != nil || netCfg.GetIpv4() != nil || netCfg.GetIpv6() != nil {
		return nil, nil, fmt.Errorf("cannot import routes for network group %q with any other routes/attributes configured", netCfg.GetName())
	}

	if !hasBGPv4Peer && !hasBGPv6Peer {
		return nil, nil, fmt.Errorf("cannot import routes for network group %q without associated BGP peer", netCfg.GetName())
	}

	rts := &routeTables{
		ipv4:             imported.GetIpv4RoutesPath(),
		ipv6:             imported.GetIpv6RoutesPath(),
		overwriteNexthop: imported.GetOverwriteNexthop(),
	}

	switch imported.GetRouteTableFormat() {
	case opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_UNSPECIFIED:
		return nil, nil, fmt.Errorf("route table format not specified")
	case opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_CISCO:
		rts.format = routeTableFormatCisco
	case opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_JUNIPER:
		rts.format = routeTableFormatJuniper
	case opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_CSV:
		rts.format = routeTableFormatCSV
	default:
		return nil, nil, fmt.Errorf("unrecognized route table format %s", imported.GetRouteTableFormat())
	}

	netToRouteTables[netCfg.GetName()] = rts
	var ipv4Pools []*ixconfig.TopologyIpv4PrefixPools
	var ipv6Pools []*ixconfig.TopologyIpv6PrefixPools
	if rts.ipv4 != "" {
		ipv4Pool := &ixconfig.TopologyIpv4PrefixPools{}
		if hasBGPv4Peer {
			ipv4Pool.BgpIPRouteProperty = []*ixconfig.TopologyBgpIpRouteProperty{{
				Name:   ixconfig.String(fmt.Sprintf("Imported IPv4 BGP Routes")),
				Active: ixconfig.MultivalueTrue(),
			}}
		}
		if hasBGPv6Peer {
			ipv4Pool.BgpV6IPRouteProperty = []*ixconfig.TopologyBgpV6IpRouteProperty{{
				Name:   ixconfig.String(fmt.Sprintf("Imported IPv4 BGP V6 Routes")),
				Active: ixconfig.MultivalueTrue(),
			}}
		}
		ipv4Pools = []*ixconfig.TopologyIpv4PrefixPools{ipv4Pool}
	}
	if rts.ipv6 != "" {
		ipv6Pool := &ixconfig.TopologyIpv6PrefixPools{}
		if hasBGPv4Peer {
			ipv6Pool.BgpIPRouteProperty = []*ixconfig.TopologyBgpIpRouteProperty{{
				Name:   ixconfig.String(fmt.Sprintf("Imported IPv6 BGP Routes")),
				Active: ixconfig.MultivalueTrue(),
			}}
		}
		if hasBGPv6Peer {
			ipv6Pool.BgpV6IPRouteProperty = []*ixconfig.TopologyBgpV6IpRouteProperty{{
				Name:   ixconfig.String(fmt.Sprintf("Imported IPv6 BGP V6 Routes")),
				Active: ixconfig.MultivalueTrue(),
			}}
		}
		ipv6Pools = []*ixconfig.TopologyIpv6PrefixPools{ipv6Pool}
	}
	return ipv4Pools, ipv6Pools, nil
}
