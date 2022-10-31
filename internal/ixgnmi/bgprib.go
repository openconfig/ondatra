// Copyright 2021 Google Inc.
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
	"fmt"
	"strings"

	log "github.com/golang/glog"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ondatra/telemetry"
)

// bgpLearnedInfo is the output of learned info for BGP from an Ixia device.
// This struct is used for ipv4 and ipv6 info.
type bgpLearnedInfo struct {
	IPV4Prefix     string `ixia:"IPv4 Prefix"`
	IPV6Prefix     string `ixia:"IPv6 Prefix"`
	PrefixLen      int    `ixia:"Prefix Length"`
	PathID         uint32 `ixia:"Path ID"`
	IPV4NextHop    string `ixia:"IPv4 Next Hop"`
	IPV6NextHop    string `ixia:"IPv6 Next Hop"`
	IPV6NextHop2   string `ixia:"IPv6 Next Hop 2"`
	MED            uint32 `ixia:"MED"`
	LocalPref      uint32 `ixia:"Local Preference"`
	Origin         string `ixia:"Origin"`
	ASPath         string `ixia:"AS Path"`
	Community      string `ixia:"Community"`
	AIGP           uint64 `ixia:"AIGP"`
	Color          string `ixia:"Color"`
	LargeCommunity string `ixia:"Large Community"`
}

func bgpRIBDevice(intf, protocol string, peer4Infos, peer6Infos map[string][]bgpLearnedInfo) (*telemetry.Device, error) {
	dev := &telemetry.Device{}
	rib := dev.GetOrCreateNetworkInstance(intf).
		GetOrCreateProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, protocol).
		GetOrCreateBgp().
		GetOrCreateRib()

	for peer, infos := range peer4Infos {
		ipv4RIB := rib.GetOrCreateAfiSafi(telemetry.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST).
			GetOrCreateIpv4Unicast().
			GetOrCreateNeighbor(peer).
			GetOrCreateAdjRibInPre()
		attrLen := len(rib.AttrSet)
		commLen := len(rib.Community)

		for i, info := range infos {
			if info == (bgpLearnedInfo{}) {
				log.Infof("Skipping empty BGP learned info")
				continue
			}
			attrIndex := uint64(i + attrLen)
			commIndex := uint64(i + commLen)
			if err := appendDetails(info, rib, attrIndex, commIndex, true); err != nil {
				return nil, fmt.Errorf("failed to append details for elem %d: %w", i, err)
			}
			route := &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route{
				Prefix:         ygot.String(fmt.Sprintf("%s/%d", info.IPV4Prefix, info.PrefixLen)),
				PathId:         ygot.Uint32(info.PathID),
				AttrIndex:      &attrIndex,
				CommunityIndex: &commIndex,
			}
			if err := ipv4RIB.AppendRoute(route); err != nil {
				return nil, fmt.Errorf("failed to append route for elem %d: %w", i, err)
			}
		}
	}

	for peer, infos := range peer6Infos {
		ipv6RIB := rib.GetOrCreateAfiSafi(telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST).
			GetOrCreateIpv6Unicast().
			GetOrCreateNeighbor(peer).
			GetOrCreateAdjRibInPre()
		attrLen := len(rib.AttrSet)
		commLen := len(rib.Community)

		for i, info := range infos {
			if info == (bgpLearnedInfo{}) {
				log.Infof("Skipping empty BGP learned info")
				continue
			}
			attrIndex := uint64(i + attrLen)
			commIndex := uint64(i + commLen)
			if err := appendDetails(info, rib, attrIndex, commIndex, false); err != nil {
				return nil, fmt.Errorf("failed to append details for elem %d: %w", i, err)
			}
			route := &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route{
				Prefix:         ygot.String(fmt.Sprintf("%s/%d", info.IPV6Prefix, info.PrefixLen)),
				PathId:         ygot.Uint32(info.PathID),
				AttrIndex:      &attrIndex,
				CommunityIndex: &commIndex,
			}
			if err := ipv6RIB.AppendRoute(route); err != nil {
				return nil, fmt.Errorf("failed to append route for elem %d: %w", i, err)
			}
		}
	}
	return dev, nil
}

func appendDetails(info bgpLearnedInfo, rib *telemetry.NetworkInstance_Protocol_Bgp_Rib, attrIndex, commIndex uint64, v4 bool) error {
	nextHop := info.IPV6NextHop
	if v4 {
		nextHop = info.IPV4NextHop
	}

	as := &telemetry.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
		NextHop:   &nextHop,
		Aigp:      ygot.Uint64(info.AIGP),
		LocalPref: ygot.Uint32(info.LocalPref),
		Med:       ygot.Uint32(info.MED),
		Index:     &attrIndex,
	}
	switch info.Origin {
	case "IGP":
		as.Origin = telemetry.BgpTypes_BgpOriginAttrType_IGP
	case "EGP":
		as.Origin = telemetry.BgpTypes_BgpOriginAttrType_EGP
	case "Incomplete":
		as.Origin = telemetry.BgpTypes_BgpOriginAttrType_INCOMPLETE
	default:
		return fmt.Errorf("unknown origin type: %q", info.Origin)
	}
	if err := rib.AppendAttrSet(as); err != nil {
		return err
	}

	community := &telemetry.NetworkInstance_Protocol_Bgp_Rib_Community{
		Index: &commIndex,
	}
	for _, str := range strings.Split(info.Community, ",") {
		if str != "" {
			community.Community = append(community.Community, telemetry.UnionString(str))
		}
	}
	return rib.AppendCommunity(community)
}
