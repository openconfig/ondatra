// Copyright 2019 Google Inc.
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

package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ondatra/internal/gnmigen"
)

func TestWriteGoCode(t *testing.T) {
	telemCode := &gnmigen.GeneratedTelemetryCode{
		Headers: map[string]string{
			"telemetry": "telem common header\n",
		},
		FakeRootMethods: "\ntelem root methods\n",
		GoPerNodeSnippets: map[string]gnmigen.GoPerNodeCodeSnippets{
			"telemetry": {{
				PathStructName: "PathStruct1",
				GetMethod:      "\nGet1\n",
				CollectMethod:  "\nCollect1\n",
				ConvertHelper:  "\nconvert1\n",
			}, {
				PathStructName: "PathStruct2",
				GetMethod:      "\nGet2\n",
				CollectMethod:  "\nCollect2\n",
				ConvertHelper:  "\nconvert2\n",
			}, {
				PathStructName: "PathStruct3",
				GetMethod:      "\nGet3\n",
				CollectMethod:  "\nCollect3\n",
				ConvertHelper:  "\nconvert3\n",
			}}},
		GoReturnTypeSnippets: []gnmigen.GoReturnTypeCodeSnippet{{
			TypeName:       "int64",
			QualifiedType:  "\nQualifiedInt64\n",
			CollectionType: "\nCollectionInt64\n",
		}, {
			TypeName:       "float64",
			QualifiedType:  "\nQualifiedFloat64\n",
			CollectionType: "\nCollectionFloat64\n",
		}},
	}

	telemFuncsFileName := func(filen int) string { return fmt.Sprintf(telemFuncsFileFmt, filen) }
	telemTypesFileName := func(filen int) string { return fmt.Sprintf(telemTypesFileFmt, filen) }

	tests := []struct {
		name              string
		inTelemFuncsFileN int
		inTelemTypesFileN int
		want              map[string]string
	}{{
		name:              "no outputs",
		inTelemFuncsFileN: 0,
		inTelemTypesFileN: 0,
		want: map[string]string{
			fakeRootMethodsFileName: "telem common header\n\ntelem root methods\n",
		},
	}, {
		name:              "no file splitting",
		inTelemFuncsFileN: 1,
		inTelemTypesFileN: 1,
		want: map[string]string{
			fakeRootMethodsFileName: "telem common header\n\ntelem root methods\n",
			telemTypesFileName(0): `telem common header

QualifiedInt64

CollectionInt64

QualifiedFloat64

CollectionFloat64
`,
			telemFuncsFileName(0): `telem common header

Get1

Collect1

convert1

Get2

Collect2

convert2

Get3

Collect3

convert3
`,
		},
	}, {
		name:              "split into two files",
		inTelemFuncsFileN: 2,
		inTelemTypesFileN: 2,
		want: map[string]string{
			fakeRootMethodsFileName: "telem common header\n\ntelem root methods\n",
			telemTypesFileName(0): `telem common header

QualifiedInt64

CollectionInt64
`,
			telemTypesFileName(1): `telem common header

QualifiedFloat64

CollectionFloat64
`,
			telemFuncsFileName(0): `telem common header

Get1

Collect1

convert1

Get2

Collect2

convert2
`,
			telemFuncsFileName(1): `telem common header

Get3

Collect3

convert3
`,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeFileOutputSpec(telemCode, false, true, tt.inTelemFuncsFileN, tt.inTelemTypesFileN)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Did not get expected output, diff (-want,+got):\n%s", diff)
			}
		})
	}
}

func TestWriteGoCodeSplitPackage(t *testing.T) {
	telemCode := &gnmigen.GeneratedTelemetryCode{
		Headers: map[string]string{
			"telemetry": "telem common header\n",
			"foo_path":  "foo_path common header\n",
		},
		FakeRootMethods: "telem fake root helper\n",
		GoPerNodeSnippets: map[string]gnmigen.GoPerNodeCodeSnippets{
			"telemetry": {{
				PathStructName: "PathStruct1",
				GetMethod:      "\nGet1\n",
				CollectMethod:  "\nCollect1\n",
				ConvertHelper:  "\nconvert1\n",
			}, {
				PathStructName: "PathStruct2",
				GetMethod:      "\nGet2\n",
				CollectMethod:  "\nCollect2\n",
				ConvertHelper:  "\nconvert2\n",
			}},
			"foo_path": {{
				PathStructName: "PathStruct3",
				GetMethod:      "\nGet3\n",
				CollectMethod:  "\nCollect3\n",
				ConvertHelper:  "\nconvert3\n",
			}}},
		GoReturnTypeSnippets: []gnmigen.GoReturnTypeCodeSnippet{{
			TypeName:       "int64",
			QualifiedType:  "\nQualifiedInt64\n",
			CollectionType: "\nCollectionInt64\n",
		}, {
			TypeName:       "float64",
			QualifiedType:  "\nQualifiedFloat64\n",
			CollectionType: "\nCollectionFloat64\n",
		}},
	}

	tests := []struct {
		name                 string
		inTelemFuncsFileN    int
		inTelemTypesFileN    int
		inGeneratePathHelper bool
		want                 map[string]string
	}{{
		name:              "no outputs",
		inTelemFuncsFileN: 0,
		inTelemTypesFileN: 0,
		want:              map[string]string{},
	}, {
		name:                 "generate path structs",
		inTelemFuncsFileN:    1,
		inTelemTypesFileN:    1,
		inGeneratePathHelper: true,
		want: map[string]string{
			fakeRootMethodsFileName: "telem common header\ntelem fake root helper\n",
			"foo_path/telemetry-0.go": `foo_path common header

Get3

Collect3

convert3
`,
			"device_telem.go": `telem common header

Get1

Collect1

convert1

Get2

Collect2

convert2
`,
			"telem_types-0.go": `telem common header

QualifiedInt64

CollectionInt64

QualifiedFloat64

CollectionFloat64
`,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := makeFileOutputSpec(telemCode, true, tt.inGeneratePathHelper, tt.inTelemFuncsFileN, tt.inTelemTypesFileN)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Did not get expected output, diff (-want,+got):\n%s", diff)
			}
		})
	}
}
