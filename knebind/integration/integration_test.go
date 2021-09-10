package integration_test

import (
	"testing"

	kinit "github.com/openconfig/ondatra/knebind/init"
	"github.com/openconfig/ondatra"
)

func TestMain(m *testing.M) {
	ondatra.RunTests(m, kinit.Init)
}

func TestGNMI(t *testing.T) {
	dut1 := ondatra.DUT(t, "dut1")
	dut2 := ondatra.DUT(t, "dut2")
	sys1 := dut1.Telemetry().System().GetFull(t)
	sys2 := dut2.Telemetry().System().GetFull(t)
	t.Logf("dut1 system: %v", sys1)
	t.Logf("dut2 system: %v", sys2)
}
