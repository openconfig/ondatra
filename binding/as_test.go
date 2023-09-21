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

package binding

import (
	"strings"
	"testing"
)

var target interface {
	testFunc()
}

func TestDUTAs(t *testing.T) {
	tests := []struct {
		desc    string
		fn      func() error
		wantErr string
	}{{
		desc: "target nil",
		fn: func() error {
			return DUTAs(&testDUT{}, nil)
		},
		wantErr: "target must",
	}, {
		desc: "target non-pointer",
		fn: func() error {
			return DUTAs(&testDUT{}, testDUT{})
		},
		wantErr: "target must",
	}, {
		desc: "src nil",
		fn: func() error {
			return DUTAs(nil, &target)
		},
		wantErr: "src must",
	}, {
		desc: "src non-pointer",
		fn: func() error {
			return DUTAs(wrapDUT{}, &target)
		},
		wantErr: "src must",
	}, {
		desc: "no match",
		fn: func() error {
			return DUTAs(&wrapDUT{}, &target)
		},
		wantErr: "no match",
	}, {
		desc: "no match - wrapped",
		fn: func() error {
			return DUTAs(&wrapDUT{&AbstractDUT{}}, &target)
		},
		wantErr: "no match",
	}, {
		desc: "wrong embed type",
		fn: func() error {
			return DUTAs(&struct {
				*AbstractDUT
				DUT string
			}{
				AbstractDUT: &AbstractDUT{},
			}, &target)
		},
		wantErr: "embedded field",
	}, {
		desc: "match",
		fn: func() error {
			return DUTAs(&testDUT{}, &target)
		},
	}, {
		desc: "match - wrapped",
		fn: func() error {
			return DUTAs(&wrapDUT{&testDUT{}}, &target)
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			target = nil
			gotErr := test.fn()
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("DUTAs() got err %v, want err %v", gotErr, test.wantErr)
			}
			if gotTarget, wantTarget := target != nil, test.wantErr == ""; gotTarget != wantTarget {
				t.Errorf("DUTAs() got target? %v, want target? %v", gotTarget, wantTarget)
			}
		})
	}
}

type wrapDUT struct {
	DUT
}

type testDUT struct {
	DUT
}

func (*testDUT) testFunc() {}

func TestATEAs(t *testing.T) {
	tests := []struct {
		desc    string
		fn      func() error
		wantErr string
	}{{
		desc: "target nil",
		fn: func() error {
			return ATEAs(&testATE{}, nil)
		},
		wantErr: "target",
	}, {
		desc: "target non-pointer",
		fn: func() error {
			return ATEAs(&testATE{}, testATE{})
		},
		wantErr: "target must",
	}, {
		desc: "src nil",
		fn: func() error {
			return ATEAs(nil, &target)
		},
		wantErr: "src must",
	}, {
		desc: "src non-pointer",
		fn: func() error {
			return ATEAs(wrapATE{}, &target)
		},
		wantErr: "src must",
	}, {
		desc: "no match",
		fn: func() error {
			return ATEAs(&wrapATE{}, &target)
		},
		wantErr: "no match",
	}, {
		desc: "no match - wrapped",
		fn: func() error {
			return ATEAs(&wrapATE{&AbstractATE{}}, &target)
		},
		wantErr: "no match",
	}, {
		desc: "wrong embed type",
		fn: func() error {
			return ATEAs(&struct {
				*AbstractATE
				ATE string
			}{
				AbstractATE: &AbstractATE{},
			}, &target)
		},
		wantErr: "embedded field",
	}, {
		desc: "match",
		fn: func() error {
			return ATEAs(&testATE{}, &target)
		},
	}, {
		desc: "match - wrapped",
		fn: func() error {
			return ATEAs(&wrapATE{&testATE{}}, &target)
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			target = nil
			gotErr := test.fn()
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("ATEAs() got err %v, want err %v", gotErr, test.wantErr)
			}
			if gotTarget, wantTarget := target != nil, test.wantErr == ""; gotTarget != wantTarget {
				t.Errorf("ATEAs() got target? %v, want target? %v", gotTarget, wantTarget)
			}
		})
	}
}

type wrapATE struct {
	ATE
}

type testATE struct {
	ATE
}

func (*testATE) testFunc() {}
