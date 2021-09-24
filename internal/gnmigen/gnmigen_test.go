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

package gnmigen

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ygot/testutil"
	"github.com/openconfig/ygot/ypathgen"
)

func generateCodeTestHelper(t *testing.T, dirName string, generatedFileName string, config bool) {
	t.Helper()
	const (
		datapath = "testdata"
		// deflakeRuns specifies the number of runs of code generation that
		// should be performed to check for flakes.
		deflakeRuns = 10
	)

	tests := []struct {
		// name is the identifier for the test.
		name string
		// resultsFolder specifies the folder containing the result files.
		resultsFolder string
		// inFiles is the set of inputFiles for the test.
		inFiles []string
		// inIncludePaths is the set of paths that should be searched for imports.
		inIncludePaths []string
		// inConfig is the configuration used for telemetry method generation.
		inConfig *GenConfig
		// inListBuilderKeyThreshold is the threshold equal or over which the
		// builder API is used for key population.
		inListBuilderKeyThreshold uint
		// wantErr specifies whether the test should expect an error.
		wantErr bool
	}{{
		name:          "container_test",
		resultsFolder: "container_test",
		inFiles:       []string{"openconfig-simple.yang"},
		inConfig:      NewDefaultConfig(),
	}, {
		name:          "list_test",
		resultsFolder: "list_test",
		inFiles:       []string{"openconfig-withlist.yang"},
		inConfig:      NewDefaultConfig(),
	}, {
		name:                      "list_test with builder API",
		resultsFolder:             "list_test",
		inFiles:                   []string{"openconfig-withlist.yang"},
		inConfig:                  NewDefaultConfig(),
		inListBuilderKeyThreshold: 2,
	}}

	codeDiff := func(t *testing.T, want, got string) string {
		// Use difflib to generate a unified diff between the
		// two code snippets such that this is simpler to debug
		// in the test output.
		diff, err := testutil.GenerateUnifiedDiff(want, got)
		if err != nil {
			t.Fatalf("error generating diff: %v", err)
		}
		return diff
	}

	for _, tt := range tests {
		// Set up
		tt.inConfig.GeneratingBinary = "gnmigen-tests"
		tt.inConfig.ListBuilderKeyThreshold = tt.inListBuilderKeyThreshold
		for i := range tt.inFiles {
			tt.inFiles[i] = filepath.Join(datapath, tt.resultsFolder, tt.inFiles[i])
		}

		genCode := tt.inConfig.GenerateTelemetryCode
		if config {
			genCode = tt.inConfig.GenerateConfigCode
		}

		t.Run(tt.name, func(t *testing.T) {
			code, errs := genCode(tt.inFiles, tt.inIncludePaths)
			if gotErr := errs != nil; gotErr != tt.wantErr {
				t.Fatalf("got errors: %v, wantErr: %v", errs, tt.wantErr)
			} else if gotErr {
				// Received expected error.
				return
			}

			wantFile := filepath.Join(datapath, tt.resultsFolder, dirName, generatedFileName)
			got := code.String()
			diffDetected := false
			wantCodeBytes, err := ioutil.ReadFile(wantFile)
			if err != nil {
				t.Fatalf("ioutil.ReadFile(%q) error: %v", wantFile, err)
			}
			if diff := codeDiff(t, string(wantCodeBytes), got); diff != "" {
				diffDetected = true
				t.Errorf("GenerateTelemetryCode(%v, %v), Config: %v, did not return correct code (file: %v), (-want, +got):\n%s",
					tt.inFiles, tt.inIncludePaths, tt.inConfig, wantFile, diff)
			}
			if diffDetected {
				// Don't check flakes unless diff succeeds.
				return
			}

			for i := 0; i < deflakeRuns; i++ {
				gotAttempt, _ := genCode(tt.inFiles, tt.inIncludePaths)
				if diff := codeDiff(t, got, gotAttempt.String()); diff != "" {
					t.Fatalf("flaky code generation, (-want, +got):\n%s", diff)
				}
			}
		})
	}
}

func TestGenerateTelemetryCode(t *testing.T) {
	generateCodeTestHelper(t, "telemetry", "telemetry.txt", false)
}

func TestGenerateConfigCode(t *testing.T) {
	generateCodeTestHelper(t, "config", "config.txt", true)
}

func TestGeneratePerNodeConfigSnippet(t *testing.T) {
	tests := []struct {
		name                      string
		inPathStructName          string
		inNodeData                *ypathgen.NodeData
		inFakeRootTypeName        string
		inSchemaStructPkgAccessor string
		wantSnippet               GoPerNodeCodeSnippet
		wantGoTypeData            goTypeData
		wantErr                   bool
	}{{
		name:             "scalar leaf",
		inPathStructName: "Container_Leaf",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "int32",
			LocalGoTypeName:       "int32",
			SubsumingGoStructName: "Container",
			GoFieldName:           "Leaf",
			IsLeaf:                true,
			IsScalarField:         true,
			HasDefault:            true,
			YANGPath:              "/container/leaf",
		},
		inFakeRootTypeName: "Device",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "Container_Leaf",
			GetMethod: `
// GetFull retrieves a sample for /container/leaf.
func (n *Container_Leaf) GetFull(t testing.TB) *QualifiedInt32 {
	t.Helper()
	goStruct := &Container{}
	ret := &QualifiedInt32{
		QualifiedType: getFull(t, n, "Container", goStruct, true),
	}
	if !ret.IsPresent() {
		ret.SetVal(goStruct.GetLeaf())
		ret.Present = true
		return ret
	}
	return convertContainer_Leaf(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /container/leaf, erroring out if it is not present.
func (n *Container_Leaf) Get(t testing.TB) int32 {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /container/leaf.
func (n *Container_LeafAny) GetFull(t testing.TB) []*QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &Container{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "Container", goStruct, queryPath, true, false)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertContainer_Leaf(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /container/leaf.
func (n *Container_LeafAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent Container
// and combines the update with an existing QualifiedType to return a *QualifiedInt32.
func convertContainer_Leaf(t testing.TB, qt *genutil.QualifiedType, parent *Container) *QualifiedInt32 {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &QualifiedInt32{
		QualifiedType: qt,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(*val)
	} else {
		qv.Present = false
	}
	return qv
}
`,

			ReplaceMethod: `
// Delete deletes the configuration at /container/leaf.
func (n *Container_Leaf) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchDelete(t testing.TB, b *SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /container/leaf.
func (n *Container_Leaf) Replace(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchReplace(t testing.TB, b *SetRequestBatch, val int32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /container/leaf.
func (n *Container_Leaf) Update(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchUpdate(t testing.TB, b *SetRequestBatch, val int32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "int32",
			TransformedGoTypeName: "Int32",
			IsLeaf:                true,
			HasDefault:            true,
		},
	}, {
		name:             "scalar leaf with schemaStructPkgAccessor",
		inPathStructName: "Container_Leaf",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "int32",
			LocalGoTypeName:       "int32",
			SubsumingGoStructName: "Container",
			GoFieldName:           "Leaf",
			IsLeaf:                true,
			IsScalarField:         true,
			YANGPath:              "/container/leaf",
		},
		inFakeRootTypeName:        "oc.Device",
		inSchemaStructPkgAccessor: "oc.",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "Container_Leaf",
			GetMethod: `
// GetFull retrieves a sample for /container/leaf.
func (n *Container_Leaf) GetFull(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.Container{}
	ret := &oc.QualifiedInt32{
		QualifiedType: getFull(t, n, "Container", goStruct, true),
	}
	return convertContainer_Leaf(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /container/leaf, erroring out if it is not present.
func (n *Container_Leaf) Get(t testing.TB) int32 {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /container/leaf.
func (n *Container_LeafAny) GetFull(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Container{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "Container", goStruct, queryPath, true, false)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertContainer_Leaf(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /container/leaf.
func (n *Container_LeafAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent oc.Container
// and combines the update with an existing QualifiedType to return a *oc.QualifiedInt32.
func convertContainer_Leaf(t testing.TB, qt *genutil.QualifiedType, parent *oc.Container) *oc.QualifiedInt32 {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &oc.QualifiedInt32{
		QualifiedType: qt,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(*val)
	} else {
		qv.Present = false
	}
	return qv
}
`,

			ReplaceMethod: `
// Delete deletes the configuration at /container/leaf.
func (n *Container_Leaf) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchDelete(t testing.TB, b *SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /container/leaf.
func (n *Container_Leaf) Replace(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchReplace(t testing.TB, b *SetRequestBatch, val int32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /container/leaf.
func (n *Container_Leaf) Update(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchUpdate(t testing.TB, b *SetRequestBatch, val int32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "int32",
			TransformedGoTypeName: "Int32",
			IsLeaf:                true,
		},
	}, {
		name:             "non-scalar leaf",
		inPathStructName: "List_Key",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "[]Binary",
			LocalGoTypeName:       "[]Binary",
			GoFieldName:           "Key",
			SubsumingGoStructName: "List",
			IsLeaf:                true,
			IsScalarField:         false,
			YANGPath:              "/lists/list/key",
		},
		inFakeRootTypeName: "FakeRoot",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "List_Key",
			GetMethod: `
// GetFull retrieves a sample for /lists/list/key.
func (n *List_Key) GetFull(t testing.TB) *QualifiedBinarySlice {
	t.Helper()
	goStruct := &List{}
	ret := &QualifiedBinarySlice{
		QualifiedType: getFull(t, n, "List", goStruct, true),
	}
	return convertList_Key(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /lists/list/key, erroring out if it is not present.
func (n *List_Key) Get(t testing.TB) []Binary {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /lists/list/key.
func (n *List_KeyAny) GetFull(t testing.TB) []*QualifiedBinarySlice {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*QualifiedBinarySlice
	for _, prefix := range sortedPrefixes {
		goStruct := &List{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "List", goStruct, queryPath, true, false)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertList_Key(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /lists/list/key.
func (n *List_KeyAny) Get(t testing.TB) [][]Binary {
	t.Helper()
	fulldata := n.GetFull(t)
	var data [][]Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent List
// and combines the update with an existing QualifiedType to return a *QualifiedBinarySlice.
func convertList_Key(t testing.TB, qt *genutil.QualifiedType, parent *List) *QualifiedBinarySlice {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &QualifiedBinarySlice{
		QualifiedType: qt,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(val)
	} else {
		qv.Present = false
	}
	return qv
}
`,
			ReplaceMethod: `
// Delete deletes the configuration at /lists/list/key.
func (n *List_Key) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchDelete(t testing.TB, b *SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /lists/list/key.
func (n *List_Key) Replace(t testing.TB, val []Binary) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchReplace(t testing.TB, b *SetRequestBatch, val []Binary) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /lists/list/key.
func (n *List_Key) Update(t testing.TB, val []Binary) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchUpdate(t testing.TB, b *SetRequestBatch, val []Binary) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "[]Binary",
			TransformedGoTypeName: "BinarySlice",
			IsLeaf:                true,
		},
	}, {
		name:             "non-scalar leaf with schemaStructPkgAccessor",
		inPathStructName: "List_Key",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "[]oc.Binary",
			LocalGoTypeName:       "[]Binary",
			GoFieldName:           "Key",
			SubsumingGoStructName: "List",
			IsLeaf:                true,
			IsScalarField:         false,
			YANGPath:              "/lists/list/key",
		},
		inFakeRootTypeName:        "oc.Device",
		inSchemaStructPkgAccessor: "oc.",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "List_Key",
			GetMethod: `
// GetFull retrieves a sample for /lists/list/key.
func (n *List_Key) GetFull(t testing.TB) *oc.QualifiedBinarySlice {
	t.Helper()
	goStruct := &oc.List{}
	ret := &oc.QualifiedBinarySlice{
		QualifiedType: getFull(t, n, "List", goStruct, true),
	}
	return convertList_Key(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /lists/list/key, erroring out if it is not present.
func (n *List_Key) Get(t testing.TB) []oc.Binary {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /lists/list/key.
func (n *List_KeyAny) GetFull(t testing.TB) []*oc.QualifiedBinarySlice {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*oc.QualifiedBinarySlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.List{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "List", goStruct, queryPath, true, false)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertList_Key(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /lists/list/key.
func (n *List_KeyAny) Get(t testing.TB) [][]oc.Binary {
	t.Helper()
	fulldata := n.GetFull(t)
	var data [][]oc.Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent oc.List
// and combines the update with an existing QualifiedType to return a *oc.QualifiedBinarySlice.
func convertList_Key(t testing.TB, qt *genutil.QualifiedType, parent *oc.List) *oc.QualifiedBinarySlice {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &oc.QualifiedBinarySlice{
		QualifiedType: qt,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(val)
	} else {
		qv.Present = false
	}
	return qv
}
`,
			ReplaceMethod: `
// Delete deletes the configuration at /lists/list/key.
func (n *List_Key) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchDelete(t testing.TB, b *SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /lists/list/key.
func (n *List_Key) Replace(t testing.TB, val []oc.Binary) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchReplace(t testing.TB, b *SetRequestBatch, val []oc.Binary) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /lists/list/key.
func (n *List_Key) Update(t testing.TB, val []oc.Binary) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchUpdate(t testing.TB, b *SetRequestBatch, val []oc.Binary) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "[]oc.Binary",
			TransformedGoTypeName: "BinarySlice",
			IsLeaf:                true,
		},
	}, {
		name:             "non-leaf",
		inPathStructName: "SuperContainer_Container",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "*SuperContainer_Container",
			LocalGoTypeName:       "*SuperContainer_Container",
			GoFieldName:           "Container",
			SubsumingGoStructName: "SuperContainer_Container",
			IsLeaf:                false,
			IsScalarField:         false,
			YANGPath:              "/super-container/container",
		},
		inFakeRootTypeName: "Device",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "SuperContainer_Container",
			GetMethod: `
// GetFull retrieves a sample for /super-container/container.
func (n *SuperContainer_Container) GetFull(t testing.TB) *QualifiedSuperContainer_Container {
	t.Helper()
	goStruct := &SuperContainer_Container{}
	ret := &QualifiedSuperContainer_Container{
		QualifiedType: getFull(t, n, "SuperContainer_Container", goStruct, false),
	}
	if ret.IsPresent() {
		ret.SetVal(goStruct)
	}
	return ret
}

// Get retrieves a value sample for /super-container/container, erroring out if it is not present.
func (n *SuperContainer_Container) Get(t testing.TB) *SuperContainer_Container {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /super-container/container.
func (n *SuperContainer_ContainerAny) GetFull(t testing.TB) []*QualifiedSuperContainer_Container {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), false)

	var data []*QualifiedSuperContainer_Container
	for _, prefix := range sortedPrefixes {
		goStruct := &SuperContainer_Container{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "SuperContainer_Container", goStruct, queryPath, false, false)
		if err != nil {
			t.Fatal(err)
		}
		if !qt.IsPresent() {
			continue
		}
		qv := (&QualifiedSuperContainer_Container{
			QualifiedType: qt,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /super-container/container.
func (n *SuperContainer_ContainerAny) Get(t testing.TB) []*SuperContainer_Container {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []*SuperContainer_Container
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ReplaceMethod: `
// Delete deletes the configuration at /super-container/container.
func (n *SuperContainer_Container) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchDelete(t testing.TB, b *SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /super-container/container.
func (n *SuperContainer_Container) Replace(t testing.TB, val *SuperContainer_Container) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchReplace(t testing.TB, b *SetRequestBatch, val *SuperContainer_Container) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /super-container/container.
func (n *SuperContainer_Container) Update(t testing.TB, val *SuperContainer_Container) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchUpdate(t testing.TB, b *SetRequestBatch, val *SuperContainer_Container) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "*SuperContainer_Container",
			TransformedGoTypeName: "SuperContainer_Container",
			IsLeaf:                false,
		},
	}, {
		name:             "non-leaf with schemaStructPkgAccessor",
		inPathStructName: "SuperContainer_Container",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "*oc.SuperContainer_Container",
			LocalGoTypeName:       "*SuperContainer_Container",
			GoFieldName:           "Container",
			SubsumingGoStructName: "SuperContainer_Container",
			IsLeaf:                false,
			IsScalarField:         false,
			YANGPath:              "/super-container/container",
		},
		inFakeRootTypeName:        "oc.Device",
		inSchemaStructPkgAccessor: "oc.",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "SuperContainer_Container",
			GetMethod: `
// GetFull retrieves a sample for /super-container/container.
func (n *SuperContainer_Container) GetFull(t testing.TB) *oc.QualifiedSuperContainer_Container {
	t.Helper()
	goStruct := &oc.SuperContainer_Container{}
	ret := &oc.QualifiedSuperContainer_Container{
		QualifiedType: getFull(t, n, "SuperContainer_Container", goStruct, false),
	}
	if ret.IsPresent() {
		ret.SetVal(goStruct)
	}
	return ret
}

// Get retrieves a value sample for /super-container/container, erroring out if it is not present.
func (n *SuperContainer_Container) Get(t testing.TB) *oc.SuperContainer_Container {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /super-container/container.
func (n *SuperContainer_ContainerAny) GetFull(t testing.TB) []*oc.QualifiedSuperContainer_Container {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), false)

	var data []*oc.QualifiedSuperContainer_Container
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.SuperContainer_Container{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "SuperContainer_Container", goStruct, queryPath, false, false)
		if err != nil {
			t.Fatal(err)
		}
		if !qt.IsPresent() {
			continue
		}
		qv := (&oc.QualifiedSuperContainer_Container{
			QualifiedType: qt,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /super-container/container.
func (n *SuperContainer_ContainerAny) Get(t testing.TB) []*oc.SuperContainer_Container {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []*oc.SuperContainer_Container
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ReplaceMethod: `
// Delete deletes the configuration at /super-container/container.
func (n *SuperContainer_Container) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchDelete(t testing.TB, b *SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /super-container/container.
func (n *SuperContainer_Container) Replace(t testing.TB, val *oc.SuperContainer_Container) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchReplace(t testing.TB, b *SetRequestBatch, val *oc.SuperContainer_Container) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /super-container/container.
func (n *SuperContainer_Container) Update(t testing.TB, val *oc.SuperContainer_Container) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchUpdate(t testing.TB, b *SetRequestBatch, val *oc.SuperContainer_Container) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "*oc.SuperContainer_Container",
			TransformedGoTypeName: "SuperContainer_Container",
			IsLeaf:                false,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSnippet, gotGoTypeData, errs := generatePerNodeConfigSnippet(tt.inPathStructName, tt.inNodeData, tt.inFakeRootTypeName, tt.inSchemaStructPkgAccessor, false)
			if tt.wantErr {
				if errs == nil {
					t.Fatalf("want error but did not get any.\ngotSnippet: %v\n\ngotGoTypeData: %v", gotSnippet, gotGoTypeData)
				}
				t.Logf("success: expected any error, received %q", errs)
				return
			}

			if errs != nil {
				t.Fatal(errs)
			}
			if diff := cmp.Diff(tt.wantSnippet, gotSnippet); diff != "" {
				t.Errorf("snippet not equal (-want, +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantGoTypeData, gotGoTypeData); diff != "" {
				t.Errorf("goTypeData not equal (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestGeneratePerNodeSnippet(t *testing.T) {
	tests := []struct {
		name               string
		inPathStructName   string
		inNodeData         *ypathgen.NodeData
		inFakeRootTypeName string
		wantSnippet        GoPerNodeCodeSnippet
		wantGoTypeData     goTypeData
		wantErr            bool
	}{{
		name:             "scalar leaf",
		inPathStructName: "Container_Leaf",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "int32",
			LocalGoTypeName:       "int32",
			GoFieldName:           "Leaf",
			SubsumingGoStructName: "Container",
			IsLeaf:                true,
			IsScalarField:         true,
			HasDefault:            true,
			YANGPath:              "/container/leaf",
		},
		inFakeRootTypeName: "Device",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "Container_Leaf",
			GetMethod: `
// GetFull retrieves a sample for /container/leaf.
func (n *Container_Leaf) GetFull(t testing.TB) *QualifiedInt32 {
	t.Helper()
	goStruct := &Container{}
	ret := &QualifiedInt32{
		QualifiedType: getFull(t, n, "Container", goStruct, true),
	}
	if !ret.IsPresent() {
		ret.SetVal(goStruct.GetLeaf())
		ret.Present = true
		return ret
	}
	return convertContainer_Leaf(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /container/leaf, erroring out if it is not present.
func (n *Container_Leaf) Get(t testing.TB) int32 {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /container/leaf.
func (n *Container_LeafAny) GetFull(t testing.TB) []*QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &Container{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "Container", goStruct, queryPath, true, true)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertContainer_Leaf(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /container/leaf.
func (n *Container_LeafAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect retrieves a Collection sample for /container/leaf.
func (n *Container_Leaf) Collect(t testing.TB, duration time.Duration) *CollectionInt32 {
	t.Helper()
	return &CollectionInt32{
		c: n.CollectUntil(t, duration, func(*QualifiedInt32) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /container/leaf and evaluates the predicate on all samples.
func (n *Container_Leaf) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedInt32) bool) *CollectionUntilInt32 {
	t.Helper()
	return &CollectionUntilInt32{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &Container{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "Container", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertContainer_Leaf(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedInt32)
			return ok && predicate(val)
		}),
	}
}

// Await waits until /container/leaf is deep-equal to the val and returns all received values.
// If the timeout is exceeded, the test fails fatally.
// To avoid a fatal failure or wait for a generic predicate, use CollectUntil.
func (n *Container_Leaf) Await(t testing.TB, duration time.Duration, val int32) []*QualifiedInt32 {
	t.Helper()
	vals, success := n.CollectUntil(t, duration, func(data *QualifiedInt32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		if len(vals) == 0 {
			t.Fatalf("Await() at /container/leaf failed: no values received")
		}
		t.Fatalf("Await() at /container/leaf failed: want %v, last got %v", val, vals[len(vals) - 1])
	}
	return vals
}

// Collect retrieves a Collection sample for /container/leaf.
func (n *Container_LeafAny) Collect(t testing.TB, duration time.Duration) *CollectionInt32 {
	t.Helper()
	return &CollectionInt32{
		c: n.CollectUntil(t, duration, func(*QualifiedInt32) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /container/leaf and evaluates the predicate on all samples.
func (n *Container_LeafAny) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedInt32) bool) *CollectionUntilInt32 {
	t.Helper()
	return &CollectionUntilInt32{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &Container{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "Container", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertContainer_Leaf(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedInt32)
			return ok && predicate(val)
		}),
	}
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent Container
// and combines the update with an existing QualifiedType to return a *QualifiedInt32.
func convertContainer_Leaf(t testing.TB, qt *genutil.QualifiedType, parent *Container) *QualifiedInt32 {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &QualifiedInt32{
		QualifiedType: qt,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(*val)
	} else {
		qv.Present = false
	}
	return qv
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "int32",
			TransformedGoTypeName: "Int32",
			IsLeaf:                true,
			HasDefault:            true,
		},
	}, {
		name:             "non-scalar leaf",
		inPathStructName: "List_Key",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "[]Binary",
			LocalGoTypeName:       "[]Binary",
			GoFieldName:           "Key",
			SubsumingGoStructName: "List",
			IsLeaf:                true,
			IsScalarField:         false,
			YANGPath:              "/lists/list/key",
		},
		inFakeRootTypeName: "Device",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "List_Key",
			GetMethod: `
// GetFull retrieves a sample for /lists/list/key.
func (n *List_Key) GetFull(t testing.TB) *QualifiedBinarySlice {
	t.Helper()
	goStruct := &List{}
	ret := &QualifiedBinarySlice{
		QualifiedType: getFull(t, n, "List", goStruct, true),
	}
	return convertList_Key(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /lists/list/key, erroring out if it is not present.
func (n *List_Key) Get(t testing.TB) []Binary {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /lists/list/key.
func (n *List_KeyAny) GetFull(t testing.TB) []*QualifiedBinarySlice {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*QualifiedBinarySlice
	for _, prefix := range sortedPrefixes {
		goStruct := &List{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "List", goStruct, queryPath, true, true)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertList_Key(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /lists/list/key.
func (n *List_KeyAny) Get(t testing.TB) [][]Binary {
	t.Helper()
	fulldata := n.GetFull(t)
	var data [][]Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect retrieves a Collection sample for /lists/list/key.
func (n *List_Key) Collect(t testing.TB, duration time.Duration) *CollectionBinarySlice {
	t.Helper()
	return &CollectionBinarySlice{
		c: n.CollectUntil(t, duration, func(*QualifiedBinarySlice) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /lists/list/key and evaluates the predicate on all samples.
func (n *List_Key) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedBinarySlice) bool) *CollectionUntilBinarySlice {
	t.Helper()
	return &CollectionUntilBinarySlice{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &List{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "List", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertList_Key(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedBinarySlice)
			return ok && predicate(val)
		}),
	}
}

// Await waits until /lists/list/key is deep-equal to the val and returns all received values.
// If the timeout is exceeded, the test fails fatally.
// To avoid a fatal failure or wait for a generic predicate, use CollectUntil.
func (n *List_Key) Await(t testing.TB, duration time.Duration, val []Binary) []*QualifiedBinarySlice {
	t.Helper()
	vals, success := n.CollectUntil(t, duration, func(data *QualifiedBinarySlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		if len(vals) == 0 {
			t.Fatalf("Await() at /lists/list/key failed: no values received")
		}
		t.Fatalf("Await() at /lists/list/key failed: want %v, last got %v", val, vals[len(vals) - 1])
	}
	return vals
}

// Collect retrieves a Collection sample for /lists/list/key.
func (n *List_KeyAny) Collect(t testing.TB, duration time.Duration) *CollectionBinarySlice {
	t.Helper()
	return &CollectionBinarySlice{
		c: n.CollectUntil(t, duration, func(*QualifiedBinarySlice) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /lists/list/key and evaluates the predicate on all samples.
func (n *List_KeyAny) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedBinarySlice) bool) *CollectionUntilBinarySlice {
	t.Helper()
	return &CollectionUntilBinarySlice{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &List{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "List", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertList_Key(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedBinarySlice)
			return ok && predicate(val)
		}),
	}
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent List
// and combines the update with an existing QualifiedType to return a *QualifiedBinarySlice.
func convertList_Key(t testing.TB, qt *genutil.QualifiedType, parent *List) *QualifiedBinarySlice {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &QualifiedBinarySlice{
		QualifiedType: qt,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(val)
	} else {
		qv.Present = false
	}
	return qv
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "[]Binary",
			TransformedGoTypeName: "BinarySlice",
			IsLeaf:                true,
		},
	}, {
		name:             "ieeefloat32 special case",
		inPathStructName: "Container_Leaf",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "Binary",
			LocalGoTypeName:       "Binary",
			GoFieldName:           "Leaf",
			SubsumingGoStructName: "Container",
			IsLeaf:                true,
			IsScalarField:         false,
			YANGTypeName:          "ieeefloat32",
			YANGPath:              "/container/leaf",
		},
		inFakeRootTypeName: "Device",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "Container_Leaf",
			GetMethod: `
// GetFull retrieves a sample for /container/leaf.
func (n *Container_Leaf) GetFull(t testing.TB) *QualifiedFloat32 {
	t.Helper()
	goStruct := &Container{}
	ret := &QualifiedFloat32{
		QualifiedType: getFull(t, n, "Container", goStruct, true),
	}
	return convertContainer_Leaf(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /container/leaf, erroring out if it is not present.
func (n *Container_Leaf) Get(t testing.TB) float32 {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /container/leaf.
func (n *Container_LeafAny) GetFull(t testing.TB) []*QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &Container{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "Container", goStruct, queryPath, true, true)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertContainer_Leaf(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /container/leaf.
func (n *Container_LeafAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect retrieves a Collection sample for /container/leaf.
func (n *Container_Leaf) Collect(t testing.TB, duration time.Duration) *CollectionFloat32 {
	t.Helper()
	return &CollectionFloat32{
		c: n.CollectUntil(t, duration, func(*QualifiedFloat32) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /container/leaf and evaluates the predicate on all samples.
func (n *Container_Leaf) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedFloat32) bool) *CollectionUntilFloat32 {
	t.Helper()
	return &CollectionUntilFloat32{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &Container{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "Container", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertContainer_Leaf(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedFloat32)
			return ok && predicate(val)
		}),
	}
}

// Await waits until /container/leaf is deep-equal to the val and returns all received values.
// If the timeout is exceeded, the test fails fatally.
// To avoid a fatal failure or wait for a generic predicate, use CollectUntil.
func (n *Container_Leaf) Await(t testing.TB, duration time.Duration, val float32) []*QualifiedFloat32 {
	t.Helper()
	vals, success := n.CollectUntil(t, duration, func(data *QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		if len(vals) == 0 {
			t.Fatalf("Await() at /container/leaf failed: no values received")
		}
		t.Fatalf("Await() at /container/leaf failed: want %v, last got %v", val, vals[len(vals) - 1])
	}
	return vals
}

// Collect retrieves a Collection sample for /container/leaf.
func (n *Container_LeafAny) Collect(t testing.TB, duration time.Duration) *CollectionFloat32 {
	t.Helper()
	return &CollectionFloat32{
		c: n.CollectUntil(t, duration, func(*QualifiedFloat32) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /container/leaf and evaluates the predicate on all samples.
func (n *Container_LeafAny) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedFloat32) bool) *CollectionUntilFloat32 {
	t.Helper()
	return &CollectionUntilFloat32{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &Container{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "Container", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertContainer_Leaf(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedFloat32)
			return ok && predicate(val)
		}),
	}
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent Container
// and combines the update with an existing QualifiedType to return a *QualifiedFloat32.
func convertContainer_Leaf(t testing.TB, qt *genutil.QualifiedType, parent *Container) *QualifiedFloat32 {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &QualifiedFloat32{
		QualifiedType: qt,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(ygot.BinaryToFloat32(val))
	} else {
		qv.Present = false
	}
	return qv
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "float32",
			TransformedGoTypeName: "Float32",
			IsLeaf:                true,
		},
	}, {
		name:             "ieeefloat32 special case with unexpected type",
		inPathStructName: "Container_Leaf",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "float64", // Expect this to be Binary only.
			LocalGoTypeName:       "float64", // Expect this to be Binary only.
			GoFieldName:           "Leaf",
			SubsumingGoStructName: "Container",
			IsLeaf:                true,
			IsScalarField:         false,
			YANGTypeName:          "ieeefloat32",
		},
		wantErr: true,
	}, {
		name:             "ieeefloat32 slice special case",
		inPathStructName: "List_Key",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "[]Binary",
			LocalGoTypeName:       "[]Binary",
			GoFieldName:           "Key",
			SubsumingGoStructName: "List",
			IsLeaf:                true,
			IsScalarField:         false,
			YANGTypeName:          "ieeefloat32",
			YANGPath:              "/lists/list/key",
		},
		inFakeRootTypeName: "Device",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "List_Key",
			GetMethod: `
// GetFull retrieves a sample for /lists/list/key.
func (n *List_Key) GetFull(t testing.TB) *QualifiedFloat32Slice {
	t.Helper()
	goStruct := &List{}
	ret := &QualifiedFloat32Slice{
		QualifiedType: getFull(t, n, "List", goStruct, true),
	}
	return convertList_Key(t, ret.QualifiedType, goStruct)
}

// Get retrieves a value sample for /lists/list/key, erroring out if it is not present.
func (n *List_Key) Get(t testing.TB) []float32 {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /lists/list/key.
func (n *List_KeyAny) GetFull(t testing.TB) []*QualifiedFloat32Slice {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), true)

	var data []*QualifiedFloat32Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &List{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "List", goStruct, queryPath, true, true)
		if err != nil {
			t.Fatal(err)
		}
		qv := convertList_Key(t, qt, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /lists/list/key.
func (n *List_KeyAny) Get(t testing.TB) [][]float32 {
	t.Helper()
	fulldata := n.GetFull(t)
	var data [][]float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect retrieves a Collection sample for /lists/list/key.
func (n *List_Key) Collect(t testing.TB, duration time.Duration) *CollectionFloat32Slice {
	t.Helper()
	return &CollectionFloat32Slice{
		c: n.CollectUntil(t, duration, func(*QualifiedFloat32Slice) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /lists/list/key and evaluates the predicate on all samples.
func (n *List_Key) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedFloat32Slice) bool) *CollectionUntilFloat32Slice {
	t.Helper()
	return &CollectionUntilFloat32Slice{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &List{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "List", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertList_Key(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedFloat32Slice)
			return ok && predicate(val)
		}),
	}
}

// Await waits until /lists/list/key is deep-equal to the val and returns all received values.
// If the timeout is exceeded, the test fails fatally.
// To avoid a fatal failure or wait for a generic predicate, use CollectUntil.
func (n *List_Key) Await(t testing.TB, duration time.Duration, val []float32) []*QualifiedFloat32Slice {
	t.Helper()
	vals, success := n.CollectUntil(t, duration, func(data *QualifiedFloat32Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		if len(vals) == 0 {
			t.Fatalf("Await() at /lists/list/key failed: no values received")
		}
		t.Fatalf("Await() at /lists/list/key failed: want %v, last got %v", val, vals[len(vals) - 1])
	}
	return vals
}

// Collect retrieves a Collection sample for /lists/list/key.
func (n *List_KeyAny) Collect(t testing.TB, duration time.Duration) *CollectionFloat32Slice {
	t.Helper()
	return &CollectionFloat32Slice{
		c: n.CollectUntil(t, duration, func(*QualifiedFloat32Slice) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for /lists/list/key and evaluates the predicate on all samples.
func (n *List_KeyAny) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *QualifiedFloat32Slice) bool) *CollectionUntilFloat32Slice {
	t.Helper()
	return &CollectionUntilFloat32Slice{
		c: genutil.CollectUntil(t, n, duration, func(upd *genutil.DataPoint) (genutil.QualifiedValue, error) {
			parentPtr := &List{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := genutil.Unmarshal(t, []*genutil.DataPoint{upd}, getSchema(), "List", parentPtr, nil, true, true)
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convertList_Key(t, qv, parentPtr), nil
		},
		func(qualVal genutil.QualifiedValue) bool {
			val, ok := qualVal.(*QualifiedFloat32Slice)
			return ok && predicate(val)
		}),
	}
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent List
// and combines the update with an existing QualifiedType to return a *QualifiedFloat32Slice.
func convertList_Key(t testing.TB, qt *genutil.QualifiedType, parent *List) *QualifiedFloat32Slice {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &QualifiedFloat32Slice{
		QualifiedType: qt,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		qv.SetVal(binarySliceToFloat32(val))
	} else {
		qv.Present = false
	}
	return qv
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "[]float32",
			TransformedGoTypeName: "Float32Slice",
			IsLeaf:                true,
		},
	}, {
		name:             "non-leaf",
		inPathStructName: "SuperContainer_Container",
		inNodeData: &ypathgen.NodeData{
			GoTypeName:            "*SuperContainer_Container",
			LocalGoTypeName:       "*SuperContainer_Container",
			GoFieldName:           "Container",
			SubsumingGoStructName: "SuperContainer_Container",
			IsLeaf:                false,
			IsScalarField:         false,
			YANGPath:              "/super-container/container",
		},
		inFakeRootTypeName: "FakeRoot",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "SuperContainer_Container",
			GetMethod: `
// GetFull retrieves a sample for /super-container/container.
func (n *SuperContainer_Container) GetFull(t testing.TB) *QualifiedSuperContainer_Container {
	t.Helper()
	goStruct := &SuperContainer_Container{}
	ret := &QualifiedSuperContainer_Container{
		QualifiedType: getFull(t, n, "SuperContainer_Container", goStruct, false),
	}
	if ret.IsPresent() {
		ret.SetVal(goStruct)
	}
	return ret
}

// Get retrieves a value sample for /super-container/container, erroring out if it is not present.
func (n *SuperContainer_Container) Get(t testing.TB) *SuperContainer_Container {
	t.Helper()
	return n.GetFull(t).Val(t)
}

// GetFull retrieves a list of samples for /super-container/container.
func (n *SuperContainer_ContainerAny) GetFull(t testing.TB) []*QualifiedSuperContainer_Container {
	t.Helper()
	datapoints, queryPath := genutil.Get(t, n, false)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), false)

	var data []*QualifiedSuperContainer_Container
	for _, prefix := range sortedPrefixes {
		goStruct := &SuperContainer_Container{}
		qt, err := genutil.Unmarshal(t, datapointGroups[prefix], getSchema(), "SuperContainer_Container", goStruct, queryPath, false, true)
		if err != nil {
			t.Fatal(err)
		}
		if !qt.IsPresent() {
			continue
		}
		qv := (&QualifiedSuperContainer_Container{
			QualifiedType: qt,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for /super-container/container.
func (n *SuperContainer_ContainerAny) Get(t testing.TB) []*SuperContainer_Container {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []*SuperContainer_Container
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
		},
		wantGoTypeData: goTypeData{
			GoTypeName:            "*SuperContainer_Container",
			TransformedGoTypeName: "SuperContainer_Container",
			IsLeaf:                false,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSnippet, gotGoTypeData, errs := generatePerNodeSnippet(tt.inPathStructName, tt.inNodeData, tt.inFakeRootTypeName, "", true)
			if tt.wantErr {
				if errs == nil {
					t.Fatalf("want error but did not get any.\ngotSnippet: %v\n\ngotGoTypeData: %v", gotSnippet, gotGoTypeData)
				}
				t.Logf("success: expected any error, received %q", errs)
				return
			}

			if errs != nil {
				t.Fatal(errs)
			}
			if diff := cmp.Diff(tt.wantSnippet, gotSnippet); diff != "" {
				t.Errorf("snippet not equal (-want, +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantGoTypeData, gotGoTypeData); diff != "" {
				t.Errorf("goTypeData not equal (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestGeneratePerTypeSnippet(t *testing.T) {
	tests := []struct {
		name               string
		inGoTypeData       goTypeData
		inFakeRootTypeName string
		wantSnippet        GoReturnTypeCodeSnippet
	}{{
		name: "scalar type",
		inGoTypeData: goTypeData{
			GoTypeName:            "int32",
			TransformedGoTypeName: "Int32",
			IsLeaf:                true,
		},
		inFakeRootTypeName: "Device",
		wantSnippet: GoReturnTypeCodeSnippet{
			TypeName: "int32",
			QualifiedType: `
// QualifiedInt32 is a int32 with a corresponding timestamp.
type QualifiedInt32 struct {
	*genutil.QualifiedType
	val int32 // val is the sample value.
}

func (q *QualifiedInt32) String() string {
	return genutil.QualifiedTypeString(q.val, q.QualifiedType)
}

// Val returns the value of the int32 sample, erroring out if not present.
func (q *QualifiedInt32) Val(t testing.TB) int32 {
	t.Helper()
	if !q.Present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the int32 sample.
func (q *QualifiedInt32) SetVal(v int32) *QualifiedInt32 {
	q.val = v
	return q
}

`,
			CollectionType: `
// CollectionInt32 is a telemetry Collection whose Await method returns a slice of int32 samples.
type CollectionInt32 struct {
	c *CollectionUntilInt32
}

// Await blocks for the telemetry collection to be complete, and then returns the slice of samples received.
func (u *CollectionInt32) Await(t testing.TB) []*QualifiedInt32 {
	t.Helper()
	data, _ := u.c.Await(t)
	return data
}

// CollectionUntilInt32 is a telemetry Collection whose Await method returns a slice of int32 samples.
type CollectionUntilInt32 struct {
	c *genutil.Collection
}

// Await blocks for the telemetry collection to be complete or the predicate to be true whichever is first.
// The received data and the status of the predicate are returned.
func (u *CollectionUntilInt32) Await(t testing.TB) ([]*QualifiedInt32, bool) {
	t.Helper()
	var ret []*QualifiedInt32
	updates, predTrue := u.c.Await(t)
	for _, upd := range updates {
		ret = append(ret, upd.(*QualifiedInt32))
	}
	return ret, predTrue
}
`,
		},
	}, {
		name: "non-scalar type",
		inGoTypeData: goTypeData{
			GoTypeName:            "[]Binary",
			TransformedGoTypeName: "BinarySlice",
			IsLeaf:                true,
		},
		inFakeRootTypeName: "FakeRoot",
		wantSnippet: GoReturnTypeCodeSnippet{
			TypeName: "[]Binary",
			QualifiedType: `
// QualifiedBinarySlice is a []Binary with a corresponding timestamp.
type QualifiedBinarySlice struct {
	*genutil.QualifiedType
	val []Binary // val is the sample value.
}

func (q *QualifiedBinarySlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.QualifiedType)
}

// Val returns the value of the []Binary sample, erroring out if not present.
func (q *QualifiedBinarySlice) Val(t testing.TB) []Binary {
	t.Helper()
	if !q.Present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the []Binary sample.
func (q *QualifiedBinarySlice) SetVal(v []Binary) *QualifiedBinarySlice {
	q.val = v
	return q
}

`,
			CollectionType: `
// CollectionBinarySlice is a telemetry Collection whose Await method returns a slice of []Binary samples.
type CollectionBinarySlice struct {
	c *CollectionUntilBinarySlice
}

// Await blocks for the telemetry collection to be complete, and then returns the slice of samples received.
func (u *CollectionBinarySlice) Await(t testing.TB) []*QualifiedBinarySlice {
	t.Helper()
	data, _ := u.c.Await(t)
	return data
}

// CollectionUntilBinarySlice is a telemetry Collection whose Await method returns a slice of []Binary samples.
type CollectionUntilBinarySlice struct {
	c *genutil.Collection
}

// Await blocks for the telemetry collection to be complete or the predicate to be true whichever is first.
// The received data and the status of the predicate are returned.
func (u *CollectionUntilBinarySlice) Await(t testing.TB) ([]*QualifiedBinarySlice, bool) {
	t.Helper()
	var ret []*QualifiedBinarySlice
	updates, predTrue := u.c.Await(t)
	for _, upd := range updates {
		ret = append(ret, upd.(*QualifiedBinarySlice))
	}
	return ret, predTrue
}
`,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSnippet, errs := generatePerTypeSnippet(tt.inGoTypeData, tt.inFakeRootTypeName, true)
			if errs != nil {
				t.Fatal(errs)
			}
			if diff := cmp.Diff(tt.wantSnippet, gotSnippet); diff != "" {
				t.Errorf("snippet not equal (-want, +got):\n%s", diff)
			}
		})
	}
}
