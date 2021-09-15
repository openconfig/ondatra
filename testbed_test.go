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
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/negtest"

	opb "github.com/openconfig/ondatra/proto"
)

func TestReserveErrors(t *testing.T) {
	initFakeBinding(t)

	for _, tt := range []struct {
		name, tbProto string
		res           *reservation.Reservation
		wantErr       string
	}{{
		name:    "Bad testbed proto",
		tbProto: "gaga",
		wantErr: "proto",
	}, {
		name:    "Nonexistent device ID",
		tbProto: `duts{id:"gaga"}`,
		res:     &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{&reservation.Dims{}}}},
		wantErr: "gaga",
	}, {
		name:    "Nonexistent port ID",
		tbProto: `duts{id:"dut" ports{id:"gaga"}}`,
		res: &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{
			&reservation.Dims{Name: "d1", Vendor: opb.Device_ARISTA, HardwareModel: "m", SoftwareVersion: "v"},
		}}},
		wantErr: "gaga",
	}, {
		name:    "Disallowed device ID",
		tbProto: `duts{id:"$^&#(#"}`,
		res:     &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{&reservation.Dims{}}}},
		wantErr: "invalid testbed ID",
	}, {
		name:    "Disallowed port ID",
		tbProto: `duts{id:"dut" ports{id:"$^&#(#"}}`,
		res:     &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{&reservation.Dims{}}}},
		wantErr: "invalid testbed ID",
	}, {
		name:    "Duplicate port ID",
		tbProto: `duts{id:"dut" ports{id:"port1"} ports{id:"port1"}}`,
		wantErr: "duplicate",
	}, {
		name:    "Link from bad device ID",
		tbProto: `duts{id:"dut" ports{id:"port1"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"}} links{a:"gaga:port1", b:"ate1:port1"}`,
		wantErr: "port ID",
	}, {
		name:    "Link to bad device ID",
		tbProto: `duts{id:"dut" ports{id:"port1"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"}} links{a:"dut:port1", b:"gaga:port1"}`,
		wantErr: "port ID",
	}, {
		name:    "Link from bad port ID",
		tbProto: `duts{id:"dut" ports{id:"port1"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"}} links{a:"dut:gaga", b:"ate1:port1"}`,
		wantErr: "port ID",
	}, {
		name:    "Link to bad port ID",
		tbProto: `duts{id:"dut" ports{id:"port1"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"}} links{a:"dut:port1", b:"ate1:gaga"}`,
		wantErr: "port ID",
	}, {
		name:    "Two links from same port on DUT",
		tbProto: `duts{id:"dut" ports{id:"port1"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"} ports:{id:"port2"}} links{a:"dut:port1", b:"ate1:port1"} links{a:"dut:port1", b:"ate1:port2"}`,
		wantErr: `from "dut:port1" to "ate1:port1" and "ate1:port2"`,
	}, {
		name:    "Two links from same port on ATE",
		tbProto: `duts{id:"dut" ports{id:"port1"} ports:{id:"port2"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"}} links{a:"ate1:port1", b:"dut:port1"} links{a:"ate1:port1", b:"dut:port2"}`,
		wantErr: `from "ate1:port1" to "dut:port1" and "dut:port2"`,
	}, {
		name:    "Two links to same port on DUT",
		tbProto: `duts{id:"dut" ports{id:"port1"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"} ports:{id:"port2"}} links{a:"ate1:port1", b:"dut:port1"} links{a:"ate1:port2", b:"dut:port1"}`,
		wantErr: `from "dut:port1" to "ate1:port1" and "ate1:port2"`,
	}, {
		name:    "Two links to same port on ATE",
		tbProto: `duts{id:"dut" ports{id:"port1"} ports:{id:"port2"}} ates{vendor:IXIA, id:"ate1" ports:{id:"port1"}} links{a:"dut:port1", b:"ate1:port1"} links{a:"dut:port2", b:"ate1:port1"}`,
		wantErr: `from "ate1:port1" to "dut:port1" and "dut:port2"`,
	}, {
		name:    "No device name",
		tbProto: `duts{id:"dut"}`,
		res: &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{
			&reservation.Dims{Vendor: opb.Device_ARISTA, HardwareModel: "m", SoftwareVersion: "v"},
		}}},
		wantErr: "no name",
	}, {
		name:    "No device vendor",
		tbProto: `duts{id:"dut"}`,
		res: &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{&reservation.Dims{
			Name: "d1", HardwareModel: "m", SoftwareVersion: "v"},
		}}},
		wantErr: "no vendor",
	}, {
		name:    "No hardware model",
		tbProto: `duts{id:"dut"}`,
		res: &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{&reservation.Dims{
			Name: "d1", Vendor: opb.Device_ARISTA, SoftwareVersion: "v"},
		}}},
		wantErr: "no hardware model",
	}, {
		name:    "No software version",
		tbProto: `duts{id:"dut"}`,
		res: &reservation.Reservation{DUTs: map[string]*reservation.DUT{"dut": &reservation.DUT{&reservation.Dims{
			Name: "d1", Vendor: opb.Device_ARISTA, HardwareModel: "m"},
		}}},
		wantErr: "no software version",
	}} {
		t.Run(tt.name, func(t *testing.T) {
			fakeBind.Reservation = tt.res
			err := reserve(writeTemp(t, tt.tbProto), time.Hour, 0)
			if err == nil {
				release()
				t.Fatalf("Reserve unexpectedly succeeded, must fail")
			}
			t.Logf("Testbed reservation error: %v", err)
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("Reserve error message was '%s', must contain '%s'", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestDoubleReserveFails(t *testing.T) {
	initFakeBinding(t)
	if err := reserve(fakeTBPath, time.Hour, 0); err != nil {
		t.Fatalf("First Reserve() call failed: %v", err)
	}
	if err := reserve(fakeTBPath, time.Hour, 0); err == nil {
		t.Errorf("Second Reserve() call succeeded, but must fail")
	}
	if err := release(); err != nil {
		t.Errorf("Release() failed: %v", err)
	}
}

func TestReserve(t *testing.T) {
	initFakeBinding(t)
	if err := reserve(fakeTBPath, time.Hour, 0); err != nil {
		t.Fatal(err)
	}

	t.Run("Get DUT", func(t *testing.T) {
		id := "dut"
		d := DUT(t, id)
		if got, want := d.ID(), id; got != want {
			t.Errorf("DUT id = %q, want %q", got, want)
		}
		if got, want := d.Name(), "pf01.xxx01"; got != want {
			t.Errorf("DUT name = %q, want %q", got, want)
		}
		if got, want := d.Vendor(), ARISTA; got != want {
			t.Errorf("DUT vendor = %v, want %v", got, want)
		}
		if got, want := d.Model(), "aristaModel"; got != want {
			t.Errorf("DUT model = %v, want %v", got, want)
		}
		if got, want := d.Version(), "aristaVersion"; got != want {
			t.Errorf("DUT version = %v, want %v", got, want)
		}
	})

	t.Run("Get DUTs", func(t *testing.T) {
		id := "dut"
		duts := DUTs(t)
		d, ok := duts[id]
		if !ok {
			t.Errorf("DUT id %q not present in DUTs map %q, must be present", id, duts)
		}
		if got, want := d.ID(), id; got != want {
			t.Errorf("DUT id = %q, want %q", got, want)
		}
	})

	t.Run("Get DUT failure", func(t *testing.T) {
		id := "gaga"
		got := negtest.ExpectFatal(t, func(t testing.TB) {
			DUT(t, id)
		})
		if !strings.Contains(got, id) {
			t.Errorf("DUT(%q) failed with message %q, want %q", id, got, id)
		}
	})

	t.Run("Get ATE", func(t *testing.T) {
		id := "ate"
		a := ATE(t, id)
		if got, want := a.ID(), id; got != want {
			t.Errorf("ATE id = %q, want %q", got, want)
		}
		if got, want := a.Name(), "ix1"; got != want {
			t.Errorf("ATE name = %q, want %q", got, want)
		}
	})

	t.Run("Get ATEs", func(t *testing.T) {
		id := "ate"
		ates := ATEs(t)
		a, ok := ates[id]
		if !ok {
			t.Errorf("ATE id %q not present in ATEs map %q, must be present", id, ates)
		}
		if got, want := a.ID(), id; got != want {
			t.Errorf("ATE id = %q, want %q", got, want)
		}
	})

	t.Run("Get ATE failure", func(t *testing.T) {
		id := "gaga"
		got := negtest.ExpectFatal(t, func(t testing.TB) {
			ATE(t, id)
		})
		if !strings.Contains(got, id) {
			t.Errorf("ATE(%q) failed with message %q, want %q", id, got, id)
		}
	})

	t.Run("Get Port", func(t *testing.T) {
		did, pid := "dut", "port1"
		p := DUT(t, did).Port(t, pid)
		if got, want := p.ID(), pid; got != want {
			t.Errorf("port id = %q, want %q", got, want)
		}
		if got, want := p.Device().ID(), did; got != want {
			t.Errorf("port device id = %q, want %q", got, want)
		}
		// TODO: Restore when speed is consistently propagated back.
		// if got, want := p.Speed(), Speed10Gb; got != want {
		//	t.Errorf("port speed = %d, want %d", got, want)
		// }
	})

	t.Run("Get Port failure", func(t *testing.T) {
		d := DUT(t, "dut")
		pid := "gaga"
		got := negtest.ExpectFatal(t, func(t testing.TB) {
			d.Port(t, pid)
		})
		if !strings.Contains(got, pid) {
			t.Errorf("Port(%q) failed with message %q, want %q", pid, got, pid)
		}
	})
}

func writeTemp(t *testing.T, c string) string {
	t.Helper()
	ttd := os.Getenv("TEST_TMPDIR")
	if ttd == "" {
		ttd = os.TempDir()
	}
	f, err := ioutil.TempFile(ttd, "*.pb.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err := f.WriteString(c); err != nil {
		f.Close()
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}
	return f.Name()
}
