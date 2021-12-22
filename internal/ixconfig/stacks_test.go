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

package ixconfig

import (
	"strings"
	"testing"
)

func TestEthernetStack(t *testing.T) {
	stack := NewEthernetStack(0)
	wantStackAlias := "ethernet-1"
	wantFieldAliases := map[string]*TrafficField{
		"ethernet.header.sourceAddress-2":      stack.SourceAddress(),
		"ethernet.header.destinationAddress-1": stack.DestinationAddress(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for EthernetStack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for EthernetStack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for EthernetStack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}

func TestVlanStack(t *testing.T) {
	stack := NewVlanStack(1)
	wantStackAlias := "vlan-2"
	wantFieldAliases := map[string]*TrafficField{
		"vlan.header.vlanTag.vlanID-3": stack.VlanTagVlanID(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for VlanStack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for VlanStack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for VlanStack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}

func TestGreStack(t *testing.T) {
	stack := NewGreStack(1)
	wantStackAlias := "gre-2"
	wantFieldAliases := map[string]*TrafficField{
		"gre.header.keyHolder.key-11":              stack.KeyHolderKey(),
		"gre.header.sequenceHolder.sequenceNum-13": stack.SequenceHolderSequenceNum(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for GreStack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for GreStack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for GreStack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}

func TestIpv4Stack(t *testing.T) {
	stack := NewIpv4Stack(2)
	wantStackAlias := "ipv4-3"
	wantFieldAliases := map[string]*TrafficField{
		"ipv4.header.priority.raw-3":    stack.PriorityRaw(),
		"ipv4.header.flags.fragment-21": stack.FlagsFragment(),
		"ipv4.header.ttl-24":            stack.Ttl(),
		"ipv4.header.srcIp-27":          stack.SrcIp(),
		"ipv4.header.dstIp-28":          stack.DstIp(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for Ipv4Stack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for Ipv4Stack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for Ipv4Stack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}

func TestIpv6Stack(t *testing.T) {
	stack := NewIpv6Stack(2)
	wantStackAlias := "ipv6-3"
	wantFieldAliases := map[string]*TrafficField{
		"ipv6.header.versionTrafficClassFlowLabel.trafficClass-2": stack.VersionTrafficClassFlowLabelTrafficClass(),
		"ipv6.header.versionTrafficClassFlowLabel.flowLabel-3":    stack.VersionTrafficClassFlowLabelFlowLabel(),
		"ipv6.header.hopLimit-6":                                  stack.HopLimit(),
		"ipv6.header.srcIP-7":                                     stack.SrcIP(),
		"ipv6.header.dstIP-8":                                     stack.DstIP(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for Ipv6Stack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for Ipv6Stack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for Ipv6Stack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}

func TestMplsStack(t *testing.T) {
	stack := NewMplsStack(1)
	wantStackAlias := "mpls-2"
	wantFieldAliases := map[string]*TrafficField{
		"mpls.label.value-1":        stack.Value(),
		"mpls.label.experimental-2": stack.Experimental(),
		"mpls.label.ttl-4":          stack.Ttl(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for MplsStack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for MplsStack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for MplsStack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}

func TestTcpStack(t *testing.T) {
	stack := NewTcpStack(3)
	wantStackAlias := "tcp-4"
	wantFieldAliases := map[string]*TrafficField{
		"tcp.header.srcPort-1":        stack.SrcPort(),
		"tcp.header.dstPort-2":        stack.DstPort(),
		"tcp.header.sequenceNumber-3": stack.SequenceNumber(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for TcpStack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for TcpStack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for TcpStack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}

func TestUdpStack(t *testing.T) {
	stack := NewUdpStack(3)
	wantStackAlias := "udp-4"
	wantFieldAliases := map[string]*TrafficField{
		"udp.header.srcPort-1": stack.SrcPort(),
		"udp.header.dstPort-2": stack.DstPort(),
	}
	if !strings.Contains(stack.Xpath.String(), wantStackAlias) {
		t.Errorf("Unexpected xpath %q for UdpStack (wanted alias text to contain %q)", stack.Xpath.String(), wantStackAlias)
	}
	for alias, field := range wantFieldAliases {
		xp := field.XPath().String()
		if !strings.Contains(xp, wantStackAlias) {
			t.Errorf("Unexpected xpath %q for UdpStack field (wanted alias text to contain stack alias portion %q)", xp, wantStackAlias)
		}
		if !strings.Contains(xp, alias) {
			t.Errorf("Unexpected xpath %q for UdpStack field (wanted alias text to contain %q)", xp, alias)
		}
	}
}
