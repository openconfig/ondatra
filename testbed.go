// Copyright 2019 Google LLC
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

package ondatra

import (
	"golang.org/x/net/context"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/testbed"
)

func reserve(testbedPath string, runTime, waitTime time.Duration) error {
	return testbed.Reserve(context.Background(), testbedPath, runTime, waitTime)
}

func release() error {
	return testbed.Release(context.Background())
}

func checkRes(t testing.TB) *reservation.Reservation {
	t.Helper()
	res, err := testbed.Reservation()
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// DUT returns the DUT in the testbed with a given id.
func DUT(t testing.TB, id string) *DUTDevice {
	t.Helper()
	rd, err := checkRes(t).DUT(id)
	if err != nil {
		t.Fatalf("DUT(t, %s): %v", id, err)
	}
	return newDUT(id, rd)
}

// DUTs returns a map of DUT id to DUT in the testbed.
func DUTs(t testing.TB) map[string]*DUTDevice {
	t.Helper()
	rm := checkRes(t).DUTs
	m := make(map[string]*DUTDevice)
	for id, rd := range rm {
		m[id] = newDUT(id, rd)
	}
	return m
}

func newDUT(id string, res *reservation.DUT) *DUTDevice {
	return &DUTDevice{&Device{id: id, res: res}}
}

// ATE returns the ATE in the testbed with a given id.
func ATE(t testing.TB, id string) *ATEDevice {
	t.Helper()
	ra, err := checkRes(t).ATE(id)
	if err != nil {
		t.Fatalf("ATE(t, %s): %v", id, err)
	}
	return newATE(id, ra)
}

// ATEs returns a map of ATE id to ATE in the testbed.
func ATEs(t testing.TB) map[string]*ATEDevice {
	t.Helper()
	rm := checkRes(t).ATEs
	m := make(map[string]*ATEDevice)
	for id, ra := range rm {
		m[id] = newATE(id, ra)
	}
	return m
}

func newATE(id string, res *reservation.ATE) *ATEDevice {
	return &ATEDevice{&Device{id: id, res: res}}
}

// ReservationID returns the reservation ID for the testbed.
func ReservationID(t testing.TB) string {
	t.Helper()
	return checkRes(t).ID
}
