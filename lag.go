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

package ondatra

import (
	opb "github.com/openconfig/ondatra/proto"
)

// LAG is a link aggregation group.
type LAG struct {
	pb *opb.Lag
}

// WithPorts sets the LAG ports to the specified ports.
func (l *LAG) WithPorts(ports ...*Port) *LAG {
	l.pb.Ports = nil
	for _, p := range ports {
		l.pb.Ports = append(l.pb.Ports, p.Name())
	}
	return l
}

// LACP returns the LACP configuration of the LAG.
func (l *LAG) LACP() *LACP {
	return &LACP{pb: l.pb.Lacp}
}

// LACP is a LACP configuration on a LAG.
type LACP struct {
	pb *opb.Lag_Lacp
}

// WithEnabled sets whether LACP is enabled on the LAG.
func (l *LACP) WithEnabled(enabled bool) *LACP {
	l.pb.Enabled = enabled
	return l
}
