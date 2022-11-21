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

package orderedmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrderedMapSet(t *testing.T) {
	tests := []struct {
		desc                               string
		keysToAdd, valuesToAdd             []string
		wantOrderedKeys, wantOrderedValues []string
	}{{
		desc:              "Add unique keys",
		keysToAdd:         []string{"one", "two", "three", "four", "five"},
		valuesToAdd:       []string{"1", "2", "3", "4", "5"},
		wantOrderedKeys:   []string{"one", "two", "three", "four", "five"},
		wantOrderedValues: []string{"1", "2", "3", "4", "5"},
	}, {
		desc:              "Add overlapping keys",
		keysToAdd:         []string{"one", "two", "three", "one", "four"},
		valuesToAdd:       []string{"1", "2", "3", "11", "4"},
		wantOrderedKeys:   []string{"one", "two", "three", "four"},
		wantOrderedValues: []string{"11", "2", "3", "4"},
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			om := NewOrderedMap[string, string]()
			for i, key := range tc.keysToAdd {
				om.Set(key, tc.valuesToAdd[i])
			}
			if diff := cmp.Diff(tc.wantOrderedKeys, om.Keys()); diff != "" {
				t.Errorf("OrderedMap keys mismatched (-want + got):\n%s", diff)
			}
			for i, key := range tc.wantOrderedKeys {
				got, _ := om.Get(key)
				want := tc.wantOrderedValues[i]
				if got != want {
					t.Errorf("OrderedMap[%s] got %s, want %s", key, got, want)
				}
			}
		})
	}
}

func TestOrderedMapDelete(t *testing.T) {
	om := NewOrderedMap[string, string]()
	om.Set("one", "1")
	om.Set("two", "2")
	om.Set("three", "3")
	om.Set("four", "4")
	om.Set("five", "5")
	tests := []struct {
		desc                               string
		om                                 *OrderedMap[string, string]
		keyToDelete                        string
		wantOrderedKeys, wantOrderedValues []string
	}{{
		desc:              "Delete key",
		om:                om,
		keyToDelete:       "three",
		wantOrderedKeys:   []string{"one", "two", "four", "five"},
		wantOrderedValues: []string{"1", "2", "4", "5"},
	}, {
		desc:              "Delete non-existant key",
		om:                om,
		keyToDelete:       "does not exist",
		wantOrderedKeys:   []string{"one", "two", "four", "five"},
		wantOrderedValues: []string{"1", "2", "4", "5"},
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			om.Delete(tc.keyToDelete)
			if diff := cmp.Diff(tc.wantOrderedKeys, om.Keys()); diff != "" {
				t.Errorf("OrderedMap keys mismatched (-want + got):\n%s", diff)
			}
			if _, ok := om.Get(tc.keyToDelete); ok {
				t.Errorf("OrderedMap got value for key %s, want no value", tc.keyToDelete)
			}
			for i, key := range tc.wantOrderedKeys {
				got, _ := om.Get(key)
				want := tc.wantOrderedValues[i]
				if got != want {
					t.Errorf("OrderedMap[%s] got %s, want %s", key, got, want)
				}
			}
		})
	}
}
