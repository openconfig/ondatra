package closer

import (
	"testing"

	"github.com/pkg/errors"
)

var (
	errInput = errors.New("input error")
	errClose = errors.New("close error")

	closeErr, closeCalled bool
)

func fakeClose() error {
	closeCalled = true
	if closeErr {
		return errClose
	}
	return nil
}

func TestClose(t *testing.T) {
	tests := []struct {
		name      string
		inputErr  bool
		closeErr  bool
		wantCause error
	}{{
		name: "no errors",
	}, {
		name:      "input error only",
		inputErr:  true,
		wantCause: errInput,
	}, {
		name:      "close error only",
		closeErr:  true,
		wantCause: errClose,
	}, {
		name:      "both errors",
		inputErr:  true,
		closeErr:  true,
		wantCause: errInput,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input error
			if tt.inputErr {
				input = errInput
			}
			closeErr = tt.closeErr
			closeCalled = false
			Close(&input, fakeClose, "fake msg")
			if !closeCalled {
				t.Errorf("Close(): was not called")
			}
			if (input == nil) != (tt.wantCause == nil) {
				t.Errorf("Close(): got error %v, want cause %v", input, tt.wantCause)
			} else if input != nil && errors.Cause(input) != tt.wantCause {
				t.Errorf("Close(): got error %v, want cause %v", input, tt.wantCause)
			}
		})
	}
}

func TestCloseOnErr(t *testing.T) {
	tests := []struct {
		name      string
		inputErr  bool
		closeErr  bool
		wantCause error
	}{{
		name: "no errors",
	}, {
		name:      "input error only",
		inputErr:  true,
		wantCause: errInput,
	}, {
		name:     "close error only",
		closeErr: true,
	}, {
		name:      "both errors",
		inputErr:  true,
		closeErr:  true,
		wantCause: errInput,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input error
			if tt.inputErr {
				input = errInput
			}
			closeErr = tt.closeErr
			closeCalled = false
			CloseOnErr(&input, fakeClose, "fake msg")
			if wantCalled := tt.wantCause != nil; wantCalled != closeCalled {
				t.Errorf("CloseOnErr(): got called %v, want called %v", closeCalled, wantCalled)
			}
			if (input == nil) != (tt.wantCause == nil) {
				t.Errorf("CloseOnErr(): got error %v, want cause %v", input, tt.wantCause)
			} else if input != nil && errors.Cause(input) != tt.wantCause {
				t.Errorf("CloseOnErr(): got error %v, want cause %v", input, tt.wantCause)
			}
		})
	}
}
