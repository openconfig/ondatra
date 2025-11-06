// Copyright 2025 Google LLC
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

// Package solver provides a library for solving a testbed reservation request
// against a given inventory of available devices and their connections.
package solver

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/portgraph"
	opb "github.com/openconfig/ondatra/proto"
	"github.com/pborman/uuid"
)

// Inventory of available binding devices and links.
// This struct is provided by a specific binding implementation.
type Inventory struct {
	DUTs  []binding.DUT
	ATEs  []binding.ATE
	Links map[*binding.Port]*binding.Port // Maps a port to its peer.
}

// SolveResult contains the result of a successful testbed solve.
type SolveResult struct {
	Assignment       *portgraph.Assignment
	AbsNode2Dev      map[*portgraph.AbstractNode]*opb.Device
	AbsPort2Port     map[*portgraph.AbstractPort]*opb.Port
	ConNode2BindDev  map[*portgraph.ConcreteNode]*binding.Device
	ConPort2BindPort map[*portgraph.ConcretePort]*binding.Port
}

// Solve finds an assignment to the provided testbed from the given inventory.
// It returns the low-level portgraph.Assignment and maps to correlate
// graph elements back to the original binding and testbed elements.
func Solve(ctx context.Context, tb *opb.Testbed, inv *Inventory, partial map[string]string) (*SolveResult, error) {
	abstractGraph, absNode2Dev, absPort2Port, err := portgraph.TestbedToAbstractGraph(tb, partial)
	if err != nil {
		return nil, fmt.Errorf("could not parse specified testbed: %w", err)
	}

	superGraph, conNode2BindDev, conPort2BindPort, err := inventoryToConcreteGraph(inv)
	if err != nil {
		return nil, fmt.Errorf("could not convert inventory to concrete graph: %w", err)
	}

	assignment, err := portgraph.Solve(ctx, abstractGraph, superGraph)
	if err != nil {
		return nil, fmt.Errorf("could not solve for specified testbed: %w", err)
	}

	return &SolveResult{
		Assignment:       assignment,
		AbsNode2Dev:      absNode2Dev,
		AbsPort2Port:     absPort2Port,
		ConNode2BindDev:  conNode2BindDev,
		ConPort2BindPort: conPort2BindPort,
	}, nil
}

// inventoryToConcreteGraph converts the generic Inventory struct to a portgraph.ConcreteGraph.
func inventoryToConcreteGraph(inv *Inventory) (*portgraph.ConcreteGraph, map[*portgraph.ConcreteNode]*binding.Device, map[*portgraph.ConcretePort]*binding.Port, error) {
	cg := &portgraph.ConcreteGraph{Desc: "Generic Inventory"}
	conNode2BindDev := make(map[*portgraph.ConcreteNode]*binding.Device)
	conPort2BindPort := make(map[*portgraph.ConcretePort]*binding.Port)
	bindPortToConPort := make(map[*binding.Port]*portgraph.ConcretePort)

	addDevice := func(dev binding.Device, role string) error {
		attrs := map[string]string{
			portgraph.RoleAttr:   role,
			portgraph.VendorAttr: dev.Vendor().String(),
			portgraph.HWAttr:     dev.HardwareModel(),
			portgraph.SWAttr:     dev.SoftwareVersion(),
			portgraph.NameAttr:   dev.Name(),
		}

		var ports []*portgraph.ConcretePort
		// Accessing ports using the map from the Dims struct.
		for portID, bp := range dev.Ports() {
			if bp == nil {
				return fmt.Errorf("device %s port ID %s has a nil *binding.Port", dev.Name(), portID)
			}
			pAttrs := map[string]string{
				portgraph.NameAttr:  bp.Name,
				portgraph.SpeedAttr: bp.Speed.String(),
				portgraph.PMDAttr:   bp.PMD.String(),
			}
			if bp.Group != "" {
				pAttrs[portgraph.GroupAttr] = bp.Group
			}
			cp := &portgraph.ConcretePort{
				// Desc should be unique for each concrete port.
				Desc:  fmt.Sprintf("%s:%s", dev.Name(), bp.Name),
				Attrs: pAttrs,
			}
			ports = append(ports, cp)
			conPort2BindPort[cp] = bp
			bindPortToConPort[bp] = cp
		}

		node := &portgraph.ConcreteNode{
			Desc:  dev.Name(),
			Ports: ports,
			Attrs: attrs,
		}
		cg.Nodes = append(cg.Nodes, node)
		conNode2BindDev[node] = &dev
		return nil
	}

	for _, dut := range inv.DUTs {
		if err := addDevice(dut, portgraph.RoleDUT); err != nil {
			return nil, nil, nil, err
		}
	}
	for _, ate := range inv.ATEs {
		if err := addDevice(ate, portgraph.RoleATE); err != nil {
			return nil, nil, nil, err
		}
	}

	edgeSet := make(map[struct {
		Src *portgraph.ConcretePort
		Dst *portgraph.ConcretePort
	}]bool)
	for srcBP, dstBP := range inv.Links {
		srcCP, srcOk := bindPortToConPort[srcBP]
		if !srcOk {
			return nil, nil, nil, fmt.Errorf("link source port %s not found in inventory device ports", srcBP.Name)
		}
		dstCP, dstOk := bindPortToConPort[dstBP]
		if !dstOk {
			return nil, nil, nil, fmt.Errorf("link destination port %s not found in inventory device ports", dstBP.Name)
		}

		// Ensure we only add each edge once by creating a canonical key.
		key := struct {
			Src *portgraph.ConcretePort
			Dst *portgraph.ConcretePort
		}{Src: srcCP, Dst: dstCP}
		if srcCP.Desc > dstCP.Desc {
			key = struct {
				Src *portgraph.ConcretePort
				Dst *portgraph.ConcretePort
			}{Src: dstCP, Dst: srcCP}
		}

		if !edgeSet[key] {
			cg.Edges = append(cg.Edges, &portgraph.ConcreteEdge{Src: key.Src, Dst: key.Dst})
			edgeSet[key] = true
		}
	}
	return cg, conNode2BindDev, conPort2BindPort, nil
}

func dynDims(dev *opb.Device, bdev binding.Device, tbPort2BindPort map[*opb.Port]*binding.Port) *binding.Dims {
	dims := &binding.Dims{
		Name:            bdev.Name(),
		Vendor:          bdev.Vendor(),
		HardwareModel:   bdev.HardwareModel(),
		SoftwareVersion: bdev.SoftwareVersion(),
		Ports:           make(map[string]*binding.Port),
	}
	for _, tp := range dev.GetPorts() {
		if bp, present := tbPort2BindPort[tp]; present {
			dims.Ports[tp.GetId()] = bp
		}
	}
	return dims
}

// AssignmentToReservation converts the solveResult to a binding.Reservation.
func AssignmentToReservation(solveResult *SolveResult, tb *opb.Testbed) (*binding.Reservation, error) {
	res := &binding.Reservation{
		ID:   uuid.New(),
		DUTs: make(map[string]binding.DUT),
		ATEs: make(map[string]binding.ATE),
	}
	if solveResult == nil {
		return nil, fmt.Errorf("solveResult is nil for this reservation: %s", res.ID)
	}
	if solveResult.Assignment == nil {
		return nil, fmt.Errorf("solveResult.Assignment is nil for this reservation: %s", res.ID)
	}

	tbDev2BindDev := make(map[*opb.Device]*binding.Device)
	for absNode, conNode := range solveResult.Assignment.Node2Node {
		tbDev2BindDev[solveResult.AbsNode2Dev[absNode]] = solveResult.ConNode2BindDev[conNode]
	}
	tbPort2BindPort := make(map[*opb.Port]*binding.Port)
	for absPort, conPort := range solveResult.Assignment.Port2Port {
		tbPort2BindPort[solveResult.AbsPort2Port[absPort]] = solveResult.ConPort2BindPort[conPort]
	}

	for _, tdut := range tb.GetDuts() {
		bdev := tbDev2BindDev[tdut]
		if bdev == nil {
			return nil, fmt.Errorf("DUT %q not resolved", tdut.GetId())
		}
		bdut, ok := (*bdev).(binding.DUT)
		if !ok {
			return nil, fmt.Errorf("device %q assigned to DUT %q is not a DUT", (*bdev).Name(), tdut.GetId())
		}
		dims := dynDims(tdut, bdut, tbPort2BindPort)
		res.DUTs[tdut.GetId()] = &binding.AbstractDUT{Dims: dims}
	}
	for _, tate := range tb.GetAtes() {
		bdev := tbDev2BindDev[tate]
		if bdev == nil {
			return nil, fmt.Errorf("ATE %q not resolved", tate.GetId())
		}
		bate, ok := (*bdev).(binding.ATE)
		if !ok {
			return nil, fmt.Errorf("device %q assigned to ATE %q is not an ATE", (*bdev).Name(), tate.GetId())
		}
		dims := dynDims(tate, bate, tbPort2BindPort)
		res.ATEs[tate.GetId()] = &binding.AbstractATE{Dims: dims}
	}
	return res, nil
}
