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

package ixnet

import (
	opb "github.com/openconfig/ondatra/proto"
)

// NewDHCPV6Client returns a new DHCP v6 client configuration.
// Tests must not call this method directly.
func NewDHCPV6Client(pb *opb.DhcpV6Client) *DHCPV6Client {
	return &DHCPV6Client{pb}
}

// NewDHCPV6Server returns a new DHCP v6 server configuration.
// Tests must not call this method directly.
func NewDHCPV6Server(pb *opb.DhcpV6Server) *DHCPV6Server {
	return &DHCPV6Server{pb}
}

// DHCPV6Client is a DHCP v6 Client config on an ATE.
type DHCPV6Client struct {
	pb *opb.DhcpV6Client
}

// DHCPV6Server is a DHCP v6 Server config on an ATE.
type DHCPV6Server struct {
	pb *opb.DhcpV6Server
}
