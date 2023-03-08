// Copyright 2023 Google LLC
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

package creds

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	tpb "github.com/openconfig/kne/proto/topo"
	"gopkg.in/yaml.v2"
)

func TestUnmarshal(t *testing.T) {
	credStr := `
node:
  n1:
    username: user1
    password: pass1
  n2:
    username: user2
    password: pass2
vendor:
  OPENCONFIG:
    username: user3
    password: pass3
username: user4
password: pass4`
	got := new(Credentials)
	if err := yaml.Unmarshal([]byte(credStr), got); err != nil {
		t.Fatalf("Unmarshal got unexpected error: %v", err)
	}

	want := &Credentials{
		Node: map[string]*UserPass{
			"n1": {Username: "user1", Password: "pass1"},
			"n2": {Username: "user2", Password: "pass2"},
		},
		Vendor: map[tpb.Vendor]*UserPass{
			tpb.Vendor_OPENCONFIG: {Username: "user3", Password: "pass3"},
		},
		Default: &UserPass{Username: "user4", Password: "pass4"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Unmarshal got unexpected diff: %s", diff)
	}
}

func TestUnmarshalError(t *testing.T) {
	credStr := `
vendor:
  NONSENSE:
    username: user3
    password: pass3`
	gotErr := yaml.Unmarshal([]byte(credStr), new(Credentials))
	if want := "NONSENSE"; gotErr == nil || !strings.Contains(gotErr.Error(), want) {
		t.Errorf("Unmarshal got error %v, want error containing %q", gotErr, want)
	}
}

func TestLookup(t *testing.T) {
	credsNoDefault := &Credentials{
		Node: map[string]*UserPass{
			"n1": {Username: "nodeUser", Password: "nodePass"},
		},
		Vendor: map[tpb.Vendor]*UserPass{
			tpb.Vendor_OPENCONFIG: {Username: "vendorUser", Password: "vendorPass"},
		},
	}
	credsWithDefault := &Credentials{
		Node:    credsNoDefault.Node,
		Vendor:  credsNoDefault.Vendor,
		Default: &UserPass{Username: "defaultUser", Password: "defaultPass"},
	}

	tests := []struct {
		desc   string
		creds  *Credentials
		name   string
		vendor tpb.Vendor
		want   *UserPass
	}{{
		desc:   "node match",
		creds:  credsWithDefault,
		name:   "n1",
		vendor: tpb.Vendor_OPENCONFIG,
		want:   &UserPass{Username: "nodeUser", Password: "nodePass"},
	}, {
		desc:   "vendor match",
		creds:  credsWithDefault,
		name:   "n2",
		vendor: tpb.Vendor_OPENCONFIG,
		want:   &UserPass{Username: "vendorUser", Password: "vendorPass"},
	}, {
		desc:   "default match",
		creds:  credsWithDefault,
		name:   "n2",
		vendor: tpb.Vendor_HOST,
		want:   &UserPass{Username: "defaultUser", Password: "defaultPass"},
	}, {
		desc:   "no match",
		creds:  credsNoDefault,
		name:   "n2",
		vendor: tpb.Vendor_HOST,
		want:   nil,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.creds.Lookup(tt.name, tt.vendor)
			if !cmp.Equal(tt.want, got) {
				t.Errorf("Lookup(%s, %v) got %v, want %v", tt.name, tt.vendor, got, tt.want)
			}
		})
	}
}

func TestParseFlags(t *testing.T) {
	flags := &Flags{
		nodeCreds:    []string{"n1/user1/pass1", "n2/user2/pass2"},
		vendorCreds:  []string{"OPENCONFIG/user3/pass3"},
		defaultCreds: "user4/pass4",
	}
	got, err := flags.Parse()
	if err != nil {
		t.Fatalf("Parse got unexpected error: %v", err)
	}

	want := &Credentials{
		Node: map[string]*UserPass{
			"n1": {Username: "user1", Password: "pass1"},
			"n2": {Username: "user2", Password: "pass2"},
		},
		Vendor: map[tpb.Vendor]*UserPass{
			tpb.Vendor_OPENCONFIG: {Username: "user3", Password: "pass3"},
		},
		Default: &UserPass{Username: "user4", Password: "pass4"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Parse got unexpected diff: %s", diff)
	}
}

func TestParseFlagsErrors(t *testing.T) {
	tests := []struct {
		desc  string
		flags *Flags
		want  string
	}{{
		desc: "invalid node cred",
		flags: &Flags{
			nodeCreds: []string{"n1/user"},
		},
		want: "invalid node",
	}, {
		desc: "invalid vendor cred",
		flags: &Flags{
			vendorCreds: []string{"OPENCONFIG"},
		},
		want: "invalid vendor",
	}, {
		desc: "invalid vendor enum",
		flags: &Flags{
			vendorCreds: []string{"NONSENSE/user/pass"},
		},
		want: "invalid vendor",
	}, {
		desc: "invalid default",
		flags: &Flags{
			defaultCreds: "user",
		},
		want: "invalid default",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			_, got := test.flags.Parse()
			if got == nil || !strings.Contains(got.Error(), test.want) {
				t.Fatalf("Parse got error %v, want error containing %q", got, test.want)
			}
		})
	}
}
