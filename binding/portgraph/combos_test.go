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

package portgraph

import (
	"testing"
)

func TestCheckEnoughPorts(t *testing.T) {
	port := &ConcretePort{}
	tests := []struct {
		desc string
		np   []*needPorts
		want bool
	}{{
		desc: "ok",
		np:   []*needPorts{{1, []*ConcretePort{&ConcretePort{}, &ConcretePort{}}}, {2, []*ConcretePort{&ConcretePort{}, &ConcretePort{}}}},
		want: true,
	}, {
		desc: "not enough unique ports",
		np:   []*needPorts{{1, []*ConcretePort{port}}, {2, []*ConcretePort{port, &ConcretePort{}}}},
		want: false,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got := checkEnoughPorts(tc.np)
			if got != tc.want {
				t.Errorf("checkEnoughPorts got %t, want %t", got, tc.want)
			}
		})
	}
}
