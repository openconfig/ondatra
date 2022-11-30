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

package netutil

import (
	"sort"
	"testing"

	"github.com/openconfig/ygnmi/ygnmi"
	"golang.org/x/exp/constraints"
)

// MeanRate calculates the average rate across the supplied interval
// from the values supplied in the vals slice. It sorts the values by
// timestamp and uses the earliest and latest to calculate the rate. Errors
// are reported using the supplied testing.TB.
func MeanRate[T constraints.Integer](t testing.TB, vals []*ygnmi.Value[T]) float64 {
	t.Helper()

	if len(vals) < 2 {
		t.Fatalf("Cannot calculate rate from %d length slice, got %v", len(vals), vals)
	}

	sort.Slice(vals, func(i, j int) bool {
		return vals[i].Timestamp.Before(vals[j].Timestamp)
	})
	first, ok := vals[0].Val()
	if !ok {
		t.Fatalf("First value is not present: %v", vals[0])
	}
	last, ok := vals[len(vals)-1].Val()
	if !ok {
		t.Fatalf("Last value is not present: %v", vals[len(vals)-1])
	}
	return float64(last-first) / vals[len(vals)-1].Timestamp.Sub(vals[0].Timestamp).Seconds()
}
