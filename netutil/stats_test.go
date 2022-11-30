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
	"strings"
	"testing"
	"time"

	"github.com/openconfig/testt"
	"github.com/openconfig/ygnmi/ygnmi"
)

func TestMeanRate(t *testing.T) {
	testsPass := []struct {
		desc   string
		inVals []*ygnmi.Value[int]
		want   float64
	}{{
		desc: "two values",
		inVals: []*ygnmi.Value[int]{
			(&ygnmi.Value[int]{Timestamp: time.Unix(1, 0)}).SetVal(1),
			(&ygnmi.Value[int]{Timestamp: time.Unix(2, 0)}).SetVal(2),
		},
		want: 1.0,
	}, {
		desc: "vals must be sorted",
		inVals: []*ygnmi.Value[int]{
			(&ygnmi.Value[int]{Timestamp: time.Unix(50, 0)}).SetVal(9000),
			(&ygnmi.Value[int]{Timestamp: time.Unix(40, 0)}).SetVal(0),
			(&ygnmi.Value[int]{Timestamp: time.Unix(20, 0)}).SetVal(0),
			(&ygnmi.Value[int]{Timestamp: time.Unix(30, 0)}).SetVal(0),
		},
		want: 300.0,
	}}
	for _, test := range testsPass {
		t.Run(test.desc, func(t *testing.T) {
			if got := MeanRate(t, test.inVals); got != test.want {
				t.Errorf("MeanRate(%v) got %.4f, want %.4f", test.inVals, got, test.want)
			}
		})
	}

	testsFail := []struct {
		desc         string
		inVals       []*ygnmi.Value[int]
		wantFatalMsg string
	}{{
		desc:         "zero values",
		inVals:       nil,
		wantFatalMsg: "Cannot calculate rate from 0 length slice",
	}, {
		desc: "one value",
		inVals: []*ygnmi.Value[int]{
			(&ygnmi.Value[int]{Timestamp: time.Unix(1, 0)}).SetVal(1),
		},
		wantFatalMsg: "Cannot calculate rate from 1 length slice",
	}, {
		desc: "first value not present",
		inVals: []*ygnmi.Value[int]{
			(&ygnmi.Value[int]{Timestamp: time.Unix(1, 0)}),
			(&ygnmi.Value[int]{Timestamp: time.Unix(2, 0)}).SetVal(2),
		},
		wantFatalMsg: "First value is not present",
	}, {
		desc: "last value not present",
		inVals: []*ygnmi.Value[int]{
			(&ygnmi.Value[int]{Timestamp: time.Unix(1, 0)}).SetVal(1),
			(&ygnmi.Value[int]{Timestamp: time.Unix(2, 0)}),
		},
		wantFatalMsg: "Last value is not present",
	}}
	for _, test := range testsFail {
		t.Run(test.desc, func(t *testing.T) {
			got := testt.ExpectFatal(t, func(t testing.TB) {
				MeanRate(t, test.inVals)
			})
			if !strings.Contains(got, test.wantFatalMsg) {
				t.Errorf("MeanRate(%v) got failure %q, want %q", test.inVals, got, test.wantFatalMsg)
			}
		})
	}
}
