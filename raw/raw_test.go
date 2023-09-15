// Copyright 2019 Google LLC
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

package raw

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
	"google.golang.org/grpc"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var (
	dut     = &fakebind.DUT{AbstractDUT: &binding.AbstractDUT{&binding.Dims{Name: "fakeDUT"}}}
	ate     = &fakebind.ATE{AbstractATE: &binding.AbstractATE{&binding.Dims{Name: "fakeATE"}}}
	dutAPIs = NewDUTAPIs(dut)
	ateAPIs = NewATEAPIs(ate)
)

func TestGNMI(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad gnmi"
		dut.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			dutAPIs.GNMI(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("GNMI(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ gpb.GNMIClient }{}
		dut.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return want, nil
		}
		if got := dutAPIs.GNMI(t); got != want {
			t.Errorf("GNMI(t) got %v, want %v", got, want)
		}
	})
}

func TestGNOI(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad gnoi"
		dut.DialGNOIFn = func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			dutAPIs.GNOI(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("GNOI(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ binding.GNOIClients }{}
		dut.DialGNOIFn = func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
			return want, nil
		}
		if got := dutAPIs.GNOI(t); got != want {
			t.Errorf("GNOI(t) got %v, want %v", got, want)
		}
	})
}

func TestGNSI(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad gnsi"
		dut.DialGNSIFn = func(context.Context, ...grpc.DialOption) (binding.GNSIClients, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			dutAPIs.GNSI(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("GNSI(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ binding.GNSIClients }{}
		dut.DialGNSIFn = func(context.Context, ...grpc.DialOption) (binding.GNSIClients, error) {
			return want, nil
		}
		if got := dutAPIs.GNSI(t); got != want {
			t.Errorf("GNSI(t) got %v, want %v", got, want)
		}
	})
}

func TestGRIBI(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad gribi"
		dut.DialGRIBIFn = func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			dutAPIs.GRIBI(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("GRIBI(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ grpb.GRIBIClient }{}
		dut.DialGRIBIFn = func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
			return want, nil
		}
		if got := dutAPIs.GRIBI(t); got != want {
			t.Errorf("GRIBI(t) got %v, want %v", got, want)
		}
	})
}

func TestP4RT(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad p4rt"
		dut.DialP4RTFn = func(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			dutAPIs.P4RT(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("P4RT(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ p4pb.P4RuntimeClient }{}
		dut.DialP4RTFn = func(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
			return want, nil
		}
		if got := dutAPIs.P4RT(t); got != want {
			t.Errorf("P4RT(t) got %v, want %v", got, want)
		}
	})
}

func TestCLI(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad cli"
		dut.DialCLIFn = func(context.Context) (binding.CLIClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			dutAPIs.CLI(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("CLI(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ binding.CLIClient }{}
		dut.DialCLIFn = func(context.Context) (binding.CLIClient, error) {
			return want, nil
		}
		if got := dutAPIs.CLI(t); want != got {
			t.Errorf("CLI(t) got %v, want %v", got, want)
		}
	})
}

func TestConsole(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad cli"
		dut.DialConsoleFn = func(context.Context) (binding.ConsoleClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			dutAPIs.Console(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("Console(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ binding.ConsoleClient }{}
		dut.DialConsoleFn = func(context.Context) (binding.ConsoleClient, error) {
			return want, nil
		}
		if got := dutAPIs.Console(t); want != got {
			t.Errorf("Console(t) got %v, want %v", got, want)
		}
	})
}

func TestGNMIATE(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		wantErr := "bad gnmi"
		ate.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return nil, errors.New(wantErr)
		}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			ateAPIs.GNMI(t)
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("GNMI(t) got err %v, want %v", gotErr, wantErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		want := struct{ gpb.GNMIClient }{}
		ate.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
			return want, nil
		}
		if got := ateAPIs.GNMI(t); got != want {
			t.Errorf("GNMI(t) got %v, want %v", got, want)
		}
	})
}
