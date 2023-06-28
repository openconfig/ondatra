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
	"path"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ygot/ygot"
)

func bgpRIBFromIxia(ctx context.Context, client cfgClient, netInst *oc.NetworkInstance, nodes *cachedNodes) error {
	const (
		bgp4Path = "topology/deviceGroup/ethernet/ipv4/bgpIpv4Peer/operations/getAllLearnedInfo"
		bgp6Path = "topology/deviceGroup/ethernet/ipv6/bgpIpv6Peer/operations/getAllLearnedInfo"
	)
	peer4Infos, err := fetchBGPInfos(ctx, client, bgp4Path, nodes.bgp4Peers)
	if err != nil {
		return fmt.Errorf("failed to read bgp rib v4 learned info: %w", err)
	}
	peer6Infos, err := fetchBGPInfos(ctx, client, bgp6Path, nodes.bgp6Peers)
	if err != nil {
		return fmt.Errorf("failed to read bgp rib v6 learned info: %w", err)
	}
	return populateRIB(netInst, peer4Infos, peer6Infos)
}

func fetchBGPInfos[T ixconfig.IxiaCfgNode](ctx context.Context, client cfgClient, opPath string, bgpPeers []T) (map[string][]bgpLearnedInfo, error) {
	if len(bgpPeers) == 0 {
		return nil, nil
	}
	var nodeIDs []string
	for _, peerNode := range bgpPeers {
		nodeID, err := client.NodeID(peerNode)
		if err != nil {
			return nil, err
		}
		nodeIDs = append(nodeIDs, nodeID)
	}
	if err := client.Session().Post(ctx, opPath, ixweb.OpArgs{nodeIDs}, nil); err != nil {
		return nil, err
	}
	peer2Infos := make(map[string][]bgpLearnedInfo, len(bgpPeers))
	for _, peerNode := range bgpPeers {
		nodeID, err := client.NodeID(peerNode)
		if err != nil {
			return nil, err
		}
		table := &table{}
		if err := client.Session().Get(ctx, path.Join(nodeID, "learnedInfo/1/table/1"), table); err != nil {
			return nil, fmt.Errorf("failed to get learned info: %w", err)
		}
		var infos []bgpLearnedInfo
		unmarshalTable(table, &infos)
		peer2Infos[dutIP(peerNode)] = infos
	}
	return peer2Infos, nil
}

func dutIP(bgpPeer ixconfig.IxiaCfgNode) string {
	switch v := bgpPeer.(type) {
	case *ixconfig.TopologyBgpIpv4Peer:
		return *(v.DutIp.SingleValue.Value)
	case *ixconfig.TopologyBgpIpv6Peer:
		return *(v.DutIp.SingleValue.Value)
	default:
		log.Fatalf("impossible: not a bgp peer type %v (%T)", v, v)
	}
	return ""
}

// bgpLearnedInfo is the output of learned info for BGP from an Ixia device.
// This struct is used for ipv4 and ipv6 info.
type bgpLearnedInfo struct {
	IPV4Prefix     string `ixia:"IPv4 Prefix "`
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

func populateRIB(netInst *oc.NetworkInstance, peer4Infos, peer6Infos map[string][]bgpLearnedInfo) error {
	rib := netInst.
		GetOrCreateProtocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "0").
		GetOrCreateBgp().GetOrCreateRib()

	for peer, infos := range peer4Infos {
		ipv4RIB := rib.GetOrCreateAfiSafi(oc.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST).
			GetOrCreateIpv4Unicast().
			GetOrCreateNeighbor(peer).
			GetOrCreateAdjRibInPre()
		attrLen := len(rib.AttrSet)
		commLen := len(rib.Community)
		extCommLen := len(rib.ExtCommunity)

		for i, info := range infos {
			if info == (bgpLearnedInfo{}) {
				log.Infof("Skipping empty BGP learned info")
				continue
			}
			attrIndex := uint64(i + attrLen)
			commIndex := uint64(i + commLen)
			extCommIndex := uint64(i + extCommLen)
			if err := appendDetails(info, rib, attrIndex, commIndex, extCommIndex, true); err != nil {
				return fmt.Errorf("failed to append details for elem %d: %w", i, err)
			}
			route := &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route{
				Prefix:         ygot.String(fmt.Sprintf("%s/%d", info.IPV4Prefix, info.PrefixLen)),
				PathId:         ygot.Uint32(info.PathID),
				AttrIndex:      &attrIndex,
				CommunityIndex: &commIndex,
			}
			if err := ipv4RIB.AppendRoute(route); err != nil {
				return fmt.Errorf("failed to append route for elem %d: %w", i, err)
			}
		}
	}

	for peer, infos := range peer6Infos {
		ipv6RIB := rib.GetOrCreateAfiSafi(oc.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST).
			GetOrCreateIpv6Unicast().
			GetOrCreateNeighbor(peer).
			GetOrCreateAdjRibInPre()
		attrLen := len(rib.AttrSet)
		commLen := len(rib.Community)
		extCommLen := len(rib.ExtCommunity)

		for i, info := range infos {
			if info == (bgpLearnedInfo{}) {
				log.Infof("Skipping empty BGP learned info")
				continue
			}
			attrIndex := uint64(i + attrLen)
			commIndex := uint64(i + commLen)
			extCommIndex := uint64(i + extCommLen)
			if err := appendDetails(info, rib, attrIndex, commIndex, extCommIndex, false); err != nil {
				return fmt.Errorf("failed to append details for elem %d: %w", i, err)
			}
			route := &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route{
				Prefix:         ygot.String(fmt.Sprintf("%s/%d", info.IPV6Prefix, info.PrefixLen)),
				PathId:         ygot.Uint32(info.PathID),
				AttrIndex:      &attrIndex,
				CommunityIndex: &commIndex,
			}
			if err := ipv6RIB.AppendRoute(route); err != nil {
				return fmt.Errorf("failed to append route for elem %d: %w", i, err)
			}
		}
	}
	return nil
}

func appendDetails(info bgpLearnedInfo, rib *oc.NetworkInstance_Protocol_Bgp_Rib, attrIndex, commIndex, extCommIndex uint64, v4 bool) error {
	nextHop := info.IPV6NextHop
	if v4 {
		nextHop = info.IPV4NextHop
	}

	as := &oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
		NextHop:   &nextHop,
		Aigp:      ygot.Uint64(info.AIGP),
		LocalPref: ygot.Uint32(info.LocalPref),
		Med:       ygot.Uint32(info.MED),
		Index:     &attrIndex,
	}
	switch info.Origin {
	case "IGP":
		as.Origin = oc.RibBgp_BgpOriginAttrType_IGP
	case "EGP":
		as.Origin = oc.RibBgp_BgpOriginAttrType_EGP
	case "Incomplete":
		as.Origin = oc.RibBgp_BgpOriginAttrType_INCOMPLETE
	default:
		return fmt.Errorf("unknown origin type: %q", info.Origin)
	}

	asSegs, err := parseASSegments(info.ASPath)
	if err != nil {
		return err
	}
	for _, seg := range asSegs {
		as.AppendAsSegment(seg)
	}
	if err := rib.AppendAttrSet(as); err != nil {
		return err
	}

	comm := &oc.NetworkInstance_Protocol_Bgp_Rib_Community{
		Index: &commIndex,
	}
	for _, c := range parseCommunities(info.Community) {
		comm.Community = append(comm.Community, c)
	}
	if err := rib.AppendCommunity(comm); err != nil {
		return err
	}

	extComm := &oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity{
		Index: &extCommIndex,
	}
	for _, ec := range parseCommunities(info.Color) {
		extComm.ExtCommunity = append(extComm.ExtCommunity, ec)
	}
	if err := rib.AppendExtCommunity(extComm); err != nil {
		return err
	}

	return nil
}

func parseCommunities(s string) []oc.UnionString {
	// Needed because an empty string will split into a 1-length array containing the empty string.
	if s == "" {
		return nil
	}
	var comms []oc.UnionString
	for _, str := range strings.Split(s, ", ") {
		comms = append(comms, oc.UnionString(str))
	}
	return comms
}

var (
	asSegRE      = regexp.MustCompile("(<|{)[^>}]*(>|})")
	asSegDelimRE = regexp.MustCompile("\\s*(\\s|,)\\s*")
)

func parseASSegments(asPathStr string) ([]*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment, error) {
	// Split asPathStr into segments.
	// Apparently SEQs are space-separated and SETs are comma-separated,
	// but the parsing logic is robust to either of those separators.
	// Example: "<1 2> {3,4,5} <6>" -> ["<1 2>", "{3,4,5}", "<6>"]
	segStrs := asSegRE.FindAllString(asPathStr, -1)

	var segs []*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment
	for i, segStr := range segStrs {
		segType := oc.RibBgp_AsPathSegmentType_AS_SEQ
		if segStr[0] == '{' {
			segType = oc.RibBgp_AsPathSegmentType_AS_SET
		}

		// Chop off the open and closing bracket/brace
		// Example: "{3 4 5}" -> "3 4 5"
		segStr = segStr[1 : len(segStr)-1]
		var members []uint32

		// Split the string at whitespace or commas and convert each element to uint32.
		// Example with whitespace: "1 2" -> ["1", "2"]
		// Example with commas: "3,4,5" -> ["3", "4", "5"]
		for _, s := range asSegDelimRE.Split(segStr, -1) {
			if len(s) == 0 {
				continue
			}
			member, err := strconv.ParseUint(s, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("invalid AS segment member: %q", s)
			}
			members = append(members, uint32(member))
		}
		segs = append(segs, &oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment{
			Index:  ygot.Uint32(uint32(i)),
			Member: members,
			Type:   segType,
		})
	}
	return segs, nil
}
