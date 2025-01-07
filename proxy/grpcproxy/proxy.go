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

// Package grpcproxy provides a single proxy grpc server which can be used to
// tranparently forward both unary and bidirectional streams across different
// gRPC connections. The main use case was to provide connections across
// auth domains.
package grpcproxy

import (
	"fmt"
	"io"
	"net"
	"strings"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	lru "github.com/hashicorp/golang-lru"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	clientStreamDesc = &grpc.StreamDesc{
		ServerStreams: true,
		ClientStreams: true,
	}
)

// DestProviderFn computes a target address for this reverse proxy for every new proxy request.
// Implementations can compute an address based on input of incoming request method and headers.
// Return values should be the next target address and the metadata to pass to the target.
type DestProviderFn func(metadata.MD, string) (string, metadata.MD, error)

// GRPCDialer is interface to dial a gRPC connection.
type GRPCDialer interface {
	DialGRPC(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error)
}

type defaultGRPCDialer struct{}

func (d defaultGRPCDialer) DialGRPC(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	log.Infof("DialContext(%s)", target)
	conn, err := grpc.DialContext(ctx, target, opts...)
	log.Infof("DialContext(%s) returns %s", target, err)
	return conn, err
}

// GRPCServer is an interface to server gRPC.
type GRPCServer interface {
	Serve(net.Listener) error
	Stop()
}

// ServerProviderFn creates a GRPCServer from server options.
type ServerProviderFn func(...grpc.ServerOption) (GRPCServer, error)

func defaultServerProviderFn(opts ...grpc.ServerOption) (GRPCServer, error) {
	return grpc.NewServer(opts...), nil
}

type proxy struct {
	destProvider   DestProviderFn
	serverProvider ServerProviderFn
	grpcDialer     GRPCDialer
	cacheKeyFn     CacheKeyFn
	cacheEvictFn   CacheEvictFn
	cacheSize      int
	cache          *lru.Cache
}

// Option configures the proxy server.
type Option func(p *proxy)

// WithDialer sets the proxy server GRPCDialer implementation.
func WithDialer(d GRPCDialer) Option {
	return func(p *proxy) {
		p.grpcDialer = d
	}
}

// WithLRUCache sets the cache size and the keyFunction for caching GRPC client connections.
func WithLRUCache(keyFn CacheKeyFn, evictFn CacheEvictFn, cacheSize int) Option {
	return func(p *proxy) {
		p.cacheKeyFn = keyFn
		p.cacheEvictFn = evictFn
		p.cacheSize = cacheSize
	}
}

// WithServerProvider sets the proxy server GRPCServer provider implementation.
// REMINDER: if you are providing this option make sure you properly handle the
// appending of opts from the caller in your implementation as the server will
// append additional opts as part of it startup.
// Specifically grpc.ForceServerCodec and grpc.UnknownServiceHandler.
func WithServerProvider(sp ServerProviderFn) Option {
	return func(p *proxy) {
		p.serverProvider = sp
	}
}

// NewServer returns a GRPCServer proxy that handles unknown services.
func NewServer(dp DestProviderFn, opts ...Option) (GRPCServer, error) {
	p := &proxy{destProvider: dp}
	for _, opt := range opts {
		opt(p)
	}
	if p.grpcDialer == nil {
		p.grpcDialer = defaultGRPCDialer{}
	}
	if p.serverProvider == nil {
		p.serverProvider = defaultServerProviderFn
	}
	if err := p.initCache(); err != nil {
		return nil, err
	}
	return p.serverProvider(
		grpc.ForceServerCodec(codec()),
		grpc.UnknownServiceHandler(p.instrumentedStreamHandler),
	)
}

// dialTarget returns the outgoing context, the grpc client connection,
// a boolean suggesting whether the client connection was cached, and an error if any.
func (p *proxy) dialTarget(ctx context.Context, method string) (context.Context, *grpc.ClientConn, bool, error) {
	// Intercept header.
	name := fmt.Sprintf("dialTarget(%s)", method)
	log.Info(name)
	hd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Infof("%s proxy_error: missing metadata", name)
		return nil, nil, false, status.Error(codes.InvalidArgument, "proxy_error: missing metadata")
	}
	dest, destHd, err := p.destProvider(hd, method)
	if err != nil {
		log.Infof("%s destProvider: %v", name, err)
		return nil, nil, false, err
	}
	destCtx := metadata.NewOutgoingContext(ctx, destHd)
	c, ck, err := p.connFromCache(hd, dest)
	if err != nil {
		return nil, nil, false, err
	}
	if c != nil {
		log.InfoContextf(ctx, "%s %s connection from cache", name, dest)
		return destCtx, c, true, nil
	}
	c, err = p.grpcDialer.DialGRPC(ctx, dest, grpc.WithDefaultCallOptions(grpc.ForceCodec(codec())))
	if err != nil {
		log.Infof("%s proxy_error: cannot connect to target backend %s: %v", name, dest, err)
		return nil, nil, false, status.Errorf(codes.Unavailable, "proxy_error: cannot connect to target backend: %v", err)
	}
	log.InfoContextf(ctx, "%s %s dial successful", name, dest)
	if added := p.addToCache(c, ck); added {
		log.InfoContextf(ctx, "%s %s added to cache", name, dest)
		return destCtx, c, true, nil
	}
	return destCtx, c, false, nil
}

func (p *proxy) instrumentedStreamHandler(srv any, stream grpc.ServerStream) error {
	ctx := stream.Context()
	method, ok := grpc.MethodFromServerStream(stream)
	if !strings.HasPrefix(method, "/") {
		log.Warning("The method name doesn't start with '/', trying to add it.")
		method = "/" + method
	}
	if !ok {
		err := status.Error(codes.Internal, "proxy_error: no method name on ServerStream Context")
		mut := []tag.Mutator{tag.Upsert(errorKey, err.Error())}
		stats.RecordWithTags(ctx, mut, mProxyErrors.M(1))
		return err
	}

	// Start a new span
	ctx, span := trace.StartSpan(ctx, method)
	defer span.End()

	ctx, _ = tag.New(ctx, tag.Upsert(methodKey, method))
	stats.Record(ctx, mNumRequests.M(1))
	err := p.streamHandler(ctx, method, stream)
	if err != nil {
		mut := []tag.Mutator{tag.Upsert(errorKey, err.Error())}
		if strings.Contains(err.Error(), "proxy_error:") {
			stats.RecordWithTags(ctx, mut, mProxyErrors.M(1))
		} else {
			stats.RecordWithTags(ctx, mut, mClientErrors.M(1))
		}
	}
	return err
}

func (p *proxy) streamHandler(ctx context.Context, method string, stream grpc.ServerStream) error {
	// Establish connection to dest server.
	destCtx, destConn, fromCache, err := p.dialTarget(ctx, method)
	if err != nil {
		return err
	}
	if !fromCache {
		defer destConn.Close()
	}

	// Establish a stream to target server.
	destCtx, clientCancel := context.WithCancel(destCtx)
	defer clientCancel()
	clientStream, err := grpc.NewClientStream(destCtx, clientStreamDesc, destConn, method, grpc.WaitForReady(true))
	if err != nil {
		return status.Errorf(codes.Internal, "proxy_error: %v", err)
	}
	return p.bidiStream(clientStream, stream)
}

func (p *proxy) bidiStream(cs grpc.ClientStream, ss grpc.ServerStream) error {
	c2sErr := make(chan error, 1)
	s2cErr := make(chan error, 1)
	go func() { c2sErr <- forwardClientToServer(ss, cs) }()
	go func() { s2cErr <- forwardServerToClient(cs, ss) }()
	for {
		select {
		// If ClientToServer ends, we always want to return instead of waiting for
		// ServerToClient to finish also. If target server is not sending us any more
		// messages, we don't want to send it any new messages either.
		case err := <-c2sErr:
			return err
		// If ServerToClient ends normally, ClientToServer may still be active, we want to
		// wait for it to finish. Otherwise exit prematurely with error.
		case err := <-s2cErr:
			if err != nil {
				return err
			}
		}
	}
}

func forwardServerToClient(dst grpc.ClientStream, src grpc.ServerStream) error {
	p := &rawProto{}
	// Headers already transferred as part of grpc client connection context.
	// Only body transfer needed.
	for i := 0; ; i++ {
		if err := src.RecvMsg(p); err != nil {
			if err == io.EOF {
				if dErr := dst.CloseSend(); dErr != nil {
					return status.Errorf(codes.Internal, "proxy_error: %v", dErr)
				}
				return nil
			}
			return status.Errorf(codes.Internal, "proxy_error: %v", err)
		}
		if err := dst.SendMsg(p); err != nil {
			return status.Errorf(codes.Internal, "proxy_error: %v", err)
		}
	}
}

func forwardClientToServer(dst grpc.ServerStream, src grpc.ClientStream) error {
	p := &rawProto{}
	// Transfer Headers.
	md, err := src.Header()
	if err != nil {
		return status.Errorf(codes.Internal, "proxy_error: %v", err)
	}
	if err := dst.SendHeader(md); err != nil {
		return status.Errorf(codes.Internal, "proxy_error: %v", err)
	}

	// Transfer body.
	for i := 0; ; i++ {
		if err := src.RecvMsg(p); err != nil {
			if err == io.EOF {
				dst.SetTrailer(src.Trailer())
				return nil
			}
			return err
		}
		if err := dst.SendMsg(p); err != nil {
			return err
		}
	}
}
