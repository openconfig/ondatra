package ondatra

import (
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/telemetry"

	opb "github.com/openconfig/ondatra/proto"
)

// Device represents a network device.
type Device struct {
	id  string
	res reservation.Device
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
)

// String returns the name of the vendor.
func (v Vendor) String() string {
	return opb.Device_Vendor(v).String()
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
	rp, err := d.res.Dimensions().Port(id)
	if err != nil {
		return nil, err
	}
	return d.newPort(id, rp), nil
}

func (d *Device) newPort(id string, res *reservation.Port) *Port {
	return &Port{dev: d, id: id, res: res}
}

// Operations returns a handle to the device operations API.
func (d *Device) Operations() *Operations {
	return &Operations{d.res}
}

// Telemetry returns a telemetry path root for the device.
func (d *Device) Telemetry() *telemetry.DevicePath {
	return telemetry.DeviceRoot(d.ID())
}

// Port represents a port.
type Port struct {
	dev *Device
	id  string
	res *reservation.Port
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
)

// TODO: Restore when speed is consistently propagated back.
// Speed returns the port speed.
// func (p *Port) Speed() Speed {
//	return Speed(p.pb.GetSpeed())
//}
