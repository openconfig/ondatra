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
	"fmt"
	"net"
	"net/http"
	"sync"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/openconfig/gnmi/errlist"
	"github.com/openconfig/ondatra/proxy/grpcproxy"
	"github.com/openconfig/ondatra/proxy/httpovergrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	hpb "github.com/openconfig/ondatra/proxy/proto/httpovergrpc"
	rpb "github.com/openconfig/ondatra/proxy/proto/reservation"
)

// HTTPDoCloser provides the required interface for clients to the proxy.
type HTTPDoCloser interface {
	Do(*http.Request) (*http.Response, error)
	Close() error
}

// HTTPCloserClient provides noop wrapped http.Client.
type HTTPCloserClient struct {
	*http.Client
}

// Close is a noop close method.
func (c *HTTPCloserClient) Close() error {
	return nil
}

// WrapHTTPDoCloser wraps a standard http.Client with a noop closer.
func WrapHTTPDoCloser(c *http.Client) HTTPDoCloser {
	return &HTTPCloserClient{c}
}

// Dialer adds lower level binding interface for proxied connections
// to the base ondatra.Binding interface. Binding implementors are only expected
// to implement a proxy Dialer if they want to use the proxy as an extension to
// ondatra.Binding interface.
type Dialer interface {
	// DialGRPC creates a client connection to the specified service endpoint
	// on the target. The service will be represented as a gRPC service name.
	// Implementations must append transport security options necessary to reach
	// the server.
	DialGRPC(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error)

	// Resolve will return the resolved devices/ates and related services for the
	// resolved topology.
	Resolve() (*rpb.Reservation, error)

	// HTTPClient provides an HTTP Client that can connect to the provided target.
	HTTPClient(string) (HTTPDoCloser, error)
}

// Proxy will create a set of proxy listeners based on the provided testbed.
// The proxy implementation relies on the underlying proxy interface for
// allowing custom dial functions.
type Proxy struct {
	gProxies map[string]*grpcProxy
	hProxies map[string]*httpProxy
	dialer   Dialer
	resv     *rpb.Reservation
	sOpts    []grpc.ServerOption
}

type httpProxy struct {
	srv        grpcproxy.GRPCServer
	targetAddr string
	localAddr  string
	c          HTTPDoCloser
	errMu      sync.Mutex
	err        error
}

type grpcProxy struct {
	srv        grpcproxy.GRPCServer
	targetAddr string
	localAddr  string

	errMu sync.Mutex
	err   error
}

// New will create a new set of proxies for the provided Dialer.
func New(d Dialer, sOpts ...grpc.ServerOption) (*Proxy, error) {
	if d == nil {
		return nil, fmt.Errorf("dialer cannot be nil")
	}
	resv, err := d.Resolve()
	if err != nil {
		return nil, err
	}
	log.Infof("Setting up proxies for: %s", resv)
	p := &Proxy{
		gProxies: map[string]*grpcProxy{},
		hProxies: map[string]*httpProxy{},
		dialer:   d,
		resv:     resv,
		sOpts:    sOpts,
	}

	for k, d := range resv.GetDevices() {
		if err := p.addGRPC(k, d); err != nil {
			return nil, err
		}
	}
	for k, a := range resv.GetAtes() {
		if err := p.addHTTP(k, a); err != nil {
			return nil, err
		}
	}
	if err := grpcproxy.RegisterMetricViews(); err != nil {
		return nil, err
	}
	if err := httpovergrpc.RegisterMetricViews(); err != nil {
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
	for k, v := range p.gProxies {
		ep[k] = Endpoint{
			TargetName: k,
			TargetAddr: v.targetAddr,
			Addr:       v.localAddr,
		}
	}
	for k, v := range p.hProxies {
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
	for _, s := range p.gProxies {
		s.srv.Stop()
		func() {
			s.errMu.Lock()
			defer s.errMu.Unlock()
			errs.Add(s.err)
		}()
	}
	for _, s := range p.hProxies {
		s.srv.Stop()
		if err := s.c.Close(); err != nil {
			errs.Add(err)
		}
		func() {
			s.errMu.Lock()
			defer s.errMu.Unlock()
			errs.Add(s.err)
		}()
	}
	return errs.Err()
}

func (p *Proxy) addGRPC(key string, d *rpb.ResolvedDevice) error {
	if _, ok := p.gProxies[key]; ok {
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
	srv, err := grpcproxy.NewServer(destProvider, grpcproxy.WithDialer(p.dialer), grpcproxy.WithServerProvider(
		func(opts ...grpc.ServerOption) (grpcproxy.GRPCServer, error) {
			opts = append(opts, p.sOpts...)
			return grpc.NewServer(opts...), nil
		}))
	if err != nil {
		return fmt.Errorf("failed to create new gRPC proxy server: %v", err)
	}
	gProxy := &grpcProxy{
		srv:        srv,
		localAddr:  lis.Addr().String(),
		targetAddr: s.GetProxiedGrpc().GetAddress(),
	}
	go func() {
		gProxy.errMu.Lock()
		defer gProxy.errMu.Unlock()
		gProxy.err = srv.Serve(lis)
	}()
	p.gProxies[key] = gProxy
	log.Infof("Added grpc proxy for %q on %s", key, gProxy.localAddr)
	return nil
}

func (p *Proxy) addHTTP(key string, d *rpb.ResolvedDevice) error {
	log.Infof("Adding HTTP proxy for %q", key)
	if _, ok := p.hProxies[key]; ok {
		log.Infof("HTTP server already registered for device %q ignoring other grpc services", d.GetName())
		return nil
	}
	s, ok := d.GetServices()["http"]
	if !ok || s.GetHttpOverGrpc().GetAddress() == "" {
		log.Infof("No service found for %q", key)
		return nil
	}
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		return fmt.Errorf("failed to create listen port for %q: %w", key, err)
	}
	targetAddr := s.GetHttpOverGrpc().GetAddress()
	c, err := p.dialer.HTTPClient(targetAddr)
	if err != nil {
		return fmt.Errorf("failed to create new HTTP client: %w", err)
	}
	sProxy := httpovergrpc.New(httpovergrpc.WithClient(c))
	srv := grpc.NewServer(p.sOpts...)
	hpb.RegisterHTTPOverGRPCServer(srv, sProxy)
	hProxy := &httpProxy{
		srv:        srv,
		c:          c,
		localAddr:  lis.Addr().String(),
		targetAddr: s.GetHttpOverGrpc().GetAddress(),
	}
	go func() {
		hProxy.errMu.Lock()
		defer hProxy.errMu.Unlock()
		hProxy.err = srv.Serve(lis)
	}()
	p.hProxies[key] = hProxy
	log.Infof("Added httpovergrpc proxy for %q on %s", key, hProxy.localAddr)
	return nil
}
