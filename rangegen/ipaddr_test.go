// Copyright 2021 Google LLC
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

package rangegen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseCIDR(t *testing.T) {
	tests := []struct {
		desc        string
		cidr        string
		wantIPv6    bool
		wantMaskLen uint
		wantErr     bool
	}{{
		desc:    "invalid IPv4 CIDR",
		cidr:    "1.2.3.4",
		wantErr: true,
	}, {
		desc:        "valid IPv4 CIDR",
		cidr:        "10.10.10.0/25",
		wantMaskLen: 25,
	}, {
		desc:    "invalid IPv6 CIDR",
		cidr:    "aaaa::",
		wantErr: true,
	}, {
		desc:        "valid IPv6 CIDR",
		cidr:        "aaaa::/17",
		wantMaskLen: 17,
		wantIPv6:    true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			netw, gotErr := parseCIDR(test.cidr)
			if (gotErr != nil) != test.wantErr {
				t.Fatalf("parseCIDR{%s}: unexpected error result, got err: %v, want err? %t",
					test.cidr, gotErr, test.wantErr)
			}
			if test.wantErr {
				return
			}
			if test.wantIPv6 != netw.isIPv6 {
				t.Errorf("parseCIDR{%s}: unexpected IPv6 value, got: %t, want: %t",
					test.cidr, netw.isIPv6, test.wantIPv6)
			}
			if test.wantMaskLen != netw.maskLen {
				t.Errorf("parseCIDR{%s}: unexpected mask length, got: %d, want: %d",
					test.cidr, netw.maskLen, test.wantMaskLen)
			}
		})
	}
}

func TestCIDRRange(t *testing.T) {
	tests := []struct {
		desc     string
		cidr     string
		count    uint32
		wantNets []string
		wantErr  bool
	}{{
		desc:    "IPv4 network range overflow",
		cidr:    "255.0.0.0/8",
		count:   2,
		wantErr: true,
	}, {
		desc:     "IPv4 successful iteration",
		cidr:     "10.10.10.0/25",
		count:    3,
		wantNets: []string{"10.10.10.0/25", "10.10.10.128/25", "10.10.11.0/25"},
	}, {
		desc:     "IPv4 iteration from middle of network",
		cidr:     "10.10.10.1/32",
		count:    3,
		wantNets: []string{"10.10.10.1/32", "10.10.10.2/32", "10.10.10.3/32"},
	}, {
		desc:    "IPv6 network range overflow",
		cidr:    "ffff::/16",
		count:   2,
		wantErr: true,
	}, {
		desc:     "IPv6 successful iteration",
		cidr:     "aaaa::/17",
		count:    3,
		wantNets: []string{"aaaa::/17", "aaaa:8000::/17", "aaab::/17"},
	}, {
		desc:     "IPv6 iteration from middle of network",
		cidr:     "aaaa:0001::/17",
		count:    3,
		wantNets: []string{"aaaa:1::/17", "aaaa:8001::/17", "aaab:1::/17"},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			netsCh, gotErr := CIDRs(test.cidr, test.count)
			if (gotErr != nil) != test.wantErr {
				t.Fatalf("CIDRs{%s, %d}: unexpected error result, got err: %v, want err? %t",
					test.cidr, test.count, gotErr, test.wantErr)
			}
			if test.wantErr {
				return
			}

			gotNets := make([]string, 0)
			for net := range netsCh {
				gotNets = append(gotNets, net)
			}
			if diff := cmp.Diff(test.wantNets, gotNets); diff != "" {
				t.Fatalf("CIDRs{%s, %d}: unexpected diff in generated list of networks: %s",
					test.cidr, test.count, diff)
			}
		})
	}
}
