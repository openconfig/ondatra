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

// Package gribi performs GRIBI related operations for ONDATRA tests.
package gribi

import (
	"golang.org/x/net/context"
	"sync"

	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/testbed"

	grpb "github.com/openconfig/gribi/v1/proto/service"
)

var (
	mu     sync.Mutex
	gribis = make(map[binding.Device]grpb.GRIBIClient)
)

// NewGRIBI creates a new gRIBI client for the specified Device.
func NewGRIBI(ctx context.Context, dev *binding.DUT) (grpb.GRIBIClient, error) {
	return testbed.Bind().DialGRIBI(ctx, dev, grpc.WithBlock())
}

// FetchGRIBI fetches the gRIBI client for the specified Device.
func FetchGRIBI(ctx context.Context, dev *binding.DUT) (grpb.GRIBIClient, error) {
	mu.Lock()
	defer mu.Unlock()

	c, ok := gribis[dev]
	if !ok {
		var err error
		c, err = NewGRIBI(ctx, dev)
		if err != nil {
			return nil, err
		}
		gribis[dev] = c
	}
	return c, nil
}
