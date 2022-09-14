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

package portgraph

// genCombos yields all key->value assignment maps, given the specified range
// of possible values for each key, such that no two keys are assigned the same
// value. Returns a channel that yields the assignments and a function to stop
// the generation. Caller should immediately defer a call to the stop function
// to ensure the routine generating the assignments exits.
func genCombos[K, V comparable](m map[K][]V) (<-chan map[K]V, func()) {
	gen := &generator[K, V]{
		m:    m,
		res:  make(map[K]V),
		used: make(map[V]bool),
		ch:   make(chan map[K]V),
		stop: make(chan struct{}, 1),
		done: make(chan struct{}, 1),
	}
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	go func() {
		defer close(gen.ch)
		gen.recurse(keys)
		gen.done <- struct{}{}
	}()
	return gen.ch, gen.stopAndWait
}

type generator[K, V comparable] struct {
	m    map[K][]V
	res  map[K]V
	used map[V]bool
	ch   chan map[K]V
	stop chan struct{}
	done chan struct{}
}

// stopAndWait signals the generation to stop and waits for it to stop.
func (g *generator[K, V]) stopAndWait() {
	g.stop <- struct{}{}
	<-g.done
}

// recurse generates all the map combinations for the specified list of keys,
// returning true until it receives a stop signal.
func (g *generator[K, V]) recurse(keys []K) bool {
	if len(keys) == 0 {
		copy := make(map[K]V, len(g.m))
		for k, v := range g.res {
			copy[k] = v
		}
		select {
		case <-g.stop:
			return false
		case g.ch <- copy:
			return true
		}
	}
	first := keys[0]
	for _, i := range g.m[first] {
		if !g.used[i] {
			g.res[first] = i
			g.used[i] = true
			if !g.recurse(keys[1:]) {
				return false
			}
			delete(g.used, i)
		}
	}
	return true
}
