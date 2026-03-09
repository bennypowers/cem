/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package publishDiagnostics

import (
	"fmt"

	"bennypowers.dev/cem/internal/requirements"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// analyzeRequirementDiagnostics validates conditional requirements extracted from manifest descriptions.
func analyzeRequirementDiagnostics(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	return AnalyzeRequirementDiagnosticsForTest(ctx, doc)
}

// AnalyzeRequirementDiagnosticsForTest is the exported version for testing.
func AnalyzeRequirementDiagnosticsForTest(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	dm, err := ctx.DocumentManager()
	if err != nil {
		helpers.SafeDebugLog("[REQ_DIAGNOSTICS] Failed to get DocumentManager: %v", err)
		return diagnostics
	}

	elements, err := doc.FindCustomElements(dm)
	if err != nil {
		helpers.SafeDebugLog("[REQ_DIAGNOSTICS] Error finding custom elements: %v", err)
		return diagnostics
	}

	for _, element := range elements {
		decl := ctx.FindCustomElementDeclaration(element.TagName)
		if decl == nil {
			continue
		}

		rules := requirements.ExtractElementRules(decl)
		if len(rules) == 0 {
			continue
		}

		facts := requirements.ElementFacts{
			Attributes: make(map[string]string),
			HasAttr:    make(map[string]bool),
		}
		for attrName, attr := range element.Attributes {
			facts.Attributes[attrName] = attr.Value
			facts.HasAttr[attrName] = true
		}

		violations := requirements.Evaluate(rules, facts)
		for _, v := range violations {
			diagnostics = append(diagnostics, createRequirementDiagnostic(v, element.Range))
		}
	}

	helpers.SafeDebugLog("[REQ_DIAGNOSTICS] Generated %d requirement diagnostics", len(diagnostics))
	return diagnostics
}

func createRequirementDiagnostic(v requirements.Violation, elemRange protocol.Range) protocol.Diagnostic {
	diagnostic := protocol.Diagnostic{
		Range:    elemRange,
		Message:  v.Message,
		Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
		Source:   &[]string{"cem-lsp"}[0],
	}

	// Add autofix data for RequireAttrPresent violations (suggest adding the attribute)
	if v.Rule.Requirement.Type == requirements.RequireAttrPresent {
		autofixData := &types.AutofixData{
			Type:       types.DiagnosticTypeRequirementViolation,
			Original:   "",
			Suggestion: fmt.Sprintf(`%s=""`, v.Rule.Requirement.AttrName),
			Range:      elemRange,
		}
		diagnostic.Data = autofixData.ToMap()
	}

	return diagnostic
}
