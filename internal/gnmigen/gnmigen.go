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

// Package gnmigen is a library for generating an API for creating gNMI paths
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
package gnmigen

import (
	"regexp"
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
	// defaultConfigPackageName specifies the default name that should be
	// used for the generated Go package.
	defaultConfigPackageName = "config"
	// schemaStructPkgAlias is the package alias of the schema struct
	// package when the struct package is to be generated in a
	// separate package.
	schemaStructPkgAlias = "oc"
	// configPkgAlias is the package alias of the config
	// package when the package is to be generated in a
	// separate package.
	configPkgAlias = "config"
	// defaultFakeRootName is the default name for the root structure.
	defaultFakeRootName = "root"
	// pathStructSuffix is the suffix to be appended to generated
	// PathStructs to distinguish them from the generated GoStructs, which
	// assume a similar name.
	pathStructSuffix = "Path"
	// defaultGenUtilImportPath is the default path for the genutil library that
	// is used in the generated code.
	defaultGenUtilImportPath = "github.com/openconfig/ondatra/internal/gnmigen/genutil/genutil"
)

// NewDefaultConfig creates a GenConfig with default configuration.
func NewDefaultConfig() *GenConfig {
	return &GenConfig{
		PackageName: defaultPathPackageName,
		GoImports: GoImports{
			YgotImportPath:    genutil.GoDefaultYgotImportPath,
			YtypesImportPath:  genutil.GoDefaultYtypesImportPath,
			GNMIProtoPath:     genutil.GoDefaultGNMIImportPath,
			GenUtilImportPath: defaultGenUtilImportPath,
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
	// SplitByModule controls whether path struct helpers should be split
	// into different Go packages.
	SplitByModule bool
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
	// GNMIProtoPath specifies the path to the generated gNMI protobuf, which
	// is used when generating gNMI paths for unmarshalling a schema element.
	GNMIProtoPath string
	// GenUtilImportPath is the import path to the genutil package,
	// which contains helpers used by the generated code.
	GenUtilImportPath string
	// ConfigImportPath is the import path of the of config batch objects.
	ConfigImportPath string
}

// GeneratedTelemetryCode contains generated code snippets that when written
// into files within a directory, forms the telemetry part of a telemetry path
// API implementation.
type GeneratedTelemetryCode struct {
	// Headers is a map from package name to the header that should be used for the output Go files.
	Headers map[string]string
	// FakeRootMethods defines the helper code that must be included in the package containing the fake root.
	// This includes additional funcs to set custom options on the fake root.
	FakeRootMethods string
	// GoPerNodeSnippets is a map from package name to the set of generated telemetry methods
	// for each top-level path struct below the fake root.
	GoPerNodeSnippets map[string]GoPerNodeCodeSnippets
	// GoReturnTypeCodeSnippets is the set of return types for all telemetry methods.
	GoReturnTypeSnippets GoReturnTypeCodeSnippets
}

func (c *GeneratedTelemetryCode) String(pkg string) string {
	var b strings.Builder
	b.WriteString(c.Headers[pkg])
	b.WriteString(c.FakeRootMethods)
	b.WriteString(c.GoReturnTypeSnippets.String())
	b.WriteString(c.GoPerNodeSnippets[pkg].String())
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

var pkgReplaceExp = regexp.MustCompile("[._-]")

// GenerateCode takes a slice of strings containing the path to a set
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
func (cg *GenConfig) GenerateCode(yangFiles, includePaths []string, config bool) (*GeneratedTelemetryCode, util.Errors) {
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
		TrimOCPackage:                        true,
		SplitByModule:                        cg.SplitByModule,
	}
	if config {
		pcg.PreferOperationalState = false
		pcg.ExcludeState = true
	}
	_, nodeDataMap, errs := pcg.GeneratePathCode(yangFiles, includePaths)
	if errs != nil {
		return nil, errs
	}

	genCode := &GeneratedTelemetryCode{
		GoPerNodeSnippets: make(map[string]GoPerNodeCodeSnippets),
		Headers:           make(map[string]string),
	}

	var err error
	genCode.FakeRootMethods, err = cg.executeTemplate(cg.PackageName, config, goFakeRootMethods)
	if err != nil {
		return nil, util.AppendErr(errs, err)
	}

	// Determine whether there is an accessor for accessing generated
	// structs and telemetry types from a different package.
	var schemaStructPkgAccessor string
	if cg.GoImports.SchemaStructPkgPath != "" {
		schemaStructPkgAccessor = schemaStructPkgAlias + "."
	}
	var configPkgAccessor string
	if cg.GoImports.ConfigImportPath != "" {
		configPkgAccessor = configPkgAlias + "."
	}
	fakeRootTypeName := cg.fakeRootTypeName(false)

	// Accumulate encountered types for per-type output later.
	goTypeSet := map[string]goTypeData{}

	// Per-node code generation; order by names for deterministic output.
	for _, pathStructName := range ypathgen.GetOrderedNodeDataNames(nodeDataMap) {
		pkgName := nodeDataMap[pathStructName].GoPathPackageName
		var perNodeSnippet GoPerNodeCodeSnippet
		var goType goTypeData
		var es util.Errors

		if config {
			perNodeSnippet, goType, es = generatePerNodeConfigSnippet(pathStructName, nodeDataMap[pathStructName], fakeRootTypeName, schemaStructPkgAccessor, configPkgAccessor, cg.PreferShadowPath)
		} else {
			perNodeSnippet, goType, es = generatePerNodeSnippet(pathStructName, nodeDataMap[pathStructName], fakeRootTypeName, schemaStructPkgAccessor, cg.PreferShadowPath)
		}
		if es != nil {
			errs = util.AppendErrs(errs, es)
			continue
		}
		genCode.GoPerNodeSnippets[pkgName] = append(genCode.GoPerNodeSnippets[pkgName], perNodeSnippet)
		goTypeSet[goType.GoTypeName] = goType
	}
	for pkg := range genCode.GoPerNodeSnippets {
		genCode.Headers[pkg], err = cg.executeTemplate(pkg, config, goCommonHeaderTemplate)
		if err != nil {
			return nil, util.AppendErr(errs, err)
		}
	}

	// Per-type code generation. This only makes sense to generate if we're
	// not already importing a set of schema structs from another package.
	if cg.GoImports.SchemaStructPkgPath == "" && !config {
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

	"{{ .YgotImportPath }}"
	"{{ .YtypesImportPath }}"
	"{{ .GenUtilImportPath }}"
	{{- if .SchemaStructPkgPath }}
	{{ .SchemaStructPkgAlias }} "{{ .SchemaStructPkgPath }}"
	{{- end }}
	{{- if .ConfigImportPath }}
	{{ .ConfigPkgAlias }} "{{ .ConfigImportPath }}"
	{{- end }}

	gpb "{{ .GNMIProtoPath }}"
)
`)

	// goFakeRootMethods contains code that should live only in the
	// the package that contains the fake root path struct.
	goFakeRootMethods = mustTemplate("goFakeRootMethods", `
// WithReplica adds the replica number to the context metadata of the gNMI
// server query.
func (n *{{ .FakeRootTypePathName }}) WithReplica(replica int) *{{ .FakeRootTypePathName }} {
	genutil.PutReplica(n, replica)
	return n
}

// WithSubscriptionMode specifies the subscription mode in the underlying gNMI
// subscribe.
func (n *{{ .FakeRootTypePathName }}) WithSubscriptionMode(mode gpb.SubscriptionMode) *{{ .FakeRootTypePathName }} {
	genutil.PutSubscriptionMode(n, mode)
	return n
}

// WithClient allows the user to provide a gNMI client. This allows for creation
// of tests for multiple gNMI clients to a single DUT.
func (n *{{ .FakeRootTypePathName }}) WithClient(c gpb.GNMIClient) *{{ .FakeRootTypePathName }} {
	genutil.PutClient(n, c)
	return n
}

{{- if .GenerateConfigCode }}

// NewBatch returns a newly instantiated SetRequestBatch object for batching set requests.
func (d *{{ .FakeRootTypePathName }}) NewBatch() *{{ .ConfigPkgAccessor }}SetRequestBatch {
	return {{ .ConfigPkgAccessor }}NewSetBatch(d)
}
{{- else }}

// NewBatch returns a newly instantiated SetRequestBatch object for batching set requests.
func (d *{{ .FakeRootTypePathName }}) NewBatch() *{{ .SchemaStructPkgAccessor }}Batch {
	return {{ .SchemaStructPkgAccessor }}NewBatch(d)
}
{{- end }}
`)

	// goLeafConvertTemplate contains the per-leaf helper that retrieves the leaf
	// value from its containing parent GoStruct to complete the unmarshalling process.
	goLeafConvertTemplate = mustTemplate("goLeafConvertHelper", `{{ $QualifiedGoTypeName := printf "%sQualified%s" .SchemaStructPkgAccessor .GoType.TransformedGoTypeName }}
// convert{{ .PathStructName }} extracts the value of the leaf {{ .GoFieldName }} from its parent {{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}
// and combines the update with an existing Metadata to return a *{{ $QualifiedGoTypeName }}.
func convert{{ .PathStructName }}(t testing.TB, md *genutil.Metadata, parent *{{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}) *{{ $QualifiedGoTypeName }} {
	t.Helper()
	qv := &{{ $QualifiedGoTypeName }}{
		Metadata: md,
	}
	val := parent.{{ .GoFieldName }}
	if !reflect.ValueOf(val).IsZero() {
		{{- if .SpecialConversionFn }}
		qv.SetVal({{ .SpecialConversionFn }}({{ if .IsScalarField -}} * {{- end -}} val))
		{{- else }}
		qv.SetVal({{ if .IsScalarField -}} * {{- end -}} val)
		{{- end }}
	}
	return qv
}
`)

	// goNodeGetTemplate contains the per-node helper that makes the Get telemetry
	// call, unmarshals the payload, and returns the result to the user.
	goNodeGetTemplate = mustTemplate("goLeafGetMethod", `{{ $QualifiedGoTypeName := printf "%sQualified%s" .SchemaStructPkgAccessor .GoType.TransformedGoTypeName }}
// Lookup fetches the value at {{ .YANGPath }} with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *{{ .PathStructName }}) Lookup(t testing.TB) *{{ $QualifiedGoTypeName }} {
	t.Helper()
	goStruct := &{{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}{}
	md, ok := {{ .SchemaStructPkgAccessor }}Lookup(t, n, "{{ .GoStructTypeName }}", goStruct, {{ .GoType.IsLeaf }}, {{ .PreferShadowPath }})
	if ok {
		{{- if .GoType.IsLeaf }}
		return convert{{ .PathStructName }}(t, md, goStruct)
		{{- else }}
		return (&{{ $QualifiedGoTypeName }}{
			Metadata: md,
		}).SetVal(goStruct)
		{{- end }}
	}
	{{- if .GoType.HasDefault }}
	return (&{{ $QualifiedGoTypeName }}{
		Metadata: md,
	}).SetVal(goStruct.Get{{ .GoFieldName }}())
	{{- else }}
	return nil
	{{- end }}
}

// Get fetches the value at {{ .YANGPath }} with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *{{ .PathStructName }}) Get(t testing.TB) {{ .GoType.GoTypeName }} {
	t.Helper()
	return n.Lookup(t).Val(t)
}

{{- if not .IsRoot }}

// Lookup fetches the values at {{ .YANGPath }} with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) Lookup(t testing.TB) []*{{ $QualifiedGoTypeName }} {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*{{ $QualifiedGoTypeName }}
	for _, prefix := range sortedPrefixes {
		goStruct := &{{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], {{ .SchemaStructPkgAccessor }}GetSchema(), "{{ .GoStructTypeName }}", goStruct, queryPath, {{ .GoType.IsLeaf }}, {{ .PreferShadowPath }})
		if !ok {
			continue
		}
		{{- if .GoType.IsLeaf }}
		qv := convert{{ .PathStructName }}(t, md, goStruct)
		{{- else }}
		qv := (&{{ $QualifiedGoTypeName }}{
			Metadata: md,
		}).SetVal(goStruct)
		{{- end }}
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at {{ .YANGPath }} with a ONCE subscription.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) Get(t testing.TB) []{{ .GoType.GoTypeName }} {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []{{ .GoType.GoTypeName }}
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

{{- end }}
`)

	// goCollectTemplate contains the per-leaf Collect telemetry call. It
	// triggers the asynchronous data collection at this node, and returns a type
	// containing an Await method that blocks for the data collection to complete,
	// as well as unmarshalling that data for the user.
	goCollectTemplate = mustTemplate("goLeafCollectMethod", `{{ $QualifiedGoTypeName := printf "%sQualified%s" .SchemaStructPkgAccessor .GoType.TransformedGoTypeName }}
// Collect starts an asynchronous collection of the values at {{ .YANGPath }} with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *{{ .PathStructName }}) Collect(t testing.TB, duration time.Duration) *{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }} {
	t.Helper()
	c := &{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }}{}
	c.W = n.Watch(t, duration, func(v *{{ $QualifiedGoTypeName }}) bool {
		{{- if .GoType.IsLeaf }}
		c.Data = append(c.Data, v)
		{{- else }}
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&{{ $QualifiedGoTypeName }}{
			Metadata: v.Metadata,
		}).SetVal(copy.({{.GoType.GoTypeName}})))
		{{- end }}
		return false
	})
	return c
}

func watch_{{ .PathStructName }}(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *{{ $QualifiedGoTypeName }}) bool) *{{ .SchemaStructPkgAccessor }}{{ .GoType.TransformedGoTypeName }}Watcher {
	t.Helper()
	w := &{{ .SchemaStructPkgAccessor }}{{ .GoType.TransformedGoTypeName }}Watcher{}
	gs := &{{ .SchemaStructPkgAccessor }}{{ .GoStructTypeName }}{}
	w.W = genutil.MustWatch(t, n, nil, duration, {{ .GoType.IsLeaf }}, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, {{ .SchemaStructPkgAccessor }}GetSchema(), "{{ .GoStructTypeName }}", gs, queryPath, {{ .GoType.IsLeaf }}, {{ .PreferShadowPath }})
		{{- if .GoType.IsLeaf }}
		return convert{{ .PathStructName }}(t, md, gs), nil
		{{- else }}
		return (&{{ $QualifiedGoTypeName }}{
			Metadata: md,
		}).SetVal(gs), nil
		{{- end }}
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*{{ $QualifiedGoTypeName }})
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at {{ .YANGPath }} with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *{{ .PathStructName }}) Watch(t testing.TB, timeout time.Duration, predicate func(val *{{ $QualifiedGoTypeName }}) bool) *{{ .SchemaStructPkgAccessor }}{{ .GoType.TransformedGoTypeName }}Watcher {
	t.Helper()
	return watch_{{ .PathStructName }}(t, n, timeout, predicate)
}

// Await observes values at {{ .YANGPath }} with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *{{ .PathStructName }}) Await(t testing.TB, timeout time.Duration, val {{ .GoType.GoTypeName }}) *{{ $QualifiedGoTypeName }} {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *{{ $QualifiedGoTypeName }}) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at {{ .YANGPath }} failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds {{ .YANGPath }} to the batch object.
func (n *{{ .PathStructName }}) Batch(t testing.TB, b *{{ .SchemaStructPkgAccessor }}Batch) {
	t.Helper()
	{{ .SchemaStructPkgAccessor }}MustAddToBatch(t, b, n)
}

{{- if not .IsRoot }}

// Collect starts an asynchronous collection of the values at {{ .YANGPath }} with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) Collect(t testing.TB, duration time.Duration) *{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }} {
	t.Helper()
	c := &{{ .SchemaStructPkgAccessor }}Collection{{ .GoType.TransformedGoTypeName }}{}
	c.W = n.Watch(t, duration, func(v *{{ $QualifiedGoTypeName }}) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at {{ .YANGPath }} with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) Watch(t testing.TB, timeout time.Duration, predicate func(val *{{ $QualifiedGoTypeName }}) bool) *{{ .SchemaStructPkgAccessor }}{{ .GoType.TransformedGoTypeName }}Watcher {
	t.Helper()
	return watch_{{ .PathStructName }}(t, n, timeout, predicate)
}

// Batch adds {{ .YANGPath }} to the batch object.
func (n *{{ .PathStructName }}{{ .WildcardSuffix }}) Batch(t testing.TB, b *{{ .SchemaStructPkgAccessor }}Batch) {
	t.Helper()
	{{ .SchemaStructPkgAccessor }}MustAddToBatch(t, b, n)
}
{{- end }}
`)

	// goNodeSetTemplate is used to generate gNMI Set wrapper calls.
	goNodeSetTemplate = mustTemplate("goNodeSetMethods", `
// Delete deletes the configuration at {{ .YANGPath }}.
func (n *{{ .PathStructName }}) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at {{ .YANGPath }} in the given batch object.
func (n *{{ .PathStructName }}) BatchDelete(t testing.TB, b *{{ .ConfigPkgAccessor }}SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at {{ .YANGPath }}.
func (n *{{ .PathStructName }}) Replace(t testing.TB, val {{ .GoType.GoTypeName }}) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, {{ if .IsScalarField -}} & {{- end -}} val)
}

// BatchReplace buffers a config replace operation at {{ .YANGPath }} in the given batch object.
func (n *{{ .PathStructName }}) BatchReplace(t testing.TB, b *{{ .ConfigPkgAccessor }}SetRequestBatch, val {{ .GoType.GoTypeName }}) {
	t.Helper()
	b.BatchReplace(t, n, {{ if .IsScalarField -}} & {{- end -}} val)
}

// Update updates the configuration at {{ .YANGPath }}.
func (n *{{ .PathStructName }}) Update(t testing.TB, val {{ .GoType.GoTypeName }}) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, {{ if .IsScalarField -}} & {{- end -}} val)
}

// BatchUpdate buffers a config update operation at {{ .YANGPath }} in the given batch object.
func (n *{{ .PathStructName }}) BatchUpdate(t testing.TB, b *{{ .ConfigPkgAccessor }}SetRequestBatch, val {{ .GoType.GoTypeName }}) {
	t.Helper()
	b.BatchUpdate(t, n, {{ if .IsScalarField -}} & {{- end -}} val)
}
`)

	// goQualifiedTypeTemplate contains the per-leaf return type used to store the
	// telemetry results. A pointer to a struct is intended as the return type to
	// allow a non-existent node to be expressed as a nil pointer.
	goQualifiedTypeTemplate = mustTemplate("goQualifiedType", `{{ $QualifiedGoTypeName := printf "Qualified%s" .TransformedGoTypeName }}
// {{ $QualifiedGoTypeName }} is a {{ .GoTypeName }} with a corresponding timestamp.
type {{ $QualifiedGoTypeName }} struct {
	*genutil.Metadata
	val {{ .GoTypeName }} // val is the sample value.
	present bool
}

func (q *{{ $QualifiedGoTypeName }}) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the {{ .GoTypeName }} sample, erroring out if not present.
func (q *{{ $QualifiedGoTypeName }}) Val(t testing.TB) {{ .GoTypeName }} {
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

// SetVal sets the value of the {{ .GoTypeName }} sample.
func (q *{{ $QualifiedGoTypeName }}) SetVal(v {{ .GoTypeName }}) *{{ $QualifiedGoTypeName }} {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *{{ $QualifiedGoTypeName }}) IsPresent() bool {
	return q != nil && q.present
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
	goCollectionTypeTemplate = mustTemplate("goCollectionType", `{{ $QualifiedGoTypeName := printf "Qualified%s" .TransformedGoTypeName }}
// Collection{{ .TransformedGoTypeName }} is a telemetry Collection whose Await method returns a slice of {{ .GoTypeName }} samples.
type Collection{{ .TransformedGoTypeName }} struct {
	W *{{ .TransformedGoTypeName }}Watcher
	Data []*{{ $QualifiedGoTypeName }}
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *Collection{{ .TransformedGoTypeName }}) Await(t testing.TB) []*{{ $QualifiedGoTypeName }} {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// {{ .TransformedGoTypeName }}Watcher observes a stream of {{ .GoTypeName }} samples.
type {{ .TransformedGoTypeName }}Watcher struct {
	W *genutil.Watcher
	LastVal *{{ $QualifiedGoTypeName }}
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *{{ .TransformedGoTypeName }}Watcher) Await(t testing.TB) (*{{ $QualifiedGoTypeName }}, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
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

// executeTemplate excutes the input template using the values from GenConfig.
// packageName is the name of the package to use in the template.
// generateConfigCode controls whether to template should contain code for config or telemetry.
// The usable templates are: goCommonHelpers, goFakeRootMethods.
func (cg *GenConfig) executeTemplate(packageName string, generateConfigCode bool, tmpl *template.Template) (string, error) {
	var schemaStructPkgAccessor string
	if cg.GoImports.SchemaStructPkgPath != "" {
		schemaStructPkgAccessor = schemaStructPkgAlias + "."
	}
	var configPkgAccessor string
	if cg.GoImports.ConfigImportPath != "" {
		configPkgAccessor = configPkgAlias + "."
	}
	var header strings.Builder
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
		// ConfigPkgAlias is the package alias for the imported ygen-generated file.
		ConfigPkgAlias string
		// ConfigPkgAccessor is the accessor prefix for types in
		// the config file (if any).
		ConfigPkgAccessor string
		// GenerateConfigCode is a flag for whether config functions are being
		// generated instead of telemetry functions.
		GenerateConfigCode bool
		PreferShadowPath   bool
	}{
		GoImports:               cg.GoImports,
		PackageName:             packageName,
		FakeRootTypePathName:    cg.fakeRootTypeName(true),
		FakeRootTypeName:        cg.fakeRootTypeName(false),
		FakeRootBaseTypeName:    ygot.FakeRootBaseTypeName,
		SchemaStructPkgAlias:    schemaStructPkgAlias,
		SchemaStructPkgAccessor: schemaStructPkgAccessor,
		GenerateConfigCode:      generateConfigCode,
		PreferShadowPath:        cg.PreferShadowPath,
		ConfigPkgAlias:          configPkgAlias,
		ConfigPkgAccessor:       configPkgAccessor,
	}
	if err := tmpl.Execute(&header, s); err != nil {
		return "", err
	}

	return header.String(), nil
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
	}
	if err := goNodeGetTemplate.Execute(&getMethod, s); err != nil {
		util.AppendErr(errs, err)
	}

	if err := goCollectTemplate.Execute(&collectMethod, s); err != nil {
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
func generatePerNodeConfigSnippet(pathStructName string, nodeData *ypathgen.NodeData, fakeRootTypeName, schemaStructPkgAccessor, configPkgAccessor string, preferShadowPath bool) (GoPerNodeCodeSnippet, goTypeData, util.Errors) {
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
		ConfigPkgAccessor       string
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
		ConfigPkgAccessor:       configPkgAccessor,
	}
	var getMethod, replaceMethod, convertHelper strings.Builder
	if nodeData.IsLeaf {
		// Leaf types use their parent GoStruct to unmarshal, before
		// being retrieved out when returned to the user.
		if err := goLeafConvertTemplate.Execute(&convertHelper, s); err != nil {
			util.AppendErr(errs, err)
		}
	}
	if err := goNodeSetTemplate.Execute(&replaceMethod, s); err != nil {
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
	// Lookup()/Collect(), there hasn't been the need for non-leaves so
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
