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

	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/oc"
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
	ap1 := ate.Port(t, "port1")
	ap2 := ate.Port(t, "port2")
	top := ate.Topology().New()
	if1 := top.AddInterface("if1").WithPort(ap1)
	if2 := top.AddInterface("if2").WithPort(ap2)
	if1.IPv4().
		WithAddress("192.168.31.2/30").
		WithDefaultGateway("192.168.31.1")
	if2.IPv4().
		WithAddress("192.168.32.2/30").
		WithDefaultGateway("192.168.32.1")
	net := if2.AddNetwork("net")
	net.IPv4().
		WithAddress("192.168.40.0/30").
		WithCount(100)
	top.Push(t).StartProtocols(t)
}

func TestGenerateTraffic(t *testing.T) {
	ate := ondatra.ATE(t, "ate")
	ap1 := ate.Port(t, "port1")
	ap2 := ate.Port(t, "port2")
	top := ate.Topology().New()
	if1 := top.AddInterface("if1").WithPort(ap1)
	if2 := top.AddInterface("if2").WithPort(ap2)
	flow := ate.Traffic().NewFlow("Flow1").
		WithSrcEndpoints(if1).
		WithDstEndpoints(if2).
		WithFrameRatePct(50)
	ate.Traffic().Start(t, flow)
	time.Sleep(3 * time.Minute)
	ate.Traffic().Stop(t)
}

func TestQueryTelemetry(t *testing.T) {
	dut := ondatra.DUT(t, "dut")
	dp := dut.Port(t, "port1")
	ate := ondatra.ATE(t, "ate")
	ap := ate.Port(t, "port1")

	ds := gnmi.Get(t, dut, gnmi.OC().Interface(dp.Name()).OperStatus().State())
	as := gnmi.Get(t, ate, gnmi.OC().Interface(ap.Name()).OperStatus().State())
	if want := oc.Interface_OperStatus_UP; ds != want {
		t.Errorf("Get(DUT port1 status): got %v, want %v", ds, want)
	}
	if want := oc.Interface_OperStatus_UP; as != want {
		t.Errorf("Get(ATE port1 status): got %v, want %v", as, want)
	}

	p := gnmi.Collect(t, dut, gnmi.OC().Interface(`{{ port "port2" }}`).Counters().InPkts().State(), time.Minute)
	if got := netutil.MeanRate(t, p.Await(t)); got != expectedRate {
		t.Fatalf("Got unexpected input rate %.4f, want: %.4f", got, expectedRate)
	}
}

func TestExecuteOperation(t *testing.T) {
	dut := ondatra.DUT(t, "dut")
	dut.Operations().
		NewPing().
		WithDestination("8.8.8.8").
		Operate(t)

	ate := ondatra.ATE(t, "ate")
	ap := dut.Port(t, "port1")
	ate.Actions().
		NewSetPortState().
		WithPort(ap).
		WithEnabled(false).
		Send(t)
}

func doSomethingWith(stuff ...any) {}
