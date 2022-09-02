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
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		desc string
		c    Constraint
		v    string
		want bool
	}{{
		desc: "Equal; match",
		c:    Equal("match"),
		v:    "match",
		want: true,
	}, {
		desc: "Equal; no match",
		c:    Equal("match"),
		v:    "not matching",
		want: false,
	}, {
		desc: "NotEqual; match",
		c:    NotEqual("don't match this"),
		v:    "not matching this",
		want: true,
	}, {
		desc: "NotEqual; no match",
		c:    NotEqual("match"),
		v:    "match",
		want: false,
	}, {
		desc: "Regex; match",
		c:    Regex(regexp.MustCompile(".*")),
		v:    "anything",
		want: true,
	}, {
		desc: "Regex; no match",
		c:    Regex(regexp.MustCompile("[0-9]+")),
		v:    "these are not numbers",
		want: false,
	}, {
		desc: "NotRegex; match",
		c:    NotRegex(regexp.MustCompile("[a-z]+")),
		v:    "1234",
		want: true,
	}, {
		desc: "NotRegex; no match",
		c:    NotRegex(regexp.MustCompile("[a-z]+")),
		v:    "these are not numbers",
		want: false,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			if m := tc.c.match(tc.v); m != tc.want {
				t.Errorf("match() got %t, want %t", m, tc.want)
			}
		})
	}
}
