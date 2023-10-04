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

package binding

import (
	"fmt"
	"reflect"
)

// DUTAs tests if the specified DUT or any DUT it transitively embeds fulfills
// the interface pointed to by target, and if so, sets target to that DUT.
func DUTAs(dut DUT, target any) error {
	return as(dut, target, "DUT", func(v any) (DUT, bool) {
		d, ok := v.(DUT)
		return d, ok
	})
}

// ATEAs tests if the specified ATE or any ATE it transitively embeds fulfills
// the interface pointed to by target, and if so, sets target to that ATE.
func ATEAs(ate ATE, target any) error {
	return as(ate, target, "ATE", func(v any) (ATE, bool) {
		d, ok := v.(ATE)
		return d, ok
	})
}

func as[T Device](src T, target any, field string, assert func(any) (T, bool)) error {
	targetPtrVal, ok := nonNilPtr(target)
	if !ok {
		return fmt.Errorf("target must be non-nil pointer, got %v (type: %T)", target, target)
	}
	targetVal := targetPtrVal.Elem()
	srcVal, ok := nonNilPtr(src)
	if !ok {
		return fmt.Errorf("src must be non-nil pointer, got %v (type: %T)", src, src)
	}

	for curr, currVal := src, srcVal; ; {
		if currVal.Type().AssignableTo(targetVal.Type()) {
			targetVal.Set(currVal)
			return nil
		}
		embedField := currVal.Elem().FieldByName(field)
		if !embedField.IsValid() {
			return fmt.Errorf("no match found in %v (type: %T)", src, src)
		}
		intf := embedField.Interface()
		if intf == nil {
			return fmt.Errorf("no match found in %v (type: %T)", src, src)
		}
		curr, ok = assert(intf)
		if !ok {
			return fmt.Errorf("embedded field %v must be of type binding.%v, got %T", field, field, intf)
		}
		currVal = reflect.ValueOf(curr)
	}
}

func nonNilPtr(v any) (reflect.Value, bool) {
	val := reflect.ValueOf(v)
	ok := val.IsValid() && val.Type().Kind() == reflect.Ptr && !val.IsNil()
	return val, ok
}
