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

// Package telemgen is a library for generating an API for creating gNMI paths
// from a YANG model and allowing ONDATRA telemetry calls to be made targeting
// those paths. The generated code is a hierarchy of structs which follows a
// similar structure to ygen's compressed output, where each struct identifies
// a compressed schema node, the difference being leaf nodes are also given
// generated structs. The hierarchy, instead of being based on struct fields,
// is formed via constructor methods that produce structs mapping to the child
// schema nodes of the calling object. In addition to constructor methods,
// telemetry methods are also generated for each struct, thereby providing for
// a way of doing telemetry calls on a given schema that is by construction
// less error-prone.
package telemgen

import (
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/genutil"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ypathgen"
)

// Static default configuration values that differ from the zero value for their types.
const (
	// defaultPathPackageName specifies the default name that should be
	// used for the generated Go package.
	defaultPathPackageName = "telemetry"
	// schemaStructPkgAlias is the package alias of the schema struct
	// package when the struct package is to be generated in a
	// separate package.
	schemaStructPkgAlias = "oc"
	// defaultFakeRootName is the default name for the root structure.
	defaultFakeRootName = "root"
	// pathStructSuffix is the suffix to be appended to generated
	// PathStructs to distinguish them from the generated GoStructs, which
	// assume a similar name.
	pathStructSuffix = "Path"
	// GoDefaultProtoLibImportPath is the default path for the proto
	// library used in the generated code.
	GoDefaultProtoLibImportPath = "google.golang.org/protobuf/proto"
	// GoDefaultTelemGoImportPath is the default path for the telemgo library that
	// is used in the generated code.
	GoDefaultTelemGoImportPath = "github.com/openconfig/ondatra/telemgen/telemgo"
)

// NewDefaultConfig creates a GenConfig with default configuration.
func NewDefaultConfig() *GenConfig {
	return &GenConfig{
		PackageName: defaultPathPackageName,
		GoImports: GoImports{
			YgotImportPath:     genutil.GoDefaultYgotImportPath,
			YtypesImportPath:   genutil.GoDefaultYtypesImportPath,
			GoyangImportPath:   genutil.GoDefaultGoyangImportPath,
			ProtoLibImportPath: GoDefaultProtoLibImportPath,
			GNMIProtoPath:      genutil.GoDefaultGNMIImportPath,
			TelemGoImportPath:  GoDefaultTelemGoImportPath,
		},
		FakeRootName:     defaultFakeRootName,
		GeneratingBinary: genutil.CallerName(),
	}
}

// GenConfig stores code generation configuration.
type GenConfig struct {
	// PackageName is the name that should be used for the generating package.
	PackageName string
	// GoImports contains package import options.
	GoImports GoImports
	// FakeRootName specifies the name of the struct that should be generated
	// representing the root.
	FakeRootName string
	// ExcludeModules specifies any modules that are included within the set of
	// modules that should have code generated for them that should be ignored during
	// code generation. This is due to the fact that some schemas (e.g., OpenConfig
	// interfaces) currently result in overlapping entities (e.g., /interfaces).
	ExcludeModules []string
	// YANGParseOptions provides the options that should be handed to the
	// github.com/openconfig/goyang/pkg/yang library. These specify how the
	// input YANG files should be parsed.
	YANGParseOptions yang.Options
	// GeneratingBinary is the name of the binary calling the main
	// library, it is included in the header of output files for
	// debugging purposes. If a string is not specified, the location of
	// the library is utilised.
	GeneratingBinary string
	// ListBuilderKeyThreshold means to use the builder API format instead
	// of the key-combination API format for constructing list keys when
	// the number of keys is at least the threshold value.
	// 0 (default) means no threshold, i.e. always use the key-combination
	// API format.
	ListBuilderKeyThreshold uint
	// PreferShadowPath makes the generated code use the "shadow-path"
	// instead of the "path" struct tags when both are present while
	// processing a GoStruct for unmarshalling or marshalling.
	PreferShadowPath bool
}

// GoImports contains package import paths needed by telemgen.
// This is a superset of ypathgen's imports.
type GoImports struct {
	// SchemaStructPkgPath specifies the path to the ygen-generated structs, which
	// is used to get the enum and union type names used as the list key
	// for calling a list path accessor.
	SchemaStructPkgPath string
	// YgotImportPath specifies the path to the ygot library that should be used
	// in the generated code.
	YgotImportPath string
	// YtypesImportPath specifies the path to the ytypes library that should be used
	// in the generated code.
	YtypesImportPath string
	// GoyangImportPath specifies the path that should be used in the generated
	// code for importing the goyang/pkg/yang package.
	GoyangImportPath string
	// GNMIProtoPath specifies the path to the generated gNMI protobuf, which
	// is used when generating gNMI paths for unmarshalling a schema element.
	GNMIProtoPath string
	// ProtoLibImportPath is the import path to the proto package used by
	// the generated code.
	ProtoLibImportPath string
	// TelemGoImportPath is the import path to the telemgo package,
	// which contains helpers used by the generated code.
	TelemGoImportPath string
}

// GeneratedTelemetryCode contains generated code snippets that when written
// into files within a directory, forms the telemetry part of a telemetry path
// API implementation. The generated code is divided into the header and two
// types of telemetry struct & method objects. The header is further split into
// a common and a one-off header, and the telemetry objects is further split
// into snippets that are per-node, or per return type, as different nodes
// could have the same return types.
type GeneratedTelemetryCode struct {
	// CommonHeader is the header that should be used for all output Go files.
	CommonHeader string
	// OneOffHeader defines the header that should be included in only one output Go file.
	OneOffHeader string
	// GoPerNodeSnippets is the set of generated telemetry methods for all leaf path structs.
	GoPerNodeSnippets GoPerNodeCodeSnippets
	// GoReturnTypeCodeSnippets is the set of return types for all telemetry methods.
	GoReturnTypeSnippets GoReturnTypeCodeSnippets
}

func (c *GeneratedTelemetryCode) String() string {
	var b strings.Builder
	b.WriteString(c.CommonHeader)
	b.WriteString(c.OneOffHeader)
	b.WriteString(c.GoReturnTypeSnippets.String())
	b.WriteString(c.GoPerNodeSnippets.String())
	return b.String()
}

// GoPerNodeCodeSnippets is the set of generated telemetry methods for all path structs.
type GoPerNodeCodeSnippets []GoPerNodeCodeSnippet

func (ss GoPerNodeCodeSnippets) String() string {
	var b strings.Builder
	for _, s := range ss {
		genutil.WriteIfNotEmpty(&b, s.String())
	}
	return b.String()
}

// GoReturnTypeCodeSnippets is the set of return types for all telemetry methods.
type GoReturnTypeCodeSnippets []GoReturnTypeCodeSnippet

func (ss GoReturnTypeCodeSnippets) String() string {
	var b strings.Builder
	for _, s := range ss {
		genutil.WriteIfNotEmpty(&b, s.String())
	}
	return b.String()
}

// GoPerNodeCodeSnippet is the set of generated telemetry methods for a leaf
// path struct.
type GoPerNodeCodeSnippet struct {
	// PathStructName is the path struct name of the leaf.
	PathStructName string
	// GetMethod is the method snippet for the leaf's Get() method.
	GetMethod string
	// ReplaceMethod is the method snippet for the leaf's Replace() method.
	ReplaceMethod string
	// CollectMethod is the method snippet for the leaf's Collect() method.
	CollectMethod string
	// ConvertHelper is the function snippet corresponding to this leaf.
	ConvertHelper string
}

func (s *GoPerNodeCodeSnippet) String() string {
	var b strings.Builder
	for _, method := range []string{s.GetMethod, s.CollectMethod, s.ReplaceMethod, s.ConvertHelper} {
		genutil.WriteIfNotEmpty(&b, method)
	}
	return b.String()
}

// GoReturnTypeCodeSnippet is a set of return types for matching telemetry methods.
type GoReturnTypeCodeSnippet struct {
	// TypeName is the Go type name which these return types correspond to.
	TypeName string
	// QualifiedType is the value + timestamp combination of the retrieved update.
	QualifiedType string
	// CollectionType is returned by Collect() containing an Await() method
	// that in turn returns a slice of its corresponding QualifiedType
	// values.
	CollectionType string
}

func (s *GoReturnTypeCodeSnippet) String() string {
	var b strings.Builder
	for _, method := range []string{s.QualifiedType, s.CollectionType} {
		genutil.WriteIfNotEmpty(&b, method)
	}
	return b.String()
}

// GenerateTelemetryCode takes a slice of strings containing the path to a set
// of YANG files which contain YANG modules, and a second slice of strings
// which specifies the set of paths that are to be searched for associated
// models (e.g., modules that are included by the specified set of modules, or
// submodules of those modules). It extracts the set of modules that are to be
// generated, and returns ypathgen.GeneratedPathCode and GeneratedTelemetryCode
// struct pointers which can be written into a directory of files that together
// constitute a package for a telemetry path API.
//
// ypathgen.GeneratedPathCode contains these:
//      1. A path struct definition for each container, list, or leaf schema
//      node, as well as the fakeroot.
//      2. A Resolve() helper function, which can return the absolute path of
//      any struct.
//      3. Next-level methods for the fakeroot and each non-leaf schema node,
//      which instantiate and return the next-level structs corresponding to
//      their child schema nodes.
// GeneratedTelemetryCode in turn contains these:
//      1. Get(), Collect() telemetry methods for each path struct, apart from
//      the fake root.
//      2. Qualified<Type> and Collection<Type> structs for every return type
//      used by the Get() and Collect() methods for retrieving the telemetry
//      values of any node.
//      3. A convert() helper for each path struct used for unmarshalling a
//      leaf update to the desired ygen value.
// If errors are encountered during code generation, they are returned.
// TODO: Collect() for non-leaf nodes are currently unsupported (methods are not being generated).
func (cg *GenConfig) GenerateTelemetryCode(yangFiles, includePaths []string) (*GeneratedTelemetryCode, util.Errors) {
	// pcg is the configuration used to generate the path struct
	// portion of the generated telemetry code.
	pcg := &ypathgen.GenConfig{
		PackageName: cg.PackageName,
		GoImports: ypathgen.GoImports{
			SchemaStructPkgPath: cg.GoImports.SchemaStructPkgPath,
			YgotImportPath:      cg.GoImports.YgotImportPath,
		},
		FakeRootName:                         cg.FakeRootName,
		PathStructSuffix:                     pathStructSuffix,
		PreferOperationalState:               true,
		ShortenEnumLeafNames:                 true,
		EnumOrgPrefixesToTrim:                []string{"openconfig"},
		UseDefiningModuleForTypedefEnumNames: true,
		ExcludeModules:                       cg.ExcludeModules,
		YANGParseOptions:                     cg.YANGParseOptions,
		GeneratingBinary:                     cg.GeneratingBinary,
		ListBuilderKeyThreshold:              cg.ListBuilderKeyThreshold,
	}
	_, nodeDataMap, errs := pcg.GeneratePathCode(yangFiles, includePaths)
	if errs != nil {
		return nil, errs
	}

	genCode := &GeneratedTelemetryCode{}

	var err error
	if genCode.CommonHeader, genCode.OneOffHeader, err = cg.generateGoHeader(false, cg.PreferShadowPath); err != nil {
		return nil, util.AppendErr(errs, err)
	}

	// Determine whether there is an accessor for accessing generated
	// structs and telemetry types from a different package.
	var schemaStructPkgAccessor string
	if cg.GoImports.SchemaStructPkgPath != "" {
		schemaStructPkgAccessor = schemaStructPkgAlias + "."
	}
	fakeRootTypeName := cg.fakeRootTypeName(false)

	// Accumulate encountered types for per-type output later.
	goTypeSet := map[string]goTypeData{}

	// Per-node code generation; order by names for deterministic output.
	for _, pathStructName := range ypathgen.GetOrderedNodeDataNames(nodeDataMap) {
		perNodeSnippet, goType, es := generatePerNodeSnippet(pathStructName, nodeDataMap[pathStructName], fakeRootTypeName, schemaStructPkgAccessor, cg.PreferShadowPath)
		if es != nil {
			errs = util.AppendErrs(errs, es)
			continue
		}
		genCode.GoPerNodeSnippets = append(genCode.GoPerNodeSnippets, perNodeSnippet)
		goTypeSet[goType.GoTypeName] = goType
	}

	// Per-type code generation. This only makes sense to generate if we're
	// not already importing a set of schema structs from another package.
	if cg.GoImports.SchemaStructPkgPath == "" {
		for _, goType := range goTypeSet {
			returnTypeSnippet, es := generatePerTypeSnippet(goType, fakeRootTypeName, cg.PreferShadowPath)
			if es != nil {
				errs = util.AppendErrs(errs, es)
				continue
			}
			genCode.GoReturnTypeSnippets = append(genCode.GoReturnTypeSnippets, returnTypeSnippet)
		}
		// Order by names for deterministic output.
		sort.Slice(genCode.GoReturnTypeSnippets, func(i, j int) bool {
			return genCode.GoReturnTypeSnippets[i].TypeName < genCode.GoReturnTypeSnippets[j].TypeName
		})
	}

	return genCode, errs
}

// GenerateConfigCode takes a slice of strings containing the path to a set
// of YANG files which contain YANG modules, and a second slice of strings
// which specifies the set of paths that are to be searched for associated
// models (e.g., modules that are included by the specified set of modules, or
// submodules of those modules). It extracts the set of modules that are to be
// generated, and returns *GeneratedTelemetryCode which can be written into a
// directory of files that together constitute a package for a config path
// API.
// If errors are encountered during code generation, they are returned.
// TODO: Batch and Get functionalities are unimplemented.
func (cg *GenConfig) GenerateConfigCode(yangFiles, includePaths []string) (*GeneratedTelemetryCode, util.Errors) {
	// pcg is the configuration used to generate the path struct
	// portion of the generated telemetry code.
	pcg := &ypathgen.GenConfig{
		PackageName: cg.PackageName,
		GoImports: ypathgen.GoImports{
			SchemaStructPkgPath: cg.GoImports.SchemaStructPkgPath,
			YgotImportPath:      cg.GoImports.YgotImportPath,
		},
		FakeRootName:                         cg.FakeRootName,
		PathStructSuffix:                     pathStructSuffix,
		PreferOperationalState:               false,
		ExcludeState:                         true,
		ShortenEnumLeafNames:                 true,
		EnumOrgPrefixesToTrim:                []string{"openconfig"},
		UseDefiningModuleForTypedefEnumNames: true,
		ExcludeModules:                       cg.ExcludeModules,
		YANGParseOptions:                     cg.YANGParseOptions,
		GeneratingBinary:                     cg.GeneratingBinary,
	}
	_, nodeDataMap, errs := pcg.GeneratePathCode(yangFiles, includePaths)
	if errs != nil {
		return nil, errs
	}

	genCode := &GeneratedTelemetryCode{}

	var err error
	genCode.CommonHeader, genCode.OneOffHeader, err = cg.generateGoHeader(true, cg.PreferShadowPath)
	if err != nil {
		return nil, util.AppendErr(errs, err)
	}

	// Determine whether there is an accessor for accessing generated
	// structs and telemetry types from a different package.
	var schemaStructPkgAccessor string
	if cg.GoImports.SchemaStructPkgPath != "" {
		schemaStructPkgAccessor = schemaStructPkgAlias + "."
	}
	fakeRootTypeName := cg.fakeRootTypeName(false)

	// Accumulate encountered types for per-type output later.
	goTypeSet := map[string]goTypeData{}

	// Per-node code generation; order by names for deterministic output.
	for _, pathStructName := range ypathgen.GetOrderedNodeDataNames(nodeDataMap) {
		perNodeSnippet, goType, es := generatePerNodeConfigSnippet(pathStructName, nodeDataMap[pathStructName], fakeRootTypeName, schemaStructPkgAccessor, cg.PreferShadowPath)
		if es != nil {
			errs = util.AppendErrs(errs, es)
			continue
		}
		genCode.GoPerNodeSnippets = append(genCode.GoPerNodeSnippets, perNodeSnippet)
		goTypeSet[goType.GoTypeName] = goType
	}

	return genCode, errs
}

var (
	goCommonHeaderTemplate = mustTemplate("goCommonHeader", `
{{- /**/ -}}
package {{ .PackageName }}

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"

	"{{ .ProtoLibImportPath }}"
	"{{ .GoyangImportPath }}"
	"{{ .YgotImportPath }}"
	"{{ .YtypesImportPath }}"
	"{{ .TelemGoImportPath }}"
	{{- if .SchemaStructPkgPath }}
	{{ .SchemaStructPkgAlias }} "{{ .SchemaStructPkgPath }}"
	{{- end }}

	gpb "{{ .GNMIProtoPath }}"
)
`)

	// goOneOffHeaderTemplate contains code that should live only in one of the
	// generated telemetry files (which all live within one package).
	//
	// binarySliceToFloat32, although ideally should be part of the ygot
	// library, lives here due to "[]Binary", a type defined in the
	// ygen-generated code, not being able to be assigned to [][]byte,
	// preventing such a function from being used.
	goOneOffHeaderTemplate = mustTemplate("goOneOffHeader", `
// WithReplica adds the replica number to the context metadata of the gNMI
// server query.
func (n *{{ .FakeRootTypePathName }}) WithReplica(replica int) *{{ .FakeRootTypePathName }} {
	telemgo.PutReplica(n, replica)
	return n
}

// WithSubscriptionMode specifies the subscription mode in the underlying gNMI
// subscribe.
func (n *{{ .FakeRootTypePathName }}) WithSubscriptionMode(mode gpb.SubscriptionMode) *{{ .FakeRootTypePathName }} {
	telemgo.PutSubscriptionMode(n, mode)
	return n
}

// WithClient allows the user to provide a gNMI client. This allows for creation
// of tests for multiple gNMI clients to a single DUT.
func (n *{{ .FakeRootTypePathName }}) WithClient(c gpb.GNMIClient) *{{ .FakeRootTypePathName }} {
	telemgo.PutClient(n, c)
	return n
}

{{- if .GenerateConfigCode }}

// WithReplica adds the replica number to the context metadata of the gNMI
// server query.
func (b *Batch) WithReplica(replica int) *Batch {
	b.deviceRoot.WithReplica(replica)
	return b
}

// WithSubscriptionMode specifies the subscription mode in the underlying gNMI
// subscribe.
func (b *Batch) WithSubscriptionMode(mode gpb.SubscriptionMode) *Batch {
	b.deviceRoot.WithSubscriptionMode(mode)
	return b
}

// Batch is the root of a duplicated path tree that has batch operations instead of synchronous config operations that DevicePath has.
type Batch struct {
 	replaceReqs []telemgo.PathValuePair
 	origin string
 	deviceRoot *{{ .FakeRootTypePathName }}
}

// NewBatch returns a newly instantiated Batch object for batching set requests.
func (d *DevicePath) NewBatch() *Batch {
	return &Batch{
		deviceRoot: &{{ .FakeRootTypePathName }}{ygot.New{{ .FakeRootBaseTypeName }}(d.Id())},
		origin: "openconfig",
	}
}

func (b *Batch) batchReplace(t testing.TB, n ygot.PathStruct, val interface{}) {
	t.Helper()
	path, customData, errs := ygot.ResolvePath(n)
	if len(errs) > 0 {
		t.Fatalf("Errors resolving path %v: %v", n, errs)
	}
	if path.GetTarget() != b.deviceRoot.Id() {
		t.Fatalf("path target doesn't equal the device target for the batch: got %q, want %q", path.GetTarget(), b.deviceRoot.Id())
	}
	if len(customData) != 0 {
		t.Fatalf("batching cannot accept a path that has its custom request options; please set request options solely via the batch object.")
	}
	b.replaceReqs = append(b.replaceReqs, telemgo.PathValuePair{PathElems: path.GetElem(), Val: val})
}

// Set commits all the batched requests.
func (b *Batch) Set(t testing.TB) *gpb.SetResponse {
	return telemgo.BatchReplace(t, b.origin, b.deviceRoot.Id(), b.deviceRoot.CustomData(), b.replaceReqs)
}

{{- end }}

func binarySliceToFloat32(in []{{ .SchemaStructPkgAccessor }}Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// getFull uses gNMI Get to fill the input GoStruct with values at the input path.
func getFull(t testing.TB, n ygot.PathStruct, goStructName string, gs ygot.GoStruct, isLeaf bool) *telemgo.QualifiedType {
	datapoints, queryPath := telemgo.Get(t, n, isLeaf)
	qv, err := telemgo.Unmarshal(t, datapoints, getSchema(), goStructName, gs, queryPath, isLeaf, {{ .PreferShadowPath }})
	if err != nil {
		t.Fatal(err)
	}
	return qv
}

// getSchema return the generated ytypes schema used for unmarshaling datapoints.
func getSchema() *ytypes.Schema {
	return &ytypes.Schema{
		Root:       &{{ .FakeRootTypeName }}{},
		SchemaTree: {{ .SchemaStructPkgAccessor }}SchemaTree,
		Unmarshal:  {{ .SchemaStructPkgAccessor }}Unmarshal,
	}
}
`)

	// goLeafConvertTemplate contains the per-leaf helper that retrieves the leaf
	// value from its containing parent GoStruct to complete the unmarshalling process.
	goLeafConvertTemplate = mustTemplate("goLeafConvertHelper", `{{ $QualifiedGoTypeName := printf "%sQualified%s" .SchemaStructPkgAccessor .GoType.TransformedGoTypeName }}
// convert{{ .PathStructName }} extracts the value of the leaf {{ .GoFieldName }} from its parent {{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}
// and combines the update with an existing QualifiedType to return a *{{ $QualifiedGoTypeName }}.
func convert{{ .PathStructName }}(t testing.TB, qt *telemgo.QualifiedType, parent *{{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}) *{{ $QualifiedGoTypeName }} {
	t.Helper()
	if qt.ComplianceErrors != nil {
		t.Fatal(qt.ComplianceErrors)
	}
	qv := &{{ $QualifiedGoTypeName }}{
		QualifiedType: qt,
	}
	val := parent.{{ .GoFieldName }}
	if !reflect.ValueOf(val).IsZero() {
		qv.Present = true
		{{- if .SpecialConversionFn }}
		qv.SetVal({{ .SpecialConversionFn }}({{ if .IsScalarField -}} * {{- end -}} val))
		{{- else }}
		qv.SetVal({{ if .IsScalarField -}} * {{- end -}} val)
		{{- end }}
	} else {
		qv.Present = false
	}
	return qv
}
`)

	// goNodeGetTemplate contains the per-node helper that makes the Get telemetry
	// call, unmarshals the payload, and returns the result to the user.
	goNodeGetTemplate = mustTemplate("goLeafGetMethod", `{{ $QualifiedGoTypeName := printf "%sQualified%s" .SchemaStructPkgAccessor .GoType.TransformedGoTypeName }}
// GetFull retrieves a sample for {{ .YANGPath }}.
func (n *{{ .PathStructName }}) GetFull(t testing.TB) *{{ $QualifiedGoTypeName }} {
	t.Helper()
	goStruct := &{{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}{}
	ret := &{{ $QualifiedGoTypeName }}{
		QualifiedType: getFull(t, n, "{{ .GoStructTypeName }}", goStruct, {{ .GoType.IsLeaf }}),
	}

	{{- if .GoType.HasDefault }}
	if !ret.IsPresent() {
		ret.SetVal(goStruct.Get{{ .GoFieldName }}())
		ret.Present = true
		return ret
	}
	{{- end }}

	{{- if .GoType.IsLeaf }}
	return convert{{ .PathStructName }}(t, ret.QualifiedType, goStruct)
	{{- else }}
	if ret.IsPresent() {
		ret.SetVal(goStruct)
	}
	return ret
	{{- end }}
}

// Get retrieves a value sample for {{ .YANGPath }}, erroring out if it is not present.
func (n *{{ .PathStructName }}) Get(t testing.TB) {{ .GoType.GoTypeName }} {
	t.Helper()
	return n.GetFull(t).Val(t)
}

{{- if not .IsRoot }}

// GetFull retrieves a list of samples for {{ .YANGPath }}.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) GetFull(t testing.TB) []*{{ $QualifiedGoTypeName }} {
	t.Helper()
	datapoints, queryPath := telemgo.Get(t, n, false)
	datapointGroups, sortedPrefixes := telemgo.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)), {{- if .GoType.IsLeaf }} true {{- else }} false {{- end }})

	var data []*{{ $QualifiedGoTypeName }}
	for _, prefix := range sortedPrefixes {
		goStruct := &{{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}{}
		qt, err := telemgo.Unmarshal(t, datapointGroups[prefix], getSchema(), "{{ .GoStructTypeName }}", goStruct, queryPath, {{ .GoType.IsLeaf }}, {{ .PreferShadowPath }})
		if err != nil {
			t.Fatal(err)
		}

		{{- if .GoType.IsLeaf }}
		qv := convert{{ .PathStructName }}(t, qt, goStruct)
		{{- else }}
		if !qt.IsPresent() {
			continue
		}
		qv := (&{{ $QualifiedGoTypeName }}{
			QualifiedType: qt,
		}).SetVal(goStruct)
		{{- end }}
		data = append(data, qv)
	}
	return data
}

// Get retrieves a list of value samples for {{ .YANGPath }}.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) Get(t testing.TB) []{{ .GoType.GoTypeName }} {
	t.Helper()
	fulldata := n.GetFull(t)
	var data []{{ .GoType.GoTypeName }}
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

{{- end }}
`)

	// goLeafCollectTemplate contains the per-leaf Collect telemetry call. It
	// triggers the asynchronous data collection at this node, and returns a type
	// containing an Await method that blocks for the data collection to complete,
	// as well as unmarshalling that data for the user.
	goLeafCollectTemplate = mustTemplate("goLeafCollectMethod", `{{ $QualifiedGoTypeName := printf "%sQualified%s" .SchemaStructPkgAccessor .GoType.TransformedGoTypeName }}
// Collect retrieves a Collection sample for {{ .YANGPath }}.
func (n *{{ .PathStructName }}) Collect(t testing.TB, duration time.Duration) *{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }} {
	t.Helper()
	return &{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }}{
		c: n.CollectUntil(t, duration, func(*{{ $QualifiedGoTypeName }}) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for {{ .YANGPath }} and evaluates the predicate on all samples.
func (n *{{ .PathStructName }}) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *{{ $QualifiedGoTypeName }}) bool) *{{ .SchemaStructPkgAccessor }}CollectionUntil{{ .GoType.TransformedGoTypeName }} {
	t.Helper()
	return &{{ .SchemaStructPkgAccessor }}CollectionUntil{{ .GoType.TransformedGoTypeName }}{
		c: telemgo.CollectUntil(t, n, duration, func(upd *telemgo.DataPoint) (telemgo.QualifiedValue, error) {
			parentPtr := &{{ .GoStructTypeName }}{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := telemgo.Unmarshal(t, []*telemgo.DataPoint{upd}, getSchema(), "{{ .GoStructTypeName }}", parentPtr, nil, {{ .GoType.IsLeaf }}, {{ .PreferShadowPath }})
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convert{{ .PathStructName }}(t, qv, parentPtr), nil
		},
		func(qualVal telemgo.QualifiedValue) bool {
			val, ok := qualVal.(*{{ $QualifiedGoTypeName }})
			return ok && predicate(val)
		}),
	}
}

// Await waits until {{ .YANGPath }} is deep-equal to the val and returns all received values.
// If the timeout is exceeded, the test fails fatally.
// To avoid a fatal failure or wait for a generic predicate, use CollectUntil.
func (n *{{ .PathStructName }}) Await(t testing.TB, duration time.Duration, val {{ .GoType.GoTypeName }}) []*{{ $QualifiedGoTypeName }} {
	t.Helper()
	vals, success := n.CollectUntil(t, duration, func(data *{{ $QualifiedGoTypeName }}) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		if len(vals) == 0 {
			t.Fatalf("Await() at {{ .YANGPath }} failed: no values received")
		}
		t.Fatalf("Await() at {{ .YANGPath }} failed: want %v, last got %v", val, vals[len(vals) - 1])
	}
	return vals
}

// Collect retrieves a Collection sample for {{ .YANGPath }}.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) Collect(t testing.TB, duration time.Duration) *{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }} {
	t.Helper()
	return &{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }}{
		c: n.CollectUntil(t, duration, func(*{{ $QualifiedGoTypeName }}) bool { return false }),
	}
}

// CollectUntil retrieves a Collection sample for {{ .YANGPath }} and evaluates the predicate on all samples.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) CollectUntil(t testing.TB, duration time.Duration, predicate func(val *{{ $QualifiedGoTypeName }}) bool) *{{ .SchemaStructPkgAccessor }}CollectionUntil{{ .GoType.TransformedGoTypeName }} {
	t.Helper()
	return &{{ .SchemaStructPkgAccessor }}CollectionUntil{{ .GoType.TransformedGoTypeName }}{
		c: telemgo.CollectUntil(t, n, duration, func(upd *telemgo.DataPoint) (telemgo.QualifiedValue, error) {
			parentPtr := &{{ .GoStructTypeName }}{}
			// queryPath is not needed on leaves because full gNMI path is always returned.
			qv, err := telemgo.Unmarshal(t, []*telemgo.DataPoint{upd}, getSchema(), "{{ .GoStructTypeName }}", parentPtr, nil, {{ .GoType.IsLeaf }}, {{ .PreferShadowPath }})
			if err != nil || qv.ComplianceErrors != nil {
				return nil, fmt.Errorf("unmarshal err: %v, complianceErrs: %v", err, qv.ComplianceErrors)
			}
			return convert{{ .PathStructName }}(t, qv, parentPtr), nil
		},
		func(qualVal telemgo.QualifiedValue) bool {
			val, ok := qualVal.(*{{ $QualifiedGoTypeName }})
			return ok && predicate(val)
		}),
	}
}
`)

	// goNodeReplaceTemplate is used to generate the Replace configuration call.
	goNodeReplaceTemplate = mustTemplate("goNodeReplaceMethod", `
// Replace replaces the configuration at {{ .YANGPath }}.
func (n *{{ .PathStructName }}) Replace(t testing.TB, val {{ .GoType.GoTypeName }}) *gpb.SetResponse {
	t.Helper()
	return telemgo.Replace(t, n, {{ if .IsScalarField -}} & {{- end -}} val)
}

func (n *{{ .PathStructName }}) BatchReplace(t testing.TB, b *Batch, val {{ .GoType.GoTypeName }}) {
	b.batchReplace(t, n, {{ if .IsScalarField -}} & {{- end -}} val)
}
`)

	// goQualifiedTypeTemplate contains the per-leaf return type used to store the
	// telemetry results. A pointer to a struct is intended as the return type to
	// allow a non-existent node to be expressed as a nil pointer.
	goQualifiedTypeTemplate = mustTemplate("goQualifiedType", `{{ $QualifiedGoTypeName := printf "Qualified%s" .TransformedGoTypeName }}
// {{ $QualifiedGoTypeName }} is a {{ .GoTypeName }} with a corresponding timestamp.
type {{ $QualifiedGoTypeName }} struct {
	*telemgo.QualifiedType
	val {{ .GoTypeName }} // val is the sample value.
}

func (q *{{ $QualifiedGoTypeName }}) String() string {
	return telemgo.QualifiedTypeString(q.val, q.QualifiedType)
}

// Val returns the value of the {{ .GoTypeName }} sample, erroring out if not present.
func (q *{{ $QualifiedGoTypeName }}) Val(t testing.TB) {{ .GoTypeName }} {
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

// SetVal sets the value of the {{ .GoTypeName }} sample.
func (q *{{ $QualifiedGoTypeName }}) SetVal(v {{ .GoTypeName }}) *{{ $QualifiedGoTypeName }} {
	q.val = v
	return q
}

`)

	// goCollectionTypeTemplate contains a per-type definition of
	// intermediate information stored for a telemetry collection. Its
	// creation signals the beginning of an asynchronous data collection of
	// a particular node; this node has the type associated with this
	// collection type. The intermediate information is required to
	// disambiguate this node from other nodes having the same type, in
	// order to correctly carry out the unmarshal operation once the
	// user triggers the Await() call that blocks for the data retrieval to
	// complete.
	// TODO: DataPoints returned from Await are not guaranteed to be in
	// order. First decide whether to handle non-leaves in the near future. If not,
	// incorporate it in the template. If yes, beware that the subtree has to
	// be deep-copied for each new sample.
	goCollectionTypeTemplate = mustTemplate("goCollectionType", `{{ $QualifiedGoTypeName := printf "Qualified%s" .TransformedGoTypeName }}
// Collection{{ .TransformedGoTypeName }} is a telemetry Collection whose Await method returns a slice of {{ .GoTypeName }} samples.
type Collection{{ .TransformedGoTypeName }} struct {
	c *CollectionUntil{{ .TransformedGoTypeName }}
}

// Await blocks for the telemetry collection to be complete, and then returns the slice of samples received.
func (u *Collection{{ .TransformedGoTypeName }}) Await(t testing.TB) []*{{ $QualifiedGoTypeName }} {
	t.Helper()
	data, _ := u.c.Await(t)
	return data
}

// CollectionUntil{{ .TransformedGoTypeName }} is a telemetry Collection whose Await method returns a slice of {{ .GoTypeName }} samples.
type CollectionUntil{{ .TransformedGoTypeName }} struct {
	c *telemgo.Collection
}

// Await blocks for the telemetry collection to be complete or the predicate to be true whichever is first.
// The received data and the status of the predicate are returned.
func (u *CollectionUntil{{ .TransformedGoTypeName }}) Await(t testing.TB) ([]*{{ $QualifiedGoTypeName }}, bool) {
	t.Helper()
	var ret []*{{ $QualifiedGoTypeName }}
	updates, predTrue := u.c.Await(t)
	for _, upd := range updates {
		ret = append(ret, upd.(*{{ $QualifiedGoTypeName }}))
	}
	return ret, predTrue
}
`)
)

// mustTemplate generates a template.Template for a string template.
// This should only be used during initialization as it panics if unsuccessful.
func mustTemplate(name, src string) *template.Template {
	return template.Must(template.New(name).Parse(src))
}

// fakeRootTypeName returns the name of either the GoStruct or PathStruct fake
// root type name.
// If schemaStructPkgAccessor is provided, then it will be prefixed in front of
// the GoStruct fake root type name.
func (cg *GenConfig) fakeRootTypeName(pathStruct bool) string {
	goStructFakeRootName := yang.CamelCase(cg.FakeRootName)
	switch {
	case pathStruct:
		return goStructFakeRootName + pathStructSuffix
	case cg.GoImports.SchemaStructPkgPath != "":
		return schemaStructPkgAlias + "." + goStructFakeRootName
	default:
		return goStructFakeRootName
	}
}

// generateGoHeader returns the common header, one-off header of the file in
// that order, and error if one is encountered.
func (cg *GenConfig) generateGoHeader(generateConfigCode bool, preferShadowPath bool) (string, string, error) {
	var schemaStructPkgAccessor string
	if cg.GoImports.SchemaStructPkgPath != "" {
		schemaStructPkgAccessor = schemaStructPkgAlias + "."
	}
	var commonHeader strings.Builder
	s := struct {
		// GoImports contains package import paths needed by telemgen.
		GoImports
		// PackgeName is the name of the package to be generated.
		PackageName string
		// FakeRootTypePathName is the type name with path suffix of the fakeroot node in the generated code.
		FakeRootTypePathName string
		// FakeRootTypeName is the type name of the fakeroot node in the generated code.
		FakeRootTypeName string
		// FakeRootBaseTypeName is the type name of the fake root struct which
		// should be embedded within the fake root path struct.
		FakeRootBaseTypeName string
		// SchemaStructPkgAlias is the package alias for the imported ygen-generated file.
		SchemaStructPkgAlias string
		// SchemaStructPkgAccessor is the accessor prefix for types in
		// the imported ygen-generated file (if any).
		SchemaStructPkgAccessor string
		// GenerateConfigCode is a flag for whether config functions are being
		// generated instead of telemetry functions.
		GenerateConfigCode bool

		PreferShadowPath bool
	}{
		GoImports:               cg.GoImports,
		PackageName:             cg.PackageName,
		FakeRootTypePathName:    cg.fakeRootTypeName(true),
		FakeRootTypeName:        cg.fakeRootTypeName(false),
		FakeRootBaseTypeName:    ygot.FakeRootBaseTypeName,
		SchemaStructPkgAlias:    schemaStructPkgAlias,
		SchemaStructPkgAccessor: schemaStructPkgAccessor,
		GenerateConfigCode:      generateConfigCode,
		PreferShadowPath:        preferShadowPath,
	}
	if err := goCommonHeaderTemplate.Execute(&commonHeader, s); err != nil {
		return "", "", err
	}

	var oneOffHeader strings.Builder
	if err := goOneOffHeaderTemplate.Execute(&oneOffHeader, s); err != nil {
		return "", "", err
	}

	return commonHeader.String(), oneOffHeader.String(), nil
}

// goTypeData stores processed information of a node's type.
type goTypeData struct {
	// GoTypeName is the type name to be used to store the telemetry result of a single sample.
	GoTypeName string
	// TransformedGoTypeName is a Go-safe identifier name for a single transformed version of GoTypeName.
	TransformedGoTypeName string
	// IsLeaf indicates whether the type is a leaf type in the schema.
	IsLeaf bool
	// HasDefault indicates whether the leaf has a default value specified in its YANG definition.
	HasDefault bool
}

// generatePerNodeSnippet generates the telemetry code for a path node.
func generatePerNodeSnippet(pathStructName string, nodeData *ypathgen.NodeData, fakeRootTypeName, schemaStructPkgAccessor string, preferShadowPath bool) (GoPerNodeCodeSnippet, goTypeData, util.Errors) {
	// Special case: ieeefloat32 is represented as a 4-byte binary in YANG
	// and ygen, but float32 is more user-friendly.
	var specialConversionFn string
	if nodeData.YANGTypeName == "ieeefloat32" {
		switch nodeData.LocalGoTypeName {
		case "Binary":
			nodeData.GoTypeName = "float32"
			nodeData.LocalGoTypeName = "float32"
			specialConversionFn = "ygot.BinaryToFloat32"
		case "[]" + "Binary":
			nodeData.GoTypeName = "[]float32"
			nodeData.LocalGoTypeName = "[]float32"
			specialConversionFn = "binarySliceToFloat32"
		default:
			return GoPerNodeCodeSnippet{}, goTypeData{}, util.NewErrs(
				errors.Errorf("ieeefloat32 is expected to be a binary, got %q", nodeData.LocalGoTypeName))
		}
	}

	var errs util.Errors
	s := struct {
		PathStructName   string
		GoType           goTypeData
		GoFieldName      string
		GoStructTypeName string
		YANGPath         string
		FakeRootTypeName string
		// IsScalarField indicates a leaf that is stored as a pointer
		// in its parent struct.
		IsScalarField           bool
		IsRoot                  bool
		SchemaStructPkgAccessor string
		// WildcardSuffix is the suffix used to indicate that a path
		// node contains a wildcard.
		WildcardSuffix string
		// SpecialConversionFn is the special-case conversion function
		// to convert the field from the parent struct into the
		// qualified type returned to the user.
		SpecialConversionFn string
		PreferShadowPath    bool
	}{
		PathStructName: pathStructName,
		GoType: goTypeData{
			GoTypeName:            nodeData.GoTypeName,
			TransformedGoTypeName: transformGoTypeName(nodeData),
			IsLeaf:                nodeData.IsLeaf,
			HasDefault:            nodeData.HasDefault,
		},
		GoFieldName:             nodeData.GoFieldName,
		GoStructTypeName:        nodeData.SubsumingGoStructName,
		YANGPath:                nodeData.YANGPath,
		FakeRootTypeName:        fakeRootTypeName,
		IsScalarField:           nodeData.IsScalarField,
		IsRoot:                  nodeData.YANGPath == "/",
		WildcardSuffix:          ypathgen.WildcardSuffix,
		SpecialConversionFn:     specialConversionFn,
		SchemaStructPkgAccessor: schemaStructPkgAccessor,
		PreferShadowPath:        preferShadowPath,
	}
	var getMethod, collectMethod, convertHelper strings.Builder
	if nodeData.IsLeaf {
		// Leaf types use their parent GoStruct to unmarshal, before
		// being retrieved out when returned to the user.
		if err := goLeafConvertTemplate.Execute(&convertHelper, s); err != nil {
			util.AppendErr(errs, err)
		}
		// TODO: Collect methods for non-leaf nodes is not implemented.
		if err := goLeafCollectTemplate.Execute(&collectMethod, s); err != nil {
			util.AppendErr(errs, err)
		}
	}
	if err := goNodeGetTemplate.Execute(&getMethod, s); err != nil {
		util.AppendErr(errs, err)
	}

	return GoPerNodeCodeSnippet{
		PathStructName: pathStructName,
		GetMethod:      getMethod.String(),
		CollectMethod:  collectMethod.String(),
		ConvertHelper:  convertHelper.String(),
	}, s.GoType, errs
}

// generatePerNodeConfigSnippet generates the config code for a path node.
func generatePerNodeConfigSnippet(pathStructName string, nodeData *ypathgen.NodeData, fakeRootTypeName, schemaStructPkgAccessor string, preferShadowPath bool) (GoPerNodeCodeSnippet, goTypeData, util.Errors) {
	// TODO: See if a float32 -> binary helper should be provided
	// for setting a float32 leaf.
	var errs util.Errors
	s := struct {
		PathStructName          string
		GoType                  goTypeData
		GoFieldName             string
		GoStructTypeName        string
		YANGPath                string
		FakeRootTypeName        string
		IsScalarField           bool
		IsRoot                  bool
		SchemaStructPkgAccessor string
		WildcardSuffix          string
		SpecialConversionFn     string
		PreferShadowPath        bool
	}{
		PathStructName: pathStructName,
		GoType: goTypeData{
			GoTypeName:            nodeData.GoTypeName,
			TransformedGoTypeName: transformGoTypeName(nodeData),
			IsLeaf:                nodeData.IsLeaf,
			HasDefault:            nodeData.HasDefault,
		},
		GoFieldName:             nodeData.GoFieldName,
		GoStructTypeName:        nodeData.SubsumingGoStructName,
		YANGPath:                nodeData.YANGPath,
		FakeRootTypeName:        fakeRootTypeName,
		IsScalarField:           nodeData.IsScalarField,
		IsRoot:                  nodeData.YANGPath == "/",
		WildcardSuffix:          ypathgen.WildcardSuffix,
		SchemaStructPkgAccessor: schemaStructPkgAccessor,
		PreferShadowPath:        preferShadowPath,
	}
	var getMethod, replaceMethod, convertHelper strings.Builder
	if nodeData.IsLeaf {
		// Leaf types use their parent GoStruct to unmarshal, before
		// being retrieved out when returned to the user.
		if err := goLeafConvertTemplate.Execute(&convertHelper, s); err != nil {
			util.AppendErr(errs, err)
		}
	}
	if err := goNodeReplaceTemplate.Execute(&replaceMethod, s); err != nil {
		util.AppendErr(errs, err)
	}
	if err := goNodeGetTemplate.Execute(&getMethod, s); err != nil {
		util.AppendErr(errs, err)
	}

	return GoPerNodeCodeSnippet{
		PathStructName: pathStructName,
		GetMethod:      getMethod.String(),
		ConvertHelper:  convertHelper.String(),
		ReplaceMethod:  replaceMethod.String(),
	}, s.GoType, errs
}

// generatePerTypeSnippet generates the telemetry code for a return type.
func generatePerTypeSnippet(goType goTypeData, fakeRootTypeName string, preferShadowPath bool) (GoReturnTypeCodeSnippet, util.Errors) {
	var errs util.Errors
	var qualifiedType, collectionType strings.Builder
	if err := goQualifiedTypeTemplate.Execute(&qualifiedType, goType); err != nil {
		util.AppendErr(errs, err)
	}
	if goType.IsLeaf {
		// TODO: Collect methods for non-leaf nodes is not implemented.
		if err := goCollectionTypeTemplate.Execute(&collectionType, struct {
			goTypeData       // Embedded
			FakeRootTypeName string
			PreferShadowPath bool
		}{
			goTypeData:       goType,
			FakeRootTypeName: fakeRootTypeName,
			PreferShadowPath: preferShadowPath,
		}); err != nil {
			util.AppendErr(errs, err)
		}
	}
	return GoReturnTypeCodeSnippet{
		TypeName:       goType.GoTypeName,
		QualifiedType:  qualifiedType.String(),
		CollectionType: collectionType.String(),
	}, errs
}

// transformGoTypeName converts any unaccepted symbols (e.g. [], *) in the type
// name of a telemetry result into symbols acceptable as a unique Go
// identifier. Further, it capitalizes the first letter so that it is exported.
func transformGoTypeName(nodeData *ypathgen.NodeData) string {
	// TODO: map[] is an output for a non-leaf wildcard
	// GetFull()/Collect(), there hasn't been the need for non-leaves so
	// we haven't decided how to deal with those yet.

	// As maps, lists are never given defined types by ygen, so the input
	// type is either a pointer to a container struct (*structPkg.), or a leaf
	// type that's either a single value (structPkg.) or a slice ([]).
	s := strings.NewReplacer("*", "").Replace(nodeData.LocalGoTypeName)

	const slicePf = "[]"
	if strings.HasPrefix(s, slicePf) {
		s = strings.TrimPrefix(s, slicePf) + "Slice"
	}
	return strings.ToUpper(s[0:1]) + s[1:]
}
