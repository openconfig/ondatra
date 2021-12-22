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
)

// Files represents the IxNetwork files API.
type Files struct {
	sess *Session
}

// List returns a slice of the files that match the specified glob pattern.
func (f *Files) List(ctx context.Context, pattern string) ([]string, error) {
	out := struct {
		Files []struct {
			Name string `json:"name"`
		} `json:"files"`
	}{}
	if err := f.sess.jsonReq(ctx, get, "files?filter="+pattern, nil, &out); err != nil {
		return nil, err
	}
	var fileNames []string
	for _, f := range out.Files {
		fileNames = append(fileNames, f.Name)
	}
	return fileNames, nil
}

// UploadFile uploads a file with the specified name and content to the session.
func (f *Files) Upload(ctx context.Context, filename string, content []byte) error {
	_, err := f.sess.binaryReq(ctx, post, "files?filename="+filename, content)
	return err
}

// DownloadFile downloads a file with the specified name and returns the content.
func (f *Files) Download(ctx context.Context, filename string) ([]byte, error) {
	return f.sess.binaryReq(ctx, get, "files?filename="+filename, nil)
}

// Delete deletes a file with the specified name.
func (f *Files) Delete(ctx context.Context, filename string) error {
	_, err := f.sess.binaryReq(ctx, delete, "files?filename="+filename, nil)
	return err
}
