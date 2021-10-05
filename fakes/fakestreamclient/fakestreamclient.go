// Package fakestreamclient provides a fake StreamClient implementation for
// testing.
package fakestreamclient

import (
	"golang.org/x/net/context"
	"fmt"
	"io"
	"math/rand"
)

// FakeStreamClient provides a fake implementation of a StreamClient for use in
// testing.
type FakeStreamClient struct {
	stdin     *io.PipeWriter
	stdout    *io.PipeReader
	stderr    *io.PipeReader
	InReader  *io.PipeReader
	OutWriter *io.PipeWriter
	ErrWriter *io.PipeWriter
	cErr      error
}

// New returns a new FakeStreamClient. Users should just access the test sides
// of the pipes for testing via exported fields.
func New() *FakeStreamClient {
	inReader, inWriter := io.Pipe()
	outReader, outWriter := io.Pipe()
	errReader, errWriter := io.Pipe()

	return &FakeStreamClient{
		InReader:  inReader,
		OutWriter: outWriter,
		ErrWriter: errWriter,
		stdin:     inWriter,
		stdout:    outReader,
		stderr:    errReader,
	}
}

// SendCommand will echo whatever command is sent. If the cmd is "error" an
// error till be returned.
func (f *FakeStreamClient) SendCommand(_ context.Context, cmd string) (string, error) {
	if cmd == "error" {
		return "", fmt.Errorf("error")
	}
	return cmd, nil
}

// Close will return a configured err.
func (f *FakeStreamClient) Close() error {
	f.InReader.Close()
	f.OutWriter.Close()
	f.ErrWriter.Close()
	return f.cErr
}

// Stdin will return the configured buffer.
func (f *FakeStreamClient) Stdin() io.WriteCloser {
	return f.stdin
}

// Stdout will return the configured buffer.
func (f *FakeStreamClient) Stdout() io.ReadCloser {
	return f.stdout
}

// Stderr will return the configured buffer.
func (f *FakeStreamClient) Stderr() io.ReadCloser {
	return f.stderr
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandSeq will generate a random character string of len n.
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
