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

	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ondatra/telemetry"
)

var hashImpl = fnv.New32()

type ingressLSP struct {
	up       bool
	src, dst string
}

func rsvpTEDevice(ingressLSPsBySessionPrefix map[string][]*ingressLSP, intf string) (*telemetry.Device, error) {
	dev := &telemetry.Device{}
	rsvpTE := dev.GetOrCreateNetworkInstance(intf).GetOrCreateMpls().GetOrCreateSignalingProtocols().GetOrCreateRsvpTe()

	// Hash prefixes to maintain constant session indices for a given config.
	pfxToHash := make(map[string]uint32, len(ingressLSPsBySessionPrefix))
	hashToPfx := make(map[uint32]string, len(ingressLSPsBySessionPrefix))
	for pfx := range ingressLSPsBySessionPrefix {
		_, err := hashImpl.Write([]byte(pfx))
		if err != nil {
			return nil, fmt.Errorf("failed to hash session prefix %q: %w", pfx, err)
		}
		h := hashImpl.Sum32()
		if hPfx, exists := hashToPfx[h]; exists {
			return nil, fmt.Errorf("hash collision for session prefixes %q, %q", pfx, hPfx)
		}
		hashToPfx[h] = pfx
		pfxToHash[pfx] = h
		hashImpl.Reset()
	}

	rsvpTE.Session = map[uint64]*telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{}
	for pfx, lsps := range ingressLSPsBySessionPrefix {
		// Left-shift hash of the LSP config to the upper bits of a uint64.
		// The lower bits will be used to indicate the indices of tunnels
		// within this config.
		h := uint64(pfxToHash[pfx]) << 32
		for i, lsp := range lsps {
			//
			session := rsvpTE.GetOrCreateSession(h | uint64(i))
			session.SessionName = ygot.String(fmt.Sprintf("%s %d", pfx, i))
			session.Type = telemetry.MplsTypes_LSP_ROLE_INGRESS
			session.SourceAddress = ygot.String(lsp.src)
			session.DestinationAddress = ygot.String(lsp.dst)
			session.Status = telemetry.Session_Status_DOWN
			if lsp.up {
				session.Status = telemetry.Session_Status_UP
			}
		}
	}
	return dev, nil
}
