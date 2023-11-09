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

//go:build snippets

package snippets_test

import (
	"testing"
	"time"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/gnoigo/system"
	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/gnmi/otg"
	"github.com/openconfig/ondatra/gnoi"
	"github.com/openconfig/ondatra/netutil"

	"google3/ops/netops/lab/ondatra/b2bind/b2bindinit"
)

var expectedRate float64

func TestMain(m *testing.M) {
	ondatra.RunTests(m, b2bindinit.Init)
}

func TestGetDevicesFromTestbed(t *testing.T) {
	ate := ondatra.ATE(t, "ate")
	dut := ondatra.DUT(t, "dut")
	doSomethingWith(ate, dut)
}

func TestPushConfig(t *testing.T) {
	dut := ondatra.DUT(t, "dut")
	dut.Config().New().
		WithAristaText(`
			interface {{ port "port1" }}
			 description From Ixia
			 no switchport
			 ip address 192.168.31.1/30`).
		WithCiscoFile("path/to/cisco_config.txt").
		Push(t)
}

func TestCreateTopology(t *testing.T) {
	ate := ondatra.ATE(t, "ate")
	ap := ate.Port(t, "port")
	top := gosnappi.NewConfig()
	top.Ports().Add().SetName(ap.ID())

	dev := top.Devices().Add().SetName("dev")
	eth := dev.Ethernets().Add().SetName("eth").SetMac("02:00:01:01:01:01")
	eth.Connection().SetPortName(ap.ID())
	ip := eth.Ipv4Addresses().Add().SetName("ipv4").
		SetAddress("192.168.31.2").SetPrefix(30).
		SetGateway("192.168.31.1")
	bgp := dev.Bgp().SetRouterId(ip.Address()).
		Ipv4Interfaces().Add().SetIpv4Name(ip.Name())
	bgpPeer := bgp.Peers().Add().SetName("peer").
		SetPeerAddress("192.168.31.1").
		SetAsType(gosnappi.BgpV4PeerAsType.EBGP)
	bgpPeer.V4Routes().Add().SetName("routes").
		Addresses().Add().SetAddress("192.168.40.0").
		SetPrefix(30).SetCount(100)

	ate.OTG().PushConfig(t, top)
	ate.OTG().StartProtocols(t)
}

func addIPIntf(top gosnappi.Config, port *ondatra.Port, devName, mac, ip, gw string, prefixLen uint32) gosnappi.DeviceIpv4 {
	top.Ports().Add().SetName(port.ID())
	dev := top.Devices().Add().SetName(devName)
	eth := dev.Ethernets().Add().SetName(dev.Name() + ".Eth").SetMac(mac)
	eth.Connection().SetPortName(port.ID())
	return eth.Ipv4Addresses().Add().SetName(eth.Name() + ".IPv4").
		SetAddress(ip).SetGateway(gw).SetPrefix(prefixLen)
}

func addIntfs(t *testing.T, ate *ondatra.ATEDevice, top gosnappi.Config, srcIntfMAC string) (gosnappi.DeviceIpv4, gosnappi.DeviceIpv4) {
	return addIPIntf(top, ate.Port(t, "port1"), "dev1", srcIntfMAC, "192.168.31.2", "192.168.31.1", 30),
		addIPIntf(top, ate.Port(t, "port2"), "dev2", "02:00:01:01:01:02", "192.168.32.2", "192.168.32.1", 30)
}

func TestGenerateTraffic(t *testing.T) {
	ate := ondatra.ATE(t, "ate")
	top := gosnappi.NewConfig()
	// Two IP interfaces configured similarly to TestCreateTopology.
	srcIntfMAC := "02:00:01:01:01:01"
	ip1, ip2 := addIntfs(t, ate, top, srcIntfMAC)

	flow := top.Flows().Add().SetName("Flow1")
	flow.Metrics().SetEnable(true)
	flow.TxRx().Device().
		SetTxNames([]string{ip1.Name()}).
		SetRxNames([]string{ip2.Name()})

	flow.Packet().Add().Ethernet().Src().
		SetValue(srcIntfMAC)
	ipHdr := flow.Packet().Add().Ipv4()
	ipHdr.Src().SetValue(ip1.Address())
	ipHdr.Dst().SetValue(ip2.Address())

	ate.OTG().PushConfig(t, top)
	ate.OTG().StartProtocols(t)
	ate.OTG().StartTraffic(t)
	time.Sleep(3 * time.Minute)
	ate.OTG().StopTraffic(t)
}

func TestQueryTelemetry(t *testing.T) {
	dut := ondatra.DUT(t, "dut")
	dp := dut.Port(t, "port1")
	ate := ondatra.ATE(t, "ate")
	ap := ate.Port(t, "port1")

	ds := gnmi.Get(t, dut, gnmi.OC().Interface(dp.Name()).OperStatus().State())
	as := gnmi.Get(t, ate.OTG(), gnmi.OTG().Port(ap.ID()).Link().State())
	if want := oc.Interface_OperStatus_UP; ds != want {
		t.Errorf("Get(DUT port1 status): got %v, want %v", ds, want)
	}
	if want := otg.Port_Link_UP; as != want {
		t.Errorf("Get(ATE port1 status): got %v, want %v", as, want)
	}

	p := gnmi.Collect(t, dut, gnmi.OC().Interface(`{{ port "port2" }}`).Counters().InPkts().State(), time.Minute)
	if got := netutil.MeanRate(t, p.Await(t)); got != expectedRate {
		t.Fatalf("Got unexpected input rate %.4f, want: %.4f", got, expectedRate)
	}
}

func TestExecuteOperation(t *testing.T) {
	dut := ondatra.DUT(t, "dut")
	gnoi.Execute(t, dut, system.NewPingOperation().Destination("8.8.8.8"))

	ate := ondatra.ATE(t, "ate")
	ap := dut.Port(t, "port1")
	portStateAction := gosnappi.NewControlState()
	portStateAction.Port().Link().
		SetPortNames([]string{ap.ID()}).
		SetState(gosnappi.StatePortLinkState.DOWN)
	ate.OTG().SetControlState(t, portStateAction)
}

func doSomethingWith(stuff ...any) {}
