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
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/open-traffic-generator/snappi/gosnappi"
	closer "github.com/openconfig/gocloser"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind/solver"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/prototext"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
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
		res.DUTs[i] = kdut
		if b.cfg.SkipReset {
			continue
		}
		if err := kdut.resetConfig(); err != nil {
			return nil, err
		}
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
	_, err := kneCmdFn(d.cfg, "topology", "reset", d.cfg.TopoPath, d.Name(), "--push", "--skip")
	return err
}

func (d *kneDUT) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	//If the insecure field is set to true then set to insecure
	if d.cfg.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
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
	if !d.cfg.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))) // NOLINT
	}
	creds := newRPCCredentials(d.cfg, d.Name(), d.NodeVendor)
	if creds != nil {
		opts = append(opts, grpc.WithPerRPCCredentials(creds))
	}
	log.Infof("Dialing service %q on DUT %s@%s using credentials %+v", serviceName, d.Name(), addr, creds)
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

type rpcCredentials struct {
	username string
	password string
}

func (r *rpcCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": r.username,
		"password": r.password,
	}, nil
}

func (r *rpcCredentials) RequireTransportSecurity() bool {
	cfg := Config{}
	if !cfg.Insecure {
		return false
	} else {
		return true
	}
}

// newRPCCredentials determines the correct credentials used to access a node via rpc
// from a given node name, node vendor, and knebind config. The precedence order for determining
// the credentials is as follows:
//
// 1. credential provided for a specific node by name from the knebind config
// 2. credential provided for a vendor of the node from the knebind config
// 3. credential from the default username and password fields from the knebind config
// 4. no credentials
func newRPCCredentials(cfg *Config, name string, vendor tpb.Vendor) *rpcCredentials {
	if cfg.Credentials != nil && cfg.Credentials.Node != nil {
		if creds, ok := cfg.Credentials.Node[name]; ok {
			return &rpcCredentials{username: creds.Username, password: creds.Password}
		}
	}
	if cfg.Credentials != nil && cfg.Credentials.Vendor != nil {
		if creds, ok := cfg.Credentials.Vendor[vendor]; ok {
			return &rpcCredentials{username: creds.Username, password: creds.Password}
		}
	}
	if cfg.Username != "" || cfg.Password != "" {
		return &rpcCredentials{username: cfg.Username, password: cfg.Password}
	}
	return nil
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

func (d *kneDUT) DialCLI(context.Context) (binding.StreamClient, error) {
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
	var user, pass string
	creds := newRPCCredentials(d.cfg, d.Name(), d.NodeVendor)
	if creds != nil {
		user, pass = creds.username, creds.password
	}
	log.Infof("Using credentials %s/%s to SSH", user, pass)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			if len(questions) > 0 {
				return []string{pass}, nil
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

func (a *kneATE) DialOTG(ctx context.Context, opts ...grpc.DialOption) (gosnappi.GosnappiApi, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := a.dialGRPC(ctx, "grpc", opts...)
	if err != nil {
		return nil, err
	}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().
		SetClientConnection(conn).
		SetRequestTimeout(30 * time.Second)
	return api, nil
}

func (a *kneATE) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))) // NOLINT
	conn, err := a.dialGRPC(ctx, "gnmi", opts...)
	if err != nil {
		return nil, err
	}
	return gpb.NewGNMIClient(conn), nil
}

func (a *kneATE) dialGRPC(ctx context.Context, serviceName string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	s, err := a.Service(serviceName)
	if err != nil {
		return nil, err
	}
	addr := serviceAddr(s)
	log.Infof("Dialing service %q on ATE %s@%s", serviceName, a.Name(), addr)
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("DialContext(ctx, %s, %v): %w", addr, opts, err)
	}
	return conn, nil
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
