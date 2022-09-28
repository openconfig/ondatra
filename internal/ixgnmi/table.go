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

package ixgnmi

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// table is an Ixia API table.
type table struct {
	ID       int        `json:"id"`
	Type     string     `json:"type"`
	Columns  []string   `json:"columns"`
	Values   [][]string `json:"values"`
	RowCount int        `json:"rowCount"`
}

// unmarshalTable converts the input table in to slice of structs.
// s must be a pointer to a slice of structs or pointer to a slice of struct pointers.
// The structs must have "ixia" tags corresponding to the column names in the table.
func unmarshalTable(t *table, s any) error {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Ptr || rt.Elem().Kind() != reflect.Slice {
		return errors.New("must pass in a pointer to slice")
	}

	slice := reflect.ValueOf(s).Elem()
	newSlice := reflect.MakeSlice(slice.Type(), len(t.Values), len(t.Values))

	for i := range t.Values {
		if err := unmarshalTableRow(t, i, newSlice.Index(i)); err != nil {
			return fmt.Errorf("error unmarshaling row %d: %w", i, err)
		}
	}

	slice.Set(newSlice)
	return nil
}

func unmarshalTableRow(t *table, index int, rv reflect.Value) error {
	if rv.Kind() == reflect.Ptr {
		rv.Set(reflect.New(rv.Type().Elem()))
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct type got: %v", rv.Kind())
	}

	tm := map[string]int{}
	for i, key := range t.Columns {
		tm[key] = i
	}

	for i := 0; i < rv.Type().NumField(); i++ {
		colName, ok := rv.Type().Field(i).Tag.Lookup("ixia")
		if !ok {
			continue
		}
		colIndex, ok := tm[colName]
		if !ok {
			continue
		}
		fieldV := rv.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		strVal := t.Values[index][colIndex]
		if strVal == "removePacket[ ]" || strVal == "removePacket[N/A]" || strVal == "N/A" {
			continue
		}
		// Set the field value based on its column index and parse any integers.
		switch kind := fieldV.Kind(); kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if strVal == "NA" || strVal == "" {
				continue
			}
			val, err := strconv.ParseInt(strVal, 10, 0)
			if err != nil {
				return fmt.Errorf("parse field %q failed: %w", rv.Type().Field(i).Name, err)
			}
			fieldV.SetInt(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if strVal == "NA" || strVal == "" {
				continue
			}
			val, err := strconv.ParseUint(strVal, 10, 0)
			if err != nil {
				return fmt.Errorf("parse field %q failed: %w", rv.Type().Field(i).Name, err)
			}
			fieldV.SetUint(val)
		case reflect.String:
			fieldV.SetString(strVal)
		default:
			return fmt.Errorf("unsupported field type: %q, value: %q", kind, strVal)
		}

	}
	return nil
}
