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

// Package cli provides implementation on connecting to a DUT's cli.
package cli

import (
	"golang.org/x/net/context"
	"sync"

	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
)

var (
	mu   sync.Mutex
	clis = map[*reservation.DUT]binding.StreamClient{}
)

// FetchCLI will return a CLI client for interacting with the DUT.
func FetchCLI(ctx context.Context, dut *reservation.DUT) (binding.StreamClient, error) {
	mu.Lock()
	defer mu.Unlock()
	c, ok := clis[dut]
	if !ok {
		var err error
		c, err = binding.Get().DialCLI(ctx, dut, grpc.WithBlock())
		if err != nil {
			return nil, err
		}
		clis[dut] = c
	}
	return c, nil
}
