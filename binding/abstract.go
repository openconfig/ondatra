// Copyright 2022 Google LLC
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

package binding

import (
	"golang.org/x/net/context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var _ DUT = &AbstractDUT{}

// AbstractDUT is a reserved DUT.
// All implementations of the DUT interface must embed this type.
type AbstractDUT struct {
	*Dims
}

// Name returns the name of the DUT.
func (d *AbstractDUT) Name() string {
	return d.Dims.Name
}

// Vendor returns the vendor of the DUT.
func (d *AbstractDUT) Vendor() opb.Device_Vendor {
	return d.Dims.Vendor
}

// SoftwareVersion returns the software version of the DUT.
func (d *AbstractDUT) SoftwareVersion() string {
	return d.Dims.SoftwareVersion
}

// HardwareModel returns the hardware model of the DUT.
func (d *AbstractDUT) HardwareModel() string {
	return d.Dims.HardwareModel
}

// Ports returns the reserved ports on the DUT.
func (d *AbstractDUT) Ports() map[string]*Port {
	return d.Dims.Ports
}

func (d *AbstractDUT) String() string {
	return fmt.Sprintf("DUT%+v", *d)
}

func (d *AbstractDUT) PushConfig(ctx context.Context, config string, reset bool) error {
	return errors.New("PushConfig unimplemented")
}

func (d *AbstractDUT) DialCLI(context.Context, ...grpc.DialOption) (StreamClient, error) {
	return nil, errors.New("DialCLI unimplemented")
}

func (d *AbstractDUT) DialConsole(context.Context, ...grpc.DialOption) (StreamClient, error) {
	return nil, errors.New("DialConsole unimplemented")
}

func (d *AbstractDUT) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	return nil, errors.New("DialGNMI unimplemented")
}

func (d *AbstractDUT) DialGNOI(context.Context, ...grpc.DialOption) (GNOIClients, error) {
	return nil, errors.New("DialGNOI unimplemented")
}

func (d *AbstractDUT) DialGRIBI(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
	return nil, errors.New("DialGRIBI unimplemented")
}

func (d *AbstractDUT) DialP4RT(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
	return nil, errors.New("DialP4RT unimplemented")
}

func (d *AbstractDUT) isDUT() {}

var _ ATE = &AbstractATE{}

// AbstractATE is implementation support for the ATE interface.
// All implementations of the ATE interface must embed this type.
type AbstractATE struct {
	Dims *Dims
}

// Name returns the name of the ATE.
func (a *AbstractATE) Name() string {
	return a.Dims.Name
}

// Vendor returns the vendor of the ATE.
func (a *AbstractATE) Vendor() opb.Device_Vendor {
	return a.Dims.Vendor
}

// SoftwareVersion returns the software version of the ATE.
func (a *AbstractATE) SoftwareVersion() string {
	return a.Dims.SoftwareVersion
}

// HardwareModel returns the hardware model of the ATE.
func (a *AbstractATE) HardwareModel() string {
	return a.Dims.HardwareModel
}

// Ports returns the reserved ports on the ATE.
func (a *AbstractATE) Ports() map[string]*Port {
	return a.Dims.Ports
}

func (a *AbstractATE) String() string {
	return fmt.Sprintf("ATE%+v", *a)
}

// DialIxNetwork creates a client connection to the specified ATE's IxNetwork endpoint.
func (a *AbstractATE) DialIxNetwork(context.Context) (*IxNetwork, error) {
	return nil, errors.New("DialIxNetwork unimplemented")
}

func (a *AbstractATE) isATE() {}
