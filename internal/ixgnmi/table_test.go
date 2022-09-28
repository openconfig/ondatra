// Copyright 2021 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixgnmi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
)

type example struct {
	Str  string `ixia:"str"`
	I    int    `ixia:"i"`
	I8   int8   `ixia:"i8"`
	I16  int16  `ixia:"i16"`
	I32  int32  `ixia:"i32"`
	I64  int64  `ixia:"i64"`
	UI   uint   `ixia:"ui"`
	UI8  uint8  `ixia:"ui8"`
	UI16 uint16 `ixia:"ui16"`
	UI32 uint32 `ixia:"ui32"`
	UI64 uint64 `ixia:"ui64"`
}

var fullTable = &table{
	Columns: []string{"str", "i", "i8", "i16", "i32", "i64", "ui", "ui8", "ui16", "ui32", "ui64"},
	Values: [][]string{{
		"hello", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	}},
}

func TestUnmarshalTable(t *testing.T) {
	tests := []struct {
		name    string
		inSlice any
		inTable *table
		want    any
		wantErr string
	}{{
		name:    "not a pointer",
		inSlice: "",
		inTable: &table{},
		wantErr: "must pass in a pointer to slice",
	}, {
		name:    "not a slice",
		inSlice: &struct{}{},
		inTable: &table{},
		wantErr: "must pass in a pointer to slice",
	}, {
		name:    "not a pointer to slice of structs",
		inSlice: &[]string{},
		inTable: &table{
			Values: [][]string{{
				"Test",
			}},
		},
		wantErr: "expected struct type",
	}, {
		name:    "not a pointer to slice of structs pointer",
		inSlice: &[]*string{},
		inTable: &table{
			Values: [][]string{{
				"Test",
			}},
		},
		wantErr: "expected struct type",
	}, {
		name:    "string parse error",
		inSlice: &[]example{},
		inTable: &table{
			Columns: []string{"i"},
			Values: [][]string{{
				"Test",
			}},
		},
		wantErr: "parse field \"I\" failed",
	}, {
		name:    "full table",
		inSlice: &[]example{},
		inTable: fullTable,
		want: &[]example{{
			Str:  "hello",
			I:    1,
			I8:   2,
			I16:  3,
			I32:  4,
			I64:  5,
			UI:   6,
			UI8:  7,
			UI16: 8,
			UI32: 9,
			UI64: 10,
		}},
	}, {
		name:    "full table struct pointer",
		inSlice: &[]*example{},
		inTable: fullTable,
		want: &[]*example{{
			Str:  "hello",
			I:    1,
			I8:   2,
			I16:  3,
			I32:  4,
			I64:  5,
			UI:   6,
			UI8:  7,
			UI16: 8,
			UI32: 9,
			UI64: 10,
		}},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := unmarshalTable(tt.inTable, tt.inSlice)
			if d := errdiff.Substring(err, tt.wantErr); d != "" {
				t.Fatalf("unexpected error: %s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(tt.inSlice, tt.want); d != "" {
				t.Fatalf("unexpected result: %s", d)
			}
		})
	}
}
