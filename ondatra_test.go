package ondatra

import (
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestReserveOnRun(t *testing.T) {
	origRunTests := runTestsFn
	defer func() {
		reserveFn = reserve
		releaseFn = release
		runTestsFn = origRunTests
	}()
	tests := []struct {
		desc                  string
		sig                   os.Signal
		reservErr, releaseErr error
		wantReleased          bool
		wantErr               string
	}{{
		desc:      "error on reserve",
		reservErr: errors.New("error reserving testbed"),
		wantErr:   "error reserving testbed",
	}, {
		desc:         "error on release",
		releaseErr:   errors.New("error releasing testbed"),
		wantReleased: true,
		wantErr:      "error releasing testbed",
	}, {
		desc:         "release on signal",
		sig:          os.Interrupt,
		wantReleased: true,
	}, {
		desc:         "release on test completion",
		wantReleased: true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			reserveFn = func(_ string, _, _ time.Duration) error {
				return test.reservErr
			}
			tbReleased := false
			releaseCh := make(chan bool, 1)
			releaseMtx := &sync.Mutex{}
			defer close(releaseCh)
			releaseFn = func() error {
				releaseMtx.Lock()
				defer releaseMtx.Unlock()
				if !tbReleased {
					tbReleased = true
					releaseCh <- true
				}
				return test.releaseErr
			}

			runTestsFn = func(*fixture, *testing.M, time.Duration) int { return 0 }
			if test.sig != nil {
				// Run indefinitely if we should close on a signal.
				runTestsFn = func(*fixture, *testing.M, time.Duration) int {
					<-releaseCh
					return 0
				}
			}

			if test.sig != nil {
				sigc <- os.Interrupt
			}
			_, gotErr := doRun(nil)
			if (gotErr != nil) != (test.wantErr != "") {
				t.Fatalf("doRun: got err %v, wanted err? %t", gotErr, test.wantErr != "")
			}
			if test.wantErr != "" && !strings.Contains(gotErr.Error(), test.wantErr) {
				t.Errorf("doRun: unexpected error: got %v, wanted error with substring %q", gotErr, test.wantErr)
			}

			releaseMtx.Lock()
			defer releaseMtx.Unlock()
			if tbReleased != test.wantReleased {
				t.Errorf("doRun: testbed released? %t, wanted testbed released? %t", tbReleased, test.wantReleased)
			}
		})
	}
}
