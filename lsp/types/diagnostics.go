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
package types

import protocol "github.com/tliron/glsp/protocol_3_16"

// DiagnosticType represents the type of diagnostic suggestion
type DiagnosticType string

const (
	DiagnosticTypeSlotSuggestion           DiagnosticType = "slot-suggestion"
	DiagnosticTypeTagSuggestion            DiagnosticType = "tag-suggestion"
	DiagnosticTypeMissingImport            DiagnosticType = "missing-import"
	DiagnosticTypeAttributeSuggestion      DiagnosticType = "attribute-suggestion"
	DiagnosticTypeAttributeValueSuggestion DiagnosticType = "attribute-value-suggestion"
)

// AutofixData contains the data needed for creating autofix code actions
type AutofixData struct {
	Type       DiagnosticType `json:"type"`
	Original   string         `json:"original"`
	Suggestion string         `json:"suggestion"`
	Range      protocol.Range `json:"range"`
	// For missing import diagnostics
	ImportPath string `json:"importPath,omitempty"`
	TagName    string `json:"tagName,omitempty"`
}

// ToMap converts AutofixData to map[string]any for LSP protocol compatibility
func (d *AutofixData) ToMap() map[string]any {
	result := map[string]any{
		"type":       string(d.Type),
		"original":   d.Original,
		"suggestion": d.Suggestion,
		"range":      d.Range,
	}
	if d.ImportPath != "" {
		result["importPath"] = d.ImportPath
	}
	if d.TagName != "" {
		result["tagName"] = d.TagName
	}
	return result
}

// AutofixDataFromMap creates AutofixData from map[string]any
func AutofixDataFromMap(data map[string]any) (*AutofixData, bool) {
	typeStr, typeOk := data["type"].(string)
	original, originalOk := data["original"].(string)
	suggestion, suggestionOk := data["suggestion"].(string)
	rangeData, rangeOk := data["range"]

	if !typeOk || !originalOk || !suggestionOk || !rangeOk {
		return nil, false
	}

	// Convert range data back to protocol.Range
	var fixRange protocol.Range
	if rangeMap, ok := rangeData.(map[string]any); ok {
		if start, startOk := rangeMap["start"].(map[string]any); startOk {
			if line, lineOk := start["line"].(float64); lineOk {
				fixRange.Start.Line = uint32(line)
			}
			if char, charOk := start["character"].(float64); charOk {
				fixRange.Start.Character = uint32(char)
			}
		}
		if end, endOk := rangeMap["end"].(map[string]any); endOk {
			if line, lineOk := end["line"].(float64); lineOk {
				fixRange.End.Line = uint32(line)
			}
			if char, charOk := end["character"].(float64); charOk {
				fixRange.End.Character = uint32(char)
			}
		}
	} else if protoRange, ok := rangeData.(protocol.Range); ok {
		fixRange = protoRange
	} else {
		return nil, false
	}

	result := &AutofixData{
		Type:       DiagnosticType(typeStr),
		Original:   original,
		Suggestion: suggestion,
		Range:      fixRange,
	}

	// Optional fields for missing import diagnostics
	if importPath, ok := data["importPath"].(string); ok {
		result.ImportPath = importPath
	}
	if tagName, ok := data["tagName"].(string); ok {
		result.TagName = tagName
	}

	return result, true
}
