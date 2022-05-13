// Copyright 2019 Google LLC
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

package genutil

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/goyang/pkg/yang"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"github.com/openconfig/gnmi/errdiff"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var globalEnumMap = map[string]map[int64]ygot.EnumDefinition{
	"EnumType": {
		43:  {Name: "E_VALUE_FORTY_THREE"},
		44:  {Name: "E_VALUE_FORTY_FOUR"},
		100: {Name: "E_VALUE_ONE_HUNDRED"},
	},
	"EnumType2": {
		42: {Name: "E_VALUE_FORTY_TWO"},
	},
}

// EnumType is used as an enum type in various tests in the ytypes package.
type EnumType int64

func (EnumType) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return globalEnumMap
}

func (EnumType) IsYANGGoEnum() {}

// EnumType2 is used as an enum type in various tests in the ytypes package.
type EnumType2 int64

func (EnumType2) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return globalEnumMap
}

func (EnumType2) IsYANGGoEnum() {}

type UnionLeafType interface {
	Is_UnionLeafType()
}

type UnionLeafType_String struct {
	String string
}

func (*UnionLeafType_String) Is_UnionLeafType() {}

type UnionLeafType_Uint32 struct {
	Uint32 uint32
}

func (*UnionLeafType_Uint32) Is_UnionLeafType() {}

type UnionLeafType_EnumType struct {
	EnumType EnumType
}

func (*UnionLeafType_EnumType) Is_UnionLeafType() {}

func (*UnionLeafType_EnumType) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return globalEnumMap
}

type UnionLeafType_EnumType2 struct {
	EnumType2 EnumType2
}

func (*UnionLeafType_EnumType2) Is_UnionLeafType() {}

func (*UnionLeafType_EnumType2) ΛMap() map[string]map[int64]ygot.EnumDefinition {
	return globalEnumMap
}

// Device is the fake root.
type Device struct {
	SuperContainer *SuperContainer `path:"super-container" module:"yang-module"`
}

func (*Device) IsYANGGoStruct()                               {}
func (*Device) ΛValidate(opts ...ygot.ValidationOption) error { return nil }
func (*Device) ΛEnumTypeMap() map[string][]reflect.Type       { return nil }
func (*Device) ΛBelongingModule() string                      { return "" }

func unmarshalFunc([]byte, ygot.GoStruct, ...ytypes.UnmarshalOpt) error { return nil }

type SuperContainer struct {
	LeafContainerStruct *LeafContainerStruct `path:"leaf-container-struct"`
	Model               *Model               `path:"model"`
}

func (*SuperContainer) IsYANGGoStruct()                          {}
func (*SuperContainer) ΛValidate(...ygot.ValidationOption) error { return nil }
func (*SuperContainer) ΛEnumTypeMap() map[string][]reflect.Type  { return nil }
func (*SuperContainer) ΛBelongingModule() string                 { return "" }

type LeafContainerStruct struct {
	Uint64Leaf          *uint64       `path:"uint64-leaf"`
	EnumLeaf            EnumType      `path:"enum-leaf"`
	UnionLeaf           UnionLeafType `path:"union-leaf"`
	UnionLeaf2          EnumType      `path:"union-leaf2"`
	UnionLeafSingleType []string      `path:"union-stleaflist"`
}

func (*LeafContainerStruct) IsYANGGoStruct()                          {}
func (*LeafContainerStruct) ΛValidate(...ygot.ValidationOption) error { return nil }
func (*LeafContainerStruct) ΛBelongingModule() string                 { return "" }

type Model struct {
	SingleKey map[int32]*Model_SingleKey `path:"a/single-key"`
}

func (*Model) IsYANGGoStruct()                          {}
func (*Model) ΛValidate(...ygot.ValidationOption) error { return nil }
func (*Model) ΛEnumTypeMap() map[string][]reflect.Type  { return nil }
func (*Model) ΛBelongingModule() string                 { return "" }

type Model_SingleKey struct {
	Key   *int32 `path:"config/key|key" shadow-path:"state/key"`
	Value *int64 `path:"config/value" shadow-path:"state/value"`
}

func (*Model_SingleKey) IsYANGGoStruct()                          {}
func (*Model_SingleKey) ΛValidate(...ygot.ValidationOption) error { return nil }
func (*Model_SingleKey) ΛEnumTypeMap() map[string][]reflect.Type  { return nil }
func (*Model_SingleKey) ΛBelongingModule() string                 { return "" }

// NewSingleKey is a *generated* method for Model which may be used by an
// unmarshal function in ytype's reflect library, and is kept here in case.
func (t *Model) NewSingleKey(Key int32) (*Model_SingleKey, error) {
	// Initialise the list within the receiver struct if it has not already been
	// created.
	if t.SingleKey == nil {
		t.SingleKey = make(map[int32]*Model_SingleKey)
	}

	key := Key

	// Ensure that this key has not already been used in the
	// list. Keyed YANG lists do not allow duplicate keys to
	// be created.
	if _, ok := t.SingleKey[key]; ok {
		return nil, fmt.Errorf("duplicate key %q for list SingleKey", key)
	}

	t.SingleKey[key] = &Model_SingleKey{
		Key: &Key,
	}

	return t.SingleKey[key], nil
}

// RenameSingleKey is a *generated* method for Model which may be used by an
// unmarshal function in ytype's reflect library, and is kept here in case.
// RenameSingleKey renames an entry in the list SingleKey within
// the Model struct. The entry with key oldK is renamed to newK updating
// the key within the value.
func (t *Model) RenameSingleKey(oldK, newK int32) error {
	if _, ok := t.SingleKey[newK]; ok {
		return fmt.Errorf("key %q already exists in SingleKey", newK)
	}

	e, ok := t.SingleKey[oldK]
	if !ok {
		return fmt.Errorf("key %q not found in SingleKey", oldK)
	}
	e.Key = &newK

	t.SingleKey[newK] = e
	delete(t.SingleKey, oldK)
	return nil
}

// ΛListKeyMap returns the keys of the Model_SingleKey struct, which is a YANG list entry.
func (t *Model_SingleKey) ΛListKeyMap() (map[string]interface{}, error) {
	if t.Key == nil {
		return nil, fmt.Errorf("nil value for key Key")
	}

	return map[string]interface{}{
		"key": *t.Key,
	}, nil
}

func (*LeafContainerStruct) ΛEnumTypeMap() map[string][]reflect.Type {
	return map[string][]reflect.Type{
		"/super-container/leaf-container-struct/enum-leaf":   {reflect.TypeOf(EnumType(0))},
		"/super-container/leaf-container-struct/union-leaf":  {reflect.TypeOf(EnumType(0)), reflect.TypeOf(EnumType2(0))},
		"/super-container/leaf-container-struct/union-leaf2": {reflect.TypeOf(EnumType(0))},
	}
}

func (*LeafContainerStruct) To_UnionLeafType(i interface{}) (UnionLeafType, error) {
	switch v := i.(type) {
	case string:
		return &UnionLeafType_String{v}, nil
	case uint32:
		return &UnionLeafType_Uint32{v}, nil
	case EnumType:
		return &UnionLeafType_EnumType{v}, nil
	case EnumType2:
		return &UnionLeafType_EnumType2{v}, nil
	default:
		return nil, fmt.Errorf("cannot convert %v to To_UnionLeafType, unknown union type, got: %T, want any of [string, uint32]", i, i)
	}
}

// addParents adds parent pointers for a schema tree.
func addParents(e *yang.Entry) {
	for _, c := range e.Dir {
		c.Parent = e
		addParents(c)
	}
}

// mustPath converts a string to its path proto.
func gnmiPath(t *testing.T, s string) *gpb.Path {
	p, err := ygot.StringToStructuredPath(s)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func TestUnmarshal(t *testing.T) {
	rootSchema := &yang.Entry{
		Name:       "device",
		Kind:       yang.DirectoryEntry,
		Annotation: map[string]interface{}{"isFakeRoot": true},
		Dir: map[string]*yang.Entry{
			"super-container": {
				Name: "super-container",
				Kind: yang.DirectoryEntry,
				Dir: map[string]*yang.Entry{
					"model": {
						Name: "model",
						Kind: yang.DirectoryEntry,
						Dir: map[string]*yang.Entry{
							"a": {
								Name: "a",
								Kind: yang.DirectoryEntry,
								Dir: map[string]*yang.Entry{
									"single-key": {
										Name:     "single-key",
										Kind:     yang.DirectoryEntry,
										ListAttr: yang.NewDefaultListAttr(),
										Key:      "key",
										Dir: map[string]*yang.Entry{
											"key": {
												Name: "key",
												Kind: yang.LeafEntry,
												Type: &yang.YangType{Kind: yang.Yleafref, Path: "../config/key"},
											},
											"config": {
												Name: "config",
												Kind: yang.DirectoryEntry,
												Dir: map[string]*yang.Entry{
													"key": {
														Name: "key",
														Kind: yang.LeafEntry,
														Type: &yang.YangType{Kind: yang.Yint32},
													},
													"value": {
														Name: "value",
														Kind: yang.LeafEntry,
														Type: &yang.YangType{Kind: yang.Yint64},
													},
												},
											},
											"state": {
												Name: "state",
												Kind: yang.DirectoryEntry,
												Dir: map[string]*yang.Entry{
													"key": {
														Name: "key",
														Kind: yang.LeafEntry,
														Type: &yang.YangType{Kind: yang.Yint32},
													},
													"value": {
														Name: "value",
														Kind: yang.LeafEntry,
														Type: &yang.YangType{Kind: yang.Yint64},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"leaf-container-struct": {
						Name: "leaf-container-struct",
						Kind: yang.DirectoryEntry,
						Dir: map[string]*yang.Entry{
							"uint64-leaf": {
								Name: "uint64-leaf",
								Kind: yang.LeafEntry,
								Type: &yang.YangType{
									Kind: yang.Yuint64,
								},
							},
							"enum-leaf": {
								Name: "enum-leaf",
								Kind: yang.LeafEntry,
								Type: &yang.YangType{
									Kind: yang.Yenum,
								},
							},
							"union-leaf": {
								Name: "union-leaf",
								Kind: yang.LeafEntry,
								Type: &yang.YangType{
									Kind: yang.Yunion,
									Type: []*yang.YangType{
										{
											Kind:    yang.Ystring,
											Pattern: []string{"a+"},
										},
										{
											Kind: yang.Yuint32,
										},
										{
											Kind: yang.Yenum,
										},
										{
											Kind: yang.Yenum,
										},
										{
											Kind: yang.Yleafref,
											Path: "../enum-leaf",
										},
									},
								},
							},
							"union-leaf2": {
								Name: "union-leaf2",
								Kind: yang.LeafEntry,
								Type: &yang.YangType{
									Kind: yang.Yunion,
									Type: []*yang.YangType{
										{
											Kind: yang.Yenum,
										},
									},
								},
							},
							"union-stleaflist": {
								Name:     "union-stleaflist",
								Kind:     yang.LeafEntry,
								ListAttr: yang.NewDefaultListAttr(),
								Type: &yang.YangType{
									Kind: yang.Yunion,
									Type: []*yang.YangType{
										{
											// Note that Validate is not called as part of Unmarshal,
											// therefore any string pattern will actually match.
											Kind:    yang.Ystring,
											Pattern: []string{"a+"},
										},
										{
											Kind:    yang.Ystring,
											Pattern: []string{"b+"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	addParents(rootSchema)

	schemaStruct := func() *ytypes.Schema {
		return &ytypes.Schema{
			Root:       &Device{},
			SchemaTree: map[string]*yang.Entry{"Device": rootSchema},
			Unmarshal:  unmarshalFunc,
		}
	}

	superContainerSchema := rootSchema.Dir["super-container"]

	passingTests := []struct {
		name                 string
		inData               []*DataPoint
		inQueryPath          *gpb.Path
		inStructSchema       *yang.Entry
		inStruct             ygot.GoStruct
		inLeaf               bool
		inPreferShadowPath   bool
		wantUnmarshalledData []*DataPoint
		wantStruct           ygot.GoStruct
	}{{
		name: "retrieve uint64",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		wantStruct: &LeafContainerStruct{Uint64Leaf: ygot.Uint64(43)},
	}, {
		name: "retrieve uint64 into fake root",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: rootSchema,
		inStruct:       &Device{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		wantStruct: &Device{SuperContainer: &SuperContainer{LeafContainerStruct: &LeafContainerStruct{Uint64Leaf: ygot.Uint64(43)}}},
	}, {
		name: "successfully retrieve uint64 with positive int",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 42}},
			Timestamp: time.Unix(1, 1),
		}},
		wantStruct: &LeafContainerStruct{Uint64Leaf: ygot.Uint64(42)},
	}, {
		name: "successfully retrieve uint64 with zero",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 0}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 0}},
			Timestamp: time.Unix(1, 1),
		}},
		wantStruct: &LeafContainerStruct{Uint64Leaf: ygot.Uint64(0)},
	}, {
		name: "delete uint64",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{Uint64Leaf: ygot.Uint64(0)},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Timestamp: time.Unix(1, 1),
		}},
		wantStruct: &LeafContainerStruct{},
	}, {
		name: "retrieve union",
		inData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "aaaa",
				},
			},
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "aaaa",
				},
			},
			Timestamp: time.Unix(2, 2),
		}},
		wantStruct: &LeafContainerStruct{UnionLeaf: &UnionLeafType_String{"aaaa"}},
	}, {
		name: "delete union",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{UnionLeaf: &UnionLeafType_String{"forty two"}},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Timestamp: time.Unix(2, 2),
		}},
		wantStruct: &LeafContainerStruct{},
	}, {
		name: "delete union that's already deleted",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Timestamp: time.Unix(2, 2),
		}},
		wantStruct: &LeafContainerStruct{},
	}, {
		name: "retrieve union with a single enum inside",
		inData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf2"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_FORTY_FOUR",
				},
			},
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf2"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf2"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_FORTY_FOUR",
				},
			},
			Timestamp: time.Unix(2, 2),
		}},
		wantStruct: &LeafContainerStruct{UnionLeaf2: EnumType(44)},
	}, {
		name: "retrieve leaflist",
		inData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "aaaaa"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "b"}},
						},
					},
				},
			},
			Timestamp: time.Unix(1, 3),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{UnionLeaf2: EnumType(44)},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "aaaaa"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "b"}},
						},
					},
				},
			},
			Timestamp: time.Unix(1, 3),
		}},
		wantStruct: &LeafContainerStruct{
			UnionLeaf2:          EnumType(44),
			UnionLeafSingleType: []string{"aaaaa", "b"},
		},
	}, {
		name: "delete leaflist",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Timestamp: time.Unix(1, 3),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct: &LeafContainerStruct{
			UnionLeaf2:          EnumType(44),
			UnionLeafSingleType: []string{"forty two", "forty three"},
		},
		inLeaf: true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Timestamp: time.Unix(1, 3),
		}},
		wantStruct: &LeafContainerStruct{
			UnionLeaf2: EnumType(44),
		},
	}, {
		name: "retrieve leaf container, setting uint64 and enum",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(1, 1),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/enum-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_FORTY_THREE",
				},
			},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 42}},
			Timestamp: time.Unix(1, 1),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/enum-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_FORTY_THREE",
				},
			},
			Timestamp: time.Unix(1, 1),
		}},
		wantStruct: &LeafContainerStruct{
			Uint64Leaf: ygot.Uint64(42),
			EnumLeaf:   43,
		},
	}, {
		name: "retrieve leaf container, setting uint64, enum, union, union with single enum, and leaflist",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/enum-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_ONE_HUNDRED",
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_UintVal{
					UintVal: 100,
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf2"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_ONE_HUNDRED",
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "aa"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "bb"}},
						},
					},
				},
			},
			Timestamp: time.Unix(10, 10),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct: &LeafContainerStruct{
			Uint64Leaf: ygot.Uint64(42),
			EnumLeaf:   43,
		},
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 100}},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/enum-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_ONE_HUNDRED",
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_UintVal{
					UintVal: 100,
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf2"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "E_VALUE_ONE_HUNDRED",
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "aa"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "bb"}},
						},
					},
				},
			},
			Timestamp: time.Unix(10, 10),
		}},
		wantStruct: &LeafContainerStruct{
			Uint64Leaf:          ygot.Uint64(100),
			EnumLeaf:            EnumType(100),
			UnionLeaf:           &UnionLeafType_Uint32{100},
			UnionLeaf2:          EnumType(100),
			UnionLeafSingleType: []string{"aa", "bb"},
		},
	}, {
		name: "set union uint leaf with positive int value",
		inData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_IntVal{
					IntVal: 100,
				},
			},
			Timestamp: time.Unix(10, 10),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_UintVal{
					UintVal: 100,
				},
			},
			Timestamp: time.Unix(10, 10),
		}},
		wantStruct: &LeafContainerStruct{
			UnionLeaf: &UnionLeafType_Uint32{100},
		},
	}, {
		name:           "empty datapoint slice",
		inData:         []*DataPoint{},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantStruct:     &LeafContainerStruct{},
	}, {
		name:           "nil datapoint slice",
		inData:         nil,
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantStruct:     &LeafContainerStruct{},
	}, {
		name: "not all timestamps are the same -- the values should apply in order they are in the slice",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
			Timestamp: time.Unix(2, 1),
		}, {
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 200}},
			Timestamp: time.Unix(1, 0),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "aa"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "bb"}},
						},
					},
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "a"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "bbb"}},
						},
					},
				},
			},
			Timestamp: time.Unix(20, 20),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct"),
		inStructSchema: superContainerSchema,
		inStruct:       &SuperContainer{},
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 100}},
			Timestamp: time.Unix(2, 1),
		}, {
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 200}},
			Timestamp: time.Unix(1, 0),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "aa"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "bb"}},
						},
					},
				},
			},
			Timestamp: time.Unix(10, 10),
		}, {
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_LeaflistVal{
					LeaflistVal: &gpb.ScalarArray{
						Element: []*gpb.TypedValue{
							{Value: &gpb.TypedValue_StringVal{StringVal: "a"}},
							{Value: &gpb.TypedValue_StringVal{StringVal: "bbb"}},
						},
					},
				},
			},
			Timestamp: time.Unix(20, 20),
		}},
		wantStruct: &SuperContainer{
			LeafContainerStruct: &LeafContainerStruct{
				Uint64Leaf: ygot.Uint64(200),
				// TODO: If Collect calls are to be
				// implemented, then need to add tests for adds
				// and deletes to same and different children,
				// whether leaf or non-leaf, under a non-leaf.
				UnionLeafSingleType: []string{"a", "bbb"},
			},
		},
	}, {
		name: "retrieve single list key, setting the key value",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=42]/config/key"),
		inStructSchema: superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:       &Model_SingleKey{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model_SingleKey{Key: ygot.Int32(42)},
	}, {
		name: "retrieve single shadow-path list key, not setting the key value",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=42]/state/key"),
		inStructSchema: superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:       &Model_SingleKey{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model_SingleKey{},
	}, {
		name: "retrieve single shadow-path list key, setting the key value with preferShadowPath=true",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:        gnmiPath(t, "super-container/model/a/single-key[key=42]/state/key"),
		inStructSchema:     superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:           &Model_SingleKey{},
		inLeaf:             true,
		inPreferShadowPath: true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model_SingleKey{Key: ygot.Int32(42)},
	}, {
		name: "retrieve single non-shadow-path list value, setting the value",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
		inStructSchema: superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:       &Model_SingleKey{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model_SingleKey{Value: ygot.Int64(4242)},
	}, {
		name: "retrieve single shadow-path list value, not setting the value",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=42]/state/value"),
		inStructSchema: superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:       &Model_SingleKey{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model_SingleKey{},
	}, {
		name: "retrieve single non-shadow-path list value, not setting the value with preferShadowPath=true",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:        gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
		inStructSchema:     superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:           &Model_SingleKey{},
		inLeaf:             true,
		inPreferShadowPath: true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model_SingleKey{},
	}, {
		name: "retrieve single shadow-path list value, setting the value with preferShadowPath=true",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:        gnmiPath(t, "super-container/model/a/single-key[key=42]/state/value"),
		inStructSchema:     superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:           &Model_SingleKey{},
		inLeaf:             true,
		inPreferShadowPath: true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/state/value"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 4242}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model_SingleKey{Value: ygot.Int64(4242)},
	}, {
		name: "retrieve entire list, setting multiple list keys",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}, {
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=43]/config/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key"),
		inStructSchema: superContainerSchema.Dir["model"],
		inStruct:       &Model{},
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}, {
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=43]/config/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model{SingleKey: map[int32]*Model_SingleKey{
			42: {Key: ygot.Int32(42)},
			43: {Key: ygot.Int32(43)},
		}},
	}, {
		name: "delete a list key value, success",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
		inStructSchema: superContainerSchema.Dir["model"],
		inStruct: &Model{SingleKey: map[int32]*Model_SingleKey{
			42: {Key: ygot.Int32(42), Value: ygot.Int64(4242)},
		}},
		inLeaf: true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model{SingleKey: map[int32]*Model_SingleKey{
			42: {Key: ygot.Int32(42)},
		}},
	}, {
		name: "delete a list key, no-op since preferShadowPath=true",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
		inStructSchema: superContainerSchema.Dir["model"],
		inStruct: &Model{SingleKey: map[int32]*Model_SingleKey{
			42: {Key: ygot.Int32(42), Value: ygot.Int64(4242)},
		}},
		inLeaf:             true,
		inPreferShadowPath: true,
		wantUnmarshalledData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=42]/config/value"),
			Timestamp: time.Unix(4, 4),
		}},
		wantStruct: &Model{SingleKey: map[int32]*Model_SingleKey{
			42: {Key: ygot.Int32(42), Value: ygot.Int64(4242)},
		}},
	}}

	for _, tt := range passingTests {
		t.Run(tt.name, func(t *testing.T) {
			unmarshalledData, complianceErrs, err := unmarshal(tt.inData, tt.inStructSchema, tt.inStruct, tt.inQueryPath, schemaStruct(), tt.inLeaf, tt.inPreferShadowPath)
			if err != nil {
				t.Fatalf("unmarshal: got error, want none: %v", err)
			}
			if complianceErrs != nil {
				t.Fatalf("unmarshal: got compliance errors, want none: %v", complianceErrs)
			}

			if diff := cmp.Diff(tt.wantUnmarshalledData, unmarshalledData, protocmp.Transform()); diff != "" {
				t.Errorf("unmarshal: successfully unmarshalled datapoints do not match (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantStruct, tt.inStruct); diff != "" {
				t.Errorf("unmarshal: struct after unmarshalling does not match (-want +got):\n%s", diff)
			}
		})
	}

	failingTests := []struct {
		name                  string
		inData                []*DataPoint
		inQueryPath           *gpb.Path
		inStructSchema        *yang.Entry
		inStruct              ygot.GoStruct
		inLeaf                bool
		inPreferShadowPath    bool
		wantUnmarshalledData  []*DataPoint
		wantStruct            ygot.GoStruct
		wantErrSubstr         string
		wantPathErrSubstr     *TelemetryError
		wantTypeErrSubstr     *TelemetryError
		wantValidateErrSubstr string
	}{{
		name: "fail to retrieve uint64 due to wrong type",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantTypeErrSubstr: &TelemetryError{
			Path:  gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_StringVal{StringVal: "foo"}},
			Err:   errors.New("failed to unmarshal"),
		},
	}, {
		name: "multiple datapoints for leaf node",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 42}},
			Timestamp: time.Unix(1, 1),
		}, {
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_UintVal{UintVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Err: errors.New("got multiple"),
		},
	}, {
		name: "failed to retrieve uint64 with negative int",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: -42}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantTypeErrSubstr: &TelemetryError{
			Path:  gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: -42}},
			Err:   errors.New("failed to unmarshal"),
		},
	}, {
		name: "fail to retrieve uint64 due to wrong path",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/xxxxxxxxx/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Path:  gnmiPath(t, "super-container/xxxxxxxxx/uint64-leaf"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Err:   errors.New(`does not match the query path "/super-container/leaf-container-struct/uint64-leaf"`),
		},
	}, {
		name: "retrieve union with field that doesn't match regex",
		inData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "forty two",
				},
			},
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantUnmarshalledData: []*DataPoint{{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
			Value: &gpb.TypedValue{
				Value: &gpb.TypedValue_StringVal{
					StringVal: "forty two",
				},
			},
			Timestamp: time.Unix(2, 2),
		}},
		wantValidateErrSubstr: "does not match regular expression pattern",
	}, {
		name: "delete at a non-existent path",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/dne"),
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Path: gnmiPath(t, "super-container/leaf-container-struct/dne"),
			Err:  errors.New("does not match the query path"),
		},
	}, {
		name: "fail to delete union with a single enum inside due to wrong path prefix",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "not-valid-prefix/leaf-container-struct/union-leaf2"),
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf2"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{UnionLeaf2: EnumType(44)},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Path: gnmiPath(t, "not-valid-prefix/leaf-container-struct/union-leaf2"),
			Err:  errors.New(`does not match the query path "/super-container/leaf-container-struct/union-leaf2"`),
		},
	}, {
		name: "fail to delete union with a single enum inside due to wrong path suffix",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/leaf-container-struct/union-needle2"),
			Timestamp: time.Unix(2, 2),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/union-leaf2"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{UnionLeaf2: EnumType(44)},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Path: gnmiPath(t, "super-container/leaf-container-struct/union-needle2"),
			Err:  errors.New(`does not match the query path "/super-container/leaf-container-struct/union-leaf2"`),
		},
	}, {
		name: "retrieve a list value with an invalid list key",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "super-container/model/a/single-key[key=forty-four]/config/key"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Timestamp: time.Unix(4, 4),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=44]/config/key"),
		inStructSchema: superContainerSchema.Dir["model"].Dir["a"].Dir["single-key"],
		inStruct:       &Model_SingleKey{},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Path:  gnmiPath(t, "super-container/model/a/single-key[key=forty-four]/config/key"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 42}},
			Err:   errors.New(`does not match the query path "/super-container/model/a/single-key[key=44]/config/key"`),
		},
	}, {
		name: "invalid input: parent schema is not parent of input data's path.",
		inData: []*DataPoint{{
			Path:      gnmiPath(t, "different-container/leaf-container-struct/uint64-leaf"),
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Path:  gnmiPath(t, "different-container/leaf-container-struct/uint64-leaf"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Err:   errors.New(`does not match the query path "/super-container/leaf-container-struct/uint64-leaf"`),
		},
	}, {
		name: "invalid input: deprecated elements field in path",
		inData: []*DataPoint{{
			Path: &gpb.Path{
				Element: []string{"super-container", "model", "a", "single-key", "forty-four", "config", "key"},
			},
			Value:     &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Timestamp: time.Unix(1, 1),
		}},
		inQueryPath:    gnmiPath(t, "super-container/model/a/single-key[key=44]/config/key"),
		inStructSchema: superContainerSchema.Dir["leaf-container-struct"],
		inStruct:       &LeafContainerStruct{},
		inLeaf:         true,
		wantPathErrSubstr: &TelemetryError{
			Path: &gpb.Path{
				Element: []string{"super-container", "model", "a", "single-key", "forty-four", "config", "key"},
			},
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 43}},
			Err:   errors.New(`path uses deprecated and unsupported Element field`),
		},
	}}

	for _, tt := range failingTests {
		t.Run(tt.name, func(t *testing.T) {
			unmarshalledData, complianceErrs, errs := unmarshal(tt.inData, tt.inStructSchema, tt.inStruct, tt.inQueryPath, schemaStruct(), tt.inLeaf, tt.inPreferShadowPath)
			if errs != nil {
				t.Fatalf("unmarshal: got more than one error: %v", errs)
			}
			if diff := cmp.Diff(tt.wantUnmarshalledData, unmarshalledData, protocmp.Transform()); diff != "" {
				t.Errorf("unmarshal: successfully unmarshalled datapoints do not match (-want +got):\n%s", diff)
			}

			var pathErrs, typeErrs []*TelemetryError
			var validateErrs []error
			if complianceErrs != nil {
				pathErrs = complianceErrs.PathErrors
				typeErrs = complianceErrs.TypeErrors
				validateErrs = complianceErrs.ValidateErrors
			}
			if len(pathErrs) > 1 {
				t.Fatalf("unmarshal: got more than one path unmarshal error: %v", pathErrs)
			}
			if len(typeErrs) > 1 {
				t.Fatalf("unmarshal: got more than one type unmarshal error: %v", typeErrs)
			}
			if len(validateErrs) > 1 {
				t.Fatalf("unmarshal: got more than one validate error: %v", validateErrs)
			}

			// Populate errors for validation.
			var err, validateErr error
			var pathErr, typeErr *TelemetryError
			if len(pathErrs) == 1 {
				pathErr = pathErrs[0]
			}
			if len(typeErrs) == 1 {
				typeErr = typeErrs[0]
			}
			if len(validateErrs) == 1 {
				validateErr = validateErrs[0]
			}

			// Validate expected errors
			if diff := errdiff.Substring(err, tt.wantErrSubstr); diff != "" {
				t.Fatalf("unmarshal: did not get expected error substring:\n%s", diff)
			}

			verifyTelemetryError := func(t *testing.T, gotErr, wantErrSubstr *TelemetryError) {
				t.Helper()
				// Only do exact verification on the Path and Value fields of the Telemetry errors.
				if diff := cmp.Diff(wantErrSubstr, gotErr, protocmp.Transform(), cmp.FilterPath(
					func(p cmp.Path) bool {
						return p.String() == "Err"
					},
					cmp.Ignore(),
				)); diff != "" {
					t.Fatalf("unmarshal: did not get expected path compliance error (-want, +got):\n%s", diff)
				}
				if gotErr != nil && wantErrSubstr != nil {
					if diff := errdiff.Substring(gotErr.Err, wantErrSubstr.Err.Error()); diff != "" {
						t.Fatalf("unmarshal: did not get expected compliance error substring:\n%s", diff)
					}
				}
			}
			verifyTelemetryError(t, pathErr, tt.wantPathErrSubstr)
			verifyTelemetryError(t, typeErr, tt.wantTypeErrSubstr)

			if diff := errdiff.Substring(validateErr, tt.wantValidateErrSubstr); diff != "" {
				t.Fatalf("unmarshal: did not get expected validateErr substring:\n%s", diff)
			}
		})
	}
}

func TestLatestTimestamp(t *testing.T) {
	tests := []struct {
		desc     string
		in       []*DataPoint
		want     time.Time
		wantRecv time.Time
	}{{
		desc: "basic",
		in: []*DataPoint{{
			Path:          gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:         &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
			Timestamp:     time.Unix(0, 1),
			RecvTimestamp: time.Unix(5, 5),
		}, {
			Path:          gnmiPath(t, "super-container/leaf-container-struct/uint64-leaf"),
			Value:         &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 200}},
			Timestamp:     time.Unix(3, 3),
			RecvTimestamp: time.Unix(4, 4),
		}, {
			Path:          gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value:         &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 300}},
			Timestamp:     time.Unix(2, 2),
			RecvTimestamp: time.Unix(3, 3),
		}, {
			Path:          gnmiPath(t, "super-container/leaf-container-struct/union-stleaflist"),
			Value:         &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 400}},
			Timestamp:     time.Unix(1, 1),
			RecvTimestamp: time.Unix(2, 2),
		}},
		want:     time.Unix(3, 3),
		wantRecv: time.Unix(5, 5),
	}, {
		desc:     "empty list",
		want:     time.Time{},
		wantRecv: time.Time{},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := LatestTimestamp(tt.in); !got.Equal(tt.want) {
				t.Errorf("LatestTimestamp: got %v, want %v", got, tt.want)
			}
			if got := LatestRecvTimestamp(tt.in); !got.Equal(tt.wantRecv) {
				t.Errorf("LatestRecvTimestamp: got %v, want %v", got, tt.wantRecv)
			}
		})
	}
}

func TestBundleDatapoints(t *testing.T) {
	tests := []struct {
		desc         string
		inDatapoints []*DataPoint
		inPrefixLen  uint
		want         map[string][]*DataPoint
		wantPrefixes []string
		wantErr      bool
	}{{
		desc: "leaf-paths",
		inDatapoints: []*DataPoint{{
			Path:  gnmiPath(t, "alpha[key=un]/bravo/leaf0"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
		}, {
			Path:  gnmiPath(t, "alpha[key=deux]/bravo/leaf0"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 200}},
		}, {
			Path:  gnmiPath(t, "alpha[key=un]/bravo/leaf0/leaf-under-leaf"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 300}},
		}},
		inPrefixLen: 3,
		want: map[string][]*DataPoint{
			"/alpha[key=un]/bravo/leaf0": {{
				Path:  gnmiPath(t, "alpha[key=un]/bravo/leaf0"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
			}, {
				Path:  gnmiPath(t, "alpha[key=un]/bravo/leaf0/leaf-under-leaf"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 300}},
			}},
			"/alpha[key=deux]/bravo/leaf0": {{
				Path:  gnmiPath(t, "alpha[key=deux]/bravo/leaf0"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 200}},
			}},
		},
		wantPrefixes: []string{
			"/alpha[key=deux]/bravo/leaf0",
			"/alpha[key=un]/bravo/leaf0",
		},
	}, {
		desc: "non-leaf-paths",
		inDatapoints: []*DataPoint{{
			Path:  gnmiPath(t, "alpha/bravo[key=un]/leaf0"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=deux]/leaf2"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 300}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf3"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 400}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf4"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 500}},
		}, {
			// duplicate path
			Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf3"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 1000}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=deux]/foo/leaf5"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 600}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=un]/foo/bar/leaf6"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 700}},
		}},
		inPrefixLen: 2,
		want: map[string][]*DataPoint{
			"/alpha/bravo[key=un]": {{
				Path:  gnmiPath(t, "alpha/bravo[key=un]/leaf0"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
			}, {
				Path:  gnmiPath(t, "alpha/bravo[key=un]/foo/bar/leaf6"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 700}},
			}},
			"/alpha/bravo[key=deux]": {{
				Path:  gnmiPath(t, "alpha/bravo[key=deux]/leaf2"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 300}},
			}, {
				Path:  gnmiPath(t, "alpha/bravo[key=deux]/foo/leaf5"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 600}},
			}},
			"/alpha/bravo[key=trois]": {{
				Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf3"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 400}},
			}, {
				Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf4"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 500}},
			}, {
				Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf3"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 1000}},
			}},
		},
		wantPrefixes: []string{
			"/alpha/bravo[key=deux]",
			"/alpha/bravo[key=trois]",
			"/alpha/bravo[key=un]",
		},
	}, {
		desc: "path-shorter-than-prefixLen",
		inDatapoints: []*DataPoint{{
			Path:  gnmiPath(t, "alpha/"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 10}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=un]/leaf0"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=deux]/leaf2"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 300}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf3"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 400}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf4"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 500}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=deux]/foo/leaf5"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 600}},
		}, {
			Path:  gnmiPath(t, "alpha/bravo[key=un]/foo/bar/leaf6"),
			Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 700}},
		}},
		inPrefixLen: 2,
		want: map[string][]*DataPoint{
			"/": {{
				Path:  gnmiPath(t, "alpha/"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 10}},
			}},
			"/alpha/bravo[key=un]": {{
				Path:  gnmiPath(t, "alpha/bravo[key=un]/leaf0"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 100}},
			}, {
				Path:  gnmiPath(t, "alpha/bravo[key=un]/foo/bar/leaf6"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 700}},
			}},
			"/alpha/bravo[key=deux]": {{
				Path:  gnmiPath(t, "alpha/bravo[key=deux]/leaf2"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 300}},
			}, {
				Path:  gnmiPath(t, "alpha/bravo[key=deux]/foo/leaf5"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 600}},
			}},
			"/alpha/bravo[key=trois]": {{
				Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf3"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 400}},
			}, {
				Path:  gnmiPath(t, "alpha/bravo[key=trois]/leaf4"),
				Value: &gpb.TypedValue{Value: &gpb.TypedValue_IntVal{IntVal: 500}},
			}},
		},
		wantPrefixes: []string{
			"/",
			"/alpha/bravo[key=deux]",
			"/alpha/bravo[key=trois]",
			"/alpha/bravo[key=un]",
		},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, gotPrefixes, err := bundleDatapoints(tt.inDatapoints, tt.inPrefixLen)
			if gotErr := err != nil; gotErr != tt.wantErr {
				t.Errorf("Got error: %v, want error: %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want, protocmp.Transform()); diff != "" {
				t.Errorf("Datapoint groups (-got, +want):\n%s", diff)
			}
			if diff := cmp.Diff(gotPrefixes, tt.wantPrefixes); diff != "" {
				t.Errorf("Prefixes (-got, +want):\n%s", diff)
			}
		})
	}
}

func TestQualifiedTypeString(t *testing.T) {
	tests := []struct {
		desc  string
		want  string
		value interface{}
		time  time.Time
		path  *gpb.Path
	}{{
		desc:  "nil interface",
		value: nil,
		time:  time.Time{},
		path: &gpb.Path{
			Origin: "openconfig",
			Elem: []*gpb.PathElem{{
				Name: "network-instances",
			}, {
				Name: "network-instance",
				Key:  map[string]string{"name": "master"},
			}}},
		want: "\nTimestamp: 0001-01-01 00:00:00 +0000 UTC\ntarget: , path: /network-instances/network-instance[name=master]\nVal: <nil> (not present)\n\n",
	}, {
		desc:  "typed nil",
		value: (*LeafContainerStruct)(nil),
		time:  time.Time{},
		path: &gpb.Path{
			Origin: "openconfig",
			Elem: []*gpb.PathElem{{
				Name: "network-instances",
			}, {
				Name: "network-instance",
				Key:  map[string]string{"name": "master"},
			}}},
		want: "\nTimestamp: 0001-01-01 00:00:00 +0000 UTC\ntarget: , path: /network-instances/network-instance[name=master]\nVal: <nil> (not present)\n\n",
	}, {
		desc:  "typed nil and GoStruct",
		value: (*Device)(nil),
		time:  time.Time{},
		path: &gpb.Path{
			Origin: "openconfig",
			Elem: []*gpb.PathElem{{
				Name: "network-instances",
			}, {
				Name: "network-instance",
				Key:  map[string]string{"name": "master"},
			}}},
		want: "\nTimestamp: 0001-01-01 00:00:00 +0000 UTC\ntarget: , path: /network-instances/network-instance[name=master]\nVal: <nil> (not present)\n\n",
	}, {
		desc:  "non-nil validated struct",
		value: &Device{},
		time:  time.Time{},
		want:  "\nTimestamp: 0001-01-01 00:00:00 +0000 UTC\n\nVal: \n{}\n\n",
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := QualifiedTypeString(tt.value, &Metadata{
				Timestamp: tt.time,
				Path:      tt.path,
			})
			if got != tt.want {
				t.Fatalf("Got string %q, want %q", got, tt.want)
			}
		})
	}
}
