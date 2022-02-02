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
	"github.com/openconfig/ondatra/binding/ixweb"
)

func TestParsePortStats(t *testing.T) {
	table := ixweb.StatTable{{
		"Port Name":        "1/2",
		"Line Speed":       "100GE",
		"Link State":       "Link Up",
		"Frames Tx.":       "1000",
		"Valid Frames Rx.": "750",
		"Bytes Tx.":        "10000",
		"Bytes Rx.":        "7500",
		"Tx. Rate (bps)":   "100.5",
		"Rx. Rate (bps)":   "110.6",
		"CRC Errors":       "5",
	}, {
		"Port Name": "3/4",
	}}
	want := []*portRow{{
		portName:  "1/2",
		lineSpeed: "100GE",
		linkState: "Link Up",
		framesTx:  pUint(1000),
		framesRx:  pUint(750),
		bytesTx:   pUint(10000),
		bytesRx:   pUint(7500),
		txRate:    pFloat(100.5),
		rxRate:    pFloat(110.6),
		crcErrs:   pUint(5),
	}, {
		portName: "3/4",
	}}

	got, err := parsePortStats(table)
	if err != nil {
		t.Fatalf("parsePortStats err = %v", err)
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(portRow{})); diff != "" {
		t.Errorf("parsePortStats unexpected diff (-want,+got): %s", diff)
	}
}

func TestParsePortCPUStats(t *testing.T) {
	table := ixweb.StatTable{{
		"Port Name":        "1/2",
		"Total Memory(KB)": "15000",
		"Free Memory(KB)":  "10000",
		"%CPU Load":        "50",
	}, {
		"Port Name": "3/4",
	}}
	want := []*portCPURow{{
		portName:    "1/2",
		totalMemory: pUint(15000),
		freeMemory:  pUint(10000),
		cpuLoad:     pUint(50),
	}, {
		portName: "3/4",
	}}

	got, err := parsePortCPUStats(table)
	if err != nil {
		t.Fatalf("parsePortCPUStats err = %v", err)
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(portCPURow{})); diff != "" {
		t.Errorf("parsePortCPUStats unexpected diff (-want,+got): %s", diff)
	}
}

func TestParseFlowStats(t *testing.T) {
	table := ixweb.StatTable{{
		"Traffic Item":              "flowA",
		"Tx Port":                   "1/2",
		"Rx Port":                   "3/4",
		"Tx Frames":                 "1000",
		"Rx Frames":                 "750",
		"Tx Frame Rate":             "50.5",
		"Rx Frame Rate":             "51.6",
		"Rx Bytes":                  "10000",
		"Tx Rate (bps)":             "100.5",
		"Rx Rate (bps)":             "110.6",
		"Loss %":                    "11.25",
		"IPv4 :Source Address":      "1.2.3.4",
		"IPv4 :Destination Address": "2.3.4.5",
		"IPv6 :Source Address":      "1:2:3:4:5:6:7:8",
		"IPv6 :Destination Address": "2:3:4:5:6:7:8:9",
		"MPLS:Label Value":          "12345",
	}, {
		"Traffic Item": "flowB",
	}}
	want := []*flowRow{{
		trafficItem: "flowA",
		txPort:      "1/2",
		rxPort:      "3/4",
		txFrames:    pUint(1000),
		rxFrames:    pUint(750),
		txFrameRate: pFloat(50.5),
		rxFrameRate: pFloat(51.6),
		rxBytes:     pUint(10000),
		txRate:      pFloat(100.5),
		rxRate:      pFloat(110.6),
		lossPct:     pFloat(11.25),
		srcIPv4:     "1.2.3.4",
		dstIPv4:     "2.3.4.5",
		srcIPv6:     "1:2:3:4:5:6:7:8",
		dstIPv6:     "2:3:4:5:6:7:8:9",
		mplsLabel:   pUint(12345),
	}, {
		trafficItem: "flowB",
	}}

	got, err := parseFlowStats(table)
	if err != nil {
		t.Fatalf("parseFlowStats err = %v", err)
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(flowRow{})); diff != "" {
		t.Errorf("parseFlowStats unexpected diff (-want,+got): %s", diff)
	}
}

func TestParseEgressStats(t *testing.T) {
	table := ixweb.StatTable{{
		"Egress Tracking":           "Custom: (4 bits offset 128)",
		"Traffic Item":              "flowA",
		"Tx Port":                   "1/2",
		"Rx Port":                   "3/4",
		"Tx Frames":                 "1000",
		"Rx Frames":                 "750",
		"Tx Frame Rate":             "50.5",
		"Rx Frame Rate":             "51.6",
		"Rx Bytes":                  "10000",
		"Tx Rate (bps)":             "100.5",
		"Rx Rate (bps)":             "110.6",
		"Loss %":                    "11.25",
		"IPv4 :Source Address":      "1.2.3.4",
		"IPv4 :Destination Address": "2.3.4.5",
		"IPv6 :Source Address":      "1:2:3:4:5:6:7:8",
		"IPv6 :Destination Address": "2:3:4:5:6:7:8:9",
		"MPLS:Label Value":          "12345",
	}, {
		"Egress Tracking": "2",
	}}
	want := []*egressRow{{
		filter: "Custom: (4 bits offset 128)",
		flowRow: &flowRow{
			trafficItem: "flowA",
			txPort:      "1/2",
			rxPort:      "3/4",
			txFrames:    pUint(1000),
			rxFrames:    pUint(750),
			txFrameRate: pFloat(50.5),
			rxFrameRate: pFloat(51.6),
			rxBytes:     pUint(10000),
			txRate:      pFloat(100.5),
			rxRate:      pFloat(110.6),
			lossPct:     pFloat(11.25),
			srcIPv4:     "1.2.3.4",
			dstIPv4:     "2.3.4.5",
			srcIPv6:     "1:2:3:4:5:6:7:8",
			dstIPv6:     "2:3:4:5:6:7:8:9",
			mplsLabel:   pUint(12345),
		},
	}, {
		filter:  "2",
		flowRow: &flowRow{},
	}}

	got, err := parseEgressStats(table)
	if err != nil {
		t.Fatalf("parseEgressStats err = %v", err)
	}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(egressRow{}, flowRow{})); diff != "" {
		t.Errorf("parseEgressStats unexpected diff (-want,+got): %s", diff)
	}
}

func pUint(i uint64) *uint64 {
	return &i
}

func pFloat(f float32) *float32 {
	return &f
}
