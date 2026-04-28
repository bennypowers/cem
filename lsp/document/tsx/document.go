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
package tsx

import (
	"fmt"
	"strings"

	tsxlang "bennypowers.dev/cem/internal/languages/tsx"
	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/document/base"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

const (
	SCORE_TAG_NAME_BASE   = 10
	SCORE_ATTR_NAME_BASE  = 20
	SCORE_ATTR_VALUE_BASE = 30

	SCORE_CURSOR_BONUS          = 50
	SCORE_BROAD_CAPTURE_PENALTY = 5

	MAX_REASONABLE_TAG_LENGTH = 10
)

// TSXDocument represents a TSX document with tree-sitter parsing
type TSXDocument struct {
	*base.BaseDocument
}

// NewTSXDocument creates a new TSX document
func NewTSXDocument(uri, content string, version int32) *TSXDocument {
	doc := &TSXDocument{}
	doc.BaseDocument = base.NewBaseDocument(uri, content, version, "tsx", tsxlang.ReturnParser, doc)
	return doc
}

// Parse parses the TSX content using tree-sitter
func (d *TSXDocument) Parse(content string) error {
	d.UpdateContent(content, d.Version())

	parser := tsxlang.BorrowParser()
	if parser == nil {
		return fmt.Errorf("failed to get TSX parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	return nil
}

// FindModuleScript returns EOF position for TSX files.
func (d *TSXDocument) FindModuleScript() (protocol.Position, bool) {
	content, err := d.Content()
	if err != nil {
		return protocol.Position{}, false
	}

	lines := strings.Split(content, "\n")
	return protocol.Position{Line: uint32(len(lines)), Character: 0}, true
}

// CompletionPrefix extracts the prefix being typed for filtering completions
func (d *TSXDocument) CompletionPrefix(analysis *types.CompletionAnalysis) string {
	if analysis == nil {
		return ""
	}

	switch analysis.Type {
	case types.CompletionTagName:
		return analysis.TagName
	case types.CompletionAttributeName:
		return analysis.AttributeName
	case types.CompletionAttributeValue:
		return analysis.AttributeName
	default:
		return ""
	}
}

// findCustomElements finds custom elements in TSX documents
func (d *TSXDocument) findCustomElements(handler *Handler) ([]types.CustomElementMatch, error) {
	tree, content := d.TreeAndContent()
	if tree == nil {
		return nil, fmt.Errorf("no tree available for document")
	}

	var elements []types.CustomElementMatch
	elementMap := make(map[string]types.CustomElementMatch)

	customElementsMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "tsx", "customElements")
	if err != nil {
		return nil, fmt.Errorf("failed to create custom elements matcher: %w", err)
	}
	defer customElementsMatcher.Close()

	parentCaptureNames := []string{"element", "self.closing.tag", "start.tag", "incomplete.element"}

	for _, parentName := range parentCaptureNames {
		for captureMap := range customElementsMatcher.ParentCaptures(tree.RootNode(), []byte(content), parentName) {
			if tagNames, exists := captureMap["tag.name"]; exists {
				for _, tagCapture := range tagNames {
					if strings.Contains(tagCapture.Text, "-") {
						elementKey := fmt.Sprintf("%s:%d:%d", tagCapture.Text, tagCapture.StartByte, tagCapture.EndByte)

						if _, exists := elementMap[elementKey]; exists {
							continue
						}

						element := types.CustomElementMatch{
							TagName:    tagCapture.Text,
							Range:      d.ByteRangeToProtocolRange(content, tagCapture.StartByte, tagCapture.EndByte),
							Attributes: make(map[string]types.AttributeMatch),
						}

						if attrNames, hasAttrs := captureMap["attr.name"]; hasAttrs {
							for _, attrCapture := range attrNames {
								attr := types.AttributeMatch{
									Name:  attrCapture.Text,
									Range: d.ByteRangeToProtocolRange(content, attrCapture.StartByte, attrCapture.EndByte),
								}

								if attrValues, hasValues := captureMap["attr.value"]; hasValues {
									for _, valueCapture := range attrValues {
										if valueCapture.StartByte > attrCapture.EndByte &&
											valueCapture.StartByte < attrCapture.EndByte+100 {
											attr.Value = valueCapture.Text
											break
										}
									}
								}

								element.Attributes[attr.Name] = attr
							}
						}

						elementMap[elementKey] = element
					}
				}
			}
		}
	}

	for _, element := range elementMap {
		elements = append(elements, element)
	}

	return elements, nil
}

// analyzeCompletionContext analyzes completion context for TSX documents
func (d *TSXDocument) analyzeCompletionContext(position protocol.Position, handler *Handler) *types.CompletionAnalysis {
	analysis := &types.CompletionAnalysis{
		Type: types.CompletionUnknown,
	}

	tree, content := d.TreeAndContent()
	if tree == nil {
		return analysis
	}

	byteOffset := d.PositionToByteOffset(position, content)

	completionMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "tsx", "completionContext")
	if err != nil {
		return analysis
	}
	defer completionMatcher.Close()

	var possibleMatches []struct {
		captureType string
		capture     Q.CaptureInfo
		captureMap  Q.CaptureMap
	}

	for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
		for captureName, captures := range captureMap {
			if captureName == "context" {
				continue
			}
			for _, capture := range captures {
				if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
					possibleMatches = append(possibleMatches, struct {
						captureType string
						capture     Q.CaptureInfo
						captureMap  Q.CaptureMap
					}{captureName, capture, captureMap})
				} else if captureName == "attr.name.completion" && capture.EndByte <= byteOffset && byteOffset-capture.EndByte <= 10 {
					possibleMatches = append(possibleMatches, struct {
						captureType string
						capture     Q.CaptureInfo
						captureMap  Q.CaptureMap
					}{captureName, capture, captureMap})
				}
			}
		}
	}

	if len(possibleMatches) == 0 {
		for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
			for captureName, captures := range captureMap {
				for _, capture := range captures {
					if captureName == "tag.name.completion" &&
						capture.EndByte < byteOffset &&
						byteOffset-capture.EndByte <= 2 {
						if strings.Contains(capture.Text, "-") {
							analysis.Type = types.CompletionAttributeName
							analysis.TagName = extractCustomElementFromText(capture.Text)
							return analysis
						}
					}
				}
			}
		}
		return analysis
	}

	var bestMatch *struct {
		captureType string
		capture     Q.CaptureInfo
		captureMap  Q.CaptureMap
	}

	for i, match := range possibleMatches {
		if bestMatch == nil {
			bestMatch = &possibleMatches[i]
			continue
		}

		currentScore := scoreMatch(match.captureType, match.captureMap, match.capture, byteOffset)
		bestScore := scoreMatch(bestMatch.captureType, bestMatch.captureMap, bestMatch.capture, byteOffset)

		if currentScore < bestScore {
			bestMatch = &possibleMatches[i]
		}
	}

	if bestMatch != nil {
		switch bestMatch.captureType {
		case "tag.name.completion":
			analysis.Type = types.CompletionTagName
			analysis.TagName = bestMatch.capture.Text
		case "attr.name.completion":
			analysis.Type = types.CompletionAttributeName
			analysis.AttributeName = bestMatch.capture.Text
			analysis.TagName = d.findTagNameForAttribute(byteOffset, bestMatch.captureMap)
		case "attr.value.completion":
			analysis.Type = types.CompletionAttributeValue
			tagName := d.findBestTagName(byteOffset, possibleMatches)
			attrName := d.findBestAttributeName(byteOffset, possibleMatches)
			analysis.TagName = tagName
			analysis.AttributeName = attrName
		case "interpolation":
			analysis.Type = types.CompletionUnknown
			return analysis
		}
	}

	return analysis
}

func (d *TSXDocument) findTagNameForAttribute(byteOffset uint, captureMap Q.CaptureMap) string {
	if tagNames, exists := captureMap["tag.name.completion"]; exists {
		var closestTag *Q.CaptureInfo
		var closestDistance = ^uint(0)

		for _, tagCapture := range tagNames {
			if (tagCapture.StartByte <= byteOffset && byteOffset <= tagCapture.EndByte) ||
				(tagCapture.EndByte <= byteOffset) {

				var distance uint
				if byteOffset <= tagCapture.EndByte {
					distance = 0
				} else {
					distance = byteOffset - tagCapture.EndByte
				}

				if distance < closestDistance {
					closestDistance = distance
					closestTag = &tagCapture
				}
			}
		}

		if closestTag != nil {
			tagName := extractCustomElementFromText(closestTag.Text)
			if tagName == "" {
				if strings.Contains(closestTag.Text, "-") && !strings.Contains(closestTag.Text, " ") {
					tagName = closestTag.Text
				}
			}
			return tagName
		}
	}

	return ""
}

func (d *TSXDocument) findBestTagName(byteOffset uint, possibleMatches []struct {
	captureType string
	capture     Q.CaptureInfo
	captureMap  Q.CaptureMap
}) string {
	var bestTagName string
	var bestDistance = ^uint(0)

	for _, match := range possibleMatches {
		if tagNames, exists := match.captureMap["tag.name.completion"]; exists {
			for _, tagCapture := range tagNames {
				var distance uint
				if tagCapture.StartByte <= byteOffset && byteOffset <= tagCapture.EndByte {
					distance = 0
				} else if tagCapture.EndByte <= byteOffset {
					distance = byteOffset - tagCapture.EndByte
				} else {
					continue
				}

				if distance < bestDistance {
					tagName := extractCustomElementFromText(tagCapture.Text)
					if tagName != "" {
						bestTagName = tagName
						bestDistance = distance
					}
				}
			}
		}
	}

	return bestTagName
}

func (d *TSXDocument) findBestAttributeName(byteOffset uint, possibleMatches []struct {
	captureType string
	capture     Q.CaptureInfo
	captureMap  Q.CaptureMap
}) string {
	var bestAttrName string
	var bestDistance = ^uint(0)

	for _, match := range possibleMatches {
		if attrNames, exists := match.captureMap["attr.name.completion"]; exists {
			for _, attrCapture := range attrNames {
				if attrCapture.EndByte <= byteOffset {
					distance := byteOffset - attrCapture.EndByte
					if distance < bestDistance {
						attrName := extractAttributeNameFromText(attrCapture.Text)
						if attrName != "" {
							bestAttrName = attrName
							bestDistance = distance
						}
					}
				}
			}
		}

		if match.captureType == "attr.name.completion" {
			if match.capture.EndByte <= byteOffset {
				distance := byteOffset - match.capture.EndByte
				if distance < bestDistance {
					attrName := extractAttributeNameFromText(match.capture.Text)
					if attrName != "" {
						bestAttrName = attrName
						bestDistance = distance
					}
				}
			}
		}
	}

	return bestAttrName
}

func extractCustomElementFromText(text string) string {
	parts := strings.Split(text, "<")
	if len(parts) < 2 {
		return ""
	}

	lastPart := parts[len(parts)-1]

	tagName := lastPart
	if spaceIdx := strings.Index(tagName, " "); spaceIdx != -1 {
		tagName = tagName[:spaceIdx]
	}
	if gtIdx := strings.Index(tagName, ">"); gtIdx != -1 {
		tagName = tagName[:gtIdx]
	}

	if strings.Contains(tagName, "-") {
		return tagName
	}

	return ""
}

func extractAttributeNameFromText(text string) string {
	if !strings.Contains(text, " ") && !strings.Contains(text, "<") && !strings.Contains(text, "=") {
		return text
	}

	words := strings.Fields(text)
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		word = strings.TrimSuffix(word, "=")
		if len(word) > 0 && !strings.Contains(word, "<") && !strings.Contains(word, ">") {
			return word
		}
	}

	return ""
}

func scoreMatch(captureType string, captureMap Q.CaptureMap, capture Q.CaptureInfo, byteOffset uint) int {
	score := 0

	switch captureType {
	case "tag.name.completion":
		score += SCORE_TAG_NAME_BASE
	case "attr.name.completion":
		score += SCORE_ATTR_NAME_BASE
	case "attr.value.completion":
		score += SCORE_ATTR_VALUE_BASE
	}

	if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
		score -= SCORE_CURSOR_BONUS
	}

	if len(captureMap) > 2 {
		score += len(captureMap) * SCORE_BROAD_CAPTURE_PENALTY
	}

	captureLength := capture.EndByte - capture.StartByte
	if captureLength > MAX_REASONABLE_TAG_LENGTH {
		score += int(captureLength)
	}

	return score
}
