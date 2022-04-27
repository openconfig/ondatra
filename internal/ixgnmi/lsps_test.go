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
	"hash"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/telemetry"
)

type constantHash struct {
	hash.Hash32
}

func (constantHash) Write(p []byte) (int, error) { return len(p), nil }
func (constantHash) Reset()                      {}
func (constantHash) Sum32() uint32               { return 0 }

func TestLSPs(t *testing.T) {
	origHashImpl := hashImpl
	tests := []struct {
		name        string
		ingressLSPs map[string][]*ingressLSP
		hashImpl    hash.Hash32
		inLSPs      *telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe
		wantLSPs    *telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe
		wantErr     string
	}{{
		name: "hash collision on LSP session prefixes",
		ingressLSPs: map[string][]*ingressLSP{
			"lsp0": []*ingressLSP{},
			"lsp1": []*ingressLSP{},
		},
		hashImpl: constantHash{},
		wantErr:  "hash collision",
	}, {
		name: "full LSP update",
		ingressLSPs: map[string][]*ingressLSP{
			"lsp0": []*ingressLSP{
				{up: true, src: "1.1.1.1", dst: "1.1.1.2"},
				{up: false, src: "2.2.2.1", dst: "2.2.2.2"},
			},
			"lsp1": []*ingressLSP{
				{up: true, src: "3.3.3.3", dst: "4.4.4.4"},
			},
		},
		// Non-empty initial data to test that it is reset.
		inLSPs: &telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe{
			Session: map[uint64]*telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
				0: &telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
					LocalIndex:         ygot.Uint64(0),
					SessionName:        ygot.String("some session"),
					Type:               telemetry.MplsTypes_LSP_ROLE_INGRESS,
					SourceAddress:      ygot.String("10.10.10.10"),
					DestinationAddress: ygot.String("11.11.11.11"),
					Status:             telemetry.Session_Status_DOWN,
				},
			},
		},
		wantLSPs: &telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe{
			Session: map[uint64]*telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
				2765635037960339456: &telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
					LocalIndex:         ygot.Uint64(2765635037960339456),
					SessionName:        ygot.String("lsp0 0"),
					Type:               telemetry.MplsTypes_LSP_ROLE_INGRESS,
					SourceAddress:      ygot.String("1.1.1.1"),
					DestinationAddress: ygot.String("1.1.1.2"),
					Status:             telemetry.Session_Status_UP,
				},
				2765635037960339457: &telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
					LocalIndex:         ygot.Uint64(2765635037960339457),
					SessionName:        ygot.String("lsp0 1"),
					Type:               telemetry.MplsTypes_LSP_ROLE_INGRESS,
					SourceAddress:      ygot.String("2.2.2.1"),
					DestinationAddress: ygot.String("2.2.2.2"),
					Status:             telemetry.Session_Status_DOWN,
				},
				2765635042255306752: &telemetry.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
					LocalIndex:         ygot.Uint64(2765635042255306752),
					SessionName:        ygot.String("lsp1 0"),
					Type:               telemetry.MplsTypes_LSP_ROLE_INGRESS,
					SourceAddress:      ygot.String("3.3.3.3"),
					DestinationAddress: ygot.String("4.4.4.4"),
					Status:             telemetry.Session_Status_UP,
				},
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hashImpl != nil {
				hashImpl = tt.hashImpl
			}
			defer func() { hashImpl = origHashImpl }()

			err := lsps(tt.ingressLSPs, tt.inLSPs)
			if d := errdiff.Substring(err, tt.wantErr); d != "" {
				t.Fatalf("unexpected error string\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(tt.wantLSPs, tt.inLSPs); d != "" {
				t.Errorf("unexpected LSPs diff (-want +got)\n%s", d)
			}
		})
	}
}
