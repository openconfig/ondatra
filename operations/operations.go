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

package operations

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	closer "github.com/openconfig/gocloser"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/events"
	"github.com/openconfig/ondatra/internal/rawapis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	frpb "github.com/openconfig/gnoi/factory_reset"
	ospb "github.com/openconfig/gnoi/os"
	spb "github.com/openconfig/gnoi/system"
	tpb "github.com/openconfig/gnoi/types"
)

// New constructs a new instance of the operations API.
// Tests must not call this directly
func New(dut binding.DUT) *Operations {
	return &Operations{dut: dut}
}

// Operations is the device operations API.
type Operations struct {
	dut binding.DUT
}

// NewInstall creates a new install operation.
func (o *Operations) NewInstall() *InstallOp {
	return &InstallOp{dut: o.dut}
}

// InstallOp is an OS install operation.
type InstallOp struct {
	dut     binding.DUT
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
	t = events.ActionStarted(t, "Installing package on %s", i.dut)
	if err := install(context.Background(), i.dut, i.version, i.standby, i.reader); err != nil {
		t.Fatalf("Operate(t) on %s: %v", i, err)
	}
}

// The gNOI install scenarios are documented on the Install function here:
// https://github.com/openconfig/gnoi/blob/master/os/os.proto
func install(ctx context.Context, dut binding.DUT, version string, standby bool, reader io.Reader) error {
	if version == "" {
		return fmt.Errorf("version not set in install operation on DUT: %v", dut)
	}
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	ic, err := gnoi.OS().Install(ctx)
	if err != nil {
		return fmt.Errorf("error creating gnoi install client: %w", err)
	}
	sreq := &ospb.InstallRequest{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{
		Version:           version,
		StandbySupervisor: standby,
	}}}
	if err := ic.Send(sreq); err != nil {
		return fmt.Errorf("error sending gnoi install request: %w", err)
	}
	validated, err := awaitPackageInstall(ic)
	if err != nil {
		return err
	}
	if validated == nil { // if content is needed
		if reader == nil {
			return fmt.Errorf("DUT %v does not have package and no package specified in install operation", dut)
		}
		awaitChan := make(chan error)
		go func() {
			validated, err = awaitPackageInstall(ic)
			awaitChan <- err
		}()
		if err := transferContent(ic, reader); err != nil {
			return err
		}
		if err := <-awaitChan; err != nil {
			return err
		}
	}
	if gotVersion := validated.GetVersion(); gotVersion != version {
		return fmt.Errorf("installed version %q does not match requested version %q", gotVersion, version)
	}
	return nil
}

func transferContent(ic ospb.OS_InstallClient, reader io.Reader) error {
	// The gNOI SetPackage operation sets the maximum chunk size at 64K,
	// so assuming the install operation allows for up to the same size.
	buf := make([]byte, 64*1024)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			if err := ic.Send(&ospb.InstallRequest{Request: &ospb.InstallRequest_TransferContent{buf[0:n]}}); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return ic.Send(&ospb.InstallRequest{Request: &ospb.InstallRequest_TransferEnd{&ospb.TransferEnd{}}})
}

// awaitPackageInstall receives messages from the client until learns that either
// (a) the package is installed and validated, in which case it returns the validated message
// (b) the device does not have the package, in which case it returns a nil validated message
// (c) an error occurs, in which case it returns the error
func awaitPackageInstall(ic ospb.OS_InstallClient) (*ospb.Validated, error) {
	for {
		cresp, err := ic.Recv()
		if err != nil {
			return nil, err
		}
		switch v := cresp.GetResponse().(type) {
		case *ospb.InstallResponse_Validated:
			return v.Validated, nil
		case *ospb.InstallResponse_TransferReady:
			return nil, nil
		case *ospb.InstallResponse_InstallError:
			errName := ospb.InstallError_Type_name[int32(v.InstallError.Type)]
			return nil, fmt.Errorf("installation error %q: %s", errName, v.InstallError.GetDetail())
		case *ospb.InstallResponse_TransferProgress:
			log.Infof("installation progress: %v bytes received from client", v.TransferProgress.GetBytesReceived())
		case *ospb.InstallResponse_SyncProgress:
			log.Infof("installation progress: %v%% synced from supervisor", v.SyncProgress.GetPercentageTransferred())
		default:
			return nil, fmt.Errorf("unexpected client install response: %v (%T)", v, v)
		}
	}
}

// NewPing creates a new ping operation.
func (o *Operations) NewPing() *PingOp {
	return &PingOp{dut: o.dut}
}

// PingOp is a ping operation.
type PingOp struct {
	dut        binding.DUT
	dest       string
	count      int32
	packetSize int32
}

func (p *PingOp) String() string {
	return fmt.Sprintf("PingOp%+v", *p)
}

// WithDestination specifies the destination address of the Ping operation.
func (p *PingOp) WithDestination(dest string) *PingOp {
	p.dest = dest
	return p
}

// WithCount specifies the number of packets used by a Ping operation.
func (p *PingOp) WithCount(count int32) *PingOp {
	p.count = count
	return p
}

// WithPacketSize specifies the size of each packet (in bytes) used by a Ping operation.
func (p *PingOp) WithPacketSize(packetSize int32) *PingOp {
	p.packetSize = packetSize
	return p
}

// Operate performs the Ping operation.
func (p *PingOp) Operate(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Pinging from %s", p.dut)
	if err := ping(context.Background(), p.dut, p.dest, p.count, p.packetSize); err != nil {
		t.Fatalf("Operate(t) on %s: %v", p, err)
	}
}

func ping(ctx context.Context, dut binding.DUT, dest string, count, packetSize int32) error {
	if dest == "" {
		return fmt.Errorf("no destination for ping operation: %v", dest)
	}
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	ping, err := gnoi.System().Ping(ctx, &spb.PingRequest{
		Destination: dest,
		Count:       count,
		Size:        packetSize,
	})
	if err != nil {
		return fmt.Errorf("error on gnoi ping of %s from %v: %w", dest, dut, err)
	}

	// Ping result is included in the last message of the stream.
	lastPingResp := &spb.PingResponse{}
	for {
		resp, err := ping.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error receiving ping response: %w", err)
		}
		lastPingResp = resp
	}

	if sent, recv := lastPingResp.GetSent(), lastPingResp.GetReceived(); sent != recv {
		return fmt.Errorf("ping sent %d packets, received %d", sent, recv)
	}
	return nil
}

// Time returns the current system time.
func (o *Operations) Time(t testing.TB) time.Time {
	t.Helper()
	t = events.ActionStarted(t, "Requesting System Time from %s", o.dut)
	time, err := sysTime(context.Background(), o.dut)
	if err != nil {
		t.Fatalf("SystemTime(t) on %s: %v", o.dut, err)
	}
	return time
}

func sysTime(ctx context.Context, dut binding.DUT) (time.Time, error) {
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return time.Time{}, err
	}
	resp, err := gnoi.System().Time(ctx, &spb.TimeRequest{})
	if err != nil {
		return time.Time{}, fmt.Errorf("error fetching system time: %w", err)
	}
	return time.Unix(0, int64(resp.GetTime())), nil
}

// NewFactoryReset creates a new Factory Reset Operation.
// By default the FactoryReset is performed without zero_fill and factory_os.
func (o *Operations) NewFactoryReset() *FactoryResetOp {
	return &FactoryResetOp{dut: o.dut}
}

// FactoryResetOp is a factory reset operation.
type FactoryResetOp struct {
	dut       binding.DUT
	factoryOS bool
	zeroFill  bool
}

// WithFactoryOS instructs the device to rollback to its original OS version.
func (s *FactoryResetOp) WithFactoryOS(factoryOS bool) *FactoryResetOp {
	s.factoryOS = factoryOS
	return s
}

// WithZeroFill instructs the device to zero fill persistent storage state data.
func (s *FactoryResetOp) WithZeroFill(zeroFill bool) *FactoryResetOp {
	s.zeroFill = zeroFill
	return s
}

// String representation of the method.
func (s *FactoryResetOp) String() string {
	return fmt.Sprintf("FactoryResetOp%+v", *s)
}

// Operate performs the FactoryReset operation.
func (s *FactoryResetOp) Operate(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Performing Factory Reset on %s", s.dut)
	if err := factoryReset(context.Background(), s.dut, s.factoryOS, s.zeroFill); err != nil {
		t.Fatalf("Operate(t) on %s: %v", s, err)
	}
}

func factoryReset(ctx context.Context, dut binding.DUT, factoryOS, zeroFill bool) error {
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	resp, err := gnoi.FactoryReset().Start(ctx, &frpb.StartRequest{
		FactoryOs: factoryOS,
		ZeroFill:  zeroFill,
	})
	if err != nil {
		return fmt.Errorf("error performing factory reset: %w ", err)
	}
	if rpcErr := resp.GetResetError(); rpcErr != nil {
		return fmt.Errorf("error performing factory reset: %v ", rpcErr)
	}
	return nil
}

// NewReboot creates a new reboot operation.
func (o *Operations) NewReboot() *RebootOp {
	return &RebootOp{dut: o.dut}
}

// RebootOp is a reboot operation.
type RebootOp struct {
	dut     binding.DUT
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
	t = events.ActionStarted(t, "Rebooting %s", r.dut)
	if err := reboot(context.Background(), r.dut, r.timeout); err != nil {
		t.Fatalf("Operate(t) on %s: %v", r, err)
	}
}

func reboot(ctx context.Context, dut binding.DUT, timeout time.Duration) error {
	const (
		defaultRebootTimeout = 30 * time.Minute
		defaultStatusWait    = 10 * time.Second
	)
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	if _, err := gnoi.System().Reboot(ctx, &spb.RebootRequest{Method: spb.RebootMethod_COLD}); err != nil {
		return fmt.Errorf("error on gnoi reboot: %w", err)
	}
	rebootTimeout := timeout
	switch {
	case rebootTimeout == 0:
		rebootTimeout = defaultRebootTimeout
	case rebootTimeout < 0:
		return errors.New("reboot timeout must be a positive duration")
	}
	rebootDeadline := time.Now().Add(rebootTimeout)
	retry := true
	for retry {
		if time.Now().After(rebootDeadline) {
			retry = false
			break
		}
		resp, err := gnoi.System().RebootStatus(ctx, &spb.RebootStatusRequest{})
		switch {
		case status.Code(err) == codes.Unimplemented:
			// Unimplemented means we don't have a valid way
			// to validate health of reboot.
			return nil
		case err == nil:
			if !resp.GetActive() {
				return nil
			}
		default:
			// any other error just sleep.
		}
		statusWait := time.Duration(resp.GetWait()) * time.Nanosecond
		if statusWait <= 0 {
			statusWait = defaultStatusWait
		}
		time.Sleep(statusWait)
	}
	return fmt.Errorf("reboot of %s timed out after %s", dut, rebootTimeout)
}

// NewKillProcess creates a new kill process operation.
// By default the process is killed with a SIGTERM signal.
func (o *Operations) NewKillProcess() *KillProcessOp {
	return &KillProcessOp{
		dut: o.dut,
		req: &spb.KillProcessRequest{Signal: spb.KillProcessRequest_SIGNAL_TERM},
	}
}

// KillProcessOp is an operation that kills a process on a device.
type KillProcessOp struct {
	dut binding.DUT
	req *spb.KillProcessRequest
}

// WithPID sets the PID to be killed.
func (r *KillProcessOp) WithPID(pid uint32) *KillProcessOp {
	r.req.Pid = pid
	return r
}

// WithName sets the name of the process to be killed.
func (r *KillProcessOp) WithName(name string) *KillProcessOp {
	r.req.Name = name
	return r
}

// WithSIGTERM sets the kill signal to SIGTERM.
func (r *KillProcessOp) WithSIGTERM() *KillProcessOp {
	r.req.Signal = spb.KillProcessRequest_SIGNAL_TERM
	return r
}

// WithSIGKILL sets the kill signal to SIGKILL.
func (r *KillProcessOp) WithSIGKILL() *KillProcessOp {
	r.req.Signal = spb.KillProcessRequest_SIGNAL_KILL
	return r
}

// WithSIGHUP sets the kill signal to SIGHUP.
func (r *KillProcessOp) WithSIGHUP() *KillProcessOp {
	r.req.Signal = spb.KillProcessRequest_SIGNAL_HUP
	return r
}

// WithRestart sets whether the process should restart after being killed.
func (r *KillProcessOp) WithRestart(restart bool) *KillProcessOp {
	r.req.Restart = true
	return r
}

func (r *KillProcessOp) String() string {
	return fmt.Sprintf("KillProcessOp%+v", *r)
}

// Operate performs the kill process operation.
func (r *KillProcessOp) Operate(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Killing process on %s", r.dut)
	if err := killProcess(context.Background(), r.dut, r.req); err != nil {
		t.Fatalf("Operate(t) on %s: %v", r, err)
	}
}

func killProcess(ctx context.Context, dut binding.DUT, req *spb.KillProcessRequest) error {
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	_, err = gnoi.System().KillProcess(ctx, req)
	return err
}

// NewSwitchControlProcessor creates a new switch control processor operation.
func (o *Operations) NewSwitchControlProcessor() *SwitchControlProcessorOp {
	return &SwitchControlProcessorOp{dut: o.dut}
}

// SwitchControlProcessorOp is an operation that switches from the current
// route procesor to a provided destination route processor.
type SwitchControlProcessorOp struct {
	dut  binding.DUT
	dest string
}

func (s *SwitchControlProcessorOp) String() string {
	return fmt.Sprintf("SwitchControlProcessorOp%+v", *s)
}

// WithDestination sets the destination route processor.
func (s *SwitchControlProcessorOp) WithDestination(dest string) *SwitchControlProcessorOp {
	s.dest = dest
	return s
}

// Operate performs the SwitchControlProcessor operation.
func (s *SwitchControlProcessorOp) Operate(t testing.TB) {
	t.Helper()
	if err := switchControlProcessor(context.Background(), s.dut, s.dest); err != nil {
		t.Fatalf("Operate(t) on %s: %v", s, err)
	}
}

func switchControlProcessor(ctx context.Context, dut binding.DUT, dest string) error {
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	procPath := &tpb.Path{
		Origin: "openconfig-platform",
		Elem: []*tpb.PathElem{
			{Name: "components"},
			{Name: "component", Key: map[string]string{"name": dest}},
			{Name: "state"},
			{Name: "location"},
		},
	}
	_, err = gnoi.System().SwitchControlProcessor(ctx, &spb.SwitchControlProcessorRequest{ControlProcessor: procPath})
	return err
}
