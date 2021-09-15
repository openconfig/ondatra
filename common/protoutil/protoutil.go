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
