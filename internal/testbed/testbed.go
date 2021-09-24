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
	"golang.org/x/net/context"
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/prototext"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/usererr"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	resMu sync.RWMutex
	res   *reservation.Reservation
)

// Reservation returns the current reservation.
func Reservation() (*reservation.Reservation, error) {
	resMu.RLock()
	defer resMu.RUnlock()
	if res == nil {
		return nil, errors.New("testbed is not reserved; RunTests was not called")
	}
	return res, nil
}

// Reserve reserves a testbed, typically reading its definition for the test.
func Reserve(ctx context.Context, testbedPath string, runTime, waitTime time.Duration) error {
	if testbedPath == "" {
		return errors.New("testbed path not specified")
	}
	if runTime < 0 {
		return errors.Errorf("run timeout is negative: %d", runTime)
	}
	if waitTime < 0 {
		return errors.Errorf("wait timeout is negative: %d", waitTime)
	}
	resMu.Lock()
	defer resMu.Unlock()
	if res != nil {
		return errors.New("testbed is already reserved; RunTests was already called")
	}
	tb := &opb.Testbed{}
	s, err := ioutil.ReadFile(testbedPath)
	if err != nil {
		return errors.Wrapf(err, "failed to read testbed proto %s", testbedPath)
	}
	err = prototext.Unmarshal(s, tb)
	if err != nil {
		return errors.Wrapf(err, "failed to parse testbed proto %s", testbedPath)
	}
	if err := validateTB(tb); err != nil {
		return err
	}
	r, err := binding.Get().Reserve(ctx, tb, runTime, waitTime)
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
				return usererr.New("duplicate port %q", pid)
			}
			pm[pid] = ""
		}
	}
	for _, ln := range tb.GetLinks() {
		dupB, ok := pm[ln.GetA()]
		if !ok {
			return usererr.New("nonexistent linked port ID %q", ln.GetA())
		}
		dupA, ok := pm[ln.GetB()]
		if !ok {
			return usererr.New("nonexistent linked port ID %q", ln.GetB())
		}
		if dupB != "" {
			return usererr.New("conflicting connections from %q to %q and %q", ln.GetA(), dupB, ln.GetB())
		}
		if dupA != "" {
			return usererr.New("conflicting connections from %q to %q and %q", ln.GetB(), dupA, ln.GetA())
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
		return usererr.New("invalid testbed ID %q: must start with a letter and contain only letters, numbers, or underscore", id)
	}
	return nil
}

func validateRes(tb *opb.Testbed, res *reservation.Reservation) error {
	for _, dut := range tb.GetDuts() {
		rd, err := res.DUT(dut.GetId())
		if err != nil {
			return err
		}
		if err := validateDevice(dut, rd); err != nil {
			return err
		}
	}
	for _, ate := range tb.GetAtes() {
		ra, err := res.ATE(ate.GetId())
		if err != nil {
			return err
		}
		if err := validateDevice(ate, ra); err != nil {
			return err
		}
	}
	return nil
}

func validateDevice(dev *opb.Device, rd reservation.Device) error {
	dims := rd.Dimensions()
	if dims.Name == "" {
		return errors.Errorf("no name for reserved device: %v", rd)
	}
	if dims.Vendor == opb.Device_UNKNOWN {
		return errors.Errorf("no vendor for reserved device: %v", rd)
	}
	if dims.HardwareModel == "" {
		return errors.Errorf("no hardware model for reserved device: %v", rd)
	}
	if dims.SoftwareVersion == "" {
		return errors.Errorf("no software version for reserved device: %v", rd)
	}
	for _, p := range dev.GetPorts() {
		rp, err := dims.Port(p.GetId())
		if err != nil {
			return err
		}
		if rp.Name == "" {
			return errors.Errorf("no name for reserved port: %v", rp)
		}
	}
	return nil
}

// Release releases the testbed.
func Release(ctx context.Context) error {
	resMu.Lock()
	defer resMu.Unlock()
	if res == nil {
		return nil
	}
	res = nil
	return binding.Get().Release(ctx)
}
