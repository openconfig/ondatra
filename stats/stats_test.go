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

package stats

import (
	"strings"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/negtest"
	"github.com/openconfig/ondatra/telemetry"
)

func TestMeanRateUint64(t *testing.T) {
	testsPass := []struct {
		desc   string
		inVals []*telemetry.QualifiedUint64
		want   float64
	}{{
		desc: "two values",
		inVals: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(1, 0),
				}}).SetVal(1),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(2, 0),
				}}).SetVal(2),
		},
		want: 1.0,
	}, {
		desc: "vals must be sorted",
		inVals: []*telemetry.QualifiedUint64{
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(50, 0),
				}}).SetVal(9000),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(40, 0),
				}}).SetVal(0),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(20, 0),
				},
			}).SetVal(0),
			(&telemetry.QualifiedUint64{
				Metadata: &genutil.Metadata{
					Timestamp: time.Unix(30, 0),
				}}).SetVal(0),
		},
		want: 300.0,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			if got := MeanRateUint64(t, tt.inVals); got != tt.want {
				t.Errorf("did not get expected rate, got: %.4f, want: %.4f", got, tt.want)
			}
		})
	}

	testsFail := []struct {
		desc         string
		inVals       []*telemetry.QualifiedUint64
		wantFatalMsg string
	}{{
		desc:         "zero values",
		inVals:       nil,
		wantFatalMsg: "cannot calculate rate from 0 length slice",
	}, {
		desc:         "one value",
		inVals:       []*telemetry.QualifiedUint64{{}},
		wantFatalMsg: "cannot calculate rate from 1 length slice",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				MeanRateUint64(t, tt.inVals)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("MeanRateUint64(%v) failed with message %q, want: %q", tt.inVals, got, tt.wantFatalMsg)
			}
		})
	}
}
