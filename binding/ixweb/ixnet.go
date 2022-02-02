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
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	sessionsPath = "/api/v1/sessions"
)

func sessionPath(id int) string {
	return path.Join(sessionsPath, strconv.Itoa(id))
}

var (
	sessionDisallowRE = regexp.MustCompile(`[<>:"/\|?*]`)
)

// IxNetwork represents the IxNetwork application on Ixia Web Platform.
type IxNetwork struct {
	ixweb *IxWeb
}

// NewSession creates a new IxNetwork session with the specified name.
func (n *IxNetwork) NewSession(ctx context.Context, name string) (*Session, error) {
	in := struct {
		ApplicationType string `json:"applicationType"`
	}{
		ApplicationType: "ixnrest",
	}
	out := struct {
		ID int `json:"id"`
	}{}
	if err := n.ixweb.jsonReq(ctx, post, sessionsPath, in, &out); err != nil {
		return nil, fmt.Errorf("Error creating session: %w", err)
	}
	data := sessionData{
		Name: encodeSessionName(name),
	}
	spath := sessionPath(out.ID)
	if err := n.ixweb.jsonReq(ctx, patch, spath, data, nil); err != nil {
		return nil, fmt.Errorf("Error naming session: %w", err)
	}
	if err := n.ixweb.jsonReq(ctx, post, path.Join(spath, "operations/start"), nil, nil); err != nil {
		return nil, fmt.Errorf("Error starting session: %w", err)
	}
	return &Session{ixweb: n.ixweb, id: out.ID, name: name}, nil
}

type sessionData struct {
	Name string `json:"sessionName"`
}

// encodeSessionName encodes the given session name into an allowed format.
func encodeSessionName(s string) string {
	const ellipses = "..."
	s = sessionDisallowRE.ReplaceAllLiteralString(s, "_")
	if maxLen := 64; len(s) > maxLen {
		s = s[:maxLen-len(ellipses)] + ellipses
	}
	return s
}

// FetchSession fetches the IxNetwork session with the specified ID.
func (n *IxNetwork) FetchSession(ctx context.Context, id int) (*Session, error) {
	data := sessionData{}
	if err := n.ixweb.jsonReq(ctx, get, sessionPath(id), nil, &data); err != nil {
		return nil, fmt.Errorf("Error deleting session: %w", err)
	}
	return &Session{ixweb: n.ixweb, id: id, name: data.Name}, nil
}

// DeleteSession deletes the IxNetwork session with the specified ID.
func (n *IxNetwork) DeleteSession(ctx context.Context, id int) error {
	spath := sessionPath(id)
	if err := n.ixweb.jsonReq(ctx, post, path.Join(spath, "operations/stop"), nil, nil); err != nil {
		return fmt.Errorf("Error stopping session: %w", err)
	}
	if err := n.ixweb.jsonReq(ctx, delete, spath, nil, nil); err != nil {
		return fmt.Errorf("Error deleting session: %w", err)
	}
	return nil
}

// Session represents an IxNetwork session.
type Session struct {
	ixweb *IxWeb
	id    int
	name  string

	// IxNetwork doesn't like concurrent HTTP requests to the same session.
	// Use mutex to prevent concurrency.
	mu sync.Mutex
}

// ID returns the unique ID of the session.
func (s *Session) ID() int {
	return s.id
}

// Name returns the descriptive name of the session.
func (s *Session) Name() string {
	return s.name
}

func (s *Session) String() string {
	return fmt.Sprintf("Session %d %s", s.id, s.name)
}

// AbsPath returns an absolute path, given a path relative to the IxNetwork session.
// If relPath is already an absolute path, this method returns the input string.
func (s *Session) AbsPath(relPath string) string {
	absPrefix := path.Join(sessionPath(s.id), "ixnetwork")
	if strings.HasPrefix(relPath, absPrefix) {
		return relPath
	}
	return path.Join(absPrefix, relPath)
}

// Get submits a JSON GET request, at a path relative to the IxNetwork session.
func (s *Session) Get(ctx context.Context, path string, out interface{}) error {
	return s.jsonReq(ctx, get, path, nil, out)
}

// Delete submits a JSON DELETE request, at a path relative to the IxNetwork session.
func (s *Session) Delete(ctx context.Context, path string) error {
	// Content-type (JSON or binary) does not matter for a DELETE.
	return s.jsonReq(ctx, delete, path, nil, nil)
}

// Patch submits a JSON PATCH request, at a path relative to the IxNetwork session.
func (s *Session) Patch(ctx context.Context, path string, in interface{}) error {
	return s.jsonReq(ctx, patch, path, in, nil)
}

// Post submits a JSON POST request, at a path relative to the IxNetwork session.
func (s *Session) Post(ctx context.Context, path string, in, out interface{}) error {
	return s.jsonReq(ctx, post, path, in, out)
}

// Config returns the config API for this session.
func (s *Session) Config() *Config {
	return &Config{sess: s}
}

// Files returns the file API for this session.
func (s *Session) Files() *Files {
	return &Files{sess: s}
}

// Stats returns the statistics API for this session.
func (s *Session) Stats() *Stats {
	return &Stats{sess: s}
}

func (s *Session) jsonReq(ctx context.Context, method httpMethod, path string, in, out interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ixweb.jsonReq(ctx, method, s.AbsPath(path), in, out)
}

func (s *Session) binaryReq(ctx context.Context, method httpMethod, path string, content []byte) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ixweb.binaryReq(ctx, method, s.AbsPath(path), content)
}

// OpArgs is a list of arguments for an operation to use as input to a POST request.
type OpArgs []interface{}

// MarshalJSON marshals the operation arguments to JSON.
func (a OpArgs) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for i, v := range a {
		m["arg"+strconv.Itoa(i+1)] = v
	}
	return json.Marshal(m)
}
