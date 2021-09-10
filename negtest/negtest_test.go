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

func TestErrorMsg(t *testing.T) {
	tests := []struct {
		desc               string
		fn                 func(t testing.TB)
		wantMsgs           []string
		wantFatalSubstring string
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
		desc:               "Fatal due to no error",
		fn:                 func(t testing.TB) {},
		wantFatalSubstring: "did noit raise an error as was expected",
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if tt.wantFatalSubstring != "" {
				if got := ExpectFatal(t, func(t testing.TB) { ExpectError(t, tt.fn) }); strings.Contains(got, tt.wantFatalSubstring) {
					t.Fatalf("did not get expected fatal message, got: %v, want: %v", got, tt.wantFatalSubstring)
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
		t.Errorf("panic arg = %q, want %q", got, wantPanicArg)
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
