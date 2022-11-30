// Copyright 2020 Google Inc.
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
	"testing"

	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ygot/ygot"
	"google.golang.org/protobuf/encoding/prototext"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// isEmptyDiff is a simple helper to determine whether the gNMI Notification
// returned by the ygot.Diff function reflects no diff between the original and
// modified structs provided.
func isEmptyDiff(n *gpb.Notification) bool {
	return (len(n.Update) == 0 && len(n.Delete) == 0)
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		name             string
		tables           map[string]ixweb.StatTable
		want             *oc.Root
		wantErrSubstring string
	}{{
		name: "single view, success",
		tables: map[string]ixweb.StatTable{
			portCPUStatsCaption: {{
				"Port Name":        "port1",
				"Total Memory(KB)": "420",
				"Free Memory(KB)":  "100",
				"%CPU Load":        "100",
			}},
		},
		want: func() *oc.Root {
			d := &oc.Root{}
			p := d.GetOrCreateComponent("port1")
			c := d.GetOrCreateComponent("port1_CPU")

			p.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
			c.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU

			p.GetOrCreateSubcomponent("port1_CPU")
			c.Parent = ygot.String("port1")

			p.Memory = &oc.Component_Memory{
				Available: ygot.Uint64(100),
				Utilized:  ygot.Uint64(320),
			}

			c.GetOrCreateCpu().GetOrCreateUtilization().Instant = ygot.Uint8(100)

			return d
		}(),
	}, {
		name: "one valid and one unknown view, success",
		tables: map[string]ixweb.StatTable{
			portCPUStatsCaption: {{
				"Port Name": "port1",
				"%CPU Load": "100",
			}},
			"Some Unknown View": {{
				"Colour": "burnt-umber",
			}},
		},
		want: func() *oc.Root {
			d := &oc.Root{}
			p := d.GetOrCreateComponent("port1")
			c := d.GetOrCreateComponent("port1_CPU")

			p.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
			c.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU

			p.GetOrCreateSubcomponent("port1_CPU")
			c.Parent = ygot.String("port1")

			c.GetOrCreateCpu().GetOrCreateUtilization().Instant = ygot.Uint8(100)

			return d
		}(),
	}, {
		name: "invalid known view, err logging and no failure",
		tables: map[string]ixweb.StatTable{
			portCPUStatsCaption: {{
				"%CPU Load": "42",
			}},
		},
	}, {
		name: "multiple stats",
		tables: map[string]ixweb.StatTable{
			portCPUStatsCaption: {{
				"Port Name": "port1",
				"%CPU Load": "100",
			}},
			portStatsCaption: {{
				"Port Name": "port1",
				"Bytes Rx.": "10",
				"Bytes Tx.": "20",
			}},
		},
		want: func() *oc.Root {
			d := &oc.Root{}
			p := d.GetOrCreateComponent("port1")
			p.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
			p.GetOrCreateSubcomponent("port1_CPU")

			c := d.GetOrCreateComponent("port1_CPU")
			c.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU
			c.GetOrCreateCpu().Utilization = &oc.Component_Cpu_Utilization{
				Instant: ygot.Uint8(100),
			}
			c.Parent = ygot.String("port1")

			i := d.GetOrCreateInterface("port1")
			i.Type = oc.IETFInterfaces_InterfaceType_ethernetCsmacd
			ic := i.GetOrCreateCounters()
			ic.InOctets = ygot.Uint64(10)
			ic.OutOctets = ygot.Uint64(20)

			return d
		}(),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translate(&Stats{Tables: tt.tables})

			if err != nil {
				t.Fatalf("Want: No Error, Got: %v", err)
			}

			diff, err := ygot.Diff(tt.want, got)
			if err != nil {
				t.Fatalf("cannot diff received output, %v", err)
			}
			if !isEmptyDiff(diff) {
				t.Fatalf("did not get expected mapped structure, diff (orig=want, modified=got):\n%s", prototext.Format(diff))
			}
		})
	}
}

func TestTranslatePortCPUStats(t *testing.T) {
	tests := []struct {
		name             string
		table            ixweb.StatTable
		want             *oc.Root
		wantErrSubstring string
	}{{
		name: "complete single component mapping",
		table: ixweb.StatTable{{
			"Port Name":        "port1",
			"Total Memory(KB)": "42",
			"Free Memory(KB)":  "21",
			"%CPU Load":        "100",
		}},
		want: func() *oc.Root {
			s := &oc.Root{}
			p := s.GetOrCreateComponent("port1")
			p.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
			p.NewSubcomponent("port1_CPU")

			m := p.GetOrCreateMemory()
			m.Available = ygot.Uint64(21)
			m.Utilized = ygot.Uint64(21)

			c := s.GetOrCreateComponent("port1_CPU")
			c.Parent = ygot.String("port1")
			c.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU
			c.GetOrCreateCpu().GetOrCreateUtilization().Instant = ygot.Uint8(uint8(100))

			return s
		}(),
	}, {
		name: "partial component mapping",
		table: ixweb.StatTable{{
			"Port Name": "port2",
			"%CPU Load": "42",
		}},
		want: func() *oc.Root {
			s := &oc.Root{}

			p := s.GetOrCreateComponent("port2")
			p.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
			p.NewSubcomponent("port2_CPU")

			c := s.GetOrCreateComponent("port2_CPU")
			c.Parent = ygot.String("port2")
			c.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU
			c.GetOrCreateCpu().GetOrCreateUtilization().Instant = ygot.Uint8(uint8(42))

			return s
		}(),
	}, {
		name: "multiple components",
		table: ixweb.StatTable{{
			"Port Name": "port1",
			"%CPU Load": "1",
		}, {
			"Port Name": "port2",
			"%CPU Load": "2",
		}},
		want: func() *oc.Root {
			s := &oc.Root{}

			for _, ep := range []struct {
				name string
				load uint8
			}{
				{name: "port1", load: 1},
				{name: "port2", load: 2},
			} {
				p := s.GetOrCreateComponent(ep.name)
				p.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_PORT
				p.NewSubcomponent(fmt.Sprintf("%s_CPU", ep.name))

				c := s.GetOrCreateComponent(fmt.Sprintf("%s_CPU", ep.name))
				c.Parent = ygot.String(ep.name)
				c.Type = oc.PlatformTypes_OPENCONFIG_HARDWARE_COMPONENT_CPU
				c.GetOrCreateCpu().GetOrCreateUtilization().Instant = ygot.Uint8(ep.load)
			}

			return s
		}(),
	}, {
		name: "missing port name",
		table: ixweb.StatTable{{
			"%CPU Load": "42",
		}},
		wantErrSubstring: "required key",
	}, {
		name: "invalid total memory",
		table: ixweb.StatTable{{
			"Port Name":        "ixia2/foo",
			"Total Memory(KB)": "invalid",
			"Free Memory(KB)":  "20",
		}},
		wantErrSubstring: "Total Memory",
	}, {
		name: "invalid free memory",
		table: ixweb.StatTable{{
			"Port Name":        "ixia2/bar",
			"Total Memory(KB)": "10",
			"Free Memory(KB)":  "invalid",
		}},
		wantErrSubstring: "Free Memory",
	}, {
		name: "invalid CPU utilisation",
		table: ixweb.StatTable{{
			"Port Name": "ixia2/baz",
			"%CPU Load": "invalid",
		}},
		wantErrSubstring: "CPU Load",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translatePortCPUStats(tt.table, nil, nil)
			if diff := errdiff.Substring(err, tt.wantErrSubstring); diff != "" {
				t.Fatalf("did not get expected error, %s", diff)
			}
			if err != nil {
				return
			}

			diff, err := ygot.Diff(tt.want, got)
			if err != nil {
				t.Fatalf("cannot diff received output, %v", err)
			}
			if !isEmptyDiff(diff) {
				t.Fatalf("did not get expected mapped structure, diff (orig=want, modified=got):\n%s", prototext.Format(diff))
			}
		})
	}
}

func TestTranslatePortStats(t *testing.T) {
	tests := []struct {
		name             string
		table            ixweb.StatTable
		want             *oc.Root
		wantErrSubstring string
	}{{
		name: "single port statistics",
		table: ixweb.StatTable{{
			"Port Name":        "port1",
			"Valid Frames Rx.": "10",
			"Frames Tx.":       "20",
			"Bytes Rx.":        "30",
			"Bytes Tx.":        "40",
			"Tx. Rate (bps)":   "1024",
			"Rx. Rate (bps)":   "2048",
			"CRC Errors":       "42",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			i := d.GetOrCreateInterface("port1")
			i.Counters = &oc.Interface_Counters{
				InOctets:  ygot.Uint64(30),
				InPkts:    ygot.Uint64(10),
				OutOctets: ygot.Uint64(40),
				OutPkts:   ygot.Uint64(20),
			}

			i.GetOrCreateEthernet().GetOrCreateCounters().InCrcErrors = ygot.Uint64(42)

			i.InRate = float32Bytes(2048.0)
			i.OutRate = float32Bytes(1024.0)
			i.Type = oc.IETFInterfaces_InterfaceType_ethernetCsmacd
			return d

		}(),
	}, {
		name: "two ports statistics",
		table: ixweb.StatTable{{
			"Port Name": "port1",
			"Bytes Rx.": "10",
			"Bytes Tx.": "20",
		}, {
			"Port Name": "port2",
			"Bytes Rx.": "100",
			"Bytes Tx.": "200",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			for intf, counters := range map[string][]uint64{
				"port1": []uint64{10, 20},
				"port2": []uint64{100, 200},
			} {
				i := d.GetOrCreateInterface(intf)
				i.Type = oc.IETFInterfaces_InterfaceType_ethernetCsmacd
				i.Counters = &oc.Interface_Counters{
					InOctets:  ygot.Uint64(counters[0]),
					OutOctets: ygot.Uint64(counters[1]),
				}
			}

			return d
		}(),
	}, {
		name: "single interface oper status up",
		table: ixweb.StatTable{{
			"Port Name":  "p4",
			"Link State": "Link Up",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			i := d.GetOrCreateInterface("p4")
			i.OperStatus = oc.Interface_OperStatus_UP
			i.Type = oc.IETFInterfaces_InterfaceType_ethernetCsmacd
			return d
		}(),
	}, {
		name: "single interface link down",
		table: ixweb.StatTable{{
			"Port Name":  "p5",
			"Link State": "Link Down",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			i := d.GetOrCreateInterface("p5")
			i.OperStatus = oc.Interface_OperStatus_DOWN
			i.Type = oc.IETFInterfaces_InterfaceType_ethernetCsmacd
			return d
		}(),
	}, {
		name: "single interface no pcs lock",
		table: ixweb.StatTable{{
			"Port Name":  "p5",
			"Link State": "No PCS Lock",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			i := d.GetOrCreateInterface("p5")
			i.OperStatus = oc.Interface_OperStatus_DOWN
			i.Type = oc.IETFInterfaces_InterfaceType_ethernetCsmacd
			return d
		}(),
	}, {
		name: "invalid input for uint64 statistic",
		table: ixweb.StatTable{{
			"Port Name": "port42",
			"Bytes Rx.": "four hundred and ninety seven",
		}},
		wantErrSubstring: "Bytes Rx.",
	}, {
		name: "invalid input for float32 statistic",
		table: ixweb.StatTable{{
			"Port Name":      "port48",
			"Tx. Rate (bps)": "one point eight gigabits per second",
		}},
		wantErrSubstring: "Tx. Rate (bps)",
	}, {
		name: "missing port name",
		table: ixweb.StatTable{{
			"Bytes Rx.": "42",
		}},
		wantErrSubstring: "required key",
	}, {
		name: "invalid port status",
		table: ixweb.StatTable{{
			"Port Name":  "port42",
			"Link State": "Unmappable Value",
		}},
		wantErrSubstring: "unmappable port link state",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translatePortStats(tt.table, nil, nil)
			if diff := errdiff.Substring(err, tt.wantErrSubstring); diff != "" {
				t.Fatalf("did not get expected error, %s", diff)
			}
			if err != nil {
				return
			}

			diff, err := ygot.Diff(tt.want, got)
			if err != nil {
				t.Fatalf("cannot diff received output, %v", err)
			}
			if !isEmptyDiff(diff) {
				t.Fatalf("did not get expected mapped struct, delta from want to got:\n%s", prototext.Format(diff))
			}
		})
	}
}

func TestTranslateTrafficItemStats(t *testing.T) {
	tests := []struct {
		name             string
		table            ixweb.StatTable
		itFlows          []string
		want             *oc.Root
		wantErrSubstring string
	}{{
		name: "single flow statistics",
		table: ixweb.StatTable{{
			"Traffic Item":  "traffic1",
			"Loss %":        "2.1",
			"Rx Bytes":      "100",
			"Rx Frames":     "10",
			"Rx Frame Rate": "1",
			"Rx Rate (bps)": "1024",
			"Tx Frames":     "20",
			"Tx Frame Rate": "2",
			"Tx Rate (bps)": "2048",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			f := d.GetOrCreateFlow("traffic1")
			f.Counters = &oc.Flow_Counters{
				InOctets: ygot.Uint64(100),
				InPkts:   ygot.Uint64(10),
				OutPkts:  ygot.Uint64(20),
			}
			f.LossPct = float32Bytes(2.1)
			f.InFrameRate = float32Bytes(1)
			f.OutFrameRate = float32Bytes(2)
			f.InRate = float32Bytes(1024)
			f.OutRate = float32Bytes(2048)
			return d
		}(),
	}, {
		name: "missing traffic item",
		table: ixweb.StatTable{{
			"Loss %": "2.1",
		}},
		wantErrSubstring: "required key",
	}, {
		name: "invalid input for uint64 statistic",
		table: ixweb.StatTable{{
			"Traffic Item": "traffic1",
			"Rx Bytes":     "one hundred",
		}},
		wantErrSubstring: "Rx Bytes",
	}, {
		name: "invalid input for float32 statistic",
		table: ixweb.StatTable{{
			"Traffic Item": "traffic1",
			"Loss %":       "two point one",
		}},
		wantErrSubstring: "Loss %",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translateTrafficItemStats(tt.table, tt.itFlows, nil)
			if diff := errdiff.Substring(err, tt.wantErrSubstring); diff != "" {
				t.Fatalf("did not get expected error, %s", diff)
			}
			if err != nil {
				return
			}

			diff, err := ygot.Diff(tt.want, got)
			if err != nil {
				t.Fatalf("cannot diff received output, %v", err)
			}
			if !isEmptyDiff(diff) {
				t.Fatalf("did not get expected mapped struct, delta from want to got:\n%s", prototext.Format(diff))
			}
		})
	}
}

func TestTranslateFlowStats(t *testing.T) {
	tests := []struct {
		name             string
		table            ixweb.StatTable
		want             *oc.Root
		wantErrSubstring string
	}{{
		name: "single flow statistics - ingress tracking disabled",
		table: ixweb.StatTable{{
			"Traffic Item":  "traffic1",
			"Loss %":        "2.1",
			"Rx Bytes":      "100",
			"Rx Frames":     "10",
			"Rx Frame Rate": "1",
			"Rx Rate (bps)": "1024",
			"Tx Frames":     "20",
			"Tx Frame Rate": "2",
			"Tx Rate (bps)": "2048",
			"Rx Port":       "port1",
			"Tx Port":       "Eth1",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			f := d.GetOrCreateFlow("traffic1")
			it := f.GetOrCreateIngressTracking("Eth1", "port1", oc.IngressTracking_MplsLabel_NO_LABEL, "", "", "", "", 0)
			it.Counters = &oc.Flow_IngressTracking_Counters{
				InOctets: ygot.Uint64(100),
				InPkts:   ygot.Uint64(10),
				OutPkts:  ygot.Uint64(20),
			}
			it.LossPct = float32Bytes(2.1)
			it.InFrameRate = float32Bytes(1)
			it.OutFrameRate = float32Bytes(2)
			it.InRate = float32Bytes(1024)
			it.OutRate = float32Bytes(2048)
			return d
		}(),
	}, {
		name: "ingress tracking statistics",
		table: ixweb.StatTable{{
			"Traffic Item":              "traffic1",
			"Loss %":                    "2.1",
			"Rx Bytes":                  "100",
			"Rx Frames":                 "10",
			"Rx Frame Rate":             "1",
			"Rx Rate (bps)":             "1024",
			"Tx Frames":                 "20",
			"Tx Frame Rate":             "2",
			"Tx Rate (bps)":             "2048",
			"Rx Port":                   "port1",
			"Tx Port":                   "Eth1",
			"MPLS:Label Value":          "0",
			"Source Endpoint":           "255.255.255.255",
			"Dest Endpoint":             "de:ad:be:ee:ee:ef",
			"IPv4 :Source Address":      "1.1.1.1",
			"IPv4 :Destination Address": "2.2.2.2",
			"IPv6 :Source Address":      "1::",
			"IPv6 :Destination Address": "EE::",
			"IPv4 :Precedence":          "3",
			"VLAN:VLAN-ID":              "1",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			f := d.GetOrCreateFlow("traffic1")
			it := f.GetOrCreateIngressTracking("Eth1", "port1", oc.IngressTracking_MplsLabel_IPV4_EXPLICIT_NULL, "1.1.1.1", "2.2.2.2", "1::", "EE::", 1)
			it.Counters = &oc.Flow_IngressTracking_Counters{
				InOctets: ygot.Uint64(100),
				InPkts:   ygot.Uint64(10),
				OutPkts:  ygot.Uint64(20),
			}
			it.LossPct = float32Bytes(2.1)
			it.InFrameRate = float32Bytes(1)
			it.OutFrameRate = float32Bytes(2)
			it.InRate = float32Bytes(1024)
			it.OutRate = float32Bytes(2048)
			return d
		}(),
	}, {
		name: "missing traffic item",
		table: ixweb.StatTable{{
			"Loss %": "2.1",
		}},
		wantErrSubstring: "required key",
	}, {
		name: "invalid input for uint64 statistic",
		table: ixweb.StatTable{{
			"Traffic Item": "traffic1",
			"Rx Bytes":     "one hundred",
		}},
		wantErrSubstring: "Rx Bytes",
	}, {
		name: "invalid input for float32 statistic",
		table: ixweb.StatTable{{
			"Traffic Item": "traffic1",
			"Loss %":       "two point one",
		}},
		wantErrSubstring: "Loss %",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translateFlowStats(tt.table, nil, nil)
			if diff := errdiff.Substring(err, tt.wantErrSubstring); diff != "" {
				t.Fatalf("did not get expected error, %s", diff)
			}
			if err != nil {
				return
			}

			diff, err := ygot.Diff(tt.want, got)
			if err != nil {
				t.Fatalf("cannot diff received output, %v", err)
			}
			if !isEmptyDiff(diff) {
				t.Fatalf("did not get expected mapped struct, delta from want to got:\n%s", prototext.Format(diff))
			}
		})
	}
}

func TestTranslateEgressStats(t *testing.T) {
	tests := []struct {
		name             string
		table            ixweb.StatTable
		itFlows          []string
		want             *oc.Root
		wantErrSubstring string
	}{{
		name: "single flow",
		table: ixweb.StatTable{{
			"Egress Tracking": "MyFilter",
		}, {
			"Egress Tracking": "1",
			"Loss %":          "2.1",
			"Rx Bytes":        "100",
			"Rx Frames":       "10",
			"Rx Frame Rate":   "1",
			"Rx Rate (bps)":   "1024",
			"Tx Frames":       "20",
			"Tx Frame Rate":   "2",
			"Tx Rate (bps)":   "2048",
		}},
		want: func() *oc.Root {
			d := &oc.Root{}
			f := d.GetOrCreateFlow("traffic1")
			f.Filter = ygot.String("MyFilter")
			et := f.GetOrCreateEgressTracking("1")
			et.Counters = &oc.Flow_EgressTracking_Counters{
				InOctets: ygot.Uint64(100),
				InPkts:   ygot.Uint64(10),
				OutPkts:  ygot.Uint64(20),
			}
			et.LossPct = float32Bytes(2.1)
			et.InFrameRate = float32Bytes(1)
			et.OutFrameRate = float32Bytes(2)
			et.InRate = float32Bytes(1024)
			et.OutRate = float32Bytes(2048)
			return d
		}(),
	}, {
		name: "ingress tracking statistics",
		table: ixweb.StatTable{{
			"Egress Tracking":           "Custom: (8 bits at offset 184)",
			"Rx Port":                   "port1",
			"Tx Port":                   "Eth1",
			"MPLS:Label Value":          "0",
			"Source Endpoint":           "255.255.255.255",
			"Dest Endpoint":             "de:ad:be:ee:ee:ef",
			"IPv4 :Source Address":      "1.1.1.1",
			"IPv4 :Destination Address": "2.2.2.2",
			"IPv6 :Source Address":      "1::",
			"IPv6 :Destination Address": "EE::",
			"IPv4 :Precedence":          "3",
			"VLAN:VLAN-ID":              "1",
		}, {
			"Egress Tracking": "1",
			"Loss %":          "2.1",
			"Rx Bytes":        "100",
			"Rx Frames":       "10",
			"Rx Frame Rate":   "1",
			"Rx Rate (bps)":   "1024",
			"Tx Frames":       "20",
			"Tx Frame Rate":   "2",
			"Tx Rate (bps)":   "2048",
		}, {
			"Egress Tracking": "2",
			"Loss %":          "3.1",
			"Rx Bytes":        "200",
			"Rx Frames":       "20",
			"Rx Frame Rate":   "2",
			"Rx Rate (bps)":   "2048",
			"Tx Frames":       "30",
			"Tx Frame Rate":   "3",
			"Tx Rate (bps)":   "3072",
		}, {
			"Egress Tracking":           "Custom: (8 bits at offset 184)",
			"Rx Port":                   "port1",
			"Tx Port":                   "Eth1",
			"MPLS:Label Value":          "0",
			"Source Endpoint":           "255.255.255.255",
			"Dest Endpoint":             "de:ad:be:ee:ee:ef",
			"IPv4 :Source Address":      "1.1.1.1",
			"IPv4 :Destination Address": "2.2.2.2",
			"IPv6 :Source Address":      "1::",
			"IPv6 :Destination Address": "EE::",
			"IPv4 :Precedence":          "3",
			"VLAN:VLAN-ID":              "2",
		}, {
			"Egress Tracking": "3",
			"Loss %":          "4.1",
			"Rx Bytes":        "300",
			"Rx Frames":       "30",
			"Rx Frame Rate":   "3",
			"Rx Rate (bps)":   "3072",
			"Tx Frames":       "40",
			"Tx Frame Rate":   "4",
			"Tx Rate (bps)":   "4096",
		}, {
			"Egress Tracking": "4",
			"Loss %":          "5.1",
			"Rx Bytes":        "400",
			"Rx Frames":       "40",
			"Rx Frame Rate":   "4",
			"Rx Rate (bps)":   "4096",
			"Tx Frames":       "50",
			"Tx Frame Rate":   "5",
			"Tx Rate (bps)":   "5120",
		}},
		itFlows: []string{"traffic1"},
		want: func() *oc.Root {
			d := &oc.Root{}
			f := d.GetOrCreateFlow("traffic1")
			it := f.GetOrCreateIngressTracking("Eth1", "port1", oc.IngressTracking_MplsLabel_IPV4_EXPLICIT_NULL, "1.1.1.1", "2.2.2.2", "1::", "EE::", 1)
			it.Filter = ygot.String("Custom: (8 bits at offset 184)")
			et := it.GetOrCreateEgressTracking("1")
			et.Counters = &oc.Flow_IngressTracking_EgressTracking_Counters{
				InOctets: ygot.Uint64(100),
				InPkts:   ygot.Uint64(10),
				OutPkts:  ygot.Uint64(20),
			}
			et.LossPct = float32Bytes(2.1)
			et.InFrameRate = float32Bytes(1)
			et.OutFrameRate = float32Bytes(2)
			et.InRate = float32Bytes(1024)
			et.OutRate = float32Bytes(2048)
			et = it.GetOrCreateEgressTracking("2")
			et.Counters = &oc.Flow_IngressTracking_EgressTracking_Counters{
				InOctets: ygot.Uint64(200),
				InPkts:   ygot.Uint64(20),
				OutPkts:  ygot.Uint64(30),
			}
			et.LossPct = float32Bytes(3.1)
			et.InFrameRate = float32Bytes(2)
			et.OutFrameRate = float32Bytes(3)
			et.InRate = float32Bytes(2048)
			et.OutRate = float32Bytes(3072)
			it = f.GetOrCreateIngressTracking("Eth1", "port1", oc.IngressTracking_MplsLabel_IPV4_EXPLICIT_NULL, "1.1.1.1", "2.2.2.2", "1::", "EE::", 2)
			it.Filter = ygot.String("Custom: (8 bits at offset 184)")
			et = it.GetOrCreateEgressTracking("3")
			et.Counters = &oc.Flow_IngressTracking_EgressTracking_Counters{
				InOctets: ygot.Uint64(300),
				InPkts:   ygot.Uint64(30),
				OutPkts:  ygot.Uint64(40),
			}
			et.LossPct = float32Bytes(4.1)
			et.InFrameRate = float32Bytes(3)
			et.OutFrameRate = float32Bytes(4)
			et.InRate = float32Bytes(3072)
			et.OutRate = float32Bytes(4096)
			et = it.GetOrCreateEgressTracking("4")
			et.Counters = &oc.Flow_IngressTracking_EgressTracking_Counters{
				InOctets: ygot.Uint64(400),
				InPkts:   ygot.Uint64(40),
				OutPkts:  ygot.Uint64(50),
			}
			et.LossPct = float32Bytes(5.1)
			et.InFrameRate = float32Bytes(4)
			et.OutFrameRate = float32Bytes(5)
			et.InRate = float32Bytes(4096)
			et.OutRate = float32Bytes(5120)
			return d
		}(),
	}, {
		name: "missing traffic item",
		table: ixweb.StatTable{{
			"Loss %": "2.1",
		}},
		wantErrSubstring: "required key",
	}, {
		name: "invalid input for uint64 statistic",
		table: ixweb.StatTable{{
			"Traffic Item": "traffic1",
			"Rx Bytes":     "one hundred",
		}},
		wantErrSubstring: "Rx Bytes",
	}, {
		name: "invalid input for float32 statistic",
		table: ixweb.StatTable{{
			"Traffic Item": "traffic1",
			"Loss %":       "two point one",
		}},
		wantErrSubstring: "Loss %",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := translateEgressStats(tt.table, tt.itFlows, []string{"traffic1"})
			if diff := errdiff.Substring(err, tt.wantErrSubstring); diff != "" {
				t.Fatalf("translateEgressStats got unexpected error, %s", diff)
			}
			if err != nil {
				return
			}

			diff, err := ygot.Diff(tt.want, got)
			if err != nil {
				t.Fatalf("cannot diff received output, %v", err)
			}
			if !isEmptyDiff(diff) {
				t.Fatalf("translateEgressStats got unexpected diff:\n%s", prototext.Format(diff))
			}
		})
	}
}

func TestMplsLabelFromUint(t *testing.T) {
	tests := []struct {
		desc string
		in   *uint64
		want oc.Flow_IngressTracking_MplsLabel_Union
	}{{
		desc: "nil",
		want: oc.IngressTracking_MplsLabel_NO_LABEL,
	}, {
		desc: "enum",
		in:   ygot.Uint64(1),
		want: oc.IngressTracking_MplsLabel_ROUTER_ALERT,
	}, {
		desc: "other uint value",
		in:   ygot.Uint64(200),
		want: oc.UnionUint32(200),
	}, {
		desc: "overflow",
		in:   ygot.Uint64(uint64(2) << 31),
		want: oc.IngressTracking_MplsLabel_UNSET,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := mplsLabelFromUint(tt.in); got != tt.want {
				t.Errorf("mplsLabelFromUint(%v): got %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
