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

package ixweb

import (
	"path"

	"golang.org/x/net/context"
)

// Chassis represents the chassis management application on Ixia Web Platform.
type Chassis struct {
	ixweb *IxWeb
}

func (c *Chassis) absPath(relPath string) string {
	return path.Join("/chassis/api/v2", relPath)
}

// Get submits a GET request, at a relative path to the chassis API.
func (c *Chassis) Get(ctx context.Context, path string, out any) error {
	return c.ixweb.jsonReq(ctx, get, c.absPath(path), nil, out)
}

// Delete submits a DELETE request, at a relative path to the chassis API.
func (c *Chassis) Delete(ctx context.Context, path string) error {
	return c.ixweb.jsonReq(ctx, delete, c.absPath(path), nil, nil)
}

// Patch submits a PATCH request, at a relative path to the chassis API.
func (c *Chassis) Patch(ctx context.Context, path string, in any) error {
	return c.ixweb.jsonReq(ctx, patch, c.absPath(path), in, nil)
}

// Post submits a POST request, at a relative path to the chassis API.
func (c *Chassis) Post(ctx context.Context, path string, in, out any) error {
	return c.ixweb.jsonReq(ctx, post, c.absPath(path), in, out)
}
