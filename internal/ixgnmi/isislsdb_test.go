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
	"golang.org/x/net/context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ondatra/telemetry"
)

func TestISISLSDBFromIxia(t *testing.T) {
	const isisID = "/fake/isis/id"
	isisXP := parseXPath(t, "/xpath/to/isis")
	nodes := &cachedNodes{
		isisL3s: []*ixconfig.TopologyIsisL3{{Xpath: isisXP}},
	}

	tests := []struct {
		desc    string
		postErr map[string]error
		getErr  map[string]error
		getRsps map[string]string
		want    *telemetry.NetworkInstance
		wantErr string
	}{{
		desc: "run op error",
		postErr: map[string]error{
			"topology/deviceGroup/ethernet/isisL3/operations/getLearnedInfo": errors.New("op fail"),
		},
		wantErr: "op fail",
	}, {
		desc: "get error",
		getErr: map[string]error{
			isisID + "/learnedInfo/1/table/1": errors.New("get fail"),
		},
		wantErr: "get fail",
	}, {
		desc: "unknown level",
		getRsps: map[string]string{
			isisID + "/learnedInfo/1/table/1": `{
				"columns": ["Learned Via"],
				"values": [["L3"]]
			}`,
		},
		wantErr: "level",
	}, {
		desc: "no data",
		want: &telemetry.NetworkInstance{
			Protocol: map[telemetry.NetworkInstance_Protocol_Key]*telemetry.NetworkInstance_Protocol{
				telemetry.NetworkInstance_Protocol_Key{
					Identifier: telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS,
					Name:       "0",
				}: {
					Identifier: telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS,
					Name:       ygot.String("0"),
					Isis:       &telemetry.NetworkInstance_Protocol_Isis{},
				},
			},
		},
	}, {
		desc: "full data",
		getRsps: map[string]string{
			isisID + "/learnedInfo/1/table/1": `{
				"columns": ["Learned Via", "System ID", "Pseudo Node Index", "LSP Index", "Sequence Number"],
				"values": [
					["L1", "01 23 45 67 89 AB",  "2", "34", "5"],
					["L2", "CD EF 01 23 45 67", "98",  "7", "0"]
				]
			}`,
		},
		want: &telemetry.NetworkInstance{
			Protocol: map[telemetry.NetworkInstance_Protocol_Key]*telemetry.NetworkInstance_Protocol{
				telemetry.NetworkInstance_Protocol_Key{
					Identifier: telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS,
					Name:       "0",
				}: {
					Identifier: telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS,
					Name:       ygot.String("0"),
					Isis: &telemetry.NetworkInstance_Protocol_Isis{
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
				},
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			getRsps := make(map[string][]string)
			for p, r := range test.getRsps {
				getRsps[p] = []string{r}
			}
			client := &fakeCfgClient{
				sess: &fakeSession{
					postErrs: test.postErr,
					getErrs:  test.getErr,
					getRsps:  getRsps,
				},
				xpathToID: map[string]string{
					isisXP.String(): isisID,
				},
			}

			got := new(telemetry.NetworkInstance)
			err := isisLSDBFromIxia(context.Background(), client, got, nodes)
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("isisLSDBFromIxia() got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(test.want, got); d != "" {
				t.Errorf("isisLSDBFromIxia() got unexpected diff (-want +got)\n%s", d)
			}
		})
	}
}
