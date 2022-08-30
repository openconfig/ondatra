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
	"bytes"
	"golang.org/x/net/context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"

	log "github.com/golang/glog"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/prototext"
	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/gocloser"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind/solver"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var (
	// to be stubbed out by tests
	kneCmdFn  = kneCmd
	sshExecFn = sshExec
)

// New returns a new KNE bind instance.
func New(cfg *Config) (*Bind, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	return &Bind{cfg: cfg}, nil
}

// Bind implements the ondatra Binding interface for KNE
type Bind struct {
	binding.Binding
	cfg *Config
}

// Reserve implements the binding Reserve method by finding nodes and links in
// the topology specified in the config file that match the requested testbed.
func (b *Bind) Reserve(ctx context.Context, tb *opb.Testbed, runTime time.Duration, waitTime time.Duration, partial map[string]string) (*binding.Reservation, error) {
	if len(partial) > 0 {
		return nil, errors.New("KNEBind Reserve does not yet support partial mappings")
	}
	out, err := kneCmdFn(b.cfg, "topology", "service", b.cfg.TopoPath)
	if err != nil {
		return nil, err
	}
	topo := new(tpb.Topology)
	if err := prototext.Unmarshal(out, topo); err != nil {
		return nil, fmt.Errorf("error unmarshalling KNE topology proto: %w", err)
	}
	res, err := solver.Solve(tb, topo)
	if err != nil {
		return nil, err
	}
	for i, dut := range res.DUTs {
		kdut := &kneDUT{
			ServiceDUT: dut.(*solver.ServiceDUT),
			cfg:        b.cfg,
		}
		if err := kdut.resetConfig(); err != nil {
			return nil, err
		}
		res.DUTs[i] = kdut
	}
	for i, ate := range res.ATEs {
		res.ATEs[i] = &kneATE{
			ServiceATE: ate.(*solver.ServiceATE),
			cfg:        b.cfg,
		}
	}
	return res, nil
}

// Release is a no-op because there's need to reserve local VMs.
func (b *Bind) Release(context.Context) error {
	return nil
}

type kneDUT struct {
	*solver.ServiceDUT
	cfg *Config
}

func (d *kneDUT) resetConfig() error {
	_, err := kneCmdFn(d.cfg, "topology", "reset", d.cfg.TopoPath, d.Name(), "--push")
	return err
}

func (d *kneDUT) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	conn, err := d.dialGRPC(ctx, "gnmi", opts...)
	if err != nil {
		return nil, err
	}
	return gpb.NewGNMIClient(conn), nil
}

func (d *kneDUT) DialGNOI(ctx context.Context, opts ...grpc.DialOption) (binding.GNOIClients, error) {
	conn, err := d.dialGRPC(ctx, "gnoi", opts...)
	if err != nil {
		return nil, err
	}
	return &gnoiClients{conn: conn}, nil
}

func (d *kneDUT) DialGRIBI(ctx context.Context, opts ...grpc.DialOption) (grpb.GRIBIClient, error) {
	conn, err := d.dialGRPC(ctx, "gribi", opts...)
	if err != nil {
		return nil, err
	}
	return grpb.NewGRIBIClient(conn), nil
}

func (d *kneDUT) DialP4RT(ctx context.Context, opts ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
	conn, err := d.dialGRPC(ctx, "p4rt", opts...)
	if err != nil {
		return nil, err
	}
	return p4pb.NewP4RuntimeClient(conn), nil
}

func (d *kneDUT) dialGRPC(ctx context.Context, serviceName string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	s, err := d.Service(serviceName)
	if err != nil {
		return nil, err
	}
	addr := serviceAddr(s)
	log.Infof("Dialing service %q on dut %s@%s", serviceName, d.Name(), addr)
	opts = append(opts,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithPerRPCCredentials(&passCred{
			username: d.cfg.Username,
			password: d.cfg.Password,
		}))
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("DialContext(ctx, %s, %v): %w", addr, opts, err)
	}
	return conn, nil
}

// serviceAddr returns the external IP address of a KNE service.
func serviceAddr(s *tpb.Service) string {
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

func (d *kneDUT) PushConfig(ctx context.Context, config string, reset bool) error {
	if d.Vendor() != opb.Device_ARISTA {
		return errors.New("KNEBind PushConfig only supports Arista devices")
	}
	if reset {
		if err := d.resetConfig(); err != nil {
			return err
		}
	}
	_, err := d.exec("enable\nconfig terminal\n" + config)
	return err
}

func (d *kneDUT) DialCLI(context.Context, ...grpc.DialOption) (binding.StreamClient, error) {
	return &kneCLI{dut: d}, nil
}

type kneCLI struct {
	*binding.AbstractStreamClient
	dut *kneDUT
}

func (c *kneCLI) SendCommand(_ context.Context, cmd string) (string, error) {
	return c.dut.exec(cmd)
}

func (d *kneDUT) exec(cmd string) (string, error) {
	s, err := d.Service("ssh")
	if err != nil {
		return "", err
	}
	addr := serviceAddr(s)
	config := &ssh.ClientConfig{
		User: d.cfg.Username,
		Auth: []ssh.AuthMethod{ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			if len(questions) > 0 {
				return []string{d.cfg.Password}, nil
			}
			return nil, nil
		})},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return sshExecFn(addr, config, cmd)
}

func sshExec(addr string, cfg *ssh.ClientConfig, cmd string) (_ string, rerr error) {
	client, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		return "", fmt.Errorf("could not dial SSH server %s: %w", addr, err)
	}
	defer closer.Close(&rerr, client.Close, "error closing SSH client")
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("could not create ssh session: %w", err)
	}
	defer closer.Close(&rerr, func() error {
		if err := session.Close(); err != io.EOF {
			return err
		}
		return nil
	}, "error closing SSH session")
	var buf bytes.Buffer
	session.Stdout = &buf
	if err := session.Run(cmd); err != nil {
		return "", fmt.Errorf("could not execute %q\noutput: %q: %w", cmd, buf.String(), err)
	}
	return buf.String(), nil
}

type kneATE struct {
	*solver.ServiceATE
	cfg *Config
}

func (a *kneATE) DialOTG(context.Context) (gosnappi.GosnappiApi, error) {
	s, err := a.Service("grpc")
	if err != nil {
		return nil, err
	}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().
		SetLocation(serviceAddr(s)).
		SetRequestTimeout(30 * time.Second)
	return api, nil
}

func (a *kneATE) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	s, err := a.Service("gnmi")
	if err != nil {
		return nil, err
	}
	addr := serviceAddr(s)
	log.Infof("Dialing service %q on ate %s@%s", addr, a.Name(), addr)
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))) // NOLINT
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("DialContext(ctx, %s, %v): %w", addr, opts, err)
	}
	return gpb.NewGNMIClient(conn), nil
}

func kneCmd(cfg *Config, args ...string) ([]byte, error) {
	if cfg.KubecfgPath != "" {
		args = append(append([]string{}, args...), fmt.Sprintf("--kubecfg=%s", cfg.KubecfgPath))
	}
	cmd := exec.Command(cfg.CLIPath, args...)
	out, err := cmd.Output()
	if err != nil {
		if execErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("error executing command %v: %s: %w", cmd, execErr.Stderr, err)
		}
		return nil, fmt.Errorf("error executing command %v: %w", cmd, err)
	}
	return out, nil
}
