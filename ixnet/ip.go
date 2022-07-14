// Copyright 2020 Google LLC
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

package ixnet

import (
	opb "github.com/openconfig/ondatra/proto"
)

// NewIP returns a new IP configuration.
// Tests must not call this method directly.
func NewIP(pb *opb.IpConfig) *IP {
	return &IP{pb}
}

// IP is an representation of IP config on the ATE.
type IP struct {
	pb *opb.IpConfig
}

// WithAddress sets the IP address in CIDR notation.
func (i *IP) WithAddress(address string) *IP {
	i.pb.AddressCidr = address
	return i
}

// WithDefaultGateway sets the default gateway.
func (i *IP) WithDefaultGateway(gateway string) *IP {
	i.pb.DefaultGateway = gateway
	return i
}
