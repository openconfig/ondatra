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
	"errors"
	"fmt"
	"io"
	"time"

	log "github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/rawapis"

	frpb "github.com/openconfig/gnoi/factory_reset"
	ospb "github.com/openconfig/gnoi/os"
	spb "github.com/openconfig/gnoi/system"
	tpb "github.com/openconfig/gnoi/types"
	opb "github.com/openconfig/ondatra/proto"
)

const (
	defaultRebootTimeout = 30 * time.Minute
	defaultStatusWait    = 10 * time.Second
)

// Install executes an install operation.
// The gNOI install scenarios are documented on the Install function here:
// https://github.com/openconfig/gnoi/blob/master/os/os.proto
func Install(ctx context.Context, dut binding.DUT, version string, standby bool, reader io.Reader) error {
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

// Ping executes the ping command from target device to a specified destination.
func Ping(ctx context.Context, dut binding.DUT, dest string, count int32) error {
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

// SystemTime gives the current system time of the device.
func SystemTime(ctx context.Context, dut binding.DUT) (time.Time, error) {
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

// FactoryReset performs a factory reset of the device.
//
//	factoryOS instructs the device to rollback the its original OS version
//	zeroFill instructs the device to zero fill persistent storage state data
func FactoryReset(ctx context.Context, dut binding.DUT, factoryOS, zeroFill bool) error {
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

// SetInterfaceState sets the state of a specified interface on a device.
func SetInterfaceState(ctx context.Context, dut binding.DUT, intf string, enabled *bool) error {
	if intf == "" {
		return fmt.Errorf("no interface provided in set interface state operation on DUT %v", dut)
	}
	if enabled == nil {
		return fmt.Errorf("no enabled state provided in set interface state operation on DUT %v", dut)
	}
	configs := disableConfigs
	if *enabled {
		configs = enableConfigs
	}
	cfgFormat, ok := configs[dut.Vendor()]
	if !ok {
		return fmt.Errorf("SetInterfaceState not yet supported for vendor %v", dut.Vendor())
	}
	cfg := fmt.Sprintf(cfgFormat, intf)
	if err := dut.PushConfig(ctx, cfg, false); err != nil {
		return fmt.Errorf("failed to set interface state: %w", err)
	}
	return nil
}

// Reboot reboots a device.
func Reboot(ctx context.Context, dut binding.DUT, timeout time.Duration) error {
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

// KillProcess kills a process on a device, and optionally restarts it.
func KillProcess(ctx context.Context, dut binding.DUT, req *spb.KillProcessRequest) error {
	gnoi, err := rawapis.FetchGNOI(ctx, dut)
	if err != nil {
		return err
	}
	_, err = gnoi.System().KillProcess(ctx, req)
	return err
}

// SwitchControlProcessor switches to a provided destination route processor.
func SwitchControlProcessor(ctx context.Context, dut binding.DUT, dest string) error {
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
