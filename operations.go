// Copyright 2020 Google LLC
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
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/openconfig/ondatra/closer"
	"github.com/openconfig/ondatra/internal/operations"
	"github.com/openconfig/ondatra/internal/reservation"
)

// Operations is the device operations API.
type Operations struct {
	dev reservation.Device
}

// NewInstall creates a new install operation.
func (o *Operations) NewInstall() *InstallOp {
	return &InstallOp{dev: o.dev}
}

// InstallOp is an OS install operation.
type InstallOp struct {
	dev     reservation.Device
	version string
	standby bool
	reader  io.Reader
}

func (i *InstallOp) String() string {
	return fmt.Sprintf("InstallOp%+v", *i)
}

// WithVersion specifies the version of the install operation.
func (i *InstallOp) WithVersion(version string) *InstallOp {
	i.version = version
	return i
}

// WithStandbySupervisor specifies whether the installation applies to the
// Standby Supervisor instead of the Active Supervisor.
func (i *InstallOp) WithStandbySupervisor(standby bool) *InstallOp {
	i.standby = standby
	return i
}

// WithPackageReader specifies the content of the OS package to be installed, in
// the form of a reader of its bytes. It may be omitted if the device already
// has the package. The specified function must return io.EOF when all the
// content has been read.
func (i *InstallOp) WithPackageReader(reader io.Reader) *InstallOp {
	i.reader = reader
	return i
}

// WithPackageFile specifies the content of the OS package to be installed, in
// the form of path to an image file. It may be omitted if the device already
// has the package.
func (i *InstallOp) WithPackageFile(path string) *InstallOp {
	return i.WithPackageReader(&fileReader{path: path})
}

type fileReader struct {
	path string
	file *os.File
}

func (fr *fileReader) Read(p []byte) (int, error) {
	if fr.file == nil {
		var err error
		fr.file, err = os.Open(fr.path)
		if err != nil {
			return 0, err
		}
	}
	n, err := fr.file.Read(p)
	if err == io.EOF {
		defer closer.CloseAndLog(fr.file.Close, "error closing package file")
	}
	return n, err
}

// Operate performs the Install operation.
func (i *InstallOp) Operate(t testing.TB) {
	t.Helper()
	logAction(t, "Installing package on %s", i.dev)
	if err := operations.Install(context.Background(), i.dev, i.version, i.standby, i.reader); err != nil {
		t.Fatalf("Operate(t) on %s: %v", i, err)
	}
}

// NewPing creates a new ping operation.
func (o *Operations) NewPing() *PingOp {
	return &PingOp{dev: o.dev}
}

// PingOp is a ping operation.
type PingOp struct {
	dev  reservation.Device
	dest string
}

func (p *PingOp) String() string {
	return fmt.Sprintf("PingOp%+v", *p)
}

// WithDestination specifies the destination address of the Ping operation.
func (p *PingOp) WithDestination(dest string) *PingOp {
	p.dest = dest
	return p
}

// Operate performs the Ping operation.
func (p *PingOp) Operate(t testing.TB) {
	t.Helper()
	logAction(t, "Pinging from %s", p.dev)
	if err := operations.Ping(context.Background(), p.dev, p.dest); err != nil {
		t.Fatalf("Operate(t) on %s: %v", p, err)
	}
}

// NewSetInterfaceState creates a new set interface state operation.
func (o *Operations) NewSetInterfaceState() *SetInterfaceStateOp {
	return &SetInterfaceStateOp{dev: o.dev}
}

// SetInterfaceStateOp is a set interface state operation that sets the state of an interface on a DUT.
type SetInterfaceStateOp struct {
	dev     reservation.Device
	intf    string
	enabled *bool
}

func (s *SetInterfaceStateOp) String() string {
	return fmt.Sprintf("SetInterfaceStateOp%+v", *s)
}

// WithPhysicalInterface specifies the target physcial interface of the set interface state operation.
// Only one of the logical interface and physical interface properties can be set at a time.
func (s *SetInterfaceStateOp) WithPhysicalInterface(port *Port) *SetInterfaceStateOp {
	s.intf = port.Name()
	return s
}

// WithLogicalInterface specifies the target logical interface of the set interface state operation.
// Only one of the logical interface and physical interface properties can be set at a time.
func (s *SetInterfaceStateOp) WithLogicalInterface(intf string) *SetInterfaceStateOp {
	s.intf = intf
	return s
}

// WithStateEnabled specifies that the target interface should be enabled.
func (s *SetInterfaceStateOp) WithStateEnabled(e bool) *SetInterfaceStateOp {
	s.enabled = &e
	return s
}

// Operate performs the set interface state operation.
func (s *SetInterfaceStateOp) Operate(t testing.TB) {
	t.Helper()
	logAction(t, "Setting interface state on %s", s.dev)
	if err := operations.SetInterfaceState(context.Background(), s.dev, s.intf, s.enabled); err != nil {
		t.Fatalf("Operate(t) on %s: %v", s, err)
	}
}

// NewReboot creates a new reboot operation.
func (o *Operations) NewReboot() *RebootOp {
	return &RebootOp{dev: o.dev}
}

// RebootOp is a reboot operation.
type RebootOp struct {
	dev     reservation.Device
	timeout time.Duration
}

func (r *RebootOp) String() string {
	return fmt.Sprintf("RebootOp%+v", *r)
}

// WithTimeout specifies the timeout on the reboot operation.
func (r *RebootOp) WithTimeout(timeout time.Duration) *RebootOp {
	r.timeout = timeout
	return r
}

// Operate performs the Reboot operation.
func (r *RebootOp) Operate(t testing.TB) {
	t.Helper()
	logAction(t, "Rebooting %s", r.dev)
	if err := operations.Reboot(context.Background(), r.dev, r.timeout); err != nil {
		t.Fatalf("Operate(t) on %s: %v", r, err)
	}
}

// NewRestartRouting creates a new restart routing operation.
func (o *Operations) NewRestartRouting() *RestartRoutingOp {
	return &RestartRoutingOp{o.dev}
}

// RestartRoutingOp is an operation that restarts the routing process on a device
type RestartRoutingOp struct {
	dev reservation.Device
}

func (r *RestartRoutingOp) String() string {
	return fmt.Sprintf("RestartRoutingOp%+v", *r)
}

// Operate performs the RestartRouting operation.
func (r *RestartRoutingOp) Operate(t testing.TB) {
	t.Helper()
	logAction(t, "Restarting routing on %s", r.dev)
	if err := operations.RestartRouting(context.Background(), r.dev); err != nil {
		t.Fatalf("Operate(t) on %s: %v", r, err)
	}
}
