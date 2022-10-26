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

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/telemetry"
)

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

func isisLSDBDevice(learnedISIS []isisLearnedInfo, intf, protocol string) (*telemetry.Device, error) {
	dev := &telemetry.Device{}
	isis := dev.GetOrCreateNetworkInstance(intf).
		GetOrCreateProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS, protocol).
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
			return nil, fmt.Errorf("unknown ISIS level: %q", info.LearnedVia)
		}
		lspID := fmt.Sprintf("%s%s.%s%s.%s%s.%02d-%02d", info.SystemID[0:2], info.SystemID[3:5],
			info.SystemID[6:8], info.SystemID[9:11], info.SystemID[12:14], info.SystemID[15:17],
			info.PseudoNodeIndex, info.LSPIndex)
		var flags []telemetry.E_Lsp_Flags
		// TODO: infer the flags
		lsp := isis.GetOrCreateLevel(levelNum).GetOrCreateLsp(lspID)
		lsp.SetIsType(levelNum)
		lsp.SetFlags(flags)
		lsp.SetSequenceNumber(info.SeqNumber)
	}

	return dev, nil
}
