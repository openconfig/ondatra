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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/telemetry"
)

func TestISISLSDBDevice(t *testing.T) {
	tests := []struct {
		desc     string
		info     []isisLearnedInfo
		wantISIS *telemetry.NetworkInstance_Protocol_Isis
		wantErr  string
	}{{
		desc:     "empty learned info",
		info:     []isisLearnedInfo{{}},
		wantISIS: &telemetry.NetworkInstance_Protocol_Isis{},
	}, {
		desc: "success",
		info: []isisLearnedInfo{{
			LearnedVia:      "L1",
			SystemID:        "01 23 45 67 89 AB",
			PseudoNodeIndex: 2,
			LSPIndex:        34,
			SeqNumber:       5,
		}, {
			LearnedVia:      "L2",
			SystemID:        "CD EF 01 23 45 67",
			PseudoNodeIndex: 98,
			LSPIndex:        7,
			SeqNumber:       0,
		}},
		wantISIS: &telemetry.NetworkInstance_Protocol_Isis{
			Level: map[uint8]*telemetry.NetworkInstance_Protocol_Isis_Level{
				1: &telemetry.NetworkInstance_Protocol_Isis_Level{
					LevelNumber: ygot.Uint8(1),
					Lsp: map[string]*telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
						"0123.4567.89AB.02-34": &telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
							IsType:         ygot.Uint8(1),
							LspId:          ygot.String("0123.4567.89AB.02-34"),
							SequenceNumber: ygot.Uint32(5),
						},
					},
				},
				2: &telemetry.NetworkInstance_Protocol_Isis_Level{
					LevelNumber: ygot.Uint8(2),
					Lsp: map[string]*telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
						"CDEF.0123.4567.98-07": &telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
							IsType:         ygot.Uint8(2),
							LspId:          ygot.String("CDEF.0123.4567.98-07"),
							SequenceNumber: ygot.Uint32(0),
						},
					},
				},
			},
		},
	}, {
		desc:    "unknown level",
		info:    []isisLearnedInfo{{LearnedFrom: "L3"}},
		wantErr: "level",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := isisLSDBDevice(test.info, "interface", "protocol")
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("isisLSDBDevice got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			gotISIS := got.GetNetworkInstance("interface").
				GetProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS, "protocol").GetIsis()
			if d := cmp.Diff(test.wantISIS, gotISIS); d != "" {
				t.Errorf("isisLSDBDevice got unexpected ISIS diff (-want +got)\n%s", d)
			}
		})
	}
}
