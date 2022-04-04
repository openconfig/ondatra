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

// Package operations controls device operations for ONDATRA tests.
package operations

import (
	"golang.org/x/net/context"
	"fmt"
	"io"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"github.com/openconfig/gnoi/types"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/usererr"
	"github.com/openconfig/ondatra/internal/ate"
	"github.com/openconfig/ondatra/internal/testbed"

	ospb "github.com/openconfig/gnoi/os"
	spb "github.com/openconfig/gnoi/system"
	opb "github.com/openconfig/ondatra/proto"
)

const (
	defaultRebootTimeout = 30 * time.Minute
	defaultStatusWait    = 10 * time.Second
)

var (
	mu    sync.Mutex
	gnois = make(map[*binding.DUT]binding.GNOIClients)
)

// NewGNOI creates a gNOI client for the specified DUT.
func NewGNOI(ctx context.Context, dut *binding.DUT) (binding.GNOIClients, error) {
	return testbed.Bind().DialGNOI(ctx, dut, grpc.WithBlock())
}

// FetchGNOI fetches a cached gNOI client for the given DUT.
func FetchGNOI(ctx context.Context, dut *binding.DUT) (binding.GNOIClients, error) {
	mu.Lock()
	defer mu.Unlock()
	gnoi, ok := gnois[dut]
	if !ok {
		var err error
		gnoi, err = NewGNOI(ctx, dut)
		if err != nil {
			return nil, fmt.Errorf("error dialing gNOI: %w", err)
		}
		gnois[dut] = gnoi
	}
	return gnoi, nil
}

// Install executes an install operation.
// The gNOI install scenarios are documented on the Install function here:
// https://github.com/openconfig/gnoi/blob/master/os/os.proto
func Install(ctx context.Context, dev binding.Device, version string, standby bool, reader io.Reader) error {
	dut, err := checkDUT(dev, "ping")
	if err != nil {
		return err
	}
	if version == "" {
		return usererr.New("version not set in install operation on device: %v", dev)
	}
	gnoi, err := FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	ic, err := gnoi.OS().Install(ctx)
	if err != nil {
		return errors.Wrap(err, "error creating gnoi install client")
	}
	sreq := &ospb.InstallRequest{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{
		Version:           version,
		StandbySupervisor: standby,
	}}}
	if err := ic.Send(sreq); err != nil {
		return errors.Wrap(err, "error sending gnoi install request")
	}
	validated, err := awaitPackageInstall(ic)
	if err != nil {
		return err
	}
	if validated == nil { // if content is needed
		if reader == nil {
			return errors.Errorf("device %v does not have package and no package specified in install operation", dev)
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
		return errors.Errorf("installed version %q does not match requested version %q", gotVersion, version)
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
			return nil, usererr.New("installation error %q: %s", errName, v.InstallError.GetDetail())
		case *ospb.InstallResponse_TransferProgress:
			log.Infof("installation progress: %v bytes received from client", v.TransferProgress.GetBytesReceived())
		case *ospb.InstallResponse_SyncProgress:
			log.Infof("installation progress: %v%% synced from supervisor", v.SyncProgress.GetPercentageTransferred())
		default:
			return nil, errors.Errorf("unexpected client install response: %v (%T)", v, v)
		}
	}
}

// Ping executes the ping command from target device to a specified destination.
func Ping(ctx context.Context, dev binding.Device, dest string, count int32) error {
	dut, err := checkDUT(dev, "ping")
	if err != nil {
		return err
	}
	if dest == "" {
		return errors.Errorf("no destination for ping operation: %v", dest)
	}
	gnoi, err := FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	ping, err := gnoi.System().Ping(ctx, &spb.PingRequest{
		Destination: dest,
		Count:       count,
	})
	if err != nil {
		return errors.Wrapf(err, "error on gnoi ping of %s from %v", dest, dev)
	}

	// Ping result is included in the last message of the stream.
	lastPingResp := &spb.PingResponse{}
	for {
		resp, err := ping.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "error receiving ping response")
		}
		lastPingResp = resp
	}

	if sent, recv := lastPingResp.GetSent(), lastPingResp.GetReceived(); sent != recv {
		return fmt.Errorf("ping sent %d packets, received %d", sent, recv)
	}
	return nil
}

// SetInterfaceState sets the state of a specified interface on a device.
func SetInterfaceState(ctx context.Context, dev binding.Device, intf string, enabled *bool) error {
	if intf == "" {
		return errors.Errorf("no interface provided in set interface state operation on device %v", dev)
	}
	if enabled == nil {
		return errors.Errorf("no enabled state provided in set interface state operation on device %v", dev)
	}
	if dut, ok := dev.(*binding.DUT); ok {
		return setDUTInterfaceState(ctx, dut, intf, *enabled)
	}
	return ate.SetInterfaceState(ctx, dev.(*binding.ATE), intf, *enabled)
}

var (
	enableConfigs = map[opb.Device_Vendor]string{
		opb.Device_ARISTA: `
		  interface %s
        no shutdown
      !`,
		opb.Device_CISCO: `
		  interface %s
        no shutdown
      !`,
		opb.Device_JUNIPER: `
		  interfaces {
		    %s {
	        delete: disable;
	      }
	    }`,
	}
	disableConfigs = map[opb.Device_Vendor]string{
		opb.Device_ARISTA: `
		  interface %s
        shutdown
      !`,
		opb.Device_CISCO: `
		  interface %s
        shutdown
      !`,
		opb.Device_JUNIPER: `
		  interfaces {
		    %s {
	        disable;
	      }
	    }`,
	}
)

func setDUTInterfaceState(ctx context.Context, dut *binding.DUT, intf string, enabled bool) error {
	configs := disableConfigs
	if enabled {
		configs = enableConfigs
	}
	cfgFormat, ok := configs[dut.Vendor]
	if !ok {
		return fmt.Errorf("SetInterfaceState not yet supported for vendor %v", dut.Vendor)
	}
	cfg := fmt.Sprintf(cfgFormat, intf)
	if err := testbed.Bind().PushConfig(ctx, dut, cfg, false); err != nil {
		return errors.Wrap(err, "failed to set interface state")
	}
	return nil
}

// Reboot reboots a device.
func Reboot(ctx context.Context, dev binding.Device, timeout time.Duration) error {
	dut, err := checkDUT(dev, "restart routing")
	if err != nil {
		return err
	}
	gnoi, err := FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	if _, err := gnoi.System().Reboot(ctx, &spb.RebootRequest{Method: spb.RebootMethod_COLD}); err != nil {
		return errors.Wrap(err, "error on gnoi reboot")
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
	return errors.Errorf("reboot of %s timed out after %s", dev, rebootTimeout)
}

// KillProcess kills a process on a device, and optionally restarts it.
func KillProcess(ctx context.Context, dev binding.Device, req *spb.KillProcessRequest) error {
	dut, err := checkDUT(dev, "restart routing")
	if err != nil {
		return err
	}
	gnoi, err := FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	_, err = gnoi.System().KillProcess(ctx, req)
	return err
}

func checkDUT(dev binding.Device, op string) (*binding.DUT, error) {
	if _, ok := dev.(*binding.ATE); ok {
		return nil, errors.Errorf("%s operation not supported on ATEs: %v", op, dev)
	}
	return dev.(*binding.DUT), nil
}


// SwitchControlProcessor switches to destination route processor.
func SwitchControlProcessor(ctx context.Context, dev binding.Device, dest string) error {
       dut, err := checkDUT(dev, "switch control processor")
       if err != nil {
               return err
       }
       gnoi, err := FetchGNOI(ctx, dut)
       if err != nil {
              return err
       }
       controlProcessor := &types.Path{
               Origin: "openconfig-platform",
               Elem: []*types.PathElem{
                       {Name: "components",},
                       {Name: "component",Key: map[string]string{"name": dest}},
                       {Name: "state"},
                       {Name: "location"},
               },
       }
       _, err = gnoi.System().SwitchControlProcessor(ctx,&spb.SwitchControlProcessorRequest{ControlProcessor:controlProcessor})
       return err
}
