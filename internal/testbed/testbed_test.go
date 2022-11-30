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

package testbed_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/ondatra/internal/flags"
	"github.com/openconfig/ondatra/internal/testbed"

	opb "github.com/openconfig/ondatra/proto"
)

func TestReserve(t *testing.T) {
	tbPath := writeTestbedFile(t, t.TempDir(), `duts{id:"dut"}`)
	tbFlags := &flags.Values{TestbedPath: tbPath}
	tbAndIDFlags := &flags.Values{TestbedPath: tbPath, ResvID: "1234"}
	wantRes := &binding.Reservation{DUTs: map[string]binding.DUT{
		"dut": &binding.AbstractDUT{&binding.Dims{Name: "myDUT"}},
	}}

	t.Run("success", func(t *testing.T) {
		bind := fakebind.Setup()
		bind.ReserveFn = func(context.Context, *opb.Testbed, time.Duration, time.Duration, map[string]string) (*binding.Reservation, error) {
			return wantRes, nil
		}
		if err := testbed.Reserve(context.Background(), tbFlags); err != nil {
			t.Fatalf("Reserve() got unexpected error: %v", err)
		}
		gotRes, err := testbed.Reservation()
		if err != nil {
			t.Fatalf("Reservation() got unexpected error: %v", err)
		}
		if gotRes != wantRes {
			t.Errorf("Reserve() got %v, want %v", gotRes, wantRes)
		}
	})

	t.Run("error", func(t *testing.T) {
		wantErr := "reserve error"
		bind := fakebind.Setup()
		bind.ReserveFn = func(context.Context, *opb.Testbed, time.Duration, time.Duration, map[string]string) (*binding.Reservation, error) {
			return nil, fmt.Errorf(wantErr)
		}
		gotErr := testbed.Reserve(context.Background(), tbFlags)
		if gotErr == nil || !strings.Contains(gotErr.Error(), wantErr) {
			t.Errorf("Reserve() got error %v, want %q", gotErr, wantErr)
		}
	})

	t.Run("double reserve fails", func(t *testing.T) {
		bind := fakebind.Setup()
		bind.ReserveFn = func(context.Context, *opb.Testbed, time.Duration, time.Duration, map[string]string) (*binding.Reservation, error) {
			return wantRes, nil
		}
		if err := testbed.Reserve(context.Background(), tbFlags); err != nil {
			t.Fatalf("Reserve() got unexpected error on first call: %v", err)
		}
		if err := testbed.Reserve(context.Background(), tbFlags); err == nil {
			t.Errorf("Reserve() got success on second call, but want error")
		}
	})

	t.Run("fetch success", func(t *testing.T) {
		bind := fakebind.Setup()
		bind.FetchReservationFn = func(context.Context, string) (*binding.Reservation, error) {
			return wantRes, nil
		}
		if err := testbed.Reserve(context.Background(), tbAndIDFlags); err != nil {
			t.Error(err)
		}
		gotRes, err := testbed.Reservation()
		if err != nil {
			t.Fatalf("Reservation() got unexpected error: %v", err)
		}
		if gotRes != wantRes {
			t.Fatalf("Reserve() got %v, want %v", gotRes, wantRes)
		}
	})

	t.Run("fetch error", func(t *testing.T) {
		wantErr := "fetch error"
		bind := fakebind.Setup()
		bind.FetchReservationFn = func(context.Context, string) (*binding.Reservation, error) {
			return nil, fmt.Errorf(wantErr)
		}
		gotErr := testbed.Reserve(context.Background(), tbAndIDFlags)
		if gotErr == nil || !strings.Contains(gotErr.Error(), wantErr) {
			t.Errorf("Reserve() got error %v, want %q", gotErr, wantErr)
		}
	})
}

func TestReserveBadTestbeds(t *testing.T) {
	bind := fakebind.Setup()
	tmpDir := t.TempDir()

	tests := []struct {
		name, tbProto string
		res           *binding.Reservation
		wantErr       string
	}{{
		name:    "Bad testbed proto",
		tbProto: "gaga",
		wantErr: "proto",
	}, {
		name:    "Nonexistent device ID",
		tbProto: `duts{id:"gaga"}`,
		res:     &binding.Reservation{DUTs: map[string]binding.DUT{"dut": &binding.AbstractDUT{&binding.Dims{}}}},
		wantErr: "gaga",
	}, {
		name:    "Nonexistent port ID",
		tbProto: `duts{id:"dut" ports{id:"gaga"}}`,
		res: &binding.Reservation{DUTs: map[string]binding.DUT{"dut": &binding.AbstractDUT{
			&binding.Dims{Name: "d1", Vendor: opb.Device_ARISTA, HardwareModel: "m", SoftwareVersion: "v"},
		}}},
		wantErr: "gaga",
	}, {
		name:    "Disallowed device ID",
		tbProto: `duts{id:"$^&#(#"}`,
		res:     &binding.Reservation{DUTs: map[string]binding.DUT{"dut": &binding.AbstractDUT{&binding.Dims{}}}},
		wantErr: "invalid testbed ID",
	}, {
		name:    "Disallowed port ID",
		tbProto: `duts{id:"dut" ports{id:"$^&#(#"}}`,
		res:     &binding.Reservation{DUTs: map[string]binding.DUT{"dut": &binding.AbstractDUT{&binding.Dims{}}}},
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
		res: &binding.Reservation{DUTs: map[string]binding.DUT{"dut": &binding.AbstractDUT{
			&binding.Dims{Vendor: opb.Device_ARISTA, HardwareModel: "m", SoftwareVersion: "v"},
		}}},
		wantErr: "no name",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bind.ReserveFn = func(context.Context, *opb.Testbed, time.Duration, time.Duration, map[string]string) (*binding.Reservation, error) {
				return test.res, nil
			}
			tbPath := writeTestbedFile(t, tmpDir, test.tbProto)
			gotErr := testbed.Reserve(context.Background(), &flags.Values{TestbedPath: tbPath})
			if gotErr == nil || !strings.Contains(gotErr.Error(), test.wantErr) {
				t.Errorf("Reserve() got error %v, want %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestRelease(t *testing.T) {
	t.Run("not reserved", func(t *testing.T) {
		bind := fakebind.Setup()
		bind.ReleaseFn = func(context.Context) error {
			return nil
		}
		if err := testbed.Release(context.Background()); err != nil {
			t.Errorf("Release() got unexpected error: %v", err)
		}
	})

	t.Run("reserved", func(t *testing.T) {
		bind := fakebind.Setup().WithReservation(new(binding.Reservation))
		bind.ReleaseFn = func(context.Context) error {
			return nil
		}
		if err := testbed.Release(context.Background()); err != nil {
			t.Errorf("Release() got unexpected error: %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		wantErr := "release error"
		bind := fakebind.Setup().WithReservation(new(binding.Reservation))
		bind.ReleaseFn = func(context.Context) error {
			return fmt.Errorf(wantErr)
		}
		gotErr := testbed.Release(context.Background())
		if gotErr == nil || !strings.Contains(gotErr.Error(), wantErr) {
			t.Errorf("Release() got error %v, want %q", gotErr, wantErr)
		}
	})
}

func writeTestbedFile(t *testing.T, dir, text string) string {
	t.Helper()
	f, err := os.CreateTemp(dir, "*.textproto")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Failed to close temp file: %v", err)
		}
	}()
	if _, err := f.WriteString(text); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	return f.Name()
}
