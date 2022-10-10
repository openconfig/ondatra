// Copyright 2022 Google LLC
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

// Package gnmifix fixes deprecated Ondatra Telemetry code.
package gnmifix

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer fixes deprecated Ondatra Telemetry code.
var Analyzer = &analysis.Analyzer{
	Name:     "gnmifix",
	Doc:      "fix deprecated ondatra usage",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var telemFns = map[string]bool{
	"Get":     true,
	"Lookup":  true,
	"Watch":   true,
	"Collect": true,
	"Await":   true,
}
var (
	ondatraPath = regexp.MustCompile("openconfig/ondatra/[telemetry|config]")
)

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	var typesDiag []*analysis.Diagnostic
	var funcsDiag []*analysis.Diagnostic

	inspect.Preorder([]ast.Node{&ast.SelectorExpr{}}, func(n ast.Node) {
		if !telemetryTypesToYGNMITypes(n) {
			return
		}
		buf := &strings.Builder{}
		printer.Fprint(buf, pass.Fset, n)
		diag := analysis.Diagnostic{
			Pos:     n.Pos(),
			End:     n.End(),
			Message: "Deprecated usage of ondatra generated types",
			SuggestedFixes: []analysis.SuggestedFix{{
				Message: "ygmni type",
				TextEdits: []analysis.TextEdit{{
					Pos:     n.Pos(),
					End:     n.End(),
					NewText: []byte(buf.String()),
				}},
			}},
		}
		typesDiag = append(typesDiag, &diag)
	})

	inspect.Preorder([]ast.Node{&ast.CallExpr{}}, func(n ast.Node) {
		expr := n.(*ast.CallExpr)
		selector, ok := expr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		// If Get, Lookup, etc is called on an Ondatra telemetry/config struct.
		if telemFns[selector.Sel.Name] && ondatraPath.Match([]byte(pass.TypesInfo.TypeOf(selector.X).String())) {
			buf := &strings.Builder{}
			for i := 1; i < len(expr.Args); i++ {
				fmt.Fprint(buf, ", ")
				printer.Fprint(buf, pass.Fset, expr.Args[i])
			}
			i := &callInfo{
				call: selector.Sel.Name,
				args: buf.String(),
				pos:  expr.Pos(),
				end:  expr.End(),
			}
			handleTelemCalls(pass, selector.X, i, &funcsDiag)
		}
	})

	for _, typeD := range funcsDiag {
		pass.Report(*typeD)
	}

	// Prevent overlapping diagnostics as it prevents fixes from being applied.
	for _, typeD := range typesDiag {
		report := true
		for _, funcD := range funcsDiag {
			if typeD.Pos >= funcD.Pos && typeD.End <= funcD.End {
				report = false
			}
		}
		if report {
			pass.Report(*typeD)
		}
	}

	return nil, nil
}

type callInfo struct {
	path      string
	config    bool
	caller    string
	pathFound bool
	client    string
	call      string
	args      string
	wildcard  bool
	indirect  bool
	pos       token.Pos
	end       token.Pos
	pathOnly  bool
}

func (c callInfo) String() string {
	if c.pathOnly && c.pathFound {
		return fmt.Sprintf("ocpath.Root().%s", strings.TrimSuffix(c.path, "."))
	} else if c.pathOnly {
		return fmt.Sprintf("%s.%s", c.caller, strings.TrimSuffix(c.path, "."))
	}
	queryType := "State()"
	if c.config {
		queryType = "Config()"
	}
	var callSuffix string
	if c.wildcard {
		callSuffix = "All"
	}
	caller := "ocpath.Root()"
	if c.indirect {
		caller = c.caller
	}
	return fmt.Sprintf("gnmi.%s%s(t, %s, %s.%s%s%s)", c.call, callSuffix, c.client, caller, c.path, queryType, c.args)
}

// handleCallExpr builds the chained path calls.
func handleCallExpr(expr *ast.CallExpr, fset *token.FileSet, info *callInfo) ast.Expr {
	buf := &strings.Builder{}
	selector := expr.Fun.(*ast.SelectorExpr)
	// If reached a Telemetry() or State() call means full path is found.
	if selector.Sel.Name == "Telemetry" {
		info.pathFound = true
	} else if selector.Sel.Name == "Config" {
		info.config = true
		info.pathFound = true
	}
	// If haven't found the last path elem or this is a different call chain.
	if !info.pathFound && !info.indirect {
		printer.Fprint(buf, fset, expr)
		fullCall := buf.String()
		call := fullCall[strings.LastIndex(fullCall, selector.Sel.Name):]
		if strings.Contains(call, "Any(") {
			info.wildcard = true
		}
		info.path = fmt.Sprintf("%s.%s", call, info.path)
	}
	return selector.X
}

func buildDiagnostic(expr *ast.Ident, info *callInfo) *analysis.Diagnostic {
	info.client = expr.Name
	msg := "Deprecated usage of Ondatra Get, Lookup, etc"
	if info.pathOnly {
		msg = "Deprecated Ondatra telemetry path"
	}
	diag := &analysis.Diagnostic{
		Pos:     info.pos,
		End:     info.end,
		Message: msg,
		SuggestedFixes: []analysis.SuggestedFix{{
			Message: "Fix for deprecated ondatra telemetry",
			TextEdits: []analysis.TextEdit{{
				Pos:     info.pos,
				End:     info.end,
				NewText: []byte(info.String()),
			}},
		}},
	}
	return diag
}

// handleTelemCalls recursively walks up a path to build a callInfo for a telemetry call.
func handleTelemCalls(pass *analysis.Pass, n ast.Node, info *callInfo, diags *[]*analysis.Diagnostic) {
	switch expr := n.(type) {
	case *ast.CallExpr:
		next := handleCallExpr(expr, pass.Fset, info)
		handleTelemCalls(pass, next, info, diags)
	case *ast.Ident: // At an Ident, reached the end of the call chain.
		if !info.indirect {
			info.caller = expr.Name
		}
		// If the root is a DUT, then we're done.
		typ := pass.TypesInfo.TypeOf(expr)
		if strings.Contains(typ.String(), "ondatra.DUTDevice") || strings.Contains(typ.String(), "ondatra.ATEDevice") { // If called on DUTDevice have everything we need.
			diag := buildDiagnostic(expr, info)
			*diags = append(*diags, diag)
			return
		}
		// If not, try to find device by looking at its declaration.
		if expr.Obj != nil && expr.Obj.Decl != nil {
			assign, ok := expr.Obj.Decl.(*ast.AssignStmt)
			if ok {
				info.indirect = true
				// Walk up the assignment to find dut.
				handleTelemCalls(pass, assign.Rhs[0], info, diags)
				// Start a new walk, fixing the intermediate path.
				handleTelemCalls(pass, assign.Rhs[0], &callInfo{
					pos:      assign.Rhs[0].Pos(),
					end:      assign.Rhs[0].End(),
					pathOnly: true,
				}, diags)
			}
		}
	}
}

var (
	globalType = map[string]bool{
		"Uint64": true,
		"Uint32": true,
		"Uint16": true,
		"Uint8":  true,
		"String": true,
		"Bool":   true,
	}
)

// telemetryTypesToYGNMITypes converts telemetry.foo in the equivalent ygnmi type.
func telemetryTypesToYGNMITypes(n ast.Node) bool {
	expr := n.(*ast.SelectorExpr) // Match against expression like foo.bar
	x, ok := expr.X.(*ast.Ident)
	if !ok {
		return false
	}
	if x.Name != "telemetry" { // Match against telemetry.bar.
		return false
	}
	name := expr.Sel.Name
	if !strings.HasPrefix(name, "Qualified") { // Turn telemetry.bar in to oc.bar.
		x.Name = "oc"
		return true
	}
	typeStr := ""
	typ := strings.TrimPrefix(name, "Qualified") // Turn telemetry.QualifiedFoo into ygnmi.Value[foo].
	if strings.HasSuffix(typ, "Slice") {
		typeStr += "[]"
	}
	typ = strings.TrimSuffix(typ, "Slice")
	if globalType[typ] {
		typeStr += strings.ToLower(typ)
	} else if strings.Contains(typ, "_Union") || strings.Contains(typ, "E_") {
		typeStr += fmt.Sprintf("oc.%s", typ)
	} else {
		typeStr += fmt.Sprintf("*oc.%s", typ)
	}
	expr.Sel.Name = fmt.Sprintf("Value[%s]", typeStr)
	x.Name = "ygnmi"
	return true
}
