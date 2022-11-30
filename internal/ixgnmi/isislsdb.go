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

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/ixconfig"
)

func isisLSDBFromIxia(ctx context.Context, client cfgClient, netInst *oc.NetworkInstance, nodes *cachedNodes) error {
	info, err := fetchISISInfo(ctx, client, nodes.isisL3s)
	if err != nil {
		return err
	}
	return populateISIS(netInst, info)
}

type isisLearnedInfo struct {
	LearnedVia      string `ixia:"Learned Via"`
	LearnedFrom     string `ixia:"Learned From"`
	Metric          int    `ixia:"Metric"`
	IPv4Prefix      string `ixia:"IPv4 Prefix"`
	Mask            int    `ixia:"Mask"`
	PrefixAttrFlags string `ixia:"Prefix Attribute Flags"`
	IPv4SrcRouterID string `ixia:"IPv4 Source Router ID"`
	IPv6SrcRouterID string `ixia:"IPv6 Source Router ID"`
	SIDLabel        int    `ixia:"SID/Label"`
	Algorithm       int    `ixia:"Algorithm"`
	FAPMMetric      int    `ixia:"FAPM Metric"`
	SystemID        string `ixia:"System ID"`
	PseudoNodeIndex int    `ixia:"Pseudo Node Index"`
	LSPIndex        int    `ixia:"LSP Index"`
	HostName        string `ixia:"Host Name"`
	SeqNumber       uint32 `ixia:"Sequence Number"`
	AgeSecs         int    `ixia:"Age (in secs)"`
}

func fetchISISInfo(ctx context.Context, client cfgClient, isisL3s []*ixconfig.TopologyIsisL3) ([]isisLearnedInfo, error) {
	const opPath = "topology/deviceGroup/ethernet/isisL3/operations/getLearnedInfo"
	var nodeIDs []string
	for _, isisNode := range isisL3s {
		nodeID, err := client.NodeID(isisNode)
		if err != nil {
			return nil, err
		}
		nodeIDs = append(nodeIDs, nodeID)
	}
	if err := client.Session().Post(ctx, opPath, ixweb.OpArgs{nodeIDs}, nil); err != nil {
		return nil, fmt.Errorf("failed to run op: %w", err)
	}
	var allInfos []isisLearnedInfo
	for _, nodeID := range nodeIDs {
		table := &table{}
		if err := client.Session().Get(ctx, path.Join(nodeID, "learnedInfo/1/table/1"), table); err != nil {
			return nil, fmt.Errorf("failed to get learned info: %w", err)
		}
		var infos []isisLearnedInfo
		unmarshalTable(table, &infos)
		allInfos = append(allInfos, infos...)
	}
	return allInfos, nil
}

func populateISIS(netInst *oc.NetworkInstance, learnedISIS []isisLearnedInfo) error {
	isis := netInst.
		GetOrCreateProtocol(oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS, "0").
		GetOrCreateIsis()

	for _, info := range learnedISIS {
		if info == (isisLearnedInfo{}) {
			log.Infof("Skipping empty ISIS learned info")
			continue
		}

		var levelNum uint8
		switch info.LearnedVia {
		case "L1":
			levelNum = 1
		case "L2":
			levelNum = 2
		default:
			return fmt.Errorf("unknown ISIS level: %q", info.LearnedVia)
		}
		lspID := fmt.Sprintf("%s%s.%s%s.%s%s.%02d-%02d", info.SystemID[0:2], info.SystemID[3:5],
			info.SystemID[6:8], info.SystemID[9:11], info.SystemID[12:14], info.SystemID[15:17],
			info.PseudoNodeIndex, info.LSPIndex)
		var flags []oc.E_Lsp_Flags
		// TODO(greg-dennis): Infer the flags when we learn what they look like in learned info.
		lsp := isis.GetOrCreateLevel(levelNum).GetOrCreateLsp(lspID)
		lsp.SetIsType(levelNum)
		lsp.SetFlags(flags)
		lsp.SetSequenceNumber(info.SeqNumber)
		// TODO(team): Populate IPv4 prefixes when IxNetwork provides support in learned info.
		extV4Prefix := lsp.
			GetOrCreateTlv(oc.IsisLsdbTypes_ISIS_TLV_TYPE_EXTENDED_IPV4_REACHABILITY).
			GetOrCreateExtendedIpv4Reachability().
			GetOrCreatePrefix(info.IPv4Prefix)
		extV4Prefix.SetMetric(uint32(info.Metric))
	}
	return nil
}
