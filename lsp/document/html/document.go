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
package html

import (
	"fmt"
	"maps"
	"strings"
	"sync"

	htmllang "bennypowers.dev/cem/internal/languages/html"
	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/document/base"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// HTMLDocument represents an HTML document with tree-sitter parsing
type HTMLDocument struct {
	*base.BaseDocument
	importMap map[string]string
	importMu  sync.RWMutex
}

// NewHTMLDocument creates a new HTML document
func NewHTMLDocument(uri, content string, version int32, language string) *HTMLDocument {
	doc := &HTMLDocument{}
	doc.BaseDocument = base.NewBaseDocument(uri, content, version, language, htmllang.ReturnParser, doc)
	return doc
}

// ImportMap returns the import map from <script type="importmap"> tag
func (d *HTMLDocument) ImportMap() map[string]string {
	d.importMu.RLock()
	defer d.importMu.RUnlock()
	if d.importMap == nil {
		return nil
	}
	result := make(map[string]string, len(d.importMap))
	maps.Copy(result, d.importMap)
	return result
}

// SetImportMap sets the import map
func (d *HTMLDocument) SetImportMap(importMap map[string]string) {
	d.importMu.Lock()
	defer d.importMu.Unlock()
	d.importMap = importMap
}

// Parse parses the HTML content using tree-sitter
func (d *HTMLDocument) Parse(content string) error {
	d.UpdateContent(content, d.Version())

	parser := htmllang.BorrowParser()
	if parser == nil {
		return fmt.Errorf("failed to get HTML parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	return nil
}

// ParseWithRanges parses using tree-sitter language injection: the HTML parser
// is restricted to the given byte ranges (e.g. HTML regions extracted from a
// template grammar). Byte positions in the resulting tree refer to the original
// source, so line/column mapping works without a whitespace-filled buffer.
func (d *HTMLDocument) ParseWithRanges(content string, ranges []ts.Range) error {
	d.UpdateContent(content, d.Version())

	if len(ranges) == 0 {
		return nil
	}

	parser := htmllang.BorrowParser()
	if parser == nil {
		return fmt.Errorf("failed to get HTML parser")
	}
	d.SetParser(parser)

	if err := parser.SetIncludedRanges(ranges); err != nil {
		return fmt.Errorf("failed to set included ranges: %w", err)
	}

	tree := parser.Parse([]byte(content), nil)

	_ = parser.SetIncludedRanges(nil)

	d.SetTree(tree)

	return nil
}

// findCustomElements finds custom elements in the HTML document
func (d *HTMLDocument) findCustomElements(handler *Handler) ([]types.CustomElementMatch, error) {
	var elements []types.CustomElementMatch

	if handler == nil {
		helpers.SafeDebugLog("[HTML] Handler is nil in findCustomElements")
		return elements, nil
	}

	tree, content := d.TreeAndContent()
	if tree == nil {
		return elements, nil
	}

	matcher, err := Q.GetCachedQueryMatcher(handler.queryManager, handler.language, "customElements")
	if err != nil {
		helpers.SafeDebugLog("[HTML] Failed to create custom elements matcher: %v", err)
		return elements, nil
	}
	defer matcher.Close()

	root := tree.RootNode()
	contentBytes := []byte(content)

	if root.EndByte() > uint(len(contentBytes)) {
		return elements, nil
	}

	for captureMap := range matcher.ParentCaptures(root, contentBytes, "element") {
		if tagNames, ok := captureMap["tag.name"]; ok && len(tagNames) > 0 {
			tagName := tagNames[0].Text
			tagRange := d.ByteRangeToProtocolRange(content, tagNames[0].StartByte, tagNames[0].EndByte)

			attributes := make(map[string]types.AttributeMatch)
			if attrNames, ok := captureMap["attr.name"]; ok {
				valuesByPosition := make(map[uint]string)

				if attrValues, ok := captureMap["attr.value"]; ok {
					for _, attrValue := range attrValues {
						value := attrValue.Text
						if len(value) >= 2 && (value[0] == '"' || value[0] == '\'') {
							value = value[1 : len(value)-1]
						}
						valuesByPosition[attrValue.StartByte] = value
					}
				}

				if unquotedValues, ok := captureMap["attr.unquoted.value"]; ok {
					for _, unquotedValue := range unquotedValues {
						valuesByPosition[unquotedValue.StartByte] = unquotedValue.Text
					}
				}

				for _, attrName := range attrNames {
					attrMatch := types.AttributeMatch{
						Name:  attrName.Text,
						Range: d.ByteRangeToProtocolRange(content, attrName.StartByte, attrName.EndByte),
					}

					var closestValue string
					var closestDistance = ^uint(0)
					for valuePos, value := range valuesByPosition {
						if valuePos > attrName.EndByte {
							distance := valuePos - attrName.EndByte
							if distance < closestDistance {
								betweenText := string(contentBytes[attrName.EndByte:valuePos])
								trimmed := strings.TrimLeft(betweenText, " \t\r\n")
								if len(trimmed) > 0 && trimmed[0] == '=' {
									closestDistance = distance
									closestValue = value
								}
							}
						}
					}
					if closestValue != "" {
						attrMatch.Value = closestValue
					}

					attributes[attrName.Text] = attrMatch
				}
			}

			elements = append(elements, types.CustomElementMatch{
				TagName:    tagName,
				Range:      tagRange,
				Attributes: attributes,
			})
		}
	}

	return elements, nil
}

// analyzeCompletionContext analyzes completion context for HTML documents
func (d *HTMLDocument) analyzeCompletionContext(position protocol.Position, handler *Handler) *types.CompletionAnalysis {
	analysis := &types.CompletionAnalysis{
		Type: types.CompletionUnknown,
	}

	tree := d.Tree()
	if tree == nil {
		return analysis
	}

	content, err := d.Content()
	if err != nil {
		return analysis
	}

	byteOffset := d.PositionToByteOffset(position, content)

	helpers.SafeDebugLog("[HTML] analyzeCompletionContext: content=%q, position=%+v, byteOffset=%d", content, position, byteOffset)

	completionMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, handler.language, "completionContext")
	if err != nil {
		helpers.SafeDebugLog("[HTML] Failed to create completion context matcher: %v", err)
		return analysis
	}
	defer completionMatcher.Close()

	allTagNames := []Q.CaptureInfo{}
	allAttrNames := []Q.CaptureInfo{}
	allAttrValues := []Q.CaptureInfo{}

	captureCount := 0
	for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
		captureCount++
		if tagNames, exists := captureMap["tag.name"]; exists {
			allTagNames = append(allTagNames, tagNames...)
		}
		if attrNames, exists := captureMap["attr.name"]; exists {
			allAttrNames = append(allAttrNames, attrNames...)
		}
		if attrValues, exists := captureMap["attr.value"]; exists {
			allAttrValues = append(allAttrValues, attrValues...)
		}
	}

	for _, capture := range allTagNames {
		if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
			captureText := capture.Text
			if !strings.Contains(captureText, " ") &&
				!strings.Contains(captureText, "=") &&
				!strings.Contains(captureText, ">") &&
				!strings.Contains(captureText, "\n") {
				analysis.Type = types.CompletionTagName
				analysis.TagName = strings.TrimLeft(captureText, "<")
				return analysis
			}
		}
	}

	for _, capture := range allAttrValues {
		if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
			captureText := capture.Text
			openBrackets := strings.Count(captureText, "<")
			newlineCount := strings.Count(captureText, "\n")

			if openBrackets <= 1 && newlineCount <= 1 {
				analysis.Type = types.CompletionAttributeValue
				for _, attrCapture := range allAttrNames {
					if attrCapture.EndByte <= byteOffset {
						analysis.AttributeName = attrCapture.Text
					}
				}
				for _, tagCapture := range allTagNames {
					if tagCapture.EndByte <= byteOffset {
						analysis.TagName = tagCapture.Text
					}
				}
				return analysis
			}
		}
	}

	for _, capture := range allAttrNames {
		if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
			captureText := capture.Text
			if !strings.Contains(captureText, "<") &&
				!strings.Contains(captureText, ">") &&
				!strings.Contains(captureText, "\n") &&
				!strings.Contains(captureText, "=") {
				analysis.Type = types.CompletionAttributeName
				analysis.AttributeName = captureText
				for _, tagCapture := range allTagNames {
					if tagCapture.EndByte <= byteOffset {
						analysis.TagName = tagCapture.Text
					}
				}
				return analysis
			}
		}
	}

	if analysis.Type == types.CompletionUnknown && captureCount > 0 {
		for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
			for captureName, captures := range captureMap {
				if captureName == "tag.name" {
					for _, capture := range captures {
						if strings.Contains(capture.Text, "-") &&
							capture.EndByte < byteOffset &&
							byteOffset <= capture.EndByte+3 {

							afterTag := content[capture.EndByte:byteOffset]
							if strings.TrimSpace(afterTag) == "" {
								analysis.Type = types.CompletionAttributeName
								analysis.TagName = capture.Text
								return analysis
							}
						}
					}
				}
			}
		}

		for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
			var associatedTagName string
			var latestAttrEnd uint

			if tagNames, exists := captureMap["tag.name"]; exists && len(tagNames) > 0 {
				for _, tagCapture := range tagNames {
					if strings.Contains(tagCapture.Text, "-") && tagCapture.StartByte < byteOffset {
						associatedTagName = tagCapture.Text
					}
				}
			}

			if attrValues, exists := captureMap["attr.value"]; exists {
				for _, attrValue := range attrValues {
					if attrValue.EndByte < byteOffset && attrValue.EndByte > latestAttrEnd {
						latestAttrEnd = attrValue.EndByte
					}
				}
			}

			if unquotedValues, exists := captureMap["attr.unquoted.value"]; exists {
				for _, unquotedValue := range unquotedValues {
					if unquotedValue.EndByte < byteOffset && unquotedValue.EndByte > latestAttrEnd {
						latestAttrEnd = unquotedValue.EndByte
					}
				}
			}

			if associatedTagName != "" && latestAttrEnd > 0 {
				gapContent := content[latestAttrEnd:byteOffset]
				if len(gapContent) > 0 && (gapContent[0] == '"' || gapContent[0] == '\'') {
					gapContent = gapContent[1:]
				}

				if strings.TrimSpace(gapContent) == "" && len(gapContent) <= 5 {
					analysis.Type = types.CompletionAttributeName
					analysis.TagName = associatedTagName
					return analysis
				}
			}
		}
	}

	if analysis.Type == types.CompletionUnknown && byteOffset > 0 {
		node := tree.RootNode().DescendantForByteRange(byteOffset-1, byteOffset)
		if node != nil {
			kind := node.Kind()
			parentKind := ""
			if p := node.Parent(); p != nil {
				parentKind = p.Kind()
			}
			if (kind == "<" || parentKind == "ERROR" || parentKind == "start_tag") &&
				content[byteOffset-1] == '<' {
				analysis.Type = types.CompletionTagName
			}
		}
	}

	helpers.SafeDebugLog("[HTML] analyzeCompletionContext result: Type=%d, TagName=%q", analysis.Type, analysis.TagName)

	return analysis
}

// CompletionPrefix extracts the prefix being typed for filtering completions.
// Uses analysis struct fields and LineContent only -- never searches full
// document content, which may contain template tokens after injection parsing.
func (d *HTMLDocument) CompletionPrefix(analysis *types.CompletionAnalysis) string {
	if analysis == nil {
		return ""
	}

	switch analysis.Type {
	case types.CompletionTagName:
		if analysis.TagName != "" {
			return analysis.TagName
		}
		if analysis.LineContent == "" {
			return ""
		}
		if i := strings.LastIndex(analysis.LineContent, "<"); i != -1 {
			after := analysis.LineContent[i+1:]
			for j, r := range after {
				if r == ' ' || r == '>' || r == '\n' || r == '\t' {
					return after[:j]
				}
			}
			return after
		}
		return ""

	case types.CompletionAttributeName:
		if analysis.AttributeName != "" {
			return analysis.AttributeName
		}
		if analysis.LineContent == "" {
			return ""
		}
		words := strings.Fields(analysis.LineContent)
		if len(words) > 0 {
			last := words[len(words)-1]
			if !strings.Contains(last, "<") && !strings.Contains(last, "=") {
				return last
			}
		}
		return ""

	case types.CompletionAttributeValue:
		line := analysis.LineContent
		if line == "" {
			return ""
		}
		lastQuote := max(strings.LastIndex(line, "\""), strings.LastIndex(line, "'"))
		if lastQuote != -1 {
			return line[lastQuote+1:]
		}
		return ""

	default:
		return ""
	}
}

// FindModuleScript finds insertion point in module script
func (d *HTMLDocument) FindModuleScript() (protocol.Position, bool) {
	scriptTags := d.ScriptTags()
	for _, script := range scriptTags {
		if script.IsModule && script.Src == "" {
			return script.ContentRange.End, true
		}
	}
	return protocol.Position{}, false
}

// findHeadInsertionPoint finds insertion point just before </head> using tree-sitter.
func (d *HTMLDocument) findHeadInsertionPoint(handler *Handler) (protocol.Position, bool) {
	tree, content := d.TreeAndContent()
	if tree == nil {
		return protocol.Position{}, false
	}

	matcher, err := Q.GetCachedQueryMatcher(handler.queryManager, handler.language, "headElements")
	if err != nil {
		helpers.SafeDebugLog("[HTML] findHeadInsertionPoint: failed to create matcher: %v", err)
		return protocol.Position{}, false
	}
	defer matcher.Close()

	root := tree.RootNode()
	for captureMap := range matcher.ParentCaptures(root, []byte(content), "head.element") {
		if endTags, exists := captureMap["end.tag"]; exists && len(endTags) > 0 {
			pos := d.ByteOffsetToPosition(endTags[0].StartByte)
			return pos, true
		}
	}

	return protocol.Position{}, false
}
