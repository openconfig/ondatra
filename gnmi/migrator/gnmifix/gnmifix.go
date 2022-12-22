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
	"Replace": true,
	"Delete":  true,
	"Update":  true,
}

var (
	ondatraPath = regexp.MustCompile("openconfig/ondatra/(?:telemetry|config).*Path")
	telemCaller = regexp.MustCompile(`ondatra\.DUTDevice|ondatra\.ATEDevice|otg\.OTG`)
)

const newImports = `
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/oc"
`

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	var typesDiag []*analysis.Diagnostic
	var funcsDiag []*analysis.Diagnostic

	inspect.Preorder([]ast.Node{&ast.SelectorExpr{}}, func(n ast.Node) {
		pos, end := n.Pos(), n.End() // Save the pos before the ast is modified.
		if !telemetryTypesToYGNMITypes(n) {
			return
		}
		buf := &strings.Builder{}
		printer.Fprint(buf, pass.Fset, n)
		diag := analysis.Diagnostic{
			Pos:     pos,
			End:     end,
			Message: "Deprecated usage of ondatra generated types",
			SuggestedFixes: []analysis.SuggestedFix{{
				Message: "ygmni type",
				TextEdits: []analysis.TextEdit{{
					Pos:     pos,
					End:     end,
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
		if typeStr := pass.TypesInfo.TypeOf(selector.X).String(); telemFns[selector.Sel.Name] && ondatraPath.MatchString(typeStr) {
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

	type diagPos struct {
		Start token.Pos
		End   token.Pos
	}

	diagMap := map[diagPos]*analysis.Diagnostic{}

	// When there are multiple telemetry calls on the same path variable, may have duplicate diagnostics.
	for _, funcsD := range funcsDiag {
		diagMap[diagPos{
			Start: funcsD.Pos,
			End:   funcsD.End,
		}] = funcsD

	}
	for _, diag := range diagMap {
		pass.Report(*diag)
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

	for _, f := range pass.Files {
		// If there are diagnostics, add new imports.
		if len(funcsDiag) == 0 && len(typesDiag) == 0 {
			continue
		}
		importPos := f.Imports[len(f.Imports)-1].End()
		diag := analysis.Diagnostic{
			Pos:     importPos,
			End:     importPos,
			Message: "Adding new telemetry imports",
			SuggestedFixes: []analysis.SuggestedFix{{
				Message: "ygmni imports",
				TextEdits: []analysis.TextEdit{{
					Pos:     importPos,
					End:     importPos,
					NewText: []byte(newImports),
				}},
			}},
		}
		pass.Report(diag)
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
	otg       bool
}

func (c callInfo) pathToString() string {
	left := strings.TrimSuffix(c.caller, ".")
	if c.pathFound && !c.indirect {
		left = "gnmi.OC()"
		if c.otg {
			left = "gnmi.OTG()"
		}
	}

	right := strings.TrimSuffix(c.path, ".")
	if right == "" {
		return left
	}
	return fmt.Sprintf("%s.%s", left, right)
}

func (c callInfo) String() string {
	path := c.pathToString()
	if c.pathOnly {
		return path
	}

	queryType := "State()"
	if c.config {
		queryType = "Config()"
	}

	var callSuffix string
	if c.wildcard {
		callSuffix = "All"
	} else if c.config && (c.call == "Get" || c.call == "Lookup") {
		callSuffix = "Config"
	}
	return fmt.Sprintf("gnmi.%s%s(t, %s, %s.%s%s)", c.call, callSuffix, c.client, path, queryType, c.args)
}

// handleCallExpr builds the chained path calls.
func handleCallExpr(expr *ast.CallExpr, selector *ast.SelectorExpr, fset *token.FileSet, info *callInfo) ast.Expr {
	// If reached a Telemetry() or State() call means full path is found.
	if selector.Sel.Name == "Telemetry" {
		info.pathFound = true
	} else if selector.Sel.Name == "Config" {
		info.config = true
		info.pathFound = true
	}

	xbuf := &strings.Builder{}
	buf := &strings.Builder{}
	printer.Fprint(xbuf, fset, selector.X)
	printer.Fprint(buf, fset, expr)
	fullCall := buf.String()

	call := fullCall[len(xbuf.String())+1:]
	if strings.Contains(call, "Any(") {
		info.wildcard = true
	}

	// If haven't found the last path elem or this is a different call chain.
	if !info.pathFound && !info.indirect {
		info.path = fmt.Sprintf("%s.%s", call, info.path)
	}
	return selector.X
}

func buildDiagnostic(client string, info *callInfo) *analysis.Diagnostic {
	info.client = client
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
		if reportIfValid(pass, expr, info, diags) {
			return
		}
		selector, ok := expr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		next := handleCallExpr(expr, selector, pass.Fset, info)
		handleTelemCalls(pass, next, info, diags)
	case *ast.Ident: // At an Ident, reached the end of the call chain.
		if reportIfValid(pass, expr, info, diags) {
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
	case *ast.SelectorExpr: // Handles case where dut is a struct member.
		reportIfValid(pass, expr, info, diags)
	}
}

func reportIfValid(pass *analysis.Pass, caller ast.Expr, info *callInfo, diags *[]*analysis.Diagnostic) bool {
	buf := &strings.Builder{}
	printer.Fprint(buf, pass.Fset, caller)
	callerStr := buf.String()
	if !info.indirect {
		info.caller = callerStr
	}

	typeStr := pass.TypesInfo.TypeOf(caller).String()
	if !telemCaller.MatchString(typeStr) {
		return false
	}

	info.otg = strings.Contains(typeStr, "otg.OTG")
	diag := buildDiagnostic(callerStr, info)
	*diags = append(*diags, diag)
	return true
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
	if name == "Device" { // The root object was renamed from Device to Root.
		expr.Sel.Name = "Root"
	}

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
