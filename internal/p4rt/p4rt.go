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
	"sync"

	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"

	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var (
	mu    sync.Mutex
	p4rts = make(map[*reservation.DUT]p4pb.P4RuntimeClient)
)

// FetchP4RT fetches the P4RT client for the given DUT.
func FetchP4RT(ctx context.Context, dut *reservation.DUT) (p4pb.P4RuntimeClient, error) {
	mu.Lock()
	defer mu.Unlock()
	p4rt, ok := p4rts[dut]
	if !ok {
		var err error
		p4rt, err = binding.Get().DialP4RT(ctx, dut, grpc.WithBlock())
		if err != nil {
			return nil, err
		}
		p4rts[dut] = p4rt
	}
	return p4rt, nil
}
