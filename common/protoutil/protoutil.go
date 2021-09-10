// Package protoutil provides utilities for handling protos.
package protoutil

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// ReadText reads a text protocol buffer from a given path.
// If the path is a plain filename, interprets it relative to the target directory.
func ReadText(p string, pb proto.Message) error {
	s, err := ioutil.ReadFile(p)
	if err != nil {
		return errors.Wrapf(err, "failed to read %s", p)
	}
	err = prototext.Unmarshal(s, pb)
	if err != nil {
		return errors.Wrapf(err, "failed to parse %s", p)
	}
	return nil
}
