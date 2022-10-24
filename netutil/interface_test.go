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

package netutil

import (
	"golang.org/x/net/context"
	"strings"
	"testing"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/testt"

	opb "github.com/openconfig/ondatra/proto"
)

func TestLoopbackInterface(t *testing.T) {
	tests := []struct {
		desc    string
		vendor  ondatra.Vendor
		num     int
		want    string
		wantErr string
	}{{
		desc:   "arista",
		vendor: ondatra.ARISTA,
		num:    1,
		want:   "Loopback1",
	}, {
		desc:   "cisco",
		vendor: ondatra.CISCO,
		num:    2,
		want:   "Loopback2",
	}, {
		desc:   "juniper",
		vendor: ondatra.JUNIPER,
		num:    3,
		want:   "lo3",
	}, {
		desc:   "nokia",
		vendor: ondatra.NOKIA,
		num:    4,
		want:   "lo4",
	}, {
		desc:    "no prefix",
		vendor:  ondatra.IXIA,
		wantErr: "no loopback interface prefix",
	}, {
		desc:    "negative num",
		vendor:  ondatra.ARISTA,
		num:     -3,
		wantErr: "negative number",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := loopbackInterface(test.vendor, test.num)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("loopbackInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("loopbackInterface got %s, want %s", got, test.want)
			}
		})
	}
}

func TestBundleInterface(t *testing.T) {
	tests := []struct {
		desc    string
		vendor  ondatra.Vendor
		num     int
		want    string
		wantErr string
	}{{
		desc:   "arista",
		vendor: ondatra.ARISTA,
		num:    3,
		want:   "Port-Channel3",
	}, {
		desc:   "cisco",
		vendor: ondatra.CISCO,
		num:    13,
		want:   "Bundle-Ether13",
	}, {
		desc:   "juniper",
		vendor: ondatra.JUNIPER,
		num:    113,
		want:   "ae113",
	}, {
		desc:   "nokia",
		vendor: ondatra.NOKIA,
		num:    24,
		want:   "lag24",
	}, {
		desc:    "no prefix",
		vendor:  ondatra.IXIA,
		wantErr: "no bundle interface prefix",
	}, {
		desc:    "negative num",
		vendor:  ondatra.ARISTA,
		num:     -3,
		wantErr: "negative number",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := bundleInterface(test.vendor, test.num)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("bundleInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("bundleInterface got %s, want %s", got, test.want)
			}
		})
	}
}

func TestVLANInterface(t *testing.T) {
	tests := []struct {
		desc    string
		vendor  ondatra.Vendor
		num     int
		want    string
		wantErr string
	}{{
		desc:   "arista",
		vendor: ondatra.ARISTA,
		num:    105,
		want:   "Vlan105",
	}, {
		desc:   "cisco",
		vendor: ondatra.CISCO,
		num:    65,
		want:   "BVI65",
	}, {
		desc:   "juniper",
		vendor: ondatra.JUNIPER,
		num:    5,
		want:   "irb.5",
	}, {
		desc:   "nokia",
		vendor: ondatra.NOKIA,
		num:    6,
		want:   "irb1.6",
	}, {
		desc:    "no prefix",
		vendor:  ondatra.IXIA,
		wantErr: "no VLAN interface prefix",
	}, {
		desc:    "negative num",
		vendor:  ondatra.CISCO,
		num:     -5,
		wantErr: "negative number",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := vlanInterface(test.vendor, test.num)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("vlanInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("vlanInterface got %s, want %s", got, test.want)
			}
		})
	}
}

func TestNextBundleInterface(t *testing.T) {
	tests := []struct {
		desc    string
		vendor  ondatra.Vendor
		intfs   map[string]*telemetry.Interface
		want    string
		wantErr string
	}{{
		desc:   "arista first",
		vendor: ondatra.ARISTA,
		want:   "Port-Channel1",
	}, {
		desc:   "cisco first",
		vendor: ondatra.CISCO,
		want:   "Bundle-Ether1",
	}, {
		desc:   "juniper first",
		vendor: ondatra.JUNIPER,
		want:   "ae1",
	}, {
		desc:   "nokia first",
		vendor: ondatra.NOKIA,
		want:   "lag1",
	}, {
		desc:   "between intfs",
		vendor: ondatra.ARISTA,
		intfs: map[string]*telemetry.Interface{
			"Port-Channel1": &telemetry.Interface{},
			"Port-Channel3": &telemetry.Interface{},
		},
		want: "Port-Channel2",
	}, {
		desc:    "no range",
		vendor:  ondatra.IXIA,
		wantErr: "no bundle interface",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			val := &telemetry.QualifiedDevice{}
			if test.intfs != nil {
				val.SetVal(&telemetry.Device{
					Interface: test.intfs,
				})
			}
			got, err := nextBundleInterface(t, test.vendor, val)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("nextBundleInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("nextBundleInterface got %s, want %s", got, test.want)
			}
		})
	}
}

func TestNextVLANInterface(t *testing.T) {
	tests := []struct {
		desc    string
		vendor  ondatra.Vendor
		intfs   map[string]*telemetry.Interface
		want    string
		wantErr string
	}{{
		desc:   "arista first",
		vendor: ondatra.ARISTA,
		want:   "Vlan1",
	}, {
		desc:   "cisco first",
		vendor: ondatra.CISCO,
		want:   "BVI1",
	}, {
		desc:   "juniper first",
		vendor: ondatra.JUNIPER,
		want:   "irb.1",
	}, {
		desc:   "nokia first",
		vendor: ondatra.NOKIA,
		want:   "irb1.1",
	}, {
		desc:   "between intfs",
		vendor: ondatra.CISCO,
		intfs: map[string]*telemetry.Interface{
			"BVI1": &telemetry.Interface{},
			"BVI3": &telemetry.Interface{},
		},
		want: "BVI2",
	}, {
		desc:    "no range",
		vendor:  ondatra.IXIA,
		wantErr: "no VLAN interface",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			val := &telemetry.QualifiedDevice{}
			if test.intfs != nil {
				val.SetVal(&telemetry.Device{
					Interface: test.intfs,
				})
			}
			got, err := nextVLANInterface(t, test.vendor, val)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("nextVLANInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("nextVLANInterface got %s, want %s", got, test.want)
			}
		})
	}
}

func TestSetInterfaceState(t *testing.T) {
	var gotConfig string
	bdut := &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{}},
		PushConfigFn: func(_ context.Context, config string, _ bool) error {
			gotConfig = config
			return nil
		},
	}
	fakebind.Setup().WithReservation(
		&binding.Reservation{DUTs: map[string]binding.DUT{"dut": bdut}})
	dut := ondatra.DUT(t, "dut")

	tests := []struct {
		desc       string
		vendor     opb.Device_Vendor
		enable     bool
		wantConfig string
	}{{
		desc:       "Arista enable",
		vendor:     opb.Device_ARISTA,
		enable:     true,
		wantConfig: "no shutdown",
	}, {
		desc:       "Cisco enable",
		vendor:     opb.Device_CISCO,
		enable:     true,
		wantConfig: "no shutdown",
	}, {
		desc:       "Juniper enable",
		vendor:     opb.Device_JUNIPER,
		enable:     true,
		wantConfig: "delete: disable",
	}, {
		desc:       "Arista disable",
		vendor:     opb.Device_ARISTA,
		enable:     false,
		wantConfig: "  shutdown",
	}, {
		desc:       "Cisco disable",
		vendor:     opb.Device_CISCO,
		enable:     false,
		wantConfig: "  shutdown",
	}, {
		desc:       "Juniper disable",
		vendor:     opb.Device_JUNIPER,
		enable:     false,
		wantConfig: "  disable",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotConfig = ""
			bdut.AbstractDUT.Dims.Vendor = test.vendor
			SetInterfaceState(t, dut, "intf", test.enable)
			if !strings.Contains(gotConfig, test.wantConfig) {
				t.Errorf("SetInterfaceState() got config %q, want %q", gotConfig, test.wantConfig)
			}
		})
	}

	t.Run("unsupported vendor", func(t *testing.T) {
		bdut.AbstractDUT.Dims.Vendor = opb.Device_DELL
		msg := testt.ExpectFatal(t, func(t testing.TB) {
			SetInterfaceState(t, dut, "intf", true)
		})
		if want := "supported"; !strings.Contains(msg, want) {
			t.Errorf("SetInterfaceState got fatal msg %q, want %q", msg, want)
		}
	})
}
