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
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/internal/testbed"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
)

var _ gnmi.DeviceOrOpts = (*Device)(nil)

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
	return d.res.Name()
}

// Vendor is a DUT vendor.
type Vendor int

const (
	// ARISTA vendor.
	ARISTA = Vendor(opb.Device_ARISTA)
	// CISCO vendor.
	CISCO = Vendor(opb.Device_CISCO)
	// DELL vendor.
	DELL = Vendor(opb.Device_DELL)
	// JUNIPER vendor.
	JUNIPER = Vendor(opb.Device_JUNIPER)
	// IXIA vendor.
	IXIA = Vendor(opb.Device_IXIA)
	// CIENA vendor.
	CIENA = Vendor(opb.Device_CIENA)
	// OPENCONFIG vendor.
	OPENCONFIG = Vendor(opb.Device_OPENCONFIG)
	// PALOALTO vendor.
	PALOALTO = Vendor(opb.Device_PALOALTO)
	// ZPE vendor.
	ZPE = Vendor(opb.Device_ZPE)
	// NOKIA vendor.
	NOKIA = Vendor(opb.Device_NOKIA)
)

// String returns the name of the vendor.
func (v Vendor) String() string {
	return opb.Device_Vendor(v).String()
}

// GNMIOpts returns a new set of options to customize gNMI queries.
func (d *Device) GNMIOpts() *gnmi.Opts {
	useGetForCfg := d.Vendor() == CISCO || d.Vendor() == JUNIPER
	return gnmi.NewOpts(d.Name(), useGetForCfg, d.clientFn)
}

// Vendor returns the device vendor.
func (d *Device) Vendor() Vendor {
	return Vendor(d.res.Vendor())
}

// Model returns the device hardware model.
func (d *Device) Model() string {
	return d.res.HardwareModel()
}

// Version returns the device software version.
func (d *Device) Version() string {
	return d.res.SoftwareVersion()
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
	for id, p := range d.res.Ports() {
		ports = append(ports, d.newPort(id, p))
	}
	return ports
}

func (d *Device) port(id string) (*Port, error) {
	rp, err := testbed.Port(d.res, id)
	if err != nil {
		return nil, err
	}
	return d.newPort(id, rp), nil
}

func (d *Device) newPort(id string, res *binding.Port) *Port {
	return &Port{dev: d, id: id, res: res}
}

// CustomData returns custom data for the specified key.
func (d *Device) CustomData(key string) any {
	return d.res.CustomData()[key]
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
	// Speed1Gb is a port speed of 1Gbps.
	Speed1Gb = Speed(opb.Port_S_1GB)
	// Speed5Gb is a port speed of 5Gbps.
	Speed5Gb = Speed(opb.Port_S_5GB)
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

// CardModel returns the card model.
func (p *Port) CardModel() string {
	return p.res.CardModel
}

// PMD is a Physical Medium Dependent .
type PMD opb.Port_Pmd

const (
	// PMD10GBASELRM is the PMD 10GBASE-LRM
	PMD10GBASELRM = PMD(opb.Port_PMD_10GBASE_LRM)
	// PMD10GBASELR is the PMD 10GBASE-LR.
	PMD10GBASELR = PMD(opb.Port_PMD_10GBASE_LR)
	// PMD10GBASEZR is the PMD 10GBASE-ZR.
	PMD10GBASEZR = PMD(opb.Port_PMD_10GBASE_ZR)
	// PMD10GBASEER is the PMD 10GBASE-ER.
	PMD10GBASEER = PMD(opb.Port_PMD_10GBASE_ER)
	// PMD10GBASESR is the PMD 10GBASE-SR.
	PMD10GBASESR = PMD(opb.Port_PMD_10GBASE_SR)
	// PMD40GBASECR4 is the PMD 40GBASE-CR4.
	PMD40GBASECR4 = PMD(opb.Port_PMD_40GBASE_CR4)
	// PMD40GBASESR4 is the PMD 40GBASE-SR4.
	PMD40GBASESR4 = PMD(opb.Port_PMD_40GBASE_SR4)
	// PMD40GBASELR4 is the PMD 40GBASE-LR4.
	PMD40GBASELR4 = PMD(opb.Port_PMD_40GBASE_LR4)
	// PMD40GBASEER4 is the PMD 40GBASE-ER4.
	PMD40GBASEER4 = PMD(opb.Port_PMD_40GBASE_ER4)
	// PMD40GBASEPSM4 is the PMD 40GBASE-PSM4.
	PMD40GBASEPSM4 = PMD(opb.Port_PMD_40GBASE_PSM4)
	// PMD4X10GBASELR is the PMD 4X10GBASE-LR.
	PMD4X10GBASELR = PMD(opb.Port_PMD_4X10GBASE_LR)
	// PMD4X10GBASESR is the PMD 4X10GBASE-SR.
	PMD4X10GBASESR = PMD(opb.Port_PMD_4X10GBASE_SR)
	// PMD100GAOC is the PMD 100G-AOC.
	PMD100GAOC = PMD(opb.Port_PMD_100G_AOC)
	// PMD100GACC is the PMD 100G-ACC.
	PMD100GACC = PMD(opb.Port_PMD_100G_ACC)
	// PMD100GBASESR10 is the PMD 100GBASE-SR10.
	PMD100GBASESR10 = PMD(opb.Port_PMD_100GBASE_SR10)
	// PMD100GBASESR4 is the PMD 100GBASE-SR4.
	PMD100GBASESR4 = PMD(opb.Port_PMD_100GBASE_SR4)
	// PMD100GBASELR4 is the PMD 100GBASE-LR4.
	PMD100GBASELR4 = PMD(opb.Port_PMD_100GBASE_LR4)
	// PMD100GBASEER4 is the PMD 100GBASE-ER4.
	PMD100GBASEER4 = PMD(opb.Port_PMD_100GBASE_ER4)
	// PMD100GBASECWDM4 is the PMD 100GBASE-CWDM4.
	PMD100GBASECWDM4 = PMD(opb.Port_PMD_100GBASE_CWDM4)
	// PMD100GBASECLR4 is the PMD 100GBASE-CLR4.
	PMD100GBASECLR4 = PMD(opb.Port_PMD_100GBASE_CLR4)
	// PMD100GBASEPSM4 is the PMD 100GBASE-PSM4.
	PMD100GBASEPSM4 = PMD(opb.Port_PMD_100GBASE_PSM4)
	// PMD100GBASECR4 is the PMD 100GBASE-CR4.
	PMD100GBASECR4 = PMD(opb.Port_PMD_100GBASE_CR4)
	// PMD100GBASEFR is the PMD 100GBASE-FR.
	PMD100GBASEFR = PMD(opb.Port_PMD_100GBASE_FR)
	// PMD400GBASEZR is the PMD 400GBASE-ZR.
	PMD400GBASEZR = PMD(opb.Port_PMD_400GBASE_ZR)
	// PMD400GBASELR4 is the PMD 400GBASE-LR4.
	PMD400GBASELR4 = PMD(opb.Port_PMD_400GBASE_LR4)
	// PMD400GBASEFR4 is the PMD 400GBASE-FR4.
	PMD400GBASEFR4 = PMD(opb.Port_PMD_400GBASE_FR4)
	// PMD400GBASELR8 is the PMD 400GBASE-LR8.
	PMD400GBASELR8 = PMD(opb.Port_PMD_400GBASE_LR8)
	// PMD400GBASEDR4 is the PMD 400GBASE-DR4.
	PMD400GBASEDR4 = PMD(opb.Port_PMD_400GBASE_DR4)
)

func (pmd PMD) String() string {
	return opb.Port_Pmd_name[int32(pmd)]
}

// PMD returns the Physical Medium Dependent.
func (p *Port) PMD() PMD {
	return PMD(p.res.PMD)
}
