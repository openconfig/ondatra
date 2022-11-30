// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flags

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseReserve(t *testing.T) {
	tests := []struct {
		desc        string
		res         string
		wantID      string
		wantPartial map[string]string
		wantErr     string
	}{{
		desc: "nothing",
	}, {
		desc:   "reservation id",
		res:    "1234-5678-abcd-efgh",
		wantID: "1234-5678-abcd-efgh",
	}, {
		desc:        "single assignment",
		res:         "dut=ft00.abc99",
		wantPartial: map[string]string{"dut": "ft00.abc99"},
	}, {
		desc: "multiple assignment",
		res:  "dut=ft00.abc99,dut:port1=Ethernet00",
		wantPartial: map[string]string{
			"dut":       "ft00.abc99",
			"dut:port1": "Ethernet00",
		},
	}, {
		desc: "multiple assignment with spaces",
		res:  "dut = ft00.abc99, dut:port1 = Ethernet00",
		wantPartial: map[string]string{
			"dut":       "ft00.abc99",
			"dut:port1": "Ethernet00",
		},
	}, {
		desc: "ate with no session",
		res:  "ate=blahblah.google.blah,ate:port1=0/0",
		wantPartial: map[string]string{
			"ate":       "blahblah.google.blah",
			"ate:port1": "0/0",
		},
	}, {
		desc: "ate with session",
		res:  "ate=blahblah.google.blah#123,ate:port1=0/0",
		wantPartial: map[string]string{
			"ate":       "blahblah.google.blah#123",
			"ate:port1": "0/0",
		},
	}, {
		desc:    "invalid mapping",
		res:     "dut=ft00=abc99",
		wantErr: "cannot parse",
	}, {
		desc:    "duplicate ID",
		res:     "dut=ft00,dut=ft01",
		wantErr: "duplicate ID",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotID, gotPartial, err := parseReserve(tt.res)
			if (err == nil) != (tt.wantErr == "") || err != nil && !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("parseReserve(): got error %v, want %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(gotID, tt.wantID); diff != "" {
				t.Errorf("parseReserve(): got incorrect ID diff:(-got,+want)\n%v", diff)
			}
			if diff := cmp.Diff(gotPartial, tt.wantPartial); diff != "" {
				t.Errorf("parseReserve(): got incorrect partial mapping diff:(-got,+want)\n%v", diff)
			}
		})
	}
}
