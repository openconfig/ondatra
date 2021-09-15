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

// This binary generates Go code corresponding to an input YANG schema.
// The input set of YANG modules are read, parsed using Goyang, and handed as
// input to the codegen package which generates the corresponding Go code.
package main

import (
	"flag"
	"fmt"
	"math"
	"path/filepath"
	"strings"

	log "github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/genutil"
	"github.com/openconfig/ygot/util"
	"github.com/openconfig/ondatra/telemgen"
)

// Flags for telemgen's configuration
var (
	yangPaths               = flag.String("path", "", "Comma separated list of paths to be recursively searched for included modules or submodules within the defined YANG modules.")
	packageName             = flag.String("package_name", "telemetry", "The name of the Go package that should be generated.")
	schemaStructPath        = flag.String("schema_struct_path", "", "The Go import path for the schema structs package. This should be specified if and only if schema structs are not being generated at the same time as the path structs.")
	excludeModules          = flag.String("exclude_modules", "", "Comma separated set of module names that should be excluded from code generation this can be used to ensure overlapping namespaces can be ignored.")
	outputDir               = flag.String("output_dir", "", "The directory that the Go package should be written to.")
	generateConfigFunc      = flag.Bool("generate_config_func", false, "Whether config functions should be generated instead of telemetry functions.")
	preferShadowPath        = flag.Bool("prefer_shadow_path", false, "Whether to use the \"shadow-path\" or the \"path\" struct tags when both are present when marshalling or unmarshalling a GoStruct.")
	listBuilderKeyThreshold = flag.Uint("list_builder_key_threshold", 0, "The threshold equal or over which the builder API is used for key population. 0 means to never use the builder API.")
	// The number of files into which to split the telemetry functions.
	telemFuncsFileN int
	// The number of files into which to split the telemetry type definitions.
	telemTypesFileN int
)

func init() {
	flag.IntVar(&telemFuncsFileN, "telemetry_funcs_file_split", 10, "The number of files into which to split the telemetry functions.")
	flag.IntVar(&telemTypesFileN, "telemetry_types_file_split", 10, "The number of files into which to split the telemetry type definitions.")
}

// Static values for telemgen's configuration
const (
	// If set to true, circular dependencies between submodules are ignored.
	ignoreCircDeps = false
	// The name of the fake root entity. This name will be capitalized for exporting.
	fakeRootName = "device"
	// Import path for gNMI's proto package.
	gnmiProtoPath = "github.com/openconfig/gnmi/proto/gnmi"
	// Import path for ygot.
	ygotImportPath = "github.com/openconfig/ygot/ygot"
	// Import path for ytypes.
	ytypesImportPath = "github.com/openconfig/ygot/ytypes"
	// Import path for goyang's yang package.
	goyangImportPath = "google3/third_party/golang/goyang/pkg/yang/yang"
	// Import path for ONDATRA's telemgo package.
	telemgoImportPath = "github.com/openconfig/ondatra/telemgen/telemgo"
	// Import path for the proto library.
	protoLibImportPath = "google3/third_party/golang/protobuf/v2/proto/proto"
)

// File names for the output files.
const (
	// telemFuncsFileFmt is the format string filename (missing index) to
	// be used by the telemetry calls when Go code is output to a
	// directory.
	telemFuncsFileFmt = "telemetry-%d.go"
	// telemTypesFileFmt is the format string filename (missing index) to
	// be used for the return types used by telemetry calls when Go code is
	// output to a directory.
	telemTypesFileFmt = "telem_types-%d.go"
	// telemHelpersFileName is the file name containing hand-written helper
	// functions.
	telemHelpersFileName = "telem_helpers.go"
)

// makeFileOutputSpec generates a map, keyed by filenames, to strings
// containing the code to be output to the corresponding filename. It puts the
// telemetry methods and telemetry method return types each in their own files,
// while also allowing the first to be split into multiple files using the
// telemFuncsFileN parameter as it can be quite large. The file splitting is
// done roughly by splitting on the number of nodes, so it's an approximation
// instead of an exact split, which does have some variance.
func makeFileOutputSpec(telemCode *telemgen.GeneratedTelemetryCode, telemFuncsFileN, telemTypesFileN int) (map[string]string, error) {
	var b strings.Builder
	b.WriteString(telemCode.CommonHeader)
	b.WriteString(telemCode.OneOffHeader)
	out := map[string]string{telemHelpersFileName: b.String()}

	// Split the telemetry API code into files.
	var snippets []string
	for _, s := range telemCode.GoReturnTypeSnippets {
		snippets = append(snippets, s.String())
	}
	files, err := splitCodeSnippets(snippets, telemCode.CommonHeader, telemTypesFileN)
	if err != nil {
		return nil, err
	}
	for i, file := range files {
		out[fmt.Sprintf(telemTypesFileFmt, i)] = file
	}

	snippets = []string{}
	for _, s := range telemCode.GoPerNodeSnippets {
		snippets = append(snippets, s.String())
	}
	files, err = splitCodeSnippets(snippets, telemCode.CommonHeader, telemFuncsFileN)
	if err != nil {
		return nil, err
	}
	for i, file := range files {
		out[fmt.Sprintf(telemFuncsFileFmt, i)] = file
	}

	return out, nil
}

// splitCodeSnippets splits the code snippets roughly evenly into the
// specified number of files.
func splitCodeSnippets(ss []string, commonHeader string, fileN int) ([]string, error) {
	if fileN == 0 {
		return nil, nil
	}
	snippetN := len(ss)
	if fileN < 1 || fileN > snippetN {
		return nil, fmt.Errorf("requested %d files, but must be between 1 and %d (number of snippets)", fileN, snippetN)
	}

	files := make([]string, 0, fileN)
	snippetsPerFile := int(math.Ceil(float64(snippetN) / float64(fileN)))
	// Empty files could appear with certain snippetN/fileN combinations due
	// to the ceiling numbers being used for snippetsPerFile.
	// e.g. 4/3 gives two files of two structs.
	// This is a little more complex, but spreads out the structs more evenly.
	// If we instead use the floor number, and put all remainder structs in
	// the last file, we might double the last file's number of structs if we get unlucky.
	// e.g. 99/10 assigns 18 structs to the last file.
	emptyFiles := fileN - int(math.Ceil(float64(snippetN)/float64(snippetsPerFile)))
	var b strings.Builder
	b.WriteString(commonHeader)
	for i, s := range ss {
		b.WriteString(s)
		// The last file contains the remainder of the structs.
		if i == snippetN-1 || (i+1)%snippetsPerFile == 0 {
			files = append(files, b.String())
			b.Reset()
			b.WriteString(commonHeader)
		}
	}
	for i := 0; i != emptyFiles; i++ {
		files = append(files, commonHeader)
	}

	return files, nil
}

// writeFiles creates or truncates files in a given base directory and writes
// to them. Keys of the contents map are file names, and values are the
// contents to be written. An error is returned if the base directory does not
// exist. If a file cannot be written, the function aborts with the error,
// leaving an unspecified set of the other input files written with their given
// contents.
func writeFiles(dir string, out map[string]string) error {
	for filename, contents := range out {
		if len(contents) == 0 {
			continue
		}
		fh := genutil.OpenFile(filepath.Join(dir, filename))
		if fh == nil {
			return errors.Errorf("could not open file %q", filename)
		}
		if _, err := fh.WriteString(contents); err != nil {
			return err
		}
		// flush & close written files before function finishes.
		defer genutil.SyncFile(fh)
	}

	return nil
}

// main parses command-line flags to determine the set of YANG modules for
// which code generation should be performed, and calls the generation library
// to generate Go code corresponding to their schema. The output is written
// to the specified file.
func main() {
	flag.Parse()
	// Extract the set of modules that code is to be generated for,
	// throwing an error if the set is empty.
	generateModules := flag.Args()
	if len(generateModules) == 0 {
		log.Exitln("Error: no input modules specified")
	}

	// Determine the set of paths that should be searched for included
	// modules. This is supplied by the user as a set of comma-separated
	// paths, so we split the string. Additionally, for each path
	// specified, we append "..." to ensure that the directory is
	// recursively searched.
	var includePaths []string
	if len(*yangPaths) > 0 {
		pathParts := strings.Split(*yangPaths, ",")
		for _, path := range pathParts {
			includePaths = append(includePaths, filepath.Join(path, "..."))
		}
	}

	// Determine which modules the user has requested to be excluded from
	// code generation.
	var modsExcluded []string
	if len(*excludeModules) > 0 {
		modParts := strings.Split(*excludeModules, ",")
		for _, mod := range modParts {
			modsExcluded = append(modsExcluded, mod)
		}
	}

	if *outputDir == "" {
		log.Exitln("Error: outputDir unspecified")
	}

	// Perform the code generation.
	cg := &telemgen.GenConfig{
		PackageName: *packageName,
		GoImports: telemgen.GoImports{
			YgotImportPath:      ygotImportPath,
			YtypesImportPath:    ytypesImportPath,
			GoyangImportPath:    goyangImportPath,
			GNMIProtoPath:       gnmiProtoPath,
			TelemGoImportPath:   telemgoImportPath,
			ProtoLibImportPath:  protoLibImportPath,
			SchemaStructPkgPath: *schemaStructPath,
		},
		FakeRootName:   fakeRootName,
		ExcludeModules: modsExcluded,
		YANGParseOptions: yang.Options{
			IgnoreSubmoduleCircularDependencies: ignoreCircDeps,
		},
		GeneratingBinary:        genutil.CallerName(),
		ListBuilderKeyThreshold: *listBuilderKeyThreshold,
		PreferShadowPath:        *preferShadowPath,
	}

	var genCode *telemgen.GeneratedTelemetryCode
	var errs util.Errors
	if *generateConfigFunc {
		if genCode, errs = cg.GenerateConfigCode(generateModules, includePaths); errs != nil {
			log.Exitf("Error: Generating Config Functions: %s\n", errs)
		}
	} else {
		if genCode, errs = cg.GenerateTelemetryCode(generateModules, includePaths); errs != nil {
			log.Exitf("Error: Generating Telemetry Functions: %s\n", errs)
		}
	}

	out, err := makeFileOutputSpec(genCode, telemFuncsFileN, telemTypesFileN)
	if err != nil {
		log.Exitf("Error: Generating Code: %s\n", err)
	}
	if err := writeFiles(*outputDir, out); err != nil {
		log.Exitf("Error: %v", err)
	}
}
