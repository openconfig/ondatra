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

// Package p4rt performs P4RT related operations for ONDATRA tests.
package p4rt

import (
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/testbed"

	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

// NewP4RT creates a P4RT client for the specified DUT.
func NewP4RT(ctx context.Context, dut *binding.DUT) (p4pb.P4RuntimeClient, error) {
	return testbed.Bind().DialP4RT(ctx, dut, grpc.WithBlock())
}
