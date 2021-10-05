package clientwrapper

import (
	"bytes"
	"golang.org/x/net/context"
	"testing"
	"time"

	"github.com/openconfig/ondatra/fakes/fakestreamclient"
)

type action struct {
	wait   time.Duration
	stdout []byte
	stderr []byte
	err    bool
}

func TestClientWrapper(t *testing.T) {
	overBuffer := fakestreamclient.RandSeq(8000)
	tests := []struct {
		desc    string
		actions []action
	}{{
		desc: "simple writes",
		actions: []action{
			{stdout: []byte("1 multiple writes\n")},
			{stdout: []byte("2 multiple writes\n")},
			{stdout: []byte("3 multiple writes\n")},
			{stdout: []byte("4 multiple writes\n")},
			{stdout: []byte("5 multiple writes\n")},
			{stdout: []byte("6 multiple writes\n")},
			{stdout: []byte("7 multiple writes\n")},
			{stdout: []byte("8 multiple writes\n")},
		},
	}, {
		desc: "simple err and out",
		actions: []action{
			{stdout: []byte("1 multiple writes\n")},
			{stderr: []byte("2 multiple writes\n")},
			{stdout: []byte("3 multiple writes\n")},
			{stdout: []byte("4 multiple writes\n")},
			{stderr: []byte("5 multiple writes\n")},
			{stdout: []byte("6 multiple writes\n")},
			{stdout: []byte("7 multiple writes\n")},
			{stderr: []byte("8 multiple writes\n")},
		},
	}, {
		desc: "simple err and out with err",
		actions: []action{
			{stdout: []byte("1 multiple writes\n")},
			{stderr: []byte("2 multiple writes\n")},
			{stdout: []byte("3 multiple writes\n")},
			{stdout: []byte("4 multiple writes\n")},
			{stderr: []byte("5 multiple writes\n")},
			{stdout: []byte("6 multiple writes\n")},
			{stdout: []byte("7 multiple writes\n")},
			{err: true},
		},
	}, {
		desc: "oversize buffer",
		actions: []action{
			{stdout: []byte(overBuffer)},
			{stderr: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
			{stderr: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			b := bytes.NewBuffer([]byte{})
			sc := fakestreamclient.New()
			ctx := context.Background()
			l := Start(ctx, sc, b)
			bLen := 0
			for _, a := range tt.actions {
				time.Sleep(a.wait)
				switch {
				case a.err:
					sc.OutWriter.Close()
				case len(a.stdout) != 0:
					bLen += len(a.stdout) + 9 // prefix
					sc.OutWriter.Write(a.stdout)
				case len(a.stderr) != 0:
					bLen += len(a.stdout) + 8 // prefix
					sc.ErrWriter.Write(a.stderr)
				}
			}
			l.Stop()
			if bLen > len(b.Bytes()) {
				t.Fatalf("Unexpected buffer output length: got %d, want %d", len(b.Bytes()), bLen)
			}
		})
	}
}
