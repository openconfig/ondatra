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

package rawapis

import (
	"fmt"
	"sync/atomic"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// withUnaryAnnotateErrors annotates every gRPC error returned by an RPC with
// the RPC request message.
func withUnaryAnnotateErrors() grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(
		func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc, opts...)
			return maybeAnnotateErr(err, req)
		})
}

// withStreamAnnotateErrors annotates every gRPC error returned by an RPC with
// the preceding RPC request message.
func withStreamAnnotateErrors() grpc.DialOption {
	return grpc.WithChainStreamInterceptor(
		func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			client, err := streamer(ctx, desc, cc, method, opts...)
			if client != nil {
				client = &annotateErrClient{ClientStream: client}
			}
			return client, err
		})
}

type annotateErrClient struct {
	grpc.ClientStream
	req atomic.Value
}

func (c *annotateErrClient) SendMsg(m any) error {
	c.req.Store(m)
	err := c.ClientStream.SendMsg(m)
	return maybeAnnotateErr(err, m)
}

func (c *annotateErrClient) RecvMsg(m any) error {
	err := c.ClientStream.RecvMsg(m)
	return maybeAnnotateErr(err, c.req.Load())
}

func maybeAnnotateErr(err error, req any) error {
	// Do not annotate the error if:
	// 1. err is nil: there is no error to annotate
	// 2. req is nil: there is no request to annotate with
	// 3. err is not a grpc error: sentinel errors like io.EOF must be preserved
	if err != nil && req != nil {
		if st, ok := status.FromError(err); ok {
			var reqText string
			if pm, ok := req.(proto.Message); ok {
				reqText = prototext.Format(pm)
			} else {
				reqText = fmt.Sprint(req)
			}
			pb := st.Proto()
			pb.Message = fmt.Sprintf("error on request {\n%s}: %s", reqText, pb.Message)
			return status.FromProto(pb).Err()
		}
	}
	return err
}
