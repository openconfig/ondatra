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
	"golang.org/x/net/context"
	"fmt"
	"io"
	"strings"

	log "github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
)

var (
	clientStreamDesc = &grpc.StreamDesc{
		ServerStreams: true,
		ClientStreams: true,
	}
)

// DestProviderFn computes a target address for this reverse proxy for every new proxy request.
// Implementations can compute an address based on input of incoming request method and headers.
type DestProviderFn func(metadata.MD, string) (string, metadata.MD, error)

// GRPCDialer is interface to dial a gRPC connection.
type GRPCDialer interface {
	DialGRPC(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error)
}

type proxy struct {
	destProvider DestProviderFn
	grpcDialer   GRPCDialer
	targetOpts   []grpc.DialOption
	keyMethod    tag.Key
	targetConn   *grpc.ClientConn
}

type defaultGRPCDialer struct{}

func (d defaultGRPCDialer) DialGRPC(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, target, opts...)
}

// NewWithGRPCDialer is similar to NewServer, but user can provide a custom gRPC dialer.
func NewWithGRPCDialer(dp DestProviderFn, dialer GRPCDialer, tOpts []grpc.DialOption, sOpts ...grpc.ServerOption) *grpc.Server {
	p := &proxy{
		destProvider: dp,
		grpcDialer:   dialer,
		targetOpts:   tOpts,
	}
	opts := append(
		sOpts,
		grpc.ForceServerCodec(codec()),
		grpc.UnknownServiceHandler(p.instrumentedStreamHandler))
	return grpc.NewServer(opts...)
}

func (p *proxy) dialTarget(ctx context.Context, method string) (context.Context, *grpc.ClientConn, error) {
	// Intercept header.
	hd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, status.Errorf(codes.InvalidArgument, "proxy_error: missing metadata")
	}
	if p.targetConn != nil {
		destCtx := metadata.NewOutgoingContext(ctx, hd)
		return destCtx, p.targetConn, nil
	}
	dest, destHd, err := p.destProvider(hd, method)
	if err != nil {
		return nil, nil, err
	}
	destCtx := metadata.NewOutgoingContext(ctx, destHd)
	opts := append(p.targetOpts, grpc.WithDefaultCallOptions(grpc.ForceCodec(codec())))
	c, err := p.grpcDialer.DialGRPC(ctx, dest, opts...)
	if err != nil {
		return nil, nil, status.Errorf(codes.Unavailable, fmt.Sprintf("proxy_error: cannot connect to target backend: %v", err))
	}
	return destCtx, c, nil
}

func (p *proxy) instrumentedStreamHandler(srv interface{}, stream grpc.ServerStream) error {
	ctx := stream.Context()
	method, ok := grpc.MethodFromServerStream(stream)
	if !strings.HasPrefix(method, "/") {
		log.Warning("The method name doesn't start with '/', trying to add it.")
		method = "/" + method
	}
	if !ok {
		err := status.Errorf(codes.Internal, "proxy_error: no method name on ServerStream Context")
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
	destCtx, destConn, err := p.dialTarget(ctx, method)
	if err != nil {
		return err
	}
	if p.targetConn == nil {
		defer destConn.Close()
	}

	// Establish a stream to target server.
	destCtx, clientCancel := context.WithCancel(destCtx)
	defer clientCancel()
	clientStream, err := grpc.NewClientStream(destCtx, clientStreamDesc, destConn, method, grpc.WaitForReady(true))
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("proxy_error: %v", err))
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
	p := &proto{}
	// Headers already transferred as part of grpc client connection context.
	// Only body transfer needed.
	for i := 0; ; i++ {
		if err := src.RecvMsg(p); err != nil {
			if err == io.EOF {
				if dErr := dst.CloseSend(); dErr != nil {
					return status.Errorf(codes.Internal, fmt.Sprintf("proxy_error: %v", dErr))
				}
				return nil
			}
			return status.Errorf(codes.Internal, fmt.Sprintf("proxy_error: %v", err))
		}
		if err := dst.SendMsg(p); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("proxy_error: %v", err))
		}
	}
}

func forwardClientToServer(dst grpc.ServerStream, src grpc.ClientStream) error {
	p := &proto{}

	// Transfer Headers.
	md, err := src.Header()
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("proxy_error: %v", err))
	}
	if err := dst.SendHeader(md); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("proxy_error: %v", err))
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
