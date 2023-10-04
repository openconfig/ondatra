// Copyright 2023 Google LLC
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

package otg_test

import (
	"testing"
	"time"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/otg"
)

// Example demonstrates how to use the OTG API.
func Example() {
	t := new(testing.T) // In a real test, *testing.T would be a parameter.

	const (
		eth1Name   = "intf1.eth"
		eth2Name   = "intf2.eth"
		ipv41Name  = "intf1.ipv4"
		ipv42Name  = "intf2.ipv4"
		ateSrcMAC  = "02:00:01:01:01:01"
		ateDstMAC  = "02:00:02:01:01:01"
		ateSrcIP   = "192.0.2.2"
		ateDstIP   = "192.0.2.6"
		ateSrcGway = "192.0.2.1"
		ateDstGway = "192.0.2.5"
	)

	ate := ondatra.ATE(t, "ate")
	ap1 := ate.Port(t, "port1")
	ap2 := ate.Port(t, "port2")

	// Create an empty OTG config and add the ports.
	cfg := gosnappi.NewConfig()
	cfg.Ports().Add().SetName(ap1.ID())
	cfg.Ports().Add().SetName(ap2.ID())

	// Configure an interface on port 1.
	intf1 := cfg.Devices().Add().SetName("intf1")
	eth1 := intf1.Ethernets().Add().
		SetName(eth1Name).
		SetMac(ateSrcMAC)
	eth1.Connection().SetPortName(ap1.ID())
	eth1.Ipv4Addresses().Add().
		SetName(ipv41Name).
		SetAddress(ateSrcIP).
		SetPrefix(30).
		SetGateway(ateSrcGway)

	// Configure an interface on port 2.
	intf2 := cfg.Devices().Add().SetName("intf2")
	eth2 := intf2.Ethernets().Add().
		SetName(eth2Name).
		SetMac(ateDstMAC)
	eth2.Connection().SetPortName(ap2.ID())
	eth2.Ipv4Addresses().Add().
		SetName(ipv42Name).
		SetAddress(ateDstIP).
		SetPrefix(30).
		SetGateway(ateDstGway)

	// Setup a traffic flow.
	flow := cfg.Flows().Add().SetName("flow1")
	flow.TxRx().Device().
		SetTxNames([]string{ipv41Name}).
		SetRxNames([]string{ipv42Name})
	flow.Metrics().SetEnable(true)
	eth := flow.Packet().Add().Ethernet()
	eth.Src().SetValue(ateSrcMAC)
	ipv4 := flow.Packet().Add().Ipv4()
	ipv4.Src().SetValue(ateSrcIP)
	ipv4.Dst().SetValue("20.21.21.21")

	// Push config and start control-plane protocols.
	ate.OTG().PushConfig(t, cfg)
	ate.OTG().StartProtocols(t)

	// Use gNMI to wait for the configured ports to be up.
	gnmi.Await(t, ate.OTG(), gnmi.OTG().Port(ap1.ID()).Link().State(), time.Minute, otg.Port_Link_UP)
	gnmi.Await(t, ate.OTG(), gnmi.OTG().Port(ap2.ID()).Link().State(), time.Minute, otg.Port_Link_UP)

	// Run data-plane traffic and then stop traffic and protocols.
	ate.OTG().StartTraffic(t)
	time.Sleep(30 * time.Second)
	ate.OTG().StopTraffic(t)
	ate.OTG().StopProtocols(t)

	// Gather flow statistics and assert that there was no traffic loss.
	flowMetrics := gnmi.Get(t, ate.OTG(), gnmi.OTG().Flow(flow.Name()).Counters().State())
	outPkts := flowMetrics.GetOutPkts()
	inPkts := flowMetrics.GetInPkts()
	lossPct := 100 * float64(outPkts-inPkts) / float64(outPkts)
	if lossPct < 0 || lossPct > 1 {
		t.Errorf("Unexpected loss percentage for traffic flow: %f", lossPct)
	}
}
