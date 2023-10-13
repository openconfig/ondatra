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

// Package gnoi provides convenience functions for running gNOI operations in ondatra.
package gnoi

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/gnoigo"
	"github.com/openconfig/ondatra"
)

var (
	// Stub for testing only.
	stubExecuteFn func(context.Context, gnoigo.Clients, any) (any, error)
)

// Execute performs an operation and returns one or more response protos.
// For example, a PingOperation returns a slice of PingResponse messages.
func Execute[T any](t testing.TB, dut *ondatra.DUTDevice, op gnoigo.Operation[T]) T {
	resp, err := execute(context.Background(), dut.RawAPIs().GNOI(t), op)
	if err != nil {
		t.Fatal(err)
	}
	return resp
}

func execute[T any](ctx context.Context, clients gnoigo.Clients, op gnoigo.Operation[T]) (T, error) {
	if stubExecuteFn != nil {
		r, err := stubExecuteFn(ctx, clients, op)
		var result T
		if r != nil {
			result = r.(T)
		}
		// If r is nil, zero value of T is returned.
		return result, err
	}
	return gnoigo.Execute(ctx, clients, op)
}
