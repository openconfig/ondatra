// Copyright 2021 Cisco System
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

// Package griibi controls gribi operation  for ONDATRA tests.


package gribi

import (
	"sync"

	"golang.org/x/net/context"
	gribi "github.com/openconfig/gribi/v1/proto/service"
	fgribi "github.com/openconfig/gribigo/fluent"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"google.golang.org/grpc"
)

const (
	maxUint64 = 1<<64 - 1
)

var (
	mu          sync.Mutex
	gribiMaster = make(map[*reservation.DUT]*binding.GRIBIClient)
	gribiSecond  = make(map[*reservation.DUT]*binding.GRIBIClient)
	electionID  = make(map[*reservation.DUT]gribi.Uint128)
)

func GetElectionID(ctx context.Context, dut *reservation.DUT) (low, high uint64) {
	mu.Lock()
	defer mu.Unlock()
	_, ok := electionID[dut]
	if !ok {
		electionID[dut] = gribi.Uint128{Low: 0,
			High: 0,
		}
	}
	return electionID[dut].Low, electionID[dut].High
}

func SetElectionID(ctx context.Context, dut *reservation.DUT, low, high uint64) {
	mu.Lock()
	defer mu.Unlock()
	electionID[dut] = gribi.Uint128{Low: low,
		High: high,
	}
}

func IncElectionID(ctx context.Context, dut *reservation.DUT) (high, low uint64) {
	mu.Lock()
	defer mu.Unlock()
	_, ok := electionID[dut]
	if !ok {
		electionID[dut] = gribi.Uint128{Low: 1,
			High: 0,
		}
	} else if electionID[dut].Low == maxUint64 {
		electionID[dut] = gribi.Uint128{Low: 0,
			High: electionID[dut].High + 1,
		}
	} else {
		electionID[dut] = gribi.Uint128{Low: electionID[dut].Low + 1,
			High: electionID[dut].High,
		}
	}
	return electionID[dut].High, electionID[dut].Low
}

// NewGRIBI creates a new gRIBI client for the specified DUT.
func NewGRIBIClient(ctx context.Context, dut *reservation.DUT) (*binding.GRIBIClient, error) {
	return binding.Get().DialGRIBI(ctx, dut, grpc.WithBlock())
}

// fetchGRIBI fetches a cached gRIBI client for the given DUT.
func FetchMasterGRIBI(ctx context.Context, dut *reservation.DUT) (*fgribi.GRIBIClient, error) {
	mu.Lock()
	defer mu.Unlock()
	gribiClient, ok := gribiMaster[dut]
	if !ok {
		var err error
		gribiClient, err = NewGRIBIClient(ctx, dut)
		if err != nil {
			return nil, err
		}
		gribiMaster[dut] = gribiClient
	}
	return gribiClient.FluentAPIHandle, nil
}

func ResetMasterGRIBI(ctx context.Context, dut *reservation.DUT) (*fgribi.GRIBIClient, error) {
	mu.Lock()
	defer mu.Unlock()
	gribiClient, ok := gribiMaster[dut]
	if ok {
		gribiClient.GRPCClient.Close()
		gribiMaster[dut] = nil
	}
	var err error
	gribiClient, err = NewGRIBIClient(ctx, dut)
	if err != nil {
		return nil, err
	}
	gribiMaster[dut] = gribiClient
	return gribiClient.FluentAPIHandle, nil
}

// fetchGRIBI fetches a cached gRIBI client for the given DUT.
func FetchSecondGRIBI(ctx context.Context, dut *reservation.DUT) (*fgribi.GRIBIClient, error) {
	mu.Lock()
	defer mu.Unlock()
	gribiClient, ok := gribiSecond[dut]
	if !ok {
		var err error
		gribiClient, err = NewGRIBIClient(ctx, dut)
		if err != nil {
			return nil, err
		}
		gribiSecond[dut] = gribiClient
	}
	return gribiClient.FluentAPIHandle, nil
}

func ResetSecondGRIBI(ctx context.Context, dut *reservation.DUT) (*fgribi.GRIBIClient, error) {
	mu.Lock()
	defer mu.Unlock()
	gribiClient, ok := gribiSecond[dut]
	if ok {
		gribiClient.GRPCClient.Close()
		gribiSecond[dut] = nil
	}
	var err error
	gribiClient, err = NewGRIBIClient(ctx, dut)
	if err != nil {
		return nil, err
	}
	gribiSecond[dut] = gribiClient
	return gribiClient.FluentAPIHandle, nil
}

func CloseSecondGRIBI(ctx context.Context, dut *reservation.DUT) {
	mu.Lock()
	defer mu.Unlock()
	gribiClient, ok := gribiSecond[dut]
	if ok {
		gribiClient.GRPCClient.Close()
		delete(gribiSecond, dut)
	}
}

func CloseMasterGRIBI(ctx context.Context, dut *reservation.DUT) {
	mu.Lock()
	defer mu.Unlock()
	gribiClient, ok := gribiMaster[dut]
	if ok {
		gribiClient.GRPCClient.Close()
		delete(gribiMaster, dut)
	}
}
