// Package operations controls device operations for ONDATRA tests.
package operations

import (
	"golang.org/x/net/context"
	"fmt"
	"io"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/closer"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/usererr"

	ospb "github.com/openconfig/gnoi/os"
	spb "github.com/openconfig/gnoi/system"
)

const (
	defaultRebootTimeout = 30 * time.Minute
	defaultStatusWait    = 10 * time.Second
)

var (
	mu    sync.Mutex
	gnois = make(map[*reservation.DUT]binding.GNOIClients)
)

// NewGNOI creates a new gNOI client for the specified DUT.
func NewGNOI(ctx context.Context, dut *reservation.DUT) (binding.GNOIClients, error) {
	return binding.Get().DialGNOI(ctx, dut, grpc.WithBlock())
}

// fetchGNOI fetches a cached gNOI client for the given DUT.
func fetchGNOI(ctx context.Context, dut *reservation.DUT) (binding.GNOIClients, error) {
	mu.Lock()
	defer mu.Unlock()
	gnoi, ok := gnois[dut]
	if !ok {
		var err error
		gnoi, err = NewGNOI(ctx, dut)
		if err != nil {
			return nil, err
		}
		gnois[dut] = gnoi
	}
	return gnoi, nil
}

// Install executes an install operation.
// The gNOI install scenarios are documented on the Install function here:
// https://github.com/openconfig/gnoi/blob/master/os/os.proto
func Install(ctx context.Context, dev reservation.Device, version string, standby bool, reader io.Reader) error {
	dut, err := checkDUT(dev, "ping")
	if err != nil {
		return err
	}
	if version == "" {
		return usererr.New("version not set in install operation on device: %v", dev)
	}
	gnoi, err := fetchGNOI(ctx, dut)
	if err != nil {
		return errors.Wrap(err, "error dialing gnoi")
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
func Ping(ctx context.Context, dev reservation.Device, dest string) (rerr error) {
	dut, err := checkDUT(dev, "ping")
	if err != nil {
		return err
	}
	if dest == "" {
		return errors.Errorf("no destination for ping operation: %v", dest)
	}
	gnoi, err := fetchGNOI(ctx, dut)
	if err != nil {
		return errors.Wrap(err, "error dialing gnoi")
	}
	ping, err := gnoi.System().Ping(ctx, &spb.PingRequest{Destination: dest})
	if err != nil {
		return errors.Wrapf(err, "error on gnoi ping of %s from %v", dest, dev)
	}
	defer closer.Close(&rerr, ping.CloseSend, "error closing gnoi ping client")
	if _, err := ping.Recv(); err != nil {
		return errors.Wrapf(err, "error on ping recv")
	}
	return nil
}

// SetInterfaceState sets the state of a specified interface on a device.
func SetInterfaceState(ctx context.Context, dev reservation.Device, intf string, enabled *bool) error {
	if intf == "" {
		return errors.Errorf("no interface provided in set interface state operation on device %v", dev)
	}
	if enabled == nil {
		return errors.Errorf("no enabled state provided in set interface state operation on device %v", dev)
	}
	if dut, ok := dev.(*reservation.DUT); ok {
		return setDUTInterfaceState(ctx, dut, intf, *enabled)
	}
	return binding.Get().SetATEPortState(dev.(*reservation.ATE), intf, *enabled)
}

func setDUTInterfaceState(ctx context.Context, dut *reservation.DUT, intf string, enabled bool) error {
	const ocFormat = `{
   "interfaces": {
      "interface": [
         {
            "config": {
               "enabled": %v,
               "name": "%s",
               "type": "ethernetCsmacd"
            },
            "name": "%s"
         }
      ]
   }
}`
	ocCfg := fmt.Sprintf(ocFormat, enabled, intf, intf)
	opts := &binding.ConfigOptions{OpenConfig: true, Append: true}
	if err := binding.Get().PushConfig(ctx, dut, ocCfg, opts); err != nil {
		return errors.Wrap(err, "failed to set interface state")
	}
	return nil
}

// Reboot reboots a device.
func Reboot(ctx context.Context, dev reservation.Device, timeout time.Duration) error {
	dut, err := checkDUT(dev, "restart routing")
	if err != nil {
		return err
	}
	gnoi, err := fetchGNOI(ctx, dut)
	if err != nil {
		return errors.Wrap(err, "error dialing gnoi")
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

// RestartRouting restarts routing on a device.
func RestartRouting(ctx context.Context, dev reservation.Device) error {
	dut, err := checkDUT(dev, "restart routing")
	if err != nil {
		return err
	}
	return binding.Get().RestartRouting(dut)
}

func checkDUT(dev reservation.Device, op string) (*reservation.DUT, error) {
	if _, ok := dev.(*reservation.ATE); ok {
		return nil, errors.Errorf("%s operation not supported on ATEs: %v", op, dev)
	}
	return dev.(*reservation.DUT), nil
}
