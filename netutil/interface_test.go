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
	"strings"
	"testing"

	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
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
		num:    0,
		want:   "lo0",
	}, {
		desc:   "nokia",
		vendor: ondatra.NOKIA,
		num:    4,
		want:   "lo4",
	}, {
		desc:    "not supported",
		vendor:  ondatra.IXIA,
		wantErr: "not supported",
	}, {
		desc:    "negative index",
		vendor:  ondatra.ARISTA,
		num:     -3,
		wantErr: "negative",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			dut := &fakeDUT{vendor: test.vendor}
			got, err := loopbackInterface(dut, test.num)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("loopbackInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("loopbackInterface got %s, want %s", got, test.want)
			}
		})
	}
}

func TestAggregateInterface(t *testing.T) {
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
		want:   "Port-Channel2",
	}, {
		desc:   "cisco",
		vendor: ondatra.CISCO,
		num:    2,
		want:   "Bundle-Ether3",
	}, {
		desc:   "juniper",
		vendor: ondatra.JUNIPER,
		num:    0,
		want:   "ae0",
	}, {
		desc:   "nokia",
		vendor: ondatra.NOKIA,
		num:    4,
		want:   "lag5",
	}, {
		desc:    "not supported",
		vendor:  ondatra.IXIA,
		wantErr: "not supported",
	}, {
		desc:    "negative index",
		vendor:  ondatra.ARISTA,
		num:     -3,
		wantErr: "negative",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			dut := &fakeDUT{vendor: test.vendor}
			got, err := aggregateInterface(dut, test.num)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("aggregateInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("aggregateInterface got %s, want %s", got, test.want)
			}
		})
	}
}

func TestNextAggregateInterface(t *testing.T) {
	tests := []struct {
		desc    string
		vendor  ondatra.Vendor
		intfs   map[string]*oc.Interface
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
		want:   "ae0",
	}, {
		desc:   "nokia first",
		vendor: ondatra.NOKIA,
		want:   "lag1",
	}, {
		desc:   "between intfs",
		vendor: ondatra.ARISTA,
		intfs: map[string]*oc.Interface{
			"Port-Channel1": &oc.Interface{},
			"Port-Channel3": &oc.Interface{},
		},
		want: "Port-Channel2",
	}, {
		desc:    "not supported",
		vendor:  ondatra.IXIA,
		wantErr: "not supported",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			val := &ygnmi.Value[*oc.Root]{}
			if test.intfs != nil {
				val.SetVal(&oc.Root{
					Interface: test.intfs,
				})
			}
			dut := &fakeDUT{vendor: test.vendor}
			got, err := nextAggregateInterface(dut, val)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("nextAggregateInterface got err %v, want %s", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("nextAggregateInterface got %s, want %s", got, test.want)
			}
		})
	}
}

type fakeDUT struct {
	vendor ondatra.Vendor
}

func (d *fakeDUT) Vendor() ondatra.Vendor {
	return d.vendor
}

func (d *fakeDUT) Model() string {
	return ""
}
