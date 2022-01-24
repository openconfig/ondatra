// Copyright 2021 Google LLC
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

// Package knebind provides an Ondatra binding for KNE devices.
package knebind

import (
	"golang.org/x/net/context"
	"crypto/tls"
	"fmt"
	"os/exec"
	"sync"
	"time"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/knebind/solver"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	kpb "github.com/google/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var (
	fetchTopo = fetchTopology // to be stubbed out by tests
)

// Bind implements the ondatra Binding interface for KNE
type Bind struct {
	binding.Binding
	services solver.ServiceMap
	mu       sync.Mutex
	cfg      *Config
}

// New returns a new KNE bind instance.
func New(cfg *Config) (*Bind, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	return &Bind{
		cfg: cfg,
	}, nil
}

// Reserve implements the binding Reserve method by finding nodes and links in
// the topology specified in the config file that match the requested testbed.
func (b *Bind) Reserve(ctx context.Context, tb *opb.Testbed, runTime time.Duration, waitTime time.Duration) (*reservation.Reservation, error) {
	topo, err := fetchTopo(b.cfg)
	if err != nil {
		return nil, err
	}
	sol, err := solver.Solve(tb, topo)
	if err != nil {
		return nil, err
	}
	b.services = sol.Services
	return sol.Reservation, nil
}

func fetchTopology(cfg *Config) (*kpb.Topology, error) {
	args := []string{"topology", "service", cfg.TopoPath}
	if cfg.KubecfgPath != "" {
		args = append(args, fmt.Sprintf("--kubecfg=%s", cfg.KubecfgPath))
	}
	cmd := exec.Command(cfg.CLIPath, args...)
	out, err := cmd.Output()
	if err != nil {
		if execErr, ok := err.(*exec.ExitError); ok {
			return nil, errors.Wrapf(err, "error executing command %v: %s", cmd, execErr.Stderr)
		}
		return nil, errors.Wrapf(err, "error executing command %v", cmd)
	}
	topo := new(kpb.Topology)
	if err := prototext.Unmarshal(out, topo); err != nil {
		return nil, errors.Wrap(err, "error unmarshalling KNE topology proto")
	}
	return topo, nil
}

// Release is a no-op because there's need to reserve local VMs.
func (b *Bind) Release(_ context.Context) error {
	return nil
}

// SetTestMetadata is unused for KNE.
func (b *Bind) SetTestMetadata(_ *binding.TestMetadata) error {
	return nil
}

func (b *Bind) DialGNMI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	conn, err := b.dialGRPC(ctx, dut, "gnmi", opts...)
	if err != nil {
		return nil, err
	}
	return gpb.NewGNMIClient(conn), nil
}

func (b *Bind) DialP4RT(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
	conn, err := b.dialGRPC(ctx, dut, "p4rt", opts...)
	if err != nil {
		return nil, err
	}
	return p4pb.NewP4RuntimeClient(conn), nil
}

func (b *Bind) dialGRPC(ctx context.Context, dut *reservation.DUT, serviceName string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	s, err := b.services.Lookup(dut.Name, serviceName)
	if err != nil {
		return nil, err
	}
	addr := serviceAddr(s)
	log.Infof("Dialing service %q on dut %s@%s", serviceName, dut.Name, addr)
	opts = append(opts,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithPerRPCCredentials(&passCred{
			username: b.cfg.Username,
			password: b.cfg.Password,
		}))
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "DialContext(ctx, %s, %v)", addr, opts)
	}
	return conn, nil
}

// serviceAddr returns the external IP address of a KNE service.
func serviceAddr(s *kpb.Service) string {
	return fmt.Sprintf("%s:%d", s.GetOutsideIp(), s.GetOutside())
}

type passCred struct {
	username string
	password string
}

func (c *passCred) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": c.username,
		"password": c.password,
	}, nil
}

func (c *passCred) RequireTransportSecurity() bool {
	return true
}
