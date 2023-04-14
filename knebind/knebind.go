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
	"os"
	"os/user"
	"path/filepath"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/open-traffic-generator/snappi/gosnappi"
	closer "github.com/openconfig/gocloser"
	"github.com/openconfig/kne/topo"
	"github.com/openconfig/kne/topo/node"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind/creds"
	"github.com/openconfig/ondatra/knebind/solver"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	cpb "github.com/openconfig/kne/proto/controller"
	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

// To be stubbed out by unit tests
var (
	userCurrFn = user.Current
)

// New returns a new KNE bind instance.
func New(cfg *Config) (*Bind, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	if cfg.Kubeconfig == "" {
		user, err := userCurrFn()
		if err != nil {
			return nil, err
		}
		cfg.Kubeconfig = filepath.Join(user.HomeDir, ".kube/config")
	}
	top, err := topo.Load(cfg.Topology)
	if err != nil {
		return nil, fmt.Errorf("error loading topology: %w", err)
	}
	tm, err := topo.New(top, topo.WithKubecfg(cfg.Kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("error creating topology manager: %w", err)
	}
	return &Bind{cfg: cfg, tm: tm}, nil
}

// Bind implements the ondatra Binding interface for KNE
type Bind struct {
	binding.Binding
	cfg *Config
	tm  topoManager
}

// Reserve implements the binding Reserve method by finding nodes and links in
// the topology specified in the config file that match the requested testbed.
func (b *Bind) Reserve(ctx context.Context, tb *opb.Testbed, runTime time.Duration, waitTime time.Duration, partial map[string]string) (*binding.Reservation, error) {
	resp, err := b.tm.Show(ctx)
	if err != nil {
		return nil, err
	}
	res, err := solver.Solve(tb, resp.GetTopology(), partial)
	if err != nil {
		return nil, err
	}
	for i, dut := range res.DUTs {
		kdut := &kneDUT{
			ServiceDUT: dut.(*solver.ServiceDUT),
			bind:       b,
		}
		res.DUTs[i] = kdut
		if b.cfg.SkipReset {
			continue
		}
		if err := kdut.resetConfig(ctx); err != nil {
			return nil, err
		}
	}
	for i, ate := range res.ATEs {
		res.ATEs[i] = &kneATE{
			ServiceATE: ate.(*solver.ServiceATE),
			bind:       b,
		}
	}
	return res, nil
}

// Release is a no-op because there's no need to reserve local VMs.
func (b *Bind) Release(context.Context) error {
	return nil
}

type kneDUT struct {
	*solver.ServiceDUT
	bind *Bind
}

func (d *kneDUT) resetConfig(ctx context.Context) error {
	// TODO(alexmasi): Reduce duplication between this and CLI reset implementation.
	name := d.Name()
	switch err := d.bind.tm.ResetCfg(ctx, name); {
	case status.Code(err) == codes.Unimplemented:
		log.Warningf("Node %q does not support config reset, skipping", name)
	case err != nil:
		return err
	}
	bp, err := fileRelative(d.bind.cfg.Topology)
	if err != nil {
		return fmt.Errorf("failed to find relative path for topology: %v", err)
	}
	node := d.bind.tm.Nodes()[name]
	cd := node.GetProto().GetConfig().GetConfigData()
	if cd == nil {
		log.Infof("Skipping node %q no initial config found", name)
		return nil
	}
	var b []byte
	switch v := cd.(type) {
	case *tpb.Config_Data:
		b = v.Data
	case *tpb.Config_File:
		cPath := v.File
		if !filepath.IsAbs(cPath) {
			cPath = filepath.Join(bp, cPath)
		}
		log.Infof("Pushing configuration %q to %q", cPath, name)
		var err error
		b, err = os.ReadFile(cPath)
		if err != nil {
			return err
		}
	}
	return d.pushConfig(ctx, b)
}

func (d *kneDUT) pushConfig(ctx context.Context, config []byte) error {
	switch err := d.bind.tm.ConfigPush(ctx, d.Name(), bytes.NewBuffer(config)); {
	case status.Code(err) == codes.Unimplemented:
		log.Warningf("Node %q does not support pushing config, skipping", d.Name())
	case err != nil:
		return err
	}
	return nil
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
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))) // NOLINT
	creds := d.newRPCCredentials()
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
	*creds.UserPass
}

func (r *rpcCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": r.Username,
		"password": r.Password,
	}, nil
}

func (r *rpcCredentials) RequireTransportSecurity() bool {
	return true
}

// newRPCCredentials determines the correct credentials used to access a node via rpc
// from a given node name, node vendor, and knebind config. The precedence order for determining
// the credentials is as follows:
//
// 1. credential provided for a specific node by name from the knebind config
// 2. credential provided for a vendor of the node from the knebind config
// 3. credential from the default username and password fields from the knebind config
// 4. no credentials
func (d *kneDUT) newRPCCredentials() *rpcCredentials {
	cfg := d.bind.cfg
	if userPass := cfg.Credentials.Lookup(d.Name(), d.NodeVendor); userPass != nil {
		return &rpcCredentials{userPass}
	}
	// TODO(team): Deprecate username and password fields.
	if cfg.Username != "" {
		return &rpcCredentials{&creds.UserPass{Username: cfg.Username, Password: cfg.Password}}
	}
	return nil
}

func (d *kneDUT) PushConfig(ctx context.Context, config string, reset bool) error {
	if reset {
		if err := d.resetConfig(ctx); err != nil {
			return err
		}
	}
	return d.pushConfig(ctx, []byte(config))
}

func (d *kneDUT) DialCLI(context.Context) (binding.StreamClient, error) {
	return &kneCLI{dut: d}, nil
}

type kneCLI struct {
	*binding.AbstractStreamClient
	dut *kneDUT
}

func (c *kneCLI) SendCommand(_ context.Context, cmd string) (_ string, rerr error) {
	s, err := c.dut.Service("ssh")
	if err != nil {
		return "", err
	}
	addr := serviceAddr(s)
	userPass := c.dut.newRPCCredentials()
	if userPass == nil {
		return "", errors.New("SendCommand requires node credentials be provided")
	}
	log.Infof("Using credentials %v to SSH", userPass)
	cfg := &ssh.ClientConfig{
		User: userPass.Username,
		Auth: []ssh.AuthMethod{ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			if len(questions) > 0 {
				return []string{userPass.Password}, nil
			}
			return nil, nil
		})},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
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

func (c *kneCLI) Close() error {
	return nil
}

type kneATE struct {
	*solver.ServiceATE
	bind *Bind
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

func fileRelative(p string) (string, error) {
	bp, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}
	return filepath.Dir(bp), nil
}

type topoManager interface {
	Nodes() map[string]node.Node
	Show(context.Context) (*cpb.ShowTopologyResponse, error)
	ConfigPush(context.Context, string, io.Reader) error
	ResetCfg(context.Context, string) error
}
