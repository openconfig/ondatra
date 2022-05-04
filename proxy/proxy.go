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

// Package proxy provide a standard grpc proxy for use with Ondatra.
package proxy

import (
	"golang.org/x/net/context"
	"fmt"
	"net"
	"sync"

	log "github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/openconfig/gnmi/errlist"
	"github.com/openconfig/ondatra/proxy/grpcproxy"
	rpb "github.com/openconfig/ondatra/proxy/proto/reservation"
)

// Dialer adds lower level binding interface for proxied connections
// to the base ondatra.Binding interface. Binding implementors are only expected
// to implement a proxy Dialer if they want to use the proxy as an extension to
// ondatra.Binding interface.
type Dialer interface {
	// DialGRPC creates a client connection to the specified service endpoint
	// on the target. The service will be represented as a gRPC service name.
	// Implementations must append transport security options necessary to reach
	// the server.
	DialGRPC(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error)

	// Resolve will return the resolved devices/ates and related services for the
	// resolved topology.
	Resolve() (*rpb.Reservation, error)
}

// Proxy will create a set of proxy listeners based on the provided testbed.
// The proxy implementation relies on the underlying proxy interface for
// allowing custom dial functions.
type Proxy struct {
	proxies map[string]*proxy
	dialer  Dialer
	resv    *rpb.Reservation
}

type proxy struct {
	srv        grpcproxy.GRPCServer
	targetAddr string
	localAddr  string

	errMu sync.Mutex
	err   error
}

// New will create a new set of proxies for the provided Dialer.
func New(d Dialer) (*Proxy, error) {
	if d == nil {
		return nil, fmt.Errorf("dialer cannot be nil")
	}
	resv, err := d.Resolve()
	if err != nil {
		return nil, err
	}
	p := &Proxy{
		proxies: map[string]*proxy{},
		dialer:  d,
		resv:    resv,
	}
	for k, d := range resv.GetDevices() {
		if err := p.add(k, d); err != nil {
			return nil, err
		}
	}
	if err := grpcproxy.RegisterMetricViews(); err != nil {
		return nil, err
	}
	return p, nil
}

// Endpoint is the current location of a proxy.
type Endpoint struct {
	TargetName string
	TargetAddr string
	Addr       string
}

// Endpoints returns a map of the currently configured proxy endpoints configured
// and their state.
func (p *Proxy) Endpoints() map[string]Endpoint {
	ep := map[string]Endpoint{}
	for k, v := range p.proxies {
		ep[k] = Endpoint{
			TargetName: k,
			TargetAddr: v.targetAddr,
			Addr:       v.localAddr,
		}
	}
	return ep
}

// Stop will shutdown all current proxies.
func (p *Proxy) Stop() error {
	var errs errlist.List
	for _, s := range p.proxies {
		s.srv.Stop()
		func() {
			s.errMu.Lock()
			defer s.errMu.Unlock()
			errs.Add(s.err)
		}()
	}
	return errs.Err()
}

func (p *Proxy) add(key string, d *rpb.ResolvedDevice) error {
	if _, ok := p.proxies[key]; ok {
		log.Infof("gRPC server already registered for device %q ignoring other grpc services", d.GetName())
		return nil
	}
	var s *rpb.Service
	for _, s = range d.GetServices() {
		break
	}
	// Target has no services exported, nothing to do.
	if s == nil || s.GetProxiedGrpc().GetAddress() == "" {
		return nil
	}
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return fmt.Errorf("failed to create listen port for %q: %w", key, err)
	}
	targetAddr := s.GetProxiedGrpc().GetAddress()
	destProvider := func(md metadata.MD, _ string) (string, metadata.MD, error) {
		return targetAddr, md, nil
	}
	srv, err := grpcproxy.NewServer(destProvider, grpcproxy.WithDialer(p.dialer))
	if err != nil {
		return fmt.Errorf("failed to create new gRPC proxy server: %v", err)
	}
	grpcProxy := &proxy{
		srv:        srv,
		localAddr:  lis.Addr().String(),
		targetAddr: s.GetProxiedGrpc().GetAddress(),
	}
	go func() {
		grpcProxy.errMu.Lock()
		defer grpcProxy.errMu.Unlock()
		grpcProxy.err = srv.Serve(lis)
	}()
	p.proxies[key] = grpcProxy
	log.Infof("Added grpc proxy for %q on %s", key, grpcProxy.localAddr)
	return nil
}
