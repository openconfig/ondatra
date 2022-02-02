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
	"fmt"
	"math"
	"sync"
	"testing"

	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/ate"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/internal/testbed"
	"github.com/openconfig/ondatra/telemetry/device"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
)

// Device represents a network device.
type Device struct {
	id  string
	res binding.Device
	// clientFn is the func used by the telemetry library to get the gnmi client for a DUT.
	clientFn func(context.Context) (gpb.GNMIClient, error)
}

// String returns a string repesentation of the device.
func (d *Device) String() string {
	return fmt.Sprintf("%s(%s)", d.ID(), d.Name())
}

// ID returns the device id.
func (d *Device) ID() string {
	return d.id
}

// Name returns the device name.
func (d *Device) Name() string {
	return d.res.Dimensions().Name
}

// Vendor is a DUT vendor.
type Vendor int

const (
	// ARISTA vendor.
	ARISTA = Vendor(opb.Device_ARISTA)
	// CISCO vendor.
	CISCO = Vendor(opb.Device_CISCO)
	// JUNIPER vendor.
	JUNIPER = Vendor(opb.Device_JUNIPER)
	// IXIA vendor.
	IXIA = Vendor(opb.Device_IXIA)
	// CIENA vendor.
	CIENA = Vendor(opb.Device_CIENA)
	// PALOALTO vendor.
	PALOALTO = Vendor(opb.Device_PALOALTO)
)

// String returns the name of the vendor.
func (v Vendor) String() string {
	return opb.Device_Vendor(v).String()
}

// Telemetry returns a telemetry path root for the device.
func (d *Device) Telemetry() *device.DevicePath {
	root := device.DeviceRoot(d.ID())
	// TODO: Add field to root node in ygot instead of using custom data.
	root.PutCustomData(genutil.DefaultClientKey, d.clientFn)
	return root
}

// Vendor returns the device vendor.
func (d *Device) Vendor() Vendor {
	return Vendor(d.res.Dimensions().Vendor)
}

// Model returns the device hardware model.
func (d *Device) Model() string {
	return d.res.Dimensions().HardwareModel
}

// Version returns the device software version.
func (d *Device) Version() string {
	return d.res.Dimensions().SoftwareVersion
}

// Port returns a port with a given id.
func (d *Device) Port(t testing.TB, ID string) *Port {
	t.Helper()
	p, err := d.port(ID)
	if err != nil {
		t.Fatalf("Port(t, %s) on %s: %v", ID, d, err)
	}
	return p
}

// Ports returns a slice of all ports configured on the device.
func (d *Device) Ports() []*Port {
	var ports []*Port
	for id, p := range d.res.Dimensions().Ports {
		ports = append(ports, d.newPort(id, p))
	}
	return ports
}

func (d *Device) port(id string) (*Port, error) {
	rp, err := testbed.Port(d.res.Dimensions(), id)
	if err != nil {
		return nil, err
	}
	return d.newPort(id, rp), nil
}

func (d *Device) newPort(id string, res *binding.Port) *Port {
	return &Port{dev: d, id: id, res: res}
}

// Operations returns a handle to the device operations API.
func (d *Device) Operations() *Operations {
	return &Operations{d.res}
}

// Port represents a port.
type Port struct {
	dev *Device
	id  string
	res *binding.Port
}

func (p *Port) String() string {
	return fmt.Sprintf("%s:%s(%s:%s)", p.dev.ID(), p.ID(), p.dev.Name(), p.Name())
}

// ID returns the port ID.
func (p *Port) ID() string {
	return p.id
}

// Name returns the port name.
func (p *Port) Name() string {
	return p.res.Name
}

// Device returns the device that owns the port.
func (p *Port) Device() *Device {
	return p.dev
}

// Speed is a port speed.
type Speed int

const (
	// Speed10Gb is a port speed of 10Gbps.
	Speed10Gb = Speed(opb.Port_S_10GB)
	// Speed100Gb is a port speed of 100Gbps.
	Speed100Gb = Speed(opb.Port_S_100GB)
	// Speed400Gb is a port speed of 400Gbps.
	Speed400Gb = Speed(opb.Port_S_400GB)
)

// Speed returns the port speed.
func (p *Port) Speed() Speed {
	return Speed(p.res.Speed)
}

var (
	gnmisMu sync.Mutex
	gnmis   = make(map[binding.Device]gpb.GNMIClient)
)

// newGNMI creates a new gNMI client for the specified Device.
func newGNMI(ctx context.Context, dev binding.Device) (gpb.GNMIClient, error) {
	dialGNMI := func(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
		return testbed.Bind().DialGNMI(ctx, dev.(*binding.DUT), opts...)
	}
	if rATE, ok := dev.(*binding.ATE); ok {
		dialGNMI = func(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
			return ate.DialGNMI(ctx, rATE, opts...)
		}
	}
	return dialGNMI(ctx,
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)))
}

// fetchGNMI fetches the gNMI client for the given device.
// If a GNMIClient is provided it will be just returned as is and not cached.
func fetchGNMI(ctx context.Context, dev binding.Device, c gpb.GNMIClient) (gpb.GNMIClient, error) {
	if c != nil {
		return c, nil
	}
	gnmisMu.Lock()
	defer gnmisMu.Unlock()
	c, ok := gnmis[dev]
	if !ok {
		var err error
		c, err = newGNMI(ctx, dev)
		if err != nil {
			return nil, err
		}
		gnmis[dev] = c
	}
	return c, nil
}
