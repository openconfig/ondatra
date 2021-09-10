package usererr

import (
	"testing"

	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

func TestIn(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		wantIn bool
	}{{
		name:   "new",
		err:    New("foo"),
		wantIn: true,
	}, {
		name:   "wrap",
		err:    Wrap(errors.New("foo")),
		wantIn: true,
	}, {
		name:   "wrapped with errors.Wrap",
		err:    errors.Wrap(New("foo"), "bar"),
		wantIn: true,
	}, {
		name:   "wrapped with xerrors %w",
		err:    xerrors.Errorf("foo: %w", New("bar")),
		wantIn: true,
	}, {
		name:   "wrapped multiple times",
		err:    errors.Wrap(xerrors.Errorf("foo: %w", New("bar")), "baz"),
		wantIn: true,
	}, {
		name:   "not user err",
		err:    errors.New("foo"),
		wantIn: false,
	}, {
		name:   "not user error - wrapped",
		err:    errors.Wrap(xerrors.Errorf("foo: %w", errors.New("bar")), "baz"),
		wantIn: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.err); got != tt.wantIn {
				t.Errorf("In(%v) got %v, want %v", tt.err, got, tt.wantIn)
			}
		})
	}
}
