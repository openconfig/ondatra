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
	"testing"

	"golang.org/x/net/context"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"google.golang.org/grpc"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	ppb "github.com/p4lang/p4runtime/go/p4/v1"
)

var (
	dut = &fakebind.DUT{
		DialCLIFn: func(context.Context) (binding.StreamClient, error) {
			return &struct{ binding.StreamClient }{}, nil
		},
		DialConsoleFn: func(context.Context) (binding.StreamClient, error) {
			return &struct{ binding.StreamClient }{}, nil
		},
		DialGNMIFn: func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return &struct{ gpb.GNMIClient }{}, nil
		},
		DialGNOIFn: func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
			return &struct{ binding.GNOIClients }{}, nil
		},
		DialGNSIFn: func(context.Context, ...grpc.DialOption) (binding.GNSIClients, error) {
			return &struct{ binding.GNSIClients }{}, nil
		},
		DialGRIBIFn: func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
			return &struct{ grpb.GRIBIClient }{}, nil
		},
		DialP4RTFn: func(context.Context, ...grpc.DialOption) (ppb.P4RuntimeClient, error) {
			return &struct{ ppb.P4RuntimeClient }{}, nil
		},
	}

	ate = &fakebind.ATE{
		DialIxNetworkFn: func(context.Context) (*binding.IxNetwork, error) {
			return &binding.IxNetwork{}, nil
		},
		DialOTGFn: func(context.Context, ...grpc.DialOption) (gosnappi.GosnappiApi, error) {
			return &struct{ gosnappi.GosnappiApi }{}, nil
		},
		DialGNMIFn: func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return &struct{ gpb.GNMIClient }{}, nil
		},
	}
)

func TestCLI(t *testing.T) {
	if _, err := NewCLI(context.Background(), dut); err != nil {
		t.Fatalf("NewCLI() unexpected error: %v", err)
	}
}

func TestConsole(t *testing.T) {
	if _, err := NewConsole(context.Background(), dut); err != nil {
		t.Fatalf("NewConsole() unexpected error: %v", err)
	}
}

func TestGNMI(t *testing.T) {
	gotNew, err := NewGNMI(context.Background(), dut)
	if err != nil {
		t.Fatalf("NewGNMI() unexpected error: %v", err)
	}
	wantFetch, err := FetchGNMI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGNMI() unexpected error: %v", err)
	}
	gotFetch, err := FetchGNMI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGNMI() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchGNMI() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
	if gotFetch == gotNew {
		t.Errorf("FetchGNMI() unexpected result: got %v, want unique value", gotFetch)
	}
}

func TestGNOI(t *testing.T) {
	gotNew, err := NewGNOI(context.Background(), dut)
	if err != nil {
		t.Fatalf("NewGNOI() unexpected error: %v", err)
	}
	wantFetch, err := FetchGNOI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGNOI() unexpected error: %v", err)
	}
	gotFetch, err := FetchGNOI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGNOI() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchGNOI() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
	if gotFetch == gotNew {
		t.Errorf("FetchGNOI() unexpected result: got %v, want unique value", gotFetch)
	}
}

func TestGNSI(t *testing.T) {
	gotNew, err := NewGNSI(context.Background(), dut)
	if err != nil {
		t.Fatalf("NewGNSI() unexpected error: %v", err)
	}
	wantFetch, err := FetchGNSI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGNSI() unexpected error: %v", err)
	}
	gotFetch, err := FetchGNSI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGNSI() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchGNSI() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
	if gotFetch == gotNew {
		t.Errorf("FetchGNSI() unexpected result: got %v, want unique value", gotFetch)
	}
}

func TestGRIBI(t *testing.T) {
	gotNew, err := NewGRIBI(context.Background(), dut)
	if err != nil {
		t.Fatalf("NewGRIBI() unexpected error: %v", err)
	}
	wantFetch, err := FetchGRIBI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGRIBI() unexpected error: %v", err)
	}
	gotFetch, err := FetchGRIBI(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchGRIBI() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchGRIBI() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
	if gotFetch == gotNew {
		t.Errorf("FetchGRIBI() unexpected result: got %v, want unique value", gotFetch)
	}
}

func TestP4RT(t *testing.T) {
	gotNew, err := NewP4RT(context.Background(), dut)
	if err != nil {
		t.Fatalf("NewP4RT() unexpected error: %v", err)
	}
	wantFetch, err := FetchP4RT(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchP4RT() unexpected error: %v", err)
	}
	gotFetch, err := FetchP4RT(context.Background(), dut)
	if err != nil {
		t.Fatalf("FetchP4RT() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchP4RT() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
	if gotFetch == gotNew {
		t.Errorf("FetchP4RT() unexpected result: got %v, want unique value", gotFetch)
	}
}

func TestIxNetwork(t *testing.T) {
	wantFetch, err := FetchIxNetwork(context.Background(), ate)
	if err != nil {
		t.Fatalf("FetchIxNetwork() unexpected error: %v", err)
	}
	gotFetch, err := FetchIxNetwork(context.Background(), ate)
	if err != nil {
		t.Fatalf("FetchIxNetwork() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchIxNetwork() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
}

func TestOTG(t *testing.T) {
	wantFetch, err := FetchOTG(context.Background(), ate)
	if err != nil {
		t.Fatalf("FetchOTG() unexpected error: %v", err)
	}
	gotFetch, err := FetchOTG(context.Background(), ate)
	if err != nil {
		t.Fatalf("FetchOTG() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchOTG() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
}

func TestOTGGNMI(t *testing.T) {
	wantFetch, err := FetchOTGGNMI(context.Background(), ate)
	if err != nil {
		t.Fatalf("FetchOTGGNMI() unexpected error: %v", err)
	}
	gotFetch, err := FetchOTGGNMI(context.Background(), ate)
	if err != nil {
		t.Fatalf("FetchOTGGNMI() unexpected error: %v", err)
	}
	if gotFetch != wantFetch {
		t.Errorf("FetchOTGGNMI() unexpected result: got %v, want %v", gotFetch, wantFetch)
	}
}
