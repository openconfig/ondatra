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
	"errors"
	"hash"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ygot/ygot"
)

type constantHash struct {
	hash.Hash32
}

func (constantHash) Write(p []byte) (int, error) { return len(p), nil }
func (constantHash) Reset()                      {}
func (constantHash) Sum32() uint32               { return 0 }

func TestRSVPTEFromIxia(t *testing.T) {
	const (
		lsp1ID = "/fake/rsvpte/lsp1"
		lsp2ID = "/fake/rsvpte/lsp2"
	)
	lsp1XP := parseXPath(t, "/fake/xpath/lsp1")
	lsp2XP := parseXPath(t, "/fake/xpath/lsp2")
	nodes := &cachedNodes{rsvpLSPs: []*ixconfig.TopologyRsvpP2PIngressLsps{
		{Xpath: lsp1XP, Name: ixconfig.String("lsp0")},
		{Xpath: lsp2XP, Name: ixconfig.String("lsp1")},
	}}

	tests := []struct {
		desc     string
		lspRsps  map[string]string
		lspErrs  map[string]error
		mvRsps   []string
		mvErr    error
		hashImpl hash.Hash32
		want     *oc.NetworkInstance
		wantErr  string
	}{{
		desc:    "lsp lookup error",
		lspErrs: map[string]error{lsp1ID: errors.New("some error")},
		wantErr: "failed to fetch ingress LSPs config",
	}, {
		desc:    "multivalue lookup error",
		lspRsps: map[string]string{lsp1ID: `{"state": ["up"], "sourceIP": "/api/v1/sessions/0/multivalue/1", "destIP": "/api/v1/sessions/0/multivalue/2"}`},
		mvErr:   errors.New("some error"),
		wantErr: "failed to fetch source IPs for LSP",
	}, {
		desc:     "hash collision",
		hashImpl: new(constantHash),
		wantErr:  "hash collision",
	}, {
		desc: "no data",
		want: &oc.NetworkInstance{Mpls: &oc.NetworkInstance_Mpls{
			SignalingProtocols: &oc.NetworkInstance_Mpls_SignalingProtocols{
				RsvpTe: &oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe{
					Session: map[uint64]*oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{},
				},
			},
		}},
	}, {
		desc: "full data",
		lspRsps: map[string]string{
			lsp1ID: `{"state": ["up", "notStarted"], "sourceIP": "/api/v1/sessions/0/multivalue/1", "destIP": "/api/v1/sessions/0/multivalue/2"}`,
			lsp2ID: `{"state": ["up"], "sourceIP": "/api/v1/sessions/0/multivalue/3", "destIP": "/api/v1/sessions/0/multivalue/4"}`,
		},
		mvRsps: []string{
			`["1.1.1.1", "2.2.2.1"]`,
			`["1.1.1.2", "2.2.2.2"]`,
			`["3.3.3.1"]`,
			`["3.3.3.2"]`,
		},
		want: &oc.NetworkInstance{Mpls: &oc.NetworkInstance_Mpls{
			SignalingProtocols: &oc.NetworkInstance_Mpls_SignalingProtocols{
				RsvpTe: &oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe{
					Session: map[uint64]*oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
						2765635037960339456: &oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
							LocalIndex:         ygot.Uint64(2765635037960339456),
							SessionName:        ygot.String("lsp0 0"),
							Type:               oc.MplsTypes_LSP_ROLE_INGRESS,
							SourceAddress:      ygot.String("1.1.1.1"),
							DestinationAddress: ygot.String("1.1.1.2"),
							Status:             oc.Session_Status_UP,
						},
						2765635037960339457: &oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
							LocalIndex:         ygot.Uint64(2765635037960339457),
							SessionName:        ygot.String("lsp0 1"),
							Type:               oc.MplsTypes_LSP_ROLE_INGRESS,
							SourceAddress:      ygot.String("2.2.2.1"),
							DestinationAddress: ygot.String("2.2.2.2"),
							Status:             oc.Session_Status_DOWN,
						},
						2765635042255306752: &oc.NetworkInstance_Mpls_SignalingProtocols_RsvpTe_Session{
							LocalIndex:         ygot.Uint64(2765635042255306752),
							SessionName:        ygot.String("lsp1 0"),
							Type:               oc.MplsTypes_LSP_ROLE_INGRESS,
							SourceAddress:      ygot.String("3.3.3.1"),
							DestinationAddress: ygot.String("3.3.3.2"),
							Status:             oc.Session_Status_UP,
						},
					},
				},
			},
		}},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if test.hashImpl != nil {
				origHashImpl := hashImpl
				defer func() {
					hashImpl = origHashImpl
				}()
				hashImpl = test.hashImpl
			}

			getRsps := make(map[string][]string)
			for id, rsp := range test.lspRsps {
				getRsps[id] = []string{rsp}
			}
			client := &fakeCfgClient{
				sess: &fakeSession{
					getErrs:  test.lspErrs,
					getRsps:  getRsps,
					postErrs: map[string]error{"multivalue/operations/getvalues": test.mvErr},
					postRsps: map[string][]string{"multivalue/operations/getvalues": test.mvRsps},
				},
				xpathToID: map[string]string{
					lsp1XP.String(): lsp1ID,
					lsp2XP.String(): lsp2ID,
				},
			}

			got := new(oc.NetworkInstance)
			err := rsvpTEFromIxia(context.Background(), client, got, nodes)
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("rsvpTEFromIxia() got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(test.want, got); d != "" {
				t.Errorf("rsvpTEFromIxia() got unexpected diff (-want +got)\n%s", d)
			}
		})
	}
}
