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
	"net/http"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
)

func TestList(t *testing.T) {
	files := &Files{sess: &Session{ixweb: &IxWeb{
		client: &fakeHTTPClient{doResps: []*http.Response{
			fakeResponse(200, `{"files": [{"name": "file1"}, {"name": "file2"}]}`),
		}},
	}}}
	want := []string{"file1", "file2"}
	got, err := files.List(context.Background(), "pattern")
	if err != nil {
		t.Fatalf("List got unexpected error: %v", err)
	}
	if !cmp.Equal(want, got) {
		t.Fatalf("List got unexpected files, got %v, want %v", got, want)
	}
}

func TestUpload(t *testing.T) {
	files := &Files{sess: &Session{ixweb: &IxWeb{
		client: &fakeHTTPClient{doResps: []*http.Response{
			fakeResponse(200, ""),
		}},
	}}}
	err := files.Upload(context.Background(), "myFile", []byte("myContent"))
	if err != nil {
		t.Fatalf("Upload got unexpected error: %v", err)
	}
}

func TestDownload(t *testing.T) {
	files := &Files{sess: &Session{ixweb: &IxWeb{
		client: &fakeHTTPClient{doResps: []*http.Response{
			fakeResponse(200, "myContent"),
		}},
	}}}
	content, err := files.Download(context.Background(), "myFile")
	if err != nil {
		t.Fatalf("Download got unexpected error: %v", err)
	}
	if got, want := string(content), "myContent"; got != want {
		t.Fatalf("Download got unexpected content, got %v, want %v", got, want)
	}
}

func TestDelete(t *testing.T) {
	files := &Files{sess: &Session{ixweb: &IxWeb{
		client: &fakeHTTPClient{doResps: []*http.Response{
			fakeResponse(200, ""),
		}},
	}}}
	err := files.Delete(context.Background(), "myFile")
	if err != nil {
		t.Fatalf("Delete got unexpected error: %v", err)
	}
}
