// Copyright 2023 Google LLC
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

package gnoi

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/gnoigo"
	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
	"google.golang.org/grpc"
)

type fakeOp struct {
	gnoigo.Operation[fakeResult]
}

type fakeResult struct{}

func TestExecute(t *testing.T) {
	fakebind.Setup().WithReservation(
		&binding.Reservation{DUTs: map[string]binding.DUT{"dut": &fakebind.DUT{
			AbstractDUT: &binding.AbstractDUT{&binding.Dims{}},
			DialGNOIFn: func(context.Context, ...grpc.DialOption) (gnoigo.Clients, error) {
				return nil, nil
			},
		}}},
	)
	dut := ondatra.DUT(t, "dut")
	op := &fakeOp{}

	tests := []struct {
		desc    string
		want    fakeResult
		wantErr string
	}{
		{
			desc:    "Execute with error",
			wantErr: "rpc err",
		},
		{
			desc: "Execute no error",
			want: fakeResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			stubExecuteFn = func(context.Context, gnoigo.Clients, any) (any, error) {
				if tt.wantErr == "" {
					return tt.want, nil
				}
				return tt.want, errors.New(tt.wantErr)
			}
			var got fakeResult
			gotErr := testt.CaptureFatal(t, func(t testing.TB) {
				got = Execute(t, dut, op)
			})
			if got != tt.want {
				t.Errorf("Execute() got unexpected response %v, want %v", got, tt.want)
			}
			if (gotErr == nil) != (tt.wantErr == "") || (gotErr != nil && !strings.Contains(*gotErr, tt.wantErr)) {
				t.Fatalf("Execute() got unexpected error %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
