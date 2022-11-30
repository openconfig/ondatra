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

// Package grpcutil contains gRPC utilities useful to binding implementations.
package grpcutil

import (
	"io"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var emptyCancel context.CancelFunc = func() {}

// WithDefaultTimeout returns context.WithTimeout(ctx, timeout) if ctx has no
// deadline; otherwise returns the context as-is and an empty cancel function.
func WithDefaultTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if _, ok := ctx.Deadline(); !ok {
		return context.WithTimeout(ctx, timeout)
	}
	return ctx, emptyCancel
}

// WithUnaryDefaultTimeout returns a dial option that intercepts unary requests
// and imposes the specified context timeout on the request if it does not
// already have a deadline.
func WithUnaryDefaultTimeout(timeout time.Duration) grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(
		func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			ctx, cancel := WithDefaultTimeout(ctx, timeout)
			defer cancel()
			return invoker(ctx, method, req, reply, cc, opts...)
		},
	)
}

// WithStreamDefaultTimeout returns a dial option that intercepts stream
// requests and imposes the specified context timeout on the request if it does
// not already have a deadline.
func WithStreamDefaultTimeout(timeout time.Duration) grpc.DialOption {
	return grpc.WithChainStreamInterceptor(
		func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			ctx, cancel := WithDefaultTimeout(ctx, timeout)
			cs, err := streamer(ctx, desc, cc, method, opts...)
			if err != nil {
				cancel()
				return nil, err
			}
			// Unfortunately gRPC provides no simple way to ensure a child context
			// created in an interceptor is cancelled.
			// The current recommendation is to wrap the client stream and cancel it
			// when RecvMsg returns a non-nil error, or when Header or SendMsg return
			// a non-nil, non-EOF error. See https://github.com/grpc/grpc-go/issues/1717
			return &cancelClientStream{ClientStream: cs, cancel: cancel}, nil
		},
	)
}

type cancelClientStream struct {
	grpc.ClientStream
	cancel context.CancelFunc
}

func (cs *cancelClientStream) Header() (metadata.MD, error) {
	md, err := cs.ClientStream.Header()
	if err != nil && err != io.EOF {
		cs.cancel()
	}
	return md, err
}

func (cs *cancelClientStream) SendMsg(m any) error {
	err := cs.ClientStream.SendMsg(m)
	if err != nil && err != io.EOF {
		cs.cancel()
	}
	return err
}

func (cs *cancelClientStream) RecvMsg(m any) error {
	err := cs.ClientStream.RecvMsg(m)
	if err != nil {
		cs.cancel()
	}
	return err
}
