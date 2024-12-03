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

// Package creds contains logic for DNOS credentials.
package creds

import (
	"flag"
	"fmt"
	"strings"
)

// UserPass is a username/password combination.
type UserPass struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (up *UserPass) String() string {
	return fmt.Sprintf("%s/%s", up.Username, up.Password)
}

type Credentials struct {
	Node map[string]*UserPass
}

// Lookup returns the username/password to use for the given node name and vendor.
// Returns nil if no such combination exists. This method is nil-tolerant.
func (c *Credentials) Lookup(node string) *UserPass {
	if c != nil {
		if up, ok := c.Node[node]; ok {
			return up
		}
	}
	return nil
}

func (c *Credentials) String() string {
	return fmt.Sprintf("%+v", *c)
}

// DefineFlags defines flags for allowing user-specified credentials.
func DefineFlags() *Flags {
	flags := new(Flags)
	multiStringVar(&flags.nodeCreds, "node_creds",
		"Repeated node-specific credentials in the form 'nodeName/username/password', e.g., 'n1/admin/hunter2'")
	return flags
}

func multiStringVar(v *[]string, name, usage string) {
	flag.Var((*multiStringVal)(v), name, usage)
}

type multiStringVal []string

func (v *multiStringVal) Get() any {
	return []string(*v)
}

func (v *multiStringVal) Set(s string) error {
	*v = append(*v, s)
	return nil
}

func (*multiStringVal) TypeDescription() string {
	return "repeated string"
}

func (v *multiStringVal) String() string {
	if len(*v) == 0 {
		return ""
	}
	return fmt.Sprint(*v)
}

// Flags are a collection of credentials flags.
type Flags struct {
	nodeCreds []string
}

// Parse parses the flags into Credentials.
func (f *Flags) Parse() (*Credentials, error) {
	c := new(Credentials)
	if len(f.nodeCreds) > 0 {
		c.Node = make(map[string]*UserPass, len(f.nodeCreds))
		for _, nc := range f.nodeCreds {
			parts := strings.SplitN(nc, "/", 3)
			if len(parts) != 3 {
				return nil, fmt.Errorf("invalid node credentials: %q", nc)
			}
			c.Node[parts[0]] = &UserPass{Username: parts[1], Password: parts[2]}
		}
	}
	return c, nil
}
