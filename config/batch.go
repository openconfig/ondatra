// Copyright 2021 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
)

// SetRequestBatch is an object batch multiple config operations in a single gNMI request.
// Do not construct this directly: use dut.Config().NewBatch().
type SetRequestBatch struct {
	*privateSetRequestBatch
}

// NewSetBatch creates a set batch object. Do not call this directly: use dut.Config().NewBatch().
func NewSetBatch(fakeRoot genutil.FakeRootPathStruct) *SetRequestBatch {
	return &SetRequestBatch{&privateSetRequestBatch{genutil.NewSetRequestBatch(fakeRoot)}}
}

type privateSetRequestBatch struct {
	*genutil.SetRequestBatch
}
