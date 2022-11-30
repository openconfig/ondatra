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

package grpcutil

import (
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding/grpcutil/testservice"

	tgrpcpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
	tpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
)

func TestDefaultTimeout(t *testing.T) {
	const (
		defaultTimeout = time.Minute
		givenTimeout   = 10 * defaultTimeout
	)
	spy := new(spyServer)
	ctx := context.Background()
	serv := testservice.Start(ctx, t, spy,
		WithUnaryDefaultTimeout(defaultTimeout),
		WithStreamDefaultTimeout(defaultTimeout))
	defer serv.Stop()

	t.Run("unary default deadline", func(t *testing.T) {
		spy.deadline = time.Time{}
		wantDeadline := time.Now().Add(defaultTimeout)
		serv.MustSendUnary(ctx, t, "")
		if gotDeadline := spy.deadline; !timeApproxEq(gotDeadline, wantDeadline) {
			t.Errorf("WithUnaryDefaultTimeout() got deadline %v, want approximately %v", gotDeadline, wantDeadline)
		}
	})

	t.Run("unary given deadline", func(t *testing.T) {
		spy.deadline = time.Time{}
		wantDeadline := time.Now().Add(givenTimeout)
		ctx, cancel := context.WithTimeout(ctx, givenTimeout)
		defer cancel()
		serv.MustSendUnary(ctx, t, "")
		if gotDeadline := spy.deadline; !timeApproxEq(gotDeadline, wantDeadline) {
			t.Errorf("WithUnaryDefaultTimeout() got deadline %v, want approximately %v", gotDeadline, wantDeadline)
		}
	})

	t.Run("stream default deadline", func(t *testing.T) {
		spy.deadline = time.Time{}
		wantDeadline := time.Now().Add(defaultTimeout)
		streamClient := serv.MustSendStream(ctx, t)
		for streamClient.MustRecv(t) {
		}
		if gotDeadline := spy.deadline; !timeApproxEq(gotDeadline, wantDeadline) {
			t.Errorf("WithStreamDefaultTimeout() got deadline %v, want approximately %v", gotDeadline, wantDeadline)
		}
	})

	t.Run("stream given deadline", func(t *testing.T) {
		spy.deadline = time.Time{}
		wantDeadline := time.Now().Add(givenTimeout)
		ctx, cancel := context.WithTimeout(ctx, givenTimeout)
		defer cancel()
		streamClient := serv.MustSendStream(ctx, t)
		for streamClient.MustRecv(t) {
		}
		if gotDeadline := spy.deadline; !timeApproxEq(gotDeadline, wantDeadline) {
			t.Errorf("WithStreamDefaultTimeout() got deadline %v, want approximately %v", gotDeadline, wantDeadline)
		}
	})
}

type spyServer struct {
	tgrpcpb.UnimplementedTestServer
	deadline time.Time
}

func (s *spyServer) SendUnary(ctx context.Context, _ *tpb.TestRequest) (*tpb.TestResponse, error) {
	if deadline, ok := ctx.Deadline(); ok {
		s.deadline = deadline
	}
	return new(tpb.TestResponse), nil
}

func (s *spyServer) SendStream(stream tgrpcpb.Test_SendStreamServer) error {
	if deadline, ok := stream.Context().Deadline(); ok {
		s.deadline = deadline
	}
	return nil
}

func timeApproxEq(t1, t2 time.Time) bool {
	return t1.Sub(t2).Abs().Seconds() < 1
}
