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

// Package orderedmap provides a map that remembers the order keys were inserted.
package orderedmap

// OrderedMap is a map that remembers the order that keys were assorted.
// To iterate in order, iterate over the keys and call Get on valueMap
type OrderedMap[K comparable, V any] struct {
	keys     []K
	valueMap map[K]V
}

// NewOrderedMap returns a new empty OrderedMap.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		valueMap: map[K]V{},
	}
}

// Keys returns the list of Keys to iterate over.
func (o *OrderedMap[K, V]) Keys() []K { return o.keys }

// Get returns a value corresponding the key. If the key is not found, the zero
// value of K is returned with found being false.
// Get is O(1).
func (o *OrderedMap[K, V]) Get(key K) (value V, found bool) {
	val, ok := o.valueMap[key]
	return val, ok
}

// Set stores the value and stores the order in which it was set.
// If this overrides a value, also update the order the keys were set.
// Set is O(n).
func (o *OrderedMap[K, V]) Set(key K, value V) {
	if _, ok := o.valueMap[key]; !ok {
		o.keys = append(o.keys, key)
	}
	o.valueMap[key] = value
}

// Delete removes the key and value from OrderedMap.
// Delete is O(n).
func (o *OrderedMap[K, V]) Delete(key K) {
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			delete(o.valueMap, key)
			break
		}
	}
}

// Len returns a size of OrderedMap.
func (o *OrderedMap[K, V]) Len() int {
	return len(o.keys)
}
