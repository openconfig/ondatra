// Package negtest provides utilities for writing negative tests.
package negtest

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

// ExpectFatal fails the test if the specified function does _not_ fail fatally,
// i.e. does not call any of t.{FailNow, Fatal, Fatalf}.
// If it does fail fatally, returns the fatal error message it logged.
// It is recommended the error message be checked to distinguish the
// expected failure from unrelated failures that may have occurred.
func ExpectFatal(t testing.TB, fn func(t testing.TB)) (msg string) {
	t.Helper()
	// Defer and recover to capture the expected fatal message.
	defer func() {
		switch r := recover().(type) {
		case failure:
			// panic from fatal fakeT failure, return the message
			msg = string(r)
		case nil:
			// no panic at all, do nothing
		default:
			// another panic was detected, re-raise
			panic(r)
		}
	}()
	fn(&fakeT{realT: t})
	t.Fatalf("%s did not fail fatally as expected", funcName(fn))
	return ""
}

func funcName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// ExpectError determines whether t.Errorf or t.Error was called at least
// once during a test, and returns the set of strings that were specified
// as arguments to the error calls.
func ExpectError(t testing.TB, fn func(testing.TB)) []string {
  t.Helper()
	ft := &fakeT{realT: t}
	fn(ft)
	if ft.errs == nil {
		t.Fatalf("%s did not raise an error as was expected", funcName(fn))
	}
	return ft.errs
}

// fakeT is a testing.TB implementation that can be used as an input to unit tests
// such that it is possible to check that the correct errors are raised.
type fakeT struct {
	// Any methods not explicitly implemented here will panic when called.
	testing.TB
	realT testing.TB

	// err is used to store the strings that are specified as arguments to
	// Error and Errorf when it is called.
	errs []string
}

// failure is a unique type to distinguish test failures from other panics.
type failure string

// FailNow implements the testing.TB FailNow method so that the failure can be
// retrieved by making the call within the lambda argument to ExpectFatal.
func (ft *fakeT) FailNow() {
	ft.fatal("")
}

// Fatal implements the testing.TB Fatalf method so that the failure can be
// retrieved by making the call within the lambda argument to ExpectFatal.
func (ft *fakeT) Fatal(args ...interface{}) {
	ft.fatal(fmt.Sprintln(args...))
}

// Fatalf implements the testing.TB Fatalf method so that the failure can be
// retrieved by making the call within the lambda argument to ExpectFatal.
func (ft *fakeT) Fatalf(format string, args ...interface{}) {
	ft.fatal(fmt.Sprintf(format, args...))
}

func (ft *fakeT) fatal(msg string) {
	panic(failure(msg))
}

// Log implements the testing.TB Log method by delegating to the real *testing.T.
func (ft *fakeT) Log(args ...interface{}) {
	ft.realT.Log(args...)
}

// Log implements the testing.TB Logf method by delegating to the real *testing.T.
func (ft *fakeT) Logf(format string, args ...interface{}) {
	ft.realT.Logf(format, args...)
}

// Errorf implements the testing.TB Errorf method, but rather than reporting the
// error catches it in the errs field of the fakeT.
func (ft *fakeT) Errorf(format string, args ...interface{}) {
	ft.errs = append(ft.errs, fmt.Sprintf(format, args...))
}

// Error implements the testing.TB Error method, but rather than reporting the
// error catches it in the errs field of the fakeT.
func (ft *fakeT) Error(args ...interface{}) {
	ft.errs = append(ft.errs, fmt.Sprintln(args...))
}

// Helper implements the testing.TB Helper method as a noop.
func (*fakeT) Helper() {}
