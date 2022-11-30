// Copyright 2022 Google Inc.
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
	"hash/fnv"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ygot/ygot"
)

func rsvpTEFromIxia(ctx context.Context, client cfgClient, netInst *oc.NetworkInstance, nodes *cachedNodes) error {
	ingressLSPs, err := fetchIngressLSPs(ctx, client, nodes.rsvpLSPs)
	if err != nil {
		return err
	}
	return populateIngressLSPs(netInst, ingressLSPs)
}

var hashImpl = fnv.New32()

type ingressLSP struct {
	up       bool
	src, dst string
}

func fetchIngressLSPs(ctx context.Context, client cfgClient, rsvpLSPs []*ixconfig.TopologyRsvpP2PIngressLsps) (map[string][]*ingressLSP, error) {
	const mvEP = "multivalue/operations/getvalues"
	if len(rsvpLSPs) == 0 {
		return nil, nil
	}

	ingressLSPsByPrefix := make(map[string][]*ingressLSP)
	for _, lspNode := range rsvpLSPs {
		nodeID, err := client.NodeID(lspNode)
		if err != nil {
			return nil, err
		}
		ixLSPs := new(struct {
			State    []string
			SourceIP string
			RemoteIP string
		})
		if err := client.Session().Get(ctx, nodeID, ixLSPs); err != nil {
			return nil, fmt.Errorf("failed to fetch ingress LSPs config: %w", err)
		}
		var srcIPs, dstIPs []string
		if err := client.Session().Post(ctx, mvEP, ixweb.OpArgs{ixLSPs.SourceIP, 0, len(ixLSPs.State)}, &srcIPs); err != nil {
			return nil, fmt.Errorf("failed to fetch source IPs for LSP: %w", err)
		}
		if err := client.Session().Post(ctx, mvEP, ixweb.OpArgs{ixLSPs.RemoteIP, 0, len(ixLSPs.State)}, &dstIPs); err != nil {
			return nil, fmt.Errorf("failed to fetch destination IPs for LSP: %w", err)
		}
		var lsps []*ingressLSP
		for i, state := range ixLSPs.State {
			lsps = append(lsps, &ingressLSP{
				up:  state == "up",
				src: srcIPs[i],
				dst: dstIPs[i],
			})
		}
		ingressLSPsByPrefix[*(lspNode.Name)] = lsps
	}

	return ingressLSPsByPrefix, nil
}

func populateIngressLSPs(netInst *oc.NetworkInstance, ingressLSPsBySessionPrefix map[string][]*ingressLSP) error {
	rsvpTE := netInst.GetOrCreateMpls().GetOrCreateSignalingProtocols().GetOrCreateRsvpTe()

	// Hash prefixes to maintain constant session indices for a given config.
	pfxToHash := make(map[string]uint32, len(ingressLSPsBySessionPrefix))
	hashToPfx := make(map[uint32]string, len(ingressLSPsBySessionPrefix))
	for pfx := range ingressLSPsBySessionPrefix {
		_, err := hashImpl.Write([]byte(pfx))
		if err != nil {
			return fmt.Errorf("failed to hash session prefix %q: %w", pfx, err)
		}
		h := hashImpl.Sum32()
		if hPfx, exists := hashToPfx[h]; exists {
			return fmt.Errorf("hash collision for session prefixes %q, %q", pfx, hPfx)
		}
		hashToPfx[h] = pfx
		pfxToHash[pfx] = h
		hashImpl.Reset()
	}

	rsvpTE.Session = map[uint64]*oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{}
	for pfx, lsps := range ingressLSPsBySessionPrefix {
		// Left-shift hash of the LSP config to the upper bits of a uint64.
		// The lower bits will be used to indicate the indices of tunnels
		// within this config.
		h := uint64(pfxToHash[pfx]) << 32
		for i, lsp := range lsps {
			//
			session := rsvpTE.GetOrCreateSession(h | uint64(i))
			session.SessionName = ygot.String(fmt.Sprintf("%s %d", pfx, i))
			session.Type = oc.MplsTypes_LSP_ROLE_INGRESS
			session.SourceAddress = ygot.String(lsp.src)
			session.DestinationAddress = ygot.String(lsp.dst)
			session.Status = oc.Session_Status_DOWN
			if lsp.up {
				session.Status = oc.Session_Status_UP
			}
		}
	}
	return nil
}
