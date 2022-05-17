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

		t.Run(tt.name, func(t *testing.T) {
			code, errs := tt.inConfig.GenerateCode(tt.inFiles, tt.inIncludePaths, config)
			if gotErr := errs != nil; gotErr != tt.wantErr {
				t.Fatalf("got errors: %v, wantErr: %v", errs, tt.wantErr)
			} else if gotErr {
				// Received expected error.
				return
			}

			wantFile := filepath.Join(datapath, tt.resultsFolder, dirName, generatedFileName)
			got := code.String(tt.inConfig.PackageName)
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
				gotAttempt, _ := tt.inConfig.GenerateCode(tt.inFiles, tt.inIncludePaths, config)
				if diff := codeDiff(t, got, gotAttempt.String(tt.inConfig.PackageName)); diff != "" {
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
		inConfigPkgAccessor       string
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
// Lookup fetches the value at /container/leaf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Container_Leaf) Lookup(t testing.TB) *QualifiedInt32 {
	t.Helper()
	goStruct := &Container{}
	md, ok := Lookup(t, n, "Container", goStruct, true, false)
	if ok {
		return convertContainer_Leaf(t, md, goStruct)
	}
	return (&QualifiedInt32{
		Metadata: md,
	}).SetVal(goStruct.GetLeaf())
}

// Get fetches the value at /container/leaf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Container_Leaf) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /container/leaf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Container_LeafAny) Lookup(t testing.TB) []*QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &Container{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "Container", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertContainer_Leaf(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /container/leaf with a ONCE subscription.
func (n *Container_LeafAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent Container
// and combines the update with an existing Metadata to return a *QualifiedInt32.
func convertContainer_Leaf(t testing.TB, md *genutil.Metadata, parent *Container) *QualifiedInt32 {
	t.Helper()
	qv := &QualifiedInt32{
		Metadata: md,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
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
		name:             "scalar leaf with accessors",
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
		inConfigPkgAccessor:       "config.",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "Container_Leaf",
			GetMethod: `
// Lookup fetches the value at /container/leaf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Container_Leaf) Lookup(t testing.TB) *oc.QualifiedInt32 {
	t.Helper()
	goStruct := &oc.Container{}
	md, ok := oc.Lookup(t, n, "Container", goStruct, true, false)
	if ok {
		return convertContainer_Leaf(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /container/leaf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Container_Leaf) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /container/leaf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Container_LeafAny) Lookup(t testing.TB) []*oc.QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Container{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Container", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertContainer_Leaf(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /container/leaf with a ONCE subscription.
func (n *Container_LeafAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent oc.Container
// and combines the update with an existing Metadata to return a *oc.QualifiedInt32.
func convertContainer_Leaf(t testing.TB, md *genutil.Metadata, parent *oc.Container) *oc.QualifiedInt32 {
	t.Helper()
	qv := &oc.QualifiedInt32{
		Metadata: md,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
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
func (n *Container_Leaf) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /container/leaf.
func (n *Container_Leaf) Replace(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchReplace(t testing.TB, b *config.SetRequestBatch, val int32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /container/leaf.
func (n *Container_Leaf) Update(t testing.TB, val int32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /container/leaf in the given batch object.
func (n *Container_Leaf) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val int32) {
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
// Lookup fetches the value at /lists/list/key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *List_Key) Lookup(t testing.TB) *QualifiedBinarySlice {
	t.Helper()
	goStruct := &List{}
	md, ok := Lookup(t, n, "List", goStruct, true, false)
	if ok {
		return convertList_Key(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /lists/list/key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *List_Key) Get(t testing.TB) []Binary {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /lists/list/key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *List_KeyAny) Lookup(t testing.TB) []*QualifiedBinarySlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedBinarySlice
	for _, prefix := range sortedPrefixes {
		goStruct := &List{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "List", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertList_Key(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /lists/list/key with a ONCE subscription.
func (n *List_KeyAny) Get(t testing.TB) [][]Binary {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent List
// and combines the update with an existing Metadata to return a *QualifiedBinarySlice.
func convertList_Key(t testing.TB, md *genutil.Metadata, parent *List) *QualifiedBinarySlice {
	t.Helper()
	qv := &QualifiedBinarySlice{
		Metadata: md,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
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
		name:             "non-scalar leaf with accessors",
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
		inConfigPkgAccessor:       "config.",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "List_Key",
			GetMethod: `
// Lookup fetches the value at /lists/list/key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *List_Key) Lookup(t testing.TB) *oc.QualifiedBinarySlice {
	t.Helper()
	goStruct := &oc.List{}
	md, ok := oc.Lookup(t, n, "List", goStruct, true, false)
	if ok {
		return convertList_Key(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /lists/list/key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *List_Key) Get(t testing.TB) []oc.Binary {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /lists/list/key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *List_KeyAny) Lookup(t testing.TB) []*oc.QualifiedBinarySlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBinarySlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.List{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "List", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertList_Key(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /lists/list/key with a ONCE subscription.
func (n *List_KeyAny) Get(t testing.TB) [][]oc.Binary {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]oc.Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent oc.List
// and combines the update with an existing Metadata to return a *oc.QualifiedBinarySlice.
func convertList_Key(t testing.TB, md *genutil.Metadata, parent *oc.List) *oc.QualifiedBinarySlice {
	t.Helper()
	qv := &oc.QualifiedBinarySlice{
		Metadata: md,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
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
func (n *List_Key) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /lists/list/key.
func (n *List_Key) Replace(t testing.TB, val []oc.Binary) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []oc.Binary) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /lists/list/key.
func (n *List_Key) Update(t testing.TB, val []oc.Binary) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /lists/list/key in the given batch object.
func (n *List_Key) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []oc.Binary) {
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
// Lookup fetches the value at /super-container/container with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *SuperContainer_Container) Lookup(t testing.TB) *QualifiedSuperContainer_Container {
	t.Helper()
	goStruct := &SuperContainer_Container{}
	md, ok := Lookup(t, n, "SuperContainer_Container", goStruct, false, false)
	if ok {
		return (&QualifiedSuperContainer_Container{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /super-container/container with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *SuperContainer_Container) Get(t testing.TB) *SuperContainer_Container {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /super-container/container with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *SuperContainer_ContainerAny) Lookup(t testing.TB) []*QualifiedSuperContainer_Container {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedSuperContainer_Container
	for _, prefix := range sortedPrefixes {
		goStruct := &SuperContainer_Container{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "SuperContainer_Container", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&QualifiedSuperContainer_Container{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /super-container/container with a ONCE subscription.
func (n *SuperContainer_ContainerAny) Get(t testing.TB) []*SuperContainer_Container {
	t.Helper()
	fulldata := n.Lookup(t)
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
		name:             "non-leaf with accessors",
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
		inConfigPkgAccessor:       "config.",
		wantSnippet: GoPerNodeCodeSnippet{
			PathStructName: "SuperContainer_Container",
			GetMethod: `
// Lookup fetches the value at /super-container/container with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *SuperContainer_Container) Lookup(t testing.TB) *oc.QualifiedSuperContainer_Container {
	t.Helper()
	goStruct := &oc.SuperContainer_Container{}
	md, ok := oc.Lookup(t, n, "SuperContainer_Container", goStruct, false, false)
	if ok {
		return (&oc.QualifiedSuperContainer_Container{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /super-container/container with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *SuperContainer_Container) Get(t testing.TB) *oc.SuperContainer_Container {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /super-container/container with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *SuperContainer_ContainerAny) Lookup(t testing.TB) []*oc.QualifiedSuperContainer_Container {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedSuperContainer_Container
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.SuperContainer_Container{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "SuperContainer_Container", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedSuperContainer_Container{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /super-container/container with a ONCE subscription.
func (n *SuperContainer_ContainerAny) Get(t testing.TB) []*oc.SuperContainer_Container {
	t.Helper()
	fulldata := n.Lookup(t)
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
func (n *SuperContainer_Container) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /super-container/container.
func (n *SuperContainer_Container) Replace(t testing.TB, val *oc.SuperContainer_Container) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.SuperContainer_Container) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /super-container/container.
func (n *SuperContainer_Container) Update(t testing.TB, val *oc.SuperContainer_Container) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /super-container/container in the given batch object.
func (n *SuperContainer_Container) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.SuperContainer_Container) {
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
			gotSnippet, gotGoTypeData, errs := generatePerNodeConfigSnippet(tt.inPathStructName, tt.inNodeData, tt.inFakeRootTypeName, tt.inSchemaStructPkgAccessor, tt.inConfigPkgAccessor, false)
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
// Lookup fetches the value at /container/leaf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Container_Leaf) Lookup(t testing.TB) *QualifiedInt32 {
	t.Helper()
	goStruct := &Container{}
	md, ok := Lookup(t, n, "Container", goStruct, true, true)
	if ok {
		return convertContainer_Leaf(t, md, goStruct)
	}
	return (&QualifiedInt32{
		Metadata: md,
	}).SetVal(goStruct.GetLeaf())
}

// Get fetches the value at /container/leaf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Container_Leaf) Get(t testing.TB) int32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /container/leaf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Container_LeafAny) Lookup(t testing.TB) []*QualifiedInt32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedInt32
	for _, prefix := range sortedPrefixes {
		goStruct := &Container{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "Container", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertContainer_Leaf(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /container/leaf with a ONCE subscription.
func (n *Container_LeafAny) Get(t testing.TB) []int32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []int32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect starts an asynchronous collection of the values at /container/leaf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Container_Leaf) Collect(t testing.TB, duration time.Duration) *CollectionInt32 {
	t.Helper()
	c := &CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Container_Leaf(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedInt32) bool) *Int32Watcher {
	t.Helper()
	w := &Int32Watcher{}
	gs := &Container{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, GetSchema(), "Container", gs, queryPath, true, true)
		return []genutil.QualifiedValue{convertContainer_Leaf(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /container/leaf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Container_Leaf) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedInt32) bool) *Int32Watcher {
	t.Helper()
	return watch_Container_Leaf(t, n, timeout, predicate)
}

// Await observes values at /container/leaf with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Container_Leaf) Await(t testing.TB, timeout time.Duration, val int32) *QualifiedInt32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *QualifiedInt32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /container/leaf failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /container/leaf to the batch object.
func (n *Container_Leaf) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /container/leaf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Container_LeafAny) Collect(t testing.TB, duration time.Duration) *CollectionInt32 {
	t.Helper()
	c := &CollectionInt32{}
	c.W = n.Watch(t, duration, func(v *QualifiedInt32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Container_LeafAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedInt32) bool) *Int32Watcher {
	t.Helper()
	w := &Int32Watcher{}
	structs := map[string]*Container{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &Container{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], GetSchema(), "Container", structs[pre], queryPath, true, true)
			qv := convertContainer_Leaf(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedInt32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /container/leaf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Container_LeafAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedInt32) bool) *Int32Watcher {
	t.Helper()
	return watch_Container_LeafAny(t, n, timeout, predicate)
}

// Batch adds /container/leaf to the batch object.
func (n *Container_LeafAny) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent Container
// and combines the update with an existing Metadata to return a *QualifiedInt32.
func convertContainer_Leaf(t testing.TB, md *genutil.Metadata, parent *Container) *QualifiedInt32 {
	t.Helper()
	qv := &QualifiedInt32{
		Metadata: md,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
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
// Lookup fetches the value at /lists/list/key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *List_Key) Lookup(t testing.TB) *QualifiedBinarySlice {
	t.Helper()
	goStruct := &List{}
	md, ok := Lookup(t, n, "List", goStruct, true, true)
	if ok {
		return convertList_Key(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /lists/list/key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *List_Key) Get(t testing.TB) []Binary {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /lists/list/key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *List_KeyAny) Lookup(t testing.TB) []*QualifiedBinarySlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedBinarySlice
	for _, prefix := range sortedPrefixes {
		goStruct := &List{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "List", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertList_Key(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /lists/list/key with a ONCE subscription.
func (n *List_KeyAny) Get(t testing.TB) [][]Binary {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]Binary
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect starts an asynchronous collection of the values at /lists/list/key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *List_Key) Collect(t testing.TB, duration time.Duration) *CollectionBinarySlice {
	t.Helper()
	c := &CollectionBinarySlice{}
	c.W = n.Watch(t, duration, func(v *QualifiedBinarySlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_List_Key(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedBinarySlice) bool) *BinarySliceWatcher {
	t.Helper()
	w := &BinarySliceWatcher{}
	gs := &List{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, GetSchema(), "List", gs, queryPath, true, true)
		return []genutil.QualifiedValue{convertList_Key(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedBinarySlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /lists/list/key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *List_Key) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedBinarySlice) bool) *BinarySliceWatcher {
	t.Helper()
	return watch_List_Key(t, n, timeout, predicate)
}

// Await observes values at /lists/list/key with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *List_Key) Await(t testing.TB, timeout time.Duration, val []Binary) *QualifiedBinarySlice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *QualifiedBinarySlice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /lists/list/key failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /lists/list/key to the batch object.
func (n *List_Key) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /lists/list/key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *List_KeyAny) Collect(t testing.TB, duration time.Duration) *CollectionBinarySlice {
	t.Helper()
	c := &CollectionBinarySlice{}
	c.W = n.Watch(t, duration, func(v *QualifiedBinarySlice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_List_KeyAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedBinarySlice) bool) *BinarySliceWatcher {
	t.Helper()
	w := &BinarySliceWatcher{}
	structs := map[string]*List{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &List{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], GetSchema(), "List", structs[pre], queryPath, true, true)
			qv := convertList_Key(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedBinarySlice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /lists/list/key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *List_KeyAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedBinarySlice) bool) *BinarySliceWatcher {
	t.Helper()
	return watch_List_KeyAny(t, n, timeout, predicate)
}

// Batch adds /lists/list/key to the batch object.
func (n *List_KeyAny) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent List
// and combines the update with an existing Metadata to return a *QualifiedBinarySlice.
func convertList_Key(t testing.TB, md *genutil.Metadata, parent *List) *QualifiedBinarySlice {
	t.Helper()
	qv := &QualifiedBinarySlice{
		Metadata: md,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
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
// Lookup fetches the value at /container/leaf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Container_Leaf) Lookup(t testing.TB) *QualifiedFloat32 {
	t.Helper()
	goStruct := &Container{}
	md, ok := Lookup(t, n, "Container", goStruct, true, true)
	if ok {
		return convertContainer_Leaf(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /container/leaf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Container_Leaf) Get(t testing.TB) float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /container/leaf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Container_LeafAny) Lookup(t testing.TB) []*QualifiedFloat32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedFloat32
	for _, prefix := range sortedPrefixes {
		goStruct := &Container{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "Container", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertContainer_Leaf(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /container/leaf with a ONCE subscription.
func (n *Container_LeafAny) Get(t testing.TB) []float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect starts an asynchronous collection of the values at /container/leaf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Container_Leaf) Collect(t testing.TB, duration time.Duration) *CollectionFloat32 {
	t.Helper()
	c := &CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Container_Leaf(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedFloat32) bool) *Float32Watcher {
	t.Helper()
	w := &Float32Watcher{}
	gs := &Container{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, GetSchema(), "Container", gs, queryPath, true, true)
		return []genutil.QualifiedValue{convertContainer_Leaf(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /container/leaf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Container_Leaf) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedFloat32) bool) *Float32Watcher {
	t.Helper()
	return watch_Container_Leaf(t, n, timeout, predicate)
}

// Await observes values at /container/leaf with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Container_Leaf) Await(t testing.TB, timeout time.Duration, val float32) *QualifiedFloat32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *QualifiedFloat32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /container/leaf failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /container/leaf to the batch object.
func (n *Container_Leaf) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /container/leaf with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Container_LeafAny) Collect(t testing.TB, duration time.Duration) *CollectionFloat32 {
	t.Helper()
	c := &CollectionFloat32{}
	c.W = n.Watch(t, duration, func(v *QualifiedFloat32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Container_LeafAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedFloat32) bool) *Float32Watcher {
	t.Helper()
	w := &Float32Watcher{}
	structs := map[string]*Container{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &Container{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], GetSchema(), "Container", structs[pre], queryPath, true, true)
			qv := convertContainer_Leaf(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedFloat32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /container/leaf with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Container_LeafAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedFloat32) bool) *Float32Watcher {
	t.Helper()
	return watch_Container_LeafAny(t, n, timeout, predicate)
}

// Batch adds /container/leaf to the batch object.
func (n *Container_LeafAny) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}
`,
			ConvertHelper: `
// convertContainer_Leaf extracts the value of the leaf Leaf from its parent Container
// and combines the update with an existing Metadata to return a *QualifiedFloat32.
func convertContainer_Leaf(t testing.TB, md *genutil.Metadata, parent *Container) *QualifiedFloat32 {
	t.Helper()
	qv := &QualifiedFloat32{
		Metadata: md,
	}
	val := parent.Leaf
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(ygot.BinaryToFloat32(val))
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
// Lookup fetches the value at /lists/list/key with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *List_Key) Lookup(t testing.TB) *QualifiedFloat32Slice {
	t.Helper()
	goStruct := &List{}
	md, ok := Lookup(t, n, "List", goStruct, true, true)
	if ok {
		return convertList_Key(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /lists/list/key with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *List_Key) Get(t testing.TB) []float32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /lists/list/key with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *List_KeyAny) Lookup(t testing.TB) []*QualifiedFloat32Slice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedFloat32Slice
	for _, prefix := range sortedPrefixes {
		goStruct := &List{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "List", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertList_Key(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /lists/list/key with a ONCE subscription.
func (n *List_KeyAny) Get(t testing.TB) [][]float32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]float32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect starts an asynchronous collection of the values at /lists/list/key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *List_Key) Collect(t testing.TB, duration time.Duration) *CollectionFloat32Slice {
	t.Helper()
	c := &CollectionFloat32Slice{}
	c.W = n.Watch(t, duration, func(v *QualifiedFloat32Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_List_Key(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedFloat32Slice) bool) *Float32SliceWatcher {
	t.Helper()
	w := &Float32SliceWatcher{}
	gs := &List{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, GetSchema(), "List", gs, queryPath, true, true)
		return []genutil.QualifiedValue{convertList_Key(t, md, gs)}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedFloat32Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /lists/list/key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *List_Key) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedFloat32Slice) bool) *Float32SliceWatcher {
	t.Helper()
	return watch_List_Key(t, n, timeout, predicate)
}

// Await observes values at /lists/list/key with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *List_Key) Await(t testing.TB, timeout time.Duration, val []float32) *QualifiedFloat32Slice {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *QualifiedFloat32Slice) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /lists/list/key failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /lists/list/key to the batch object.
func (n *List_Key) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /lists/list/key with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *List_KeyAny) Collect(t testing.TB, duration time.Duration) *CollectionFloat32Slice {
	t.Helper()
	c := &CollectionFloat32Slice{}
	c.W = n.Watch(t, duration, func(v *QualifiedFloat32Slice) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_List_KeyAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedFloat32Slice) bool) *Float32SliceWatcher {
	t.Helper()
	w := &Float32SliceWatcher{}
	structs := map[string]*List{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &List{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], GetSchema(), "List", structs[pre], queryPath, true, true)
			qv := convertList_Key(t, md, structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedFloat32Slice)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /lists/list/key with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *List_KeyAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedFloat32Slice) bool) *Float32SliceWatcher {
	t.Helper()
	return watch_List_KeyAny(t, n, timeout, predicate)
}

// Batch adds /lists/list/key to the batch object.
func (n *List_KeyAny) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}
`,
			ConvertHelper: `
// convertList_Key extracts the value of the leaf Key from its parent List
// and combines the update with an existing Metadata to return a *QualifiedFloat32Slice.
func convertList_Key(t testing.TB, md *genutil.Metadata, parent *List) *QualifiedFloat32Slice {
	t.Helper()
	qv := &QualifiedFloat32Slice{
		Metadata: md,
	}
	val := parent.Key
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(binarySliceToFloat32(val))
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
// Lookup fetches the value at /super-container/container with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *SuperContainer_Container) Lookup(t testing.TB) *QualifiedSuperContainer_Container {
	t.Helper()
	goStruct := &SuperContainer_Container{}
	md, ok := Lookup(t, n, "SuperContainer_Container", goStruct, false, true)
	if ok {
		return (&QualifiedSuperContainer_Container{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /super-container/container with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *SuperContainer_Container) Get(t testing.TB) *SuperContainer_Container {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /super-container/container with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *SuperContainer_ContainerAny) Lookup(t testing.TB) []*QualifiedSuperContainer_Container {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*QualifiedSuperContainer_Container
	for _, prefix := range sortedPrefixes {
		goStruct := &SuperContainer_Container{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], GetSchema(), "SuperContainer_Container", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&QualifiedSuperContainer_Container{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /super-container/container with a ONCE subscription.
func (n *SuperContainer_ContainerAny) Get(t testing.TB) []*SuperContainer_Container {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*SuperContainer_Container
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}
`,
			CollectMethod: `
// Collect starts an asynchronous collection of the values at /super-container/container with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *SuperContainer_Container) Collect(t testing.TB, duration time.Duration) *CollectionSuperContainer_Container {
	t.Helper()
	c := &CollectionSuperContainer_Container{}
	c.W = n.Watch(t, duration, func(v *QualifiedSuperContainer_Container) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&QualifiedSuperContainer_Container{
			Metadata: v.Metadata,
		}).SetVal(copy.(*SuperContainer_Container)))
		return false
	})
	return c
}

func watch_SuperContainer_Container(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedSuperContainer_Container) bool) *SuperContainer_ContainerWatcher {
	t.Helper()
	w := &SuperContainer_ContainerWatcher{}
	gs := &SuperContainer_Container{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, GetSchema(), "SuperContainer_Container", gs, queryPath, false, true)
		qv := (&QualifiedSuperContainer_Container{
			Metadata: md,
		}).SetVal(gs)
		return []genutil.QualifiedValue{qv}, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedSuperContainer_Container)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /super-container/container with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *SuperContainer_Container) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedSuperContainer_Container) bool) *SuperContainer_ContainerWatcher {
	t.Helper()
	return watch_SuperContainer_Container(t, n, timeout, predicate)
}

// Await observes values at /super-container/container with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *SuperContainer_Container) Await(t testing.TB, timeout time.Duration, val *SuperContainer_Container) *QualifiedSuperContainer_Container {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *QualifiedSuperContainer_Container) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /super-container/container failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /super-container/container to the batch object.
func (n *SuperContainer_Container) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /super-container/container with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *SuperContainer_ContainerAny) Collect(t testing.TB, duration time.Duration) *CollectionSuperContainer_Container {
	t.Helper()
	c := &CollectionSuperContainer_Container{}
	c.W = n.Watch(t, duration, func(v *QualifiedSuperContainer_Container) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_SuperContainer_ContainerAny(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *QualifiedSuperContainer_Container) bool) *SuperContainer_ContainerWatcher {
	t.Helper()
	w := &SuperContainer_ContainerWatcher{}
	structs := map[string]*SuperContainer_Container{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) ([]genutil.QualifiedValue, error) {
		t.Helper()
		datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, upd, uint(len(queryPath.Elem)))
		var currStructs []genutil.QualifiedValue
		for _, pre := range sortedPrefixes {
			if len(datapointGroups[pre]) == 0 {
				continue
			}
			if _, ok := structs[pre]; !ok {
				structs[pre] = &SuperContainer_Container{}
			}
			md, _ := genutil.MustUnmarshal(t, datapointGroups[pre], GetSchema(), "SuperContainer_Container", structs[pre], queryPath, false, true)
			qv := (&QualifiedSuperContainer_Container{
				Metadata: md,
			}).SetVal(structs[pre])
			currStructs = append(currStructs, qv)
		}
		return currStructs, nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*QualifiedSuperContainer_Container)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /super-container/container with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *SuperContainer_ContainerAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *QualifiedSuperContainer_Container) bool) *SuperContainer_ContainerWatcher {
	t.Helper()
	return watch_SuperContainer_ContainerAny(t, n, timeout, predicate)
}

// Batch adds /super-container/container to the batch object.
func (n *SuperContainer_ContainerAny) Batch(t testing.TB, b *Batch) {
	t.Helper()
	MustAddToBatch(t, b, n)
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
	*genutil.Metadata
	val int32 // val is the sample value.
	present bool
}

func (q *QualifiedInt32) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the int32 sample, erroring out if not present.
func (q *QualifiedInt32) Val(t testing.TB) int32 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
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
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInt32) IsPresent() bool {
	return q != nil && q.present
}

`,
			CollectionType: `
// CollectionInt32 is a telemetry Collection whose Await method returns a slice of int32 samples.
type CollectionInt32 struct {
	W *Int32Watcher
	Data []*QualifiedInt32
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInt32) Await(t testing.TB) []*QualifiedInt32 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Int32Watcher observes a stream of int32 samples.
type Int32Watcher struct {
	W *genutil.Watcher
	LastVal *QualifiedInt32
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Int32Watcher) Await(t testing.TB) (*QualifiedInt32, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
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
	*genutil.Metadata
	val []Binary // val is the sample value.
	present bool
}

func (q *QualifiedBinarySlice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the []Binary sample, erroring out if not present.
func (q *QualifiedBinarySlice) Val(t testing.TB) []Binary {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
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
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBinarySlice) IsPresent() bool {
	return q != nil && q.present
}

`,
			CollectionType: `
// CollectionBinarySlice is a telemetry Collection whose Await method returns a slice of []Binary samples.
type CollectionBinarySlice struct {
	W *BinarySliceWatcher
	Data []*QualifiedBinarySlice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBinarySlice) Await(t testing.TB) []*QualifiedBinarySlice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BinarySliceWatcher observes a stream of []Binary samples.
type BinarySliceWatcher struct {
	W *genutil.Watcher
	LastVal *QualifiedBinarySlice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BinarySliceWatcher) Await(t testing.TB) (*QualifiedBinarySlice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
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
