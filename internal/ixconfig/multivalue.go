// Copyright 2021 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixconfig

import (
	"strconv"
)

// String stores v in a new string value and returns a pointer to it.
func String(v string) *string { return &v }

// Bool stores v in a new bool value and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// NumberInt stores v in a new float32 value and returns a pointer to it.
func NumberInt(v int) *float32 {
	f := float32(v)
	return &f
}

// NumberFloat64 stores v in a new float32 value and returns a pointer to it.
func NumberFloat64(v float64) *float32 {
	f := float32(v)
	return &f
}

// NumberUint32 stores v in a new float32 value and returns a pointer to it.
func NumberUint32(v uint32) *float32 {
	f := float32(v)
	return &f
}

// NumberUint64 stores v in a new float32 value and returns a pointer to it.
func NumberUint64(v uint64) *float32 {
	f := float32(v)
	return &f
}

// MultivalueStr returns a single-valued Multivalue whose value is the given string.
func MultivalueStr(val string) *Multivalue {
	return &Multivalue{
		SingleValue: &MultivalueSingleValue{
			Value: String(val),
		},
	}
}

// MultivalueBool returns a single-valued Multivalue whose value is the given bool.
func MultivalueBool(val bool) *Multivalue {
	return MultivalueStr(strconv.FormatBool(val))
}

// MultivalueTrue returns a single-valued Multivalue representing a boolean 'true'.
func MultivalueTrue() *Multivalue {
	return MultivalueStr("true")
}

// MultivalueFalse returns a single-valued Multivalue representing a boolean 'false'.
func MultivalueFalse() *Multivalue {
	return MultivalueStr("false")
}

// MultivalueUint32 returns a single-valued Multivalue whose value is the given int.
func MultivalueUint32(val uint32) *Multivalue {
	return MultivalueStr(strconv.FormatUint(uint64(val), 10))
}

// MultivalueStrList returns a list-valued Multivalue containing all given strings.
func MultivalueStrList(vals ...string) *Multivalue {
	mv := &MultivalueValueList{}
	for _, v := range vals {
		mv.Values = append(mv.Values, v)
	}
	return &Multivalue{Pattern: String("valueList"), ValueList: mv}
}

// MultivalueBoolList returns a list-valued Multivalue containing all given bools.
func MultivalueBoolList(vals ...bool) *Multivalue {
	mv := &MultivalueValueList{}
	for _, v := range vals {
		mv.Values = append(mv.Values, strconv.FormatBool(v))
	}
	return &Multivalue{Pattern: String("valueList"), ValueList: mv}
}

// MultivalueUintList returns a list-valued Multivalue containing all given uints.
func MultivalueUintList(vals ...uint32) *Multivalue {
	mv := &MultivalueValueList{}
	for _, v := range vals {
		mv.Values = append(mv.Values, strconv.FormatUint(uint64(v), 10))
	}
	return &Multivalue{Pattern: String("valueList"), ValueList: mv}
}

// MultivalueIntList returns a list-valued Multivalue containing all given ints.
func MultivalueIntList(vals ...int) *Multivalue {
	mv := &MultivalueValueList{}
	for _, v := range vals {
		mv.Values = append(mv.Values, strconv.Itoa(v))
	}
	return &Multivalue{Pattern: String("valueList"), ValueList: mv}
}

// MultivalueUintIncCounter returns an incrementing Multivalue with the given
// start/step uint values.
func MultivalueUintIncCounter(start, step uint32) *Multivalue {
	return MultivalueStrIncCounter(
		strconv.FormatUint(uint64(start), 10),
		strconv.FormatUint(uint64(step), 10))
}

// MultivalueStrIncCounter returns an incrementing Multivalue with the given
// start/step string values.
func MultivalueStrIncCounter(start, step string) *Multivalue {
	return &Multivalue{
		Pattern: String("counter"),
		Counter: &MultivalueCounter{
			Direction: String("increment"),
			Start:     String(start),
			Step:      String(step),
		},
	}
}
