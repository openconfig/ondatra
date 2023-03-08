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

// Package creds contains logic for KNE credentials.
package creds

import (
	"fmt"
	"strings"

	"flag"

	tpb "github.com/openconfig/kne/proto/topo"
)

// UserPass is a username/password combination.
type UserPass struct {
	Username string
	Password string
}

func (up *UserPass) String() string {
	return fmt.Sprintf("%s/%s", up.Username, up.Password)
}

// Credentials contains credential info for nodes in the KNE topology.
type Credentials struct {
	Node    map[string]*UserPass
	Vendor  map[tpb.Vendor]*UserPass
	Default *UserPass
}

// Lookup returns the username/password to use for the given node name and vendor.
// Returns nil if no such combination exists. This method is nil-tolerant.
func (c *Credentials) Lookup(node string, vendor tpb.Vendor) *UserPass {
	if c != nil {
		if up, ok := c.Node[node]; ok {
			return up
		}
		if up, ok := c.Vendor[vendor]; ok {
			return up
		}
		if c.Default != nil {
			return c.Default
		}
	}
	return nil
}

func (c *Credentials) String() string {
	return fmt.Sprintf("%+v", *c)
}

// UnmarshalYAML allows the Credentials type to be correctly unmarshaled from yaml.
func (c *Credentials) UnmarshalYAML(unmarshal func(any) error) error {
	var u struct {
		Node     map[string]*UserPass
		Vendor   map[string]*UserPass
		Username string
		Password string
	}
	if err := unmarshal(&u); err != nil {
		return err
	}
	c.Node = u.Node
	if len(u.Vendor) > 0 {
		c.Vendor = make(map[tpb.Vendor]*UserPass, len(u.Vendor))
		for k, v := range u.Vendor {
			n, ok := tpb.Vendor_value[k]
			if !ok {
				return fmt.Errorf("kne vendor %v not recognized", k)
			}
			c.Vendor[tpb.Vendor(n)] = v
		}
	}
	if u.Username != "" {
		c.Default = &UserPass{Username: u.Username, Password: u.Password}
	}
	return nil
}

// DefineFlags defines flags for allowing user-specified credentials.
func DefineFlags() *Flags {
	flags := new(Flags)
	multiStringVar(&flags.nodeCreds, "node_creds",
		"Repeated node-specific credentials in the form 'nodeName/username/password', e.g., 'n1/admin/hunter2'")
	multiStringVar(&flags.vendorCreds, "vendor_creds",
		"Repeated vendor-specific credentials in the form 'vendorName/username/password', e.g., 'ARISTA/admin/hunter2'")
	flag.StringVar(&flags.defaultCreds, "default_creds", "",
		"Default credentials in the form 'username/password', e.g., 'admin/hunter2'")
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
	nodeCreds    []string
	vendorCreds  []string
	defaultCreds string
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
	if len(f.vendorCreds) > 0 {
		c.Vendor = make(map[tpb.Vendor]*UserPass, len(f.vendorCreds))
		for _, vc := range f.vendorCreds {
			parts := strings.SplitN(vc, "/", 3)
			if len(parts) != 3 {
				return nil, fmt.Errorf("invalid vendor credentials: %q", vc)
			}
			vendor, ok := tpb.Vendor_value[parts[0]]
			if !ok {
				return nil, fmt.Errorf("invalid vendor: %q", parts[0])
			}
			c.Vendor[tpb.Vendor(vendor)] = &UserPass{Username: parts[1], Password: parts[2]}
		}
	}
	if f.defaultCreds != "" {
		parts := strings.SplitN(f.defaultCreds, "/", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid default credentials: %q", f.defaultCreds)
		}
		c.Default = &UserPass{Username: parts[0], Password: parts[1]}
	}
	return c, nil
}
