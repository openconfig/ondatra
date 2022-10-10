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

package gnmifix

import (
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
	"testing"
)

func TestTelemetryTypesToYGNMITypes(t *testing.T) {
	tests := []struct {
		desc     string
		in       string
		wantExpr string
		wantRet  bool
	}{{
		desc:     "called on not ident",
		in:       "Foo().hello",
		wantExpr: "Foo().hello",
		wantRet:  false,
	}, {
		desc:     "called on not telemetry",
		in:       "oc.hello",
		wantExpr: "oc.hello",
		wantRet:  false,
	}, {
		desc:     "not qualified type",
		in:       "telemetry.Interface",
		wantExpr: "oc.Interface",
		wantRet:  true,
	}, {
		desc:     "telemetry.QualifiedInterface",
		in:       "telemetry.QualifiedInterface",
		wantExpr: "ygnmi.Value[*oc.Interface]",
		wantRet:  true,
	}, {
		desc:     "telemetry.QualifiedE_Interface_OperStatus",
		in:       "telemetry.QualifiedE_Interface_OperStatus",
		wantExpr: "ygnmi.Value[oc.E_Interface_OperStatus]",
		wantRet:  true,
	}, {
		desc:     "telemetry.QualifiedStringSlice",
		in:       "telemetry.QualifiedStringSlice",
		wantExpr: "ygnmi.Value[[]string]",
		wantRet:  true,
	}}
	for _, tt := range tests {
		n, err := parser.ParseExpr(tt.in)
		if err != nil {
			t.Fatalf("could parse input %q, err %v", tt.in, err)
		}
		gotRet := telemetryTypesToYGNMITypes(n)
		var buf strings.Builder
		printer.Fprint(&buf, token.NewFileSet(), n)
		if got := buf.String(); got != tt.wantExpr {
			t.Fatalf("telemetryTypesToYGNMITypes(%q) unexpected expr: got %q, want %q", tt.in, got, tt.wantExpr)
		}
		if gotRet != tt.wantRet {
			t.Fatalf("telemetryTypesToYGNMITypes(%q) unexpected return: got %t, want %t", tt.in, gotRet, tt.wantRet)
		}
	}
}
