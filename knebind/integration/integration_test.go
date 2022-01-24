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

package integration_test

import (
	"context"
	"github.com/openconfig/gribigo/chk"
	"github.com/openconfig/gribigo/fluent"
	"github.com/openconfig/gribigo/server"
	"github.com/openconfig/ondatra"
	kinit "github.com/openconfig/ondatra/knebind/init"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	ondatra.RunTests(m, kinit.Init)
}

func TestGNMI(t *testing.T) {
	dut1 := ondatra.DUT(t, "dut1")
	dut2 := ondatra.DUT(t, "dut2")
	sys1 := dut1.Telemetry().System().Lookup(t)
	sys2 := dut2.Telemetry().System().Lookup(t)
	t.Logf("dut1 system: %v", sys1)
	t.Logf("dut2 system: %v", sys2)
}

func TestGRIBISample(t *testing.T) {
	dut1 := ondatra.DUT(t, "dut1")
	dut1.GRIBI().MasterClient(t)
	defer dut1.GRIBI().CloseMasterClient(t)
	dut1.GRIBI().SynchElectionID(t, true)
	dut1.GRIBI().IncElectionID(t)
	elecLow, elecHigh := dut1.GRIBI().GetElectionID(t)
	dut1.GRIBI().MasterClient(t).Connection().WithInitialElectionID(elecLow, elecHigh).WithRedundancyMode(fluent.ElectedPrimaryClient).WithPersistence()
	dut1.GRIBI().MasterClient(t).Start(context.Background(), t)
	defer dut1.GRIBI().MasterClient(t).Stop(t)
	dut1.GRIBI().MasterClient(t).StartSending(context.Background(), t)

	subctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	dut1.GRIBI().MasterClient(t).Await(subctx, t)

	// check sessoin param result
	chk.HasResult(t, dut1.GRIBI().MasterClient(t).Results(t),
		fluent.OperationResult().
			WithSuccessfulSessionParams().
			AsResult(),
	)

	// check election id result
	chk.HasResult(t, dut1.GRIBI().MasterClient(t).Results(t),
		fluent.OperationResult().
			WithCurrentServerElectionID(elecLow, elecHigh).
			AsResult(),
	)

	// Send modify request with lower election id
	dut1.GRIBI().MasterClient(t).Modify().AddEntry(t, fluent.NextHopEntry().WithNetworkInstance(server.DefaultNetworkInstanceName).WithIndex(1).WithIPAddress("192.0.2.2").WithElectionID(elecLow+1, elecHigh))

	dut1.GRIBI().MasterClient(t).Await(subctx, t)

	// check AFT result, expecting operation fails
	chk.HasResult(t, dut1.GRIBI().MasterClient(t).Results(t), fluent.OperationResult().
		WithOperationID(1).
		WithProgrammingResult(fluent.ProgrammingFailed).
		AsResult(),
	)

}
