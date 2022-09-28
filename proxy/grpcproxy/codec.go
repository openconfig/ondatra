// Copyright 2022 Google LLC
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

package grpcproxy

import (
	"fmt"

	"google.golang.org/grpc/encoding"
)

// rawProto defines the data structure interpretation of each message proxied by this server. Since
// this is a method-agnostic proxy, treat data as raw bytes.
type rawProto struct {
	data []byte
}

// Raw codec implements encoding.Codec. It Marshals and Unmarshals data into proto as raw bytes.
type rawCodec struct{}

func (rawCodec) Marshal(msg any) ([]byte, error) {
	out, ok := msg.(*rawProto)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want bytes", msg)
	}
	return out.data, nil
}

func (rawCodec) Unmarshal(data []byte, msg any) error {
	set, ok := msg.(*rawProto)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want bytes", msg)
	}
	set.data = data
	return nil
}

const codecName = "rawcodec"

func (rawCodec) Name() string {
	return codecName
}
func (rawCodec) String() string {
	return codecName
}

// codec returns a raw bytes based encode/decode implementation of grpc.Codec.
func codec() encoding.Codec {
	return &rawCodec{}
}
