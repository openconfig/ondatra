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

package ate

import (
	"strings"
	"testing"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

const fakeSeed = 123

func init() {
	// Stub the random int function.
	randUInt32 = func() uint32 {
		return fakeSeed
	}
}

func TestSetUintRangeField(t *testing.T) {
	tests := []struct {
		desc string
		ints *opb.UIntRange
		want *ixconfig.TrafficField
	}{{
		desc: "single value",
		ints: &opb.UIntRange{Min: 4, Max: 4, Step: 1, Count: 1},
		want: &ixconfig.TrafficField{
			Auto:        ixconfig.Bool(false),
			ValueType:   ixconfig.String("singleValue"),
			SingleValue: ixconfig.String("4"),
		},
	}, {
		desc: "nonrandom range",
		ints: &opb.UIntRange{Min: 4, Max: 100, Step: 5, Count: 20},
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("increment"),
			CountValue: ixconfig.String("20"),
			StartValue: ixconfig.String("4"),
			StepValue:  ixconfig.String("5"),
		},
	}, {
		desc: "nonrandom range, default step",
		ints: &opb.UIntRange{Min: 1, Max: 10, Count: 5},
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("increment"),
			CountValue: ixconfig.String("5"),
			StartValue: ixconfig.String("1"),
			StepValue:  ixconfig.String("2"),
		},
	}, {
		desc: "random range",
		ints: &opb.UIntRange{Min: 4, Max: 100, Step: 5, Count: 20, Random: true},
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("repeatableRandomRange"),
			CountValue: ixconfig.String("20"),
			MinValue:   ixconfig.String("4"),
			MaxValue:   ixconfig.String("100"),
			StepValue:  ixconfig.String("5"),
			Seed:       uintToStr(fakeSeed),
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			field := &ixconfig.TrafficField{}
			if err := setUintRangeField(field, test.ints); err != nil {
				t.Fatalf("setUintRangeField(<field>, %v): unexpected error: %v", test.ints, err)
			}
			if diff := jsonCfgDiff(t, test.want, field); diff != "" {
				t.Errorf("setUintRangeField(<field>, %v): unexpected field diff (-want, +got)\n%s", test.ints, diff)
			}
		})
	}
}

func TestSetUintRangeFieldErrors(t *testing.T) {
	tests := []struct {
		desc    string
		ints    *opb.UIntRange
		wantErr string
	}{{
		desc:    "count zero",
		ints:    &opb.UIntRange{Min: 1, Max: 9, Step: 1, Count: 0},
		wantErr: "zero",
	}, {
		desc:    "min greater than max",
		ints:    &opb.UIntRange{Min: 9, Max: 1, Step: 1, Count: 1},
		wantErr: "greater than max",
	}, {
		desc:    "count cannot fit",
		ints:    &opb.UIntRange{Min: 1, Max: 9, Step: 3, Count: 4},
		wantErr: "cannot fit",
	}, {
		desc:    "count cannot fit, default step",
		ints:    &opb.UIntRange{Min: 1, Max: 1, Step: 0, Count: 2},
		wantErr: "cannot fit",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if err := setUintRangeField(&ixconfig.TrafficField{}, test.ints); err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("setUintRangeField(<field>, %v): got error %v, want error %q", test.ints, err, test.wantErr)
			}
		})
	}
}

func TestSetAddrRangeField(t *testing.T) {
	tests := []struct {
		desc  string
		addrs *opb.AddressRange
		at    addrType
		want  *ixconfig.TrafficField
	}{{
		desc:  "single MAC-48",
		addrs: &opb.AddressRange{Min: "01:02:03:04:05:06", Max: "01:02:03:04:05:06", Step: "00:00:00:00:00:01", Count: 1},
		at:    mac48AddrType,
		want: &ixconfig.TrafficField{
			Auto:        ixconfig.Bool(false),
			ValueType:   ixconfig.String("singleValue"),
			SingleValue: ixconfig.String("01:02:03:04:05:06"),
		},
	}, {
		desc:  "single IPv4",
		addrs: &opb.AddressRange{Min: "1.2.3.4", Max: "1.2.3.4", Step: "0.0.0.1", Count: 1},
		at:    ipv4AddrType,
		want: &ixconfig.TrafficField{
			Auto:        ixconfig.Bool(false),
			ValueType:   ixconfig.String("singleValue"),
			SingleValue: ixconfig.String("1.2.3.4"),
		},
	}, {
		desc:  "single IPv6",
		addrs: &opb.AddressRange{Min: "1:2:3:4:5:6:7:8", Max: "1:2:3:4:5:6:7:8", Step: "::1", Count: 1},
		at:    ipv6AddrType,
		want: &ixconfig.TrafficField{
			Auto:        ixconfig.Bool(false),
			ValueType:   ixconfig.String("singleValue"),
			SingleValue: ixconfig.String("1:2:3:4:5:6:7:8"),
		},
	}, {
		desc:  "nonrandom MAC-48",
		addrs: &opb.AddressRange{Min: "00:00:00:00:00:01", Max: "00:00:00:00:01:00", Count: 20},
		at:    mac48AddrType,
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("increment"),
			CountValue: ixconfig.String("20"),
			StartValue: ixconfig.String("00:00:00:00:00:01"),
			StepValue:  ixconfig.String("00:00:00:00:00:0c"),
		},
	}, {
		desc:  "nonrandom IPv4",
		addrs: &opb.AddressRange{Min: "0.0.0.1", Max: "0.0.1.0", Count: 20},
		at:    ipv4AddrType,
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("increment"),
			CountValue: ixconfig.String("20"),
			StartValue: ixconfig.String("0.0.0.1"),
			StepValue:  ixconfig.String("0.0.0.12"),
		},
	}, {
		desc:  "nonrandom IPv6",
		addrs: &opb.AddressRange{Min: "0:0:0:0:0:0:0:1", Max: "0:0:0:0:0:0:1:0", Count: 20},
		at:    ipv6AddrType,
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("increment"),
			CountValue: ixconfig.String("20"),
			StartValue: ixconfig.String("0:0:0:0:0:0:0:1"),
			StepValue:  ixconfig.String("::ccc"),
		},
	}, {
		desc:  "random MAC-48",
		addrs: &opb.AddressRange{Min: "01:02:03:04:05:06", Max: "02:03:04:05:06:07", Count: 20, Random: true},
		at:    mac48AddrType,
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("repeatableRandomRange"),
			CountValue: ixconfig.String("20"),
			MinValue:   ixconfig.String("01:02:03:04:05:06"),
			MaxValue:   ixconfig.String("02:03:04:05:06:07"),
			StepValue:  ixconfig.String("00:00:00:00:00:01"),
			Seed:       uintToStr(fakeSeed),
		},
	}, {
		desc:  "random IPv4",
		addrs: &opb.AddressRange{Min: "1.2.3.4", Max: "2.3.4.5", Count: 20, Random: true},
		at:    ipv4AddrType,
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("repeatableRandomRange"),
			CountValue: ixconfig.String("20"),
			MinValue:   ixconfig.String("1.2.3.4"),
			MaxValue:   ixconfig.String("2.3.4.5"),
			StepValue:  ixconfig.String("0.0.0.1"),
			Seed:       uintToStr(fakeSeed),
		},
	}, {
		desc:  "random IPv6",
		addrs: &opb.AddressRange{Min: "1:2:3:4:5:6:7:8", Max: "2:3:4:5:6:7:8:9", Count: 20, Random: true},
		at:    ipv6AddrType,
		want: &ixconfig.TrafficField{
			Auto:       ixconfig.Bool(false),
			FullMesh:   ixconfig.Bool(false),
			ValueType:  ixconfig.String("repeatableRandomRange"),
			CountValue: ixconfig.String("20"),
			MinValue:   ixconfig.String("1:2:3:4:5:6:7:8"),
			MaxValue:   ixconfig.String("2:3:4:5:6:7:8:9"),
			StepValue:  ixconfig.String("::1"),
			Seed:       uintToStr(fakeSeed),
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			field := &ixconfig.TrafficField{}
			if err := setAddrRangeField(field, test.at, test.addrs); err != nil {
				t.Fatalf("setAddrRangeField(<field>, %q, %v): unexpected error: %v", test.at, test.addrs, err)
			}
			if diff := jsonCfgDiff(t, test.want, field); diff != "" {
				t.Errorf("setAddrRangeField(<field>, %q, %v): unexpected field diff (-want, +got)\n%s", test.at, test.addrs, diff)
			}
		})
	}
}

func TestAddrRangeToTrafficValueErrors(t *testing.T) {
	tests := []struct {
		desc    string
		addrs   *opb.AddressRange
		at      addrType
		wantErr string
	}{{
		desc:    "count zero",
		addrs:   &opb.AddressRange{Min: "1.2.3.4", Max: "2.3.4.5", Step: "0.0.0.1", Count: 0},
		wantErr: "zero",
	}, {
		desc:    "step zero MAC-48",
		addrs:   &opb.AddressRange{Min: "01:02:03:04:05:06", Max: "02:03:04:05:06:07", Step: "00:00:00:00:00:00", Count: 1},
		at:      mac48AddrType,
		wantErr: "zero",
	}, {
		desc:    "step zero IPv4",
		addrs:   &opb.AddressRange{Min: "1.2.3.4", Max: "2.3.4.5", Step: "0.0.0.0", Count: 1},
		at:      ipv4AddrType,
		wantErr: "zero",
	}, {
		desc:    "step zero IPv6",
		addrs:   &opb.AddressRange{Min: "1:2:3:4:5:6:7:8", Max: "2:3:4:5:6:7:8:9", Step: "::", Count: 1},
		at:      ipv6AddrType,
		wantErr: "zero",
	}, {
		desc:    "invalid MAC-48",
		addrs:   &opb.AddressRange{Min: "01:02:03:04:05:", Max: "02:03:04:05:06:07", Step: "00:00:00:00:00:01", Count: 1},
		at:      mac48AddrType,
		wantErr: "not a MAC address",
	}, {
		desc:    "invalid IPv4",
		addrs:   &opb.AddressRange{Min: "1.2.3.4", Max: "2.3.4", Step: "0.0.0.1", Count: 1},
		at:      ipv4AddrType,
		wantErr: "not an IPv4 address",
	}, {
		desc:    "invalid IPv6",
		addrs:   &opb.AddressRange{Min: "1:2:3:4:5:6:7:8", Max: "2:3:4:5:6:7:8:9", Step: "::z", Count: 1},
		at:      ipv6AddrType,
		wantErr: "not an IPv6 address",
	}, {
		desc:    "not in MAC-48 format",
		addrs:   &opb.AddressRange{Min: "01:02:03:04:05:06:07:08", Max: "02:03:04:05:06:07", Step: "00:00:00:00:00:01", Count: 1},
		at:      mac48AddrType,
		wantErr: "not in MAC-48 format",
	}, {
		desc:    "not in IPv4 format",
		addrs:   &opb.AddressRange{Min: "::1", Max: "2.3.4.5", Step: "0.0.0.1", Count: 1},
		at:      ipv4AddrType,
		wantErr: "not an IPv4 address",
	}, {
		desc:    "not in IPv6 format",
		addrs:   &opb.AddressRange{Min: "1.2.3.4", Max: "2:3:4:5:6:7:8:9", Step: "::1", Count: 1},
		at:      ipv6AddrType,
		wantErr: "not an IPv6 address",
	}, {
		desc:    "min greater than max MAC-48",
		addrs:   &opb.AddressRange{Min: "02:03:04:05:06:07", Max: "01:02:03:04:05:06", Step: "00:00:00:00:00:01", Count: 1},
		at:      mac48AddrType,
		wantErr: "greater than max",
	}, {
		desc:    "min greater than max IPv4",
		addrs:   &opb.AddressRange{Min: "2.3.4.5", Max: "1.2.3.4", Step: "0.0.0.1", Count: 1},
		at:      ipv4AddrType,
		wantErr: "greater than max",
	}, {
		desc:    "min greater than max IPv6",
		addrs:   &opb.AddressRange{Min: "2:3:4:5:6:7:8:9", Max: "1:2:3:4:5:6:7:8", Step: "::1", Count: 1},
		at:      ipv6AddrType,
		wantErr: "greater than max",
	}, {
		desc:    "count cannot fit MAC-48",
		addrs:   &opb.AddressRange{Min: "00:00:00:00:00:01", Max: "00:00:00:00:00:09", Step: "00:00:00:00:00:03", Count: 4},
		at:      mac48AddrType,
		wantErr: "cannot fit",
	}, {
		desc:    "count cannot fit MAC-48, default step",
		addrs:   &opb.AddressRange{Min: "00:00:00:00:00:01", Max: "00:00:00:00:00:02", Count: 3},
		at:      mac48AddrType,
		wantErr: "cannot fit",
	}, {
		desc:    "count cannot fit IPv4",
		addrs:   &opb.AddressRange{Min: "0.0.0.1", Max: "0.0.0.9", Step: "0.0.0.3", Count: 4},
		at:      ipv4AddrType,
		wantErr: "cannot fit",
	}, {
		desc:    "count cannot fit IPv4, default step",
		addrs:   &opb.AddressRange{Min: "0.0.0.1", Max: "0.0.0.2", Count: 3},
		at:      ipv4AddrType,
		wantErr: "cannot fit",
	}, {
		desc:    "count cannot fit IPv6",
		addrs:   &opb.AddressRange{Min: "::1", Max: "::9", Step: "::3", Count: 4},
		at:      ipv6AddrType,
		wantErr: "cannot fit",
	}, {
		desc:    "count cannot fit IPv6, default step",
		addrs:   &opb.AddressRange{Min: "::1", Max: "::2", Count: 3},
		at:      ipv6AddrType,
		wantErr: "cannot fit",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if err := setAddrRangeField(&ixconfig.TrafficField{}, test.at, test.addrs); err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("setAddrRangeField(<field>, %q, %v): got error %v, want error %q", test.addrs, test.at, err, test.wantErr)
			}
		})
	}
}

func TestParseIP(t *testing.T) {
	tests := []struct {
		desc, ip       string
		wantIP, wantV6 bool
	}{{
		desc: "invalid IP",
		ip:   "1.1.1",
	}, {
		desc:   "IPv4",
		ip:     "1.1.1.1",
		wantIP: true,
	}, {
		desc:   "IPv6",
		ip:     "aa::",
		wantIP: true,
		wantV6: true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotIP, gotV6 := parseIP(test.ip)
			if (gotIP != nil) != test.wantIP {
				t.Errorf("parseIP(%q): unexpected result: parsed? %t, wanted parsed? %t", test.ip, gotIP != nil, test.wantIP)
			}
			if gotV6 != test.wantV6 {
				t.Errorf("parseIP(%q): unexpected result: is v6? %t, wanted v6? %t", test.ip, gotV6, test.wantV6)
			}
		})
	}
}
