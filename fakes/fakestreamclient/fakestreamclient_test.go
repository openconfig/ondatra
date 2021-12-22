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
package fakestreamclient

import (
	"sync"
	"testing"
)

func TestFake(t *testing.T) {
	f := New()

	inW := RandSeq(1000)
	inR := make([]byte, 4096)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		f.Stdin().Write([]byte(inW))
		f.OutWriter.Write([]byte(inW))
		f.ErrWriter.Write([]byte(inW))
		wg.Done()
	}()
	wg.Add(1)
	var gotIn, gotOut, gotErr string
	go func() {
		var n int
		n, _ = f.InReader.Read(inR)
		gotIn = string(inR[:n])
		n, _ = f.Stdout().Read(inR)
		gotOut = string(inR[:n])
		n, _ = f.Stderr().Read(inR)
		gotErr = string(inR[:n])
		wg.Done()
	}()
	wg.Wait()
	if gotIn != inW {
		t.Errorf("Stdin loop failed: got %v, want %v", gotIn, inW)
	}
	if gotOut != inW {
		t.Errorf("Stdout loop failed: got %v, want %v", gotOut, inW)
	}
	if gotErr != inW {
		t.Errorf("Stderr loop failed: got %v, want %v", gotErr, inW)
	}
}
