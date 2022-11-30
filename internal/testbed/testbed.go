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

// Package testbed implements Testing API for a testbed.
package testbed

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"sync"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/flags"
	"google.golang.org/protobuf/encoding/prototext"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	resMu   sync.RWMutex
	res     *binding.Reservation
	fetched bool
	bind    binding.Binding
)

// SetBinding sets the Ondatra binding.
func SetBinding(b binding.Binding) {
	bind = b
}

// SetReservationForTesting sets Ondatra to a state in which the specified
// reservation has been reserved. A nil reservation argument sets Ondatra to
// an unreserved state. This is only called by fakebind for testing purposes.
func SetReservationForTesting(r *binding.Reservation) {
	res = r
	fetched = false
}

// Reservation returns the current reservation.
func Reservation() (*binding.Reservation, error) {
	resMu.RLock()
	defer resMu.RUnlock()
	if res == nil {
		return nil, errors.New("testbed is not reserved; Did you forget to call ondatra.RunTests in TestMain?")
	}
	return res, nil
}

// Reserve reserves the testbed.
func Reserve(ctx context.Context, fv *flags.Values) error {
	resMu.Lock()
	defer resMu.Unlock()
	if res != nil {
		return errors.New("testbed is already reserved; Did you call ondatra.RunTests multiple times?")
	}
	tb := &opb.Testbed{}
	s, err := os.ReadFile(fv.TestbedPath)
	if err != nil {
		return fmt.Errorf("failed to read testbed proto %s: %w", fv.TestbedPath, err)
	}
	if err := prototext.Unmarshal(s, tb); err != nil {
		return fmt.Errorf("failed to parse testbed proto %s: %w", fv.TestbedPath, err)
	}
	if err := validateTB(tb); err != nil {
		return err
	}

	var r *binding.Reservation
	if fv.ResvID == "" {
		r, err = bind.Reserve(ctx, tb, fv.RunTime, fv.WaitTime, fv.ResvPartial)
	} else {
		r, err = bind.FetchReservation(ctx, fv.ResvID)
		fetched = true
	}
	if err != nil {
		return err
	}
	if err := validateRes(tb, r); err != nil {
		return err
	}
	res = r
	return nil
}

// portMap registers which ports are connected to which other ports, in the format "<device-id>:<port-id>".
// Non-connected ports map to "", which allows to check for validity of port IDs in links.
// Each pair of connected ports A and B must be in the map twice: port A's ID mapping to port B's ID and port B's ID mapping to port A's ID.
type portMap map[string]string

func validateTB(tb *opb.Testbed) error {
	pm := make(portMap)
	for _, d := range append(tb.GetDuts(), tb.GetAtes()...) {
		if err := checkID(d.GetId()); err != nil {
			return err
		}
		for _, p := range d.GetPorts() {
			if err := checkID(p.GetId()); err != nil {
				return err
			}
			pid := fmt.Sprintf("%s:%s", d.GetId(), p.GetId())
			if _, ok := pm[pid]; ok {
				return fmt.Errorf("duplicate port %q", pid)
			}
			pm[pid] = ""
		}
	}
	for _, ln := range tb.GetLinks() {
		dupB, ok := pm[ln.GetA()]
		if !ok {
			return fmt.Errorf("nonexistent linked port ID %q", ln.GetA())
		}
		dupA, ok := pm[ln.GetB()]
		if !ok {
			return fmt.Errorf("nonexistent linked port ID %q", ln.GetB())
		}
		if dupB != "" {
			return fmt.Errorf("conflicting connections from %q to %q and %q", ln.GetA(), dupB, ln.GetB())
		}
		if dupA != "" {
			return fmt.Errorf("conflicting connections from %q to %q and %q", ln.GetB(), dupA, ln.GetA())
		}
		pm[ln.GetA()] = ln.GetB()
		pm[ln.GetB()] = ln.GetA()
	}
	return nil
}

var idRE = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)

// checkID enforces testbed IDs that "look like" variable names.
// It ensures consistency and helps avoid the user mistakenly using a device name.
func checkID(id string) error {
	if !idRE.MatchString(id) {
		return fmt.Errorf("invalid testbed ID %q: must start with a letter and contain only letters, numbers, or underscore", id)
	}
	return nil
}

func validateRes(tb *opb.Testbed, res *binding.Reservation) error {
	for _, dut := range tb.GetDuts() {
		rd, err := DUT(res, dut.GetId())
		if err != nil {
			return err
		}
		if err := validateDevice(dut, rd); err != nil {
			return err
		}
	}
	for _, ate := range tb.GetAtes() {
		ra, err := ATE(res, ate.GetId())
		if err != nil {
			return err
		}
		if err := validateDevice(ate, ra); err != nil {
			return err
		}
	}
	return nil
}

func validateDevice(dev *opb.Device, rd binding.Device) error {
	if rd.Name() == "" {
		return fmt.Errorf("no name for reserved device: %v", rd)
	}
	for _, p := range dev.GetPorts() {
		rp, err := Port(rd, p.GetId())
		if err != nil {
			return err
		}
		if rp.Name == "" {
			return fmt.Errorf("no name for reserved port: %v", rp)
		}
	}
	return nil
}

// Release releases the testbed. This is a noop if the reservation is not
// currently reserved or if the reservation was fetched and not created.
func Release(ctx context.Context) error {
	resMu.Lock()
	defer resMu.Unlock()
	if res == nil || fetched {
		return nil
	}
	res = nil
	return bind.Release(ctx)
}

// Device returns the Device in the specified reservation with the specified ID.
func Device(res *binding.Reservation, id string) (binding.Device, error) {
	if d, err := DUT(res, id); err == nil { // if NO error
		return d, nil
	}
	if a, err := ATE(res, id); err == nil { // if NO error
		return a, nil
	}
	return nil, fmt.Errorf("device ID %q not found in the reservation", id)
}

// DUT returns the DUT in the specified reservation with the specified ID.
func DUT(res *binding.Reservation, id string) (binding.DUT, error) {
	if d, ok := res.DUTs[id]; ok {
		return d, nil
	}
	return nil, fmt.Errorf("DUT ID %q not found in the reservation", id)
}

// ATE returns the ATE in the specified reservation with the specified ID.
func ATE(res *binding.Reservation, id string) (binding.ATE, error) {
	if a, ok := res.ATEs[id]; ok {
		return a, nil
	}
	return nil, fmt.Errorf("ATE ID %q not found in the reservation", id)
}

// Port returns the port in the specified dimensions with the specified ID.
func Port(rd binding.Device, id string) (*binding.Port, error) {
	if p, ok := rd.Ports()[id]; ok {
		return p, nil
	}
	return nil, fmt.Errorf("port ID %q not found in reserved device %q", id, rd.Name())
}
