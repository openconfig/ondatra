// Copyright 2021 Google LLC
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

package negtest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFatalMsg(t *testing.T) {
	tests := []struct {
		desc    string
		fn      func(t testing.TB)
		wantMsg string
	}{{
		desc: "FailNow",
		fn: func(t testing.TB) {
			t.FailNow()
		},
		wantMsg: "",
	}, {
		desc: "Fatal",
		fn: func(t testing.TB) {
			t.Fatal("fatal error")
		},
		wantMsg: "fatal error\n",
	}, {
		desc: "Fatalf",
		fn: func(t testing.TB) {
			t.Fatalf("fatalf error")
		},
		wantMsg: "fatalf error",
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := ExpectFatal(t, tt.fn); got != tt.wantMsg {
				t.Errorf("ExpectFatal got msg = %q, want %q", got, tt.wantMsg)
			}
		})
	}
}

func TestCaptureFatal(t *testing.T) {
	msgEmpty := ""
	msgFatalErr := "fatal error\n"
	msgFatalfErr := "fatalf error"

	tests := []struct {
		desc    string
		fn      func(t testing.TB)
		wantMsg *string
	}{{
		desc:    "NoError",
		fn:      func(t testing.TB) {},
		wantMsg: nil,
	}, {
		desc: "FailNow",
		fn: func(t testing.TB) {
			t.FailNow()
		},
		wantMsg: &msgEmpty,
	}, {
		desc: "Fatal",
		fn: func(t testing.TB) {
			t.Fatal("fatal error")
		},
		wantMsg: &msgFatalErr,
	}, {
		desc: "Fatalf",
		fn: func(t testing.TB) {
			t.Fatalf("fatalf error")
		},
		wantMsg: &msgFatalfErr,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := CaptureFatal(t, tt.fn)
			didGet, didWant := got != nil, tt.wantMsg != nil
			if didGet != didWant {
				t.Errorf("CaptureFatal got? %v, want? %v", didGet, didWant)
			}
			if didGet && didWant && *got != *tt.wantMsg {
				t.Errorf("CaptureFatal got msg = %q, want %q", *got, *tt.wantMsg)
			}
		})
	}
}

func TestErrorMsg(t *testing.T) {
	tests := []struct {
		desc          string
		fn            func(t testing.TB)
		wantMsgs      []string
		wantSubstring string
	}{{
		desc: "Errorf called",
		fn: func(t testing.TB) {
			t.Errorf("errorf")
		},
		wantMsgs: []string{"errorf"},
	}, {
		desc: "Errorf and Error called",
		fn: func(t testing.TB) {
			t.Error("hello")
			t.Error("Planet Earth")
		},
		wantMsgs: []string{"hello\n", "Planet Earth\n"},
	}, {
		desc:          "Fatal due to no error",
		fn:            func(t testing.TB) {},
		wantSubstring: "did not raise an error as was expected",
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if tt.wantSubstring != "" {
				if got := ExpectFatal(t, func(t testing.TB) { ExpectError(t, tt.fn) }); !strings.Contains(got, tt.wantSubstring) {
					t.Fatalf("ExpectError got unexpected message %q, want substring %q", got, tt.wantSubstring)
				}
				return
			}
			if got := ExpectError(t, tt.fn); !cmp.Equal(got, tt.wantMsgs) {
				t.Errorf("ExpectError got msg = %q, want %q", got, tt.wantMsgs)
			}
		})
	}
}

func TestNoFatal(t *testing.T) {
	tt := &testT{}
	ExpectFatal(tt, func(t testing.TB) {})
	if want := "did not fail fatally"; !strings.Contains(tt.got, want) {
		t.Errorf("Expect Fatal got msg = %q, want %q", tt.got, want)
	}
}

type testT struct {
	testing.TB
	got string
}

func (*testT) Helper() {}

func (tt *testT) Fatalf(format string, args ...interface{}) {
	tt.got = fmt.Sprintf(format, args...)
}

func TestPanic(t *testing.T) {
	wantPanicArg := "my panic"
	var got interface{}
	func() {
		defer func() {
			got = recover()
		}()
		ExpectFatal(t, func(t testing.TB) {
			panic(wantPanicArg)
		})
	}()
	if got != wantPanicArg {
		t.Errorf("Panic arg = %q, want %q", got, wantPanicArg)
	}
}

func TestBenignMethods(t *testing.T) {
	ExpectFatal(t, func(t testing.TB) {
		t.Helper()
		t.Log("hello")
		t.Logf("hello %v", "there")
		// Must fail to so that the test passes
		t.FailNow()
	})
}

func TestParallelFatal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ParallelFatal(t,
			func(testing.TB) {},
			func(testing.TB) {})
	})

	t.Run("failure", func(t *testing.T) {
		failMsg1, failMsg2 := "fail1", "fail2"
		got := ExpectFatal(t, func(t testing.TB) {
			ParallelFatal(t,
				func(t testing.TB) { t.Fatal(failMsg1) },
				func(testing.TB) {},
				func(t testing.TB) { t.Fatal(failMsg2) })
		})
		if !strings.Contains(got, failMsg1) || !strings.Contains(got, failMsg2) {
			t.Errorf("ParallelFatal got unexpected message %q, want substrings %q and %q", got, failMsg1, failMsg2)
		}
	})
}
